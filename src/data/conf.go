package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Conf Config

type Config struct {
	DiscardTimeout     int `json:"discard_timeout"`     // 出牌超时时间
	OperateTimeout     int `json:"operate_timeout"`     // 碰杠胡牌超时时间
	TrusteeshipTimeout int `json:"trusteeship_timeout"` // 托管出牌超时时间
	//---
	ImageDir     string `json:"image_dir"`           //头像地址
	ImagePort    int    `json:"image_port"`          //头像端口
	FeedbackDir  string `json:"feedback_image_dir"`  //反馈地址
	FeedbackPort int    `json:"feedback_image_port"` //反馈端口
	//---
	Database       string `json:"db_addr"`          //数据库地址
	Host           string `json:"server_host"`      //服务器地址
	Port           int    `json:"server_port"`      //服务器端口
	Pprof          int    `json:"pprof_port"`       //性能监控端口
	WebPort        int    `json:"web_port"`         //后台调用端口
	PayPort        int    `json:"pay_port"`         //支付回调端口
	PayWxPattern   string `json:"pay_wx_pattern"`   //微信支付回调路径
	PayIappPattern string `json:"pay_iapp_pattern"` //爱贝支付回调路径
	ShareAddr      string `json:"share_addr"`       //分享地址
	//---
	ServerId uint64 `json:"server_id"` //服务器ID
	Version  string `json:"version"`   //版本号
}

func LoadConf(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
}
