package datasource

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	_ "github.com/rclone/rclone/backend/amazonclouddrive"
	_ "github.com/rclone/rclone/backend/azureblob"
	_ "github.com/rclone/rclone/backend/b2"
	_ "github.com/rclone/rclone/backend/box"
	_ "github.com/rclone/rclone/backend/drive"
	_ "github.com/rclone/rclone/backend/dropbox"
	_ "github.com/rclone/rclone/backend/fichier"
	_ "github.com/rclone/rclone/backend/filefabric"
	_ "github.com/rclone/rclone/backend/ftp"
	_ "github.com/rclone/rclone/backend/googlecloudstorage"
	_ "github.com/rclone/rclone/backend/googlephotos"
	_ "github.com/rclone/rclone/backend/hdfs"
	_ "github.com/rclone/rclone/backend/hidrive"
	_ "github.com/rclone/rclone/backend/http"
	_ "github.com/rclone/rclone/backend/internetarchive"
	_ "github.com/rclone/rclone/backend/jottacloud"
	_ "github.com/rclone/rclone/backend/koofr"
	_ "github.com/rclone/rclone/backend/local"
	_ "github.com/rclone/rclone/backend/mailru"
	_ "github.com/rclone/rclone/backend/mega"
	_ "github.com/rclone/rclone/backend/memory"
	_ "github.com/rclone/rclone/backend/netstorage"
	_ "github.com/rclone/rclone/backend/onedrive"
	_ "github.com/rclone/rclone/backend/opendrive"
	_ "github.com/rclone/rclone/backend/oracleobjectstorage"
	_ "github.com/rclone/rclone/backend/pcloud"
	_ "github.com/rclone/rclone/backend/premiumizeme"
	_ "github.com/rclone/rclone/backend/putio"
	_ "github.com/rclone/rclone/backend/qingstor"
	_ "github.com/rclone/rclone/backend/s3"
	_ "github.com/rclone/rclone/backend/seafile"
	_ "github.com/rclone/rclone/backend/sftp"
	_ "github.com/rclone/rclone/backend/sharefile"
	_ "github.com/rclone/rclone/backend/sia"
	_ "github.com/rclone/rclone/backend/smb"
	_ "github.com/rclone/rclone/backend/storj"
	_ "github.com/rclone/rclone/backend/sugarsync"
	_ "github.com/rclone/rclone/backend/swift"
	_ "github.com/rclone/rclone/backend/uptobox"
	_ "github.com/rclone/rclone/backend/webdav"
	_ "github.com/rclone/rclone/backend/yandex"
	_ "github.com/rclone/rclone/backend/zoho"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"io"
	"strings"
	"time"
)

type RCloneHandler struct {
	fs.Fs
}

func (h RCloneHandler) scan(ctx context.Context, path string, last string, ch chan<- Entry) error {
	entries, err := h.Fs.List(ctx, path)
	if err != nil {
		ch <- Entry{Error: err}
		return err
	}

	slices.SortFunc(entries, func(i, j fs.DirEntry) bool {
		return strings.Compare(i.Remote(), j.Remote()) < 0
	})

	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Object:
			ch <- Entry{ScannedAt: time.Now(), Info: v}
		case fs.Directory:
			err = h.scan(ctx, v.Remote(), last, ch)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (h RCloneHandler) Scan(ctx context.Context, path string, last string) <-chan Entry {
	ch := make(chan Entry)
	go func() {
		defer close(ch)
		_ = h.scan(ctx, path, last, ch)
	} ()
	return ch
}

func (h RCloneHandler) Check(ctx context.Context, path string) (fs.DirEntry, error) {
	return h.Fs.NewObject(ctx, path)
}

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.DirEntry, error) {
	object, err := h.Fs.NewObject(ctx, path)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to open object")
	}
	if length == 0 {
		return &EmptyReadCloser{}, object, nil
	}
	option := &fs.RangeOption{Start: offset, End: offset + length - 1}
	reader, err := object.Open(ctx, option)
	return reader, object, err
}

func NewRCloneHandler(ctx context.Context, source model.Source) (*RCloneHandler, error) {
	registry, err := fs.Find(source.Type)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find rclone backend")
	}

	f, err := registry.NewFs(ctx, source.Name, source.Path, configmap.Simple(source.Metadata))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rclone backend")
	}

	return &RCloneHandler{f}, nil
}

var Registry = fs.Registry

func CLIContextToSimpleMap(cctx *cli.Context) configmap.Simple {
	result := make(configmap.Simple)
	for _, flagName := range cctx.LocalFlagNames() {
		if cctx.IsSet(flagName) {
			result[flagName] = cctx.String(flagName)
		}
	}

	return result
}

func OptionsToCLIFlags(regInfo *fs.RegInfo) *cli.Command {
	cmd := &cli.Command{
		Name:      regInfo.Prefix,
		Aliases:   regInfo.Aliases,
		ArgsUsage: "<name> <path>",
		Usage:     regInfo.Description,
	}
	var usageLines []string
	var flags []cli.Flag
	var providerSet = make(map[string]struct{})
	for _, option := range regInfo.Options {
		if strings.HasPrefix(option.Provider, "!") || option.Provider == "" {
			continue
		}
		providers := strings.Split(option.Provider, ",")
		for _, provider := range providers {
			providerSet[provider] = struct{}{}
		}
	}
	for _, option := range regInfo.Options {
		option := option
		apply := func(provider string){
			category := provider
			if option.Advanced {
				if provider != "" {
					category = category + " Advanced Options"
				} else {
					category = "Advanced Options"
				}
			}
			var aliases []string
			if option.ShortOpt != "" {
				aliases = append(aliases, option.ShortOpt)
			}
			envvar := regInfo.Prefix + "-"
			if provider != "" {
				envvar += provider + "-"
			}
			envvar += option.Name
			name := strings.Replace(envvar, "_", "-", -1)
			name = strings.ToLower(name)
			envvar = strings.ReplaceAll(envvar, "-", "_")
			envvar = strings.ToUpper(envvar)
			flag := &cli.StringFlag{
				Name:     name,
				Category: category,
				Usage:    strings.Split(option.Help, "\n")[0],
				Required: option.Required,
				Hidden:   option.Hide&fs.OptionHideCommandLine != 0,
				Value:    fmt.Sprint(option.Default),
				Aliases:  aliases,
				EnvVars:  []string{envvar},
			}
			flags = append(flags, flag)

			usageLines = append(usageLines, "--"+flag.Name)
			lines := underscore.Map(strings.Split(option.Help, "\n"), func(line string) string { return "   " + line })
			usageLines = append(usageLines, lines...)
			if len(option.Examples) > 0 {
				usageLines = append(usageLines, "")
				usageLines = append(usageLines, "   Examples:")
				for _, example := range option.Examples {
					if example.Value == "" {
						example.Value = "<unset>"
					}
					usageLines = append(usageLines, "      " + example.Value)
					for _, line := range strings.Split(example.Help, "\n") {
						usageLines = append(usageLines, "         " + line)
					}
				}
			}
			usageLines = append(usageLines, "")
		}

		if option.Provider == "" {
			apply(option.Provider)
		} else if strings.HasPrefix(option.Provider, "!") {
			excludes := strings.Split(option.Provider[1:], ",")
			for provider := range providerSet {
				if !slices.Contains(excludes, provider) {
					apply(provider)
				}
			}
		} else {
			providers := strings.Split(option.Provider, ",")
			for _, provider := range providers {
				apply(provider)
			}
		}
	}
	cmd.Flags = flags
	cmd.Description = strings.Join(usageLines, "\n")
	return cmd
}
