package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

// TestFileIOManager_Write 测试写入文件
func TestFileIOManager_Write(t *testing.T) {
	name := filepath.Join(os.TempDir(), "a.data")
	defer os.Remove(name)
	fio, err := NewFileIOManager(name)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	content := []byte("hello")
	write, err := fio.Write(content)
	assert.Nil(t, err)
	assert.Equal(t, len(content), write)
}

// TestFileIOManager_Read 测试读取文件
func TestFileIOManager_Read(t *testing.T) {
	name := filepath.Join(os.TempDir(), "a.data")
	defer func() {
		_ = os.Remove(name)
	}()
	fio, _ := NewFileIOManager(name)
	content := []byte("hello")
	_, _ = fio.Write(content)

	buf := make([]byte, len(content))
	read, err := fio.Read(buf, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(content), read)

	read, err = fio.Read(buf, 2)
	assert.NotNil(t, err) // EOF error
}

// TestFileIOManager_Sync 测试同步文件
func TestFileIOManager_Sync(t *testing.T) {
	name := filepath.Join(os.TempDir(), "a.data")
	defer func() {
		_ = os.Remove(name)
	}()
	fio, _ := NewFileIOManager(name)
	content := []byte("hello")
	_, _ = fio.Write(content)

	err := fio.Sync()
	assert.Nil(t, err)
}

// TestFileIOManager_Close 测试关闭文件
func TestFileIOManager_Close(t *testing.T) {
	name := filepath.Join(os.TempDir(), "a.data")
	defer func() {
		_ = os.Remove(name)
	}()
	fio, _ := NewFileIOManager(name)
	content := []byte("hello")
	_, _ = fio.Write(content)

	err := fio.Close()
	assert.Nil(t, err)
}
