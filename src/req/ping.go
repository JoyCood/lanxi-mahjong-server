package req

import (
	"basic/socket"
	"inter"
	"protocol"

	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CPing{}
	socket.Regist(p.GetCode(), p, ping)
}

func ping(ctos *protocol.CPing, c inter.IConn) {
	stoc := &protocol.SPing{Error: proto.Uint32(0)}
	c.Send(stoc)
}
