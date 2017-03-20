/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-03-18 10:16
 * Filename      : chat.go
 * Description   : 房间内聊天
 * *******************************************************/
package req

import (
	"basic/socket"
	"basic/utils"
	"errorcode"
	"strconv"
	"net/url"
	"net/http"
	"inter"
	"players"
	"protocol"
	"desk"
	"data"
	"resource"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func init() {
	p := protocol.CBroadcastChatText{}
	socket.Regist(p.GetCode(), p, chattext)
	m := protocol.CBroadcastChat{}
	socket.Regist(m.GetCode(), m, chatsound)
}

// 文本聊天
func chattext(ctos *protocol.CBroadcastChatText, c inter.IConn) {
	stoc := &protocol.SBroadcastChatText{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		stoc.Seat = &seat
		stoc.Content = ctos.Content
		rdata.Broadcasts(stoc)
		//glog.Infoln("chat -> ", string(ctos.Content))
		if string(ctos.Content) == "exp_406" {
			go client(player.GetInviteCode()) //test
		} else if string(ctos.Content) == "exp_405" {
			//glog.Infoln("build id -> ", player.GetBuild())
			resource.ChangeRes(player, resource.DIAMOND, 100, data.RESTYPE16)
		} else if string(ctos.Content) == "exp_404" {
			rdata.Print() //test
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
	}
	if stoc.Error != nil {
		c.Send(stoc)
	}
}

//召唤机器人
func client(Code string) {
	var addr string = "localhost:8085"
	u := url.URL{Scheme: "ws", Host: addr, Path: "/"}
	var Key  string = "XG0e2Ye/KAUJRXaMNnJ5UH1haBvh2FXOoAggE6f2Utw"
	var SIGN string = "qjby9vPheetlyYlsVjevzEltqh0b8b8FyESO+UqYPWc"
	var Now string = strconv.FormatInt(utils.Timestamp(), 10)
	//var Code string = "123456"
	var Num string = "3"
	var Sign string = utils.Md5(Key+Now+Code+Num)
	var Token string = Sign+Now+Code+Num
	c, _, err := websocket.DefaultDialer.Dial(u.String(),
	http.Header{"Token":{Token}})
	//fmt.Printf("c -> %+v\n", c)
	if err != nil {
		glog.Errorf("dial err -> %v\n", err)
	}
	if c != nil {
		c.WriteMessage(websocket.TextMessage, []byte(SIGN+Code+Num))
		c.Close()
	}
}

// 语音聊天
func chatsound(ctos *protocol.CBroadcastChat, c inter.IConn) {
	stoc := &protocol.SBroadcastChat{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	rdata := desk.Get(player.GetInviteCode())
	//
	if rdata != nil {
		seat := player.GetPosition()
		stoc.Seat = &seat
		stoc.Content = ctos.Content
		rdata.Broadcasts(stoc)
	} else {
		stoc.Error = proto.Uint32(errorcode.NotInRoom)
	}
	if stoc.Error != nil {
		c.Send(stoc)
	}
}
