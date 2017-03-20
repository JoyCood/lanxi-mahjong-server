/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:22
 * Filename      : userdata.go
 * Description   : 负责处理用户详细信息的数据请求
 * *******************************************************/
package req

import (
	"basic/socket"
	"data"
	"errorcode"
	"inter"
	"protocol"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"players"
	"desk"
)

func init() {
	p := protocol.CUserData{}
	socket.Regist(p.GetCode(), p, getUserDataHdr)
	a := protocol.CArchieve{}
	socket.Regist(a.GetCode(), a, getArchieve)

}

func getUserDataHdr(ctos *protocol.CUserData, c inter.IConn) {
	stoc := &protocol.SUserData{}
	stoc.Data = &protocol.UserData{}
	stoc.Data.Roomid = proto.Uint32(0)
	stoc.Data.Invitecode = proto.String("")
	stoc.Data.Roomtype = proto.Uint32(0)
	// 获取玩家自己的详细资料
	if ctos.GetUserid() == c.GetUserid() {
		userdata := players.Get(c.GetUserid())
		stoc.Data = userdata.ConverDataUser()

		user:= &data.User{Userid:c.GetUserid()}
		photo,_:= user.GetPhotoFromDB()
		stoc.Data.Photo = proto.String(photo)
		// 再次认证是否在房间并判断房间是否过期
		if userdata.GetInviteCode() != "" && desk.Get(userdata.GetInviteCode()) != nil {
			//glog.Infoln("已经在私人局", userdata.GetInviteCode())
			stoc.Data.Roomtype = proto.Uint32(userdata.GetRoomType())
			stoc.Data.Invitecode = proto.String(userdata.GetInviteCode())
			stoc.Data.Roomid = proto.Uint32(userdata.GetRoomID())
		} else {
			if userdata.GetInviteCode() != "" {
				glog.Infoln("不在私人局，或者房间过期", userdata.GetInviteCode())
				userdata.ClearRoom()
			}
		}
	} else {
		if ctos.GetUserid() != "" {
			member := &data.User{Userid: ctos.GetUserid()}
			err := member.Get()
			//glog.Infoln("err:", err)
			if err != nil {
				// TODO : err = err: No such field: %s in obj matchid
				// 老字段不能删除
				// if member.Nickname == "" {
				stoc.Error = proto.Uint32(errorcode.UserDataNotExist)
			} else {
				stoc.Data.Coin = &member.Coin
				stoc.Data.Exp = &member.Exp
				stoc.Data.Vip = &member.Vip
				stoc.Data.Lost = &member.Lost
				stoc.Data.Win = &member.Win
				stoc.Data.Ping = &member.Ping
				stoc.Data.Exchange = &member.Exchange
				stoc.Data.Ticket = &member.Ticket
				stoc.Data.Diamond = &member.Diamond
				stoc.Data.Userid = &member.Userid
				stoc.Data.Sex = &member.Sex
				stoc.Data.Sign = &member.Sign
				stoc.Data.Nickname = &member.Nickname
				stoc.Data.Email = &member.Email
				stoc.Data.Phone = &member.Phone
				stoc.Data.Photo = &member.Photo
				stoc.Data.Ip = &member.Create_ip
				stoc.Data.Createtime = &member.Create_time
				stoc.Data.Terminal = &member.Terminal
				stoc.Data.Birth = &member.Birth
				stoc.Data.Platform = &member.Platform
				stoc.Data.Sound = &member.Sound
			}
		} else {
			stoc.Error = proto.Uint32(errorcode.UsernameEmpty)
		}
	}
	c.Send(stoc)
}

// 获取玩家的战绩
func getArchieve(ctos *protocol.CArchieve, c inter.IConn) {
	stoc := &protocol.SArchieve{}
	if ctos.GetUserid() == "" {
		stoc.Error = proto.Uint32(errorcode.UserDataNotExist)
	} /*else {
		dbase := &data.DataArchive{Userid: ctos.GetUserid()}
		err := dbase.Read()
		if err != nil {
			stoc.Error = proto.Uint32(errorcode.UserDataNotExist)
		} else {
			stoc.Userid = &dbase.Userid
			stoc.Best = &dbase.Best
			stoc.Maxcoin = &dbase.Maxcoin
			stoc.Gaincoin = &dbase.Gaincoin
			stoc.Singlecoin = &dbase.Singlecoin

			stoc.Tianhu = &dbase.Tianhu
			stoc.Dihu = &dbase.Dihu
			stoc.Qinglongdui = &dbase.Qinglongdui
			stoc.Longqi = &dbase.Longqi
			stoc.Qing = &dbase.Qing
			stoc.Qidui = &dbase.Qidui
			stoc.Qingqi = &dbase.Qingqi
			stoc.Pengpenghu = &dbase.Pengpenghu
			stoc.Hutype = &dbase.HuType
		}
	}*/
	c.Send(stoc)
}
