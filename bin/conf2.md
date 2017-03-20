//牌值
万 [1 2 3 4 5 6 7 8 9 ]
条 [17 18 19 20 21 22 23 24 25 ]
筒 [33 34 35 36 37 38 39 40 41 ]
字 [65 66 67 68 ]
风 [81 82 83]

万 W1,W2,W3,W4,W5,W6,W7,W8,W9
条 T1,T2,T3,T4,T5,T6,T7,T8,T9
筒 B1,B2,B3,B4,B5,B6,B7,B8,B9
字 Z1,Z2,Z3,Z4
风 F1,F2,F3

//掩码值
MING_KONG uint32 = 2 << 1 // 明杠
AN_KONG   uint32 = 2 << 2 // 暗杠
BU_KONG   uint32 = 2 << 3 // 补杠
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
HU_KONG_FLOWER uint32 = 2 << 20 // 杠上开花
HU_MENGQING    uint32 = 2 << 21 // 门清
HU_DANDIAO     uint32 = 2 << 22 // 单钓
TIAN_HU        uint32 = 2 << 23 // 天胡
DI_HU          uint32 = 2 << 24 // 地胡

//json data
{
	"pao":0,       //放冲位置(自摸时填自己位置)
	"dealer":0,    //庄家位置
	"seat1":{                      //位置
		"hu":"HU,PAOHU",       //胡牌类型,多个逗号隔开
		"cards":"W8,W8,W8,W8", //手牌,(大对子暗杠数,胡牌牌型)
		"pong": {           //碰  
			"seat":"2", //被碰位置 
			"card":"B1" //碰的牌值
		},
		"chow": {                  //吃
			"seat":"4",        //上家位置(被吃位置)
			"card":"T1,T2,T3"  //吃的牌
		},
		"mingKong": {         //明杠
			"seat":"2",   //被杠位置
			"card":"T7"   //明杠牌值
		},
		"anKong": {           //暗杠
			"seat":"0",   //暗杠位置可以为0
			"card":"B7"   //暗杠牌值
		},
		"buKong": {            //补杠
			"seat":"0,0",  //补杠位置可以为0
			"card":"F1,F2" //补杠牌值
		}
	},
	"seat2":{
	},
	"seat3":{
	},
	"seat4":{
	}
}
