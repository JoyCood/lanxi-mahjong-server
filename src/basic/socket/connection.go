/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 12:15
 * Filename      : connection.go
 * Description   : 封装了每个玩家连接数据结构,负责对玩家的数据发送和接收
 * *******************************************************/
package socket

import (
	"basic/event"
	"inter"
	"runtime/debug"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

const (
	// 网络掉线事件
	OFFLINE = "offline"
)

func newConnection(ip uint32, socket *websocket.Conn) *Connection {
	c := &Connection{
		writeChan: make(chan inter.IProto, 32),
		ws:        socket,
		ReadChan:  make(chan *Packet, 32),
		ipAddr:    ip,
		connected: true,
		closeChan: make(chan bool, 1),
	}
	return c
}

type Connection struct {
	writeChan        chan inter.IProto
	userid           string        // 玩家ID
	logined          bool          // true 标示已登录
	connected        bool          // false标示连接断开
	ws               *websocket.Conn
	ReadChan         chan *Packet
	closeChan        chan bool
	ipAddr           uint32 // 当前连得IP地址
	event.Dispatcher        // 事件管理器
	count            uint32
}


func (c *Connection) GetConnected() bool {
	return c.connected
}
func (c *Connection) GetIPAddr() uint32 {
	return c.ipAddr
}
func (c *Connection) SetLogin() {
	c.logined = true
}
func (c *Connection) GetLogin() bool {
	return c.logined
}

func (c *Connection) SetUserid(userid string) {
	c.userid = userid
}
func (c *Connection) GetUserid() string {
	return c.userid
}
func (c *Connection) Close() {
	c.ws.Close()
}

func (c *Connection) LoginTimeout() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}
	}()
	//建立连接后一定时间没有登录断开连接
	select {
	case <-time.After(waitForLogin):
		if !c.logined {
			c.Close()
		}
	}
}

func (c *Connection) Reader(readChan chan *Packet) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}
	}()

	for {
		select {
		// 如果管道关闭则退出for循环，因为管道关闭不会阻塞导致for进入死循环
		case packet, ok := <-readChan:
			if !ok {
				return
			}
			//glog.Infoln(packet.GetProto())
			c.count++
			c.count = c.count % 256
			if c.count != packet.count {
				glog.Errorln("count error -> ", c.count, packet.count)
				return
			}
			proxyHandle(packet.GetProto(), packet.GetContent(), c)
		}
	}
}

func (c *Connection) ReadPump() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}
	}()

	defer func() {
		//c.ws.Close()
		c.Close() //
		c.connected = false
		close(c.writeChan)
		close(c.ReadChan)
		close(c.closeChan)
		logout(c)
		if c.logined {
			c.Dispatch(OFFLINE, c)
		}
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(
		func(string) error {
			c.ws.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0, HeaderLen+1)
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			glog.Errorln("ReadMessage error ", err)
			return
		}
		// tmpBuffer, err = Unpack(c.index, append(tmpBuffer, message...), c.ReadChan)
		tmpBuffer, err = Unpack(append(tmpBuffer, message...), c.ReadChan)
		if err != nil {
			glog.Errorln("Unpack error ", err)
			return
		}
	}
}

func (c *Connection) Send(data inter.IProto) {
	defer func() {
		if err := recover(); err != nil {
			c.Close() //
			glog.Errorln(string(debug.Stack()))
		}
	}()
	if c.connected {
		c.writeChan <- data
	} else {
	}
}
func (c *Connection) write(mt int, packet inter.IProto) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	msg, _ := proto.Marshal((proto.Message)(packet))
	if len(msg) > 0 {
		b := Pack(packet.GetCode(), msg, 0)
		return c.ws.WriteMessage(mt, b)
	} else {
		return c.ws.WriteMessage(mt, msg)
	}
}

func (c *Connection) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		//c.ws.Close()
		c.Close() //
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}
	}()

	for {
		select {
		// 如果管道关闭则退出for循环，因为管道关闭不会阻塞导致for进入死循环
		case proto, ok := <-c.writeChan:
			if !ok {
				c.write(websocket.CloseMessage, nil)
				return
			}
			if err := c.write(websocket.TextMessage, proto); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
