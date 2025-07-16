package wallet

import (
   "github.com/cockroachdb/errors"
   "github.com/data-preservation-programs/singularity/cmd/cliutil"
   "github.com/data-preservation-programs/singularity/database"
   "github.com/data-preservation-programs/singularity/handler/wallet"
   "github.com/data-preservation-programs/singularity/util"
   "github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
   Name:  "list",
   Usage: "List all imported wallets",
   Flags: []cli.Flag{
	   &cli.BoolFlag{
		   Name:    "with-balance",
		   Usage:   "Fetch and display live wallet balances from Lotus",
		   Value:   false,
	   },
	   &cli.StringFlag{
		   Name:    "lotus-api",
		   Usage:   "Lotus JSON-RPC API endpoint (required for --with-balance)",
		   EnvVars: []string{"LOTUS_API"},
	   },
	   &cli.StringFlag{
		   Name:    "lotus-token",
		   Usage:   "Lotus API authorization token (required for --with-balance)",
		   EnvVars: []string{"LOTUS_TOKEN"},
	   },
   },
   Action: func(c *cli.Context) error {
	   db, closer, err := database.OpenFromCLI(c)
	   if err != nil {
		   return errors.WithStack(err)
	   }
	   defer func() { _ = closer.Close() }()

	   if c.Bool("with-balance") {
		   lotusAPI := c.String("lotus-api")
		   lotusToken := c.String("lotus-token")
		   if lotusAPI == "" || lotusToken == "" {
			   return errors.New("Both --lotus-api and --lotus-token must be provided to fetch wallet balances.")
		   }
		   lotusClient := util.NewLotusClient(lotusAPI, lotusToken)
		   wallets, err := wallet.ListWithBalanceHandler(c.Context, db, lotusClient)
		   if err != nil {
			   return errors.WithStack(err)
		   }
		   cliutil.Print(c, wallets)
		   return nil
	   }

	   wallets, err := wallet.Default.ListHandler(c.Context, db)
	   if err != nil {
		   return errors.WithStack(err)
	   }
	   cliutil.Print(c, wallets)
	   return nil
   },
}
