package module

type TrainingRole struct {
	BaseRoleID           uint64 // 对应基础角色 ID
	Assault              uint64 // 突击
	Shooting             uint64 // 射击
	Strength             uint64 // 血量
	CurrentStrength      uint64 // 当前血量
	Health               uint64 // 体力
	CurrentHealth        uint64 // 当前体力
	Coin                 uint64 // 金币数
	ScheduleRefreshCount uint64 // 日程刷新次数
	ChipRefreshCount     uint64 // 芯片刷新次数
	ChipRemakeCount      uint64 // 芯片蚀刻次数
	ShopRefreshCount     uint64 // 商店刷新次数

}

// 玩家同一时间只能有一个 training_role
type TrainingRoleModule struct {
	Role *TrainingRole
}

// 传入 base_role 的 ID，依据 base_role 的数值作为 training_role 的初始数值
// func (m *TrainingRoleModule) NewTrainingRole(id uint64, player *game.Player) {
// 	// 校验 ID 是否存在与全局的 ID 中

// 	// 校验 ID 是否存在于 base_role 表中

// 	// 从表中根据 ID 加载 base_role 的数值

// 	// 判断用户是否已经有一个训练角色

// 	// 没有，则生成一个训练角色
// 	player.TrainingRoleModule.Role = &TrainingRole{
// 		BaseRoleID:           id,
// 		Assault:              1,
// 		Shooting:             1,
// 		Strength:             1,
// 		CurrentStrength:      1,
// 		Coin:                 0,
// 		ScheduleRefreshCount: 1,
// 		ChipRefreshCount:     1,
// 		ChipRemakeCount:      1,
// 		ShopRefreshCount:     1,
// 	}
// }
