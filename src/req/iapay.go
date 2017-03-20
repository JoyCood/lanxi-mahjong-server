/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2017-02-17 15:18
 * Filename      : iapppay.go
 * Description   : 爱贝,微信支付
 * *******************************************************/
package req

import (
	"basic/socket"
	"basic/utils"
	"csv"
	"data"
	"errorcode"
	"inter"
	"players"
	"protocol"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	i := protocol.CIapppayOrder{}
	socket.Regist(i.GetCode(), i, iappOrder)
}

//爱贝支付下单
func iappOrder(ctos *protocol.CIapppayOrder, c inter.IConn) {
	stoc := &protocol.SIapppayOrder{Error: proto.Uint32(0)}
	var waresid uint32 = ctos.GetId()
	var transid string = ""
	var userid string = c.GetUserid()
	player:= players.Get(userid)
	var ipaddr uint32 = player.GetConn().GetIPAddr()
	var ip string = utils.InetTontoa(ipaddr).String()
	var parent string = player.GetBuild()
	for i := 0; i < 5; i++ { //失败重复尝试
		transid = ipayOrder(waresid, userid, ip, parent)
		if transid == "" {
			<-time.After(time.Duration(100) * time.Millisecond)
			glog.Infoln("iapp order fail:", waresid, userid)
			stoc.Error = proto.Uint32(errorcode.IpayOrderFail)
		} else {
			break
		}
	}
	stoc.Transid = &transid
	stoc.Id = &waresid
	c.Send(stoc)
}

// 下单
func ipayOrder(waresid uint32, userid, ip, parent string) string {
	d := csv.GetShop(waresid)
	if d.Paymenttype != 1 {
		return ""
	}
	var price uint32 = d.Price
	var diamond uint32 = d.Number
	var itemid string = strconv.FormatInt(int64(d.PropId), 10)
	var orderid string = data.GenCporderid(userid)
	transid, err := Iapay.Submit(price, orderid, userid, itemid)
	if err != nil {
		return ""
	}
	//transid,下单记录
	t := &data.TradeRecord{
		Id: orderid,
		Transid: transid,
		Userid: userid,
		Itemid: itemid,
		Amount: "1",
		Diamond: diamond,
		Ctime: time.Now().Unix(),
		Result: 2,
		Clientip: ip,
		Parent: parent,
	}
	err = t.Save()
	if err != nil {
		return ""
	}
	return orderid
}
