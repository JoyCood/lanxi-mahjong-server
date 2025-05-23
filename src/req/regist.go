/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-03-15 18:12
 * Filename      : regist.go
 * Description   : 注册
 * *******************************************************/
package req

import (
	"basic/socket"
	"basic/utils"
	"data"
	"errorcode"
	"inter"
	"protocol"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CRegist{}
	socket.Regist(p.GetCode(), p, regist)
	s := protocol.CSetPasswd{}
	socket.Regist(s.GetCode(), s, setpasswd)
}

func regist(ctos *protocol.CRegist, c inter.IConn) {
	stoc := &protocol.SRegist{}
	if ctos.GetNickname() == "" {
		glog.Errorln("nickname empty")
		stoc.Error = proto.Uint32(errorcode.UsernameEmpty)
		c.Send(stoc)
		return
	}

	if !utils.LegalName(ctos.GetNickname(), 7) {
		glog.Errorln("name legal", ctos.GetNickname())
		stoc.Error = proto.Uint32(errorcode.NameTooLong)
		c.Send(stoc)
		return
	}

	if ctos.GetPhone() == "" {
		glog.Errorln("phone empty")
		stoc.Error = proto.Uint32(errorcode.PhoneNumberEnpty)
		c.Send(stoc)
		return

	}

	if !utils.PhoneRegexp(ctos.GetPhone()) {
		glog.Errorln("phone error",ctos.GetPhone())
		stoc.Error = proto.Uint32(errorcode.PhoneNumberError)
		c.Send(stoc)
		return
	}

	if ctos.GetPwd() == "" {
		glog.Errorln("password empty")
		stoc.Error = proto.Uint32(errorcode.PwdEmpty)
		c.Send(stoc)
		return

	}
	if len(ctos.GetPwd()) != 32 {
		glog.Errorln("password not enough",ctos.GetPwd())
		stoc.Error = proto.Uint32(errorcode.PwdFormatError)
		c.Send(stoc)
		return

	}
	user := data.User{}
	if user.ExistsPhone(ctos.GetPhone()){
		glog.Errorln("Phone exists",ctos.GetPhone())
		stoc.Error = proto.Uint32(errorcode.PhoneRegisted)
		c.Send(stoc)
		return
	}

	userid, err := data.GenUserID()
	if len(userid) > 0 {
		auth := string(utils.GetAuth())
		user := data.User{
			Userid:      userid,
			Nickname:    ctos.GetNickname(),
			Create_ip:   c.GetIPAddr(),
			Auth:        auth,
			Pwd:         utils.Md5(ctos.GetPwd() + auth),
			Coin:        5000,
			Diamond:     10000,
			Exchange:    10,
			Ticket:      10,
			Sex:         3,
			Phone:       ctos.GetPhone(),
			Create_time: uint32( time.Now().Unix()),
		}
		if err := user.Save(); err != nil {
			glog.Errorln(err)
			stoc.Error = proto.Uint32(errorcode.RegistError)
		}
	} else {
		glog.Errorln("generate userid error",err)
		stoc.Error = proto.Uint32(errorcode.RegistError)
	}

	stoc.Userid = &userid
	//glog.Infoln(stoc.String())
	c.Send(stoc)
}

func setpasswd(ctos *protocol.CSetPasswd, c inter.IConn) {
	stoc := &protocol.SSetPasswd{Error: proto.Uint32(0)}
	if ctos.GetPwd() == "" {
		//glog.Errorln("password empty")
		stoc.Error = proto.Uint32(errorcode.PwdEmpty)
		c.Send(stoc)
		return

	}
	if len(ctos.GetPwd()) != 32 {
		//glog.Errorln("password not enough")
		stoc.Error = proto.Uint32(errorcode.PwdFormatError)
		c.Send(stoc)
		return
	}
	userid := c.GetUserid()
	user := &data.User{Userid:userid}
	err := user.UpdatePWD(ctos.GetPwd())
	if err == nil {
		stoc.Result = proto.Uint32(0)
	} else {
		stoc.Result = proto.Uint32(1)
	}
	c.Send(stoc)
}
