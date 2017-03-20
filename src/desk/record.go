/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-02-26 20:31
 * Filename      : gamerecord.go
 * Description   : 牌局记录
 * *******************************************************/
package desk

import (
	"algo"
	"data"
)

// private record
func (t *Desk) privaterecord(coins map[uint32]int32) {
	var huType uint32 = 3 //胡牌类型
	for _, v := range t.hu {
		if v&algo.PAOHU > 0 { //放冲
			huType = 2
		} else if v&algo.ZIMO > 0 { //自摸
			huType = 1
		}
	}
	//---
	rtype  := t.data.Rtype
	rname  := t.data.Rname
	ante   := t.data.Ante
	code   := t.data.Code
	rounds := t.data.Round
	zhuang := t.dealer
	round  := t.round
	id     := t.id
	expire := t.data.Expire
	payment:= t.data.Payment
	cid    := t.data.Cid
	ctime  := t.data.CTime
	//---
	roundRecord := &data.GameOverRoundRecord{
		Zhuang: zhuang,
		Hutype: huType,
		Round:  round,
	}
	//---
	var userids []string
	for k, v := range t.players {
		//var userid string = v.GetUserid()
		var coin int32 = 0
		//if n, ok := t.data.Score[userid]; ok {
		if n, ok := coins[k]; ok {
			coin = n //只记录当前局输赢
		}
		var huValue uint32
		var hucard uint32
		details := &data.GameOverUserRecord{
			Userid:  v.GetUserid(),
			Seat:    k,
			Coin:    coin, //只记录当前局输赢
			Huvalue: huValue,
			HuCard:  hucard,
		}
		roundRecord.Users = append(roundRecord.Users, details)
		userids = append(userids, v.GetUserid())
	}
	if round == 1 {
		record := &data.GameOverRecord{
			RoomId:     id,
			TotalRound: rounds,
			Invitecode: code,
			Rtype:      rtype,
			Rname:      rname,
			Ante:       ante,
			Userids:    userids,
			Cid:        cid,
			Payment:    payment,
			Expire:     expire,
			Ctime:      ctime,
		}
		record.Rounds = append(record.Rounds, roundRecord)
		record.Add()
	} else {
		roundRecord.Push(id)
	}
}
