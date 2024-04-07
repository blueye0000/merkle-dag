package merkledag

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 根据hash和path， 返回对应的文件, hash对应的类型是tree
	// 从 KVStore 中获取数据
	data, err := store.Get(hash)
	if err != nil {
		// 处理错误
		return nil
	}
	// 返回对应的文件内容
	return data
}
