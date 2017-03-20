/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:25
 * Filename      : proxy.go
 * Description   : socket协议路由
 * *******************************************************/
package socket

import (
	"reflect"
	"runtime/debug"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

type handler struct {
	f interface{}
	t reflect.Type
}

func Regist(c uint32, s interface{}, f interface{}) {
	if reflect.TypeOf(f).Kind() == reflect.Func {
		m[c] = &handler{f: f, t: reflect.TypeOf(s)}
	} else {
		glog.Errorln("must be function")
	}
}

var m map[uint32]*handler = make(map[uint32]*handler)

func proxyHandle(c uint32, b []byte, conn *Connection) {
	defer func() {
		if e := recover(); e != nil {
			glog.Errorln(c, string(debug.Stack()))
		}
	}()

	if h, ok := m[c]; ok && (conn.GetLogin() || c == 1000 || c == 1022 || c == 1024) {
		v := reflect.New(h.t)
		//glog.Infoln(c)
		if err := proto.Unmarshal(b, v.Interface().(proto.Message)); err == nil {
			reflect.ValueOf(h.f).Call([]reflect.Value{v, reflect.ValueOf(conn)})
		} else {
			glog.Errorln("protocol  unmarshal fail: ", c)
		}
	} else {
		glog.Errorln("protocol not regist:", c)
	}
}
