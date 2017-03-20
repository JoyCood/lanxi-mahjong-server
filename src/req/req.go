/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2016-12-26 10:00
 * Filename      : res.go
 * Description   : 玩牌协议消息请求
 * *******************************************************/
package req

import (
	"algo"
	"basic/socket"
	"basic/utils"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"desk"
	"data"
	"csv"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	p1 := protocol.CEnterSocialRoom{}
	socket.Regist(p1.GetCode(), p1, entryroom)
	p2 := protocol.CCreatePrivateRoom{}
	socket.Regist(p2.GetCode(), p2, create)
	p3 := protocol.CDiscard{}
	socket.Regist(p3.GetCode(), p3, discard)
	p4 := protocol.COperate{}
	socket.Regist(p4.GetCode(), p4, operate)
	p5 := protocol.CHu{}
	socket.Regist(p5.GetCode(), p5, hu)
	p6 := protocol.CQiangKong{}
	socket.Regist(p6.GetCode(), p6, qiangkong)
	p7 := protocol.CTrusteeship{}
	socket.Regist(p7.GetCode(), p7, trusteeship)
	p8 := protocol.CReady{}
	socket.Regist(p8.GetCode(), p8, ready)
	p9 := protocol.CBroken{}
	socket.Regist(p9.GetCode(), p9, dice)
	p10 := protocol.CPrivateLeave{}
	socket.Regist(p10.GetCode(), p10, leave)
	p11 := protocol.CKick{}
	socket.Regist(p11.GetCode(), p11, kick)
}

// come in private room
func entryroom(ctos *protocol.CEnterSocialRoom, c inter.IConn) {
	//glog.Infof("enter room -> %s", c.GetUserid())
	stoc := &protocol.SEnterSocialRoom{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil { //已经存在或重复进入
		rdata.Enter(player)
	} else {
		player.ClearRoom()
		rdata = desk.Get(ctos.GetInvitecode())
		if rdata == nil {
			stoc.Error = proto.Uint32(errorcode.RoomNotExist)
			c.Send(stoc)
		} else {
			d := rdata.GetData()
			if d == nil {
				stoc.Error = proto.Uint32(errorcode.RoomNotExist)
				c.Send(stoc)
				return
			}
			payment := d.(*desk.DeskData).Payment
			if payment == 1 { //AA支付
				cost := d.(*desk.DeskData).Cost
				if player.GetDiamond() < cost {
					stoc.Error = proto.Uint32(errorcode.NotEnoughDiamond)
					c.Send(stoc)
					return
				}
			}
			var code int = rdata.Enter(player)
			if code == 1 {
				stoc.Error = proto.Uint32(errorcode.RoomFull)
				c.Send(stoc)
			}
		}
	}
}

// 私人局,创建房间
func create(ctos *protocol.CCreatePrivateRoom, c inter.IConn) {
	//glog.Infof("create room -> %s", c.GetUserid())
	stoc := &protocol.SCreatePrivateRoom{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		// 如果该玩家已经在私人局直接进入
		rdata.Enter(player)
		return
	} else {
		player.ClearRoom()
	}
	//---
	roomtype := ctos.GetRtype()
	rname := ctos.GetRname()
	payment := ctos.GetPayment()
	creator := c.GetUserid()
	csvdata := csv.GetRoom(roomtype)
	if csvdata == nil {
		stoc.Error = proto.Uint32(errorcode.CreateRoomFail)
		c.Send(stoc)
		return
	}
	ante := csvdata.Ante
	cost := csvdata.Costonecount //AA支付
	round := csvdata.Number
	expire := uint32(utils.Timestamp()) + round*600
	if payment == 0 {
		cost = csvdata.Costtwocount //一人支付
	}
	if player.GetDiamond() < cost {
		stoc.Error = proto.Uint32(errorcode.NotEnoughDiamond)
		c.Send(stoc)
		return
	}
	if payment != 1 && payment != 0 {
		stoc.Error = proto.Uint32(errorcode.CreateRoomFail)
		c.Send(stoc)
		return
	}
	var macount uint32 = 0
	code := desk.GenInvitecode()
	//---
	//value, _ := gossdb.C().Incr(data.KEY_ROOM_ID, 1)
	roomid ,_:=  data.GenRoomID()
	//---
	r := desk.NewDeskData(uint32(roomid), round, expire, roomtype, ante, macount,
	cost, payment, creator, rname, code)
	roomdata := &protocol.RoomData{
		Roomid:     proto.Uint32(uint32(roomid)),
		Rtype:      proto.Uint32(roomtype),
		Expire:     proto.Uint32(expire),
		Round:      proto.Uint32(round),
		Rname:      proto.String(rname),
		Invitecode: proto.String(code),
		Count:      proto.Uint32(1),
		Userid:     proto.String(creator),
	}
	//---
	rdata = desk.NewDesk(r)
	desk.Add(code, rdata)
	stoc.Rdata = roomdata
	c.Send(stoc)
}

//打牌
func discard(ctos *protocol.CDiscard, c inter.IConn) {
	stoc := &protocol.SDiscard{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		var card uint32 = ctos.GetCard()
		seat := player.GetPosition()
		if card == 0 {
			stoc.Error = proto.Uint32(errorcode.CardValueZero)
			c.Send(stoc)
		} else {
			rdata.DiscardL(seat, byte(card), false)
		}
	} else {
		glog.Infof("discard room -> %s", c.GetUserid())
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
		c.Send(stoc)
	}
}

//操作
func operate(ctos *protocol.COperate, c inter.IConn) {
	//glog.Infof("operate room -> %s", c.GetUserid())
	stoc := &protocol.SOperate{}
	var card uint32 = ctos.GetCard()
	var value uint32 = ctos.GetValue()
	//glog.Infof("operate card -> %d, value -> %d", card, value)
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		rdata.OperateL(card, value, seat)
	} else {
		glog.Infof("operate room -> %s", c.GetUserid())
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
		c.Send(stoc)
	}
}

//胡牌请求
func hu(ctos *protocol.CHu, c inter.IConn) {
	stoc := &protocol.SHu{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		rdata.OperateL(0, algo.HU, seat)
	} else {
		glog.Infof("hu room -> %s", c.GetUserid())
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
		c.Send(stoc)
	}
}

//抢杠胡牌请求
func qiangkong(ctos *protocol.CQiangKong, c inter.IConn) {
	stoc := &protocol.SQiangKong{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		rdata.OperateL(0, algo.HU, seat)
	} else {
		glog.Infof("qiangkong room -> %s", c.GetUserid())
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
		c.Send(stoc)
	}
}

// 托管 1:托管，0：取消托管
func trusteeship(ctos *protocol.CTrusteeship, c inter.IConn) {
	stoc := &protocol.STrusteeship{}
	var kind uint32 = ctos.GetKind()
	if kind < 0 || kind > 1 {
		stoc.Error = proto.Uint32(errorcode.DataOutOfRange)
		c.Send(stoc)
	} else {
		//TODO:优化
		player := players.Get(c.GetUserid())
		rdata := desk.Get(player.GetInviteCode())
		//
		if rdata != nil {
			//seat := player.GetPosition()
			//rdata.Trust(seat, kind)
		} else {
			glog.Infof("trusteeship room -> %s", c.GetUserid())
			stoc.Error = proto.Uint32(errorcode.NotInRoom)
			c.Send(stoc)
		}
	}
}

// 私人房玩家准备
func ready(ctos *protocol.CReady, c inter.IConn) {
	//glog.Infof("ready room -> %s", c.GetUserid())
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//glog.Infof("id -> %s, code -> %s", c.GetUserid(), player.GetInviteCode())
	//
	stoc := &protocol.SReady{}
	if rdata != nil {
		seat := player.GetPosition()
		ready := ctos.GetReady()
		err := rdata.Readying(seat, ready)
		if err == 1 {
			stoc.Error = proto.Uint32(errorcode.VotingCantLaunchVote)
			c.Send(stoc)
		} else {
			//glog.Infof("id -> %s, code -> %d, %+v", c.GetUserid(), seat, ready)
		}
	} else {
		glog.Infof("id -> %s, code -> %s", c.GetUserid(), player.GetInviteCode())
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
		c.Send(stoc)
	}
}

// 玩家打骰子切牌,发牌
func dice(ctos *protocol.CBroken, c inter.IConn) {
	//glog.Infof("dice room -> %s", c.GetUserid())
	stoc := &protocol.SBroken{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		var ok bool = rdata.Diceing()
		if !ok {
			stoc.Error = proto.Uint32(errorcode.AudienceCannotOperate)
			c.Send(stoc)
			return
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotYourTurn)
		c.Send(stoc)
	}
}

//离开房间
func leave(ctos *protocol.CPrivateLeave, c inter.IConn) {
	stoc := &protocol.SPrivateLeave{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		stoc.Seat = proto.Uint32(seat)
		// 打牌中,不能离开房间,局数没有打完不能离开
		if rdata.Leave(seat) { //TODO:房主离开
			player.ClearRoom()
		} else {
			stoc.Error = proto.Uint32(errorcode.GameStartedCannotLeave)
			c.Send(stoc)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInRoomCannotLeave)
		c.Send(stoc)
	}
}

// 私人局,踢人
func kick(ctos *protocol.CKick, c inter.IConn) {
	stoc := &protocol.SKick{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		var seat uint32 = ctos.GetSeat()
		userid := c.GetUserid()
		if rdata.Kick(userid, seat) {
			stoc.Result = proto.Uint32(uint32(0))
			stoc.Seat = proto.Uint32(seat)
		} else {
			stoc.Error = proto.Uint32(errorcode.StartedNotKick)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
	}
	player.Send(stoc)
}
