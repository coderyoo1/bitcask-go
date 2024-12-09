package fio

import "os"

type FileIOManager struct {
	file *os.File
}

// NewFileIOManager 创建一个新的 FileIOManager
func NewFileIOManager(name string) (*FileIOManager, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &FileIOManager{file: file}, nil
}

func (f FileIOManager) Read(bytes []byte, i int64) (int, error) {
	return f.file.ReadAt(bytes, i)
}

func (f FileIOManager) Write(bytes []byte) (int, error) {
	return f.file.Write(bytes)
}

func (f FileIOManager) Sync() error {
	return f.file.Sync()
}

func (f FileIOManager) Close() error {
	return f.file.Close()
}
