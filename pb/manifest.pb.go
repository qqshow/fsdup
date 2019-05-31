// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/manifest.proto

package pb

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

type Slice struct {
	Checksum             []byte   `protobuf:"bytes,1,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Length               int64    `protobuf:"varint,3,opt,name=Length,proto3" json:"Length,omitempty"`
	Kind                 int32    `protobuf:"varint,4,opt,name=Kind,proto3" json:"Kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slice) Reset()         { *m = Slice{} }
func (m *Slice) String() string { return proto.CompactTextString(m) }
func (*Slice) ProtoMessage()    {}
func (*Slice) Descriptor() ([]byte, []int) {
	return fileDescriptor_2539f5857e72bfb4, []int{0}
}

func (m *Slice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slice.Unmarshal(m, b)
}
func (m *Slice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slice.Marshal(b, m, deterministic)
}
func (m *Slice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slice.Merge(m, src)
}
func (m *Slice) XXX_Size() int {
	return xxx_messageInfo_Slice.Size(m)
}
func (m *Slice) XXX_DiscardUnknown() {
	xxx_messageInfo_Slice.DiscardUnknown(m)
}

var xxx_messageInfo_Slice proto.InternalMessageInfo

func (m *Slice) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *Slice) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Slice) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Slice) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type ManifestV1 struct {
	Size                 int64    `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Slices               []*Slice `protobuf:"bytes,2,rep,name=Slices,proto3" json:"Slices,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=Comment,proto3" json:"Comment,omitempty"`
	ChunkMaxSize         int64    `protobuf:"varint,4,opt,name=ChunkMaxSize,proto3" json:"ChunkMaxSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestV1) Reset()         { *m = ManifestV1{} }
func (m *ManifestV1) String() string { return proto.CompactTextString(m) }
func (*ManifestV1) ProtoMessage()    {}
func (*ManifestV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2539f5857e72bfb4, []int{1}
}

func (m *ManifestV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestV1.Unmarshal(m, b)
}
func (m *ManifestV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestV1.Marshal(b, m, deterministic)
}
func (m *ManifestV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestV1.Merge(m, src)
}
func (m *ManifestV1) XXX_Size() int {
	return xxx_messageInfo_ManifestV1.Size(m)
}
func (m *ManifestV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestV1.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestV1 proto.InternalMessageInfo

func (m *ManifestV1) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ManifestV1) GetSlices() []*Slice {
	if m != nil {
		return m.Slices
	}
	return nil
}

func (m *ManifestV1) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ManifestV1) GetChunkMaxSize() int64 {
	if m != nil {
		return m.ChunkMaxSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Slice)(nil), "pb.Slice")
	proto.RegisterType((*ManifestV1)(nil), "pb.ManifestV1")
}

func init() { proto.RegisterFile("pb/manifest.proto", fileDescriptor_2539f5857e72bfb4) }

var fileDescriptor_2539f5857e72bfb4 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xe5, 0x38, 0x0d, 0xf4, 0xd1, 0x85, 0x37, 0x20, 0x8b, 0xc9, 0x64, 0xf2, 0x14, 0x04,
	0x7c, 0x42, 0x46, 0xa8, 0x90, 0x5c, 0x89, 0xbd, 0x2e, 0x4e, 0x63, 0x15, 0x3b, 0x16, 0x76, 0x24,
	0xc4, 0xc6, 0x9f, 0xa3, 0xbc, 0x04, 0xa4, 0x6e, 0xf7, 0x1c, 0x5b, 0xf7, 0xda, 0x70, 0x1d, 0xcd,
	0xbd, 0xdf, 0x07, 0xd7, 0xd9, 0x94, 0x9b, 0xf8, 0x39, 0xe4, 0x01, 0x8b, 0x68, 0xea, 0x23, 0xac,
	0x76, 0x1f, 0xee, 0x60, 0xf1, 0x16, 0x2e, 0xdb, 0xde, 0x1e, 0x4e, 0x69, 0xf4, 0x82, 0x49, 0xa6,
	0x36, 0xfa, 0x9f, 0xf1, 0x06, 0xaa, 0xd7, 0xae, 0x4b, 0x36, 0x8b, 0x42, 0x32, 0xc5, 0xf5, 0x42,
	0x93, 0x7f, 0xb1, 0xe1, 0x98, 0x7b, 0xc1, 0x67, 0x3f, 0x13, 0x22, 0x94, 0xcf, 0x2e, 0xbc, 0x8b,
	0x52, 0x32, 0xb5, 0xd2, 0x94, 0xeb, 0x1f, 0x06, 0xb0, 0x5d, 0xf6, 0xdf, 0x1e, 0xa6, 0x2b, 0x3b,
	0xf7, 0x6d, 0x69, 0x8a, 0x6b, 0xca, 0x78, 0x07, 0x15, 0xbd, 0x25, 0x89, 0x42, 0x72, 0x75, 0xf5,
	0xb8, 0x6e, 0xa2, 0x69, 0xc8, 0xe8, 0xe5, 0x00, 0x05, 0x5c, 0xb4, 0x83, 0xf7, 0x36, 0x64, 0x9a,
	0x5c, 0xeb, 0x3f, 0xc4, 0x1a, 0x36, 0x6d, 0x3f, 0x86, 0xd3, 0x76, 0xff, 0x45, 0xc5, 0x25, 0x15,
	0x9f, 0x39, 0x53, 0xd1, 0xbf, 0x9f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x54, 0xb8, 0x23,
	0x0c, 0x01, 0x00, 0x00,
}