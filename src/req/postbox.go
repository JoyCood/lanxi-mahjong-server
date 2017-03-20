/**********************************************************
* Author        : Michael
* Email         : dolotech@163.com
* Last modified : 2016-06-11 16:24
* Filename      : postbox.go
* Description   : 邮箱
* *******************************************************/
package req

import (
	"basic/socket"
	"data"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"resource"

	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CPost{}
	socket.Regist(p.GetCode(), p, getAllPost)
	d := protocol.CDelPost{}
	socket.Regist(d.GetCode(), d, delPost)
	r := protocol.CDelReadPost{}
	socket.Regist(r.GetCode(), r, delReadPost)
	a := protocol.CDelAllPost{}
	socket.Regist(a.GetCode(), a, delAllPost)
	o := protocol.COpenAppendix{}
	socket.Regist(o.GetCode(), o, openAppendix)
	g := protocol.CReadPost{}
	socket.Regist(g.GetCode(), g, readPost)

}

func getAllPost(ctos *protocol.CPost, c inter.IConn) {
	stoc := &protocol.SPost{}
	data := &data.DataPostbox{Receiver: c.GetUserid()}
	list, err := data.ReadAll()
	if err != nil || len(list) == 0 {
		stoc.Error = proto.Uint32(errorcode.PostboxEmpty)
	} else {
		for _, v := range list {
			p := &protocol.PostBoxData{
				Id:           &v.Id,
				Senderuserid: &v.Sender,
				Content:      &v.Content,
				Title:        &v.Title,
				Appendixname: &v.Appendixname,
				Expire:       &v.Expire,
				Read:         &v.Read,
				Kind:         &v.Kind,
				Draw:         &v.Draw,
			}
			stoc.Data = append(stoc.Data, p)
		}
	}
	c.Send(stoc)

}

func delPost(ctos *protocol.CDelPost, c inter.IConn) {
	stoc := &protocol.SDelPost{Postid: ctos.Postid}
	if ctos.GetPostid() > 0 {
		database := &data.DataPostbox{Id: ctos.GetPostid(), Receiver: c.GetUserid()}
		err := database.Delete()
		if err != nil {
			stoc.Error = proto.Uint32(errorcode.PostNotExist)
		} else {
			stoc.Error = proto.Uint32(0)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.PostNotExist)
	}
	c.Send(stoc)
}
func delReadPost(ctos *protocol.CDelReadPost, c inter.IConn) {
	stoc := &protocol.SDelReadPost{}
	database := &data.DataPostbox{Receiver: c.GetUserid()}
	err := database.CleanupRead()
	if err != nil {
		stoc.Error = proto.Uint32(errorcode.PostNotExist)
	} else {
		stoc.Error = proto.Uint32(0)
	}

	c.Send(stoc)

}
func delAllPost(ctos *protocol.CDelAllPost, c inter.IConn) {
	stoc := &protocol.SDelAllPost{}
	database := &data.DataPostbox{Receiver: c.GetUserid()}
	err := database.Cleanup()
	if err != nil {
		stoc.Error = proto.Uint32(errorcode.PostNotExist)
	} else {
		stoc.Error = proto.Uint32(0)
	}
	c.Send(stoc)

}
func openAppendix(ctos *protocol.COpenAppendix, c inter.IConn) {
	stoc := &protocol.SOpenAppendix{}
	//TODO:优化
	player := players.Get(c.GetUserid())
	//
	if ctos.GetPostid() > 0 {
		database := &data.DataPostbox{Id: ctos.GetPostid(), Receiver: c.GetUserid()}
		widgetList, err := database.OpenAppendix()
		if err != nil {
			stoc.Error = proto.Uint32(errorcode.AppendixNotExist)
		} else {
			if len(widgetList) > 0 {
				d := &protocol.PostAppendixData{}
				d.Postid = &database.Id
				d.Name = &database.Appendixname
				stoc.Data = d
				m := make(map[uint32]int32)
				for _, v := range widgetList {
					widget := &protocol.WidgetData{}
					widget.Id = &v.Id
					widget.Count = &v.Count
					d.Widgets = append(d.Widgets, widget)
					m[v.Id] = int32(v.Count)
				}
				resource.ChangeMulti(player, m, data.RESTYPE9)
			} else {
				stoc.Error = proto.Uint32(errorcode.AppendixNotExist)
			}
		}

	} else {
		stoc.Error = proto.Uint32(errorcode.PostNotExist)

	}
	c.Send(stoc)

}

func readPost(ctos *protocol.CReadPost, c inter.IConn) {
	stoc := &protocol.SReadPost{}
	stoc.Id = ctos.Id
	if ctos.GetId() > 0 {
		database := &data.DataPostbox{Id: ctos.GetId(), Receiver: c.GetUserid()}
		err := database.ReadPost()
		if err != nil {
			stoc.Error = proto.Uint32(errorcode.PostNotExist)
		}
	} else {
		stoc.Error = proto.Uint32(errorcode.PostNotExist)
	}

	c.Send(stoc)
}
