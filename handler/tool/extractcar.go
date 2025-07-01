package tool

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/oserror"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/boxo/blockservice"
	"github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs"
	"github.com/ipfs/boxo/ipld/unixfs/io"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	carblockstore "github.com/ipld/go-car/v2/blockstore"
	"github.com/urfave/cli/v2"
)

type multiBlockstore struct {
	bss []blockstore.Blockstore
}

func (m multiBlockstore) DeleteBlock(ctx context.Context, c cid.Cid) error {
	return util.ErrNotImplemented
}

func (m multiBlockstore) Has(ctx context.Context, c cid.Cid) (bool, error) {
	for _, bs := range m.bss {
		has, err := bs.Has(ctx, c)
		if err != nil {
			return false, errors.WithStack(err)
		}
		if has {
			return true, nil
		}
	}
	return false, nil
}

func (m multiBlockstore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	for _, bs := range m.bss {
		block, err := bs.Get(ctx, c)
		if errors.Is(err, ipld.ErrNotFound{}) {
			continue
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if block != nil {
			return block, nil
		}
	}
	return nil, ipld.ErrNotFound{Cid: c}
}

func (m multiBlockstore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	for _, bs := range m.bss {
		size, err := bs.GetSize(ctx, c)
		if errors.Is(err, ipld.ErrNotFound{}) {
			continue
		}
		if err != nil {
			return 0, errors.WithStack(err)
		}
		if size > 0 {
			return size, nil
		}
	}
	return 0, ipld.ErrNotFound{Cid: c}
}

func (m multiBlockstore) Put(ctx context.Context, block blocks.Block) error {
	return util.ErrNotImplemented
}

func (m multiBlockstore) PutMany(ctx context.Context, blocks []blocks.Block) error {
	return util.ErrNotImplemented
}

func (m multiBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, util.ErrNotImplemented
}

func (m multiBlockstore) HashOnRead(enabled bool) {
}

// ExtractCarHandler extracts data from CAR (Content Addressable Archives) files found in the input directory.
// The extracted data is written to the specified output directory or file based on the given CID.
//
// CAR files are used in IPFS and other content-addressable file systems to represent file content. They are
// essential for offline data transfer, backups, and other use cases where you might need to work with data
// without IPFS or another file system.
//
// Parameters:
//   - ctx: A pointer to a cli.Context object which provides access to the current command-line context.
//   - inputDir: The directory path where the function will look for CAR files.
//   - output: The destination path where the extracted content will be saved.
//   - c: A CID (Content Identifier) which identifies the data block to be extracted from the CAR file.
//
// Returns:
//   - An error if any issues arise during the extraction process, otherwise nil.
//
// Note:
//
//	The function only supports CAR files with CIDs of type raw or DagProtobuf.
//	It recursively walks the input directory and gathers all .car files.
//	For each CAR file, it reads the content and extracts the data for the given CID, then writes the extracted
//	content to the specified output directory or file. If no CAR files are found in the input directory,
//	an error is returned.
func ExtractCarHandler(ctx *cli.Context, inputDir string, output string, c cid.Cid) error {
	if c.Type() != cid.Raw && c.Type() != cid.DagProtobuf {
		return errors.New("unsupported CID type")
	}

	var files []string
	err := filepath.WalkDir(inputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "failed to walk input directory %s", path)
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".car") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to walk input directory")
	}

	if len(files) == 0 {
		return errors.New("no CAR files found in input directory")
	}

	var bss []blockstore.Blockstore
	for _, f := range files {
		bs, err := carblockstore.OpenReadOnly(f)
		if err != nil {
			return errors.Wrapf(err, "failed to open CAR file %s", f)
		}
		bss = append(bss, bs)
		defer bs.Close()
	}

	bs := &multiBlockstore{bss: bss}
	bserv := blockservice.New(bs, nil)
	dagServ := merkledag.NewDAGService(bserv)
	return writeToOutput(ctx, dagServ, output, c, true)
}

func getOutPathForFile(outPath string, c cid.Cid) (string, error) {
	stat, err := os.Stat(outPath)
	// If the user supply /a/b.txt but the file does not exist, then we need to mkdir -p /a
	if errors.Is(err, oserror.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(outPath), 0o755)
		if err != nil {
			return "", errors.Wrapf(err, "failed to create output directory %s", filepath.Dir(outPath))
		}
		return outPath, nil
	}

	if err != nil {
		return "", errors.Wrapf(err, "failed to stat output path %s", outPath)
	}

	// If the user supply /a/b but b is a directory, then we need to use /a/b/<cid> as the filename
	if stat.IsDir() {
		return filepath.Join(outPath, c.String()), nil
	}

	return outPath, nil
}

func writeToOutput(ctx *cli.Context, dagServ ipld.DAGService, outPath string, c cid.Cid, isRoot bool) error {
	node, err := dagServ.Get(ctx.Context, c)
	if err != nil {
		return errors.Wrapf(err, "failed to get node for CID %s", c)
	}

	switch c.Type() {
	case cid.Raw:
		if isRoot {
			outPath, err = getOutPathForFile(outPath, c)
			if err != nil {
				return errors.Wrapf(err, "failed to get output path for CID %s", c)
			}
		}
		_, _ = fmt.Fprintf(ctx.App.Writer, "Writing to %s\n", outPath)
		return os.WriteFile(outPath, node.RawData(), 0o600)
	case cid.DagProtobuf:
		fsnode, err := unixfs.ExtractFSNode(node)
		if err != nil {
			return errors.Wrapf(err, "failed to extract FSNode for CID %s", c)
		}
		switch fsnode.Type() {
		case unixfs.TFile:
			reader, err := io.NewDagReader(ctx.Context, node, dagServ)
			if err != nil {
				return errors.Wrapf(err, "failed to create dag reader for CID %s", c)
			}
			if isRoot {
				outPath, err = getOutPathForFile(outPath, c)
				if err != nil {
					return errors.Wrapf(err, "failed to get output path for CID %s", c)
				}
			}
			f, err := os.Create(outPath)
			if err != nil {
				return errors.Wrapf(err, "failed to create output file %s", outPath)
			}
			defer f.Close()
			_, _ = fmt.Fprintf(ctx.App.Writer, "Writing to %s\n", outPath)
			_, err = reader.WriteTo(f)
			if err != nil {
				return errors.Wrapf(err, "failed to write to output file %s", outPath)
			}
		case unixfs.TDirectory, unixfs.THAMTShard:
			dirNode, err := io.NewDirectoryFromNode(dagServ, node)
			if err != nil {
				return errors.Wrapf(err, "failed to create directory from node for CID %s", c)
			}
			_, _ = fmt.Fprintf(ctx.App.Writer, "Create Dir %s\n", outPath)
			err = os.MkdirAll(outPath, 0o755)
			if err != nil {
				return errors.Wrapf(err, "failed to create output directory %s", outPath)
			}
			err = dirNode.ForEachLink(ctx.Context, func(link *ipld.Link) error {
				return writeToOutput(ctx, dagServ, filepath.Join(outPath, link.Name), link.Cid, false)
			})
			if err != nil {
				return errors.Wrapf(err, "failed to iterate links for CID %s", c)
			}
		case unixfs.TRaw, unixfs.TMetadata, unixfs.TSymlink:
			return errors.Errorf("unsupported node type %d", fsnode.Type())
		}
	default:
		return errors.Errorf("unsupported CID type %d", c.Type())
	}
	return nil
}
