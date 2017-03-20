/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 11:01
 * Filename      : iuser.go
 * Description   : 玩家自己的详细数据接口
 * *******************************************************/
package inter

import "protocol"

type IPlayer interface {
	GetUserid() string
	GetPosition() uint32
	GetNickname() string
	GetPhone() string

	SetUserid(string)
	SetPosition(uint32)
	SetNickname(string)
	SetSex(uint32)
	SetPwd(string)
	SetReady(bool)
	GetReady() bool
	SetCoin(uint32)
	GetCoin() uint32

	GetExp() uint32
	SetExp(uint32)

	GetDiamond() uint32
	SetDiamond(uint32)

	GetExchange() uint32
	SetExchange(uint32)

	SetTicket(uint32)
	GetTicket() uint32

	GetVip() uint32
	GetWin() uint32
	GetLost() uint32
	SetVip(uint32)
	SetLost(uint32)
	SetWin(uint32)
	SetPing(uint32)
	GetPing() uint32
	ConverDataUser() *protocol.UserData
	ConverProtoUser() *protocol.ProtoUser
	GetInviteCode() string // 私人局邀请码
	GetRoomType() uint32   // 房间类型ID,对应房间表
	GetRoomID() uint32     // 比赛场或金币场房间id
	// 分别为：房间类型ID，房间号，房间邀请码
	SetRoom(uint32, uint32, uint32, string)
	ClearRoom()
	GetPlatform() uint32
	GetVipExpire() uint32
	SetVipExpire(uint32)
	GetChenmi() int32
	SetChenmi(int32)
	GetChenmiTime() uint32
	SetChenmiTime(uint32)
	GetSound() bool
	SetSound(bool)

	SetConn(IConn)
	GetConn() IConn
	Send(IProto)

	GetRoomCard() uint32
	SetRoomCard(uint32)

	GetBuild() string
	SetBuild(string)
	UserSave()
}
