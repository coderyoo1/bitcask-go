package index

import (
	"bitcask/data"
	"bytes"
	"github.com/google/btree"
	"sync"
)

// BTree 是基于 google/btree 的 BTree 的索引实现
type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

type item struct {
	key   []byte
	value *data.LogRecord
}

// NewBTree 创建一个新的 BTree
func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (i *item) Less(than btree.Item) bool {
	return bytes.Compare(i.key, than.(*item).key) == -1
}

func (b BTree) Put(key []byte, pos *data.LogRecord) bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.tree.ReplaceOrInsert(&item{key, pos})
	return true
}

func (b BTree) Get(key []byte) *data.LogRecord {
	e := b.tree.Get(&item{key, nil})
	return e.(*item).value
}

func (b BTree) Delete(key []byte) bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.tree.Delete(&item{key, nil}) != nil
}
