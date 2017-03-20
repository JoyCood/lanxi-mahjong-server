/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-05-29 11:39
 * Filename      : userdatamodify.go
 * Description   :  用户资料修改
 * *******************************************************/
package req

import (
	"basic/socket"
	"basic/utils"
	"errorcode"
	"inter"
	"players"
	"protocol"

	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CChangeNickname{}
	socket.Regist(p.GetCode(), p, modifyUsername)
	s := protocol.CChangeSex{}
	socket.Regist(s.GetCode(), s, modifyUserSex)
}

// 修改昵称
func modifyUsername(ctos *protocol.CChangeNickname, c inter.IConn) {
	//TODO:优化
	player := players.Get(c.GetUserid())
	//
	stoc := &protocol.SChangeNickname{}
	if !utils.LegalName(ctos.GetNickname(), 7) {
		stoc.Error = proto.Uint32(errorcode.NameTooLong)
	} else {
		nickname := ctos.GetNickname()
		stoc.Nickname = &nickname
		player.SetNickname(ctos.GetNickname())
	}

	c.Send(stoc)
}

// 修改性别
func modifyUserSex(ctos *protocol.CChangeSex, c inter.IConn) {
	stoc := &protocol.SChangeSex{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	//
	if ctos.GetSex() > 3 || ctos.GetSex() < 1 {
		stoc.Error = proto.Uint32(errorcode.SexValueRangeout)
	} else {
		sex := ctos.GetSex()
		stoc.Sex = &sex
		player.SetSex(ctos.GetSex())
	}
	c.Send(stoc)
}
