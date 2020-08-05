# Override Parameters

When provisioning an instance of a service Minibroker simply forwards user-
provided provisioning parameters to helm. This can be a problem in untrusted
environments, e.g. when malicious users try to configure a very high number of
replicas to take down the cluster.

To mitigate that risk override parameters can be defined per service that will
take precedence over user-provided parameters. The user-provided parameters are
discarded alltogether in this case.

Example:

```yaml
provisioning:
  rabbitmq:
    replicas: 1
    ingress:
      enabled: false
```

Override params are defined per service, it is not possible to have a different
set of parameters per plan. As a result override parameters need to be chosen
wisely to work with all plans.

**Where do I find the helm charts Minibroker uses?**

Unless configured differently Minibroker pulls helm charts from https://github.com/helm/charts/tree/master/stable
which also is the source for https://hub.helm.sh/. The hub can be used to browse
the charts and list their versions .

**How do I know which plans are enabled and which chart version belongs to them?**

Minibroker scans the chart repository for all available app versions and
creates one plan for each of them, always pointing to the respective highest
chart version. All of these plans are available to the users.

If using Minibroker within CloudFoundry only the plans explicitly enabled by the
admin are available.

**How can I find the available configuration values for a given plan?**

Plans are named after the according chart version, e.g. plan "3-8-2" corresponds
to chart version 3.8.1.
The easiest way to find the options for a given chart version is by finding the
chart on the hub and selecting the version from the sidebar.
Please note that many charts were moved to the [bitnami repository](https://hub.helm.sh/charts/bitnami)
due to the upcoming [stable deprecation](https://github.com/helm/charts#deprecation-timeline).
