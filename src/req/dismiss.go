package req

import (
	"basic/socket"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"desk"

	//"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	l := protocol.CLaunchVote{}
	socket.Regist(l.GetCode(), l, launchVote)
	v := protocol.CVote{}
	socket.Regist(v.GetCode(), v, vote)
}

// 发起房间解散投票
func launchVote(ctos *protocol.CLaunchVote, c inter.IConn) {
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	seat := player.GetPosition()
	stoc := &protocol.SLaunchVote{Seat: proto.Uint32(seat)}
	if rdata != nil {
		err := rdata.Vote(true, seat, 0)
		if err == 1 {
			stoc.Error = proto.Uint32(errorcode.RunningNotVote)
			c.Send(stoc)
		} else if err == 2 {
			stoc.Error = proto.Uint32(errorcode.VotingCantLaunchVote)
			c.Send(stoc)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInPrivateRoom)
		c.Send(stoc)
	}
}

// 玩家进行房间解散投票
func vote(ctos *protocol.CVote, c inter.IConn) {
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	seat := player.GetPosition()
	var vote uint32 = ctos.GetVote() //0同意,1不同意
	stoc := &protocol.SVote{
		Vote: proto.Uint32(vote),
		Seat: proto.Uint32(seat),
	}
	if rdata != nil {
		err := rdata.Vote(false, seat, vote)
		if err == 1 {
			stoc.Error = proto.Uint32(errorcode.RunningNotVote)
			c.Send(stoc)
		} else if err == 2 {
			stoc.Error = proto.Uint32(errorcode.NotVoteTime)
			c.Send(stoc)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInPrivateRoom)
		c.Send(stoc)
	}
}
