/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-06-11 16:18
 * Filename      : iapppay.go
 * Description   : 苹果支付
 * *******************************************************/
package req

/*
import (
	"basic/socket"
	"bytes"
	"crypto/tls"
	"data"
	"encoding/json"
	"errorcode"
	"errors"
	"inter"
	"io/ioutil"
	"net/http"
	"protocol"
	"strconv"
	"csv"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

// 苹果验证返回码对应信息
//Status 	Description
//0 	The receipt provided is valid.
//21000 	The App Store could not read the JSON object you provided.
//21002 	The data in the receipt-data property was malformed.
//21003 	The receipt could not be authenticated.
//21004 	The shared secret you provided does not match the shared secret on file for your account.Only returned for iOS 6 style transaction receipts for auto-renewable subscriptions.
//21005 	The receipt server is not currently available.
//21006 	This receipt is valid but the subscription has expired. When this status code is returned to your server, the receipt data is also decoded and returned as part of the response.Only returned for iOS 6 style transaction receipts for auto-renewable subscriptions.
//21007 	This receipt is a sandbox receipt, but it was sent to the production server.
//21008 	This receipt is a production receipt, but it was sent to the sandbox server.
const (
	//ApplePayUrl string = "https://buy.itunes.apple.com/verifyReceipt"     // 正式验证地址
	ApplePayUrl string = "https://sandbox.itunes.apple.com/verifyReceipt" //测试验证地址
)

// 苹果验证成功返回数据结构
//{"status":0, "environment":"Sandbox",
//"receipt":{"receipt_type":"ProductionSandbox", "adam_id":0, "app_item_id":0, "bundle_id":"gzmj01.ysgame.com", "application_version":"1.0", "download_id":0, "version_external_identifier":0, "receipt_creation_date":"2016-09-09 02:47:38 Etc/GMT", "receipt_creation_date_ms":"1473389258000", "receipt_creation_date_pst":"2016-09-08 19:47:38 America/Los_Angeles", "request_date":"2016-09-09 02:47:40 Etc/GMT", "request_date_ms":"1473389260393", "request_date_pst":"2016-09-08 19:47:40 America/Los_Angeles", "original_purchase_date":"2013-08-01 07:00:00 Etc/GMT", "original_purchase_date_ms":"1375340400000", "original_purchase_date_pst":"2013-08-01 00:00:00 America/Los_Angeles", "original_application_version":"1.0",
//"in_app":[
//{"quantity":"1", "product_id":"1", "transaction_id":"1000000234931388", "original_transaction_id":"1000000234931388", "purchase_date":"2016-09-09 02:47:38 Etc/GMT", "purchase_date_ms":"1473389258000", "purchase_date_pst":"2016-09-08 19:47:38 America/Los_Angeles", "original_purchase_date":"2016-09-09 02:47:38 Etc/GMT", "original_purchase_date_ms":"1473389258000", "original_purchase_date_pst":"2016-09-08 19:47:38 America/Los_Angeles", "is_trial_period":"false"}]}}
func init() {
	p := protocol.CApplePay{}
	socket.Regist(p.GetCode(), p, appleOrder)
}

type InAppReceipt struct {
	Quantity                   string `json:"quantity"`
	Product_id                 string `json:"product_id"`
	Transaction_id             string `json:"transaction_id"`
	Original_transaction_id    string `json:"original_transaction_id"`
	Purchase_date              string `json:"original_transaction_id"`
	Purchase_date_ms           string `json:"purchase_date_ms"`
	Purchase_date_pst          string `json:"purchase_date_pst"`
	Original_purchase_date     string `json:"original_purchase_date"`
	Original_purchase_date_ms  string `json:"original_purchase_date_ms"`
	original_purchase_date_pst string `json:"original_purchase_date_pst"`
	Is_trial_periodstring      string `json:"is_trial_period"`
}
type AppleReceipt struct {
	Receipt_type                 string `json:"receipt_type"`
	Adam_id                      int    `json:"adam_id"`
	App_item_id                  int    `json:"app_item_id"`
	Bundle_id                    string `json:"bundle_id"`
	Application_version          string `json:"application_version"`
	Download_id                  int    `json:"download_id"`
	Version_external_identifier  int    `json:"version_external_identifier"`
	Receipt_creation_date        string `json:"Receipt_creation_date"`
	Receipt_creation_date_ms     string `json:"receipt_creation_date_ms"`
	Receipt_creation_date_pst    string `json:"receipt_creation_date_pst"`
	Request_date                 string `json:"request_date"`
	Request_date_ms              string `json:"request_date_ms"`
	Request_date_pst             string `json:"request_date_pst"`
	Original_purchase_date       string `json:"original_purchase_date"`
	Original_purchase_date_ms    string `json:"original_purchase_date_ms"`
	Original_purchase_date_pst   string `json:"original_purchase_date_pst"`
	Original_application_version string `json:"original_application_version"`

	InApp []InAppReceipt `json:"in_app"`
}

// POST方式提交到苹果支付验证服务器的json数据结构
type AppleRequst struct {
	ReceiptData string `json:"receipt-data"`
}

// 苹果验证返回数据结构
type AppleTradingResult struct {
	Status      int          `json:"status"`
	Environment string       `json:"environment"`
	Receipt     AppleReceipt `json:"receipt"`
}

func appleOrder(ctos *protocol.CApplePay, c inter.IConn) {
	stoc := &protocol.SApplePay{}
	stoc.Id = ctos.Id
	t, err := applepayPost(ctos.GetReceipt())
	glog.Infoln(ctos.GetId(), ctos.GetReceipt(), t, err)
	if err != nil {
		stoc.Error = proto.Uint32(errorcode.AppleOrderFail)
	} else {
		if t.Status == 0 {
			for _, v := range t.Receipt.InApp {
				id, err := strconv.Atoi(v.Product_id)
				glog.Infoln("id:", id, "err:", err, "Product_id:", v.Product_id)
				var money float32
				if err == nil {
					d := csv.GetShop(uint32(id))
					money = float32(d.Price)
				}
				tracd := &data.TradingResults{
					Appuserid: c.GetUserid(),
					Cporderid: v.Transaction_id,
					Cpprivate: v.Product_id,
					Money: money,
					Currency: "RMB",
				}
				//发货
				ChargeCallBack(tracd)

			}
		} else {
			stoc.Error = proto.Uint32(errorcode.AppleOrderFail)
		}
	}
	c.Send(stoc)
}
func applepayPost(receipt string) (*AppleTradingResult, error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	appleRequst := &AppleRequst{ReceiptData: receipt}
	b, err := json.Marshal(appleRequst)
	if err != nil {
		glog.Infoln("json err:", err)
		return nil, errors.New("json marshall error")
	}

	content := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", ApplePayUrl, content)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		//解析返回的JSON数据
		ret := &AppleTradingResult{}
		glog.Errorln(string(data))
		err = json.Unmarshal(data, &ret)
		if err != nil {
			return nil, err
		}
		return ret, nil
	} else {
		return nil, err
	}
}
*/
