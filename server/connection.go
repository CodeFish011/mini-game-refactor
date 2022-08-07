package server

import (
	"fmt"
	"log"
	"net"

	"sunborn.com/mini-game-refactor/core/game"
	"sunborn.com/mini-game-refactor/pb"
	"sunborn.com/mini-game-refactor/transport"
)

type IConnection interface {
	Start()
	Stop()
}

type Connection struct {
	ID        uint64
	Conn      net.Conn                 // 实际的链接
	ReadChan  chan *transport.Datapack // 读取客户端的数据
	WriteChan chan *transport.Datapack // 向客户端发送数据
	P         *game.Player             // 当前链接登录的游戏账户
}

func NewConnection(conn net.Conn, config *Config) *Connection {
	return &Connection{
		Conn:      conn,
		ReadChan:  make(chan *transport.Datapack, config.ConnectionReadChanBufSize),  // 设置读缓冲区大小
		WriteChan: make(chan *transport.Datapack, config.ConnectionWriteChanBufSize), // 设置写缓冲区大小
		P:         game.NewPlayer(),                                                  // 链接一个账号信息
	}
}

func (c *Connection) Start() {
	go c.Read()
	go c.Write()
}

func (c *Connection) Stop() {

}

// 从自己的链接中读入数据，反序列化得到 datapack
// 到目前仍是 「独立」 状态
func (c *Connection) Read() {
	for {
		d, err := transport.Unmarshal(c.Conn)
		if err != nil {
			continue
		}
		fmt.Println(d.ID.Enum())
		// 当前用户没有登陆，且发送的不是登陆请求
		if c.P.IsLogined || d.ID != pb.MessageID_CS_ACCOUNT_LOGIN && d.ID != pb.MessageID_CS_ACCOUNT_REGISTER {
			fmt.Println(c.Conn.RemoteAddr().String(), "必须在登录后才能执行操作")
			data, _ := transport.Pack(pb.MessageID_LOGIN_REQUIRED, &pb.CS_LoginRequired{
				Tip: "您必须在登录后才能执行操作",
			})
			c.WriteChan <- data
			continue
		}

		// 此时获取到了 datapack
		// 需要发送到系统内部处理对应 MessageID 的模块进行处理
		go Dispatch(c.ID, d)
	}
}

func (c *Connection) Write() {
	for {
		select {
		case data := <-c.WriteChan:
			var err error
			sendMsg, err := data.Marshal()
			if err != nil {
				log.Printf("序列化发送消息失败, ID: %d, 地址: %s, 消息类型: %s\n", data.ID, c.Conn.RemoteAddr().String(), data.ID.Type())
				continue
			}
			_, err = c.Conn.Write(sendMsg)
			if err != nil {
				log.Printf("向客户端发送数据失败, ID: %d, 地址: %s, 消息类型: %s\n", data.ID, c.Conn.RemoteAddr().String(), data.ID.Type())
				continue
			}
		}
	}
}
