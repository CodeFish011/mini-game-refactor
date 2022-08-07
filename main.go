package main

import (
	"sunborn.com/mini-game-refactor/core/game/user/account"
	"sunborn.com/mini-game-refactor/transport"
)

func main() {
	// s := server.NewServer(server.Config{
	// 	Network: "tcp4",
	// 	Address: "127.0.0.1",
	// 	Port:    "8888",
	// })
	// account.Register()
	// s.Start()
	a := account.AccountModule{}
	a.HandleRegister(&transport.Datapack{})
}
