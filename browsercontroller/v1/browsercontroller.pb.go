// Code generated by protoc-gen-go. DO NOT EDIT.
// source: browsercontroller/v1/browsercontroller.proto

package browsercontroller

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/nlnwa/veidemann-api-go/commons/v1"
	v1 "github.com/nlnwa/veidemann-api-go/config/v1"
	v11 "github.com/nlnwa/veidemann-api-go/frontier/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type NotifyActivity_Activity int32

const (
	NotifyActivity_DATA_RECEIVED     NotifyActivity_Activity = 0
	NotifyActivity_ALL_DATA_RECEIVED NotifyActivity_Activity = 1
)

var NotifyActivity_Activity_name = map[int32]string{
	0: "DATA_RECEIVED",
	1: "ALL_DATA_RECEIVED",
}

var NotifyActivity_Activity_value = map[string]int32{
	"DATA_RECEIVED":     0,
	"ALL_DATA_RECEIVED": 1,
}

func (x NotifyActivity_Activity) String() string {
	return proto.EnumName(NotifyActivity_Activity_name, int32(x))
}

func (NotifyActivity_Activity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{1, 0}
}

type RegisterNew struct {
	ProxyId              int32         `protobuf:"varint,1,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	Uri                  string        `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	CrawlExecutionId     string        `protobuf:"bytes,3,opt,name=crawl_execution_id,json=crawlExecutionId,proto3" json:"crawl_execution_id,omitempty"`
	JobExecutionId       string        `protobuf:"bytes,4,opt,name=job_execution_id,json=jobExecutionId,proto3" json:"job_execution_id,omitempty"`
	CollectionRef        *v1.ConfigRef `protobuf:"bytes,5,opt,name=collection_ref,json=collectionRef,proto3" json:"collection_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RegisterNew) Reset()         { *m = RegisterNew{} }
func (m *RegisterNew) String() string { return proto.CompactTextString(m) }
func (*RegisterNew) ProtoMessage()    {}
func (*RegisterNew) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{0}
}

func (m *RegisterNew) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNew.Unmarshal(m, b)
}
func (m *RegisterNew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNew.Marshal(b, m, deterministic)
}
func (m *RegisterNew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNew.Merge(m, src)
}
func (m *RegisterNew) XXX_Size() int {
	return xxx_messageInfo_RegisterNew.Size(m)
}
func (m *RegisterNew) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNew.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNew proto.InternalMessageInfo

func (m *RegisterNew) GetProxyId() int32 {
	if m != nil {
		return m.ProxyId
	}
	return 0
}

func (m *RegisterNew) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *RegisterNew) GetCrawlExecutionId() string {
	if m != nil {
		return m.CrawlExecutionId
	}
	return ""
}

func (m *RegisterNew) GetJobExecutionId() string {
	if m != nil {
		return m.JobExecutionId
	}
	return ""
}

func (m *RegisterNew) GetCollectionRef() *v1.ConfigRef {
	if m != nil {
		return m.CollectionRef
	}
	return nil
}

type NotifyActivity struct {
	Activity             NotifyActivity_Activity `protobuf:"varint,1,opt,name=activity,proto3,enum=veidemann.api.browsercontroller.v1.NotifyActivity_Activity" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NotifyActivity) Reset()         { *m = NotifyActivity{} }
func (m *NotifyActivity) String() string { return proto.CompactTextString(m) }
func (*NotifyActivity) ProtoMessage()    {}
func (*NotifyActivity) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{1}
}

func (m *NotifyActivity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyActivity.Unmarshal(m, b)
}
func (m *NotifyActivity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyActivity.Marshal(b, m, deterministic)
}
func (m *NotifyActivity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyActivity.Merge(m, src)
}
func (m *NotifyActivity) XXX_Size() int {
	return xxx_messageInfo_NotifyActivity.Size(m)
}
func (m *NotifyActivity) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyActivity.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyActivity proto.InternalMessageInfo

func (m *NotifyActivity) GetActivity() NotifyActivity_Activity {
	if m != nil {
		return m.Activity
	}
	return NotifyActivity_DATA_RECEIVED
}

type Completed struct {
	CrawlLog             *v11.CrawlLog `protobuf:"bytes,1,opt,name=crawl_log,json=crawlLog,proto3" json:"crawl_log,omitempty"`
	Cached               bool          `protobuf:"varint,2,opt,name=cached,proto3" json:"cached,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Completed) Reset()         { *m = Completed{} }
func (m *Completed) String() string { return proto.CompactTextString(m) }
func (*Completed) ProtoMessage()    {}
func (*Completed) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{2}
}

func (m *Completed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Completed.Unmarshal(m, b)
}
func (m *Completed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Completed.Marshal(b, m, deterministic)
}
func (m *Completed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Completed.Merge(m, src)
}
func (m *Completed) XXX_Size() int {
	return xxx_messageInfo_Completed.Size(m)
}
func (m *Completed) XXX_DiscardUnknown() {
	xxx_messageInfo_Completed.DiscardUnknown(m)
}

var xxx_messageInfo_Completed proto.InternalMessageInfo

func (m *Completed) GetCrawlLog() *v11.CrawlLog {
	if m != nil {
		return m.CrawlLog
	}
	return nil
}

func (m *Completed) GetCached() bool {
	if m != nil {
		return m.Cached
	}
	return false
}

type DoRequest struct {
	// Types that are valid to be assigned to Action:
	//	*DoRequest_New
	//	*DoRequest_Notify
	//	*DoRequest_Completed
	Action               isDoRequest_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DoRequest) Reset()         { *m = DoRequest{} }
func (m *DoRequest) String() string { return proto.CompactTextString(m) }
func (*DoRequest) ProtoMessage()    {}
func (*DoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{3}
}

func (m *DoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoRequest.Unmarshal(m, b)
}
func (m *DoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoRequest.Marshal(b, m, deterministic)
}
func (m *DoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoRequest.Merge(m, src)
}
func (m *DoRequest) XXX_Size() int {
	return xxx_messageInfo_DoRequest.Size(m)
}
func (m *DoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoRequest proto.InternalMessageInfo

type isDoRequest_Action interface {
	isDoRequest_Action()
}

type DoRequest_New struct {
	New *RegisterNew `protobuf:"bytes,1,opt,name=new,proto3,oneof"`
}

type DoRequest_Notify struct {
	Notify *NotifyActivity `protobuf:"bytes,2,opt,name=notify,proto3,oneof"`
}

type DoRequest_Completed struct {
	Completed *Completed `protobuf:"bytes,3,opt,name=completed,proto3,oneof"`
}

func (*DoRequest_New) isDoRequest_Action() {}

func (*DoRequest_Notify) isDoRequest_Action() {}

func (*DoRequest_Completed) isDoRequest_Action() {}

func (m *DoRequest) GetAction() isDoRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *DoRequest) GetNew() *RegisterNew {
	if x, ok := m.GetAction().(*DoRequest_New); ok {
		return x.New
	}
	return nil
}

func (m *DoRequest) GetNotify() *NotifyActivity {
	if x, ok := m.GetAction().(*DoRequest_Notify); ok {
		return x.Notify
	}
	return nil
}

func (m *DoRequest) GetCompleted() *Completed {
	if x, ok := m.GetAction().(*DoRequest_Completed); ok {
		return x.Completed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoRequest_New)(nil),
		(*DoRequest_Notify)(nil),
		(*DoRequest_Completed)(nil),
	}
}

type NewReply struct {
	CrawlExecutionId     string            `protobuf:"bytes,1,opt,name=crawl_execution_id,json=crawlExecutionId,proto3" json:"crawl_execution_id,omitempty"`
	JobExecutionId       string            `protobuf:"bytes,2,opt,name=job_execution_id,json=jobExecutionId,proto3" json:"job_execution_id,omitempty"`
	CollectionRef        *v1.ConfigRef     `protobuf:"bytes,4,opt,name=collection_ref,json=collectionRef,proto3" json:"collection_ref,omitempty"`
	ReplacementScript    *v1.BrowserScript `protobuf:"bytes,5,opt,name=replacement_script,json=replacementScript,proto3" json:"replacement_script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NewReply) Reset()         { *m = NewReply{} }
func (m *NewReply) String() string { return proto.CompactTextString(m) }
func (*NewReply) ProtoMessage()    {}
func (*NewReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{4}
}

func (m *NewReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewReply.Unmarshal(m, b)
}
func (m *NewReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewReply.Marshal(b, m, deterministic)
}
func (m *NewReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewReply.Merge(m, src)
}
func (m *NewReply) XXX_Size() int {
	return xxx_messageInfo_NewReply.Size(m)
}
func (m *NewReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NewReply.DiscardUnknown(m)
}

var xxx_messageInfo_NewReply proto.InternalMessageInfo

func (m *NewReply) GetCrawlExecutionId() string {
	if m != nil {
		return m.CrawlExecutionId
	}
	return ""
}

func (m *NewReply) GetJobExecutionId() string {
	if m != nil {
		return m.JobExecutionId
	}
	return ""
}

func (m *NewReply) GetCollectionRef() *v1.ConfigRef {
	if m != nil {
		return m.CollectionRef
	}
	return nil
}

func (m *NewReply) GetReplacementScript() *v1.BrowserScript {
	if m != nil {
		return m.ReplacementScript
	}
	return nil
}

type DoReply struct {
	// Types that are valid to be assigned to Action:
	//	*DoReply_New
	//	*DoReply_Cancel
	Action               isDoReply_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoReply) Reset()         { *m = DoReply{} }
func (m *DoReply) String() string { return proto.CompactTextString(m) }
func (*DoReply) ProtoMessage()    {}
func (*DoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b5373b641737e0, []int{5}
}

func (m *DoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReply.Unmarshal(m, b)
}
func (m *DoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReply.Marshal(b, m, deterministic)
}
func (m *DoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReply.Merge(m, src)
}
func (m *DoReply) XXX_Size() int {
	return xxx_messageInfo_DoReply.Size(m)
}
func (m *DoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReply.DiscardUnknown(m)
}

var xxx_messageInfo_DoReply proto.InternalMessageInfo

type isDoReply_Action interface {
	isDoReply_Action()
}

type DoReply_New struct {
	New *NewReply `protobuf:"bytes,1,opt,name=new,proto3,oneof"`
}

type DoReply_Cancel struct {
	Cancel string `protobuf:"bytes,4,opt,name=cancel,proto3,oneof"`
}

func (*DoReply_New) isDoReply_Action() {}

func (*DoReply_Cancel) isDoReply_Action() {}

func (m *DoReply) GetAction() isDoReply_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *DoReply) GetNew() *NewReply {
	if x, ok := m.GetAction().(*DoReply_New); ok {
		return x.New
	}
	return nil
}

func (m *DoReply) GetCancel() string {
	if x, ok := m.GetAction().(*DoReply_Cancel); ok {
		return x.Cancel
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoReply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoReply_New)(nil),
		(*DoReply_Cancel)(nil),
	}
}

func init() {
	proto.RegisterEnum("veidemann.api.browsercontroller.v1.NotifyActivity_Activity", NotifyActivity_Activity_name, NotifyActivity_Activity_value)
	proto.RegisterType((*RegisterNew)(nil), "veidemann.api.browsercontroller.v1.RegisterNew")
	proto.RegisterType((*NotifyActivity)(nil), "veidemann.api.browsercontroller.v1.NotifyActivity")
	proto.RegisterType((*Completed)(nil), "veidemann.api.browsercontroller.v1.Completed")
	proto.RegisterType((*DoRequest)(nil), "veidemann.api.browsercontroller.v1.DoRequest")
	proto.RegisterType((*NewReply)(nil), "veidemann.api.browsercontroller.v1.NewReply")
	proto.RegisterType((*DoReply)(nil), "veidemann.api.browsercontroller.v1.DoReply")
}

func init() {
	proto.RegisterFile("browsercontroller/v1/browsercontroller.proto", fileDescriptor_25b5373b641737e0)
}

var fileDescriptor_25b5373b641737e0 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0xc5, 0x24, 0x21, 0xb0, 0x28, 0x08, 0x56, 0xfa, 0xfd, 0x44, 0x68, 0x0f, 0x91, 0x2b, 0x55,
	0x48, 0x05, 0xbb, 0xa1, 0xbd, 0xe5, 0x12, 0x20, 0x48, 0x20, 0xd1, 0xa8, 0xda, 0xf4, 0x8f, 0xd4,
	0x0b, 0x32, 0xcb, 0xe0, 0x6c, 0x64, 0x76, 0xdc, 0xb5, 0x81, 0xa0, 0x5e, 0xfb, 0x11, 0x7a, 0xef,
	0x27, 0xeb, 0xe7, 0xe8, 0xb5, 0xf2, 0xda, 0x10, 0x08, 0x89, 0x4a, 0x7a, 0x62, 0x76, 0xe7, 0xcd,
	0xdb, 0xd9, 0x37, 0x6f, 0x31, 0xa9, 0x0d, 0x15, 0xce, 0x03, 0x50, 0x1c, 0x65, 0xa8, 0xd0, 0xf3,
	0x40, 0xd9, 0xb3, 0x53, 0x7b, 0x6b, 0xd3, 0xf2, 0x15, 0x86, 0x48, 0xcd, 0x19, 0x88, 0x11, 0x4c,
	0x1c, 0x29, 0x2d, 0xc7, 0x17, 0xd6, 0x36, 0x6c, 0x76, 0x5a, 0xa9, 0x70, 0x9c, 0x4c, 0x50, 0x06,
	0x11, 0x8f, 0x82, 0x00, 0xa7, 0x8a, 0x43, 0x10, 0xd7, 0x57, 0x8e, 0x39, 0xca, 0xb1, 0x70, 0x1f,
	0x4a, 0x3d, 0x1b, 0x2b, 0x94, 0xa1, 0x88, 0xcf, 0xbf, 0x9f, 0x7c, 0xee, 0x22, 0xba, 0x1e, 0xd8,
	0x8e, 0x2f, 0x6c, 0x47, 0x4a, 0x0c, 0x9d, 0x50, 0xa0, 0x5c, 0x66, 0x6b, 0xfa, 0x87, 0xd7, 0x5d,
	0x90, 0xf5, 0x60, 0xee, 0xb8, 0x2e, 0x28, 0x1b, 0x7d, 0x8d, 0xd8, 0x46, 0x9b, 0xbf, 0x0c, 0x92,
	0x67, 0xe0, 0x8a, 0x20, 0x04, 0x75, 0x09, 0x73, 0x7a, 0x4c, 0xb2, 0xbe, 0xc2, 0xdb, 0xc5, 0x40,
	0x8c, 0xca, 0xc6, 0x89, 0x51, 0x3d, 0x60, 0x87, 0x7a, 0xdd, 0x1b, 0xd1, 0x22, 0xd9, 0x9b, 0x2a,
	0x51, 0x4e, 0x9f, 0x18, 0xd5, 0x1c, 0x8b, 0x42, 0x5a, 0x23, 0x94, 0x2b, 0x67, 0xee, 0x0d, 0xe0,
	0x16, 0xf8, 0x34, 0xa2, 0x8d, 0xca, 0xf6, 0x34, 0xa0, 0xa8, 0x33, 0x9d, 0x65, 0xa2, 0x37, 0xa2,
	0x55, 0x52, 0xbc, 0xc1, 0xe1, 0x26, 0x76, 0x5f, 0x63, 0x0b, 0x37, 0x38, 0x5c, 0x47, 0xf6, 0x48,
	0x81, 0x47, 0x0a, 0x72, 0x0d, 0x53, 0x30, 0x2e, 0x1f, 0x9c, 0x18, 0xd5, 0x7c, 0xc3, 0xb4, 0x36,
	0x15, 0x8f, 0xf5, 0xb3, 0x66, 0xa7, 0x56, 0x5b, 0x47, 0x0c, 0xc6, 0xec, 0xe8, 0xae, 0x92, 0xc1,
	0xd8, 0xfc, 0x69, 0x90, 0xc2, 0x25, 0x86, 0x62, 0xbc, 0x68, 0xf2, 0x50, 0xcc, 0x44, 0xb8, 0xa0,
	0x9f, 0x49, 0xd6, 0x49, 0x62, 0x7d, 0xc5, 0x42, 0xe3, 0xcc, 0xfa, 0xfb, 0x24, 0xad, 0x4d, 0x16,
	0x6b, 0x19, 0xb0, 0x15, 0x99, 0xf9, 0x96, 0x64, 0x57, 0x87, 0x94, 0xc8, 0xd1, 0x45, 0xf3, 0x43,
	0x73, 0xc0, 0x3a, 0xed, 0x4e, 0xef, 0x53, 0xe7, 0xa2, 0x98, 0xa2, 0xff, 0x91, 0x52, 0xb3, 0xdf,
	0x1f, 0x6c, 0x6e, 0x1b, 0x26, 0x90, 0x5c, 0x1b, 0x27, 0xbe, 0x07, 0x21, 0x8c, 0xe8, 0x39, 0xc9,
	0xc5, 0x8a, 0x7a, 0xe8, 0xea, 0xe6, 0xf2, 0x8d, 0x17, 0xf7, 0x9a, 0x5b, 0x3a, 0x43, 0x5f, 0x3b,
	0xc2, 0xf6, 0xd1, 0x65, 0x59, 0x9e, 0x44, 0xf4, 0x7f, 0x92, 0xe1, 0x0e, 0xbf, 0x86, 0x91, 0x1e,
	0x54, 0x96, 0x25, 0x2b, 0xf3, 0xb7, 0x41, 0x72, 0x17, 0xc8, 0xe0, 0xeb, 0x14, 0x82, 0x90, 0xb6,
	0xc9, 0x9e, 0x84, 0x79, 0x72, 0x82, 0xbd, 0xcb, 0xf5, 0xd7, 0x4c, 0xd2, 0x4d, 0xb1, 0xa8, 0x9a,
	0xf6, 0x49, 0x46, 0x6a, 0x51, 0xf4, 0x51, 0xf9, 0x46, 0xe3, 0xe9, 0x32, 0x76, 0x53, 0x2c, 0xe1,
	0xa0, 0xef, 0x48, 0x8e, 0x2f, 0x75, 0xd0, 0x1e, 0xca, 0x37, 0xea, 0xbb, 0x10, 0xae, 0xc4, 0xeb,
	0xa6, 0xd8, 0x1d, 0x43, 0x2b, 0x4b, 0x32, 0x8e, 0x76, 0x81, 0xf9, 0x3d, 0x4d, 0xb2, 0x97, 0x30,
	0x67, 0xe0, 0x7b, 0x8b, 0x47, 0x2c, 0x6b, 0x3c, 0xc1, 0xb2, 0xe9, 0x1d, 0x2d, 0xbb, 0xff, 0x8f,
	0x96, 0xa5, 0x1f, 0x09, 0x55, 0xe0, 0x7b, 0x0e, 0x87, 0x09, 0xc8, 0x70, 0x10, 0x70, 0x25, 0xfc,
	0x30, 0x79, 0x01, 0x2f, 0x1f, 0xa5, 0x6b, 0xc5, 0xda, 0x5c, 0x69, 0x34, 0x2b, 0xad, 0x31, 0xc4,
	0x5b, 0x26, 0x92, 0xc3, 0x68, 0xfe, 0x91, 0x08, 0xe7, 0xeb, 0xd3, 0xaf, 0xed, 0x34, 0xb5, 0x44,
	0xbf, 0xe5, 0xe8, 0xcb, 0x91, 0xcb, 0x24, 0x07, 0x2f, 0x7e, 0xc1, 0xd1, 0x18, 0xe3, 0xf5, 0x9d,
	0xee, 0x8d, 0x6f, 0xa4, 0x94, 0x34, 0xd5, 0x5e, 0x71, 0xd1, 0x31, 0x49, 0x8f, 0x90, 0xee, 0x34,
	0xd8, 0x95, 0x5b, 0x2b, 0xaf, 0x76, 0x85, 0xfb, 0xde, 0xc2, 0x4c, 0x55, 0x8d, 0xd7, 0x46, 0xeb,
	0x87, 0x41, 0x6a, 0x12, 0x2d, 0x39, 0xb4, 0xa4, 0x74, 0x76, 0x28, 0x6f, 0x95, 0xb7, 0x7a, 0xbd,
	0x02, 0x35, 0x13, 0x1c, 0xde, 0x1b, 0x5f, 0xba, 0xae, 0x08, 0xaf, 0xa7, 0x43, 0x8b, 0xe3, 0xc4,
	0x96, 0x9e, 0x9c, 0x3b, 0xf6, 0x8a, 0xb0, 0xee, 0xf8, 0xa2, 0xee, 0xa2, 0xfd, 0xd0, 0x87, 0xe3,
	0x6c, 0x6b, 0x73, 0x98, 0xd1, 0xff, 0xba, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x0b,
	0xf5, 0xe5, 0x69, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrowserControllerClient is the client API for BrowserController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrowserControllerClient interface {
	Do(ctx context.Context, opts ...grpc.CallOption) (BrowserController_DoClient, error)
}

type browserControllerClient struct {
	cc *grpc.ClientConn
}

func NewBrowserControllerClient(cc *grpc.ClientConn) BrowserControllerClient {
	return &browserControllerClient{cc}
}

func (c *browserControllerClient) Do(ctx context.Context, opts ...grpc.CallOption) (BrowserController_DoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrowserController_serviceDesc.Streams[0], "/veidemann.api.browsercontroller.v1.BrowserController/do", opts...)
	if err != nil {
		return nil, err
	}
	x := &browserControllerDoClient{stream}
	return x, nil
}

type BrowserController_DoClient interface {
	Send(*DoRequest) error
	Recv() (*DoReply, error)
	grpc.ClientStream
}

type browserControllerDoClient struct {
	grpc.ClientStream
}

func (x *browserControllerDoClient) Send(m *DoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *browserControllerDoClient) Recv() (*DoReply, error) {
	m := new(DoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrowserControllerServer is the server API for BrowserController service.
type BrowserControllerServer interface {
	Do(BrowserController_DoServer) error
}

func RegisterBrowserControllerServer(s *grpc.Server, srv BrowserControllerServer) {
	s.RegisterService(&_BrowserController_serviceDesc, srv)
}

func _BrowserController_Do_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrowserControllerServer).Do(&browserControllerDoServer{stream})
}

type BrowserController_DoServer interface {
	Send(*DoReply) error
	Recv() (*DoRequest, error)
	grpc.ServerStream
}

type browserControllerDoServer struct {
	grpc.ServerStream
}

func (x *browserControllerDoServer) Send(m *DoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *browserControllerDoServer) Recv() (*DoRequest, error) {
	m := new(DoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BrowserController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.browsercontroller.v1.BrowserController",
	HandlerType: (*BrowserControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "do",
			Handler:       _BrowserController_Do_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "browsercontroller/v1/browsercontroller.proto",
}
