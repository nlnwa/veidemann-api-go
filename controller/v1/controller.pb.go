// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: controller/v1/controller.proto

package controller

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/nlnwa/veidemann-api-go/config/v1"
	v11 "github.com/nlnwa/veidemann-api-go/frontier/v1"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RunStatus int32

const (
	RunStatus_RUNNING         RunStatus = 0
	RunStatus_PAUSED          RunStatus = 1
	RunStatus_PAUSE_REQUESTED RunStatus = 2
)

// Enum value maps for RunStatus.
var (
	RunStatus_name = map[int32]string{
		0: "RUNNING",
		1: "PAUSED",
		2: "PAUSE_REQUESTED",
	}
	RunStatus_value = map[string]int32{
		"RUNNING":         0,
		"PAUSED":          1,
		"PAUSE_REQUESTED": 2,
	}
)

func (x RunStatus) Enum() *RunStatus {
	p := new(RunStatus)
	*p = x
	return p
}

func (x RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_v1_controller_proto_enumTypes[0].Descriptor()
}

func (RunStatus) Type() protoreflect.EnumType {
	return &file_controller_v1_controller_proto_enumTypes[0]
}

func (x RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStatus.Descriptor instead.
func (RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{0}
}

// Kick of a crawl job immediately
// If a job is already running for this job_id, then new seeds are added to the job instead of starting a new one.
type RunCrawlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// If seed id is submitted, only this seed will be harvested.
	// If empty, all seeds configured with the submitted job id will be harvested.
	SeedId string `protobuf:"bytes,6,opt,name=seed_id,json=seedId,proto3" json:"seed_id,omitempty"`
}

func (x *RunCrawlRequest) Reset() {
	*x = RunCrawlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCrawlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCrawlRequest) ProtoMessage() {}

func (x *RunCrawlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCrawlRequest.ProtoReflect.Descriptor instead.
func (*RunCrawlRequest) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{0}
}

func (x *RunCrawlRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *RunCrawlRequest) GetSeedId() string {
	if x != nil {
		return x.SeedId
	}
	return ""
}

type RunCrawlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobExecutionId string `protobuf:"bytes,1,opt,name=job_execution_id,json=jobExecutionId,proto3" json:"job_execution_id,omitempty"`
}

func (x *RunCrawlReply) Reset() {
	*x = RunCrawlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCrawlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCrawlReply) ProtoMessage() {}

func (x *RunCrawlReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCrawlReply.ProtoReflect.Descriptor instead.
func (*RunCrawlReply) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{1}
}

func (x *RunCrawlReply) GetJobExecutionId() string {
	if x != nil {
		return x.JobExecutionId
	}
	return ""
}

type RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []v1.Role `protobuf:"varint,1,rep,packed,name=role,proto3,enum=veidemann.api.config.v1.Role" json:"role,omitempty"`
}

func (x *RoleList) Reset() {
	*x = RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleList) ProtoMessage() {}

func (x *RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleList.ProtoReflect.Descriptor instead.
func (*RoleList) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{2}
}

func (x *RoleList) GetRole() []v1.Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type OpenIdConnectIssuerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenIdConnectIssuer string `protobuf:"bytes,1,opt,name=open_id_connect_issuer,json=openIdConnectIssuer,proto3" json:"open_id_connect_issuer,omitempty"`
}

func (x *OpenIdConnectIssuerReply) Reset() {
	*x = OpenIdConnectIssuerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenIdConnectIssuerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenIdConnectIssuerReply) ProtoMessage() {}

func (x *OpenIdConnectIssuerReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenIdConnectIssuerReply.ProtoReflect.Descriptor instead.
func (*OpenIdConnectIssuerReply) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{3}
}

func (x *OpenIdConnectIssuerReply) GetOpenIdConnectIssuer() string {
	if x != nil {
		return x.OpenIdConnectIssuer
	}
	return ""
}

type CrawlerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunStatus RunStatus `protobuf:"varint,1,opt,name=runStatus,proto3,enum=veidemann.api.controller.v1.RunStatus" json:"runStatus,omitempty"`
}

func (x *CrawlerStatus) Reset() {
	*x = CrawlerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerStatus) ProtoMessage() {}

func (x *CrawlerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerStatus.ProtoReflect.Descriptor instead.
func (*CrawlerStatus) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{4}
}

func (x *CrawlerStatus) GetRunStatus() RunStatus {
	if x != nil {
		return x.RunStatus
	}
	return RunStatus_RUNNING
}

var File_controller_v1_controller_proto protoreflect.FileDescriptor

var file_controller_v1_controller_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x65, 0x64, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x3d, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4f,
	0x0a, 0x18, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x70, 0x65, 0x6e,
	0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22,
	0x55, 0x0a, 0x0d, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x44, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x39, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x32, 0xf3, 0x05, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x25, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x08, 0x52, 0x75,
	0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x12, 0x2c, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x72, 0x0a, 0x13, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x1a, 0x2f, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4a,
	0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x2d, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65,
	0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0c, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x6e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x78, 0x0a, 0x25, 0x6e, 0x6f, 0x2e, 0x6e, 0x62,
	0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_v1_controller_proto_rawDescOnce sync.Once
	file_controller_v1_controller_proto_rawDescData = file_controller_v1_controller_proto_rawDesc
)

func file_controller_v1_controller_proto_rawDescGZIP() []byte {
	file_controller_v1_controller_proto_rawDescOnce.Do(func() {
		file_controller_v1_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_v1_controller_proto_rawDescData)
	})
	return file_controller_v1_controller_proto_rawDescData
}

var file_controller_v1_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_v1_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_controller_v1_controller_proto_goTypes = []interface{}{
	(RunStatus)(0),                   // 0: veidemann.api.controller.v1.RunStatus
	(*RunCrawlRequest)(nil),          // 1: veidemann.api.controller.v1.RunCrawlRequest
	(*RunCrawlReply)(nil),            // 2: veidemann.api.controller.v1.RunCrawlReply
	(*RoleList)(nil),                 // 3: veidemann.api.controller.v1.RoleList
	(*OpenIdConnectIssuerReply)(nil), // 4: veidemann.api.controller.v1.OpenIdConnectIssuerReply
	(*CrawlerStatus)(nil),            // 5: veidemann.api.controller.v1.CrawlerStatus
	(v1.Role)(0),                     // 6: veidemann.api.config.v1.Role
	(*empty.Empty)(nil),              // 7: google.protobuf.Empty
	(*ExecutionId)(nil),              // 8: veidemann.api.controller.v1.ExecutionId
	(*v11.CrawlExecutionStatus)(nil), // 9: veidemann.api.frontier.v1.CrawlExecutionStatus
	(*v11.JobExecutionStatus)(nil),   // 10: veidemann.api.frontier.v1.JobExecutionStatus
}
var file_controller_v1_controller_proto_depIdxs = []int32{
	6,  // 0: veidemann.api.controller.v1.RoleList.role:type_name -> veidemann.api.config.v1.Role
	0,  // 1: veidemann.api.controller.v1.CrawlerStatus.runStatus:type_name -> veidemann.api.controller.v1.RunStatus
	7,  // 2: veidemann.api.controller.v1.Controller.GetRolesForActiveUser:input_type -> google.protobuf.Empty
	1,  // 3: veidemann.api.controller.v1.Controller.RunCrawl:input_type -> veidemann.api.controller.v1.RunCrawlRequest
	8,  // 4: veidemann.api.controller.v1.Controller.AbortCrawlExecution:input_type -> veidemann.api.controller.v1.ExecutionId
	8,  // 5: veidemann.api.controller.v1.Controller.AbortJobExecution:input_type -> veidemann.api.controller.v1.ExecutionId
	7,  // 6: veidemann.api.controller.v1.Controller.GetOpenIdConnectIssuer:input_type -> google.protobuf.Empty
	7,  // 7: veidemann.api.controller.v1.Controller.PauseCrawler:input_type -> google.protobuf.Empty
	7,  // 8: veidemann.api.controller.v1.Controller.UnPauseCrawler:input_type -> google.protobuf.Empty
	7,  // 9: veidemann.api.controller.v1.Controller.Status:input_type -> google.protobuf.Empty
	3,  // 10: veidemann.api.controller.v1.Controller.GetRolesForActiveUser:output_type -> veidemann.api.controller.v1.RoleList
	2,  // 11: veidemann.api.controller.v1.Controller.RunCrawl:output_type -> veidemann.api.controller.v1.RunCrawlReply
	9,  // 12: veidemann.api.controller.v1.Controller.AbortCrawlExecution:output_type -> veidemann.api.frontier.v1.CrawlExecutionStatus
	10, // 13: veidemann.api.controller.v1.Controller.AbortJobExecution:output_type -> veidemann.api.frontier.v1.JobExecutionStatus
	4,  // 14: veidemann.api.controller.v1.Controller.GetOpenIdConnectIssuer:output_type -> veidemann.api.controller.v1.OpenIdConnectIssuerReply
	7,  // 15: veidemann.api.controller.v1.Controller.PauseCrawler:output_type -> google.protobuf.Empty
	7,  // 16: veidemann.api.controller.v1.Controller.UnPauseCrawler:output_type -> google.protobuf.Empty
	5,  // 17: veidemann.api.controller.v1.Controller.Status:output_type -> veidemann.api.controller.v1.CrawlerStatus
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_controller_v1_controller_proto_init() }
func file_controller_v1_controller_proto_init() {
	if File_controller_v1_controller_proto != nil {
		return
	}
	file_controller_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_v1_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCrawlRequest); i {
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
		file_controller_v1_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCrawlReply); i {
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
		file_controller_v1_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleList); i {
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
		file_controller_v1_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenIdConnectIssuerReply); i {
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
		file_controller_v1_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlerStatus); i {
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
			RawDescriptor: file_controller_v1_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_v1_controller_proto_goTypes,
		DependencyIndexes: file_controller_v1_controller_proto_depIdxs,
		EnumInfos:         file_controller_v1_controller_proto_enumTypes,
		MessageInfos:      file_controller_v1_controller_proto_msgTypes,
	}.Build()
	File_controller_v1_controller_proto = out.File
	file_controller_v1_controller_proto_rawDesc = nil
	file_controller_v1_controller_proto_goTypes = nil
	file_controller_v1_controller_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerClient interface {
	GetRolesForActiveUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RoleList, error)
	RunCrawl(ctx context.Context, in *RunCrawlRequest, opts ...grpc.CallOption) (*RunCrawlReply, error)
	// Abort a crawl execution
	AbortCrawlExecution(ctx context.Context, in *ExecutionId, opts ...grpc.CallOption) (*v11.CrawlExecutionStatus, error)
	// Abort a job execution
	AbortJobExecution(ctx context.Context, in *ExecutionId, opts ...grpc.CallOption) (*v11.JobExecutionStatus, error)
	// Get the configured OpenID connect issuer address
	GetOpenIdConnectIssuer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*OpenIdConnectIssuerReply, error)
	PauseCrawler(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	UnPauseCrawler(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CrawlerStatus, error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) GetRolesForActiveUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RoleList, error) {
	out := new(RoleList)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/GetRolesForActiveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) RunCrawl(ctx context.Context, in *RunCrawlRequest, opts ...grpc.CallOption) (*RunCrawlReply, error) {
	out := new(RunCrawlReply)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/RunCrawl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) AbortCrawlExecution(ctx context.Context, in *ExecutionId, opts ...grpc.CallOption) (*v11.CrawlExecutionStatus, error) {
	out := new(v11.CrawlExecutionStatus)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/AbortCrawlExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) AbortJobExecution(ctx context.Context, in *ExecutionId, opts ...grpc.CallOption) (*v11.JobExecutionStatus, error) {
	out := new(v11.JobExecutionStatus)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/AbortJobExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) GetOpenIdConnectIssuer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*OpenIdConnectIssuerReply, error) {
	out := new(OpenIdConnectIssuerReply)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/GetOpenIdConnectIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) PauseCrawler(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/PauseCrawler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) UnPauseCrawler(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/UnPauseCrawler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CrawlerStatus, error) {
	out := new(CrawlerStatus)
	err := c.cc.Invoke(ctx, "/veidemann.api.controller.v1.Controller/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerServer is the server API for Controller service.
type ControllerServer interface {
	GetRolesForActiveUser(context.Context, *empty.Empty) (*RoleList, error)
	RunCrawl(context.Context, *RunCrawlRequest) (*RunCrawlReply, error)
	// Abort a crawl execution
	AbortCrawlExecution(context.Context, *ExecutionId) (*v11.CrawlExecutionStatus, error)
	// Abort a job execution
	AbortJobExecution(context.Context, *ExecutionId) (*v11.JobExecutionStatus, error)
	// Get the configured OpenID connect issuer address
	GetOpenIdConnectIssuer(context.Context, *empty.Empty) (*OpenIdConnectIssuerReply, error)
	PauseCrawler(context.Context, *empty.Empty) (*empty.Empty, error)
	UnPauseCrawler(context.Context, *empty.Empty) (*empty.Empty, error)
	Status(context.Context, *empty.Empty) (*CrawlerStatus, error)
}

// UnimplementedControllerServer can be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (*UnimplementedControllerServer) GetRolesForActiveUser(context.Context, *empty.Empty) (*RoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesForActiveUser not implemented")
}
func (*UnimplementedControllerServer) RunCrawl(context.Context, *RunCrawlRequest) (*RunCrawlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCrawl not implemented")
}
func (*UnimplementedControllerServer) AbortCrawlExecution(context.Context, *ExecutionId) (*v11.CrawlExecutionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortCrawlExecution not implemented")
}
func (*UnimplementedControllerServer) AbortJobExecution(context.Context, *ExecutionId) (*v11.JobExecutionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortJobExecution not implemented")
}
func (*UnimplementedControllerServer) GetOpenIdConnectIssuer(context.Context, *empty.Empty) (*OpenIdConnectIssuerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenIdConnectIssuer not implemented")
}
func (*UnimplementedControllerServer) PauseCrawler(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseCrawler not implemented")
}
func (*UnimplementedControllerServer) UnPauseCrawler(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnPauseCrawler not implemented")
}
func (*UnimplementedControllerServer) Status(context.Context, *empty.Empty) (*CrawlerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_GetRolesForActiveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).GetRolesForActiveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/GetRolesForActiveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).GetRolesForActiveUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_RunCrawl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunCrawlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).RunCrawl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/RunCrawl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).RunCrawl(ctx, req.(*RunCrawlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_AbortCrawlExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).AbortCrawlExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/AbortCrawlExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).AbortCrawlExecution(ctx, req.(*ExecutionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_AbortJobExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).AbortJobExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/AbortJobExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).AbortJobExecution(ctx, req.(*ExecutionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_GetOpenIdConnectIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).GetOpenIdConnectIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/GetOpenIdConnectIssuer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).GetOpenIdConnectIssuer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_PauseCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).PauseCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/PauseCrawler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).PauseCrawler(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_UnPauseCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).UnPauseCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/UnPauseCrawler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).UnPauseCrawler(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.controller.v1.Controller/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.controller.v1.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRolesForActiveUser",
			Handler:    _Controller_GetRolesForActiveUser_Handler,
		},
		{
			MethodName: "RunCrawl",
			Handler:    _Controller_RunCrawl_Handler,
		},
		{
			MethodName: "AbortCrawlExecution",
			Handler:    _Controller_AbortCrawlExecution_Handler,
		},
		{
			MethodName: "AbortJobExecution",
			Handler:    _Controller_AbortJobExecution_Handler,
		},
		{
			MethodName: "GetOpenIdConnectIssuer",
			Handler:    _Controller_GetOpenIdConnectIssuer_Handler,
		},
		{
			MethodName: "PauseCrawler",
			Handler:    _Controller_PauseCrawler_Handler,
		},
		{
			MethodName: "UnPauseCrawler",
			Handler:    _Controller_UnPauseCrawler_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Controller_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/v1/controller.proto",
}
