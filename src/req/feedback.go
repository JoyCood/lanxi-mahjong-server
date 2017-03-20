/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-06-11 16:18
 * Filename      : config.go
 * Description   : 反馈
 * *******************************************************/
package req

import (
	"basic/socket"
	"data"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"time"

	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CFeedback{}
	socket.Regist(p.GetCode(), p, fbconfig)
	n := protocol.CNotice{}
	socket.Regist(n.GetCode(), n, notice)
	a := protocol.CActivity{}
	socket.Regist(a.GetCode(), a, getAllActivity)
	r := protocol.CGetActivityRewards{}
	socket.Regist(r.GetCode(), r, getActivityRewards)
	w := protocol.CWechatShare{}
	socket.Regist(w.GetCode(), w, wechatShare)
	u := protocol.CGetCurrency{}
	socket.Regist(u.GetCode(), u, getCurrency)
}

func getCurrency(ctos *protocol.CGetCurrency, c inter.IConn) {
	//TODO:优化
	player := players.Get(c.GetUserid())
	//
	stoc := &protocol.SGetCurrency{
		Coin:     proto.Uint32(player.GetCoin()),
		Diamond:  proto.Uint32(player.GetDiamond()),
		Exchange: proto.Uint32(player.GetExchange()),
		Ticket:   proto.Uint32(player.GetTicket()),
		Roomcard: proto.Uint32(player.GetRoomCard()),
	}
	c.Send(stoc)
}

func fbconfig(ctos *protocol.CFeedback, c inter.IConn) {
	stoc := &protocol.SFeedback{Error: proto.Uint32(0)}
	database := &data.DataFeedback{
		Userid:     c.GetUserid(),
		Createtime: uint32(time.Now().Unix()),
		Kind:       ctos.GetKind(),
		Content:    ctos.GetContent(),
	}
	if err := database.Save(); err != nil {
		stoc.Error = proto.Uint32(errorcode.FeedfackError)
	}
	c.Send(stoc)
}

func notice(ctos *protocol.CNotice, c inter.IConn) {
	stoc := &protocol.SNotice{}
	// test
	// now := uint32(utils.Timestamp())
	// var l []*data.Notice
	// n := &data.Notice{
	// 	Id: now,
	// 	Title: "test",
	// 	Content: "have fun",
	// 	CTime: now,
	// 	Expire: now + 86400,
	// }
	// d := &data.DataNotice{List:append(l, n)}
	// d.Save()
	// utils.Sleep(1)
	// test end
	database := &data.DataNotice{}
	list := database.GetList()
	if len(list) == 0 {
		stoc.Error = proto.Uint32(errorcode.NoticeListEnpty)
	}
	for _, v := range list {
		notice := &protocol.Notice{}
		notice.Id = &v.Id
		notice.Type = &v.Type
		notice.Title = &v.Title
		notice.Content = &v.Content
		notice.Time = &v.CTime
		stoc.List = append(stoc.List, notice)
	}
	c.Send(stoc)
}

func getAllActivity(ctos *protocol.CActivity, c inter.IConn) {
	//list, _ := data.GetActivityList(c.GetUserid())
	stoc := &protocol.SActivity{}
	//for _, v := range list {
	//	aList := &protocol.ProtoActivity{}
	//	aList.Id = &v.Id
	//	aList.Type = &v.Type
	//	aList.Count = &v.Count
	//	aList.Rewards = &v.Rewards
	//	aList.Starttime = &v.Starttime
	//	aList.Endtime = &v.Endtime
	//	stoc.List = append(stoc.List, aList)
	//}
	c.Send(stoc)
}

func getActivityRewards(ctos *protocol.CGetActivityRewards, c inter.IConn) {
	stoc := &protocol.SGetActivityRewards{}
	//aId := ctos.GetId()
	////TODO:优化
	//player := players.Get(c.GetUserid())
	////
	//count, rewards, err := data.GetActivity(aId, c.GetUserid())
	//if err != nil {
	//	stoc.Error = proto.Uint32(errorcode.ActivityIdError)
	//	c.Send(stoc)
	//} else {
	//	t := csv.GetActivity(aId)
	//	if count < t.Count || rewards != 0 {
	//		stoc.Error = proto.Uint32(errorcode.ActivityRewardFail)
	//		c.Send(stoc)
	//	} else {
	//		// 发奖励 t.Rewards
	//		resource.Rewards(player, t.Rewards, data.RESTYPE13)
	//		err := data.SetActivityRewards(aId, c.GetUserid())
	//		glog.Errorln("SetActivityRewards ERROR:", err)
	//		c.Send(stoc)
	//	}
	//}
	c.Send(stoc)
}

func wechatShare(ctos *protocol.CWechatShare, c inter.IConn) {
	//err := data.UpdateActivity(c.GetUserid(), 1)
	//if err != nil {
	//	//
	//}
	stoc := &protocol.SUpdateActivity{}

	//list, _ := data.GetActivityList(c.GetUserid())
	//for _, v := range list {
	//	aList := &protocol.ProtoActivity{}
	//	aList.Id = &v.Id
	//	aList.Type = &v.Type
	//	aList.Count = &v.Count
	//	aList.Rewards = &v.Rewards
	//	aList.Starttime = &v.Starttime
	//	aList.Endtime = &v.Endtime
	//	stoc.List = append(stoc.List, aList)
	//}
	c.Send(stoc)
}
