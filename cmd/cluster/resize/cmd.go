package resize

import (
	"fmt"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift/osdctl/pkg/k8s"
	"github.com/openshift/osdctl/pkg/utils"
	"github.com/spf13/cobra"

	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Resize struct {
	client    client.Client
	hive      client.Client
	hiveAdmin client.Client

	cluster      *cmv1.Cluster
	clusterId    string
	instanceType string
}

func NewCmdResize() *cobra.Command {
	resize := &cobra.Command{
		Use:  "resize",
		Args: cobra.NoArgs,
	}

	resize.AddCommand(
		newCmdResizeInfra(),
	)

	return resize
}

func (r *Resize) New() error {
	scheme := runtime.NewScheme()

	// Register machinev1beta1 for Machines
	fmt.Println("machine register")
	if err := machinev1beta1.Install(scheme); err != nil {
		return err
	}

	// Register hivev1 for MachinePools
	fmt.Println("register hivev1 for machinepool")
	if err := hivev1.AddToScheme(scheme); err != nil {
		return err
	}

	fmt.Println("register corev1 to pool")
	if err := corev1.AddToScheme(scheme); err != nil {
		return err
	}

	fmt.Println("create ocm connection")
	ocmClient, err := utils.CreateConnection()
	if err != nil {
		return err
	}
	fmt.Println("defer close and get cluster status")
	defer ocmClient.Close()
	cluster, err := utils.GetClusterAnyStatus(ocmClient, r.clusterId)
	if err != nil {
		return fmt.Errorf("failed to get OCM cluster info for %s: %s", r.clusterId, err)
	}
	r.cluster = cluster
	r.clusterId = cluster.ID()

	fmt.Println("get hive cluster")
	hive, err := utils.GetHiveCluster(cluster.ID())
	if err != nil {
		return err
	}

	fmt.Printf("new k8s client %s", r.clusterId)
	c, err := k8s.New(cluster.ID(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	fmt.Println("New hive client")
	hc, err := k8s.New(hive.ID(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	fmt.Println("ocmb admin")
	hac, err := k8s.NewAsBackplaneClusterAdmin(hive.ID(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	fmt.Println("assign variables")
	r.clusterId = cluster.ID()
	r.client = c
	r.hive = hc
	r.hiveAdmin = hac

	fmt.Println("Done with cmd.go for infra")

	return nil
}
