package merkle

type Proof interface {
	Verify(value string, rootHash string) bool
}
