// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/login/merge.proto

// 账户合并相关

package login

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TelephoneRequest struct {
	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 新手机号
	TargetTelephone string `protobuf:"bytes,2,opt,name=targetTelephone,proto3" json:"targetTelephone,omitempty"`
	// 是否走新流程
	IsNewProcess         bool     `protobuf:"varint,3,opt,name=isNewProcess,proto3" json:"isNewProcess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelephoneRequest) Reset()         { *m = TelephoneRequest{} }
func (m *TelephoneRequest) String() string { return proto.CompactTextString(m) }
func (*TelephoneRequest) ProtoMessage()    {}
func (*TelephoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885f2e802da2b346, []int{0}
}

func (m *TelephoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelephoneRequest.Unmarshal(m, b)
}
func (m *TelephoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelephoneRequest.Marshal(b, m, deterministic)
}
func (m *TelephoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelephoneRequest.Merge(m, src)
}
func (m *TelephoneRequest) XXX_Size() int {
	return xxx_messageInfo_TelephoneRequest.Size(m)
}
func (m *TelephoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TelephoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TelephoneRequest proto.InternalMessageInfo

func (m *TelephoneRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TelephoneRequest) GetTargetTelephone() string {
	if m != nil {
		return m.TargetTelephone
	}
	return ""
}

func (m *TelephoneRequest) GetIsNewProcess() bool {
	if m != nil {
		return m.IsNewProcess
	}
	return false
}

type WeChatRequest struct {
	// 登陆用户
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 要合并的用户
	TargetUid            int64    `protobuf:"varint,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatRequest) Reset()         { *m = WeChatRequest{} }
func (m *WeChatRequest) String() string { return proto.CompactTextString(m) }
func (*WeChatRequest) ProtoMessage()    {}
func (*WeChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885f2e802da2b346, []int{1}
}

func (m *WeChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatRequest.Unmarshal(m, b)
}
func (m *WeChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatRequest.Marshal(b, m, deterministic)
}
func (m *WeChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatRequest.Merge(m, src)
}
func (m *WeChatRequest) XXX_Size() int {
	return xxx_messageInfo_WeChatRequest.Size(m)
}
func (m *WeChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatRequest proto.InternalMessageInfo

func (m *WeChatRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *WeChatRequest) GetTargetUid() int64 {
	if m != nil {
		return m.TargetUid
	}
	return 0
}

type WeChatUnMergeRequest struct {
	// 登陆用户
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatUnMergeRequest) Reset()         { *m = WeChatUnMergeRequest{} }
func (m *WeChatUnMergeRequest) String() string { return proto.CompactTextString(m) }
func (*WeChatUnMergeRequest) ProtoMessage()    {}
func (*WeChatUnMergeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885f2e802da2b346, []int{2}
}

func (m *WeChatUnMergeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatUnMergeRequest.Unmarshal(m, b)
}
func (m *WeChatUnMergeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatUnMergeRequest.Marshal(b, m, deterministic)
}
func (m *WeChatUnMergeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatUnMergeRequest.Merge(m, src)
}
func (m *WeChatUnMergeRequest) XXX_Size() int {
	return xxx_messageInfo_WeChatUnMergeRequest.Size(m)
}
func (m *WeChatUnMergeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatUnMergeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatUnMergeRequest proto.InternalMessageInfo

func (m *WeChatUnMergeRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type Reply struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_885f2e802da2b346, []int{3}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Reply) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *Reply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TelephoneRequest)(nil), "login.TelephoneRequest")
	proto.RegisterType((*WeChatRequest)(nil), "login.WeChatRequest")
	proto.RegisterType((*WeChatUnMergeRequest)(nil), "login.WeChatUnMergeRequest")
	proto.RegisterType((*Reply)(nil), "login.Reply")
	proto.RegisterMapType((map[string]string)(nil), "login.Reply.DataEntry")
}

func init() {
	proto.RegisterFile("examples/proto/login/merge.proto", fileDescriptor_885f2e802da2b346)
}

var fileDescriptor_885f2e802da2b346 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0x38, 0x69, 0x9b, 0x4d, 0x0b, 0xd1, 0x36, 0x40, 0x9a, 0x46, 0xb5, 0xb1, 0xf8,
	0x88, 0x8a, 0x12, 0x8b, 0x52, 0x09, 0x14, 0x89, 0x4b, 0x0a, 0x37, 0x40, 0x68, 0x95, 0x2a, 0x47,
	0xb4, 0x71, 0x06, 0xc7, 0xaa, 0xed, 0x0d, 0xf6, 0x9a, 0x90, 0x2b, 0xaf, 0xc0, 0x1b, 0xf0, 0x38,
	0x5c, 0x91, 0x38, 0xfa, 0xc4, 0xc9, 0x4f, 0x81, 0x76, 0xd7, 0x49, 0x1a, 0x2b, 0x12, 0xb9, 0x8c,
	0x66, 0x66, 0x77, 0x7e, 0xff, 0xf1, 0xce, 0x18, 0x99, 0xf0, 0x8d, 0x06, 0x33, 0x1f, 0x62, 0x7b,
	0x16, 0x31, 0xce, 0x6c, 0x9f, 0xb9, 0x5e, 0x68, 0x07, 0x10, 0xb9, 0xd0, 0x93, 0x19, 0x5c, 0x91,
	0xa9, 0x56, 0xd7, 0xf5, 0xf8, 0x34, 0x19, 0xf7, 0x1c, 0x16, 0xd8, 0x2e, 0x73, 0x99, 0xba, 0x3f,
	0x4e, 0x3e, 0xcb, 0x48, 0x15, 0x0b, 0x4f, 0x55, 0xb5, 0xda, 0x2e, 0x63, 0xae, 0x0f, 0x36, 0x9d,
	0x79, 0x36, 0x0d, 0x43, 0xc6, 0x29, 0xf7, 0x58, 0x18, 0xab, 0x53, 0xeb, 0xa7, 0x86, 0xea, 0x43,
	0xf0, 0x61, 0x36, 0x65, 0x21, 0x10, 0xf8, 0x92, 0x40, 0xcc, 0xf1, 0x09, 0xd2, 0x13, 0x6f, 0xd2,
	0xd4, 0x4c, 0xad, 0xa3, 0x0f, 0xf6, 0xb3, 0xd4, 0x10, 0x21, 0x11, 0x06, 0xbf, 0x46, 0x77, 0x39,
	0x8d, 0x5c, 0xe0, 0xab, 0xa2, 0x66, 0xc9, 0xd4, 0x3a, 0xd5, 0xc1, 0x71, 0x96, 0x1a, 0xc5, 0x23,
	0x52, 0x4c, 0xe0, 0x4b, 0x74, 0xe8, 0xc5, 0x1f, 0x60, 0xfe, 0x31, 0x62, 0x0e, 0xc4, 0x71, 0x53,
	0x37, 0xb5, 0xce, 0xc1, 0xa0, 0x9e, 0xa5, 0xc6, 0x46, 0x9e, 0x6c, 0x44, 0xd6, 0x08, 0x1d, 0x8d,
	0xe0, 0x6a, 0x4a, 0xf9, 0x0e, 0x0d, 0x3e, 0x43, 0x55, 0x25, 0x7a, 0xed, 0x4d, 0x64, 0x6b, 0xfa,
	0xe0, 0x28, 0x4b, 0x8d, 0x75, 0x92, 0xac, 0x5d, 0xeb, 0x39, 0x6a, 0x28, 0xf0, 0x75, 0xf8, 0x5e,
	0x3c, 0xf4, 0xff, 0xf9, 0xd6, 0x1f, 0x0d, 0x55, 0x08, 0xcc, 0xfc, 0x05, 0x6e, 0xa3, 0xb2, 0xc3,
	0x26, 0x90, 0xdf, 0x3a, 0xc8, 0x52, 0x43, 0xc6, 0x44, 0x5a, 0x81, 0x08, 0x62, 0x37, 0x7f, 0x1c,
	0x89, 0x08, 0x62, 0x97, 0x08, 0x83, 0x1f, 0xa3, 0xfd, 0x90, 0xcd, 0x87, 0x5e, 0x00, 0xf2, 0xfb,
	0xf5, 0x41, 0x2d, 0x4b, 0x8d, 0x65, 0x8a, 0x2c, 0x1d, 0x7c, 0x89, 0xca, 0x13, 0xca, 0x69, 0xb3,
	0x6c, 0xea, 0x9d, 0xda, 0xc5, 0xfd, 0x9e, 0x9c, 0x7e, 0x4f, 0x6a, 0xf7, 0xde, 0x50, 0x4e, 0xdf,
	0x86, 0x3c, 0x5a, 0x28, 0x5d, 0x71, 0x8f, 0x48, 0xdb, 0x7a, 0x89, 0xaa, 0xab, 0x43, 0x5c, 0x47,
	0xfa, 0x0d, 0x2c, 0x64, 0x87, 0x55, 0x22, 0x5c, 0xdc, 0x40, 0x95, 0xaf, 0xd4, 0x4f, 0xf2, 0xa9,
	0x11, 0x15, 0xf4, 0x4b, 0xaf, 0xb4, 0x8b, 0x5f, 0x3a, 0xaa, 0xc8, 0x47, 0xc0, 0x43, 0x54, 0x5d,
	0x4f, 0xec, 0x41, 0xae, 0x5b, 0x5c, 0x92, 0xd6, 0xe1, 0xed, 0x86, 0x2c, 0xf3, 0xfb, 0xef, 0xbf,
	0x3f, 0x4a, 0x2d, 0xeb, 0x9e, 0x9d, 0xc4, 0x10, 0xa9, 0xb5, 0xb5, 0xf9, 0xb2, 0xa6, 0xaf, 0x9d,
	0xe3, 0x4f, 0xe8, 0xce, 0x8a, 0x71, 0x35, 0x05, 0xe7, 0x66, 0x57, 0xf4, 0x13, 0x89, 0x36, 0xad,
	0xd3, 0xad, 0xe8, 0xae, 0x23, 0x58, 0x42, 0xe0, 0x1d, 0xda, 0x53, 0xc3, 0xc4, 0x8d, 0xbc, 0x7e,
	0x63, 0x69, 0x0a, 0xd4, 0x33, 0x49, 0x6d, 0x5a, 0xc7, 0xb7, 0xa9, 0x73, 0x81, 0xa3, 0x5c, 0xd0,
	0x46, 0xa8, 0xa6, 0xca, 0x55, 0xaf, 0xbb, 0x20, 0x1f, 0x49, 0xe4, 0x99, 0x75, 0xb2, 0x05, 0xb9,
	0x6e, 0xd3, 0x59, 0x2e, 0x73, 0xbe, 0x73, 0xf8, 0x74, 0x03, 0xbd, 0xb9, 0x89, 0x05, 0x85, 0xa7,
	0x52, 0xe1, 0xa1, 0xd5, 0xde, 0xa6, 0x90, 0x84, 0x5d, 0x99, 0xe8, 0x6b, 0xe7, 0xe3, 0x3d, 0xf9,
	0x77, 0xbf, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x65, 0x0d, 0x69, 0x55, 0x04, 0x00, 0x00,
}
