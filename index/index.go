package index

import "bitcask/data"

// Indexer 抽象内存索引接口
type Indexer interface {
	// Put 向索引中添加 data.LogRecord
	Put(key []byte, pos *data.LogRecord) bool

	// Get 从索引中获取 data.LogRecord
	Get(key []byte) *data.LogRecord

	// Delete 从索引中删除 data.LogRecord
	Delete(key []byte) bool
}
