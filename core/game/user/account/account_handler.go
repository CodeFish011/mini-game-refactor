package account

import (
	"fmt"

	"sunborn.com/mini-game-refactor/core/game"
	"sunborn.com/mini-game-refactor/pb"
	"sunborn.com/mini-game-refactor/server"
	"sunborn.com/mini-game-refactor/transport"
)

func Register() {
	// 将本模块下的所有函数注册到 Dispatcher 下
	server.HandlerRegister(pb.MessageID_CS_ACCOUNT_REGISTER, func(i game.IPlayer, d *transport.Datapack) {
		i.HandleRegister(d)
	})
}

func (m *AccountModule) HandleRegister(data *transport.Datapack) {

	fmt.Println("进入到注册处理")
	// fmt.Println(data)
	// 断言将类型转化为 pb.CS_AccountRegister_Req
	// registerParam := param.(*pb.CS_AccountRegister_Req)

	// 判断用户名是否为空

	// 判断密码是否为空

	// 判断用户是否存在

	// 新增一个用户

	// 向客户端返回消息
}

// func (m *AccountModule) HandleLogin(param proto.Message) {

// }
