# 存储

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage" method="get" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/acd" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/azureblob" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/b2" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/box" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/drive" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/dropbox" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/fichier" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/filefabric" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/storage/ftp" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

...

{'info': {'version': '2.0', 'title': 'Storage'}, 'tags': [{'name': 'storage', 'description': 'Storage operations'}], 'paths': {'/storage': {'get': {'tags': ['storage'], 'summary': 'List all storage backends', 'operationId': 'listStorageBackends', 'produces': ['application/json'], 'responses': {'200': {'description': 'OK', 'schema': {'$ref': '#/definitions/StorageBackendList'}}, 'default': {'description': 'unexpected error', 'schema': {'$ref': '#/definitions/error'}}}}}, '/storage/...