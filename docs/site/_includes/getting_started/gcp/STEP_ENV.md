You need to create a service account so that Deckhouse Platform can manage resources in the {{ page.platform_name[page.lang] }}. The detailed instructions for creating a service account are available in the [documentation](/en/documentation/v1/modules/030-cloud-provider-gcp/environment.html). Below is a brief sequence of required actions:

> List of roles required:
> - `roles/compute.admin`
> - `roles/iam.serviceAccountUser`
> - `roles/networkmanagement.admin`

Export environment variables:
{% snippetcut %}
```shell
export PROJECT=sandbox
export SERVICE_ACCOUNT_NAME=deckhouse
```
{% endsnippetcut %}

Select a project:
{% snippetcut %}
```shell
gcloud config set project $PROJECT
```
{% endsnippetcut %}

Create a service account:
{% snippetcut %}
```shell
gcloud iam service-accounts create $SERVICE_ACCOUNT_NAME
```
{% endsnippetcut %}

Verify service account roles:
{% snippetcut %}
```shell
gcloud projects get-iam-policy ${PROJECT} --flatten="bindings[].members" --format='table(bindings.role)' \
    --filter="bindings.members:${SERVICE_ACCOUNT_NAME}@${PROJECT}.iam.gserviceaccount.com"
```
{% endsnippetcut %}

Create a service account key:
{% snippetcut %}
```shell
gcloud iam service-accounts keys create --iam-account ${SERVICE_ACCOUNT_NAME}@${PROJECT}.iam.gserviceaccount.com \
    ~/service-account-key-${PROJECT}-${SERVICE_ACCOUNT_NAME}.json
```
{% endsnippetcut %}
