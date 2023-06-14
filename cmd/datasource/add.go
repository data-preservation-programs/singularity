package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

var AddCmd = &cli.Command{
	Name:  "add",
	Usage: "Add a new data source to the dataset",
	Subcommands: underscore.Map(fs.Registry, func(r *fs.RegInfo) *cli.Command {
		cmd := datasource.OptionsToCLIFlags(r)
		cmd.Action = func(c *cli.Context) error {
			datasetName := c.Args().Get(0)
			path := c.Args().Get(1)
			db := database.MustOpenFromCLI(c)
			dataset, err := database.FindDatasetByName(db, datasetName)
			if err != nil {
				return handler.NewBadRequestError(err)
			}
			if path == "" {
				return handler.NewBadRequestString("path is required")
			}
			if r.Prefix == "local" {
				path, err = filepath.Abs(path)
				if err != nil {
					return errors.Wrap(err, "failed to get absolute path")
				}
			}
			deleteAfterExport := c.Bool("delete-after-export")
			result := map[string]string{}
			for _, flag := range c.Command.Flags {
				flagName := flag.Names()[0]
				if strings.HasPrefix(flagName, r.Prefix) && c.String(flagName) != "" {
					optionName := strings.SplitN(strings.ReplaceAll(flagName, "-", "_"), "_", 2)[1]
					reg, err := fs.Find(r.Prefix)
					if err != nil {
						return errors.Wrap(err, "failed to find fs")
					}
					option, err := underscore.Find(reg.Options, func(o fs.Option) bool {
						return o.Name == optionName
					})
					if err != nil {
						return errors.Wrap(err, "failed to find option")
					}
					result[option.Name] = c.String(flagName)
				}
			}
			source := model.Source{
				DatasetID:           dataset.ID,
				Type:                r.Prefix,
				Path:                path,
				Metadata:            model.Metadata(result),
				ScanIntervalSeconds: 0,
				ScanningState:       model.Ready,
				DeleteAfterExport:   deleteAfterExport,
				DagGenState:         model.Created,
			}

			handler, err := datasource.DefaultHandlerResolver{}.Resolve(c.Context, source)
			if err != nil {
				return errors.Wrap(err, "failed to resolve handler")
			}

			_, err = handler.List(c.Context, "")
			if err != nil {
				return errors.Wrap(err, "failed to check source")
			}

			err = database.DoRetry(func() error {
				return db.Transaction(func(db *gorm.DB) error {
					err := db.Create(&source).Error
					if err != nil {
						return errors.Wrap(err, "failed to create source")
					}
					dir := model.Directory{
						Name:     path,
						SourceID: source.ID,
					}
					err = db.Create(&dir).Error
					if err != nil {
						return errors.Wrap(err, "failed to create directory")
					}
					return nil
				})
			})

			cliutil.PrintToConsole(source, c.Bool("json"))
			return nil
		}
		return cmd
	}),
}
