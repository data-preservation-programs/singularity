package daggen

import (
	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
)

type DummyNode struct {
	size uint64
	cid  cid.Cid
}

var ErrEmptyNode error = errors.New("fake fs node")

func NewDummyNode(size uint64, cid cid.Cid) DummyNode {
	return DummyNode{size: size, cid: cid}
}

func (f DummyNode) RawData() []byte {
	return nil
}

func (f DummyNode) Cid() cid.Cid {
	return f.cid
}

func (f DummyNode) String() string {
	return "DummyNode - " + f.cid.String()
}

func (f DummyNode) Loggable() map[string]interface{} {
	return nil
}

func (f DummyNode) Resolve(path []string) (interface{}, []string, error) {
	return nil, nil, ErrEmptyNode
}

func (f DummyNode) Tree(path string, depth int) []string {
	return nil
}

func (f DummyNode) ResolveLink(path []string) (*ipld.Link, []string, error) {
	return nil, nil, ErrEmptyNode
}

func (f DummyNode) Copy() ipld.Node {
	return &DummyNode{size: f.size, cid: f.cid}
}

func (f DummyNode) Links() []*ipld.Link {
	return nil
}

func (f DummyNode) Stat() (*ipld.NodeStat, error) {
	return &ipld.NodeStat{}, nil
}

func (f DummyNode) Size() (uint64, error) {
	return f.size, nil
}

