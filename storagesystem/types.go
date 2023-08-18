package storagesystem

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

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
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

// Entry is a struct that represents a single file during a data source scan.
type Entry struct {
	Error error
	Info  fs.Object
}

// Handler is an interface for all relevant operations allowed by an RClone backend.
type Handler interface {
	Reader
	Writer
	Lister
	fs.Abouter
}

type Lister interface {
	// List lists the files at the given path.
	List(ctx context.Context, path string) ([]fs.DirEntry, error)

	// Scan scans the data source starting at the given path and returns a channel of entries.
	// The `last` parameter is used to resume scanning from the last entry returned by a previous scan. It is exclusive.
	// The returned entries must be sorted by path in ascending order.
	Scan(ctx context.Context, path string, last string) <-chan Entry

	// Check checks the size and last modified time of the file at the given path.
	Check(ctx context.Context, path string) (fs.DirEntry, error)
}

type Writer interface {
	// Write writes data to the output storage file.
	Write(ctx context.Context, path string, in io.Reader) (fs.Object, error)
}

type Reader interface {
	// Read reads data from the source storage file starting at the given path and offset, and returns a ReadCloser.
	// The `length` parameter specifies the number of bytes to read.
	// This method is most likely used for retrieving a single block of data.
	Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error)
}

// EmptyReadCloser is a ReadCloser that always returns EOF.
type EmptyReadCloser struct{}

func (e *EmptyReadCloser) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (e *EmptyReadCloser) Close() error {
	return nil
}

type Backend struct {
	LongName   string
	ShortName  string
	ShortUsage string
	LongUsage  string
	Flags      []cli.Flag
	Options    []Option
}

type Option struct {
	Name     string
	FlagName string
	Default  string
}

var Backends []Backend
var BackendMap = make(map[string]Backend)

func init() {
	for _, regInfo := range fs.Registry {
		if slices.Contains([]string{"crypt", "tardigrade"}, regInfo.Prefix) {
			continue
		}
		backend := Backend{}
		backend.ShortName = regInfo.Prefix
		backend.LongName = regInfo.Name
		backend.ShortUsage = regInfo.Description

		var usageLines []string
		var flags []cli.Flag
		providerSet := make(map[string]struct{})
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
			defaultValue := fmt.Sprint(options[0].Default)
			if regInfo.Prefix == "local" && name == "encoding" {
				defaultValue = "Slash,Dot"
			}
			flag := &cli.StringFlag{
				Name:     flagName,
				Category: category,
				Usage:    strings.Split(options[0].Help, "\n")[0],
				Required: options[0].Required,
				Hidden:   options[0].Hide&fs.OptionHideCommandLine != 0,
				Value:    defaultValue,
				Aliases:  aliases,
				EnvVars:  []string{envvar},
			}
			flags = append(flags, flag)
			backend.Options = append(backend.Options, Option{
				Name:     name,
				FlagName: flagName,
				Default:  defaultValue,
			})
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

		backend.Flags = flags
		backend.LongUsage = strings.Join(usageLines, "\n")
		Backends = append(Backends, backend)
		BackendMap[backend.ShortName] = backend
	}
}
