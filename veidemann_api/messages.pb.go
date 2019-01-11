// Code generated by protoc-gen-go. DO NOT EDIT.
// source: veidemann_api/messages.proto

package veidemann_api

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

type ExtractedText struct {
	WarcId               string   `protobuf:"bytes,1,opt,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	SentenceCount        int64    `protobuf:"varint,3,opt,name=sentence_count,json=sentenceCount,proto3" json:"sentence_count,omitempty"`
	WordCount            int64    `protobuf:"varint,4,opt,name=word_count,json=wordCount,proto3" json:"word_count,omitempty"`
	LongWordCount        int64    `protobuf:"varint,5,opt,name=long_word_count,json=longWordCount,proto3" json:"long_word_count,omitempty"`
	CharacterCount       int64    `protobuf:"varint,6,opt,name=character_count,json=characterCount,proto3" json:"character_count,omitempty"`
	Lix                  int64    `protobuf:"varint,7,opt,name=lix,proto3" json:"lix,omitempty"`
	Language             string   `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractedText) Reset()         { *m = ExtractedText{} }
func (m *ExtractedText) String() string { return proto.CompactTextString(m) }
func (*ExtractedText) ProtoMessage()    {}
func (*ExtractedText) Descriptor() ([]byte, []int) {
	return fileDescriptor_a749149430e4086f, []int{0}
}

func (m *ExtractedText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractedText.Unmarshal(m, b)
}
func (m *ExtractedText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractedText.Marshal(b, m, deterministic)
}
func (m *ExtractedText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractedText.Merge(m, src)
}
func (m *ExtractedText) XXX_Size() int {
	return xxx_messageInfo_ExtractedText.Size(m)
}
func (m *ExtractedText) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractedText.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractedText proto.InternalMessageInfo

func (m *ExtractedText) GetWarcId() string {
	if m != nil {
		return m.WarcId
	}
	return ""
}

func (m *ExtractedText) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ExtractedText) GetSentenceCount() int64 {
	if m != nil {
		return m.SentenceCount
	}
	return 0
}

func (m *ExtractedText) GetWordCount() int64 {
	if m != nil {
		return m.WordCount
	}
	return 0
}

func (m *ExtractedText) GetLongWordCount() int64 {
	if m != nil {
		return m.LongWordCount
	}
	return 0
}

func (m *ExtractedText) GetCharacterCount() int64 {
	if m != nil {
		return m.CharacterCount
	}
	return 0
}

func (m *ExtractedText) GetLix() int64 {
	if m != nil {
		return m.Lix
	}
	return 0
}

func (m *ExtractedText) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type Screenshot struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExecutionId          string   `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	Uri                  string   `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Img                  []byte   `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screenshot) Reset()         { *m = Screenshot{} }
func (m *Screenshot) String() string { return proto.CompactTextString(m) }
func (*Screenshot) ProtoMessage()    {}
func (*Screenshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_a749149430e4086f, []int{1}
}

func (m *Screenshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screenshot.Unmarshal(m, b)
}
func (m *Screenshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screenshot.Marshal(b, m, deterministic)
}
func (m *Screenshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screenshot.Merge(m, src)
}
func (m *Screenshot) XXX_Size() int {
	return xxx_messageInfo_Screenshot.Size(m)
}
func (m *Screenshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Screenshot.DiscardUnknown(m)
}

var xxx_messageInfo_Screenshot proto.InternalMessageInfo

func (m *Screenshot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Screenshot) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *Screenshot) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Screenshot) GetImg() []byte {
	if m != nil {
		return m.Img
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtractedText)(nil), "veidemann.api.ExtractedText")
	proto.RegisterType((*Screenshot)(nil), "veidemann.api.Screenshot")
}

func init() { proto.RegisterFile("veidemann_api/messages.proto", fileDescriptor_a749149430e4086f) }

var fileDescriptor_a749149430e4086f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x69, 0xda, 0xd7, 0x1f, 0xf3, 0x9a, 0xf6, 0xb1, 0x97, 0x86, 0x87, 0x42, 0x2d, 0xa8,
	0xbd, 0x34, 0x39, 0x78, 0x14, 0x2f, 0x15, 0x0f, 0x3d, 0x08, 0x12, 0x05, 0xc1, 0x4b, 0xd8, 0x6e,
	0x86, 0xed, 0x42, 0x32, 0x1b, 0x92, 0x8d, 0xcd, 0xd5, 0xff, 0x5c, 0x76, 0x53, 0x83, 0xbd, 0xcd,
	0x7c, 0xe6, 0x43, 0x92, 0x6f, 0xbe, 0x70, 0xf1, 0x89, 0x2a, 0xc5, 0x9c, 0x13, 0x25, 0xbc, 0x50,
	0x51, 0x8e, 0x55, 0xc5, 0x25, 0x56, 0x61, 0x51, 0x6a, 0xa3, 0x99, 0xdf, 0x5d, 0x43, 0x5e, 0xa8,
	0xd5, 0x97, 0x07, 0xfe, 0x53, 0x63, 0x4a, 0x2e, 0x0c, 0xa6, 0x6f, 0xd8, 0x18, 0xb6, 0x80, 0xd1,
	0x91, 0x97, 0x22, 0x51, 0x69, 0xd0, 0x5b, 0xf6, 0xd6, 0x93, 0x78, 0x68, 0xd7, 0x5d, 0xca, 0x18,
	0x0c, 0x0c, 0x36, 0x26, 0xf0, 0x1c, 0x75, 0x33, 0xbb, 0x86, 0x59, 0x85, 0x64, 0x90, 0x04, 0x26,
	0x42, 0xd7, 0x64, 0x82, 0xfe, 0xb2, 0xb7, 0xee, 0xc7, 0xfe, 0x0f, 0x7d, 0xb4, 0x90, 0x5d, 0x02,
	0x1c, 0x75, 0x99, 0x9e, 0x94, 0x81, 0x53, 0x26, 0x96, 0xb4, 0xe7, 0x1b, 0x98, 0x67, 0x9a, 0x64,
	0xf2, 0xcb, 0xf9, 0xd3, 0x3e, 0xc6, 0xe2, 0xf7, 0xce, 0xbb, 0x85, 0xb9, 0x38, 0x70, 0xf7, 0xad,
	0xe5, 0xc9, 0x1b, 0x3a, 0x6f, 0xd6, 0xe1, 0x56, 0xfc, 0x07, 0xfd, 0x4c, 0x35, 0xc1, 0xc8, 0x1d,
	0xed, 0xc8, 0xfe, 0xc3, 0x38, 0xe3, 0x24, 0x6b, 0x2e, 0x31, 0x18, 0xbb, 0x00, 0xdd, 0xbe, 0x12,
	0x00, 0xaf, 0xa2, 0x44, 0xa4, 0xea, 0xa0, 0x0d, 0x9b, 0x81, 0xd7, 0x45, 0xf7, 0x54, 0xca, 0xae,
	0x60, 0x8a, 0x0d, 0x8a, 0xda, 0x28, 0x4d, 0xf6, 0xa7, 0xb4, 0xf1, 0xff, 0x76, 0x6c, 0x97, 0xda,
	0xd7, 0xd5, 0xa5, 0x72, 0xd1, 0x27, 0xb1, 0x1d, 0x2d, 0x51, 0xb9, 0x74, 0x49, 0xa7, 0xb1, 0x1d,
	0xb7, 0x12, 0x16, 0xa4, 0x43, 0xda, 0x87, 0x44, 0x3c, 0x3c, 0xeb, 0x60, 0xeb, 0x3f, 0x9f, 0x2a,
	0x7a, 0xb1, 0x0d, 0x7d, 0x3c, 0x48, 0x65, 0x0e, 0xf5, 0x3e, 0x14, 0x3a, 0x8f, 0x28, 0xa3, 0x23,
	0x8f, 0x3a, 0x7d, 0xc3, 0x0b, 0xb5, 0x91, 0x3a, 0x3a, 0x6b, 0xf8, 0xfe, 0x6c, 0xdb, 0x0f, 0x5d,
	0xcf, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xc8, 0x1f, 0xbf, 0x07, 0x02, 0x00, 0x00,
}
