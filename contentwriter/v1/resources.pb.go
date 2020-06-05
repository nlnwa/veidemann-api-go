// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: contentwriter/v1/resources.proto

package contentwriter

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RecordType int32

const (
	RecordType_WARCINFO     RecordType = 0
	RecordType_RESPONSE     RecordType = 1
	RecordType_RESOURCE     RecordType = 2
	RecordType_REQUEST      RecordType = 3
	RecordType_METADATA     RecordType = 4
	RecordType_REVISIT      RecordType = 5
	RecordType_CONVERSION   RecordType = 6
	RecordType_CONTINUATION RecordType = 7
)

// Enum value maps for RecordType.
var (
	RecordType_name = map[int32]string{
		0: "WARCINFO",
		1: "RESPONSE",
		2: "RESOURCE",
		3: "REQUEST",
		4: "METADATA",
		5: "REVISIT",
		6: "CONVERSION",
		7: "CONTINUATION",
	}
	RecordType_value = map[string]int32{
		"WARCINFO":     0,
		"RESPONSE":     1,
		"RESOURCE":     2,
		"REQUEST":      3,
		"METADATA":     4,
		"REVISIT":      5,
		"CONVERSION":   6,
		"CONTINUATION": 7,
	}
)

func (x RecordType) Enum() *RecordType {
	p := new(RecordType)
	*p = x
	return p
}

func (x RecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_contentwriter_v1_resources_proto_enumTypes[0].Descriptor()
}

func (RecordType) Type() protoreflect.EnumType {
	return &file_contentwriter_v1_resources_proto_enumTypes[0]
}

func (x RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordType.Descriptor instead.
func (RecordType) EnumDescriptor() ([]byte, []int) {
	return file_contentwriter_v1_resources_proto_rawDescGZIP(), []int{0}
}

type CrawledContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest    string               `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	WarcId    string               `protobuf:"bytes,2,opt,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	TargetUri string               `protobuf:"bytes,3,opt,name=target_uri,json=targetUri,proto3" json:"target_uri,omitempty"`
	Date      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CrawledContent) Reset() {
	*x = CrawledContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contentwriter_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawledContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawledContent) ProtoMessage() {}

func (x *CrawledContent) ProtoReflect() protoreflect.Message {
	mi := &file_contentwriter_v1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawledContent.ProtoReflect.Descriptor instead.
func (*CrawledContent) Descriptor() ([]byte, []int) {
	return file_contentwriter_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *CrawledContent) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CrawledContent) GetWarcId() string {
	if x != nil {
		return x.WarcId
	}
	return ""
}

func (x *CrawledContent) GetTargetUri() string {
	if x != nil {
		return x.TargetUri
	}
	return ""
}

func (x *CrawledContent) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type StorageRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarcId     string     `protobuf:"bytes,1,opt,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	RecordType RecordType `protobuf:"varint,2,opt,name=record_type,json=recordType,proto3,enum=veidemann.api.contentwriter.v1.RecordType" json:"record_type,omitempty"`
	StorageRef string     `protobuf:"bytes,3,opt,name=storage_ref,json=storageRef,proto3" json:"storage_ref,omitempty"`
}

func (x *StorageRef) Reset() {
	*x = StorageRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contentwriter_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageRef) ProtoMessage() {}

func (x *StorageRef) ProtoReflect() protoreflect.Message {
	mi := &file_contentwriter_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageRef.ProtoReflect.Descriptor instead.
func (*StorageRef) Descriptor() ([]byte, []int) {
	return file_contentwriter_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *StorageRef) GetWarcId() string {
	if x != nil {
		return x.WarcId
	}
	return ""
}

func (x *StorageRef) GetRecordType() RecordType {
	if x != nil {
		return x.RecordType
	}
	return RecordType_WARCINFO
}

func (x *StorageRef) GetStorageRef() string {
	if x != nil {
		return x.StorageRef
	}
	return ""
}

var File_contentwriter_v1_resources_proto protoreflect.FileDescriptor

var file_contentwriter_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x77, 0x61, 0x72, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x77, 0x61, 0x72, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x72, 0x63, 0x49, 0x64, 0x12, 0x4b,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x2a, 0x80, 0x01, 0x0a,
	0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x57,
	0x41, 0x52, 0x43, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x04,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x56, 0x49, 0x53, 0x49, 0x54, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x42,
	0x86, 0x01, 0x0a, 0x28, 0x6e, 0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contentwriter_v1_resources_proto_rawDescOnce sync.Once
	file_contentwriter_v1_resources_proto_rawDescData = file_contentwriter_v1_resources_proto_rawDesc
)

func file_contentwriter_v1_resources_proto_rawDescGZIP() []byte {
	file_contentwriter_v1_resources_proto_rawDescOnce.Do(func() {
		file_contentwriter_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_contentwriter_v1_resources_proto_rawDescData)
	})
	return file_contentwriter_v1_resources_proto_rawDescData
}

var file_contentwriter_v1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contentwriter_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contentwriter_v1_resources_proto_goTypes = []interface{}{
	(RecordType)(0),             // 0: veidemann.api.contentwriter.v1.RecordType
	(*CrawledContent)(nil),      // 1: veidemann.api.contentwriter.v1.CrawledContent
	(*StorageRef)(nil),          // 2: veidemann.api.contentwriter.v1.StorageRef
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_contentwriter_v1_resources_proto_depIdxs = []int32{
	3, // 0: veidemann.api.contentwriter.v1.CrawledContent.date:type_name -> google.protobuf.Timestamp
	0, // 1: veidemann.api.contentwriter.v1.StorageRef.record_type:type_name -> veidemann.api.contentwriter.v1.RecordType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contentwriter_v1_resources_proto_init() }
func file_contentwriter_v1_resources_proto_init() {
	if File_contentwriter_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contentwriter_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawledContent); i {
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
		file_contentwriter_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageRef); i {
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
			RawDescriptor: file_contentwriter_v1_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contentwriter_v1_resources_proto_goTypes,
		DependencyIndexes: file_contentwriter_v1_resources_proto_depIdxs,
		EnumInfos:         file_contentwriter_v1_resources_proto_enumTypes,
		MessageInfos:      file_contentwriter_v1_resources_proto_msgTypes,
	}.Build()
	File_contentwriter_v1_resources_proto = out.File
	file_contentwriter_v1_resources_proto_rawDesc = nil
	file_contentwriter_v1_resources_proto_goTypes = nil
	file_contentwriter_v1_resources_proto_depIdxs = nil
}
