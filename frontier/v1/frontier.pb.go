// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: frontier/v1/frontier.proto

package frontier

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	v11 "github.com/nlnwa/veidemann-api-go/commons/v1"
	v1 "github.com/nlnwa/veidemann-api-go/config/v1"
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

type CrawlSeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobExecutionId string           `protobuf:"bytes,1,opt,name=job_execution_id,json=jobExecutionId,proto3" json:"job_execution_id,omitempty"`
	Job            *v1.ConfigObject `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	Seed           *v1.ConfigObject `protobuf:"bytes,6,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *CrawlSeedRequest) Reset() {
	*x = CrawlSeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontier_v1_frontier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlSeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlSeedRequest) ProtoMessage() {}

func (x *CrawlSeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frontier_v1_frontier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlSeedRequest.ProtoReflect.Descriptor instead.
func (*CrawlSeedRequest) Descriptor() ([]byte, []int) {
	return file_frontier_v1_frontier_proto_rawDescGZIP(), []int{0}
}

func (x *CrawlSeedRequest) GetJobExecutionId() string {
	if x != nil {
		return x.JobExecutionId
	}
	return ""
}

func (x *CrawlSeedRequest) GetJob() *v1.ConfigObject {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *CrawlSeedRequest) GetSeed() *v1.ConfigObject {
	if x != nil {
		return x.Seed
	}
	return nil
}

// The execution id for a seed crawl
type CrawlExecutionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CrawlExecutionId) Reset() {
	*x = CrawlExecutionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontier_v1_frontier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlExecutionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlExecutionId) ProtoMessage() {}

func (x *CrawlExecutionId) ProtoReflect() protoreflect.Message {
	mi := &file_frontier_v1_frontier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlExecutionId.ProtoReflect.Descriptor instead.
func (*CrawlExecutionId) Descriptor() ([]byte, []int) {
	return file_frontier_v1_frontier_proto_rawDescGZIP(), []int{1}
}

func (x *CrawlExecutionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Message sent from Harvester to request a new page to fetch and also used to return the harvest result.
// First message should set requestNextPage to true to tell frontier to respond with a page to fetch.
// When the fetch is done, a stream of PageHarvest objects are returned:
// The first object contains metrics.
// Subsequent objects contain outlinks until all outlinks are sent.
// Finally the client should complete the request.
type PageHarvest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*PageHarvest_RequestNextPage
	//	*PageHarvest_Metrics_
	//	*PageHarvest_Outlink
	//	*PageHarvest_Error
	Msg isPageHarvest_Msg `protobuf_oneof:"msg"`
}

func (x *PageHarvest) Reset() {
	*x = PageHarvest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontier_v1_frontier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageHarvest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageHarvest) ProtoMessage() {}

func (x *PageHarvest) ProtoReflect() protoreflect.Message {
	mi := &file_frontier_v1_frontier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageHarvest.ProtoReflect.Descriptor instead.
func (*PageHarvest) Descriptor() ([]byte, []int) {
	return file_frontier_v1_frontier_proto_rawDescGZIP(), []int{2}
}

func (m *PageHarvest) GetMsg() isPageHarvest_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *PageHarvest) GetRequestNextPage() bool {
	if x, ok := x.GetMsg().(*PageHarvest_RequestNextPage); ok {
		return x.RequestNextPage
	}
	return false
}

func (x *PageHarvest) GetMetrics() *PageHarvest_Metrics {
	if x, ok := x.GetMsg().(*PageHarvest_Metrics_); ok {
		return x.Metrics
	}
	return nil
}

func (x *PageHarvest) GetOutlink() *QueuedUri {
	if x, ok := x.GetMsg().(*PageHarvest_Outlink); ok {
		return x.Outlink
	}
	return nil
}

func (x *PageHarvest) GetError() *v11.Error {
	if x, ok := x.GetMsg().(*PageHarvest_Error); ok {
		return x.Error
	}
	return nil
}

type isPageHarvest_Msg interface {
	isPageHarvest_Msg()
}

type PageHarvest_RequestNextPage struct {
	// True if this is the initial request to start a new fetch
	RequestNextPage bool `protobuf:"varint,1,opt,name=requestNextPage,proto3,oneof"`
}

type PageHarvest_Metrics_ struct {
	// Collected metrics for the page fetched
	Metrics *PageHarvest_Metrics `protobuf:"bytes,2,opt,name=metrics,proto3,oneof"`
}

type PageHarvest_Outlink struct {
	// The outlinks found in the harvested page
	Outlink *QueuedUri `protobuf:"bytes,3,opt,name=outlink,proto3,oneof"`
}

type PageHarvest_Error struct {
	// If the overall page fetch failed. Should not be used for a singel uri failure
	Error *v11.Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*PageHarvest_RequestNextPage) isPageHarvest_Msg() {}

func (*PageHarvest_Metrics_) isPageHarvest_Msg() {}

func (*PageHarvest_Outlink) isPageHarvest_Msg() {}

func (*PageHarvest_Error) isPageHarvest_Msg() {}

// A specification of the page to fetch.
type PageHarvestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URI to fetch
	QueuedUri *QueuedUri `protobuf:"bytes,1,opt,name=queued_uri,json=queuedUri,proto3" json:"queued_uri,omitempty"`
	// The configuration for the fetch
	CrawlConfig *v1.ConfigObject `protobuf:"bytes,2,opt,name=crawl_config,json=crawlConfig,proto3" json:"crawl_config,omitempty"`
}

func (x *PageHarvestSpec) Reset() {
	*x = PageHarvestSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontier_v1_frontier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageHarvestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageHarvestSpec) ProtoMessage() {}

func (x *PageHarvestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_frontier_v1_frontier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageHarvestSpec.ProtoReflect.Descriptor instead.
func (*PageHarvestSpec) Descriptor() ([]byte, []int) {
	return file_frontier_v1_frontier_proto_rawDescGZIP(), []int{3}
}

func (x *PageHarvestSpec) GetQueuedUri() *QueuedUri {
	if x != nil {
		return x.QueuedUri
	}
	return nil
}

func (x *PageHarvestSpec) GetCrawlConfig() *v1.ConfigObject {
	if x != nil {
		return x.CrawlConfig
	}
	return nil
}

type PageHarvest_Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of uri's downloaded. The requested uri + embedded resources
	UriCount int32 `protobuf:"varint,1,opt,name=uri_count,json=uriCount,proto3" json:"uri_count,omitempty"`
	// Byte count for the resources downloaded. Includes embedded resources
	BytesDownloaded int64 `protobuf:"varint,2,opt,name=bytes_downloaded,json=bytesDownloaded,proto3" json:"bytes_downloaded,omitempty"`
}

func (x *PageHarvest_Metrics) Reset() {
	*x = PageHarvest_Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontier_v1_frontier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageHarvest_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageHarvest_Metrics) ProtoMessage() {}

func (x *PageHarvest_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_frontier_v1_frontier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageHarvest_Metrics.ProtoReflect.Descriptor instead.
func (*PageHarvest_Metrics) Descriptor() ([]byte, []int) {
	return file_frontier_v1_frontier_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PageHarvest_Metrics) GetUriCount() int32 {
	if x != nil {
		return x.UriCount
	}
	return 0
}

func (x *PageHarvest_Metrics) GetBytesDownloaded() int64 {
	if x != nil {
		return x.BytesDownloaded
	}
	return 0
}

var File_frontier_v1_frontier_proto protoreflect.FileDescriptor

var file_frontier_v1_frontier_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x10,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x22, 0x22,
	0x0a, 0x10, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xda, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x48, 0x61, 0x72, 0x76, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x4a,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x48, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x55, 0x72,
	0x69, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x37, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x51, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x72, 0x69, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x72, 0x69, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0xa0, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x48, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x43, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x55, 0x72, 0x69, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x64, 0x55, 0x72, 0x69, 0x12, 0x48, 0x0a, 0x0c, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x32, 0xdc, 0x01, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x12,
	0x67, 0x0a, 0x09, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x53, 0x65, 0x65, 0x64, 0x12, 0x2b, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x53, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x48, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x48, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x70, 0x0a, 0x23, 0x6e, 0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x69,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frontier_v1_frontier_proto_rawDescOnce sync.Once
	file_frontier_v1_frontier_proto_rawDescData = file_frontier_v1_frontier_proto_rawDesc
)

func file_frontier_v1_frontier_proto_rawDescGZIP() []byte {
	file_frontier_v1_frontier_proto_rawDescOnce.Do(func() {
		file_frontier_v1_frontier_proto_rawDescData = protoimpl.X.CompressGZIP(file_frontier_v1_frontier_proto_rawDescData)
	})
	return file_frontier_v1_frontier_proto_rawDescData
}

var file_frontier_v1_frontier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_frontier_v1_frontier_proto_goTypes = []interface{}{
	(*CrawlSeedRequest)(nil),    // 0: veidemann.api.frontier.v1.CrawlSeedRequest
	(*CrawlExecutionId)(nil),    // 1: veidemann.api.frontier.v1.CrawlExecutionId
	(*PageHarvest)(nil),         // 2: veidemann.api.frontier.v1.PageHarvest
	(*PageHarvestSpec)(nil),     // 3: veidemann.api.frontier.v1.PageHarvestSpec
	(*PageHarvest_Metrics)(nil), // 4: veidemann.api.frontier.v1.PageHarvest.Metrics
	(*v1.ConfigObject)(nil),     // 5: veidemann.api.config.v1.ConfigObject
	(*QueuedUri)(nil),           // 6: veidemann.api.frontier.v1.QueuedUri
	(*v11.Error)(nil),           // 7: veidemann.api.commons.v1.Error
}
var file_frontier_v1_frontier_proto_depIdxs = []int32{
	5, // 0: veidemann.api.frontier.v1.CrawlSeedRequest.job:type_name -> veidemann.api.config.v1.ConfigObject
	5, // 1: veidemann.api.frontier.v1.CrawlSeedRequest.seed:type_name -> veidemann.api.config.v1.ConfigObject
	4, // 2: veidemann.api.frontier.v1.PageHarvest.metrics:type_name -> veidemann.api.frontier.v1.PageHarvest.Metrics
	6, // 3: veidemann.api.frontier.v1.PageHarvest.outlink:type_name -> veidemann.api.frontier.v1.QueuedUri
	7, // 4: veidemann.api.frontier.v1.PageHarvest.error:type_name -> veidemann.api.commons.v1.Error
	6, // 5: veidemann.api.frontier.v1.PageHarvestSpec.queued_uri:type_name -> veidemann.api.frontier.v1.QueuedUri
	5, // 6: veidemann.api.frontier.v1.PageHarvestSpec.crawl_config:type_name -> veidemann.api.config.v1.ConfigObject
	0, // 7: veidemann.api.frontier.v1.Frontier.CrawlSeed:input_type -> veidemann.api.frontier.v1.CrawlSeedRequest
	2, // 8: veidemann.api.frontier.v1.Frontier.GetNextPage:input_type -> veidemann.api.frontier.v1.PageHarvest
	1, // 9: veidemann.api.frontier.v1.Frontier.CrawlSeed:output_type -> veidemann.api.frontier.v1.CrawlExecutionId
	3, // 10: veidemann.api.frontier.v1.Frontier.GetNextPage:output_type -> veidemann.api.frontier.v1.PageHarvestSpec
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_frontier_v1_frontier_proto_init() }
func file_frontier_v1_frontier_proto_init() {
	if File_frontier_v1_frontier_proto != nil {
		return
	}
	file_frontier_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_frontier_v1_frontier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlSeedRequest); i {
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
		file_frontier_v1_frontier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlExecutionId); i {
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
		file_frontier_v1_frontier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageHarvest); i {
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
		file_frontier_v1_frontier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageHarvestSpec); i {
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
		file_frontier_v1_frontier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageHarvest_Metrics); i {
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
	file_frontier_v1_frontier_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PageHarvest_RequestNextPage)(nil),
		(*PageHarvest_Metrics_)(nil),
		(*PageHarvest_Outlink)(nil),
		(*PageHarvest_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_frontier_v1_frontier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frontier_v1_frontier_proto_goTypes,
		DependencyIndexes: file_frontier_v1_frontier_proto_depIdxs,
		MessageInfos:      file_frontier_v1_frontier_proto_msgTypes,
	}.Build()
	File_frontier_v1_frontier_proto = out.File
	file_frontier_v1_frontier_proto_rawDesc = nil
	file_frontier_v1_frontier_proto_goTypes = nil
	file_frontier_v1_frontier_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FrontierClient is the client API for Frontier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontierClient interface {
	CrawlSeed(ctx context.Context, in *CrawlSeedRequest, opts ...grpc.CallOption) (*CrawlExecutionId, error)
	// Request a URI from the Frontiers queue
	// Used by a Harvester to fetch a new page
	GetNextPage(ctx context.Context, opts ...grpc.CallOption) (Frontier_GetNextPageClient, error)
}

type frontierClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontierClient(cc grpc.ClientConnInterface) FrontierClient {
	return &frontierClient{cc}
}

func (c *frontierClient) CrawlSeed(ctx context.Context, in *CrawlSeedRequest, opts ...grpc.CallOption) (*CrawlExecutionId, error) {
	out := new(CrawlExecutionId)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/CrawlSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontierClient) GetNextPage(ctx context.Context, opts ...grpc.CallOption) (Frontier_GetNextPageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontier_serviceDesc.Streams[0], "/veidemann.api.frontier.v1.Frontier/GetNextPage", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontierGetNextPageClient{stream}
	return x, nil
}

type Frontier_GetNextPageClient interface {
	Send(*PageHarvest) error
	Recv() (*PageHarvestSpec, error)
	grpc.ClientStream
}

type frontierGetNextPageClient struct {
	grpc.ClientStream
}

func (x *frontierGetNextPageClient) Send(m *PageHarvest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontierGetNextPageClient) Recv() (*PageHarvestSpec, error) {
	m := new(PageHarvestSpec)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontierServer is the server API for Frontier service.
type FrontierServer interface {
	CrawlSeed(context.Context, *CrawlSeedRequest) (*CrawlExecutionId, error)
	// Request a URI from the Frontiers queue
	// Used by a Harvester to fetch a new page
	GetNextPage(Frontier_GetNextPageServer) error
}

// UnimplementedFrontierServer can be embedded to have forward compatible implementations.
type UnimplementedFrontierServer struct {
}

func (*UnimplementedFrontierServer) CrawlSeed(context.Context, *CrawlSeedRequest) (*CrawlExecutionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrawlSeed not implemented")
}
func (*UnimplementedFrontierServer) GetNextPage(Frontier_GetNextPageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNextPage not implemented")
}

func RegisterFrontierServer(s *grpc.Server, srv FrontierServer) {
	s.RegisterService(&_Frontier_serviceDesc, srv)
}

func _Frontier_CrawlSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).CrawlSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/CrawlSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).CrawlSeed(ctx, req.(*CrawlSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontier_GetNextPage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontierServer).GetNextPage(&frontierGetNextPageServer{stream})
}

type Frontier_GetNextPageServer interface {
	Send(*PageHarvestSpec) error
	Recv() (*PageHarvest, error)
	grpc.ServerStream
}

type frontierGetNextPageServer struct {
	grpc.ServerStream
}

func (x *frontierGetNextPageServer) Send(m *PageHarvestSpec) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontierGetNextPageServer) Recv() (*PageHarvest, error) {
	m := new(PageHarvest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Frontier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.frontier.v1.Frontier",
	HandlerType: (*FrontierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrawlSeed",
			Handler:    _Frontier_CrawlSeed_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNextPage",
			Handler:       _Frontier_GetNextPage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "frontier/v1/frontier.proto",
}
