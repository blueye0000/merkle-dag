package merkledag

import "encoding/json"

// Hash2File retrieves data from KVStore based on hash and path, then returns the corresponding file content.
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 根据 hash 和 path，返回对应的文件，hash 对应的类型是 tree
	data, err := store.Get(hash)
	if err != nil {
		panic(err)
	}

	var node Node
	err = json.Unmarshal(data, &node)
	if err != nil {
		panic(err)
	}

	if node.Type() == FILE {
		fileNode, ok := node.(File)
		if !ok {
			panic("Invalid file node type")
		}
		return fileNode.Bytes()
	}

	return nil
}
