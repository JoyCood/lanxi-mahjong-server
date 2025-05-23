package algo

// 全部牌值，
// 高四位表示色值(0:万，1：条,2:饼,4:风，5:字)，
// 低四位表示1-9的牌值
var CARDS = []byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,

	0x41, 0x42, 0x43, 0x44,
	0x41, 0x42, 0x43, 0x44,
	0x41, 0x42, 0x43, 0x44,
	0x41, 0x42, 0x43, 0x44,

	0x51, 0x52, 0x53,
	0x51, 0x52, 0x53,
	0x51, 0x52, 0x53,
	0x51, 0x52, 0x53,
}

//番值
var FAN map[uint32]int32 = make(map[uint32]int32)
//初始化
func init() {
	//牌型
	FAN[HU_PING           ] = 1	 // 平胡
	FAN[HU_SINGLE         ] = 2	 // 十三烂
	FAN[HU_SINGLE_ZI      ] = 3	 // 七星十三烂
	FAN[HU_SEVEN_PAIR_BIG ] = 3	 // 大七对
	FAN[HU_SEVEN_PAIR     ] = 4	 // 小七对
	FAN[HU_SEVEN_PAIR_KONG] = 4	 // 豪华小七对
	FAN[HU_ONE_SUIT       ] = 4	 // 清一色
	FAN[HU_ALL_ZI         ] = 8  // 字一色
	//胡牌方式
	FAN[QIANG_GANG    ] = 2  // 抢杠,其他家胡你补杠那张牌
	FAN[HU_KONG_FLOWER] = 2  // 杠上开花,杠完牌抓到的第一张牌自摸了
	FAN[HU_MENQING    ] = 2  // 门清
	FAN[HU_DANDIAO    ] = 2  // 单钓
	FAN[TIAN_HU       ] = 8  // 天胡
	FAN[DI_HU         ] = 8  // 地胡
	//杠牌
	FAN[MING_KONG] = 1 // 明杠
	FAN[AN_KONG  ] = 2 // 暗杠
	FAN[BU_KONG  ] = 1 // 补杠

	//组合牌型
	FAN[HU_ONE_SUIT_PAIR_BIG ] = 5  //清一色大七对
	FAN[HU_ONE_SUIT_PAIR     ] = 8  //清一色小七对
	FAN[HU_ONE_SUIT_PAIR_KONG] = 16 //清一色豪七对
	//---
	FAN[HU_ALL_ZI_PAIR_BIG ] = 8  //字一色大七对
	FAN[HU_ALL_ZI_PAIR     ] = 16 //字一色小七对
	FAN[HU_ALL_ZI_PAIR_KONG] = 16 //字一色豪七对
	//---
	FAN[TIAN_HU_ONE_SUIT       ] = 16 //天胡清一色
	FAN[DI_HU_ONE_SUIT         ] = 16 //地胡清一色
	FAN[TIAN_HU_ALL_ZI         ] = 16 //天胡字一色
	FAN[DI_HU_ALL_ZI           ] = 16 //地胡字一色
	FAN[TIAN_HU_SEVEN_PAIR     ] = 16 //天胡小七对
	FAN[DI_HU_SEVEN_PAIR       ] = 16 //地胡小七对
	FAN[TIAN_HU_SEVEN_PAIR_KONG] = 16 //天胡豪七对
	FAN[DI_HU_SEVEN_PAIR_KONG  ] = 16 //地胡豪七对
}
