package fio

// IOManager 抽象文件 IO 接口
type IOManager interface {
	// Read 从文件指定位置读取数据
	Read([]byte, int64) (int, error)

	// Write 向文件指定位置写入数据
	Write([]byte) (int, error)

	// Sync 同步文件
	Sync() error

	// Close 关闭文件
	Close() error
}
