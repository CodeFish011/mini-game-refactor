package main

import (
	"sunborn.com/mini-game-refactor/client/core"
	"sunborn.com/mini-game-refactor/pb"
	"sunborn.com/mini-game-refactor/transport"
)

func main() {

	c := core.NewClient("tcp4", "127.0.0.1", "8888")
	data, _ := transport.Pack(pb.MessageID_CS_ACCOUNT_REGISTER, &pb.CS_AccountRegister_Req{
		Username: "root",
		Password: "123456",
	})
	sendMsg, _ := data.Marshal()
	c.Conn.Write(sendMsg)
	select {}
}
