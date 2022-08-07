package game

import (
	"sunborn.com/mini-game-refactor/core/game/module"
)

type IPlayer interface {
	module.IAccountModule // 账号信息模块
	// module.IBaseRoleModule // 基础角色模块

	// ISupportCardModule    // 支援卡模块
	// ICultivateGoodsModule // 养成道具模块

	// ITrainingRoleModule // 训练角色模块
	// IUsableRoleModule   // 可用角色模块
}

// 表示一个具体的玩家
// 每个玩家由各个模块组成
type Player struct {
	// IPlayer
	module.IAccountModule
	// AccountModule *account.AccountModule // 账号信息模块
	// SupportCardModule    *SupportCardModule    // 支援卡模块
	// CultivateGoodsModule *CultivateGoodsModule // 养成道具模块
	// BaseRoleModule       *BaseRoleModule       // 基础角色模块
	// TrainingRoleModule   *TrainingRoleModule   // 训练角色模块
	// UsableRoleModule     *UsableRoleModule     // 可用角色模块

	// 非模块字段
	IsLogined bool
}

func NewPlayer() *Player {
	return &Player{}
}
