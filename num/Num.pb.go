// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protofiles/Num.proto

package num

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

type NumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To     int64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Number int64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumRequest) Reset() {
	*x = NumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Num_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumRequest) ProtoMessage() {}

func (x *NumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Num_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumRequest.ProtoReflect.Descriptor instead.
func (*NumRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_Num_proto_rawDescGZIP(), []int{0}
}

func (x *NumRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *NumRequest) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *NumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type NumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentNumber int64 `protobuf:"varint,1,opt,name=currentNumber,proto3" json:"currentNumber,omitempty"`
	Remaining     int64 `protobuf:"varint,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
}

func (x *NumResponse) Reset() {
	*x = NumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Num_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumResponse) ProtoMessage() {}

func (x *NumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Num_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumResponse.ProtoReflect.Descriptor instead.
func (*NumResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_Num_proto_rawDescGZIP(), []int{1}
}

func (x *NumResponse) GetCurrentNumber() int64 {
	if x != nil {
		return x.CurrentNumber
	}
	return 0
}

func (x *NumResponse) GetRemaining() int64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Num_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Num_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_Num_proto_rawDescGZIP(), []int{2}
}

func (x *SumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_Num_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_Num_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_Num_proto_rawDescGZIP(), []int{3}
}

func (x *SumResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_protofiles_Num_proto protoreflect.FileDescriptor

var file_protofiles_Num_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x4e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22,
	0x48, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x0b, 0x4e, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x24, 0x0a, 0x0a,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x23, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x74, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x52, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x03, 0x53, 0x75, 0x6d,
	0x12, 0x13, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_Num_proto_rawDescOnce sync.Once
	file_protofiles_Num_proto_rawDescData = file_protofiles_Num_proto_rawDesc
)

func file_protofiles_Num_proto_rawDescGZIP() []byte {
	file_protofiles_Num_proto_rawDescOnce.Do(func() {
		file_protofiles_Num_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_Num_proto_rawDescData)
	})
	return file_protofiles_Num_proto_rawDescData
}

var file_protofiles_Num_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protofiles_Num_proto_goTypes = []interface{}{
	(*NumRequest)(nil),  // 0: numbers.NumRequest
	(*NumResponse)(nil), // 1: numbers.NumResponse
	(*SumRequest)(nil),  // 2: numbers.SumRequest
	(*SumResponse)(nil), // 3: numbers.SumResponse
}
var file_protofiles_Num_proto_depIdxs = []int32{
	0, // 0: numbers.NumService.Rnd:input_type -> numbers.NumRequest
	2, // 1: numbers.NumService.Sum:input_type -> numbers.SumRequest
	1, // 2: numbers.NumService.Rnd:output_type -> numbers.NumResponse
	3, // 3: numbers.NumService.Sum:output_type -> numbers.SumResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofiles_Num_proto_init() }
func file_protofiles_Num_proto_init() {
	if File_protofiles_Num_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_Num_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumRequest); i {
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
		file_protofiles_Num_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumResponse); i {
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
		file_protofiles_Num_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_protofiles_Num_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
			RawDescriptor: file_protofiles_Num_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_Num_proto_goTypes,
		DependencyIndexes: file_protofiles_Num_proto_depIdxs,
		MessageInfos:      file_protofiles_Num_proto_msgTypes,
	}.Build()
	File_protofiles_Num_proto = out.File
	file_protofiles_Num_proto_rawDesc = nil
	file_protofiles_Num_proto_goTypes = nil
	file_protofiles_Num_proto_depIdxs = nil
}
