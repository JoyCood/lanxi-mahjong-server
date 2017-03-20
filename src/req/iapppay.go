/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-06-11 16:18
 * Filename      : iapppay.go
 * Description   : 爱贝支付
 * *******************************************************/
package req

/*
import (
	"basic/socket"
	"basic/utils"
	"csv"
	"data"
	"errorcode"
	"inter"
	"protocol"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"players"
)

// func init() {
// 	p := protocol.CIapppayOrder{}
// 	socket.Regist(p.GetCode(), p, order)
// }

func order(ctos *protocol.CIapppayOrder, c inter.IConn) {
	stoc := &protocol.SIapppayOrder{Error: proto.Uint32(0)}
	waresid := ctos.GetId()

	player:= players.Get(c.GetUserid())
	var transid string = ""
	for i := 0; i < 5; i++ { //失败重复尝试
		transid = iapppayOrder(waresid, player)
		if transid == "" {
			<-time.After(time.Duration(100) * time.Millisecond)
			glog.Infoln("order fail:", waresid, c.GetUserid())
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
func iapppayOrder(waresid uint32,player inter.IPlayer) string {
	var transid string = ""
	d := csv.GetShop(waresid)
	if d.Paymenttype != 1 {
		return transid
	}

	cporderid := data.GenCporderid(player.GetUserid())
	var notifyurl string = "http://" + data.Conf.Ipayhost + ":" + strconv.Itoa(data.Conf.Ipay) + data.NotifyClUrl
	s := &data.IapppayOrder{
		Appid:     data.APPID,
		Waresid:   1,
		Waresname: "diamond",
		//Cporderid:     strconv.Itoa(int(cporderid)),
		Cporderid:     cporderid,
		Price:         d.Price,
		Currency:      "RMB",
		Appuserid:     player.GetUserid(),
		Cpprivateinfo: strconv.Itoa(int(d.Id)), // 商品id或其他
		// Notifyurl:     data.NOTIFYURL,
		Notifyurl: notifyurl,
	}
	body, err := data.ComposeReq(s)
	if err != nil {
		glog.Infoln("body err:", err)
		return ""
	}
	result, err := data.IpayRequest(body)
	if err != nil {
		glog.Infoln("request err:", err)
		return ""
	}
	Transid := data.ParseResp(result)
	LogChargeOrder(Transid, s,player)
	glog.Infoln("Transid:", Transid, "userid:", player.GetUserid())
	// go data.QueryResult("89", userid)
	// go data.QueryResult("32471610191058365861", userid)
	return s.Cporderid
}

// 交易记录
func LogChargeOrder(Transid string, s *data.IapppayOrder,player inter.IPlayer) {
	var OrderRes uint32 = 1
	if Transid != "" {
		OrderRes = 0
	}

	l := &data.ChargeOrder{
		Orderid:   s.Cporderid,                // 订单号
		Userid:    s.Appuserid,                // userid
		Transid:   Transid,                    // 流水号
		Waresid:   s.Waresid,                  // 商品编号
		Money:    s.Price, // 交易金额
		OrderRes:  OrderRes,                   // 下单结果 0成功,1失败
		Ctime:     utils.Timestamp(),          // 创建时间(下单时间)
		Result:    uint32(1),                  // 交易结果 0成功,1失败
		Transtime: "",                         // 交易完成时间
		Status:    uint32(1),                  // 发货结果 0成功,1失败
		Parent:player.GetBuild(),
	}
	l.Save()
	l.Settle(s.Appuserid, s.Price)
}
*/
