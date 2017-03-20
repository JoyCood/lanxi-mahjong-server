/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:07
 * Filename      : user.go
 * Description : 单个玩家的数据
 * *******************************************************/
package desk

import (
	//"algo"
	"data"
	"inter"
	"protocol"
	"runtime/debug"
	"sync"
	"time"

	"github.com/golang/glog"
)

func NewPlayer(data *data.User) *Player {
	return &Player{
		//client: algo.NewClient(),
		user:   data,
	}
}

//用户的全部数据，服务器存取专用
type Player struct {
	position uint32        // 玩家座位
	//client   inter.IClient // 牌型数据
	//room     inter.IRoom   // 房间数据 用于机器人绑定房间
	ready    bool
	sync.RWMutex
	user    *data.User
	timeout *time.Timer // 碰杠胡和出牌超时计时
	conn    inter.IConn

	inviteCode string
	roomType   uint32
	roomID     uint32 // 比赛场或私人局的房间ID

}

func (this *Player) UserSave() {
	this.user.Save()
}

func (this *Player) Send(value inter.IProto) {
	this.RLock()
	defer this.RUnlock()
	this.conn.Send(value)
}
func (this *Player) GetConn() inter.IConn {
	this.RLock()
	defer this.RUnlock()
	return this.conn
}
func (this *Player) SetConn(value inter.IConn) {
	this.Lock()
	defer this.Unlock()
	this.conn = value
}

// ---------------------------房间属性---------------------------------
func (this *Player) ClearRoom() {
	this.inviteCode = ""
	this.roomType = 0
	this.roomID = 0
	this.position = 0
}

// 分别为：房间类型ID，房间号，房间邀请码
func (this *Player) SetRoom(rtype uint32, rid uint32, seat uint32, invitecode string) {
	this.inviteCode = invitecode
	this.roomType = rtype
	this.roomID = rid
	this.position = seat
}

func (this *Player) GetInviteCode() string {
	return this.inviteCode
}

func (this *Player) GetRoomID() uint32 {
	return this.roomID
}

func (this *Player) GetRoomType() uint32 {
	return this.roomType
}

//---------------------------------------------------------------------------
func (this *Player) StopTime() {
	if this.timeout != nil {
		this.timeout.Stop()
		this.timeout = nil
	}
}

func (this *Player) StartTime(f func(inter.IPlayer, byte), card byte, t int) {
	this.StopTime()
	d := time.Second * time.Duration(t)
	this.timeout = time.AfterFunc(d, func() {
		defer func() {
			if e := recover(); e != nil {
				glog.Errorln(string(debug.Stack()))
			}
		}()
		this.StopTime()

		f(this, card)
	})
}

func (this *Player) GetSound() bool      { return this.user.Sound }
func (this *Player) SetSound(value bool) { this.user.Sound = value }

func (this *Player) GetBuild() string      { return this.user.Build }
func (this *Player) SetBuild(id string) { this.user.Build = id }

func (this *Player) GetPing() uint32      { return this.user.Ping }
func (this *Player) SetPing(value uint32) { this.user.Ping = value }

func (this *Player) GetLost() uint32      { return this.user.Lost }
func (this *Player) SetLost(value uint32) { this.user.Lost = value }

func (this *Player) GetWin() uint32      { return this.user.Win }
func (this *Player) SetWin(value uint32) { this.user.Win = value }

func (this *Player) GetVip() uint32      { return this.user.Vip }
func (this *Player) SetVip(value uint32) { this.user.Vip = value }

func (this *Player) GetVipExpire() uint32      { return this.user.VipExpire }
func (this *Player) SetVipExpire(value uint32) { this.user.VipExpire = value }

func (this *Player) GetChenmi() int32      { return this.user.Chenmi }
func (this *Player) SetChenmi(value int32) { this.user.Chenmi = value }

func (this *Player) GetChenmiTime() uint32      { return this.user.ChenmiTime }
func (this *Player) SetChenmiTime(value uint32) { this.user.ChenmiTime = value }

func (this *Player) GetUserid() string   { return this.user.Userid }
func (this *Player) GetPosition() uint32 { return this.position }

func (this *Player) GetNickname() string { return this.user.Nickname }

func (this *Player) GetSex() uint32 { return this.user.Sex }

func (this *Player) GetSign() string { return this.user.Sign }

func (this *Player) GetEmail() string { return this.user.Email }

func (this *Player) GetPhone() string { return this.user.Phone }

func (this *Player) GetAuth() string { return this.user.Auth }

func (this *Player) GetPwd() string { return this.user.Pwd }

func (this *Player) GetBirth() uint32 { return this.user.Birth }

func (this *Player) GetIP() uint32 { return this.user.Create_ip }

func (this *Player) GetTime() uint32 { return this.user.Create_time }

func (this *Player) GetCoin() uint32 { return this.user.Coin }

func (this *Player) GetExp() uint32 { return this.user.Exp }

func (this *Player) GetTerminal() string { return this.user.Terminal }

func (this *Player) GetStatus() uint32 { return this.user.Status }

func (this *Player) GetAddress() string { return this.user.Address }

func (this *Player) GetPhoto() string { return this.user.Photo }

//func (this *Player) GetClient() inter.IClient { return this.client }

func (this *Player) SetUserid(id string) { this.user.Userid = id }

func (this *Player) SetPosition(pos uint32) {
	this.position = pos
}

func (this *Player) SetNickname(nick string) {
	this.user.Nickname = nick
	this.user.UpdateNickname()
}

func (this *Player) SetSex(sex uint32) {
	this.user.Sex = sex
	this.user.UpdateSex()
}

func (this *Player) SetSign(sign string) { this.user.Sign = sign }

func (this *Player) SetEmail(email string) { this.user.Email = email }

func (this *Player) SetPhone(phone string) { this.user.Phone = phone }

func (this *Player) SetAuth(auth string) { this.user.Auth = auth }

func (this *Player) SetPwd(pwd string) { this.user.Pwd = pwd }

func (this *Player) SetBirth(birth uint32) { this.user.Birth = birth }

func (this *Player) SetIP(ip uint32) { this.user.Create_ip = ip }

func (this *Player) SetTime(time uint32) { this.user.Create_time = time }

func (this *Player) SetCoin(coin uint32) { this.user.Coin = coin }

func (this *Player) SetExp(exp uint32) { this.user.Exp = exp }

func (this *Player) SetDiamond(value uint32)  { this.user.Diamond = value }
func (this *Player) SetExchange(value uint32) { this.user.Exchange = value }
func (this *Player) SetRoomCard(value uint32) { this.user.RoomCard = value }
func (this *Player) SetTicket(value uint32)   { this.user.Ticket = value }
func (this *Player) GetDiamond() uint32       { return this.user.Diamond }
func (this *Player) GetExchange() uint32      { return this.user.Exchange }
func (this *Player) GetRoomCard() uint32      { return this.user.RoomCard }
func (this *Player) GetTicket() uint32        { return this.user.Ticket }

func (this *Player) SetTerminal(termianl string) { this.user.Terminal = termianl }

func (this *Player) SetStatus(status uint32) { this.user.Status = status }

func (this *Player) SetAddress(address string) { this.user.Address = address }

func (this *Player) SetPhoto(photo string) { this.user.Photo = photo }

func (this *Player) SetQQ(qq string)          { this.user.Qq_uid = qq }
func (this *Player) SetWechat(wechat string)  { this.user.Wechat_uid = wechat }
func (this *Player) GetWechat() string        { return this.user.Wechat_uid }
func (this *Player) GetPlatform() uint32      { return this.user.Platform }
func (this *Player) SetMicrobolg(blog string) { this.user.Microblog_uid = blog }

func (this *Player) SetReady(value bool) { this.ready = value }

func (this *Player) GetReady() bool { return this.ready }

func (this *Player) ConverProtoUser() *protocol.ProtoUser {
	return &protocol.ProtoUser{
		Userid:   &this.user.Userid,
		Position: &this.position,
		Nickname: &this.user.Nickname,
		Sex:      &this.user.Sex,
		Exp:      &this.user.Exp,
		Photo:    &this.user.Photo,
		Coin:     &this.user.Coin,
		Address:  &this.user.Address,
		Terminal: &this.user.Terminal,
		Email:    &this.user.Email,
		Lost:     &this.user.Lost,
		Win:      &this.user.Win,
		Ping:     &this.user.Ping,
		Vip:      &this.user.Vip,
		Ready:    &this.ready,
		Platform: &this.user.Platform,
	}
}

func (this *Player) ConverDataUser() *protocol.UserData {
	online := true
	//var roomid uint32 = 0
	//if this.room != nil {
	//	roomid = this.room.GetId()
	//}
	var roomid uint32 = this.roomID
	return &protocol.UserData{
		Userid:     &this.user.Userid,
		Nickname:   &this.user.Nickname,
		Sex:        &this.user.Sex,
		Exp:        &this.user.Exp,
		Photo:      &this.user.Photo,
		Status:     &this.user.Status,
		Online:     &online,
		Phone:      &this.user.Phone,
		Coin:       &this.user.Coin,
		Diamond:    &this.user.Diamond,
		Exchange:   &this.user.Exchange,
		Ticket:     &this.user.Ticket,
		Address:    &this.user.Address,
		Terminal:   &this.user.Terminal,
		Email:      &this.user.Email,
		Lost:       &this.user.Lost,
		Win:        &this.user.Win,
		Ping:       &this.user.Ping,
		Vip:        &this.user.Vip,
		Ip:         &this.user.Create_ip,
		Birth:      &this.user.Birth,
		Sign:       &this.user.Sign,
		Roomid:     &roomid,
		Createtime: &this.user.Create_time,
		Platform:   &this.user.Platform,
		Sound:      &this.user.Sound,
		Roomcard:   &this.user.RoomCard,
		Build:      &this.user.Build,
	}
}
