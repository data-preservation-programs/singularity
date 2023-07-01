package datasource

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/singularity/model"
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
	"sort"
	"strconv"
	"strings"
)

type RCloneHandler struct {
	fs.Fs
}

func (h RCloneHandler) List(ctx context.Context, path string) ([]fs.DirEntry, error) {
	return h.Fs.List(ctx, path)
}

func (h RCloneHandler) scan(ctx context.Context, path string, last string, ch chan<- Entry) error {
	entries, err := h.Fs.List(ctx, path)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- Entry{Error: err}:
		}
		return err
	}

	slices.SortFunc(entries, func(i, j fs.DirEntry) bool {
		return strings.Compare(i.Remote(), j.Remote()) < 0
	})

	startScanning := last == "" // Start scanning immediately if 'last' is empty.
	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Directory:
			dirPath := v.Remote()
			// If 'last' starts with directory path followed by a slash, scan inside the directory with the remaining path.
			if strings.HasPrefix(last, dirPath+"/") {
				err = h.scan(ctx, dirPath, last, ch)
			} else if startScanning || strings.Compare(dirPath, last) > 0 {
				// If we have started scanning or the directory is greater than 'last', scan inside without 'last' param.
				err = h.scan(ctx, dirPath, "", ch)
			}
			if err != nil {
				return err
			}

		case fs.Object:
			// If 'last' is specified, skip entries until the first entry greater than 'last' is found.
			if !startScanning {
				if strings.Compare(entry.Remote(), last) > 0 {
					startScanning = true // Found the first entry greater than 'last', start scanning.
				} else {
					continue
				}
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- Entry{Info: v}:
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
	}()
	return ch
}

func (h RCloneHandler) Check(ctx context.Context, path string) (fs.DirEntry, error) {
	return h.Fs.NewObject(ctx, path)
}

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
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

	f, err := registry.NewFs(ctx, source.Type, source.Path, configmap.Simple(source.Metadata))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rclone backend")
	}

	return &RCloneHandler{f}, nil
}

func OptionsToCLIFlags(regInfo *fs.RegInfo) *cli.Command {
	cmd := &cli.Command{
		Name:      regInfo.Prefix,
		Aliases:   regInfo.Aliases,
		ArgsUsage: "<dataset_name> <source_path>",
		Usage:     regInfo.Description,
	}
	var usageLines []string
	var flags []cli.Flag
	var providerSet = make(map[string]struct{})
	var optionsByName = make(map[string][]fs.Option)
	for _, option := range regInfo.Options {
		optionsByName[option.Name] = append(optionsByName[option.Name], option)
		if strings.HasPrefix(option.Provider, "!") || option.Provider == "" {
			continue
		}
		providers := strings.Split(option.Provider, ",")
		for _, provider := range providers {
			providerSet[provider] = struct{}{}
		}
	}

	var optionsByNameSorted []string
	for name := range optionsByName {
		optionsByNameSorted = append(optionsByNameSorted, name)
	}
	sort.Strings(optionsByNameSorted)
	for _, name := range optionsByNameSorted {
		options := optionsByName[name]
		category := ""
		if options[0].Advanced {
			category = "Advanced Options"
		}
		var aliases []string
		if options[0].ShortOpt != "" {
			aliases = append(aliases, options[0].ShortOpt)
		}
		envvar := strings.ToUpper(regInfo.Prefix + "_" + name)
		flagName := strings.ToLower(strings.ReplaceAll(envvar, "_", "-"))
		flag := &cli.StringFlag{
			Name:     flagName,
			Category: category,
			Usage:    strings.Split(options[0].Help, "\n")[0],
			Required: options[0].Required,
			Hidden:   options[0].Hide&fs.OptionHideCommandLine != 0,
			Value:    fmt.Sprint(options[0].Default),
			Aliases:  aliases,
			EnvVars:  []string{envvar},
		}
		flags = append(flags, flag)
		usageLines = append(usageLines, "--"+flag.Name)
		for _, option := range options {
			margin := "   "
			if option.Provider != "" {
				margin = "      "
			}
			var providers []string
			if strings.HasPrefix(option.Provider, "!") {
				excludes := strings.Split(option.Provider[1:], ",")
				for provider := range providerSet {
					if !slices.Contains(excludes, provider) {
						providers = append(providers, provider)
					}
				}
			} else if option.Provider != "" {
				providers = strings.Split(option.Provider, ",")
			}
			sort.Strings(providers)
			if option.Provider != "" {
				usageLines = append(usageLines, "   [Provider] - "+strings.Join(providers, ", "))
			}
			lines := underscore.Map(strings.Split(option.Help, "\n"), func(line string) string { return margin + line })
			usageLines = append(usageLines, lines...)
			if len(option.Examples) > 0 {
				for i, example := range option.Examples {
					if example.Value == "" {
						option.Examples[i].Value = "<unset>"
					}
				}
				usageLines = append(usageLines, "")
				usageLines = append(usageLines, margin+"Examples:")
				maxValueLen := underscore.Max(underscore.Map(option.Examples, func(example fs.OptionExample) int { return len(example.Value) }))
				for _, example := range option.Examples {
					pattern := margin + "   | %-" + strconv.Itoa(maxValueLen) + "s | %s"
					helpLines := strings.Split(example.Help, "\n")
					exampleLine := fmt.Sprintf(pattern, example.Value, helpLines[0])
					usageLines = append(usageLines, exampleLine)
					for _, helpLine := range helpLines[1:] {
						usageLines = append(usageLines, margin+"     "+strings.Repeat(" ", maxValueLen)+" | "+helpLine)
					}
				}
			}
			usageLines = append(usageLines, "")
		}
	}

	slices.SortFunc(flags, func(i, j cli.Flag) bool { return i.Names()[0] < j.Names()[0] })
	cmd.Flags = flags
	cmd.Description = strings.Join(usageLines, "\n")
	return cmd
}
