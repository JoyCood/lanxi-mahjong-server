// Code generated by protoc-gen-go.
// source: bankrupt.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

约定：首字母为C是客户端请求协议，首字母为S是服务器返回协议，其他为嵌套结构体

It is generated from these files:
	bankrupt.proto
	circle.proto
	coin.proto
	conf.proto
	login.proto
	match.proto
	postbox.proto
	ranking.proto
	room.proto
	signin.proto
	socialroom.proto
	sys.proto
	task.proto
	user.proto
	vo.proto

It has these top-level messages:
	CBankrupt
	SBankrupt
	CCreateRoom
	SCreateRoom
	SOtherCreateRoom
	CUpdateRoomPerson
	SUpdateRoomPerson
	CPrivateCircleRecord
	SPrivateCircleRecord
	CPrivateRoomRecord
	SPrivateRoomRecord
	SResource
	ResVo
	CGetCurrency
	SGetCurrency
	CConfig
	SConfig
	SysConfig
	CLogin
	SLogin
	CMiguLogin
	SMiguLogin
	ReLogin
	CRegist
	SRegist
	CSetPasswd
	SSetPasswd
	Sdisconnection
	CWechatLogin
	SWechatLogin
	CJoinMatchGame
	SJoinMatchGame
	SMatchAllOver
	CUpdatePersonNum
	SUpdatePersonNum
	CMatchExit
	SMatchExit
	CUpdateRank
	SUpdateRank
	MatchRank
	SMatchComein
	CPost
	SPost
	CDelPost
	SDelPost
	CDelReadPost
	SDelReadPost
	CDelAllPost
	SDelAllPost
	COpenAppendix
	SOpenAppendix
	CReadPost
	SReadPost
	CRankCoin
	SRankCoin
	CRankGainCoin
	SRankGainCoin
	CRankWin
	SRankWin
	CRankExp
	SRankExp
	CRankRewards
	SRankRewards
	CComein
	SComein
	SOtherComein
	SZhuang
	CComeinRoomid
	SWaitBroken
	CBroken
	SBroken
	SZhuangDeal
	SDeal
	SDraw
	CDiscard
	SOtherDraw
	SDiscard
	COperate
	SOperate
	SPengKong
	SHu
	CHu
	CLeave
	SLeave
	CReComein
	SReComein
	SGameover
	SJI
	CBroadcastChatText
	CBroadcastChat
	SBroadcastChatText
	SBroadcastChat
	CTing
	STing
	CQiangKong
	SQiangKong
	CGetCount
	SGetCount
	CTrusteeship
	STrusteeship
	CSign
	SSign
	CSignDay
	SSignDay
	CReSign
	SReSign
	CEnterSocialRoom
	SEnterSocialRoom
	CPrivateLeave
	SPrivateLeave
	SReady
	CReady
	SPrivateOver
	PrivateScore
	CCreatePrivateRoom
	SCreatePrivateRoom
	SStart
	CStartGame
	SStartGame
	CKick
	SKick
	CLaunchVote
	SLaunchVote
	CVote
	SVote
	SVoteResult
	CPrivateRecord
	SPrivateRecord
	PrivateRecord
	PrivateRecordDetails
	PrivateRecords
	PrivateDetails
	CPRecordByRid
	SPRecordByRid
	CPing
	SPing
	CFeedback
	SFeedback
	CNotice
	SNotice
	Notice
	CBuy
	SBuy
	CBuild
	SBuild
	CActivity
	ProtoActivity
	SActivity
	CGetActivityRewards
	SGetActivityRewards
	SUpdateActivity
	CWechatShare
	CTradeList
	STradeList
	ProtoTrade
	CTradeRecord
	STradeRecord
	ProtoTradeRecard
	CTradeUserInfo
	STradeUserInfo
	CTrade
	STrade
	CIapppayOrder
	SIapppayOrder
	CApplePay
	SApplePay
	CWxpayOrder
	SWxpayOrder
	CWxpayQuery
	SWxpayQuery
	CTask
	ProtoTask
	STask
	CGetTaskRewards
	SGetTaskRewards
	SUpdateTask
	CChangeNickname
	SChangeNickname
	CChangeSex
	SChangeSex
	CArchieve
	SArchieve
	CUserData
	SUserData
	SChenmi
	UserData
	ProtoUser
	CircleAttrube
	MatchAttrube
	RoomAttrube
	RoomData
	ProtoRoom
	CircleData
	RankCoin
	RankExp
	RankWin
	PrivateRecordForCircle
	PrivateRecordForRoom
	ProtoCard
	ProtoCount
	PostBoxData
	PostAppendixData
	WidgetData
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 玩家破产领取救济金
type CBankrupt struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=7801" json:"code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CBankrupt) Reset()                    { *m = CBankrupt{} }
func (m *CBankrupt) String() string            { return proto.CompactTextString(m) }
func (*CBankrupt) ProtoMessage()               {}
func (*CBankrupt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_CBankrupt_Code uint32 = 7801

func (m *CBankrupt) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CBankrupt_Code
}

// 如果今天的所有救济金已经领取完或者没达到领取条件,error字段会返回错误号
type SBankrupt struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=7801" json:"code,omitempty"`
	Count            *uint32 `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	Coin             *uint32 `protobuf:"varint,3,req,name=coin" json:"coin,omitempty"`
	Error            *uint32 `protobuf:"varint,4,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SBankrupt) Reset()                    { *m = SBankrupt{} }
func (m *SBankrupt) String() string            { return proto.CompactTextString(m) }
func (*SBankrupt) ProtoMessage()               {}
func (*SBankrupt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_SBankrupt_Code uint32 = 7801
const Default_SBankrupt_Error uint32 = 0

func (m *SBankrupt) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SBankrupt_Code
}

func (m *SBankrupt) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *SBankrupt) GetCoin() uint32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *SBankrupt) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SBankrupt_Error
}

func init() {
	proto.RegisterType((*CBankrupt)(nil), "protocol.CBankrupt")
	proto.RegisterType((*SBankrupt)(nil), "protocol.SBankrupt")
}

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4a, 0xcc, 0xcb,
	0x2e, 0x2a, 0x2d, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9,
	0x39, 0x4a, 0xf2, 0x5c, 0x9c, 0xce, 0x4e, 0x50, 0x49, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x5e, 0x2b, 0x16, 0x73, 0x0b, 0x03, 0x43, 0x25, 0x3f, 0x2e,
	0xce, 0x60, 0x7c, 0x0a, 0x84, 0x78, 0xb9, 0x58, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0x14,
	0x98, 0x34, 0x78, 0x85, 0x78, 0x40, 0x4a, 0x32, 0xf3, 0x24, 0x98, 0xc1, 0x3c, 0x01, 0x2e, 0xd6,
	0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0x09, 0x16, 0xb0, 0x0e, 0x46, 0x03, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x16, 0x69, 0x65, 0xce, 0x8b, 0x00, 0x00, 0x00,
}
