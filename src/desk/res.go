/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2016-12-27 17:00
 * Filename      : res.go
 * Description   : 玩牌协议消息响应
 * *******************************************************/
package desk

import (
	"algo"
	"inter"
	"protocol"
	"errorcode"

	"github.com/golang/protobuf/proto"
)

//操作提示响应消息
func res_operate(seat, beseat, value, card uint32) inter.IProto {
	stoc := &protocol.SOperate{}
	stoc.Seat = proto.Uint32(seat)
	stoc.Card = proto.Uint32(card)
	stoc.Value = proto.Uint32(value)
	stoc.Beseat = proto.Uint32(beseat)
	return stoc
}

//操作提示响应消息
func res_operate2(seat, beseat, value, qiang uint32, card byte) inter.IProto {
	stoc := &protocol.SOperate{}
	stoc.Seat = proto.Uint32(seat)
	stoc.Card = proto.Uint32(uint32(card))
	stoc.Value = proto.Uint32(value)
	stoc.Beseat = proto.Uint32(beseat)
	stoc.Discontinue = proto.Uint32(qiang)
	return stoc
}

//打牌响应消息
func res_discard(seat uint32, card byte) inter.IProto {
	stoc := &protocol.SDiscard{}
	stoc.Card = proto.Uint32(uint32(card))
	stoc.Seat = proto.Uint32(seat)
	return stoc
}

func res_discard2() inter.IProto {
	stoc := &protocol.SDiscard{}
	stoc.Error = proto.Uint32(errorcode.NotYourTurn)
	return stoc
}

func res_discard3(seat, v uint32, card byte) inter.IProto {
	stoc := &protocol.SDiscard{}
	stoc.Card = proto.Uint32(uint32(card))
	stoc.Seat = proto.Uint32(seat)
	stoc.Value = proto.Uint32(v)
	return stoc
}

//处理前面有玩家胡牌优先操作,如果该玩家跳过胡牌,此协议向有碰和明杠的玩家主动发送
func res_pengkong(seat, v uint32, card byte) inter.IProto {
	stoc := &protocol.SPengKong{}
	stoc.Card = proto.Uint32(uint32(card))
	stoc.Seat = proto.Uint32(seat)
	stoc.Value = proto.Uint32(v)
	return stoc
}

//进入房间响应消息
func res_othercomein(p inter.IPlayer, score int32) inter.IProto {
	userinfo := p.ConverProtoUser()
	userinfo.Score = proto.Int32(score)
	stoc := &protocol.SOtherComein{Userinfo: userinfo}
	return stoc
}

//进入房间响应消息
func res_enter(id, seat, round, expire, dealer uint32,
data *DeskData, m map[uint32]inter.IPlayer,
ready map[uint32]bool) inter.IProto {
	var len uint32 = uint32(len(m))
	stoc := &protocol.SEnterSocialRoom{}
	roomdata := &protocol.RoomData{
		Roomid:     proto.Uint32(id),
		Rtype:      proto.Uint32(data.Rtype),
		Rname:      proto.String(data.Rname),
		Expire:     proto.Uint32(uint32(expire)),
		Round:      proto.Uint32(round),
		Count:      proto.Uint32(len),
		Invitecode: proto.String(data.Code),
		Zhuang:     proto.Uint32(dealer),
		Userid:     proto.String(data.Cid),
	}
	stoc.Room = roomdata
	var score int32
	for k, p := range m {
		uid := p.GetUserid()
		score1 := data.Score[uid]
		if k != seat {
			var r bool = false
			if r_v, ok := ready[k]; ok {
				r = r_v
			}
			userinfo := p.ConverProtoUser()
			userinfo.Score = proto.Int32(score1)
			userinfo.Ready = proto.Bool(r)
			stoc.Userinfo = append(stoc.Userinfo, userinfo)
		} else {
			score = score1
		}
	}
	stoc.Position = proto.Uint32(seat)
	stoc.Score = proto.Int32(score)
	stoc.Ready = proto.Bool(false)
	stoc.Beginning = proto.Bool(false)
	stoc.Ma = proto.Uint32(0)
	stoc.LaunchSeat = proto.Uint32(0)
	return stoc
}

//重新进入房间响应消息
func (t *Desk)  res_reEnter(id, seat, round, expire, dealer uint32, 
data *DeskData, m map[uint32]inter.IPlayer) inter.IProto {
	var lenght uint32 = uint32(len(m))
	stoc := &protocol.SEnterSocialRoom{}
	roomdata := &protocol.RoomData{
		Roomid:     proto.Uint32(id),
		Rtype:      proto.Uint32(data.Rtype),
		Rname:      proto.String(data.Rname),
		Expire:     proto.Uint32(uint32(expire)),
		Round:      proto.Uint32(round),
		Count:      proto.Uint32(lenght),
		Invitecode: proto.String(data.Code),
		Zhuang:     proto.Uint32(dealer),
		Userid:     proto.String(data.Cid),
	}
	stoc.Room = roomdata
	var score int32
	for k, p := range m {
		uid := p.GetUserid()
		score1 := data.Score[uid]
		if k != seat {
			var r bool = t.getReady(k)
			userinfo := p.ConverProtoUser()
			userinfo.Ready = proto.Bool(r)
			userinfo.Score = proto.Int32(score1)
			stoc.Userinfo = append(stoc.Userinfo, userinfo)
		} else {
			score = score1
		}
	}
	stoc.Position = proto.Uint32(seat)
	var r bool = t.getReady(seat)
	stoc.Ready = proto.Bool(r)
	stoc.Beginning = proto.Bool(t.state)
	stoc.Score = proto.Int32(score)
	stoc.Ma = proto.Uint32(0)
	stoc.LaunchSeat = proto.Uint32(t.vote)
	if t.vote > 0 { //投票中
		for k, v := range t.votes {
			if v == 0 {
				stoc.VoteAgree = append(stoc.VoteAgree, k)
			} else {
				stoc.VoteDisagree = append(stoc.VoteDisagree, k)
			}
		}
	}
	//重连数据 TODO:优化
	if !t.state {
		stoc.Beginning = proto.Bool(false)
		return stoc
	}
	stoc.Beginning = proto.Bool(true)
	stoc.Turn = proto.Uint32(t.seat)
	stoc.Dice = proto.Uint32(t.dice)
	stoc.CardsCount = proto.Uint32(uint32(len(t.cards)))
	stoc.Handcards = t.getHandCards(seat)
	var value uint32
	if v_h, ok := t.hu[seat]; ok {
		value = v_h
	} else if v_p, ok := t.pongkong[seat]; ok {
		value = v_p
	} else if v_c, ok := t.chow[seat]; ok {
		value = v_c
	}
	stoc.Value = proto.Uint32(value) //操作值
	var kongnum uint32
	for i, _ := range t.players {
		pcard := &protocol.ProtoCard{}
		pcard.Seat = proto.Uint32(i)
		pongs := t.getPongCards(i) //碰牌数据
		for _, v := range pongs {
			j, card := algo.DecodePeng(v)
			var peng uint32 = j << 24
			peng |= (uint32(card) << 16)
			pcard.Peng = append(pcard.Peng, peng)
		}
		kongs := t.getKongCards(i) //杠牌数据
		for _, v := range kongs {
			j, card, classify := algo.DecodeKong(v)
			var kong uint32 = j << 24
			kong |= (uint32(card) << 16)
			kong |= (classify << 8)
			pcard.Kong = append(pcard.Kong, kong)
			kongnum = kongnum + 1
		}
		chows := t.getChowCards(i) //吃牌数据
		pcard.Chow = chows
		pcard.Outcards = t.getOutCards(i)
		//加上进行中的牌
		if t.seat == i && t.discard != 0 {
			pcard.Outcards = append(pcard.Outcards, t.discard)
		}
		stoc.Cards = append(stoc.Cards, pcard)
	}
	// 桌上有几个杠(未摸起的牌，牌尾被摸走几张牌)
	stoc.KongCount = proto.Uint32(kongnum)
	return stoc
}

//打庄响应消息
func res_dealer(dealer uint32) inter.IProto {
	stoc := &protocol.SZhuang{}
	stoc.Zhuang = proto.Uint32(dealer)
	//stoc.Lian = proto.Uint32(t.lian)
	return stoc
}

//庄家响应消息
func res_zhuangDeal(v, dice uint32, cards []byte) inter.IProto {
	stoc := &protocol.SZhuangDeal{}
	stoc.Value = proto.Uint32(v)
	stoc.Dice = proto.Uint32(dice)
	stoc.Cards = cards
	return stoc
}

//闲家响应消息
func res_deal(v, dice uint32, cards []byte) inter.IProto {
	stoc := &protocol.SDeal{}
	stoc.Value = proto.Uint32(v)
	stoc.Dice = proto.Uint32(dice)
	stoc.Cards = cards
	return stoc
}

//闲家响应消息
func res_otherDraw(seat, value uint32) inter.IProto {
	stoc := &protocol.SOtherDraw{}
	stoc.Seat = proto.Uint32(seat)
	stoc.Kong = proto.Uint32(value)
	return stoc
}

//摸牌协议消息响应消息
func res_draw(kong, v uint32, card byte, cards []byte) inter.IProto {
	stoc := &protocol.SDraw{}
	stoc.Card = proto.Uint32(uint32(card))
	stoc.Cards = cards
	stoc.Kong = proto.Uint32(kong)
	stoc.Value = proto.Uint32(v)
	return stoc
}

//玩家准备消息响应消息
func res_ready(seat uint32, ready bool) inter.IProto {
	stoc := &protocol.SReady{}
	stoc.Ready = proto.Bool(ready)
	stoc.Seat = proto.Uint32(seat)
	return stoc
}

//结束牌局响应消息
func res_he(seat uint32, card byte) inter.IProto {
	stoc := &protocol.SHu{}
	stoc.Seat = proto.Uint32(seat)
	stoc.Card = proto.Uint32(uint32(card))
	return stoc
}

//结束牌局响应消息,huType:0:黄庄，1:自摸，2:炮胡
func res_over(seat uint32,
handCards map[uint32][]byte, hu map[uint32]uint32,
huFan, mingKong, beMingKong, anKong,
buKong, total, coin map[uint32]int32) inter.IProto {
	stoc := &protocol.SGameover{
		Data: make([]*protocol.ProtoCount, 4),
	}
	var huType uint32 = 0 //胡牌类型
	var paoSeat uint32 = 0 //放冲玩家
	var maCards []byte = []byte{}
	var i uint32
	for i = 1; i <= 4; i++ {
		var val uint32 = 0 //胡牌掩码
		if v, ok := hu[i]; ok {
			if v&algo.PAOHU > 0 { //放冲
				huType = 2
				paoSeat = seat
			} else if v&algo.ZIMO > 0 { //自摸
				huType = 1
			}
			//val = v //胡牌掩码
			val = res_over_hu(v) //显示处理
		}
		protoCount := &protocol.ProtoCount{}
		protoCount.Seat = proto.Uint32(i)
		protoCount.Ting = proto.Uint32(0)
		protoCount.Hu = proto.Uint32(val)
		// 结算时要显示该玩家的手牌
		protoCount.Cards = handCards[i]
		protoCount.Total = proto.Int32(total[i])
		protoCount.Coin = proto.Int32(coin[i])
		stoc.Data[i-1] = protoCount
		//算番
		//protoCount.HuTypeFan = proto.Int32(huTypeFan[i])
		protoCount.HuTypeFan = proto.Int32(huFan[i]) //前端只显示牌型分,所以给的一样
		protoCount.HuFan = proto.Int32(huFan[i])
		protoCount.MingKong = proto.Int32(mingKong[i])
		protoCount.BeMingKong = proto.Int32(beMingKong[i])
		protoCount.AnKong = proto.Int32(anKong[i])
		protoCount.BuKong = proto.Int32(buKong[i])
		protoCount.Lian = proto.Int32(0)
		protoCount.Ma = proto.Int32(0)
		protoCount.MaCards = maCards
	}
	stoc.AllMaCards = maCards
	stoc.PaoSeat = proto.Uint32(paoSeat)
	stoc.HuType = proto.Uint32(huType)
	return stoc
}

//自摸平胡(只显示自摸),点炮(放冲)平胡(只显示点炮)
func res_over_hu(val uint32) uint32 {
	if val&algo.PAOHU > 0 && val&algo.HU_PING > 0 {
		return val^algo.HU_PING
	}
	if val&algo.ZIMO > 0 && val&algo.HU_PING > 0 {
		return val^algo.HU_PING
	}
	return val
}

//离开房间响应消息
func res_leave(seat uint32) inter.IProto {
	stoc := &protocol.SPrivateLeave{}
	stoc.Seat = proto.Uint32(seat)
	return stoc
}

//私人局结束响应消息
func res_privateOver(id, round, expire uint32,
m map[uint32]inter.IPlayer, n map[string]int32) inter.IProto {
	stoc := &protocol.SPrivateOver{}
	// 如果是私人房，且房间过期 ，踢掉房间玩家
	stoc.Cid = proto.Uint32(0)
	stoc.Roomid = proto.Uint32(id)
	stoc.Round = proto.Uint32(round)
	stoc.Expire = proto.Uint32(expire)
	for _, p := range m {
		userid := p.GetUserid()
		s := &protocol.PrivateScore{
			Userid: proto.String(userid),
			Score:  proto.Int32(n[userid]),
		}
		stoc.List = append(stoc.List, s)
	}
	return stoc
}

//发起投票申请解散房间
func res_voteStart(seat uint32) inter.IProto {
	stoc := &protocol.SLaunchVote{Seat: proto.Uint32(seat)}
	return stoc
}

//投票解散房间事件结果
func res_voteResult(vote uint32) inter.IProto {
	stoc := &protocol.SVoteResult{Vote: proto.Uint32(vote)}
	return stoc
}

//投票
func res_vote(seat, vote uint32) inter.IProto {
	stoc := &protocol.SVote{
		Vote: proto.Uint32(vote),
		Seat: proto.Uint32(seat),
	}
	return stoc
}

//托管消息响应
func res_trust(seat, kind uint32) inter.IProto {
	stoc := &protocol.STrusteeship{
		Seat: proto.Uint32(seat),
		Kind: proto.Uint32(kind),
	}
	return stoc
}

/*
庄自摸	    庄炮胡	        闲自摸	            闲炮胡(闲放)	闲炮胡(庄放)
牌型*庄*3	牌型*放炮*庄	牌型+牌型+牌型*庄	牌型*放炮	    牌型*庄*放炮
6*牌型	    4*牌型	        4*牌型	            2*牌型	        4*牌型
庄家自摸时,其他玩家需每人多支付2倍,如:庄家平胡自摸所得番数为:1*2*3
庄家炮胡时,放炮者需支付给庄家2倍,其他玩家不需要给,如:庄家平胡炮胡所得番数为:1*2*2
其他玩家自摸时,庄家需要额外支付2倍,如:闲家平胡自摸所得番数为:1+1+1*2
其他玩家炮胡时,如果是闲家放炮,则庄家不需要给番数,所得番数为:1*2
其他玩家炮胡时,如果是庄家放炮,则所得番数为:1*2*2
*/
//结算,(明杠,放冲,庄家 - 收一家)
func (t *Desk) gameOver(l_h bool) (huFan, mingKong, beMingKong, anKong, buKong, total map[uint32]int32) {
	//huTypeFan  = make(map[uint32]int32)// 胡牌方式番数
	huFan      = make(map[uint32]int32)// 胡牌牌型番数
	mingKong   = make(map[uint32]int32)// 闷豆的番数
	beMingKong = make(map[uint32]int32)// 被点豆的负番数
	anKong     = make(map[uint32]int32)// 明豆的番数
	buKong     = make(map[uint32]int32)// 拐弯豆的番数
	total      = make(map[uint32]int32)// 总番数
	if !l_h { //黄庄
		return huFan, mingKong, beMingKong, anKong, buKong, total
	}
	var k uint32
	for k = 1; k <= 4; k++ {
		//杠牌分
		kongs := t.getKongCards(k) //杠牌数据
		//牌型分
		if v, ok := t.hu[k]; ok {
			v = algo.DanDiao(v, kongs) //有暗杠时不算点炮(放冲,单钓)
			t.hu[k] = v //
			handCards := t.getHandCards(k)   //手牌
			f_t := algo.HuType(v, handCards) //胡牌牌型,多个牌型时相乘
			f_w := algo.HuWay(v)             //胡牌方式,多个方式时相乘
			f_tw := f_t * f_w                //牌型分
			//牌型分,t.seat=出牌(放冲)位置
			huFan = fanType(t.dealer, t.seat, k, f_tw, huFan)
		}
		//杠牌分
		//kongs := t.getKongCards(k) //杠牌数据
		for _, v := range kongs {  //杠牌分
			_, _, cy := algo.DecodeKong(v) //解码杠值
			f_k := algo.HuKong(cy) //杠
			if cy == algo.MING_KONG {
				//mingKong[k] += f_k       //收一家
				//beMingKong[i] += 0 - f_k //被收一家
				mingKong = over3(mingKong, k, f_k) //收三家
			} else if cy == algo.BU_KONG {
				buKong = over3(buKong, k, f_k) //收三家
			} else if cy == algo.AN_KONG {
				anKong = over3(anKong, k, f_k) //收三家
			}
		}
	}
	//总番数
	for k = 1; k <= 4; k++ {
		total[k] += huFan[k] + mingKong[k] + beMingKong[k] + anKong[k] + buKong[k]
	}
	return huFan, mingKong, beMingKong, anKong, buKong, total
}

//倍数(放冲,庄家 - 双倍) TODO:优化
//dealer=庄家位置,paoseat=放炮位置,huseat=胡牌位置
func fanType(dealer, paoseat, huseat uint32, f_tw int32,
hf map[uint32]int32) map[uint32]int32 {
	if paoseat == huseat { //自摸,收三家
		if huseat == dealer { //庄家自摸
			//6 //其它3家*1 (庄家*1倍)
			hf = over3(hf, dealer, f_tw * 1) //收三家
		} else { //闲家自摸
			//4 //庄家*1 + 其它2家*1
			var i uint32
			for i = 1; i <= 4; i++ {
				if i == huseat {
					continue
				}
				if i == dealer {
					hf = over1(hf, huseat, i, f_tw * 1) //收一家
				} else {
					hf = over1(hf, huseat, i, f_tw * 1) //收一家
				}
			}
		}
	} else { //炮胡,收一家
		if huseat == dealer { //庄家胡(肯定闲家放炮)
			//4//收放炮的*1 (放1倍*庄1倍)
			hf = over1(hf, huseat, paoseat, f_tw * 1) //收一家
		} else { //闲家胡
			if paoseat == dealer { //庄家放炮
				//4//收庄家的*1 (放炮1倍*庄家1倍)
				hf = over1(hf, huseat, dealer, f_tw * 1) //收一家
			} else { //闲家放炮
				//2//收闲家的*1 (放炮1倍)
				hf = over1(hf, huseat, paoseat, f_tw * 1) //收一家
			}
		}
	}
	return hf
}

//收三家,seat=收的位置,val=收的番数
func over3(m map[uint32]int32, seat uint32, val int32) map[uint32]int32 {
	var i uint32
	for i = 1; i <= 4; i++ {
		var value int32
		if i != seat {
			value = 0 - val //为负数
		} else {
			value = 3 * val //收三家
		}
		m[i] += value
	}
	return m
}

//收一家,s1=收的位置,s2=出的位置,val=收的番数
func over1(m map[uint32]int32, s1, s2 uint32, val int32) map[uint32]int32 {
	m[s1] += val
	m[s2] -= val
	return m
}
