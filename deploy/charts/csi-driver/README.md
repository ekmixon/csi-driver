# cert-manager-csi-driver

![Version: v0.2.1](https://img.shields.io/badge/Version-v0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v0.2.0](https://img.shields.io/badge/AppVersion-v0.2.0-informational?style=flat-square)

A Helm chart for cert-manager-csi-driver

**Homepage:** <https://github.com/cert-manager/csi-driver>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| joshvanl | joshua.vanleeuwen@jetstack.io | https://cert-manager.io |

## Source Code

* <https://github.com/cert-manager/csi-driver>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| app.driver | object | `{"csiDataDir":"/tmp/cert-manager-csi-driver","name":"csi.cert-manager.io","useTokenRequest":false}` | Options for CSI driver |
| app.driver.csiDataDir | string | `"/tmp/cert-manager-csi-driver"` | Configures the hostPath directory that the driver will write and mount volumes from. |
| app.driver.name | string | `"csi.cert-manager.io"` | Name of the driver which will be registered with Kubernetes. |
| app.driver.useTokenRequest | bool | `false` | If enabled, will use CSI token request for creating CertificateRequests. CertificateRequests will be created via mounting pod's service accounts. |
| app.livenessProbe | object | `{"port":9809}` | Options for the liveness container. |
| app.livenessProbe.port | int | `9809` | The port that will expose the livness of the csi-driver |
| app.logLevel | int | `1` | Verbosity of cert-manager-csi-driver logging. |
| image.pullPolicy | string | `"IfNotPresent"` | Kubernetes imagePullPolicy on csi-driver. |
| image.repository | string | `"quay.io/jetstack/cert-manager-csi-driver"` | Target image repository. |
| image.tag | string | `"v0.2.0"` | Target image version tag. |
| livenessProbeImage.pullPolicy | string | `"IfNotPresent"` | Kubernetes imagePullPolicy on liveness probe. |
| livenessProbeImage.repository | string | `"k8s.gcr.io/sig-storage/livenessprobe"` | Target image repository. |
| livenessProbeImage.tag | string | `"v2.4.0"` | Target image version tag. |
| nodeDriverRegistrarImage.pullPolicy | string | `"IfNotPresent"` | Kubernetes imagePullPolicy on node-driver. |
| nodeDriverRegistrarImage.repository | string | `"k8s.gcr.io/sig-storage/csi-node-driver-registrar"` | Target image repository. |
| nodeDriverRegistrarImage.tag | string | `"v2.3.0"` | Target image version tag. |
| resources | object | `{}` |  |
| tolerations | list | `[]` |  |

