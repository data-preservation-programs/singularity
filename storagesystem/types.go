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

type Name interface {
	Name() string
}

type Lister interface {
	Name
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
	Name
	// Write writes data to the output storage file.
	Write(ctx context.Context, path string, in io.Reader) (fs.Object, error)
	// Move moves the given object to the given path.
	Move(ctx context.Context, from fs.Object, to string) (fs.Object, error)
	// Remove removes the given object.
	Remove(ctx context.Context, obj fs.Object) error
}

type Reader interface {
	Name
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
	// Name of this fs
	Name string
	// Description of this fs - defaults to Name
	Description string
	// Prefix for command line flags for this fs - defaults to Name if not set
	Prefix string
	// Different Config by Provider
	ProviderOptions []ProviderOptions
}

type Option fs.Option

type ProviderOptions struct {
	Provider            string
	ProviderDescription string
	Options             []Option
}

func (option *Option) ToCLIFlag(prefix string, useBuiltIn bool, category string) cli.Flag {
	if option.Advanced && category == "" {
		category = "Advanced"
	}
	var aliases []string
	if useBuiltIn && option.ShortOpt != "" && prefix == "" {
		aliases = []string{option.ShortOpt}
	}
	var required bool
	if useBuiltIn {
		required = option.Required
	}
	var flag cli.Flag
	name := prefix + option.Name
	usage := strings.Split(option.Help, "\n")[0]
	switch (*fs.Option)(option).Type() {
	case "string":
		//nolint:forcetypeassert
		flag = &cli.StringFlag{
			Category: category,
			Name:     name,
			Required: required,
			Usage:    usage,
			Value:    option.Default.(string),
			Aliases:  aliases,
			EnvVars:  []string{strings.ToUpper(strings.ReplaceAll(name, "-", "_"))},
		}
	case "int":
		//nolint:forcetypeassert
		flag = &cli.IntFlag{
			Category: category,
			Name:     name,
			Required: required,
			Usage:    usage,
			Value:    option.Default.(int),
			Aliases:  aliases,
			EnvVars:  []string{strings.ToUpper(strings.ReplaceAll(name, "-", "_"))},
		}
	case "bool":
		//nolint:forcetypeassert
		flag = &cli.BoolFlag{
			Category: category,
			Name:     name,
			Required: required,
			Usage:    usage,
			Value:    option.Default.(bool),
			Aliases:  aliases,
			EnvVars:  []string{strings.ToUpper(strings.ReplaceAll(name, "-", "_"))},
		}
	default:
		//nolint:forcetypeassert
		flag = &cli.StringFlag{
			Category: category,
			Name:     name,
			Required: required,
			Usage:    usage,
			Value: option.Default.(interface {
				String() string
			}).String(),
			Aliases: aliases,
			EnvVars: []string{strings.ToUpper(strings.ReplaceAll(name, "-", "_"))},
		}
	}
	return flag
}

func (p ProviderOptions) ToCLICommand(short string, long string, description string) *cli.Command {
	command := &cli.Command{
		Name:  short,
		Usage: description,
	}
	var helpLines []string
	margin := "   "
	for _, option := range p.Options {
		flag := option.ToCLIFlag("", true, "")
		command.Flags = append(command.Flags, flag)
		lines := underscore.Map(strings.Split(option.Help, "\n"), func(line string) string { return margin + line })
		helpLines = append(helpLines, "--"+flag.Names()[0])
		helpLines = append(helpLines, lines...)
		if len(option.Examples) > 0 {
			for i, example := range option.Examples {
				if example.Value == "" {
					option.Examples[i].Value = "<unset>"
				}
			}
			helpLines = append(helpLines, "")
			helpLines = append(helpLines, margin+"Examples:")
			maxValueLen := underscore.Max(underscore.Map(option.Examples, func(example fs.OptionExample) int { return len(example.Value) }))
			for _, example := range option.Examples {
				pattern := margin + "   | %-" + strconv.Itoa(maxValueLen) + "s | %s"
				help := strings.Split(example.Help, "\n")
				exampleLine := fmt.Sprintf(pattern, example.Value, help[0])
				helpLines = append(helpLines, exampleLine)
				for _, helpLine := range help[1:] {
					exampleLine = fmt.Sprintf(pattern, "", helpLine)
					helpLines = append(helpLines, exampleLine)
				}
			}
		}
		helpLines = append(helpLines, "")
	}
	command.Description = strings.Join(helpLines, "\n")
	return command
}

var Backends []Backend
var BackendMap = make(map[string]Backend)

func init() {
	for _, regInfo := range fs.Registry {
		if slices.Contains([]string{"crypt", "tardigrade"}, regInfo.Prefix) {
			continue
		}
		backend := Backend{}
		backend.Prefix = regInfo.Prefix
		backend.Name = regInfo.Name
		backend.Description = regInfo.Description

		providerMap := make(map[string]*ProviderOptions)
		var allProviders []string
		for _, option := range regInfo.Options {
			if option.Name == "provider" {
				for _, example := range option.Examples {
					providerMap[example.Value] = &ProviderOptions{
						Provider:            example.Value,
						ProviderDescription: example.Help,
					}
					allProviders = append(allProviders, example.Value)
				}
				continue
			}

			if len(allProviders) == 0 {
				allProviders = []string{""}
				providerMap[""] = &ProviderOptions{}
			}
			var providers []string
			switch {
			case option.Provider == "":
				providers = allProviders
			case strings.HasPrefix(option.Provider, "!"):
				excludes := strings.Split(option.Provider[1:], ",")
				providers = underscore.Difference(allProviders, excludes)
			default:
				providers = strings.Split(option.Provider, ",")
			}

			for _, provider := range providers {
				option := option.Copy()
				option.Examples = underscore.Filter(option.Examples, func(example fs.OptionExample) bool {
					return example.Provider == "" || example.Provider == provider
				})
				_, ok := providerMap[provider]
				if !ok {
					panic("provider not found")
				}
				providerMap[provider].Options = append(providerMap[provider].Options, Option(*option))
			}
		}

		for _, provider := range providerMap {
			backend.ProviderOptions = append(backend.ProviderOptions, *provider)
		}

		sort.Slice(backend.ProviderOptions, func(i, j int) bool {
			return backend.ProviderOptions[i].Provider < backend.ProviderOptions[j].Provider
		})

		Backends = append(Backends, backend)
		BackendMap[backend.Prefix] = backend
	}
}
