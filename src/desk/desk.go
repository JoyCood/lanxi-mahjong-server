/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2016-12-25 10:00
 * Filename      : desk.go
 * Description   : 玩牌逻辑
 * *******************************************************/
package desk

import (
	"algo"
	"data"
	"inter"
	"sync"
	"time"
	"fmt"
	"basic/utils"
	"math/rand"
	"resource"

	"github.com/golang/glog"
)

//房间牌桌数据结构
type Desk struct {
	id      uint32                   //房间id
	data    *DeskData                //房间类型基础数据
	dealer  uint32                   //庄家的座位
	dice    uint32                   //骰子
	cards   []byte                   //没摸起的海底牌
	players map[uint32]inter.IPlayer //房间玩家
	//-------------
	lian    uint32            //私人房间连庄数
	round   uint32            //私人房间打牌局数
	//-------------
	vote    uint32            //投票发起者座位号
	votes   map[uint32]uint32 //投票同意解散de玩家座位号列表
	voteT   *time.Timer //投票定时器
	record  utils.Array //打牌记录
	//-------------
	//操作提示优先级 : (胡)-(碰,杠)-(吃)
	//如果是同一玩家 : (胡,碰,杠), (胡,碰,杠,吃)
	//如果多个玩家胡 : (胡)-(碰,杠)-(吃)
	//map[位置]操作值

	// TODO 下面的数据类型可也自定义需要的类型，这样也方便程序自动初始化
	hu       map[uint32]uint32
	pongkong map[uint32]uint32
	chow     map[uint32]uint32
	skip     map[uint32]bool //过圈
	//-------------
	discard  byte   //出牌
	draw     byte   //模牌
	seat     uint32 //当前模牌|出牌位置
	timer    int    //计时
	operate  int    //操作状态
	huing    int    //一炮多响胡牌
	//-------------
	sync.Mutex  //房间锁
	state  bool //房间状态
	kong   bool //是否杠牌出牌
	//-------------
	closeCh chan bool //关闭通道
	//-------------
	trusteeship map[uint32]bool //托管
	ready       map[uint32]bool //是否准备
	//-------------

	// TODO 下面的数据类型可也自定义需要的类型，这样也方便程序自动初始化
	outCards  map[uint32][]byte   //海底牌
	pongCards map[uint32][]uint32 //碰牌
	kongCards map[uint32][]uint32 //杠牌
	chowCards map[uint32][]uint32 //吃牌(8bit-8-8)
	handCards map[uint32][]byte   //手牌
}

//// external function

//新建一张牌桌
func NewDesk(data *DeskData) inter.IDesk {

	// TODO 复杂类型自定义，就不用手动初始化
	desk := &Desk{
		id:      data.Rid,
		data:    data,
		players: make(map[uint32]inter.IPlayer),
		votes:   make(map[uint32]uint32),
		record:  utils.NewArray(true),
		//------
		trusteeship: make(map[uint32]bool),
		ready:       make(map[uint32]bool),
	}
	return desk
}

//投票解散,ok=true发起,=false投票
func (t *Desk) Vote(ok bool, seat, vote uint32) int {
	t.Lock() //房间加锁
	defer t.Unlock()
	if t.state {
		return 1
	}
	if t.vote != 0 && ok { //发起
		return 2
	}
	if t.vote == 0 && !ok { //投票
		return 2
	}
	t.votes[seat] = vote //投票
	if ok { //发起投票
		t.vote = seat //发起投票者
		t.voteT = time.AfterFunc(2 * time.Minute,
		func() { t.dismiss(true) }) //超时设置
		msg := res_voteStart(seat)
		t.broadcast(msg)
		//return 0
	}
	msg := res_vote(seat, vote)
	t.broadcast(msg)
	t.dismiss(false)
	return 0
}

//投票解散,agree > unagree
func (t *Desk) dismiss(ok bool) {
	var agree int = 0
	var unagree int = 0
	var i uint32
	for i = 1; i <= 4; i++ {
		if _, ok2 := t.votes[i]; !ok2 && ok {
			t.votes[i] = 0 //超时不投票默认算同意
		}
		if _, ok3 := t.players[i]; !ok3 && !ok {
			t.votes[i] = 0 //空位置提前默认同意
		}
	}
	for i = 1; i <= 4; i++ {
		if v, ok4 := t.votes[i]; ok4 && v == 0 {
			agree++
		} else {
			unagree++
		}
	}
	if agree > unagree { //一半以上通过即可
		msg := res_voteResult(0) //0解散,1不解散
		t.broadcast(msg)
		if t.voteT != nil {
			t.voteT.Stop()
		}
		round, expire := t.getRound()
		t.printf("dismiss") //test
		t.close(round, expire, true) //解散
	} else if ok || len(t.votes) == 4 { //结束投票
		msg := res_voteResult(1)
		t.broadcast(msg)
		t.vote = 0
		t.votes = make(map[uint32]uint32)
		if t.voteT != nil {
			t.voteT.Stop()
		}
	}
}

//托管处理,Kind=1:托管;Kind=0:取消托管
func (t *Desk) Trust(seat, kind uint32) {
	t.Lock() //房间加锁
	defer t.Unlock()
	t.trust_(seat, kind)
}

//托管处理,Kind=1:托管;Kind=0:取消托管
func (t *Desk) trust_(seat, kind uint32) {
	var tru bool = false
	if kind == 1 { //托管
		tru = true
	}
	if t.getTrust(seat) == tru { //相同设置不处理
		return
	}
	t.trusteeship[seat] = tru //设置状态
	msg := res_trust(seat, kind)
	t.broadcast(msg) //广播消息
}

//获取牌桌数据
func (t *Desk) GetData() interface{} {
	return t.data
}

//房间消息广播,聊天
func (t *Desk) Broadcasts(msg inter.IProto) {
	t.broadcast(msg)
}

//关闭房间,停服or过期清除等等.TODO:玩牌中是否关闭?
func (t *Desk) Closed(ok bool) {
	round, expire := t.getRound()
	t.printf("Closed") //test
	t.close(round, expire, ok) //ok=true强制解散,=false清理
}

//玩家准备
func (t *Desk) Readying(seat uint32, ready bool) int {
	t.Lock() //房间加锁
	defer t.Unlock()
	if t.vote != 0 { //投票中不能准备
		return 1
	}
	t.ready[seat] = ready //设置状态
	msg := res_ready(seat, ready)
	t.broadcast(msg) //广播消息
	t.autoDiceing() //自动打骰
	return 0
}

//打印牌局状态信息,test
func (t *Desk) Print() {
	t.Lock() //房间加锁
	defer t.Unlock()
	t.print()
}

//打印牌局状态信息,test
func (t *Desk) print() {
	glog.Infof("t.id -> %d, t.he -> %+x", t.id, t.hu)
	glog.Infof("t.seat -> %d, t.dealer -> %d", t.seat, t.dealer)
	glog.Infof("t.discard -> %x, t.draw -> %x", t.discard, t.draw)
	glog.Infof("handCards -> %+x", t.handCards)
	glog.Infof("pongCards -> %+x", t.pongCards)
	glog.Infof("kongCards -> %+x", t.kongCards)
	glog.Infof("chowCards -> %+x", t.chowCards)
	glog.Infof("outCards -> %+x", t.outCards)
	glog.Infof("ready -> %+v", t.ready)
	glog.Infof("len(payers) -> %d", len(t.players))
	glog.Infof("score -> %+v", t.data.Score)
	glog.Infof("t.state -> %+v", t.state)
}

//玩家打骰子切牌,发牌
func (t *Desk) Diceing() bool {
	t.Lock() //房间加锁
	defer t.Unlock()
	if t.isDiceing() {//是否打骰
		t.dealer_()   //打庄
		t.gameStart() //开始牌局
		return true
	} else {
		return false
	}
}

//进入
func (t *Desk) Enter(p inter.IPlayer) int {
	t.Lock() //房间加锁
	defer t.Unlock()
	round, expire := t.getRound()
	if t.reEnter(round, expire, p) { //检测重复进入
		return 0
	}
	var seat uint32 = t.add(p)
	if seat == 0 {
		return 1
	}
	msg1 := res_enter(t.id, seat, round, expire,
	t.dealer, t.data, t.players, t.ready)
	p.Send(msg1)
	uid := p.GetUserid()
	score := t.data.Score[uid]
	msg2 := res_othercomein(p, score)
	t.broadcast_(seat, msg2)
	t.printf("enter") //test
	return 0
}

func (t *Desk) printf(act string) {
	glog.Infof("act -> %s, t.id -> %d, t.data -> %+v", act, t.id, t.data)
	for k, v := range t.players {
		glog.Infof("seat %d -> uid %s", k, v.GetUserid())
	}
}

//玩家离开
func (t *Desk) Leave(seat uint32) bool {
	t.Lock() //房间加锁
	defer t.Unlock()
	if t.state { //游戏中不能离开
		return false
	}
	//广播消息
	msg := res_leave(seat)
	t.broadcast(msg)
	//清除数据
	delete(t.ready, seat)
	delete(t.players, seat)
	return true
}

//踢除玩家
func (t *Desk) Kick(cid string, seat uint32) bool {
	t.Lock() //房间加锁
	defer t.Unlock()
	if t.state { //游戏中不能离开
		return false
	}
	if cid != t.data.Cid {
		return false
	}
	if p, ok := t.players[seat]; ok {
		p.ClearRoom() //清除玩家房间数据
	}
	//广播消息
	msg := res_leave(seat)
	t.broadcast(msg)
	//清除数据
	delete(t.ready, seat)
	delete(t.players, seat)
	return true
}

//// internal function

//计时器
func (t *Desk) ticker() {
	tick := time.Tick(time.Second)
	glog.Infof("ticker -> %d", t.id)
	for {
		select {
		case <-tick:
			//超时判断
			if t.state {
				if t.timer == OT || t.timer == DT {
					t.timerL() //逻辑处理
				} else {
					t.timer++
				}
			}
		case <-t.closeCh:
			glog.Infof("close desk -> %d", t.id)
			return
		}
	}
}

//超时处理
func (t *Desk) timerL() {
	//操作(胡,碰杠,吃)超时处理
	if t.timer == OT && t.discard != 0 {
		t.timer = 0
		t.TurnL() //操作(自动胡牌,操作超时...)
	} else if t.timer == DT && t.draw != 0 {
		//出牌超时处理
		t.timer = 0
		t.DiscardL(t.seat, t.draw, true) //超时打出摸到的牌
	} else {
		t.timer++
	}
}

//加入牌局
func (t *Desk) add(p inter.IPlayer) uint32 {
	var i uint32
	for i = 1; i <= 4; i++ {
		if _, ok := t.players[i]; !ok {
			t.players[i] = p
			p.SetRoom(t.data.Rtype, t.id, i, t.data.Code)
			return i
		}
	}
	return 0
}

//检测重复进入
func (t *Desk) reEnter(round, expire uint32, p inter.IPlayer) bool {
	var seat uint32 = p.GetPosition()
	if _, ok := t.players[seat]; !ok {
		return false
	}
	msg1 := t.res_reEnter(t.id, seat, round, expire,
	t.dealer, t.data, t.players)
	p.Send(msg1)
	return true
}

//是否全部准备状态
func (t *Desk) allReady() bool {
	if len(t.ready) != 4 {
		return false
	}
	for _, ok := range t.ready {
		if !ok {
			return false
		}
	}
	return true
}

//是否可以打骰
func (t *Desk) isDiceing() bool {
	if t.state  { //已经开始
		return false
	}
	if t.dealer != 0 { //已经打庄
		return false
	}
	if len(t.players) != 4 { //人数不够
		return false
	}
	if !t.allReady() { //没有准备
		return false
	}
	return true
}

//自动打骰(打骰请求超时或断开情况)
func (t *Desk) autoDiceing() {
	go func() {
		utils.Sleep(2) //延迟2秒
		t.Diceing()    //主动打骰
	}()
}

//打庄处理
func (t *Desk) dealer_() {
	//选择庄家
	if t.lian == 0 {
		t.dealer = uint32(utils.RandInt32N(4) + 1)
	} else {
		t.dealer = t.lian
	}
	msg := res_dealer(t.dealer)
	t.broadcast(msg) //打庄消息通知
}

//开始游戏
func (t *Desk) gameStart() {
	t.gameStartInit() //初始化
	glog.Infof("gameStart -> %d, seat -> %d", t.id, t.seat)
	//打庄消息通知+抽水
	//var coin int32 = -100 //TODO:抽水
	var cost int32 = -1 * int32(t.data.Cost)
	var payment uint32 = t.data.Payment
	for k, p := range t.players {
		p.SetReady(false) //设置人物游戏状态
		//resource.ChangeRes(p, resource.COIN, coin, 1) //抽水
		//开始游戏扣除创建房间钻石
		if t.round != 0 {
			continue
		}
		if payment != 1 && k != t.dealer {
			continue
		}
		resource.ChangeRes(p, resource.DIAMOND, cost, data.RESTYPE4)
	}
	//打骰(两个骰子)
	dice1 := uint32(utils.RandInt32N(5) + 1)
	dice2 := uint32(utils.RandInt32N(5) + 1)
	t.dice = (dice1 << 16) + dice2 //TODO:优化
	t.shuffle() //洗牌
	t.deal() //发牌
	//等待玩家操作
}

//初始化
func (t *Desk) gameStartInit() {
	t.state = true //设置房间状态
	t.timer = -6 //重置计时器,6秒给前端展示动画
	if t.closeCh == nil {
		t.closeCh = make(chan bool, 1)
		//go t.ticker() //计时器goroutine
	}
	//------
	t.operateInit()
	t.skip =       make(map[uint32]bool)
	//------
	t.outCards =  make(map[uint32][]byte)   //海底牌
	t.pongCards = make(map[uint32][]uint32) //碰牌
	t.kongCards = make(map[uint32][]uint32) //杠牌
	t.chowCards = make(map[uint32][]uint32) //吃牌(8bit-8-8)
	t.handCards = make(map[uint32][]byte)   //手牌
}

//初始化操作值
func (t *Desk) operateInit() {
	t.hu =         make(map[uint32]uint32)
	t.pongkong =   make(map[uint32]uint32)
	t.chow =       make(map[uint32]uint32)
}

//发牌
func (t *Desk) deal() {
	for s, _ := range t.players {
		var hand int = int(algo.HAND)
		if s == t.dealer { //判断庄家发14张牌
			hand += 1
		}
		cards := make([]byte, hand, hand)
		tmp := t.cards[:hand]
		copy(cards, tmp)
		t.handCards[s] = cards
		t.cards = t.cards[hand:]
	}
	//第一个操作为庄家
	t.seat = t.dealer
	//发牌协议消息
	for s, p := range t.players {
		var cards []byte = t.getHandCards(s)
		if t.dealer == s {
			//庄家提示处理
			var v uint32 = algo.DrawDetect(byte(0), cards, []uint32{}, []uint32{}, []uint32{})
			v |= t.heType(v, t.seat, 0, cards) //大七对,清一色,字一色
			if v > 0 {
				t.hu[s] = v //设置操作状态值
			}
			t.draw = cards[len(cards)-1] //庄家最后一张默认为摸牌
			//庄家消息
			msg := res_zhuangDeal(v, t.dice, cards)
			p.Send(msg)
		} else {
			//闲家消息
			msg := res_deal(0, t.dice, cards)
			p.Send(msg)
		}
	}
}

//洗牌
func (t *Desk) shuffle() {
	rand.Seed(time.Now().UnixNano())
	d := make([]byte, algo.TOTAL, algo.TOTAL)
	copy(d, algo.CARDS)
	//测试暂时去掉洗牌
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	t.cards = d
}

//模牌,kong==false普通摸牌,kong==true扛后摸牌
func (t *Desk) drawcard() {
	if len(t.cards) == 0 {
		t.he(0) //结束牌局
		return
	}
	var value uint32 = 0
	//var seat uint32 = t.seat //过圈
	if t.kong { //杠后摸牌
		value = 1
	} else { //普通摸牌, t.discrad != 0
		var outs []byte = t.getOutCards(t.seat)
		outs = append(outs, t.discard)
		t.outCards[t.seat] = outs
		//位置切换
		t.seat = algo.NextSeat(t.seat)
	}
	//t.skip_(seat, t.seat) //过圈
	t.unskipL(t.seat) //摸牌一定过圈
	var card byte = t.cards[0]
	t.draw = card //设置摸牌状态
	t.discard = 0 //清除打牌状态
	t.operate = 0 //清除操作状态
	t.huing = 0   //清除胡牌状态
	t.cards = t.cards[1:]
	var cards []byte = t.in(t.seat, card)
	var ps []uint32 = t.getPongCards(t.seat)
	var ks []uint32 = t.getKongCards(t.seat)
	var cs []uint32 = t.getChowCards(t.seat)
	var v uint32 = algo.DrawDetect(card, cards, cs, ps, ks)
	v |= t.heType(v, t.seat, 0, cards) //大七对,清一色,字一色
	//杠上开花
	if t.kong && v&algo.HU > 0 {
		v = v | algo.HU_KONG_FLOWER
	}
	if v > 0 { //摸牌全部记录为胡(只自己操作)
		t.hu[t.seat] = v
	}
	//记录 TODO:暂时没添加
	//t.Record(algo.DRAW, card, t.seat)
	//其他玩家消息
	msg1 := res_otherDraw(t.seat, value)
	//摸牌协议消息
	msg2 := res_draw(value, v, card, cards)
	//摸牌协议消息通知
	for s, o := range t.players {
		if s == t.seat {
			//摸牌玩家消息
			o.Send(msg2)
		} else {
			//其他玩家消息
			o.Send(msg1)
		}
	}
	//托管处理
	if t.getTrust(t.seat) {
		t.timer = DT - 2 //出牌时间
	} else {
		t.timer = 0 //重置计时
	}
}

//出牌,加锁,托管自摸自动胡牌
func (t *Desk) DiscardL(seat uint32, card byte, ok bool) {
	t.Lock()
	defer t.Unlock()
	//glog.Infof("DiscardL -> %d, seat -> %d, card -> %x", t.id, t.seat, card)
	if !t.state {
		glog.Infof("DiscardL err -> %d", seat)
		return
	}
	if seat < 1 || seat > 4 {
		glog.Infof("Discard seat err -> %d", seat)
		return
	}
	if seat != t.seat || t.draw == 0 {
		p := t.getPlayer(seat)
		msg := res_discard2()
		p.Send(msg)
		glog.Infof("Discard seat err -> %d", seat)
		return
	}
	if ok { //超时操作
		t.trust_(seat, 1) //设置玩家超时托管
		if _, ok := t.hu[seat]; ok {
			t.he(seat) //托管自摸自动胡牌
			return
		}
	} else {
		t.trust_(seat, 0) //设置玩家超时托管
	}
	t.discard_(card)
}

//出牌,没加锁
func (t *Desk) discard_(card byte) {
	//记录 TODO:暂时没添加
	//t.Record(algo.DISCARD, card, t.seat)
	t.operateInit()  //清除操作记录
	t.discard = card //设置打牌状态
	t.draw = 0       //清除摸牌状态
	t.operate = 0    //清除操作状态
	t.huing = 0      //清除胡牌状态
	//检测(胡,碰杠,吃)
	for s, _ := range t.players {
		if s == t.seat { //出牌人跳过
			t.out(t.seat, card)  //移除牌
			continue
		}
		var cards []byte = t.getHandCards(s)
		var ok bool = t.getSkip(s) //是否过圈
		if !ok {
			//胡,杠碰,吃检测
			v_h := algo.DiscardHu(card, cards) //胡
			v_h |= t.heType(v_h, s, card, cards) //大七对,清一色,字一色
			if v_h > 0 {
				t.hu[s] = v_h
			}
		}
		v_p := algo.DiscardPong(card, cards) //碰杠
		if v_p > 0 {
			t.pongkong[s] = v_p
		}
		v_c := algo.DiscardChow(t.seat, s, card, cards) //吃
		if v_c > 0 {
			t.chow[s] = v_c
		}
	}
	if t.kong { //杠操作出牌标识
		t.kong = false //杠后出牌清除
	}
	t.order() //按规则优先级设置
}

//按规则优先级设置
func (t *Desk) order() {
	//TODO:优化order + turn
	var l_h, l_p, l_c int = t.getLength()
	if l_h == 0 && l_p == 0 && l_c == 0 { //无操作
		//出牌协议消息通知
		msg := res_discard(t.seat, t.discard)
		t.broadcast(msg) //消息广播
		t.drawcard() //摸牌
		return
	}
	//多人胡,碰杠一人,吃一人
	if l_p == 1 {
		for k, v := range t.pongkong {
			if vc, ok := t.chow[k]; ok { //碰杠吃为同一人
				t.pongkong[k] = v | vc
				t.chow = make(map[uint32]uint32) //清除
				break
			}
		}
	} else if l_p == 0 && l_c != 0 { //吃,胡一起提示,TODO:优化
		t.pongkong = t.chow
		t.chow = make(map[uint32]uint32) //清除
	}
	if l_h == 1 { //一人胡
		for k, v := range t.hu {
			if vc, ok := t.pongkong[k]; ok { //胡碰杠为同一人
				t.hu[k] = v | vc
				t.pongkong = make(map[uint32]uint32) //清除
				break
			}
		}
	}
	//多人胡按顺序操作
	t.turn(false)
}

//超时操作,加锁
func (t *Desk) TurnL() {
	t.Lock()
	defer t.Unlock()
	t.turn(true)
}

//超时操作,没加锁 ok表示是否超时处理
func (t *Desk) turn(ok bool) {
	var l_h, l_p, l_c int = t.getLength()
	if l_h == 0 && l_p == 0 && l_c == 0 { //无操作
		t.drawcard() //摸牌
		return
	}
	if ok { //超时
		if l_h != 0 {
			for k, _ := range t.hu {
				t.he(k) //超时自动胡牌
				return
			}
			//t.skipL(0) //跳过胡牌,不自动胡
			//t.hu = make(map[uint32]uint32)
		} else if l_p != 0 {
			t.pongkong = make(map[uint32]uint32)
		} else if l_c != 0 {
			t.chow = make(map[uint32]uint32)
		}
		//操作超时不托管,只有打牌超时才托管
		//消除提示操作消息,下个操作消息同时进行
		t.turn(false) //发送下一个操作消息
		return
	}
	//操作按胡-碰杠-吃顺序,相同操作(胡),TODO:优化
	//t.hu -> t.pongkong -> t.chow
	var m map[uint32]uint32 = make(map[uint32]uint32)
	if l_h != 0 {
		m = t.hu
	} else if l_p != 0 {
		m = t.pongkong
	} else if l_c != 0 {
		m = t.chow
	} else {
		panic(fmt.Sprintf("turn error:%d", t.seat))
	}
	var card byte = t.operateCard() //所操作的牌
	//广播操作消息
	for k, v := range m {
		p := t.getPlayer(k)
		if t.operate == 0 { //第一次提示操作
			t.operate += 1 //操作状态变化
			msg1 := res_discard3(t.seat, v, card)
			p.Send(msg1)
			msg2 := res_discard(t.seat, card)
			t.broadcast_(k, msg2) //消息广播
		} else { //二次提示操作
			t.operate += 1 //操作状态变化
			msg3 := res_pengkong(t.seat, v, card)
			p.Send(msg3)
		}
		//托管处理
		if t.getTrust(k) {
			t.timer = OT - 2 //出牌时间
		} else {
			t.timer = 0 //重置计时
		}
	}
}

//胡碰杠吃操作 TODO:优化
func (t *Desk) OperateL(card, value, seat uint32) {
	t.Lock()
	defer t.Unlock()
	if !t.state {
		glog.Infof("operate err -> %d", seat)
		return
	}
	if seat < 1 || seat > 4 {
		glog.Infof("operate seat err -> %d", seat)
		return
	}
	var l_h, l_p, l_c int = t.getLength()
	if l_h == 0 && l_p == 0 && l_c == 0 { //无操作
		glog.Infof("operate err -> %d", seat)
		return
	}
	t.trust_(seat, 0) //设置玩家超时托管
	if l_h > 1 { //多人胡,有胡
		t.operateH(card, value, seat) //TODO:区别处理??
	} else if l_h == 1 { //一人胡,有胡碰杠吃
		t.operateH(card, value, seat)
	} else if l_p == 1 { //没有胡,有碰杠吃 
		t.operateP(card, value, seat)
	} else if l_c == 1 { //没有碰杠,有吃
		t.operateC(card, value, seat)
	} else { //错误
		glog.Infof("unexpected operate err -> %d", seat)
	}
}

//胡碰杠吃操作,TODO:优化
func (t *Desk) operateH(card, value, seat uint32) {
	if v, ok := t.hu[seat]; ok {
		if value == 0 {
			t.skipL(seat) //跳过胡牌
			delete(t.hu, seat)
			t.cancelOperate(seat, t.seat, value, card)
			if len(t.hu) == t.huing && t.huing != 0 {
				for k, _ := range t.hu {
					t.he(k) //选择位置,TODO:超时时huing处理
					return
				}
			} else if seat != t.seat { //如果暗杠时取消不应该摸牌,应该打牌
				t.turn(false) //取消时进入下一个操作
			}
		} else if value&algo.HU > 0 {
			t.huing++ //选择胡牌人数
			if len(t.hu) == t.huing { //多人胡
				t.he(seat)
			}
		} else {
			t.operateA(card, value, seat, v)
		}
	} else {
		//不存在操作
		glog.Errorf("not exist operate err -> %d", seat)
	}
}

//胡碰杠吃操作
func (t *Desk) operateP(card, value, seat uint32) {
	if v, ok := t.pongkong[seat]; ok {
		if value == 0 {
			delete(t.pongkong, seat)
			t.cancelOperate(seat, t.seat, value, card)
			t.turn(false) //取消时进入下一个操作
		} else {
			t.operateA(card, value, seat, v)
		}
	} else {
		//不存在操作
		glog.Errorf("not exist operate err -> %d", seat)
	}
}

//胡碰杠吃操作
func (t *Desk) operateC(card, value, seat uint32) {
	if v, ok := t.chow[seat]; ok {
		if value == 0 {
			delete(t.chow, seat) //先清除
			t.cancelOperate(seat, t.seat, value, card)
			t.turn(false) //取消时进入下一个操作
		} else {
			t.operateA(card, value, seat, v)
		}
	} else {
		//不存在操作
		glog.Errorf("not exist operate err -> %d", seat)
	}
}

//胡碰杠吃操作,TODO:碰杠验证处理
func (t *Desk) operateA(card, value, seat, v uint32) {
	if value&algo.CHOW > 0 && v&algo.CHOW > 0 {
		t.chow_(card, value, seat)
	} else if value&algo.PENG > 0 && v&algo.PENG > 0 {
		t.pong_(card, value, seat)
	} else if value&algo.MING_KONG > 0 && v&algo.MING_KONG > 0 {
		t.kong_(card, value, seat, v, algo.MING_KONG)
	} else if value&algo.BU_KONG > 0 && v&algo.BU_KONG > 0 {
		t.kong_(card, value, seat, v, algo.BU_KONG)
	} else if value&algo.AN_KONG > 0 && v&algo.AN_KONG > 0 {
		t.kong_(card, value, seat, v, algo.AN_KONG)
	} else {
		//不存在操作
		glog.Errorf("not exist operate err -> %d", seat)
	}
}

//胡牌,(多人胡牌时同时胡),t.seat=出牌(放冲)位置
func (t *Desk) he(seat uint32) {
	for k, v := range t.players {
		glog.Infof("he seat %d -> uid %s", k, v.GetUserid())
	}
	glog.Infof("id -> %d, he -> %+x", t.id, t.hu)
	glog.Infof("seat -> %d, t.seat -> %d", seat, t.seat)
	glog.Infof("discard -> %x, draw -> %x", t.discard, t.draw)
	glog.Infof("handCards -> %+x", t.handCards)
	glog.Infof("pongCards -> %+x", t.pongCards)
	glog.Infof("kongCards -> %+x", t.kongCards)
	glog.Infof("chowCards -> %+x", t.chowCards)
	glog.Infof("outCards -> %+x", t.outCards)
	var card byte = t.operateCard() //所胡的牌
	t.tianHe() //天地胡
	//去掉平胡
	for k, v := range t.hu {
		v = algo.CancelHuPing(v)
		t.hu[k] = algo.MenQingFan2(v)
	}
	glog.Infof("id -> %d, he -> %+x", t.id, t.hu)
	glog.Infof("seat -> %d, t.seat -> %d", seat, t.seat)
	glog.Infof("seat -> %d, t.dealer -> %d", seat, t.dealer)
	t.qiangKongHe(seat, card) //抢杠胡处理
	var l_h bool = len(t.hu) > 0 //是否胡牌
	//算番
	huFan, mingKong, beMingKong, anKong, buKong, total := t.gameOver(l_h)
	coin := make(map[uint32]int32)
	//TODO:金币,积分结算
	for i, v := range total {
		p := t.getPlayer(i)
		uid := p.GetUserid()
		v *= int32(t.data.Ante) //番*底分
		t.data.Score[uid] += v  //总分
		coin[i] = v //当局分
	}
	glog.Infof("score -> %+v", t.data.Score)
	glog.Infof("huFan:%+v, mingKong:%+v, beMingKong:%+v, anKong:%+v, buKong:%+v, total:%+v",
	huFan, mingKong, beMingKong, anKong, buKong, total)
	if l_h { //胡牌
		//胡牌消息
		msg1 := res_he(seat, card)
		t.broadcast(msg1)
	}
	//结算消息广播
	msg2 := res_over(t.seat, t.handCards, t.hu, huFan, mingKong, beMingKong, anKong, buKong, total, coin)
	t.broadcast(msg2)
	t.round++ //局数
	round, expire := t.getRound()
	msg := res_privateOver(t.id, round, expire, t.players, t.data.Score)
	t.broadcast(msg)  //私人局结束消息广播
	t.privaterecord(coin) //日志记录 TODO:goroutine
	t.lianDealer(l_h, seat) //连庄
	t.overSet()       //重置状态
	t.close(round, expire, false) //结束牌局
}

//1.胡的牌(自摸,放炮),2.抢杠胡时操作的牌(补杠人的摸牌),TODO:优化
func (t *Desk) operateCard() byte {
	if t.discard != 0 {
		return t.discard
	} else {
		return t.draw
	}
}

//抢杠胡处理,TODO:多人抢杠胡
func (t *Desk) qiangKongHe(seat uint32, card byte) {
	if v, ok := t.hu[seat]; ok {
		if v&algo.QIANG_GANG > 0 {
			msg := res_operate2(seat, t.seat, algo.QIANG_GANG,
			algo.QIANG_GANG, card)
			t.broadcast(msg)
			//去掉被抢杠玩家(t.seat)的补杠,TODO:要不要还原之前的碰?
			kongs := t.getKongCards(t.seat)
			var cs []uint32
			for i, v2 := range kongs { //杠
				_, c, mask := algo.DecodeKong(v2) //解码
				if c == card && mask == algo.BU_KONG {
					cs = append(kongs[:i], kongs[i+1:]...)
					t.kongCards[t.seat] = cs
					break
				}
			}
		}
	}
}

//天地胡
func (t *Desk) tianHe() {
	var l_s uint32 = uint32(len(t.cards))
	var l_h uint32 = algo.TOTAL - (algo.HAND * 4 + 1)
	if l_s != l_h {
		return
	}
	//TIAN_HU,DI_HU
	var l_o int = len(t.outCards)
	var l_p int = len(t.pongCards)
	var l_k int = len(t.kongCards)
	var l_c int = len(t.chowCards)
	if l_o == 0 && l_p == 0 &&
	l_k == 0 && l_c == 0 {
		if t.seat == t.dealer && t.discard == 0 {
			for k, v := range t.hu {
				t.hu[k] = v|algo.TIAN_HU
			}
		} else {
			for k, v := range t.hu {
				t.hu[k] = v|algo.DI_HU
			}
		}
	}
}

//大七对,清一色,字一色
func (t *Desk) heType(val, seat uint32, card byte, hands []byte) uint32 {
	var cs []byte = []byte{}
	if card != 0 {
		cs = append(cs, card)
	}
	cs = append(cs, hands...)
	chows := t.getChowCards(seat)
	for _, v1 := range chows { //吃
		c1, c2, c3 := algo.DecodeChow(v1) //解码
		cs = append(cs, c1, c2, c3)
	}
	kongs := t.getKongCards(seat)
	for _, v2 := range kongs { //杠
		_, c, _ := algo.DecodeKong(v2) //解码
		cs = append(cs, c, c, c, c)
	}
	pongs := t.getPongCards(seat)
	for _, v3 := range pongs { //碰
		_, c := algo.DecodePeng(v3) //解码
		cs = append(cs, c, c, c)
	}
	var hu bool = false //是否有胡牌牌型
	if val&algo.HU > 0 {
		hu = true //有牌型
	}
	var chow bool = len(chows) != 0 //有吃
	var kong bool = algo.KongDetect(hands) //手牌中是否有杠
	var v uint32 = algo.HuTypeDetect(hu, chow, kong, cs)
	if v > 0 {
		if card == 0 { //自摸
			return v|algo.ZIMO
		} else { //放炮
			return v|algo.PAOHU
		}
	}
	return v
}

//连庄
func (t *Desk) lianDealer(l_h bool, seat uint32) {
	if l_h { //胡牌
		if len(t.hu) > 1 {
			t.lian = t.seat //一炮多响,放炮者当庄
		} else {
			t.lian = seat //胡牌玩家连庄
		}
	} else { //黄庄
		t.lian = t.seat //最后摸牌玩家
	}
}

//结束牌局重置状态数据
func (t *Desk) overSet() {
	t.state = false //牌局状态
	t.dealer = 0    //庄家重置
	t.timer = 0     //重置计时
	t.discard = 0   //重置打牌
	t.draw = 0      //重置摸牌
	t.dice = 0      //重置骰子
	t.seat = 0      //清除位置
	t.huing = 0     //清除胡牌
	t.kong = false  //清除杠牌
	t.trusteeship = make(map[uint32]bool)
	t.ready =       make(map[uint32]bool)
}

//结束牌局,ok=true投票解散
func (t *Desk) close(round, expire uint32, ok bool) {
	var n uint32 = uint32(utils.Timestamp())
	if (round > 0 && expire > n) && !ok {
		return
	}
	if t.closeCh != nil {
		close(t.closeCh) //关闭计时器
		t.closeCh = nil  //消除计时器
	}
	for k, p := range t.players {
		msg := res_leave(k)
		t.broadcast(msg)
		p.ClearRoom() //清除玩家房间数据
	}
	Del(t.data.Code) //从房间列表中清除
}

//胡牌过圈
func (t *Desk) skipL(seat uint32) {
	//多人超时时处理
	//var m map[uint32]uint32 = make(map[uint32]uint32)
	//if v, ok := t.hu[seat]; ok {
	//	m[seat] = v
	//} else { //seat=0表示全部跳过
	//	m = t.hu
	//}
	//for k, v := range m {
	//	if v&algo.ZIMO > 0 || k == t.seat { //自摸跳过
	//		continue
	//	}
	//	t.skip[k] = true //跳过胡牌过圈设置
	//}
	//没有超时时处理
	t.skip[seat] = true //跳过胡牌过圈设置
}

//清除过圈(摸牌,抢杠)
func (t *Desk) unskipL(seat uint32) {
	if _, ok := t.skip[seat]; ok {
		delete(t.skip, seat)
	}
}

//胡牌过圈,1.杠后出牌,2.庄家出牌,3.正常过圈
func (t *Desk) skip_(s1, s2 uint32) {
	if len(t.skip) == 0 { //empty
		return
	}
	if s2 == t.dealer || t.kong { //1,2
		t.skip = make(map[uint32]bool)
		return
	}
	var s []uint32 = make([]uint32, 0)
	for k, _ := range t.skip { //3
		if skiped(s1, s2, k) {
			s = append(s, k)
		}
	}
	for _, v := range s { //过圈位置
		delete(t.skip, v)
	}
}

//s1上一个位置,s2当前位置,s3胡牌位置
func skiped(s1, s2, s3 uint32) bool {
	if s1 == s2 { //s3不在s1,s2之间
		return false
	}
	var s4 = algo.NextSeat(s1)
	if s4 == s3 { //s3在s1,s2之间
		return true
	}
	return skiped(s4, s2, s3)
}

//碰操作,已经验证通过
func (t *Desk) pong_(card, value, seat uint32) {
	var cards []byte = t.ponging(seat, byte(card))
	//碰操作协议消息通知
	msg := res_operate(seat, t.seat, value, card)
	t.broadcast(msg)
	t.skip_(t.seat, seat) //过圈
	//状态设置
	t.seat = seat  //位置切换
	t.timer = 0    //重置计时
	t.draw = cards[len(cards)-1] //设置摸牌,超时出牌时打出
	t.discard = 0 //重置出牌
	t.operateInit() //清除操作记录,操作成功后消除,防止重复提示
	//等待出牌
}

//杠操作,已经验证通过
func (t *Desk) kong_(card1, value, seat, v, mask uint32) {
	var card byte = byte(card1)
	switch mask {
	case algo.BU_KONG:
		t.buKong(seat, card)
	case algo.MING_KONG:
		t.mingKong(seat, card)
	case algo.AN_KONG:
		t.anKong(seat, card)
	}
	//杠操作协议消息通知
	msg := res_operate(seat, t.seat, value, card1)
	t.broadcast(msg)
	t.skip_(t.seat, seat) //过圈
	//状态设置
	t.kong = true //杠操作出牌标识
	t.seat = seat //位置切换
	t.timer = 0   //重置计时
	//抢杠处理
	var ok bool = t.qiangKong(card, mask) //检测是否抢杠
	if ok { //TODO:优化
		//t.operate += 1 //操作状态变化
		t.turn(false) //抢杠操作
	} else {
		t.drawcard() //摸牌
	}
}

//吃操作,已经验证通过
func (t *Desk) chow_(card, value, seat uint32) {
	c1, c2 := algo.DecodeChow2(card)
	var ok bool = t.chowing(seat, c1, c2)
	if !ok {
		glog.Errorf("chow card error -> %d", card)
	}
	//吃操作协议消息通知
	var card2 uint32 = algo.EncodeChow(c1,c2,t.discard)
	msg := res_operate(seat, t.seat, value, card2)
	t.broadcast(msg)
	t.skip_(t.seat, seat) //过圈
	//状态设置
	t.seat = seat  //位置切换
	t.timer = 0    //重置计时
	var cards []byte = t.handCards[seat]
	t.draw = cards[len(cards)-1] //设置摸牌,超时出牌时打出
	t.discard = 0 //重置出牌
	t.operateInit() //清除操作记录
	//等待出牌
}

//补杠被抢杠,抢杠处理,抢杠胡牌
func (t *Desk) qiangKong(card byte, mask uint32) bool {
	t.operateInit() //清除操作记录
	if mask != algo.BU_KONG {
		return false
	}
	//检测(抢杠胡)
	for s, _ := range t.players {
		if s == t.seat { //出牌人跳过
			continue
		}
		//抢杠不用过圈
		//var ok bool = t.getSkip(s) //是否过圈
		//if !ok {
		var cards []byte = t.getHandCards(s)
		//胡,杠碰,吃检测
		v_h := algo.DiscardHu(card, cards) //胡
		v_h |= t.heType(v_h, s, card, cards) //大七对,清一色,字一色
		if v_h > 0 {
			t.unskipL(s) //抢杠玩家一定过圈
			t.hu[s] = v_h | algo.QIANG_GANG
		}
	}
	if len(t.hu) > 0 {
		return true
	}
	return false
}

//取消操作时消息通知
func (t *Desk) cancelOperate(seat, beseat, value, card uint32) {
	msg := res_operate(seat, beseat, value, card)
	p := t.getPlayer(seat)
	p.Send(msg)
}

//获取玩家
func (t *Desk) getPlayer(seat uint32) inter.IPlayer {
	if v, ok := t.players[seat]; ok && v != nil {
		return v
	}
	panic(fmt.Sprintf("getPlayer error:%d", seat))
}

//获取手牌
func (t *Desk) getHandCards(seat uint32) []byte {
	if v, ok := t.handCards[seat]; ok && v != nil {
		return v
	}
	panic(fmt.Sprintf("getHandCards error:%d", seat))
}

//获取海底牌
func (t *Desk) getOutCards(seat uint32) []byte {
	if v, ok := t.outCards[seat]; ok && v != nil {
		return v
	}
	return []byte{}
}

//获取碰牌
func (t *Desk) getPongCards(seat uint32) []uint32 {
	if v, ok := t.pongCards[seat]; ok && v != nil {
		return v
	}
	return []uint32{}
}

//获取杠牌
func (t *Desk) getKongCards(seat uint32) []uint32 {
	if v, ok := t.kongCards[seat]; ok && v != nil {
		return v
	}
	return []uint32{}
}

//获取吃牌
func (t *Desk) getChowCards(seat uint32) []uint32 {
	if v, ok := t.chowCards[seat]; ok && v != nil {
		return v
	}
	return []uint32{}
}

//获取玩家托管状态
func (t *Desk) getTrust(seat uint32) bool {
	if v, ok := t.trusteeship[seat]; ok {
		return v
	}
	return false
}

//获取玩家准备状态
func (t *Desk) getReady(seat uint32) bool {
	if v, ok := t.ready[seat]; ok {
		return v
	}
	return false
}

//获取玩家过圈状态
func (t *Desk) getSkip(seat uint32) bool {
	if v, ok := t.skip[seat]; ok {
		return v
	}
	return false
}

//获取剩余局数,结束时间
func (t *Desk) getRound() (uint32, uint32) {
	var expire uint32 = 0
	var now int64 = utils.Timestamp()
	if int64(t.data.Expire) > now {
		expire = t.data.Expire
	}
	var round uint32 = t.data.Round - t.round
	if round < 0 {
		round = 0
	}
	return round, expire
}

//获取操作值长度
func (t *Desk) getLength() (int, int, int) {
	return len(t.hu), len(t.pongkong), len(t.chow)
}

//房间消息广播
func (t *Desk) broadcast(msg inter.IProto) {
	for _, p := range t.players {
		p.Send(msg)
	}
}

//房间消息广播(除seat外)
func (t *Desk) broadcast_(seat uint32, msg inter.IProto) {
	for i, p := range t.players {
		if i != seat {
			p.Send(msg)
		}
	}
}

//--------操作
//摸牌
func (t *Desk) in(seat uint32, card byte) []byte {
	var cards []byte = t.getHandCards(seat)
	cards = append(cards, card)
	t.handCards[seat] = cards
	return cards
}

//玩家出牌
func (t *Desk) out(seat uint32, card byte) {
	var cards []byte = t.getHandCards(seat)
	cards = algo.Remove(card, cards)
	t.handCards[seat] = cards
}

//吃牌操作
func (t *Desk) chowing(seat uint32, c1, c2 byte) bool {
	var c []byte = []byte{t.discard, c1, c2}
	algo.Sort(c, 0, 2)
	//验证吃, c1,c2,c3有序
	if !algo.VerifyChow(c[0], c[1], c[2]) {
		return false
	}
	var cards []byte = t.getHandCards(seat)
	var isExist1 bool = algo.Exist(c1, cards, 1)
	if !isExist1 {
		glog.Errorf("chowing card error -> %d", c1)
		return false
	}
	var isExist2 bool = algo.Exist(c2, cards, 1)
	if !isExist2 {
		glog.Errorf("chowing card error -> %d", c2)
		return false
	}
	cards = algo.Remove(c1, cards)
	cards = algo.Remove(c2, cards)
	var cs []uint32 = t.getChowCards(seat)
	cs = append(cs, algo.EncodeChow(c1, c2, t.discard))
	t.handCards[seat] = cards
	t.chowCards[seat] = cs
	return true
}

//碰牌操作
func (t *Desk) ponging(seat uint32, card byte) []byte {
	var cards []byte = t.getHandCards(seat)
	var isExist bool = algo.Exist(card, cards, 2)
	if !isExist {
		glog.Errorf("ponging card error -> %d", card)
		return cards
	}
	cards = algo.RemoveN(card, cards, 2)
	var cs []uint32 = t.getPongCards(seat)
	cs = append(cs, algo.EncodePeng(seat, card))
	t.handCards[seat] = cards
	t.pongCards[seat] = cs
	return cards
}

//暗扛操作
func (t *Desk) anKong(seat uint32, card byte) {
	var cards []byte = t.getHandCards(seat)
	var isExist bool = algo.Exist(card, cards, 4)
	if !isExist {
		glog.Errorf("anKong card error -> %d", card)
		return
	}
	cards = algo.RemoveN(card, cards, 4)
	var cs []uint32 = t.getKongCards(seat)
	cs = append(cs, algo.EncodeKong(0, card, algo.AN_KONG))
	t.handCards[seat] = cards
	t.kongCards[seat] = cs
}

//明杠操作
func (t *Desk) mingKong(seat uint32, card byte) {
	var cards []byte = t.getHandCards(seat)
	var isExist bool = algo.Exist(card, cards, 3)
	if !isExist {
		glog.Errorf("mingKong card error -> %d", card)
		return
	}
	cards = algo.RemoveN(card, cards, 3)
	var cs []uint32 = t.getKongCards(seat)
	cs = append(cs, algo.EncodeKong(t.seat, card, algo.MING_KONG))
	t.handCards[seat] = cards
	t.kongCards[seat] = cs
}

//补杠操作
func (t *Desk) buKong(seat uint32, card byte) {
	var cards []byte = t.getHandCards(seat)
	var isExist bool = algo.Exist(card, cards, 1)
	if !isExist {
		glog.Errorf("buKong card error -> %d", card)
		return
	}
	var pongs []uint32 = t.getPongCards(seat)
	for i, v := range pongs {
		_, c := algo.DecodePeng(v)
		if c == card {
			pongs = append(pongs[:i], pongs[i+1:]...)
			break
		}
	}
	cards = algo.Remove(card, cards)
	var cs []uint32 = t.getKongCards(seat)
	cs = append(cs, algo.EncodeKong(0, card, algo.BU_KONG))
	t.handCards[seat] = cards
	t.kongCards[seat] = cs
	t.pongCards[seat] = pongs
}
