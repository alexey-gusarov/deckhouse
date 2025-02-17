{{- define "image-availability-alerts-by-mode" }}
{{- $controllerKind := . }}

- alert: {{ $controllerKind }}ImageAbsent
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_absent == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Следует проверить наличие образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }}{{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.
    summary: Образ `{{`{{ $labels.image }}`}}` отсутствует в registry.

- alert: {{ $controllerKind }}BadImageFormat
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_bad_image_format == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Следует формат имени образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }}{{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.
    summary: Некорректный формат имени образа `{{`{{ $labels.image }}`}}`.

- alert: {{ $controllerKind }}RegistryUnavailable
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_registry_unavailable == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Container registry недоступен для образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }}{{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.
    summary: Container registry недоступен для образа `{{`{{ $labels.image }}`}}`.

- alert: {{ $controllerKind }}AuthenticationFailure
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_authentication_failure == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Невозможно аутентифицироваться в container registry с указанными `imagePullSecrets` для образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }}{{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.
    summary: Невозможно аутентифицироваться в container registry с указанными `imagePullSecrets` для образа `{{`{{ $labels.image }}`}}`.

- alert: {{ $controllerKind }}AuthorizationFailure
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_authorization_failure == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Не хватает прав для загрузки с указанными `imagePullSecrets` для образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }}{{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.
    summary: Не хватает прав для загрузки с указанными `imagePullSecrets` для образа `{{`{{ $labels.image }}`}}`.

- alert: {{ $controllerKind }}UnknownError
  expr: |
    max by (namespace, {{ $controllerKind | lower }}, container, image) (
      k8s_image_availability_exporter_{{ $controllerKind | lower }}_unknown_error == 1
    )
    * on (namespace) group_left()
    max by (namespace) (extended_monitoring_enabled)
  labels:
    severity_level: "7"
    d8_module: extended-monitoring
    d8_component: image-availability-exporter
  annotations:
    plk_protocol_version: "1"
    plk_markup_format: "markdown"
    plk_pending_until_firing_for: "5m"
    plk_grouped_by__main: "UnavailableImagesInNamespace,namespace={{`{{ $labels.namespace }}`}},prometheus=deckhouse"
    description: >
      Произошла неизвестная ошибка для образа `{{`{{ $labels.image }}`}}`
      в Namespace `{{`{{ $labels.namespace }}`}}`
      в {{ $controllerKind }} {{`{{ $labels.`}}{{ $controllerKind | lower }} {{` }}`}}`
      в контейнере `{{`{{ $labels.container }}`}}` в registry.

      Подробнее в логах экспортера: `kubectl -n d8-monitoring logs -l app=image-availability-exporter -c image-availability-exporter`
    summary: Произошла неизвестная ошибка для образа `{{`{{ $labels.image }}`}}`.
{{- end }}

- name: d8.extended-monitoring.image-availability-exporter.image-checks
  rules:

{{- range list "Deployment" "StatefulSet" "DaemonSet" "CronJob" }}
{{- include "image-availability-alerts-by-mode" . | indent 2 }}
{{- end }}

  - alert: UnavailableImagesInNamespace
    expr: (count by (namespace) (ALERTS{alertname=~".+ImageAbsent|.+BadImageFormat|.+RegistryUnavailable|.+AuthenticationFailure|.+AuthorizationFailure|.+UnknownError", alertstate="firing"})) > 0
    labels:
      d8_module: extended-monitoring
      d8_component: image-availability-exporter
    annotations:
      plk_protocol_version: "1"
      plk_markup_format: "markdown"
      plk_alert_type: "group"
      plk_create_group_if_not_exists__main: UnavailableImagesInCluster,prometheus=deckhouse,tier=~tier
      plk_grouped_by__main: "UnavailableImagesInCluster,prometheus=deckhouse,tier=~tier"
      summary: В Namespace `{{`{{ $labels.namespace }}`}}` наличествует отсутствие образов в container registry.
      description: Подробнее в связанных алертах.
