{% assign revision=include.revision %}

To install the **Deckhouse Platform**, we will use a prebuilt Docker image. It is necessary to transfer configuration files to the container, as well as ssh-keys for access to the master nodes:

{%- if revision == 'ee' %}
{% snippetcut selector="docker-login" %}
```shell
docker login -u license-token -p <LICENSE_TOKEN> registry.deckhouse.io
docker run -it -v "$PWD/config.yml:/config.yml" -v "$HOME/.ssh/:/tmp/.ssh/" \
{% if page.platform_type == "existing" %} -v "$PWD/kubeconfig:/kubeconfig" \
{% endif %}{% if page.platform_type == "cloud" %} -v "$PWD/resources.yml:/resources.yml" -v "$PWD/dhctl-tmp:/tmp/dhctl" {% endif %} registry.deckhouse.io/deckhouse/ee/install:stable bash
```
{% endsnippetcut %}
{% else %}
{% snippetcut %}
```shell
docker run -it -v "$PWD/config.yml:/config.yml" -v "$HOME/.ssh/:/tmp/.ssh/" \
{% if page.platform_type == "existing" %} -v "$PWD/kubeconfig:/kubeconfig" \
{% endif %}{% if page.platform_type == "cloud" %} -v "$PWD/resources.yml:/resources.yml" -v "$PWD/dhctl-tmp:/tmp/dhctl" {% endif %} registry.deckhouse.io/deckhouse/ce/install:stable bash
```
{% endsnippetcut %}
{% endif %}

{%- if page.platform_type == "existing" %}
Notes:
- Kubeconfig with access to Kubernetes API must be used in kubeconfig mount.
{% endif %}

Now, to initiate the process of installation, you need to execute inside the container:

{% snippetcut %}
```shell
{%- if page.platform_type == "existing" %}
dhctl bootstrap-phase install-deckhouse \
  --kubeconfig=/kubeconfig \
  --config=/config.yml
{%- elsif page.platform_type == "baremetal" %}
dhctl bootstrap \
  --ssh-user=<username> \
  --ssh-host=<master_ip> \
  --ssh-agent-private-keys=/tmp/.ssh/id_rsa \
  --config=/config.yml
{%- elsif page.platform_type == "cloud" %}
dhctl bootstrap \
  --ssh-user=<username> \
  --ssh-agent-private-keys=/tmp/.ssh/id_rsa \
  --config=/config.yml \
  --resources=/resources.yml
{%- endif %}
```
{% endsnippetcut %}

{%- if page.platform_type == "baremetal" or page.platform_type == "cloud" %}
{%- if page.platform_type == "baremetal" %}
`username` variable here refers to the user that generated the SSH key.
{%- else %}
`username` variable here refers to
{%- if page.platform_code == "openstack" %} the default user for the relevant VM image.
{%- elsif page.platform_code == "azure" %} `azureuser` (for the images suggested in this documentation).
{%- elsif page.platform_code == "gcp" %} `user` (for the images suggested in this documentation).
{%- else %} `ubuntu` (for the images suggested in this documentation).
{%- endif %}
{%- endif %}

{%- if page.platform_type == "cloud" %}
Notes:
<ul>
<li>
<div markdown="1">
The `-v "$PWD/dhctl-tmp:/tmp/dhctl"` parameter enables saving the state of the Terraform installer to a temporary directory on the startup host. It allows the installation to continue correctly in case of a failure of the installer's container.
</div>
</li>
<li><p>If any problems {% if page.platform_type="cloud" %}on the cloud provider side {% endif %}occur, you can cancel the process of installation and remove all created objects using the following command (the configuration file should be the same you’ve used to initiate the installation):</p>
<div markdown="0">
{% snippetcut %}
```shell
dhctl bootstrap-phase abort --config=/config.yml
```
{% endsnippetcut %}
</div></li>
</ul>
{%- endif %}
{%- endif %}

After the installation is complete, you will be returned to the command line.

Almost everything is ready for a fully-fledged Deckhouse Platform to work!
