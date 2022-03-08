package merkle

import (
	"github.com/zrbecker/merkle/hash"
)

func NewInnerNode(left Node, right Node) Node {
	minKey := left.MinKey()
	if right.MinKey() < minKey {
		minKey = right.MinKey()
	}
	height := 1 + left.Height()
	if height < 1+right.Height() {
		height = 1 + right.Height()
	}
	var n = &innerNode{
		parent: nil,
		left:   left,
		right:  right,
		hash:   hash.Hash("1" + left.Hash() + right.Hash()),
		minKey: minKey,
		height: height,
	}
	left.SetParent(n)
	right.SetParent(n)
	return n
}

type innerNode struct {
	parent Node
	left   Node
	right  Node
	hash   string
	minKey string
	height int
}

func (n *innerNode) Parent() Node {
	return n.parent
}

func (n *innerNode) SetParent(parent Node) {
	n.parent = parent
}

func (n *innerNode) Hash() string {
	return n.hash
}

func (n *innerNode) Proof() Proof {
	var parentProof Proof = nil
	if n.parent != nil {
		parentProof = n.parent.Proof()
	}
	return &innerNodeProof{
		leftHash:    n.left.Hash(),
		rightHash:   n.right.Hash(),
		parentProof: parentProof,
	}
}

func (n *innerNode) MinKey() string {
	return n.minKey
}

func (n *innerNode) Height() int {
	return n.height
}

type innerNodeProof struct {
	leftHash    string
	rightHash   string
	parentProof Proof
}

func (p *innerNodeProof) Verify(childHash string, rootHash string) bool {
	if childHash != p.leftHash && childHash != p.rightHash {
		return false
	}
	nodeHash := hash.Hash("1" + p.leftHash + p.rightHash)
	if p.parentProof == nil {
		return nodeHash == rootHash
	}
	return p.parentProof.Verify(nodeHash, rootHash)
}
