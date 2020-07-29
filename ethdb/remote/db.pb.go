// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: remote/db.proto

package remote

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SizeRequest) Reset() {
	*x = SizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeRequest) ProtoMessage() {}

func (x *SizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeRequest.ProtoReflect.Descriptor instead.
func (*SizeRequest) Descriptor() ([]byte, []int) {
	return file_remote_db_proto_rawDescGZIP(), []int{0}
}

type SizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SizeReply) Reset() {
	*x = SizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeReply) ProtoMessage() {}

func (x *SizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeReply.ProtoReflect.Descriptor instead.
func (*SizeReply) Descriptor() ([]byte, []int) {
	return file_remote_db_proto_rawDescGZIP(), []int{1}
}

func (x *SizeReply) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type BucketSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName []byte `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
}

func (x *BucketSizeRequest) Reset() {
	*x = BucketSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketSizeRequest) ProtoMessage() {}

func (x *BucketSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketSizeRequest.ProtoReflect.Descriptor instead.
func (*BucketSizeRequest) Descriptor() ([]byte, []int) {
	return file_remote_db_proto_rawDescGZIP(), []int{2}
}

func (x *BucketSizeRequest) GetBucketName() []byte {
	if x != nil {
		return x.BucketName
	}
	return nil
}

type BucketSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *BucketSizeReply) Reset() {
	*x = BucketSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketSizeReply) ProtoMessage() {}

func (x *BucketSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketSizeReply.ProtoReflect.Descriptor instead.
func (*BucketSizeReply) Descriptor() ([]byte, []int) {
	return file_remote_db_proto_rawDescGZIP(), []int{3}
}

func (x *BucketSizeReply) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_remote_db_proto protoreflect.FileDescriptor

var file_remote_db_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x25,
	0x0a, 0x0f, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0x76, 0x0a, 0x02, 0x44, 0x42, 0x12, 0x2e, 0x0a, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x0a,
	0x10, 0x69, 0x6f, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x65, 0x74, 0x68, 0x2e, 0x64,
	0x62, 0x42, 0x02, 0x44, 0x42, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_db_proto_rawDescOnce sync.Once
	file_remote_db_proto_rawDescData = file_remote_db_proto_rawDesc
)

func file_remote_db_proto_rawDescGZIP() []byte {
	file_remote_db_proto_rawDescOnce.Do(func() {
		file_remote_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_db_proto_rawDescData)
	})
	return file_remote_db_proto_rawDescData
}

var file_remote_db_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_db_proto_goTypes = []interface{}{
	(*SizeRequest)(nil),       // 0: remote.SizeRequest
	(*SizeReply)(nil),         // 1: remote.SizeReply
	(*BucketSizeRequest)(nil), // 2: remote.BucketSizeRequest
	(*BucketSizeReply)(nil),   // 3: remote.BucketSizeReply
}
var file_remote_db_proto_depIdxs = []int32{
	0, // 0: remote.DB.Size:input_type -> remote.SizeRequest
	2, // 1: remote.DB.BucketSize:input_type -> remote.BucketSizeRequest
	1, // 2: remote.DB.Size:output_type -> remote.SizeReply
	3, // 3: remote.DB.BucketSize:output_type -> remote.BucketSizeReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_remote_db_proto_init() }
func file_remote_db_proto_init() {
	if File_remote_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeRequest); i {
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
		file_remote_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeReply); i {
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
		file_remote_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketSizeRequest); i {
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
		file_remote_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketSizeReply); i {
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
			RawDescriptor: file_remote_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_db_proto_goTypes,
		DependencyIndexes: file_remote_db_proto_depIdxs,
		MessageInfos:      file_remote_db_proto_msgTypes,
	}.Build()
	File_remote_db_proto = out.File
	file_remote_db_proto_rawDesc = nil
	file_remote_db_proto_goTypes = nil
	file_remote_db_proto_depIdxs = nil
}
