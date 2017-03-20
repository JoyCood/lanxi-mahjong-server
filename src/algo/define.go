package algo

// 胡牌方式：自摸,放冲,枪杠胡,杠上开花
// 胡牌类型：门清,单钓,天胡,地胡
// 胡牌类型：平胡,十三烂,七星十三烂,大七对,小七对,豪华小七对,清一色,字一色

const (
	// 牌局基础的常量
	TOTAL uint32 = 136 //一副贵州麻将的总数
	BING  uint32 = 2   //同子类型
	TIAO  uint32 = 1   //条子类型
	WAN   uint32 = 0   //万字类型
	FENG  uint32 = 4   //风牌类型
	ZI    uint32 = 5   //字牌类型

	HAND uint32 = 13 //手牌数量
	SEAT uint32 = 4  //最多可参与一桌打牌的玩家数量,不算旁观

	// 碰杠胡掩码,用32位每位代表不同的状态
	DRAW      uint32 = 0      // 摸牌 
	DISCARD   uint32 = 1      // 打牌
	PENG      uint32 = 2 << 0 // 碰
	MING_KONG uint32 = 2 << 1 // 明杠
	AN_KONG   uint32 = 2 << 2 // 暗杠
	BU_KONG   uint32 = 2 << 3 // 补杠
	KONG      uint32 = 2 << 4 // 杠(代表广义的杠)
	CHOW      uint32 = 2 << 5 // 吃
	HU        uint32 = 2 << 6 // 胡(代表广义的胡)

	//胡牌类型
	HU_PING            uint32 = 2 << 8  // 平胡
	HU_SINGLE          uint32 = 2 << 9  // 十三烂
	HU_SINGLE_ZI       uint32 = 2 << 10 // 七星十三烂
	HU_SEVEN_PAIR_BIG  uint32 = 2 << 11 // 大七对
	HU_SEVEN_PAIR      uint32 = 2 << 12 // 小七对
	HU_SEVEN_PAIR_KONG uint32 = 2 << 13 // 豪华小七对
	HU_ONE_SUIT        uint32 = 2 << 14 // 清一色
	HU_ALL_ZI          uint32 = 2 << 15 // 字一色

	ZIMO           uint32 = 2 << 17 // 自摸
	PAOHU          uint32 = 2 << 18 // 炮胡,也叫放冲
	QIANG_GANG     uint32 = 2 << 19 // 抢杠,其他家胡你补杠那张牌
	HU_KONG_FLOWER uint32 = 2 << 20 // 杠上开花,杠完牌抓到的第一张牌自摸了

	HU_MENQING     uint32 = 2 << 21 // 门清
	HU_DANDIAO     uint32 = 2 << 22 // 单钓
	TIAN_HU        uint32 = 2 << 23 // 天胡
	DI_HU          uint32 = 2 << 24 // 地胡

	//组合牌型
	HU_ONE_SUIT_PAIR_BIG  uint32 = HU_ONE_SUIT | HU_SEVEN_PAIR_BIG
	HU_ONE_SUIT_PAIR      uint32 = HU_ONE_SUIT | HU_SEVEN_PAIR
	HU_ONE_SUIT_PAIR_KONG uint32 = HU_ONE_SUIT | HU_SEVEN_PAIR_KONG
	//---
	HU_ALL_ZI_PAIR_BIG  uint32 = HU_ALL_ZI | HU_SEVEN_PAIR_BIG
	HU_ALL_ZI_PAIR      uint32 = HU_ALL_ZI | HU_SEVEN_PAIR
	HU_ALL_ZI_PAIR_KONG uint32 = HU_ALL_ZI | HU_SEVEN_PAIR_KONG
	//---
	TIAN_HU_ONE_SUIT        uint32 = TIAN_HU | HU_ONE_SUIT
	DI_HU_ONE_SUIT          uint32 = DI_HU   | HU_ONE_SUIT
	TIAN_HU_ALL_ZI          uint32 = TIAN_HU | HU_ALL_ZI
	DI_HU_ALL_ZI            uint32 = DI_HU   | HU_ALL_ZI
	TIAN_HU_SEVEN_PAIR      uint32 = TIAN_HU | HU_SEVEN_PAIR
	DI_HU_SEVEN_PAIR        uint32 = DI_HU   | HU_SEVEN_PAIR
	TIAN_HU_SEVEN_PAIR_KONG uint32 = TIAN_HU | HU_SEVEN_PAIR_KONG
	DI_HU_SEVEN_PAIR_KONG   uint32 = DI_HU   | HU_SEVEN_PAIR_KONG
)
