# singularity deal-template create

Create a new deal template with reusable deal parameters.

## Usage

```bash
singularity deal-template create [flags]
```

## Required Flags

- `--name` - Unique name for the deal template

## Optional Flags

- `--description` - Human-readable description of the template
- `--deal-price-per-gb` - Price in FIL per GiB for storage deals (default: 0.0)
- `--deal-price-per-gb-epoch` - Price in FIL per GiB per epoch for storage deals (default: 0.0)
- `--deal-price-per-deal` - Price in FIL per deal for storage deals (default: 0.0)
- `--deal-duration` - Duration for storage deals (e.g., 535days, 1y, 8760h)
- `--deal-start-delay` - Start delay for storage deals (e.g., 72h, 3days)
- `--deal-verified` - Whether deals should be verified (datacap deals)
- `--deal-keep-unsealed` - Whether to keep unsealed copy of deals
- `--deal-announce-to-ipni` - Whether to announce deals to IPNI
- `--deal-provider` - Storage Provider ID for deals (e.g., f01000)
- `--deal-url-template` - URL template for deals
- `--deal-http-headers` - HTTP headers for deals in JSON format

## Examples

### Basic Template
```bash
singularity deal-template create \
  --name "basic-archive" \
  --description "Basic archival storage" \
  --deal-price-per-gb 0.0000000001 \
  --deal-duration 535days \
  --deal-verified
```

### Enterprise Template
```bash
singularity deal-template create \
  --name "enterprise-tier" \
  --description "Enterprise-grade storage with 3-year retention" \
  --deal-duration 1095days \
  --deal-price-per-gb 0.0000000002 \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni \
  --deal-start-delay 72h \
  --deal-provider f01000
```

### With Custom Headers
```bash
singularity deal-template create \
  --name "authenticated-storage" \
  --deal-http-headers '{"Authorization":"Bearer token123","X-Custom":"value"}' \
  --deal-url-template "https://api.example.com/piece/{PIECE_CID}" \
  --deal-duration 365days
```

## See Also

- [singularity deal-template list](list.md) - List all templates
- [singularity prep create](../prep/create.md) - Use templates in preparations
- [Deal Templates Guide](../../deal-templates.md) - Complete guide to deal templates