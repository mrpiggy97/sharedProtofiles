// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protofiles/Calculation.proto

package calculation

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

type SumStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SumStreamRequest) Reset() {
	*x = SumStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Calculation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumStreamRequest) ProtoMessage() {}

func (x *SumStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Calculation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumStreamRequest.ProtoReflect.Descriptor instead.
func (*SumStreamRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_Calculation_proto_rawDescGZIP(), []int{0}
}

func (x *SumStreamRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SumStreamRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type SumStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,3,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumStreamResponse) Reset() {
	*x = SumStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Calculation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumStreamResponse) ProtoMessage() {}

func (x *SumStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Calculation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumStreamResponse.ProtoReflect.Descriptor instead.
func (*SumStreamResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_Calculation_proto_rawDescGZIP(), []int{1}
}

func (x *SumStreamResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_protofiles_Calculation_proto protoreflect.FileDescriptor

var file_protofiles_Calculation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x6d, 0x79, 0x70, 0x6b, 0x67, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x62, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x74, 0x0a, 0x12, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5e, 0x0a, 0x09, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x2e,
	0x6d, 0x79, 0x70, 0x6b, 0x67, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x79, 0x70, 0x6b, 0x67, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_Calculation_proto_rawDescOnce sync.Once
	file_protofiles_Calculation_proto_rawDescData = file_protofiles_Calculation_proto_rawDesc
)

func file_protofiles_Calculation_proto_rawDescGZIP() []byte {
	file_protofiles_Calculation_proto_rawDescOnce.Do(func() {
		file_protofiles_Calculation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_Calculation_proto_rawDescData)
	})
	return file_protofiles_Calculation_proto_rawDescData
}

var file_protofiles_Calculation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protofiles_Calculation_proto_goTypes = []interface{}{
	(*SumStreamRequest)(nil),  // 0: mypkg.calculator.v1.SumStreamRequest
	(*SumStreamResponse)(nil), // 1: mypkg.calculator.v1.SumStreamResponse
}
var file_protofiles_Calculation_proto_depIdxs = []int32{
	0, // 0: mypkg.calculator.v1.CalculationService.SumStream:input_type -> mypkg.calculator.v1.SumStreamRequest
	1, // 1: mypkg.calculator.v1.CalculationService.SumStream:output_type -> mypkg.calculator.v1.SumStreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofiles_Calculation_proto_init() }
func file_protofiles_Calculation_proto_init() {
	if File_protofiles_Calculation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_Calculation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumStreamRequest); i {
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
		file_protofiles_Calculation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumStreamResponse); i {
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
			RawDescriptor: file_protofiles_Calculation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_Calculation_proto_goTypes,
		DependencyIndexes: file_protofiles_Calculation_proto_depIdxs,
		MessageInfos:      file_protofiles_Calculation_proto_msgTypes,
	}.Build()
	File_protofiles_Calculation_proto = out.File
	file_protofiles_Calculation_proto_rawDesc = nil
	file_protofiles_Calculation_proto_goTypes = nil
	file_protofiles_Calculation_proto_depIdxs = nil
}
