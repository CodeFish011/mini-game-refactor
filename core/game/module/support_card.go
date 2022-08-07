package module

import "sunborn.com/mini-game-refactor/core/def"

type SupportCard struct {
	ID    uint64 // ID
	Level uint64 // 支援卡的突破度
}

type SupportCardModule struct {
	OwnSupportCard map[uint64]*SupportCard // 使用 map 存储玩家拥有的支援卡
}

func (m *SupportCardModule) Add(id uint64, count uint64) {
	// 判断该 ID 存在于物品全局表

	// 判断该 ID 存在于支援卡表

	for i := 0; i < int(count); i++ {
		// 判定拥有的支援卡是否超出限制
		if len(m.OwnSupportCard)+1 > def.MAX_SUPPORT_CARD_COUNT {

			break
		}

		// 判定玩家是否已经拥有该支援卡
		// 如果玩家已经拥有，则对该支援卡突破度 +1
		if card, ok := m.OwnSupportCard[id]; ok {
			// 如果该支援卡的突破度超过上限，则无法继续突破
			if card.Level >= def.MAX_SUPPORT_CARD_LEVEL {
				//

				continue
			}
			card.Level++
		}

		// 如果玩家还未拥有
		m.OwnSupportCard[id] = &SupportCard{
			ID:    id,
			Level: 1,
		}
	}
}
