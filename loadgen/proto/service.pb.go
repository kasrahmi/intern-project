// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/service.proto

package proto

import (
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

// The request message.
type MyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DurationSeconds int32  `protobuf:"varint,2,opt,name=duration_seconds,json=durationSeconds,proto3" json:"duration_seconds,omitempty"` // Duration of CPU-bound task
}

func (x *MyRequest) Reset() {
	*x = MyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyRequest) ProtoMessage() {}

func (x *MyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyRequest.ProtoReflect.Descriptor instead.
func (*MyRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *MyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MyRequest) GetDurationSeconds() int32 {
	if x != nil {
		return x.DurationSeconds
	}
	return 0
}

// The response message.
type MyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	EndToEndLatencyMs int64  `protobuf:"varint,2,opt,name=end_to_end_latency_ms,json=endToEndLatencyMs,proto3" json:"end_to_end_latency_ms,omitempty"`
	WorkerId          string `protobuf:"bytes,3,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"` // Add Worker ID to the response
}

func (x *MyResponse) Reset() {
	*x = MyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyResponse) ProtoMessage() {}

func (x *MyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyResponse.ProtoReflect.Descriptor instead.
func (*MyResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *MyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MyResponse) GetEndToEndLatencyMs() int64 {
	if x != nil {
		return x.EndToEndLatencyMs
	}
	return 0
}

func (x *MyResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x4a, 0x0a, 0x09, 0x4d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x75, 0x0a, 0x0a,
	0x4d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x15, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x45, 0x6e, 0x64, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x4d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x32, 0x4a, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_proto_goTypes = []any{
	(*MyRequest)(nil),  // 0: myproject.MyRequest
	(*MyResponse)(nil), // 1: myproject.MyResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: myproject.MyService.ProcessRequest:input_type -> myproject.MyRequest
	1, // 1: myproject.MyService.ProcessRequest:output_type -> myproject.MyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MyRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MyResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
