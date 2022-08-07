package account

import "sunborn.com/mini-game-refactor/core/game/module"

// 该模块对应数据库中的 account 表
// 点击登陆后，获得 account 信息
type AccountModule struct {
	UID           uint64 // uid
	Nickname      string // 昵称
	Signature     string // 签名
	ShowCharacter uint64 // 看板娘
	Diamond       int64  // 钻石数量
	TP            int64  // 玩家体力值

	Username string // 用户名
	Password string // 密码

	// 一下属于玩家不可见的字段
	AccountType uint64 // 账号类型
}

func NewAccountModule() module.IAccountModule {
	return &AccountModule{}
}
