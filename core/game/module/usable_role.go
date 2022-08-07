package module

type UsableRole struct {
	BaseRoleID      uint64 // 表明当前的可用角色源自哪一个基础角色
	KeyID           uint64 // 针对同一个基础角色训练出的可用角色做区分
	Assault         uint64 // 突击
	Shooting        uint64 // 射击
	Strength        uint64 // 血量
	CurrentStrength uint64 // 当前血量
}

type UsableRoleModule struct {
	// 由于同一个 base_role 可以通过训练生成不同的 usable_role
	// 因此每个 usable_role 需要通过唯一的 keyID 区分
	OwnUsableRole map[uint64]*UsableRole // 记录玩家训练完成后生成的所有可用角色
}

// 在玩家的训练角色训练完毕后，会依照训练角色基础属性，生成一个可用角色
// func (m *UsableRoleModule) NewUsableRole(id uint64, player *Player) {
// 	// 判断 ID 对应的基础角色是否在全局表中

// 	// 判断 ID 对应的基础角色是否在基础角色表中

// 	// 判断玩家的 training_role 的 ID 是否与该 ID 相同

// 	// 判断玩家拥有的训练角色是否到达上限
// 	if len(m.OwnUsableRole) > def.MAX_USABLE_ROLE_COUNT {

// 		return
// 	}

// 	// 复制训练角色属性，生成一个新的 training_role
// 	newUsableRole := &UsableRole{
// 		BaseRoleID:      id,
// 		Assault:         player.TrainingRoleModule.Role.Assault,
// 		Shooting:        player.TrainingRoleModule.Role.Shooting,
// 		Strength:        player.TrainingRoleModule.Role.Strength,
// 		CurrentStrength: player.TrainingRoleModule.Role.CurrentStrength,
// 	}

// 	m.OwnUsableRole[newUsableRole.KeyID] = newUsableRole

// }
