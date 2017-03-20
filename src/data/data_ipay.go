package data

/*
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/89hmdys/toast/crypto"
	"github.com/89hmdys/toast/rsa"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"basic/utils"
)

// 回调url
var notifyCpUrl string = "http://" + Conf.Ipayhost + ":" + strconv.Itoa(Conf.Ipay)

var NotifyClUrl string = "/mahjong/iapppay/notice"

// 回调通知url
var NOTIFYURL string = notifyCpUrl + NotifyClUrl

//爱贝商户后台接入url
const iapppayCpUrl = "http://ipay.iapppay.com:9999"

//下单接口 url
const ORDERURL = iapppayCpUrl + "/payapi/order"

//支付结果查询接口 url
const queryResultUrl = iapppayCpUrl + "/payapi/queryresult"

//应用编号
const APPID = "3006208675"


// 交易结果通知
type TradingResults struct {
	Transtype int     `json:"transtype" bson:"Transtype"` // 交易类型0–支付交易；
	Cporderid string  `json:"cporderid" bson:"_id"` // 商户订单号
	Transid   string  `json:"transid" bson:"Transid"`   // 交易流水号
	Appuserid string  `json:"appuserid" bson:"Appuserid"` // 用户在商户应
	Appid     string  `json:"appid" bson:"Appid"`     // 游戏id
	Waresid   uint32  `json:"waresid" bson:"Waresid"`   // 商品编码
	Feetype   int     `json:"feetype" bson:"Feetype"`   // 计费方式
	Money     float32 `json:"money" bson:"Money"`     // 交易金额
	Currency  string  `json:"currency" bson:"Currency"`  // 货币类型	RMB
	Result    int     `json:"result" bson:"Result"`    // 交易结果	0–交易成功 1–交易失败
	Transtime string  `json:"transtime" bson:"Transtime"` // 交易完成时间 yyyy-mm-dd hh24:mi:ss
	Cpprivate string  `json:"cpprivate" bson:"Cpprivate"` // 商户私有信息
	Paytype   uint32  `json:"paytype" bson:"Paytype"`   // 支付方式
	Parent    string  `json:"paytype" bson:"Parent"`// 绑定的父级代理商游戏ID
}

// 下单请求参数
type IapppayOrder struct {
	Appid         string `json:"appid"  bson:"Appid"`         // 应用编号
	Waresid       uint32 `json:"waresid" bson:"Waresid"`       // 商品编号
	Waresname     string `json:"waresname" bson:"Waresname"`     // 商品名称
	Cporderid     string `json:"cporderid" bson:"Cporderid"`     // 商户订单号
	Price         uint32 `json:"price" bson:"Price"`         // 支付金额
	Currency      string `json:"currency" bson:"Currency"`      // 货币类型
	Appuserid     string `json:"appuserid" bson:"Appuserid"`     // 用户在商户应用的唯一标识
	Cpprivateinfo string `json:"cpprivateinfo" bson:"Cpprivateinfo"` // 商户私有信息
	Notifyurl     string `json:"notifyurl" bson:"Notifyurl"`     // 支付结果通知地址
}

// 下单请求失败结果
type TransData struct {
	Code   json.Number `json:"code,Number" bson:"Code"`
	Errmsg string      `json:"errmsg" bson:"Errmsg"`
}

// 下单请求成功结果
type TransDataId struct {
	Transid string `json:"transid" bson:"Transid"`
}

// 主动请求交易结果请求参数
type IapppayQuery struct {
	Appid     string `json:"appid" bson:"Appid"`     // 应用编号
	Cporderid string `json:"cporderid" bson:"Cporderid"` // 商户订单号
}

// 主动请求交易结果
type QueryResults struct {
	Cporderid string  `json:"cporderid" bson:"Cporderid"` // 商户订单号
	Transid   string  `json:"transid" bson:"Transid"`   // 交易流水号
	Appuserid string  `json:"appuserid" bson:"Appuserid"` // 用户在商户应
	Appid     string  `json:"appid" bson:"Appid"`     // 游戏id
	Waresid   uint32  `json:"waresid" bson:"Waresid"`   // 商品编码
	Feetype   int     `json:"feetype" bson:"Feetype"`   // 计费方式
	Money     float32 `json:"money" bson:"Money"`     // 交易金额
	Currency  string  `json:"currency" bson:"Currency"`  // 货币类型	RMB
	Result    int     `json:"result" bson:"Result"`    // 交易结果	0–交易成功 1–交易失败
	Transtime string  `json:"transtime" bson:"Transtime"` // 交易完成时间 yyyy-mm-dd hh24:mi:ss
	Cpprivate string  `json:"cpprivate" bson:"Cpprivate"` // 商户私有信息
	Paytype   uint32  `json:"paytype" bson:"Paytype"`   // 支付方式
}

// 交易记录
type ChargeOrder struct {
	Orderid   string `bson:"_id"`// 订单号
	Userid    string `bson:"Userid"` // userid
	Phone     string `bson:"Phone"`// phone
	Transid   string `bson:"Transid"`// 流水号
	Waresid   uint32 `bson:"Waresid"` // 商品编号
	Money     uint32 `bson:"Money"`// 交易金额
	Platform  uint32 `bson:"Platform"` // 平台
	OrderRes  uint32 `bson:"OrderRes"` // 下单结果 0成功,1失败
	Ctime     int64  `bson:"Ctime"` // 创建时间(下单时间)
	Result    uint32 `bson:"Result"`// 交易结果 0成功,1失败
	Transtime string `bson:"Transtime"`// 交易完成时间
	Parent    string `bson:"Parent"`// 绑定的父级代理商游戏ID
	Status    uint32 `bson:"Status"`// 发货结果 0成功,1失败
}

// 交易记录
type TradeRecord struct {
	Id        string `json:"id" bson:"_id"` //商户订单号(游戏内自定义订单号)
	Transid   string `json:"transid" bson:"Transid"` //交易流水号(计费支付平台的交易流水号,微信订单号)
	Userid    string `json:"userid" bson:"Userid"` //用户在商户应用的唯一标识(userid)
	//--
	Itemid    string `json:"itemid" bson:"Itemid"`   //购买商品ID
	Amount    string `json:"amount" bson:"Amount"`   //购买商品数量
	Diamond   uint32 `json:"diamond" bson:"Diamond"` //购买钻石数量
	Money     uint32 `json:"money" bson:"Money"`     //交易总金额(单位为分)
	Ctime     string `json:"ctime" bson:"Ctime"`     //订单生成时间 yyyy-mm-dd hh24:mi:ss
	Transtime string `json:"transtime" bson:"Transtime"` //交易完成时间 yyyy-mm-dd hh24:mi:ss
	Result    int    `json:"result" bson:"Result"`    //交易结果(0–交易成功,1–交易失败,2-交易中,3-发货中)
	//--iapppay
	Appid     string `json:"appid" bson:"Appid"`         //游戏id(平台为商户应用分配的唯一代码)
	Waresid   uint32 `json:"waresid" bson:"Waresid"`     //商品编码(平台为应用内需计费商品分配的编码)
	Currency  string `json:"currency" bson:"Currency"`   //货币类型(RMB,CNY)
	Transtype int    `json:"transtype" bson:"Transtype"` //交易类型(0–支付交易)
	Feetype   int    `json:"feetype" bson:"Feetype"`     //计费方式(表示商品采用的计费方式)
	Paytype   uint32 `json:"paytype" bson:"Paytype"`     //支付方式(表示用户采用的支付方式,403-微信支付)
	//--
	Clientip  string `json:"clientip" bson:"Clientip"`         //客户端ip
	Parent    string `json:"paytype" bson:"Parent"`//绑定的父级代理商游戏ID
}

var cipher rsa.Cipher

// 接收交易结果通知
func init() {
	key, err := rsa.LoadKeyFromPEMFile(
		`rsa_public_key.pem`,
		`rsa_private_key.pem`,
		rsa.ParsePKCS1Key)
	if err != nil {
		glog.Errorln(err)
		return
	}

	cip, err := crypto.NewRSA(key)
	if err != nil {
		glog.Errorln(err)
		return
	}
	cipher = cip
}


// todo 对已绑定上级的玩家充值时累计记录该玩家的父级和祖父级de周期代理返利
func (this *ChargeOrder) Settle(gameid string,money uint32) error {
	users:= Agent_Users{}
	users.GetParent2(gameid)
	for _,v:=range  users{
		if v.Lv == 1{
			rebate:=RebateSettleData{Gameid:v.Gameid,Money:uint32(float32(money) * 0.4)}
			rebate.Inc()
		}else if v.Lv == 2{
			rebate:=RebateSettleData{Gameid:v.Gameid,Money:uint32(float32(money) * 0.06)}
			rebate.Inc()
		}
	}

	return nil
}

// 发送post报文
func IpayRequest(body string) ([]byte, error) {
	r, err := http.Post(ORDERURL,
		"application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		glog.Errorln("order request err: ", err)
		return nil, err
	}
	result, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		glog.Errorln("order response err: ", err)
		return nil, err
	}
	glog.Errorln("result:", string(result))
	return result, err
}

// 组装request报文
func ComposeReq(i *IapppayOrder) (string, error) {
	tran, err := json.Marshal(i)
	glog.Infoln("IapppayOrder:", i, "tran:", string(tran))
	if err != nil {
		glog.Infoln("json err:", err)
		return "", err
	}
	sign, err := IpaySign(tran)
	if err != nil {
		glog.Errorln("sign err:", err)
		return "", err
	}
	t := url.QueryEscape(string(tran))
	s := url.QueryEscape(base64.StdEncoding.EncodeToString(sign))
	b := fmt.Sprintf("transdata=%s&sign=%s&signtype=RSA", t, s)
	return b, nil
}

// 解析response报文, 返回transid
// transdata={"transid":"11111"}&sign=xxxxxx&signtype=RSA
// transdata={"code":"1001","errmsg":"签名验证失败"}
func ParseResp(result []byte) string {
	m := ParseResponse(result)
	if ParseVerify(m) {
		var transid TransDataId
		err := json.Unmarshal([]byte(m["transdata"]), &transid)
		if err == nil {
			return transid.Transid
		}
	}
	var trans TransData
	err := json.Unmarshal([]byte(m["transdata"]), &trans)
	if err == nil {
		glog.Infoln("trans:", trans, "Code:", string(trans.Code))
	}
	return ""
}

func ParseVerify(m map[string]string) bool {
	if sign, ok := m["sign"]; ok {
		bufs, err := base64.StdEncoding.DecodeString(sign)
		if err != nil {
			glog.Errorln("Decode(%q) failed: %v", sign, err)
			return false
		}
		err = IpayVerify([]byte(m["transdata"]), bufs)
		if err == nil && m["signtype"] == "RSA" {
			return true
		}
	}
	return false
}

func ParseResponse(result []byte) map[string]string {
	m := make(map[string]string)
	r := strings.Split(string(result), "&")
	for _, v := range r {
		h := strings.Split(v, "=")
		// m[h[0]] = h[1]
		arg, err := url.QueryUnescape(h[1])
		if err != nil {
			glog.Errorln("QueryUnescape err : ", err)
			m[h[0]] = ""
			continue
		}
		m[h[0]] = arg
	}
	return m
}

// RSA签名
func IpaySign(tran []byte) ([]byte, error) {
	return cipher.Sign(tran)
}

// RSA验签
func IpayVerify(tran []byte, sign []byte) error {
	return cipher.Verify(tran, sign)
}

// begin 主动请求 发送post报文
func IpayQuery(body string) ([]byte, error) {
	r, err := http.Post(queryResultUrl,
		"application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		glog.Errorln("Query request err: ", err)
		return nil, err
	}
	result, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		glog.Errorln("Query response err: ", err)
		return nil, err
	}
	glog.Errorln("result:", string(result))
	return result, err
}

// 主动请求 组装request报文
func ComposeQueryReq(i *IapppayQuery) (string, error) {
	tran, err := json.Marshal(i)
	glog.Infoln("IapppayQuery:", i, "tran:", string(tran))
	if err != nil {
		glog.Infoln("Query json err:", err)
		return "", err
	}
	sign, err := IpaySign(tran)
	if err != nil {
		glog.Errorln("Query sign err:", err)
		return "", err
	}
	t := url.QueryEscape(string(tran))
	s := url.QueryEscape(base64.StdEncoding.EncodeToString(sign))
	b := fmt.Sprintf("transdata=%s&sign=%s&signtype=RSA", t, s)
	return b, nil
}

// 解析response报文, 返回transid
// transdata={"transid":"11111"}&sign=xxxxxx&signtype=RSA
// transdata={"code":"1001","errmsg":"签名验证失败"}
func ParseQueryResp(result []byte) string {
	m := ParseResponse(result)
	if ParseVerify(m) {
		// glog.Infoln("m[transdata]:", m["transdata"])
		var message TradingResults
		err := json.Unmarshal([]byte(m["transdata"]), &message)
		// glog.Infoln("message:", message, "err:", err)
		if err == nil {
			message.Save()
			return ""
		}
		var message2 QueryResults
		err = json.Unmarshal([]byte(m["transdata"]), &message2)
		// glog.Infoln("message2:", message2, "err:", err)
		if err == nil {
			// message2.Save()
			return ""
		}
		// TODO:检测是否发货
	}
	var trans TransData
	err := json.Unmarshal([]byte(m["transdata"]), &trans)
	if err == nil {
		glog.Infoln("trans:", trans, "Code:", string(trans.Code))
	}
	return ""
}

// 检测是否发货
func CheckDelivery(transid, cporderid string) bool {
	d := &ChargeOrder{Orderid: cporderid, Transid: transid}
	err := d.Get()
	if err == nil && d.Result == 0 && d.Transtime != "" {
		return true // 已经发货
	}
	return false // 没有发货
}

// 交易结果主动查询
// cporderid = data.ChargeOrder.Orderid,
// cporderid: 89, Transid: 32471610191058365861
func QueryResult(cporderid, userid string) {
	s := &IapppayQuery{
		Appid:     APPID,     // 应用编号
		Cporderid: cporderid, // 商户订单号
	}
	body, err := ComposeQueryReq(s)
	if err != nil {
		glog.Infoln("body err:", err)
		// return ""
	}
	result, err := IpayQuery(body)
	if err != nil {
		glog.Infoln("request err:", err)
		// return ""
	}
	Transid := ParseQueryResp(result)
	glog.Infoln("cporderid:", cporderid, "userid:", userid, "Transid:", Transid)
}


func ChargeOrderLog(Phone string, Platform uint32, Status uint32, t *TradingResults)error {
	return C(_CHARGEORDER).UpdateId(t.Cporderid,bson.M{"$set":bson.M{
		"Phone":Phone,
		"Platform":Platform,
		"Result":t.Result,		// 交易结果 0成功,1失败
		"Transtime":t.Transtime,		// 交易完成时间
		"Status":Status,		// 发货结果 0成功,1失败
	}})
}

//
func GenCporderid(userid string) string {
	return utils.Base62encode(uint64(time.Now().Unix())) + userid
}

// 交易记录
func (this *ChargeOrder) Exist() (bool, error) {
	count,err:=C(_CHARGEORDER).FindId(this.Orderid).Count()
	return count > 0,err
}

func (this *ChargeOrder) Get() error {
	return C(_CHARGEORDER).FindId(this.Orderid).One(this)
}

func (this *ChargeOrder) Save() error {
	return C(_CHARGEORDER).Insert(this)

}

// 交易结果通知记录
func (this *TradingResults) Get() error {
	return C(_TRADINGRESULTS).FindId(this.Cporderid).One(this)
}

func (this *TradingResults) Save() error {
	return C(_TRADINGRESULTS).Insert(this)
}

// 获取某玩家的所有离线订单，用于上线补单
func GetTradingOff(userid string) ([]*TradingResults, error) {
	list := make([]*TradingResults, 0)
	err:= C(_TRADINGOFFLINE).Find(bson.M{"Appuserid":userid}).All(&list)
	return list,err
}

// 保存离线订单,用于下次上线补单
func (this *TradingResults) SaveTradingOff() error {
	return C(_TRADINGOFFLINE).Insert(this)
}

// 补单完成，删除订单
func DelTradingOff(userid string) error {
	return C(_TRADINGOFFLINE).Remove(bson.M{"Appuserid":userid})

}
*/
