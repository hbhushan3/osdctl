## osdctl cluster context

Shows the context of a specified cluster

```
osdctl cluster context [flags]
```

### Options

```
  -C, --cluster-id string           Cluster ID
  -d, --days int                    Command will display X days of Error SLs sent to the cluster. Days is set to 30 by default (default 30)
      --full                        Run full suite of checks.
  -h, --help                        help for context
      --jiratoken jira_token        Pass in the Jira access token directly. If not passed in, by default will read jira_token from ~/.config/osdctl.
                                    Jira access tokens can be registered by visiting https://issues.redhat.com/secure/ViewProfile.jspa?selectedTab=com.atlassian.pats.pats-plugin:jira-user-personal-access-tokens
      --oauthtoken pd_oauth_token   Pass in PD oauthtoken directly. If not passed in, by default will read pd_oauth_token from ~/.config/osdctl.
                                    PD OAuth tokens can be generated by visiting https://martindstone.github.io/PDOAuth/
      --pages int                   Command will display X pages of Cloud Trail logs for the cluster. Pages is set to 40 by default (default 40)
  -p, --profile string              AWS Profile
      --usertoken pd_user_token     Pass in PD usertoken directly. If not passed in, by default will read pd_user_token from ~/config/osdctl
      --verbose                     Verbose output
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --as string                        Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --cluster string                   The name of the kubeconfig cluster to use
      --context string                   The name of the kubeconfig context to use
      --insecure-skip-tls-verify         If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string                Path to the kubeconfig file to use for CLI requests.
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
  -o, --output string                    Valid formats are ['', 'json', 'yaml', 'env']
      --request-timeout string           The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                    The address and port of the Kubernetes API server
  -S, --skip-version-check               skip checking to see if this is the most recent release
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [osdctl cluster](osdctl_cluster.md)	 - Provides information for a specified cluster

