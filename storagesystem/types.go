package storagesystem

import (
	"context"
	"io"
	"sort"
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

type ProviderOptions struct {
	Provider            string
	ProviderDescription string
	Options             []fs.Option
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
			if option.Provider == "" {
				providers = allProviders
			} else if option.Provider == "" {
				providers = []string{""}
			} else if strings.HasPrefix(option.Provider, "!") {
				excludes := strings.Split(option.Provider[1:], ",")
				providers = underscore.Difference(allProviders, excludes)
			} else {
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
				providerMap[provider].Options = append(providerMap[provider].Options, *option)
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
