// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: alert/alert.proto

package alert

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event is inspired by Google Analytics events
// https://developers.google.com/analytics/devguides/collection/analyticsjs/events
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is not required for inserts
	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category string            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Action   string            `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Label    string            `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Value    uint64            `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_alert_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_alert_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Event) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Event) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Event) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Event) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ReportEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ReportEventRequest) Reset() {
	*x = ReportEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventRequest) ProtoMessage() {}

func (x *ReportEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventRequest.ProtoReflect.Descriptor instead.
func (*ReportEventRequest) Descriptor() ([]byte, []int) {
	return file_alert_alert_proto_rawDescGZIP(), []int{1}
}

func (x *ReportEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type ReportEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportEventResponse) Reset() {
	*x = ReportEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventResponse) ProtoMessage() {}

func (x *ReportEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventResponse.ProtoReflect.Descriptor instead.
func (*ReportEventResponse) Descriptor() ([]byte, []int) {
	return file_alert_alert_proto_rawDescGZIP(), []int{2}
}

var File_alert_alert_proto protoreflect.FileDescriptor

var file_alert_alert_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4f, 0x0a, 0x05, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alert_alert_proto_rawDescOnce sync.Once
	file_alert_alert_proto_rawDescData = file_alert_alert_proto_rawDesc
)

func file_alert_alert_proto_rawDescGZIP() []byte {
	file_alert_alert_proto_rawDescOnce.Do(func() {
		file_alert_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_alert_alert_proto_rawDescData)
	})
	return file_alert_alert_proto_rawDescData
}

var file_alert_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_alert_alert_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: alert.Event
	(*ReportEventRequest)(nil),  // 1: alert.ReportEventRequest
	(*ReportEventResponse)(nil), // 2: alert.ReportEventResponse
	nil,                         // 3: alert.Event.MetadataEntry
}
var file_alert_alert_proto_depIdxs = []int32{
	3, // 0: alert.Event.metadata:type_name -> alert.Event.MetadataEntry
	0, // 1: alert.ReportEventRequest.event:type_name -> alert.Event
	1, // 2: alert.Alert.ReportEvent:input_type -> alert.ReportEventRequest
	2, // 3: alert.Alert.ReportEvent:output_type -> alert.ReportEventResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_alert_alert_proto_init() }
func file_alert_alert_proto_init() {
	if File_alert_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alert_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alert_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alert_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alert_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alert_alert_proto_goTypes,
		DependencyIndexes: file_alert_alert_proto_depIdxs,
		MessageInfos:      file_alert_alert_proto_msgTypes,
	}.Build()
	File_alert_alert_proto = out.File
	file_alert_alert_proto_rawDesc = nil
	file_alert_alert_proto_goTypes = nil
	file_alert_alert_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlertClient is the client API for Alert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertClient interface {
	// ReportEvent does event ingestions.
	ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*ReportEventResponse, error)
}

type alertClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertClient(cc grpc.ClientConnInterface) AlertClient {
	return &alertClient{cc}
}

func (c *alertClient) ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*ReportEventResponse, error) {
	out := new(ReportEventResponse)
	err := c.cc.Invoke(ctx, "/alert.Alert/ReportEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServer is the server API for Alert service.
type AlertServer interface {
	// ReportEvent does event ingestions.
	ReportEvent(context.Context, *ReportEventRequest) (*ReportEventResponse, error)
}

// UnimplementedAlertServer can be embedded to have forward compatible implementations.
type UnimplementedAlertServer struct {
}

func (*UnimplementedAlertServer) ReportEvent(context.Context, *ReportEventRequest) (*ReportEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportEvent not implemented")
}

func RegisterAlertServer(s *grpc.Server, srv AlertServer) {
	s.RegisterService(&_Alert_serviceDesc, srv)
}

func _Alert_ReportEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).ReportEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alert/ReportEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).ReportEvent(ctx, req.(*ReportEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Alert",
	HandlerType: (*AlertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportEvent",
			Handler:    _Alert_ReportEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert/alert.proto",
}
