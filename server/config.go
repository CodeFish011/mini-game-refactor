package server

var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		MaxConnectionCount:         5,
		ConnectionReadChanBufSize:  5,
		ConnectionWriteChanBufSize: 5,
	}
}

type Config struct {
	Network string // 链接类型
	Address string // 地址
	Port    string // IP

	MaxConnectionCount         uint64 // 全局最大连接数
	ConnectionReadChanBufSize  uint64 // 链接读最大缓冲
	ConnectionWriteChanBufSize uint64 // 链接写最大缓冲
}
