// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/api/proto/api.proto

package go_micro_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad827baf0ac38ed4, []int{0}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad827baf0ac38ed4, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad827baf0ac38ed4, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "go.micro.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Response.HeaderEntry")
}

func init() { proto.RegisterFile("client/api/proto/api.proto", fileDescriptor_ad827baf0ac38ed4) }

var fileDescriptor_ad827baf0ac38ed4 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0xb3, 0x75, 0x0c, 0xf8, 0xf0, 0x60, 0x7a, 0x30, 0x0d, 0x07, 0xc4, 0x9d, 0x76, 0x1a,
	0x04, 0x2e, 0xca, 0xd5, 0x18, 0x8d, 0xc6, 0x84, 0xf4, 0x1f, 0x14, 0xf6, 0x05, 0x16, 0x07, 0x9d,
	0x6d, 0x67, 0xc2, 0x4f, 0xf4, 0x87, 0xf8, 0x3f, 0x4c, 0xbb, 0x81, 0xd3, 0xec, 0x86, 0xb7, 0xaf,
	0xcd, 0xfb, 0xbe, 0x7b, 0xfb, 0x7c, 0x83, 0xe1, 0x3a, 0xcf, 0x70, 0x6f, 0x26, 0xa2, 0xc8, 0x26,
	0x85, 0x92, 0x46, 0xda, 0x29, 0x71, 0x13, 0xbd, 0xd8, 0xc8, 0x64, 0x97, 0xad, 0x95, 0x4c, 0x44,
	0x91, 0x45, 0x53, 0x08, 0x96, 0x22, 0x53, 0xf4, 0x12, 0xc8, 0x1b, 0x1e, 0x98, 0x37, 0xf6, 0xe2,
	0x3e, 0xb7, 0x23, 0xbd, 0x82, 0xf0, 0x43, 0xe4, 0x25, 0x6a, 0xe6, 0x8f, 0x49, 0xdc, 0xe7, 0xf5,
	0x29, 0xfa, 0x22, 0xd0, 0xe5, 0xf8, 0x5e, 0xa2, 0x36, 0x56, 0xb3, 0x43, 0xb3, 0x95, 0x69, 0x6d,
	0xac, 0x4f, 0x94, 0x42, 0x50, 0x08, 0xb3, 0x65, 0xbe, 0xbb, 0x75, 0x33, 0xbd, 0x83, 0x70, 0x8b,
	0x22, 0x45, 0xc5, 0xc8, 0x98, 0xc4, 0x83, 0xd9, 0x4d, 0xd2, 0x2c, 0x92, 0xd4, 0x91, 0xc9, 0x93,
	0xd3, 0x3c, 0xec, 0x8d, 0x3a, 0xf0, 0xda, 0x40, 0xa7, 0x40, 0x36, 0x68, 0x58, 0xe0, 0x7c, 0xa3,
	0x76, 0xdf, 0x23, 0x9a, 0xca, 0x64, 0xa5, 0x74, 0x0e, 0x41, 0x21, 0xb5, 0x61, 0x1d, 0x67, 0xb9,
	0x6e, 0xb7, 0x2c, 0xa5, 0xae, 0x3d, 0x4e, 0x6c, 0x5b, 0xaf, 0x64, 0x7a, 0x60, 0x61, 0xd5, 0xda,
	0xce, 0x96, 0x4b, 0xa9, 0x72, 0xd6, 0xad, 0xb8, 0x94, 0x2a, 0x1f, 0xbe, 0xc2, 0xa0, 0xd1, 0xb1,
	0x05, 0x5c, 0x0c, 0x1d, 0x87, 0xca, 0xbd, 0x7e, 0x30, 0xa3, 0xbf, 0x3f, 0x6e, 0x69, 0xf3, 0x4a,
	0xb0, 0xf0, 0x6f, 0xbd, 0xe1, 0x33, 0xf4, 0x8e, 0xd5, 0xcf, 0xce, 0x7a, 0x81, 0xfe, 0xe9, 0x4d,
	0xe7, 0x86, 0x45, 0x9f, 0x1e, 0xf4, 0x38, 0xea, 0x42, 0xee, 0x35, 0xd2, 0x11, 0x80, 0x36, 0xc2,
	0x94, 0xfa, 0x5e, 0xa6, 0xe8, 0x32, 0x3b, 0xbc, 0x71, 0x43, 0x17, 0xa7, 0xe5, 0xfa, 0x8e, 0x78,
	0xf4, 0x97, 0x78, 0x95, 0xd3, 0xba, 0xdd, 0x23, 0x76, 0xf2, 0x83, 0xfd, 0x9f, 0x21, 0xaf, 0x42,
	0xf7, 0xeb, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x57, 0x77, 0xa9, 0x18, 0x03, 0x00,
	0x00,
}
