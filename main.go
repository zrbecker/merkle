package main

import (
	"fmt"

	"github.com/zrbecker/merkle/merkle"
)

func findAndVerify(root merkle.Node, searchKey string, verifyKey string, verifyValue string) bool {
	if root == nil {
		return false
	}
	leaf := merkle.FindBST(root, searchKey)
	proof := leaf.Proof()
	return proof.Verify(verifyKey+verifyValue, root.Hash())
}

func main() {
	var root merkle.Node = nil
	root, _ = merkle.InsertBST(root, "key2", "value2")
	root, _ = merkle.InsertBST(root, "key1", "value1")
	root, _ = merkle.InsertBST(root, "key3", "value3")
	root, _ = merkle.InsertBST(root, "key0", "value0")
	root, _ = merkle.InsertBST(root, "key9", "value9")
	root, _ = merkle.InsertBST(root, "key8", "value8")
	root, _ = merkle.InsertBST(root, "key4", "value4")
	root, _ = merkle.InsertBST(root, "key5", "value5")
	root, _ = merkle.InsertBST(root, "key7", "value7")
	root, _ = merkle.InsertBST(root, "key6", "value6")
	root, _ = merkle.InsertBST(root, "key3", "value?")

	merkle.PrintBST(root)
	fmt.Println()

	fmt.Println("Expect: true, Actual:", findAndVerify(root, "key4", "key4", "value4"))
	fmt.Println("Expect: true, Actual:", findAndVerify(root, "key3", "key3", "value?"))
	fmt.Println("Expect: false, Actual:", findAndVerify(root, "key4", "key4", "value3"))
	fmt.Println("Expect: false, Actual:", findAndVerify(root, "key4", "key3", "value3"))
	fmt.Println("Expect: false, Actual:", findAndVerify(nil, "key4", "key4", "value4"))
	fmt.Println()

	fmt.Println("Check Balance:", merkle.CheckBalance(root))
}
