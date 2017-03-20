/**********************************************************
 * Author : Michael
 * Email : dolotech@163.com
 * Last modified : 2016-03-28 20:44
 * Filename :resource.go
 * Description :变更玩家的经济资源,并推变更消息送给玩家
 * *******************************************************/

package resource

import (
	"data"
	"inter"
	"protocol"

	"github.com/golang/protobuf/proto"
	"github.com/golang/glog"
)

const (
	COIN     uint32 = 1
	EXCHANGE uint32 = 2
	TICKET   uint32 = 3
	DIAMOND  uint32 = 4
	VIP      uint32 = 14
	SOUND    uint32 = 15

	EXP  uint32 = 100
	WIN  uint32 = 101
	LOST uint32 = 102
	PING uint32 = 103

	RMB uint32 = 1
	DIA uint32 = 2
)

var RES_HASH = map[uint32]string{
	1:   "Coin",
	2:   "Exchange",
	3:   "Ticket",
	4:   "Diamond",
	100: "Exp",
	101: "win",
	102: "Lost",
	103: "Ping",
}

// 更改单个资源
func ChangeRes(c inter.IPlayer, id uint32, count int32, Type uint32) error {
	m := make(map[uint32]int32)
	m[id] = count
	return ChangeMulti(c, m, Type)
}

// 更改多个资源
func ChangeMulti(userdata inter.IPlayer, res map[uint32]int32, Type uint32) error {
	list := make([]*protocol.ResVo, 0, len(res))
	var record data.DataResChanges
	updataValue := map[string]int32{}
	for id, count := range res {
		var current int32
		switch id {
		case COIN:
			current = int32(userdata.GetCoin()) + count
			if current < 0 {
				current = 0
			}
			userdata.SetCoin(uint32(current))

		case EXCHANGE:
			current = int32(userdata.GetExchange()) + count
			if current < 0 {
				current = 0
			}
			userdata.SetExchange(uint32(current))

		case TICKET:
			current = int32(userdata.GetTicket()) + count
			if current < 0 {
				current = 0
			}
			userdata.SetTicket(uint32(current))

		case DIAMOND:
			current = int32(userdata.GetDiamond()) + count
			if current < 0 {
				current = 0
			}

			userdata.SetDiamond(uint32(current))

		case EXP:
			oldexp := userdata.GetExp()
			current = int32(oldexp) + count
			if current < 0 {
				current = 0
			}

			userdata.SetExp(uint32(current))

		case WIN:
			if count > 0 {
				current = int32(userdata.GetWin()) + count
				userdata.SetWin(uint32(current))
			}
		case LOST:
			if count > 0 {
				current = int32(userdata.GetLost()) + count
				userdata.SetLost(uint32(current))
			}
		case PING:
			if count > 0 {
				current = int32(userdata.GetPing()) + count
				userdata.SetPing(uint32(current))
			}
		case SOUND:
			u := &data.User{Userid: userdata.GetUserid(), Sound: true}
			if err := u.UpdateSound(); err == nil {
				userdata.SetSound(true)
			}
		}

		if key, ok := RES_HASH[id]; ok {
			updataValue[key] = current
		}

		list = append(list, &protocol.ResVo{
			Id:    proto.Uint32(id),
			Count: proto.Int32(current),
		})

		ch := &data.DataResChange{
			Kind:     id,
			Channel:  Type,
			Count:    count,
			Residual: uint32(current),
		}
		record = append(record, ch)
	}
	var err error
	user := &data.User{Userid: userdata.GetUserid()}
	if len(updataValue) > 0 {
		err = user.UpdateResource(updataValue)
		if err != nil {
			glog.Errorln(err)
			return err
		}
	}

	err = record.Save(userdata.GetUserid())
	if err != nil {
		glog.Errorln(err)
	}

	stoc := &protocol.SResource{List: list}
	userdata.Send(stoc)

	return err
}
