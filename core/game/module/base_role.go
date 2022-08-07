package module

import (
	"google.golang.org/protobuf/proto"
	"sunborn.com/mini-game-refactor/core/def"
)

type IBaseRoleModule interface {
	HandleListBaseRole(param proto.Message)
}

type BaseRole struct {
	ID              uint64
	Assault         uint64 // 突击
	Shooting        uint64 // 射击
	Strength        uint64 // 血量
	CurrentStrength uint64 // 当前血量
}

type BaseRoleModule struct {
	OwnBaseRole map[uint64]*BaseRole // 记录玩家已经拥有的基础角色
}

func (m *BaseRoleModule) HandleListBaseRole(param proto.Message) {

}

func (m *BaseRoleModule) Add(id uint64, count uint64) {
	// 判断全局物品表中含有该 ID

	// 判断 base_role 表中包含该 ID

	for i := 0; i < int(count); i++ {
		// 玩家拥有的基础角色数量超过限制
		if len(m.OwnBaseRole) > def.MAX_BASE_ROLE_COUNT {

			break
		}

		// 玩家还未拥有该基础角色
		m.OwnBaseRole[id] = &BaseRole{
			ID: id,
		}
	}
}
