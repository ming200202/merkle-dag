package merkledag

import "hash"

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

func Add(store KVStore, node Node, h hash.Hash) []byte {
	switch node.Type() {
	case FILE:
		fileNode := node.(File)
		store.Put([]byte("content"), fileNode.Bytes())
		h.Write(fileNode.Bytes())
		return h.Sum(nil)
	case DIR:
		dirNode := node.(Dir)
		iterator := dirNode.It()
		for iterator.Next() {
			childNode := iterator.Node()
			childMerkleRoot := Add(store, childNode, h)
			h.Write(childMerkleRoot)
		}
		return h.Sum(nil)
	default:
		return nil
	}
}
