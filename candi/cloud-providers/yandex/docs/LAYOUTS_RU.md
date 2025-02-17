---
title: "Cloud provider — Yandex.Cloud: схемы размещения"
---

## Standard

В данной схеме размещения узлы не будут иметь публичных адресов, а будут выходить в интернет через Yandex.Cloud NAT.

> ⚠️ Внимание! На момент конца 2020 года функция Yandex.Cloud NAT находится на стадии Preview. Для того, чтобы появилась возможность включения Cloud NAT в вашем облаке, необходимо заранее (лучше за неделю) обратиться в поддержку Yandex.Cloud и запросить у них доступ.

![resources](https://docs.google.com/drawings/d/e/2PACX-1vTSpvzjcEBpD1qad9u_UgvsOrYT_Xtnxwg6Pzb64HQHLqQWcZi6hhCNRPKVUdYKX32nXEVJeCzACVRG/pub?w=812&h=655)
<!--- Исходник: https://docs.google.com/drawings/d/1WI8tu-QZYcz3DvYBNlZG4s5OKQ9JKyna7ESHjnjuCVQ/edit --->

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
sshPublicKey: "ssh-rsa <SSH_PUBLIC_KEY>"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```

### Включение Cloud NAT

**Внимание!** Сразу же (в течение 3х минут) после создания базовых сетевых ресурсов, для всех подсетей необходимо включить Cloud NAT. Вручную через web-интерфейс. Если этого не сделать, то bootstrap процесс не сможет завершиться.

![Включение NAT](../../images/030-cloud-provider-yandex/enable_cloud_nat.png)

## WithoutNAT

В данной схеме размещения NAT (любого вида) не используется, а каждому узлу выдаётся публичный IP.

**Внимание!** В модуле cloud-provider-yandex пока нет поддержки Security Groups, поэтому все узлы кластера будут смотреть наружу.

![resources](https://docs.google.com/drawings/d/e/2PACX-1vTgwXWsNX6CKCRaMf5t6rl3kpKQQFHK6T8Dsg1jAwAwYaN1MRbxKFsSFQHeo1N3Qec4etPpeA0guB6-/pub?w=812&h=655)
<!--- Исходник: https://docs.google.com/drawings/d/1I7M9DquzLNu-aTjqLx1_6ZexPckL__-501Mt393W1fw/edit --->

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
sshPublicKey: "ssh-rsa <SSH_PUBLIC_KEY>"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```

## WithNATInstance

В данной схеме размещения создаётся NAT instance, а в таблицу маршрутизации добавляется правило на 0.0.0.0/0 с NAT instance nexthop'ом.

Если задан `withNATInstance.externalSubnetID` — NAT instance создастся в зоне этого subnet.
Если `withNATInstance.externalSubnetID` не задан, а `withNATInstance.internalSubnetID` задан — NAT instance создастся в зоне этого subnet.
Если ни `withNATInstance.externalSubnetID`, ни `withNATInstance.internalSubnetID` не заданы — NAT instance создастся в зоне `ru-central1-c`.

**Внимание!** Сразу же (в течение 3х минут) после создания базовых сетевых ресурсов, нужно вручную прописать маршрут к созданному NAT instance. Если этого не сделать, то bootstrap процесс не сможет завершиться.

```text
$ yc compute instance list | grep nat
| ef378c62hvqi075cp57j | kube-yc-nat | ru-central1-c | RUNNING | 130.193.44.28   | 192.168.178.22 |

$ yc vpc route-table update --name kube-yc --route destination=0.0.0.0/0,next-hop=192.168.178.22
```

![resources](https://docs.google.com/drawings/d/e/2PACX-1vSnNqebgRdwGP8lhKMJfrn5c0QXDpe9YdmIlK4eDberysLLgYiKNuwaPLHcyQhJigvQ21SANH89uipE/pub?w=812&h=655)
<!--- Исходник: https://docs.google.com/drawings/d/1oVpZ_ldcuNxPnGCkx0dRtcAdL7BSEEvmsvbG8Aif1pE/edit --->

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
sshPublicKey: "ssh-rsa <SSH_PUBLIC_KEY>"
nodeNetworkCIDR: 192.168.12.13/24
existingNetworkID: tewt243tewsdf
dhcpOptions:
  domainName: test.local
  domainNameServers:
  - 213.177.96.1
  - 231.177.97.1
```
