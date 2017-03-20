package desk

import (
	"basic/utils"
	"data"
)

const (
	ROOM_PRIVATE = 1 // 私人房房间类型
	ROOM_FREE    = 2 //  自由场房间类型
	ROOM_MATCH   = 3 //比赛场房间类型
)

var DT1 int = int(data.Conf.DiscardTimeout) //出牌超时时间
var OT1 int = int(data.Conf.OperateTimeout) //操作超时时间

var DT int = 12 //出牌超时时间
var OT int = 10 //操作超时时间

func NewDeskData(rid, round, expire, rtype, ante, macount, cost,
payment uint32, creator, rname, invitecode string) *DeskData {
	return &DeskData{
		Rid:     rid,
		Rtype:   rtype,
		Rname:   rname,
		Ante:    ante,
		Payment: payment,
		Cost:    cost,
		Cid:     creator,
		Expire:  expire,
		Round:   round,
		Code:    invitecode,
		CTime:   uint32(utils.Timestamp()),
		Score:   make(map[string]int32),
	}
}

type DeskData struct {
	Rid     uint32           //房间ID
	Rtype   uint32           //房间类型
	Rname   string           //房间名字
	Cid     string           //房间创建人
	Expire  uint32           //牌局设定的过期时间
	Code    string           //房间邀请码
	Round   uint32           //剩余牌局数
	Ante    uint32           //私人房底分
	Payment uint32           //付费方式1=AA or 0=房主支付
	Cost    uint32           //创建消耗
	CTime   uint32           //创建时间
	Score   map[string]int32 //私人局用户战绩积分
}
