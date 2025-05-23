// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 登陆
type CLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1000" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,opt,name=userid" json:"userid,omitempty"`
	Phone            *string `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
	Password         *string `protobuf:"bytes,4,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CLogin) Reset()                    { *m = CLogin{} }
func (m *CLogin) String() string            { return proto.CompactTextString(m) }
func (*CLogin) ProtoMessage()               {}
func (*CLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

const Default_CLogin_Code uint32 = 1000

func (m *CLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CLogin_Code
}

func (m *CLogin) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *CLogin) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *CLogin) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type SLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1000" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,req,name=userid" json:"userid,omitempty"`
	Unixtime         *uint32 `protobuf:"varint,3,req,name=unixtime" json:"unixtime,omitempty"`
	Error            *uint32 `protobuf:"varint,4,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SLogin) Reset()                    { *m = SLogin{} }
func (m *SLogin) String() string            { return proto.CompactTextString(m) }
func (*SLogin) ProtoMessage()               {}
func (*SLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

const Default_SLogin_Code uint32 = 1000
const Default_SLogin_Error uint32 = 0

func (m *SLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SLogin_Code
}

func (m *SLogin) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *SLogin) GetUnixtime() uint32 {
	if m != nil && m.Unixtime != nil {
		return *m.Unixtime
	}
	return 0
}

func (m *SLogin) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SLogin_Error
}

// 咪咕登陆
type CMiguLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1001" json:"code,omitempty"`
	Sessionkey       *string `protobuf:"bytes,2,opt,name=sessionkey" json:"sessionkey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMiguLogin) Reset()                    { *m = CMiguLogin{} }
func (m *CMiguLogin) String() string            { return proto.CompactTextString(m) }
func (*CMiguLogin) ProtoMessage()               {}
func (*CMiguLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

const Default_CMiguLogin_Code uint32 = 1001

func (m *CMiguLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CMiguLogin_Code
}

func (m *CMiguLogin) GetSessionkey() string {
	if m != nil && m.Sessionkey != nil {
		return *m.Sessionkey
	}
	return ""
}

type SMiguLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1001" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,req,name=userid" json:"userid,omitempty"`
	Error            *uint32 `protobuf:"varint,3,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SMiguLogin) Reset()                    { *m = SMiguLogin{} }
func (m *SMiguLogin) String() string            { return proto.CompactTextString(m) }
func (*SMiguLogin) ProtoMessage()               {}
func (*SMiguLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

const Default_SMiguLogin_Code uint32 = 1001
const Default_SMiguLogin_Error uint32 = 0

func (m *SMiguLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SMiguLogin_Code
}

func (m *SMiguLogin) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *SMiguLogin) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SMiguLogin_Error
}

// 账户重复登录协议
type ReLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1010" json:"code,omitempty"`
	Error            *uint32 `protobuf:"varint,2,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReLogin) Reset()                    { *m = ReLogin{} }
func (m *ReLogin) String() string            { return proto.CompactTextString(m) }
func (*ReLogin) ProtoMessage()               {}
func (*ReLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

const Default_ReLogin_Code uint32 = 1010
const Default_ReLogin_Error uint32 = 0

func (m *ReLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_ReLogin_Code
}

func (m *ReLogin) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_ReLogin_Error
}

type CRegist struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1022" json:"code,omitempty"`
	Nickname         *string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Phone            *string `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
	Pwd              *string `protobuf:"bytes,4,req,name=pwd" json:"pwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CRegist) Reset()                    { *m = CRegist{} }
func (m *CRegist) String() string            { return proto.CompactTextString(m) }
func (*CRegist) ProtoMessage()               {}
func (*CRegist) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

const Default_CRegist_Code uint32 = 1022

func (m *CRegist) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CRegist_Code
}

func (m *CRegist) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *CRegist) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *CRegist) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

type SRegist struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1022" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,opt,name=userid" json:"userid,omitempty"`
	Error            *uint32 `protobuf:"varint,3,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SRegist) Reset()                    { *m = SRegist{} }
func (m *SRegist) String() string            { return proto.CompactTextString(m) }
func (*SRegist) ProtoMessage()               {}
func (*SRegist) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

const Default_SRegist_Code uint32 = 1022
const Default_SRegist_Error uint32 = 0

func (m *SRegist) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SRegist_Code
}

func (m *SRegist) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *SRegist) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SRegist_Error
}

// 修改密码
type CSetPasswd struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1023" json:"code,omitempty"`
	Pwd              *string `protobuf:"bytes,2,req,name=pwd" json:"pwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSetPasswd) Reset()                    { *m = CSetPasswd{} }
func (m *CSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*CSetPasswd) ProtoMessage()               {}
func (*CSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

const Default_CSetPasswd_Code uint32 = 1023

func (m *CSetPasswd) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CSetPasswd_Code
}

func (m *CSetPasswd) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

type SSetPasswd struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1023" json:"code,omitempty"`
	Result           *uint32 `protobuf:"varint,2,req,name=result" json:"result,omitempty"`
	Error            *uint32 `protobuf:"varint,3,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SSetPasswd) Reset()                    { *m = SSetPasswd{} }
func (m *SSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*SSetPasswd) ProtoMessage()               {}
func (*SSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

const Default_SSetPasswd_Code uint32 = 1023
const Default_SSetPasswd_Error uint32 = 0

func (m *SSetPasswd) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SSetPasswd_Code
}

func (m *SSetPasswd) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *SSetPasswd) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SSetPasswd_Error
}

// 服务器主动断开网络
type Sdisconnection struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1032" json:"code,omitempty"`
	Error            *uint32 `protobuf:"varint,2,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Sdisconnection) Reset()                    { *m = Sdisconnection{} }
func (m *Sdisconnection) String() string            { return proto.CompactTextString(m) }
func (*Sdisconnection) ProtoMessage()               {}
func (*Sdisconnection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

const Default_Sdisconnection_Code uint32 = 1032
const Default_Sdisconnection_Error uint32 = 0

func (m *Sdisconnection) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_Sdisconnection_Code
}

func (m *Sdisconnection) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_Sdisconnection_Error
}

// 微信登录登陆
type CWechatLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1024" json:"code,omitempty"`
	Appid            *string `protobuf:"bytes,2,req,name=appid" json:"appid,omitempty"`
	Secret           *string `protobuf:"bytes,3,req,name=secret" json:"secret,omitempty"`
	CodeId           *string `protobuf:"bytes,4,req,name=code_id" json:"code_id,omitempty"`
	GrantType        *string `protobuf:"bytes,5,req,name=grant_type" json:"grant_type,omitempty"`
	Token            *string `protobuf:"bytes,6,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CWechatLogin) Reset()                    { *m = CWechatLogin{} }
func (m *CWechatLogin) String() string            { return proto.CompactTextString(m) }
func (*CWechatLogin) ProtoMessage()               {}
func (*CWechatLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

const Default_CWechatLogin_Code uint32 = 1024

func (m *CWechatLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_CWechatLogin_Code
}

func (m *CWechatLogin) GetAppid() string {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return ""
}

func (m *CWechatLogin) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *CWechatLogin) GetCodeId() string {
	if m != nil && m.CodeId != nil {
		return *m.CodeId
	}
	return ""
}

func (m *CWechatLogin) GetGrantType() string {
	if m != nil && m.GrantType != nil {
		return *m.GrantType
	}
	return ""
}

func (m *CWechatLogin) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type SWechatLogin struct {
	Code             *uint32 `protobuf:"varint,1,opt,name=code,def=1024" json:"code,omitempty"`
	Userid           *string `protobuf:"bytes,2,req,name=userid" json:"userid,omitempty"`
	Token            *string `protobuf:"bytes,3,req,name=token" json:"token,omitempty"`
	Error            *uint32 `protobuf:"varint,4,opt,name=error,def=0" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SWechatLogin) Reset()                    { *m = SWechatLogin{} }
func (m *SWechatLogin) String() string            { return proto.CompactTextString(m) }
func (*SWechatLogin) ProtoMessage()               {}
func (*SWechatLogin) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

const Default_SWechatLogin_Code uint32 = 1024
const Default_SWechatLogin_Error uint32 = 0

func (m *SWechatLogin) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Default_SWechatLogin_Code
}

func (m *SWechatLogin) GetUserid() string {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return ""
}

func (m *SWechatLogin) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *SWechatLogin) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return Default_SWechatLogin_Error
}

func init() {
	proto.RegisterType((*CLogin)(nil), "protocol.CLogin")
	proto.RegisterType((*SLogin)(nil), "protocol.SLogin")
	proto.RegisterType((*CMiguLogin)(nil), "protocol.CMiguLogin")
	proto.RegisterType((*SMiguLogin)(nil), "protocol.SMiguLogin")
	proto.RegisterType((*ReLogin)(nil), "protocol.ReLogin")
	proto.RegisterType((*CRegist)(nil), "protocol.CRegist")
	proto.RegisterType((*SRegist)(nil), "protocol.SRegist")
	proto.RegisterType((*CSetPasswd)(nil), "protocol.CSetPasswd")
	proto.RegisterType((*SSetPasswd)(nil), "protocol.SSetPasswd")
	proto.RegisterType((*Sdisconnection)(nil), "protocol.Sdisconnection")
	proto.RegisterType((*CWechatLogin)(nil), "protocol.CWechatLogin")
	proto.RegisterType((*SWechatLogin)(nil), "protocol.SWechatLogin")
}

var fileDescriptor4 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xef, 0x9a, 0x40,
	0x10, 0xc5, 0xc3, 0xef, 0xaf, 0xa3, 0x50, 0xb3, 0x27, 0x8e, 0x86, 0x13, 0x97, 0xb6, 0x88, 0xa6,
	0x87, 0x5e, 0x9a, 0x94, 0x63, 0x6b, 0x62, 0xe4, 0xd0, 0xa3, 0x21, 0x30, 0xc1, 0x8d, 0xba, 0x4b,
	0x76, 0x97, 0x58, 0xff, 0xfb, 0x86, 0x55, 0x6c, 0xaa, 0x68, 0xbf, 0x27, 0x23, 0xf0, 0xde, 0x9b,
	0xf9, 0xcc, 0x83, 0xf1, 0x81, 0xd7, 0x94, 0x7d, 0x6a, 0x04, 0x57, 0x9c, 0xbc, 0xe9, 0x9f, 0x92,
	0x1f, 0xa2, 0x15, 0xb8, 0xd9, 0xcf, 0xee, 0x0d, 0x21, 0x60, 0x97, 0xbc, 0xc2, 0xd0, 0x98, 0x19,
	0xb1, 0xff, 0xd5, 0x9e, 0x27, 0x49, 0x42, 0x02, 0x70, 0x5b, 0x89, 0x82, 0x56, 0xa1, 0x39, 0x33,
	0xe2, 0x11, 0xf1, 0xc1, 0x69, 0x76, 0x9c, 0x61, 0x68, 0xe9, 0xbf, 0x53, 0x78, 0x6b, 0x0a, 0x29,
	0x4f, 0x5c, 0x54, 0xa1, 0x3d, 0x33, 0xe3, 0x51, 0xb4, 0x06, 0x37, 0x7f, 0x9f, 0x9d, 0x79, 0xd1,
	0xb7, 0x8c, 0xfe, 0x56, 0xf4, 0xd8, 0x39, 0x9a, 0xb1, 0x4f, 0xa6, 0xe0, 0xa0, 0x10, 0x5c, 0x84,
	0xb6, 0x96, 0x19, 0x49, 0xb4, 0x04, 0xc8, 0x56, 0xb4, 0x6e, 0x9f, 0xb9, 0xce, 0x09, 0x01, 0x90,
	0x28, 0x25, 0xe5, 0x6c, 0x8f, 0xe7, 0xcb, 0xa0, 0xd1, 0x77, 0x80, 0xfc, 0xb5, 0xea, 0x71, 0x96,
	0x6b, 0xb2, 0xd5, 0x27, 0x7f, 0x06, 0x6f, 0x83, 0x4f, 0x0c, 0xe6, 0xc9, 0x5f, 0x81, 0xd9, 0x0b,
	0x7e, 0x80, 0x97, 0x6d, 0xb0, 0xa6, 0x52, 0x3d, 0x0a, 0xd2, 0xb4, 0xdb, 0x96, 0xd1, 0x72, 0xcf,
	0x8a, 0x23, 0x0e, 0xe3, 0x1c, 0x83, 0xd5, 0x9c, 0x7a, 0x92, 0xdf, 0xc0, 0xcb, 0x5f, 0x98, 0xdd,
	0x5f, 0xe6, 0x71, 0xfc, 0x8f, 0x00, 0x59, 0x8e, 0x6a, 0xdd, 0x1d, 0xa8, 0x1a, 0xf0, 0x58, 0xf4,
	0x79, 0x7a, 0x7f, 0x4d, 0xec, 0xf5, 0xe7, 0x01, 0xb8, 0x02, 0x65, 0x7b, 0x50, 0x5a, 0xe1, 0x0f,
	0x44, 0x7e, 0x81, 0x20, 0xaf, 0xa8, 0x2c, 0x39, 0x63, 0x58, 0x2a, 0xca, 0x07, 0xc0, 0x2d, 0xd2,
	0x01, 0x70, 0x2d, 0x4c, 0xb2, 0x5f, 0x58, 0xee, 0x0a, 0xf5, 0x04, 0x77, 0xba, 0xec, 0x58, 0x15,
	0x4d, 0x73, 0x3b, 0x57, 0x00, 0xae, 0xc4, 0x52, 0xa0, 0xd2, 0xc5, 0x19, 0x91, 0x0f, 0xe0, 0x75,
	0x92, 0x2d, 0xbd, 0xf2, 0xeb, 0x5a, 0x51, 0x8b, 0x82, 0xa9, 0xad, 0x3a, 0x37, 0x18, 0x3a, 0xfa,
	0x99, 0x0f, 0x8e, 0xe2, 0x7b, 0x64, 0xa1, 0xab, 0x57, 0xce, 0x61, 0x92, 0xff, 0x2f, 0xf6, 0xbe,
	0x26, 0x37, 0x0b, 0xeb, 0xdf, 0xd6, 0xf4, 0x7d, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x3b,
	0x1b, 0x2f, 0x68, 0x03, 0x00, 0x00,
}
