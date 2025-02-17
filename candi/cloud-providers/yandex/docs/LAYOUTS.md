---
title: "Cloud provider — Yandex.Cloud: Layouts"
---

## Standard

In this placement strategy, nodes do not have public IP addresses allocated to them; they use Yandex.Cloud NAT to connect to the Internet.

> ⚠️ Caution! The Yandex.Cloud NAT feature is at the [Preview stage](https://cloud.yandex.com/en/docs/vpc/operations/enable-nat) as of July 2021.  To enable the Cloud NAT feature for your cloud, you need to contact Yandex.Cloud support in advance (in a week or so) and request access to it.

![resources](https://docs.google.com/drawings/d/e/2PACX-1vTSpvzjcEBpD1qad9u_UgvsOrYT_Xtnxwg6Pzb64HQHLqQWcZi6hhCNRPKVUdYKX32nXEVJeCzACVRG/pub?w=812&h=655)
<!--- Source: https://docs.google.com/drawings/d/1WI8tu-QZYcz3DvYBNlZG4s5OKQ9JKyna7ESHjnjuCVQ/edit --->

```yaml
apiVersion: deckhouse.io/v1
kind: YandexClusterConfiguration
layout: Standard
provider:
  cloudID: dsafsafewf
  folderID: enh1233214367
  serviceAccountJSON: |
    {"test": "test"}
masterNodeGroup:
  replicas: 1
  zones:
  - ru-central1-a
  - ru-central1-b
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    additionalLabels:
      takes: priority
nodeGroups:
- name: khm
  replicas: 1
  zones:
  - ru-central1-a
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    coreFraction: 50
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    additionalLabels:
      toy: example
labels:
  billing: prod
sshPublicKey: "ssh-rsa ewasfef3wqefwefqf43qgqwfsd"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```

### Enabling Cloud NAT

**Caution!** Note that you must manually (using the web interface) enable Cloud NAT within 3 minutes after creating the primary network resources. The bootstrap process won't complete if you fail to do this.

![Enabling NAT](../../images/030-cloud-provider-yandex/enable_cloud_nat.png)

## WithoutNAT

In this layout, NAT (of any kind) is not used, and each node is assigned a public IP.

**Caution!** Currently, the cloud-provider-yandex module does not support Security Groups; thus, is why all cluster nodes connect directly to the Internet.

![resources](https://docs.google.com/drawings/d/e/2PACX-1vTgwXWsNX6CKCRaMf5t6rl3kpKQQFHK6T8Dsg1jAwAwYaN1MRbxKFsSFQHeo1N3Qec4etPpeA0guB6-/pub?w=812&h=655)
<!--- Source: https://docs.google.com/drawings/d/1I7M9DquzLNu-aTjqLx1_6ZexPckL__-501Mt393W1fw/edit --->

```yaml
apiVersion: deckhouse.io/v1
kind: YandexClusterConfiguration
layout: WithoutNAT
provider:
  cloudID: dsafsafewf
  folderID: enh1233214367
  serviceAccountJSON: |
    {"test": "test"}
masterNodeGroup:
  replicas: 1
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    zones:
    - ru-central1-a
    - ru-central1-b
nodeGroups:
- name: khm
  replicas: 1
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    coreFraction: 50
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    zones:
    - ru-central1-a
sshPublicKey: "ssh-rsa ewasfef3wqefwefqf43qgqwfsd"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```

## WithNATInstance

In this placement strategy, Deckhouse creates a NAT instance and adds a rule to a route table containing a route to 0.0.0.0/0 with a NAT instance as the next hop.

If the `withNATInstance.externalSubnetID` parameter is set, the NAT instance will be created in this subnet.
IF the `withNATInstance.externalSubnetID` parameter is not set and `withNATInstance.internalSubnetID` is set, the NAT instance will be created in this last subnet.
If neither `withNATInstance.externalSubnetID` nor `withNATInstance.internalSubnetID` is set, the NAT instance will be created in the  `ru-central1-c` zone.

**Caution!** Note that you must manually enter the route to the created NAT instance within 3 minutes after creating the primary network resources. The bootstrap process won't complete if you fail to do this.

```text
$ yc compute instance list | grep nat
| ef378c62hvqi075cp57j | kube-yc-nat | ru-central1-c | RUNNING | 130.193.44.28   | 192.168.178.22 |

$ yc vpc route-table update --name kube-yc --route destination=0.0.0.0/0,next-hop=192.168.178.22
```

![resources](https://docs.google.com/drawings/d/e/2PACX-1vSnNqebgRdwGP8lhKMJfrn5c0QXDpe9YdmIlK4eDberysLLgYiKNuwaPLHcyQhJigvQ21SANH89uipE/pub?w=812&h=655)
<!--- Source: https://docs.google.com/drawings/d/1oVpZ_ldcuNxPnGCkx0dRtcAdL7BSEEvmsvbG8Aif1pE/edit --->

```yaml
apiVersion: deckhouse.io/v1
kind: YandexClusterConfiguration
layout: WithNATInstance
withNATInstance:
  natInstanceExternalAddress: 30.11.34.45
  internalSubnetID: sjfwefasjdfadsfj
  externalSubnetID: etasjflsjdfiorej
provider:
  cloudID: dsafsafewf
  folderID: enh1233214367
  serviceAccountJSON: |
    {"test": "test"}
masterNodeGroup:
  replicas: 1
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    zones:
    - ru-central1-a
    - ru-central1-b
nodeGroups:
- name: khm
  replicas: 1
  instanceClass:
    cores: 4
    memory: 8192
    imageID: testtest
    coreFraction: 50
    externalIPAddresses:
    - "198.51.100.5"
    - "Auto"
    externalSubnetID: tewt243tewsdf
    zones:
    - ru-central1-a
sshPublicKey: "ssh-rsa ewasfef3wqefwefqf43qgqwfsd"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```
