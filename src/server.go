/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:02
 * Filename      : server.go
 * Description   : 服务器主文件
 * *******************************************************/
package main

import (
	"basic/imageserver"
	"basic/socket"
	"data"
	"desk"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"players"
	_ "net/http/pprof"
	"req"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/golang/glog"
)

func uploadFeedbackPhotoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	//上传参数为uploadfile
	userid := r.FormValue("userid")
	glog.Infoln(userid)
	if userid == "" {
		w.Write([]byte("Error:userid empty."))
		return
	}
	password := r.FormValue("password")
	feedbackContent := r.FormValue("feedback")
	if feedbackContent == "" {
		w.Write([]byte("Error:feedback content is empty."))
		return
	}
	glog.Infoln(feedbackContent)
	if password == "" {
		w.Write([]byte("Error:password empty."))
		return
	}
	user := &data.User{Userid: userid}
	if !user.PWDIsOK(password) {
		w.Write([]byte("Error:password or userid error."))
		return
	}
	file, _, err := r.FormFile("uploadfile")
	var imgid string
	if err == nil && file != nil {
		defer file.Close()
		imgid, err = imageserver.SaveImage(data.Conf.FeedbackDir, file)
		if err != nil {
			w.Write([]byte("Error:save iamge file error."))
			return
		}
		//	gossdb.C().Hset(data.KEY_USER+user.Userid, "Photo", imgid)
	}
	fb := &data.DataFeedback{
		Userid:    userid,
		Content:   feedbackContent,
		ImagePath: imgid,
	}
	if err := fb.Save(); err != nil {
		w.Write([]byte("Error:save content error."))
		return
	}

	w.Write([]byte("ok"))
}
func uploadHeadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	//上传参数为uploadfile
	userid := r.FormValue("userid")
	glog.Infoln(userid)
	if userid == "" {
		w.Write([]byte("Error:userid empty."))
		return

	}
	password := r.FormValue("password")
	glog.Infoln(password)
	if password == "" {
		w.Write([]byte("Error:password empty."))
		return
	}
	user := &data.User{Userid: userid}
	if !user.PWDIsOK(password) {
		w.Write([]byte("Error:password or userid error."))
		return
	}
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		w.Write([]byte("Error:get iamge data error."))
		return
	}
	defer file.Close()

	imgid, err := imageserver.SaveImage(data.Conf.ImageDir, file)
	if err != nil {
		w.Write([]byte("Error:save iamge file error."))
		return
	}
	u := &data.User{Userid:user.Userid, Photo:imgid}
	u.UpdatePhoto()
	w.Write([]byte(imgid))
}

var (
	version = "0.0.1"
	ip      = ""
)

func main() {
	var config string
	flag.StringVar(&config, "conf", "./conf.json", "config path")
	flag.Parse()
	data.LoadConf(config)
	glog.Infoln("Config: ", data.Conf)
	defer glog.Flush()
	//logcenter.Connect(data.Conf.Log.Ip, data.Conf.Log.Port)
	glog.Infoln("逻辑服务器端口:", data.Conf.Port)
	imserver := imageserver.NewServer(data.Conf.ImageDir, data.Conf.ImagePort)
	imserver.HandleFunc("/", uploadHeadPhotoHandler).Methods("POST")
	go imserver.Run()
	glog.Infoln("头像服务器监听端口: ", data.Conf.ImagePort)
	imserver = imageserver.NewServer(data.Conf.FeedbackDir, data.Conf.FeedbackPort)
	imserver.HandleFunc("/", uploadFeedbackPhotoHandler).Methods("POST")
	go imserver.Run()
	glog.Infoln("反馈服务器监听端口: ", data.Conf.FeedbackPort)
	//go imageserver.NewServer(data.Conf.FeedbackDir, data.Conf.FeedbackPort).Run()
	go pprof()
	go payRecvServe()
	go req.WebServer()
	req.IappPayInit()
	req.WxPayInit()
	req.WxLoginInit()
	data.InitIDGen()
	//go crontab()
	ln, lnCh := socket.Server(":" + strconv.Itoa(data.Conf.Port))
	glog.Infoln("version:", version)
	glog.Infoln("Server listening on", data.Conf.Port)
	glog.Infoln("Server started at", ln.Addr())
	signalProc(ln, lnCh)
}

func pprof() {
	if data.Conf.Pprof > 0 {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(data.Conf.Pprof), nil)
		glog.Infoln("性能监控端口:", data.Conf.Pprof)
		if err != nil {
			glog.Fatal("ListenAndServe error: ", err)
		}
	}
}

// 支付回调监听服务
func payRecvServe() {
	if data.Conf.PayPort > 0 {
		//go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(data.Conf.PayPort), nil)
		err := http.ListenAndServeTLS("0.0.0.0:"+strconv.Itoa(data.Conf.PayPort), "./cert.pem", "./key.pem", nil)
		glog.Infoln("支付监控端口:", data.Conf.PayPort)
		glog.Infoln("pay recv serve strat:", data.Conf.PayPort)
		if err != nil {
			glog.Fatal("ListenAndServe error: ", err)
		}
	}
}

func signalProc(ln net.Listener, lnCh chan error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln(string(debug.Stack()))
		}
	}()
	ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGUSR1, syscall.SIGUSR2)
	//signal.Notify(ch, syscall.SIGHUP)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGHUP) //监听SIGINT和SIGKILL信号
	glog.Infoln("signalProc ... ")
	for {
		msg := <-ch
		switch msg {
		default:
			//先关闭监听服务
			ln.Close()
			glog.Infoln(<-lnCh)
			//关闭连接
			socket.Close()
			//关闭服务
			desk.Close()
			players.Close()
			//延迟退出，等待连接关闭，数据回存
			glog.Infof("get sig -> %v\n", msg)
			<-time.After(10 * time.Second)
			return
		case syscall.SIGHUP:

		}
	}
}
