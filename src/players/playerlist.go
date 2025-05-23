package players

import (
	"inter"
	"sync"
	"time"

	"github.com/golang/glog"
)

var list *playerList

type playerList struct {
	sync.RWMutex
	list map[string]inter.IPlayer
	closeCh chan bool           //关闭通道
}

func init() {
	list = &playerList{
		list: make(map[string]inter.IPlayer),
		closeCh: make(chan bool, 1),
	}
	go players_ticker() //goroutine,定时清理
}

//没有被调用会造成内存泄漏
func Del(id string) {
	list.Lock()
	defer list.Unlock()
	if p, ok := list.list[id]; ok {
		p.UserSave() //
	}
	delete(list.list, id)
}

//TODO:存在数据替换可能
func Set(id string, player inter.IPlayer) {
	list.Lock()
	defer list.Unlock()
	list.list[id] = player
}

func Get(id string) inter.IPlayer {
	list.RLock()
	defer list.RUnlock()
	player, ok := list.list[id]
	if ok {
		return player
	}
	return nil
}

//关闭列表
func Close() {
	list.Lock()
	defer list.Unlock()
	for _, p := range list.list {
		p.UserSave() //
	}
	close(list.closeCh) //关闭
}

//计时器
func players_ticker() {
	tick := time.Tick(5 * time.Minute) //每5分钟
	glog.Infof("players ticker started -> %d", 1)
	for {
		select {
		case <-tick:
			players_remove() //清理
		case <-list.closeCh:
			glog.Infof("players close -> %d", len(list.list))
			return
		}
	}
}

//清理
func players_remove() {
	list.Lock()
	defer list.Unlock()
	for i, p := range list.list {
		go func(id string, player inter.IPlayer) {
			var off bool = true
			conn := player.GetConn()
			if conn != nil && conn.GetConnected() {
				off = false
			}
			if off && player.GetRoomType() == 0 {
				Del(id)
			}
		}(i, p)
	}
}
