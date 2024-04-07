package merkledag

import (
	"encoding/json"
)

type Link struct {
	Name string
	Hash []byte
	Size int
}

func Add(store KVStore, node Node, hp HashPool) []byte {
	// 将 Node 中的数据保存在 KVStore 中，然后计算出 Merkle Root
	data, err := json.Marshal(node)
	if err != nil {
		panic(err)
	}

	err = store.Put([]byte(node.Name()), data)
	if err != nil {
		panic(err)
	}

	return calculateMerkleRoot(data, hp)
}

func calculateMerkleRoot(data []byte, hp HashPool) []byte {
	hash := hp.Get()
	hash.Write(data)
	return hash.Sum(nil)
}
