package desk

import (
	"basic/utils"
	"inter"
	"strconv"
	"sync"
	"time"

	"github.com/golang/glog"
)

//全局
var (
	Rooms *SocialRooms
)

//牌桌列表结构
type SocialRooms struct {
	sync.RWMutex                //读写锁
	list map[string]inter.IDesk //牌桌列表
	closeCh chan bool           //关闭通道
}

//初始化列表
func init() {
	Rooms = &SocialRooms{
		list:    make(map[string]inter.IDesk),
		closeCh: make(chan bool, 1),
	}
	go rooms_ticker() //goroutine,定时清理
}

//生成一个牌桌邀请码,全列表中唯一
func GenInvitecode() (s string) {
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(int(utils.RandInt32N(10)))
	}
	if Exist(s) { //是否已经存在
		return GenInvitecode() //重复尝试,TODO:一定次数后放弃尝试
	}
	return
}

//添加一个私人局房间
func Add(key string, rdata inter.IDesk) {
	Rooms.Lock()
	defer Rooms.Unlock()
	Rooms.list[key] = rdata
}

//删除一个牌桌
func Del(key string) {
	Rooms.Lock()
	defer Rooms.Unlock()
	delete(Rooms.list, key)
}

//获取牌桌接口
func Get(key string) inter.IDesk {
	Rooms.RLock()
	defer Rooms.RUnlock()
	if room, ok := Rooms.list[key]; ok {
		return room
	}
	return nil
}

//是否存在
func Exist(key string) bool {
	Rooms.RLock()
	defer Rooms.RUnlock()
	_, ok := Rooms.list[key]
	return ok
}

//关闭列表
func Close() {
	Rooms.Lock()
	defer Rooms.Unlock()
	close(Rooms.closeCh) //关闭
}

//计时器
func rooms_ticker() {
	tick := time.Tick(10 * time.Minute) //每局限制了最高10分钟
	glog.Infof("rooms ticker started -> %d", 1)
	for {
		select {
		case <-tick:
			//逻辑处理
			rooms_expire(false) //过期清理
		case <-Rooms.closeCh:
			glog.Infof("rooms close -> %d", len(Rooms.list))
			//TODO:rooms_expire(true) //强制关闭
			return
		}
	}
}

//定期清理or关闭清除
func rooms_expire(ok bool) {
	Rooms.Lock()
	defer Rooms.Unlock()
	//TODO   下面开了goroutine去删除房间，上面的房间列表锁就不能锁住房间列表了
	for _, r := range Rooms.list {
		go func(room inter.IDesk) {
			room.Closed(ok) //goroutine处理,避免删除时死锁
		}(r)
	}
}
