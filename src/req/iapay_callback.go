//iapp pay 爱贝支付
package req

import (
	"bytes"
	"data"
	"fmt"
	"iapppay"
	"inter"
	"io/ioutil"
	"net/http"
	"players"
	"strconv"
	"resource"

	"github.com/golang/glog"
)

var Iapay *iapppay.IappTrans //爱贝支付

func IappPayInit() {
	host := data.Conf.Host
	port := strconv.Itoa(data.Conf.PayPort)
	//pattern := "/mahjong/iapppay/notice"
	pattern := data.Conf.PayIappPattern
	notifyUrl := "https://"+host+":"+port+pattern
	cfg := &iapppay.IappConfig{
		AppId:          "3006208675",
		PublicKeyPath:  "./rsa_public_key.pem",
		PrivateKeyPath: "./rsa_private_key.pem",
		NotifyPattern:  pattern,
		NotifyUrl:      notifyUrl,
		PlaceOrderUrl:  "http://ipay.iapppay.com:9999/payapi/order",
		QueryOrderUrl:  "http://ipay.iapppay.com:9999/payapi/queryresult",
	}
	iappTrans, err := iapppay.NewIappTrans(cfg)
	if err != nil {
		panic(err)
	}
	Iapay = iappTrans
	go Iapay.RecvNotify(iappRecvTrade) //goroutine
}

// 接收交易结果通知
func iappRecvTrade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	var buf bytes.Buffer
	var res bool
	if r.Method == "POST" {
		if result, err := ioutil.ReadAll(r.Body); err == nil {
			tradeResult, err := Iapay.ParseTradeResult(result)
			if err == nil {
				res = true
				go iapayCallback(tradeResult) //发货
			} else {
				glog.Errorf("trade result err: %v", err)
			}
		}
	}
	r.Body.Close()
	if res {
		fmt.Fprintf(&buf, "SUCCESS")
	} else {
		fmt.Fprintf(&buf, "FAILURE")
	}
	w.Write(buf.Bytes())
}

func iapayCallback(t *iapppay.TradeResult) {
	tradeRecord := &data.TradeRecord{
		Id: t.Cporderid,
		Transid: t.Transid,
	}
	err := tradeRecord.Get()
	if err != nil {
		//订单不存在或其它
		glog.Errorf("not exist orderid %v, err:, %v", t, err)
		return
	}
	if tradeRecord.Result == 0 {
		//重复发货
		glog.Errorf("repeat resp %v, err:, %v", t, err)
		return
	}
	//更新记录
	tradeRecord.Transtime = t.Transtime
	tradeRecord.Waresid = t.Waresid
	tradeRecord.Currency = t.Currency
	tradeRecord.Transtype = t.Transtype
	tradeRecord.Feetype = t.Feetype
	tradeRecord.Paytype = t.Paytype
	tradeRecord.Money = uint32(t.Money * 100) //转换为分
	tradeRecord.Result = t.Result
	// 离线状态
	player := players.Get(t.Appuserid) //TODO:优化
	if player == nil {
		tradeRecord.Result = 3 //发货中
		tradeRecord.SaveTradeOff()
	}
	//交易成功
	if t.Result == 0 && player != nil {
		sendGoods(player, tradeRecord)
	}
	//update record
	err = tradeRecord.Update()
	if err != nil {
		glog.Errorf("tradeRecord:%v, err:%v", tradeRecord, err)
	}else{
		tradeRecord.Settle(tradeRecord.Userid,tradeRecord.Money)
	}
}

//发货
func sendGoods(player inter.IPlayer, t *data.TradeRecord) {
	propid, err := strconv.Atoi(t.Itemid)
	if err != nil {
		glog.Errorf("Send Goods: %v, err: %v", t, err)
		return
	}
	var count int32 = int32(t.Diamond)
	resource.ChangeRes(player, uint32(propid), count, data.RESTYPE6)
}

// 登录检测
func tradeOff(player inter.IPlayer) {
	list, err := data.GetTradeOff(player.GetUserid())
	if err == nil {
		for _, v := range list {
			sendGoods(player, v)
		}
		data.DelTradeOff(player.GetUserid())
	}
}
