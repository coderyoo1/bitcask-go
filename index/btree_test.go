package index

import (
	"bitcask/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBTree_Put 测试 BTree 的 Put 方法
func TestBTree_Put(t *testing.T) {
	b := NewBTree()
	assert.True(t, b.Put([]byte("hello"), &data.LogRecord{Fid: 20, Offset: 30}))
	assert.True(t, b.Put([]byte("hi"), &data.LogRecord{Fid: 10, Offset: 40}))
	assert.Equal(t, 2, b.tree.Len())
	assert.True(t, b.Put([]byte("hello"), &data.LogRecord{Fid: 50, Offset: 60}))
	assert.Equal(t, 2, b.tree.Len())
}

// TestBTree_Get 测试 BTree 的 Get 方法
func TestBTree_Get(t *testing.T) {
	b := NewBTree()
	assert.True(t, b.Put([]byte("hello"), &data.LogRecord{Fid: 20, Offset: 30}))
	res := b.Get([]byte("hello"))
	assert.NotNil(t, res)
	assert.EqualValues(t, &data.LogRecord{Fid: 20, Offset: 30}, res)
}

// TestBTree_Delete 测试 BTree 的 Delete 方法
func TestBTree_Delete(t *testing.T) {
	b := NewBTree()
	assert.True(t, b.Put([]byte("hello"), &data.LogRecord{Fid: 20, Offset: 30}))
	assert.True(t, b.Delete([]byte("hello")))
	assert.Equal(t, 0, b.tree.Len())
}
