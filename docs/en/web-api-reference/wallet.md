# Wallet

## Overview

The Wallet API provides comprehensive wallet management capabilities including creation, import, and metadata management. All wallet operations support optional metadata fields for enhanced organization:

- **name**: Human-readable display name for the wallet
- **contact**: Contact information (email, etc.) 
- **location**: Geographic location or region identifier

## Metadata Support

### Request Fields
When creating or importing wallets, you can include metadata:
```json
{
  "privateKey": "...",
  "name": "My Storage Wallet",
  "contact": "admin@example.com", 
  "location": "US-East"
}
```

### Response Fields
Wallet responses include the metadata in the following fields:
```json
{
  "id": 1,
  "actorId": "f01234",
  "actorName": "My Storage Wallet",    // from 'name' field
  "address": "f1abc...", 
  "contactInfo": "admin@example.com",  // from 'contact' field
  "location": "US-East",               // from 'location' field
  "walletType": "UserWallet"
}
```

## API Endpoints

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet" method="get" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet/create" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet/{address}" method="delete" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet/{address}/init" method="post" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

{% swagger src="https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml" path="/wallet/{address}/update" method="patch" %}
[https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)
{% endswagger %}

