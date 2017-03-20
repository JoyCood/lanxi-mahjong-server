/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-06-11 16:18
 * Filename      : config.go
 * Description   : 游戏全局的数据配置表
 * *******************************************************/
package req

import (
	"basic/socket"
	"data"
	"inter"
	"protocol"
	"strconv"

	log "github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	p := protocol.CConfig{}
	socket.Regist(p.GetCode(), p, config)
}

func config(ctos *protocol.CConfig, c inter.IConn) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	stoc := &protocol.SConfig{}
	stoc.Sys = &protocol.SysConfig{}
	url := "http://" + data.Conf.Host + ":" + strconv.Itoa(data.Conf.ImagePort)
	stoc.Sys.Imageserver = proto.String(url)
	url = "http://" + data.Conf.Host + ":" + strconv.Itoa(data.Conf.FeedbackPort)
	stoc.Sys.Feedbackserver = proto.String(url)
	stoc.Sys.Discardtimeout = proto.Uint32(uint32(data.Conf.DiscardTimeout))
	stoc.Sys.Version = proto.String(data.Conf.Version)
	stoc.Sys.Shareaddr = proto.String(data.Conf.ShareAddr)
	c.Send(stoc)
}
