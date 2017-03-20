/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:22
 * Filename      : login.go
 * Description   : 手机账户和玩家账号登录
 * *******************************************************/
package req

import (
	"basic/event"
	"basic/socket"
	"basic/utils"
	"data"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"desk"
	"runtime/debug"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CLogin{}
	socket.Regist(p.GetCode(), p, login)
}

func login(ctos *protocol.CLogin, c inter.IConn) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}

	}()
	//glog.Infoln(ctos.Userid, ctos.GetPassword())
	stoc := &protocol.SLogin{}
	var member *data.User
	if ctos.GetPhone() != "" && utils.PhoneRegexp(ctos.GetPhone()) {
		member = &data.User{Phone:ctos.GetPhone()}
		if ctos.GetPassword() != "" && len(ctos.GetPassword()) == 32 && member.VerifyPwdByPhone(ctos.GetPassword()) {
			c.SetUserid(member.Userid)
		} else {
			//glog.Errorln(member.Userid)
			stoc.Error = proto.Uint32(errorcode.UsernameOrPwdError)
		}

	} else {
		stoc.Error = proto.Uint32(errorcode.UsernameOrPwdError)
	}
	if member != nil {
		err := member.Get()
		if err != nil {
			stoc.Error = proto.Uint32(errorcode.UsernameOrPwdError)
		}
	}

	if stoc.Error != nil {
		c.Send(stoc)
		time.AfterFunc(time.Millisecond*200, c.Close)
		return
	}
	//重复登录检测
	logining(member, c)
	stoc.Userid = &member.Userid
	c.Send(stoc)

}

func logining(member *data.User, c inter.IConn) {
	//	重复登录检测
	glog.Infoln("玩家ID : ", c.GetUserid())
	// 已经在房间打牌
	userdata := players.Get(member.Userid)
	if userdata != nil {
		conn := userdata.GetConn()
		conn.(event.IDispath).RemoveAll()
		conn.Close()

	} else {
		// 登陆成功把用户的数据从数据库取出存入服务内存
		userdata = desk.NewPlayer(member)
		players.Set(userdata.GetUserid(), userdata)
	}
	userdata.SetConn(c)

	c.SetLogin()

	//注释掉下面玩家下线处理，玩家数据会越积越多，造成内存泄漏
	c.(event.IDispath).ListenOnce(socket.OFFLINE, func(t string, args interface{}) {
		conn := args.(inter.IConn)
		player := players.Get(conn.GetUserid())
		if player != nil {
			//正在游戏不清理
			if player.GetRoomType() == 0 {
				//players.Del(player.GetUserid()) //TODO:直接删除有问题,延迟清理
			}
		}
		active := &data.DataUserActive{Userid: member.Userid, IP: c.GetIPAddr()}
		active.Logout()
	})
	// 记录登陆时间和IP地址
	go func() {
		active := &data.DataUserActive{Userid: member.Userid, IP: c.GetIPAddr()}
		active.Login()
		tradeOff(userdata) //发货失败订单检测
	}()

}
