# Deal Template Commands

Deal template commands allow you to create, manage, and use reusable deal configurations for data preparation workflows.

## Available Commands

* [create](create.md) - Create a new deal template
* [list](list.md) - List all deal templates  
* [get](get.md) - Get details of a specific deal template
* [delete](delete.md) - Delete a deal template

## Quick Examples

```bash
# Create a template
singularity deal-template create --name "standard" --deal-price-per-gb 0.0000000001 --deal-duration 535days

# List templates
singularity deal-template list

# Use template in preparation
singularity prep create --source /data --deal-template standard --auto-create-deals
```

For detailed usage and examples, see the [Deal Templates guide](../../deal-templates.md).