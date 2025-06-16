package dealtemplate

import (
	"encoding/json"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:     "create",
	Usage:    "Create a new deal template",
	Category: "Deal Template Management",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Name of the deal template",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "Description of the deal template",
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-gb",
			Usage: "Price in FIL per GiB for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-gb-epoch",
			Usage: "Price in FIL per GiB per epoch for storage deals",
			Value: 0.0,
		},
		&cli.Float64Flag{
			Name:  "deal-price-per-deal",
			Usage: "Price in FIL per deal for storage deals",
			Value: 0.0,
		},
		&cli.DurationFlag{
			Name:  "deal-duration",
			Usage: "Duration for storage deals (e.g., 535 days)",
			Value: 0,
		},
		&cli.DurationFlag{
			Name:  "deal-start-delay",
			Usage: "Start delay for storage deals (e.g., 72h)",
			Value: 0,
		},
		&cli.BoolFlag{
			Name:  "deal-verified",
			Usage: "Whether deals should be verified",
		},
		&cli.BoolFlag{
			Name:  "deal-keep-unsealed",
			Usage: "Whether to keep unsealed copy of deals",
		},
		&cli.BoolFlag{
			Name:  "deal-announce-to-ipni",
			Usage: "Whether to announce deals to IPNI",
		},
		&cli.StringFlag{
			Name:  "deal-provider",
			Usage: "Storage Provider ID for deals (e.g., f01000)",
		},
		&cli.StringFlag{
			Name:  "deal-url-template",
			Usage: "URL template for deals",
		},
		&cli.StringFlag{
			Name:  "deal-http-headers",
			Usage: "HTTP headers for deals in JSON format",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		db = db.WithContext(c.Context)

		// Parse deal HTTP headers if provided
		var dealHTTPHeaders model.ConfigMap
		if headersStr := c.String("deal-http-headers"); headersStr != "" {
			var tempMap map[string]string
			if err := json.Unmarshal([]byte(headersStr), &tempMap); err != nil {
				return errors.Wrapf(err, "invalid JSON format for deal-http-headers: %s", headersStr)
			}
			dealHTTPHeaders = model.ConfigMap(tempMap)
		}

		template, err := dealtemplate.Default.CreateHandler(c.Context, db, dealtemplate.CreateRequest{
			Name:                c.String("name"),
			Description:         c.String("description"),
			DealPricePerGB:      c.Float64("deal-price-per-gb"),
			DealPricePerGBEpoch: c.Float64("deal-price-per-gb-epoch"),
			DealPricePerDeal:    c.Float64("deal-price-per-deal"),
			DealDuration:        c.Duration("deal-duration"),
			DealStartDelay:      c.Duration("deal-start-delay"),
			DealVerified:        c.Bool("deal-verified"),
			DealKeepUnsealed:    c.Bool("deal-keep-unsealed"),
			DealAnnounceToIPNI:  c.Bool("deal-announce-to-ipni"),
			DealProvider:        c.String("deal-provider"),
			DealURLTemplate:     c.String("deal-url-template"),
			DealHTTPHeaders:     dealHTTPHeaders,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, *template)
		return nil
	},
}
