package module

import "sunborn.com/mini-game-refactor/core/def"

type CultivateGoods struct {
	ID    uint64
	Count uint64
}

type CultivateGoodsModule struct {
	OwnCultivateGoods map[uint64]*CultivateGoods
}

func (m *CultivateGoodsModule) Add(id uint64, count uint64) {
	// 判断该 ID 存在于物品全局表

	// 判断该 ID 存在于养成道具表

	for i := 0; i < int(count); i++ {
		// 判定拥有的养成道具是否超出限制
		if len(m.OwnCultivateGoods)+1 > def.MAX_CULTIVATE_GOODS_COUNT {

			break
		}

		// 判定玩家是否已经拥有该养成道具
		// 如果玩家已经拥有，则对该支援卡突破度 +1
		if goods, ok := m.OwnCultivateGoods[id]; ok {
			goods.Count++
		}

		// 如果玩家还未拥有
		m.OwnCultivateGoods[id] = &CultivateGoods{
			ID:    id,
			Count: 1,
		}
	}
}
