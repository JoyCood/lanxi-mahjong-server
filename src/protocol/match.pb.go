// Code generated by protoc-gen-go.
// source: match.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 报名加入比赛
type CJoinMatchGame struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8001" json:"code,omitempty"`
	Kind             *uint32 `protobuf:"varint,2,req,name=kind" json:"kind,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CJoinMatchGame) Reset()                    { *m = CJoinMatchGame{} }
func (m *CJoinMatchGame) String() string            { return proto.CompactTextString(m) }
func (*CJoinMatchGame) ProtoMessage()               {}
func (*CJoinMatchGame) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

const Default_CJoinMatchGame_Code uint32 = 8001

func (m *CJoinMatchGame) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CJoinMatchGame_Code
}

func (m *CJoinMatchGame) GetKind() uint32 {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return 0
}

type SJoinMatchGame struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8001" json:"code,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Kind             *uint32 `protobuf:"varint,3,req,name=kind" json:"kind,omitempty"`
	Playernum        *uint32 `protobuf:"varint,4,req,name=playernum" json:"playernum,omitempty"`
	Error            *uint32 `protobuf:"varint,5,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SJoinMatchGame) Reset()                    { *m = SJoinMatchGame{} }
func (m *SJoinMatchGame) String() string            { return proto.CompactTextString(m) }
func (*SJoinMatchGame) ProtoMessage()               {}
func (*SJoinMatchGame) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

const Default_SJoinMatchGame_Code uint32 = 8001
const Default_SJoinMatchGame_Error uint32 = 0

func (m *SJoinMatchGame) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SJoinMatchGame_Code
}

func (m *SJoinMatchGame) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SJoinMatchGame) GetKind() uint32 {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return 0
}

func (m *SJoinMatchGame) GetPlayernum() uint32 {
	if m != nil && m.Playernum != nil {
		return *m.Playernum
	}
	return 0
}

func (m *SJoinMatchGame) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SJoinMatchGame_Error
}

// 比赛场结束返回排行数据
type SMatchAllOver struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8003" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,req,name=userid" json:"userid,omitempty"`
	Rank             *uint32 `protobuf:"varint,3,req,name=rank" json:"rank,omitempty"`
	Score            *int32  `protobuf:"varint,4,req,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SMatchAllOver) Reset()                    { *m = SMatchAllOver{} }
func (m *SMatchAllOver) String() string            { return proto.CompactTextString(m) }
func (*SMatchAllOver) ProtoMessage()               {}
func (*SMatchAllOver) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

const Default_SMatchAllOver_Code uint32 = 8003

func (m *SMatchAllOver) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SMatchAllOver_Code
}

func (m *SMatchAllOver) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *SMatchAllOver) GetRank() uint32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *SMatchAllOver) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 获取比赛场当前报名人数
type CUpdatePersonNum struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8004" json:"code,omitempty"`
	Kind             *uint32 `protobuf:"varint,2,req,name=kind" json:"kind,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUpdatePersonNum) Reset()                    { *m = CUpdatePersonNum{} }
func (m *CUpdatePersonNum) String() string            { return proto.CompactTextString(m) }
func (*CUpdatePersonNum) ProtoMessage()               {}
func (*CUpdatePersonNum) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

const Default_CUpdatePersonNum_Code uint32 = 8004

func (m *CUpdatePersonNum) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CUpdatePersonNum_Code
}

func (m *CUpdatePersonNum) GetKind() uint32 {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return 0
}

// 更新比赛场当前人数
type SUpdatePersonNum struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8004" json:"code,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Playernum        *uint32 `protobuf:"varint,3,req,name=playernum" json:"playernum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SUpdatePersonNum) Reset()                    { *m = SUpdatePersonNum{} }
func (m *SUpdatePersonNum) String() string            { return proto.CompactTextString(m) }
func (*SUpdatePersonNum) ProtoMessage()               {}
func (*SUpdatePersonNum) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

const Default_SUpdatePersonNum_Code uint32 = 8004

func (m *SUpdatePersonNum) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SUpdatePersonNum_Code
}

func (m *SUpdatePersonNum) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SUpdatePersonNum) GetPlayernum() uint32 {
	if m != nil && m.Playernum != nil {
		return *m.Playernum
	}
	return 0
}

// 赛事未开始前,可以退出报名
type CMatchExit struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8005" json:"code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMatchExit) Reset()                    { *m = CMatchExit{} }
func (m *CMatchExit) String() string            { return proto.CompactTextString(m) }
func (*CMatchExit) ProtoMessage()               {}
func (*CMatchExit) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

const Default_CMatchExit_Code uint32 = 8005

func (m *CMatchExit) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CMatchExit_Code
}

type SMatchExit struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8005" json:"code,omitempty"`
	Error            *uint32 `protobuf:"varint,2,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SMatchExit) Reset()                    { *m = SMatchExit{} }
func (m *SMatchExit) String() string            { return proto.CompactTextString(m) }
func (*SMatchExit) ProtoMessage()               {}
func (*SMatchExit) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

const Default_SMatchExit_Code uint32 = 8005
const Default_SMatchExit_Error uint32 = 0

func (m *SMatchExit) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SMatchExit_Code
}

func (m *SMatchExit) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SMatchExit_Error
}

// 更新比赛场的排名
type CUpdateRank struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=8006" json:"code,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUpdateRank) Reset()                    { *m = CUpdateRank{} }
func (m *CUpdateRank) String() string            { return proto.CompactTextString(m) }
func (*CUpdateRank) ProtoMessage()               {}
func (*CUpdateRank) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

const Default_CUpdateRank_Code uint32 = 8006

func (m *CUpdateRank) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CUpdateRank_Code
}

func (m *CUpdateRank) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 更新比赛场的排名
type SUpdateRank struct {
	Code             *uint32      `protobuf:"varint,1,opt,name=code,def=8006" json:"code,omitempty"`
	Id               *uint32      `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Rank             []*MatchRank `protobuf:"bytes,3,rep,name=rank" json:"rank,omitempty"`
	Num              *uint32      `protobuf:"varint,4,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SUpdateRank) Reset()                    { *m = SUpdateRank{} }
func (m *SUpdateRank) String() string            { return proto.CompactTextString(m) }
func (*SUpdateRank) ProtoMessage()               {}
func (*SUpdateRank) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

const Default_SUpdateRank_Code uint32 = 8006

func (m *SUpdateRank) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SUpdateRank_Code
}

func (m *SUpdateRank) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SUpdateRank) GetRank() []*MatchRank {
	if m != nil {
		return m.Rank
	}
	return nil
}

func (m *SUpdateRank) GetNum() uint32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type MatchRank struct {
	Rank             *uint32 `protobuf:"varint,1,req,name=rank" json:"rank,omitempty"`
	Userid           *string `protobuf:"bytes,2,req,name=userid" json:"userid,omitempty"`
	Nickname         *string `protobuf:"bytes,3,req,name=nickname" json:"nickname,omitempty"`
	Score            *int32  `protobuf:"varint,4,req,name=score" json:"score,omitempty"`
	Die              *bool   `protobuf:"varint,5,req,name=die" json:"die,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MatchRank) Reset()                    { *m = MatchRank{} }
func (m *MatchRank) String() string            { return proto.CompactTextString(m) }
func (*MatchRank) ProtoMessage()               {}
func (*MatchRank) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *MatchRank) GetRank() uint32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *MatchRank) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *MatchRank) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *MatchRank) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *MatchRank) GetDie() bool {
	if m != nil && m.Die != nil {
		return *m.Die
	}
	return false
}

// 比赛开始玩家进入房间
type SMatchComein struct {
	Code             *uint32      `protobuf:"varint,1,opt,name=code,def=8007" json:"code,omitempty"`
	Position         *uint32      `protobuf:"varint,2,req,name=Position" json:"Position,omitempty"`
	Room             *ProtoRoom   `protobuf:"bytes,3,req,name=room" json:"room,omitempty"`
	Userinfo         []*ProtoUser `protobuf:"bytes,4,rep,name=userinfo" json:"userinfo,omitempty"`
	Error            *uint32      `protobuf:"varint,5,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SMatchComein) Reset()                    { *m = SMatchComein{} }
func (m *SMatchComein) String() string            { return proto.CompactTextString(m) }
func (*SMatchComein) ProtoMessage()               {}
func (*SMatchComein) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

const Default_SMatchComein_Code uint32 = 8007
const Default_SMatchComein_Error uint32 = 0

func (m *SMatchComein) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SMatchComein_Code
}

func (m *SMatchComein) GetPosition() uint32 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func (m *SMatchComein) GetRoom() *ProtoRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *SMatchComein) GetUserinfo() []*ProtoUser {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

func (m *SMatchComein) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SMatchComein_Error
}

func init() {
	proto.RegisterType((*CJoinMatchGame)(nil), "protocol.CJoinMatchGame")
	proto.RegisterType((*SJoinMatchGame)(nil), "protocol.SJoinMatchGame")
	proto.RegisterType((*SMatchAllOver)(nil), "protocol.SMatchAllOver")
	proto.RegisterType((*CUpdatePersonNum)(nil), "protocol.CUpdatePersonNum")
	proto.RegisterType((*SUpdatePersonNum)(nil), "protocol.SUpdatePersonNum")
	proto.RegisterType((*CMatchExit)(nil), "protocol.CMatchExit")
	proto.RegisterType((*SMatchExit)(nil), "protocol.SMatchExit")
	proto.RegisterType((*CUpdateRank)(nil), "protocol.CUpdateRank")
	proto.RegisterType((*SUpdateRank)(nil), "protocol.SUpdateRank")
	proto.RegisterType((*MatchRank)(nil), "protocol.MatchRank")
	proto.RegisterType((*SMatchComein)(nil), "protocol.SMatchComein")
}

var fileDescriptor5 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6b, 0xa3, 0x40,
	0x14, 0x47, 0xa3, 0x8b, 0x3e, 0x63, 0x70, 0xdd, 0x8b, 0xec, 0xc9, 0x1d, 0x58, 0xf0, 0xb2, 0x21,
	0x9b, 0xcd, 0x7e, 0xb0, 0xb7, 0x22, 0xa1, 0x50, 0x68, 0x2b, 0x91, 0xd0, 0x53, 0x0f, 0xa2, 0x53,
	0x3a, 0xa8, 0xf3, 0xc2, 0x68, 0x42, 0xfb, 0x5f, 0xf4, 0x4f, 0x2e, 0x4e, 0xcc, 0x97, 0x95, 0x36,
	0x27, 0x75, 0x7c, 0xef, 0xf7, 0x39, 0x60, 0x95, 0x49, 0x9d, 0x3e, 0x8e, 0x57, 0x02, 0x6b, 0x74,
	0x0d, 0xf9, 0x48, 0xb1, 0xf8, 0x6a, 0x6c, 0x70, 0x7b, 0x46, 0xa6, 0x30, 0x0a, 0xaf, 0x90, 0xf1,
	0xeb, 0x66, 0xee, 0x32, 0x29, 0xa9, 0xeb, 0x82, 0x96, 0x62, 0x46, 0x3d, 0xc5, 0x57, 0x02, 0xfb,
	0xbf, 0xf6, 0x6f, 0x32, 0xf9, 0xe9, 0x0e, 0x41, 0xcb, 0x19, 0xcf, 0x3c, 0xd5, 0x57, 0x03, 0x9b,
	0xa4, 0x30, 0x8a, 0x3f, 0xde, 0x01, 0x50, 0x59, 0xbb, 0xb1, 0xdf, 0x1f, 0xc8, 0xaf, 0xcf, 0x60,
	0xae, 0x8a, 0xe4, 0x99, 0x0a, 0xbe, 0x2e, 0x3d, 0x4d, 0x1e, 0x39, 0xa0, 0x53, 0x21, 0x50, 0x78,
	0xba, 0x44, 0x50, 0x26, 0x24, 0x02, 0x3b, 0x96, 0x04, 0x17, 0x45, 0x71, 0xbb, 0xa1, 0xa2, 0x87,
	0xe3, 0x97, 0x3b, 0x82, 0x4f, 0xeb, 0x8a, 0x8a, 0x96, 0xc7, 0x6c, 0x78, 0x44, 0xc2, 0xf3, 0x96,
	0xc7, 0x06, 0xbd, 0x4a, 0x51, 0x50, 0xc9, 0xa1, 0x93, 0x19, 0x38, 0xe1, 0x72, 0x95, 0x25, 0x35,
	0x8d, 0xa8, 0xa8, 0x90, 0xdf, 0xac, 0xcb, 0x1e, 0xd0, 0x59, 0xc7, 0xec, 0x1c, 0x9c, 0xf8, 0x9c,
	0xad, 0x63, 0xbb, 0x27, 0x06, 0xa5, 0x16, 0xe2, 0x03, 0x84, 0xd2, 0xce, 0xfc, 0x89, 0xd5, 0x3d,
	0x00, 0xbf, 0xc9, 0x14, 0x20, 0x7e, 0x77, 0xe2, 0x10, 0x92, 0xba, 0x0b, 0xe9, 0x07, 0x58, 0xad,
	0xa5, 0x45, 0xc2, 0xf3, 0x9e, 0xa5, 0x3f, 0xc7, 0xba, 0xc8, 0x3d, 0x58, 0xf1, 0xf9, 0xe3, 0xee,
	0xb7, 0x7d, 0x9a, 0x83, 0xc0, 0x9a, 0x7e, 0x19, 0xef, 0xae, 0xcf, 0x58, 0xca, 0x94, 0x10, 0x16,
	0x0c, 0xf6, 0x25, 0x92, 0x3b, 0x30, 0x0f, 0x7f, 0x76, 0x55, 0x28, 0x12, 0xaa, 0x5b, 0x94, 0x03,
	0x06, 0x67, 0x69, 0xce, 0x93, 0x92, 0xca, 0x80, 0xcc, 0x4e, 0x59, 0x0d, 0x70, 0xc6, 0xa8, 0xa7,
	0xfb, 0x6a, 0x60, 0x90, 0x17, 0x05, 0x86, 0xdb, 0x6c, 0x42, 0x2c, 0x29, 0xe3, 0x3d, 0xca, 0xff,
	0x36, 0x90, 0x11, 0x56, 0xac, 0x66, 0xc8, 0x8f, 0xf4, 0x23, 0x6e, 0x1b, 0x38, 0xd1, 0x1f, 0x35,
	0x2f, 0x0b, 0xc4, 0xd2, 0xfd, 0x0e, 0x86, 0xd4, 0xc5, 0x1f, 0xd0, 0xd3, 0xba, 0x36, 0xe5, 0xd8,
	0xb2, 0xa2, 0xe2, 0xed, 0xf5, 0x7c, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x21, 0xb0, 0xa4, 0xe7, 0x59,
	0x03, 0x00, 0x00,
}
