package merkle

type Node interface {
	Parent() Node
	SetParent(parent Node)
	Hash() string
	Proof() Proof
	MinKey() string
	Height() int
}

type LeafNode interface {
	Proof() Proof
	Key() string
	Value() string
}
