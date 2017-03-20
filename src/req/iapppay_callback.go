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
	"bytes"
	"csv"
	"data"
	"encoding/json"
	"fmt"
	"inter"
	"io/ioutil"
	"net/http"
	"players"
	"resource"
	"strconv"

	"github.com/golang/glog"
)

// 接收交易结果通知
//func init() {
//	http.Handle(data.NotifyClUrl, http.HandlerFunc(TradingResultsNotice))
//}

// 接收交易结果通知
func TradingResultsNotice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// w.WriteHeader(http.StatusInternalServerError)
	// fmt.Fprintf(w, "SUCCESS")
	var buf bytes.Buffer
	res := false
	//
	if r.Method == "POST" {
		if result, err := ioutil.ReadAll(r.Body); err == nil {
			m := data.ParseResponse(result)
			r.Body.Close()
			if data.ParseVerify(m) {
				//解析返回的JSON数据
				var message data.TradingResults
				err = json.Unmarshal([]byte(m["transdata"]), &message)
				glog.Infoln("m[transdata]:", m["transdata"], "message:", message, "err:", err)
				if err == nil && message.Result == 0 {
					ChargeCallBack(&message)
					res = true
				}
			}
		}
	}
	if res {
		fmt.Fprintf(&buf, "SUCCESS")
	} else {
		fmt.Fprintf(&buf, "FAILURE")
	}
	glog.Infoln("接收交易结果通知:", r.Method, res)

	w.Write(buf.Bytes())
}

// 发货
func ChargeCallBack(t *data.TradingResults) {
	//TODO:优化
	player := players.Get(t.Appuserid)
	//
	err := t.Get()
	glog.Infoln("TradingResults:", t, "err:", err)
	ok := data.CheckDelivery(t.Transid, t.Cporderid)
	if err != nil {
		if !ok {
			if player != nil {
				t.Parent = player.GetBuild()
				ChargeSend(player, t)
			} else {
				p := &data.User{Userid:t.Appuserid}
				err = p.Get()
				if err == nil {
					t.Parent = p.Build
				}
				// 离线状态
				glog.Infoln("TradingResults:", t)
				t.SaveTradingOff()
			}
			t.Save()
		} else {
			glog.Infoln("重复发货 TradingResults:", t)
		}
	} else {
		glog.Infoln("订单已经存在 TradingResults:", t)
	}
}

func ChargeSend(userdata inter.IPlayer, t *data.TradingResults) {
	id, err := strconv.Atoi(t.Cpprivate)

	glog.Infoln("id:", id, "err:", err, "Cpprivate:", t.Cpprivate)
	if err == nil {
		d := csv.GetShop(uint32(id))
		resource.ChangeRes(userdata, d.PropId, int32(d.Number), data.RESTYPE6)
		Phone := userdata.GetPhone()
		Platform := userdata.GetPlatform()
		data.ChargeOrderLog(Phone, Platform, uint32(0), t)
	}
}

// 登录检测
func CheckChargeOff(conn inter.IPlayer) {
	list, err := data.GetTradingOff(conn.GetUserid())
	if err == nil {
		for _, v := range list {
			ChargeSend(conn, v)
		}
		data.DelTradingOff(conn.GetUserid())
	}
}
*/
