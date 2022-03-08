package merkle

import (
	"github.com/zrbecker/merkle/hash"
)

func NewLeaf(key string, value string) Node {
	return &leaf{
		parent: nil,
		hash:   hash.Hash("0" + key + value),
		key:    key,
		value:  value,
	}
}

type leaf struct {
	parent Node
	hash   string
	key    string
	value  string
}

func (n *leaf) Parent() Node {
	return n.parent
}

func (n *leaf) SetParent(parent Node) {
	n.parent = parent
}

func (n *leaf) Hash() string {
	return n.hash
}

func (n *leaf) Proof() Proof {
	var parentProof Proof = nil
	if n.parent != nil {
		parentProof = n.parent.Proof()
	}
	return &leafProof{
		hash:        n.hash,
		parentProof: parentProof,
	}
}

func (n *leaf) MinKey() string {
	return n.key
}

func (n *leaf) Height() int {
	return 1
}

func (n *leaf) Key() string {
	return n.key
}

func (n *leaf) Value() string {
	return n.value
}

type leafProof struct {
	hash        string
	parentProof Proof
}

func (p *leafProof) Verify(value string, rootHash string) bool {
	nodeHash := hash.Hash("0" + value)
	if nodeHash != p.hash {
		return false
	}
	if p.parentProof == nil {
		return nodeHash == rootHash
	}
	return p.parentProof.Verify(nodeHash, rootHash)
}
