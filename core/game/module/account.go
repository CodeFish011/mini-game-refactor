package module

import (
	"sunborn.com/mini-game-refactor/transport"
)

// 用户账号的基础信息模块
type IAccountModule interface {
	HandleRegister(*transport.Datapack) // 处理注册
	// HandleLogin(param proto.Message)    // 处理登录
	// HandleLogout(param proto.Message)           // 处理登出
	// HandleQueryAccountInfo(param proto.Message) // 获取账号信息
}
