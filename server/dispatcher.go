package server

import (
	"fmt"
	"log"

	"sunborn.com/mini-game-refactor/core/game"
	"sunborn.com/mini-game-refactor/pb"
	"sunborn.com/mini-game-refactor/transport"
)

type ProcessFunc func(i game.IPlayer, d *transport.Datapack)

type Dispatcher struct {
	// HandlerSet map[pb.MessageID]func(proto.Message) // 记录服务器中所有处理的函数
}

var handlerSet map[pb.MessageID]ProcessFunc

func init() {
	handlerSet = make(map[pb.MessageID]ProcessFunc)
}

// 对整个系统内部暴露一个注册接口
// 各个模块通过该函数将自己模块下的处理函数注册来
func HandlerRegister(id pb.MessageID, f ProcessFunc) {
	handlerSet[id] = f
}

func Dispatch(connID uint64, d *transport.Datapack) {
	fmt.Println(len(handlerSet))
	_, ok := handlerSet[d.ID]
	if !ok {
		log.Printf("没有处理 %s 的函数\n", d.ID.Enum())
		return
	}

	conn, ok := GetConnectionManager().GetConnectionByID(connID)
	if !ok {
		// ID 对应 Connection 不存在
	}
	fmt.Println("此处是dispatch")

	fmt.Printf("%v\n", conn.P.HandleRegister)
	// f(conn.P, d)
}

// func GetProcessFuncByMessageID(id pb.MessageID) ProcessFunc {
// 	s, _ := handlerSet[id]

// 	p := &pb.CS_AccountLogin_Req{}
// 	a := game.Player{}
// 	s(a, p)
// }
