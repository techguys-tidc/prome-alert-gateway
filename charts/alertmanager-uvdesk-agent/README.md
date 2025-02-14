- [Helm Charts](#helm-charts)
  - [Helm Template](#helm-template)
  - [Helm Install](#helm-install)

# Helm Charts

## Helm Template

```shell
{
helm template --namespace gong --create-namespace my-release alertmanager-uvdesk-agent -f alertmanager-uvdesk-agent/values.yaml
}
```

## Helm Install

```shell
{
helm upgrade --install --namespace gong --create-namespace my-release alertmanager-uvdesk-agent -f alertmanager-uvdesk-agent/values.yaml

}
```