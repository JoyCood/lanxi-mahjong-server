package req

import (
	"protocol"
	"inter"
	"basic/socket"
	"github.com/golang/protobuf/proto"
	"resource"
	"data"
	"errorcode"
	"csv"
	"players"
)

func init() {
	s := protocol.CBuy{}
	socket.Regist(s.GetCode(), s, buy)
	b := protocol.CBuild{}
	socket.Regist(b.GetCode(), b, build)
}

func buy(ctos *protocol.CBuy, c inter.IConn) {
	stoc := &protocol.SBuy{Error: proto.Uint32(0)}
	userdata := players.Get(c.GetUserid())
	d := csv.GetShop(ctos.GetId())
	switch d.Paymenttype {
	case resource.RMB:
		if userdata.GetCoin() >= d.Price {
			resource.ChangeRes(userdata, resource.COIN, int32(d.Price) * -1, data.RESTYPE10)
			resource.ChangeRes(userdata, d.PropId, int32(d.Number), data.RESTYPE10)
			stoc.Result = proto.Uint32(0)
		} else {
			stoc.Result = proto.Uint32(1)
			stoc.Error = proto.Uint32(errorcode.NotEnoughCoin)
		}
	case resource.DIA:
		if userdata.GetDiamond() >= d.Price {
			resource.ChangeRes(userdata, resource.DIAMOND, int32(d.Price) * -1, data.RESTYPE10)
			resource.ChangeRes(userdata, d.PropId, int32(d.Number), data.RESTYPE10)
			stoc.Result = proto.Uint32(0)
		} else {
			stoc.Result = proto.Uint32(1)
			stoc.Error = proto.Uint32(errorcode.NotEnoughDiamond)
		}
	default:
		stoc.Error = proto.Uint32(errorcode.NotEnoughDiamond)
	}
	c.Send(stoc)
}

//绑定用户id
func build(ctos *protocol.CBuild, c inter.IConn) {
	stoc := &protocol.SBuild{Error: proto.Uint32(0)}
	var userid string = ctos.GetId()
	var result uint32
	userdata := players.Get(c.GetUserid())
	if userid == c.GetUserid() {
		result = 1 //不能绑定自己
	} else if userdata.GetBuild() != "" {
		result = 2 //已经绑定
	} else {
		//user := &data.User{Userid: userid}
		//err := user.GetPhotoSexName()
		//if err != nil {
		//	result = 3 //不合格id
		//} else {
		agent := data.Agent_User{Gameid:c.GetUserid()}
		if agent.IsJunior(userid) {
			result = 4 //不能绑定下级自己
		} else {
			if err := agent.Bind(userid); err == nil {
				userdata.SetBuild(userid)
				//绑定成功奖励10钻石
				resource.ChangeRes(userdata, resource.DIAMOND, 10, data.RESTYPE16)
				result = 0
			} else {
				result = 5 //代理商不存在
			}
		}
	}
	stoc.Result = proto.Uint32(result)
	c.Send(stoc)
}
