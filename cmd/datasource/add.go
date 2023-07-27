package datasource

import (
	"path/filepath"
	"strings"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var exclude = []string{
	"CreatedAt", "UpdatedAt", "ScanningWorkerID", "LastScannedTimestamp", "DagGenWorkerID", "Metadata",
}

var AddCmd = &cli.Command{
	Name:  "add",
	Usage: "Add a new data source to the dataset",
	Subcommands: underscore.Map(underscore.Filter(fs.Registry, func(r *fs.RegInfo) bool {
		return !slices.Contains([]string{"crypt", "memory", "tardigrade"}, r.Prefix)
	}), func(r *fs.RegInfo) *cli.Command {
		cmd := datasource.OptionsToCLIFlags(r)
		cmd.Flags = append(cmd.Flags, &cli.BoolFlag{
			Name:     "delete-after-export",
			Usage:    "[Dangerous] Delete the files of the dataset after exporting it to CAR files. ",
			Category: "Data Preparation Options",
		}, &cli.DurationFlag{
			Name:        "rescan-interval",
			Usage:       "Automatically rescan the source directory when this interval has passed from last successful scan",
			Category:    "Data Preparation Options",
			DefaultText: "disabled",
		})
		cmd.Action = func(c *cli.Context) error {
			datasetName := c.Args().Get(0)
			path := c.Args().Get(1)
			db, err := database.OpenFromCLI(c)
			if err != nil {
				return err
			}
			dataset, err := database.FindDatasetByName(db, datasetName)
			if err != nil {
				return handler.InvalidParameterError{Err: err}
			}
			if path == "" {
				return handler.NewInvalidParameterErr("path is required")
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
			if err != nil {
				return cli.Exit(errors.Wrap(err, "failed to add source").Error(), 1)
			}

			cliutil.PrintToConsole(source, c.Bool("json"), exclude)
			return nil
		}
		return cmd
	}),
}
