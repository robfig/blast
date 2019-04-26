// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/index/index.proto

package index

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	raft "github.com/mosuka/blast/protobuf/raft"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IndexCommand_Type int32

const (
	IndexCommand_UNKNOWN_COMMAND IndexCommand_Type = 0
	IndexCommand_SET_METADATA    IndexCommand_Type = 1
	IndexCommand_DELETE_METADATA IndexCommand_Type = 2
	IndexCommand_INDEX_DOCUMENT  IndexCommand_Type = 3
	IndexCommand_DELETE_DOCUMENT IndexCommand_Type = 4
)

var IndexCommand_Type_name = map[int32]string{
	0: "UNKNOWN_COMMAND",
	1: "SET_METADATA",
	2: "DELETE_METADATA",
	3: "INDEX_DOCUMENT",
	4: "DELETE_DOCUMENT",
}

var IndexCommand_Type_value = map[string]int32{
	"UNKNOWN_COMMAND": 0,
	"SET_METADATA":    1,
	"DELETE_METADATA": 2,
	"INDEX_DOCUMENT":  3,
	"DELETE_DOCUMENT": 4,
}

func (x IndexCommand_Type) String() string {
	return proto.EnumName(IndexCommand_Type_name, int32(x))
}

func (IndexCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{6, 0}
}

type Document struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *any.Any `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{0}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetFields() *any.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

type UpdateResult struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResult) Reset()         { *m = UpdateResult{} }
func (m *UpdateResult) String() string { return proto.CompactTextString(m) }
func (*UpdateResult) ProtoMessage()    {}
func (*UpdateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{1}
}

func (m *UpdateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResult.Unmarshal(m, b)
}
func (m *UpdateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResult.Marshal(b, m, deterministic)
}
func (m *UpdateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResult.Merge(m, src)
}
func (m *UpdateResult) XXX_Size() int {
	return xxx_messageInfo_UpdateResult.Size(m)
}
func (m *UpdateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResult proto.InternalMessageInfo

func (m *UpdateResult) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type IndexConfig struct {
	IndexMapping         *any.Any `protobuf:"bytes,1,opt,name=index_mapping,json=indexMapping,proto3" json:"index_mapping,omitempty"`
	IndexType            string   `protobuf:"bytes,2,opt,name=index_type,json=indexType,proto3" json:"index_type,omitempty"`
	IndexStorageType     string   `protobuf:"bytes,3,opt,name=index_storage_type,json=indexStorageType,proto3" json:"index_storage_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexConfig) Reset()         { *m = IndexConfig{} }
func (m *IndexConfig) String() string { return proto.CompactTextString(m) }
func (*IndexConfig) ProtoMessage()    {}
func (*IndexConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{2}
}

func (m *IndexConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexConfig.Unmarshal(m, b)
}
func (m *IndexConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexConfig.Marshal(b, m, deterministic)
}
func (m *IndexConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexConfig.Merge(m, src)
}
func (m *IndexConfig) XXX_Size() int {
	return xxx_messageInfo_IndexConfig.Size(m)
}
func (m *IndexConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IndexConfig proto.InternalMessageInfo

func (m *IndexConfig) GetIndexMapping() *any.Any {
	if m != nil {
		return m.IndexMapping
	}
	return nil
}

func (m *IndexConfig) GetIndexType() string {
	if m != nil {
		return m.IndexType
	}
	return ""
}

func (m *IndexConfig) GetIndexStorageType() string {
	if m != nil {
		return m.IndexStorageType
	}
	return ""
}

type IndexStats struct {
	Stats                *any.Any `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexStats) Reset()         { *m = IndexStats{} }
func (m *IndexStats) String() string { return proto.CompactTextString(m) }
func (*IndexStats) ProtoMessage()    {}
func (*IndexStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{3}
}

func (m *IndexStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexStats.Unmarshal(m, b)
}
func (m *IndexStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexStats.Marshal(b, m, deterministic)
}
func (m *IndexStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexStats.Merge(m, src)
}
func (m *IndexStats) XXX_Size() int {
	return xxx_messageInfo_IndexStats.Size(m)
}
func (m *IndexStats) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexStats.DiscardUnknown(m)
}

var xxx_messageInfo_IndexStats proto.InternalMessageInfo

func (m *IndexStats) GetStats() *any.Any {
	if m != nil {
		return m.Stats
	}
	return nil
}

type SearchRequest struct {
	SearchRequest        *any.Any `protobuf:"bytes,1,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{4}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetSearchRequest() *any.Any {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

type SearchResponse struct {
	SearchResult         *any.Any `protobuf:"bytes,1,opt,name=search_result,json=searchResult,proto3" json:"search_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{5}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetSearchResult() *any.Any {
	if m != nil {
		return m.SearchResult
	}
	return nil
}

type IndexCommand struct {
	Type                 IndexCommand_Type `protobuf:"varint,1,opt,name=type,proto3,enum=index.IndexCommand_Type" json:"type,omitempty"`
	Data                 *any.Any          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IndexCommand) Reset()         { *m = IndexCommand{} }
func (m *IndexCommand) String() string { return proto.CompactTextString(m) }
func (*IndexCommand) ProtoMessage()    {}
func (*IndexCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{6}
}

func (m *IndexCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexCommand.Unmarshal(m, b)
}
func (m *IndexCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexCommand.Marshal(b, m, deterministic)
}
func (m *IndexCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexCommand.Merge(m, src)
}
func (m *IndexCommand) XXX_Size() int {
	return xxx_messageInfo_IndexCommand.Size(m)
}
func (m *IndexCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexCommand.DiscardUnknown(m)
}

var xxx_messageInfo_IndexCommand proto.InternalMessageInfo

func (m *IndexCommand) GetType() IndexCommand_Type {
	if m != nil {
		return m.Type
	}
	return IndexCommand_UNKNOWN_COMMAND
}

func (m *IndexCommand) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("index.IndexCommand_Type", IndexCommand_Type_name, IndexCommand_Type_value)
	proto.RegisterType((*Document)(nil), "index.Document")
	proto.RegisterType((*UpdateResult)(nil), "index.UpdateResult")
	proto.RegisterType((*IndexConfig)(nil), "index.IndexConfig")
	proto.RegisterType((*IndexStats)(nil), "index.IndexStats")
	proto.RegisterType((*SearchRequest)(nil), "index.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "index.SearchResponse")
	proto.RegisterType((*IndexCommand)(nil), "index.IndexCommand")
}

func init() { proto.RegisterFile("protobuf/index/index.proto", fileDescriptor_7b2daf652facb3ae) }

var fileDescriptor_7b2daf652facb3ae = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xe1, 0x6e, 0xda, 0x3a,
	0x14, 0xc7, 0x03, 0x05, 0x6e, 0x7b, 0x0a, 0x94, 0xeb, 0xf6, 0x5e, 0x71, 0x73, 0x35, 0xa9, 0xb2,
	0xa6, 0x89, 0x4d, 0x5d, 0xd0, 0x98, 0xa6, 0x6d, 0xda, 0x34, 0x89, 0x91, 0x88, 0x75, 0x2d, 0xa9,
	0x14, 0xa8, 0x36, 0xed, 0x0b, 0x32, 0xc4, 0xd0, 0xa8, 0x24, 0xce, 0xb0, 0x33, 0x8d, 0x07, 0xd9,
	0x93, 0xf5, 0x85, 0x26, 0xdb, 0x84, 0xa5, 0xad, 0x82, 0xb6, 0x2f, 0x51, 0xfc, 0x3f, 0xbf, 0x73,
	0x62, 0xff, 0x7d, 0x4e, 0xc0, 0x8c, 0x97, 0x4c, 0xb0, 0x49, 0x32, 0x6b, 0x07, 0x91, 0x4f, 0xbf,
	0xeb, 0xa7, 0xa5, 0x44, 0x54, 0x56, 0x0b, 0xf3, 0xbf, 0x39, 0x63, 0xf3, 0x05, 0x6d, 0x6f, 0x48,
	0x12, 0xad, 0x34, 0x61, 0xfe, 0x7f, 0x37, 0x44, 0xc3, 0x58, 0xa4, 0xc1, 0xe6, 0x46, 0x5d, 0x92,
	0x99, 0x50, 0x0f, 0x1d, 0xc1, 0x1f, 0x60, 0xd7, 0x66, 0xd3, 0x24, 0xa4, 0x91, 0x40, 0x75, 0x28,
	0x06, 0x7e, 0xb3, 0x70, 0x5c, 0x68, 0xed, 0x79, 0xc5, 0xc0, 0x47, 0x27, 0x50, 0x99, 0x05, 0x74,
	0xe1, 0xf3, 0x66, 0xf1, 0xb8, 0xd0, 0xda, 0xef, 0x1c, 0x59, 0xfa, 0x1b, 0x56, 0x5a, 0xcd, 0xea,
	0x46, 0x2b, 0x6f, 0xcd, 0xe0, 0x87, 0x50, 0xbd, 0x8c, 0x7d, 0x22, 0xa8, 0x47, 0x79, 0xb2, 0x10,
	0xe8, 0x08, 0xca, 0x53, 0x96, 0x44, 0x42, 0x15, 0x2c, 0x7b, 0x7a, 0x81, 0x7f, 0x14, 0x60, 0xff,
	0x54, 0x9e, 0xa5, 0xc7, 0xa2, 0x59, 0x30, 0x47, 0xaf, 0xa1, 0xa6, 0x8e, 0x36, 0x0e, 0x49, 0x1c,
	0x07, 0xd1, 0x5c, 0xd1, 0x79, 0x9f, 0xaa, 0x2a, 0x74, 0xa0, 0x49, 0xf4, 0x00, 0x40, 0xa7, 0x8a,
	0x55, 0x4c, 0xd5, 0x16, 0xf7, 0xbc, 0x3d, 0xa5, 0x8c, 0x56, 0x31, 0x45, 0x27, 0x80, 0x74, 0x98,
	0x0b, 0xb6, 0x24, 0x73, 0xaa, 0xb1, 0x1d, 0x85, 0x35, 0x54, 0x64, 0xa8, 0x03, 0x92, 0xc6, 0xaf,
	0x00, 0x4e, 0xb5, 0x46, 0x04, 0x47, 0x4f, 0xa0, 0xcc, 0xe5, 0xcb, 0xd6, 0xdd, 0x68, 0x04, 0x9f,
	0x43, 0x6d, 0x48, 0xc9, 0x72, 0x7a, 0xe5, 0xd1, 0xaf, 0x09, 0xe5, 0x02, 0xbd, 0x81, 0x3a, 0x57,
	0xc2, 0x78, 0xa9, 0x95, 0xad, 0x55, 0x6a, 0x3c, 0x9b, 0x8c, 0xcf, 0xa0, 0x9e, 0x56, 0xe3, 0x31,
	0x8b, 0x38, 0x95, 0x0e, 0x6d, 0xca, 0x49, 0x63, 0xb7, 0x3b, 0x94, 0x56, 0x93, 0x24, 0xbe, 0x29,
	0x40, 0x75, 0x6d, 0x76, 0x18, 0x92, 0x48, 0xde, 0x68, 0x49, 0xb9, 0x20, 0x4b, 0xd4, 0x3b, 0x4d,
	0x4b, 0xb7, 0x58, 0x16, 0xb1, 0xa4, 0x1b, 0x9e, 0xa2, 0x50, 0x0b, 0x4a, 0x3e, 0x11, 0x64, 0xeb,
	0xed, 0x2b, 0x02, 0x5f, 0x43, 0x49, 0x79, 0x7e, 0x08, 0x07, 0x97, 0xee, 0x99, 0x7b, 0xf1, 0xc9,
	0x1d, 0xf7, 0x2e, 0x06, 0x83, 0xae, 0x6b, 0x37, 0x0c, 0xd4, 0x80, 0xea, 0xd0, 0x19, 0x8d, 0x07,
	0xce, 0xa8, 0x6b, 0x77, 0x47, 0xdd, 0x46, 0x41, 0x62, 0xb6, 0x73, 0xee, 0x8c, 0x9c, 0x5f, 0x62,
	0x11, 0x21, 0xa8, 0x9f, 0xba, 0xb6, 0xf3, 0x79, 0x6c, 0x5f, 0xf4, 0x2e, 0x07, 0x8e, 0x3b, 0x6a,
	0xec, 0x64, 0xc0, 0x8d, 0x58, 0xea, 0xdc, 0x94, 0xa0, 0xac, 0xb6, 0x2c, 0x8f, 0xf3, 0x91, 0x05,
	0x11, 0x02, 0x4b, 0x75, 0xb4, 0xcb, 0x7c, 0x6a, 0xfe, 0x7b, 0x6f, 0x9b, 0x8e, 0x1c, 0x04, 0x6c,
	0xa0, 0xa7, 0x50, 0x3e, 0xa7, 0xe4, 0x1b, 0xfd, 0x4d, 0xbc, 0x0d, 0x7f, 0xf5, 0xa9, 0x90, 0x10,
	0xca, 0x81, 0xcc, 0x4c, 0x21, 0x6c, 0xa0, 0x17, 0x00, 0x7d, 0x2a, 0x7a, 0x8b, 0x84, 0x0b, 0xba,
	0xcc, 0xcd, 0xa9, 0xe9, 0x9c, 0x35, 0x86, 0x0d, 0xf4, 0x16, 0x76, 0x87, 0x11, 0x89, 0xf9, 0x15,
	0x13, 0xb9, 0x49, 0xf9, 0xbb, 0x7c, 0x0c, 0x3b, 0x7d, 0x2a, 0xd0, 0xc1, 0xfa, 0x2a, 0xd3, 0x59,
	0x36, 0xef, 0x0a, 0xd8, 0x40, 0xcf, 0x52, 0xdb, 0xee, 0xc1, 0x87, 0x6b, 0x21, 0x3b, 0xbf, 0xd8,
	0x68, 0x15, 0x50, 0x07, 0x2a, 0x36, 0x5d, 0x50, 0x41, 0xff, 0x20, 0xe7, 0x25, 0x54, 0x74, 0x07,
	0xa3, 0xa3, 0x35, 0x72, 0x6b, 0x3c, 0xcc, 0x7f, 0xee, 0xa8, 0xba, 0xcd, 0xb1, 0x81, 0xde, 0x41,
	0xbd, 0x4f, 0x45, 0xf6, 0xe7, 0x90, 0x67, 0x07, 0xba, 0xdd, 0xb8, 0x92, 0x55, 0x46, 0xd6, 0xd2,
	0x7c, 0x3d, 0xc5, 0x79, 0xe9, 0x7f, 0x67, 0xd3, 0x15, 0x8a, 0x8d, 0xf7, 0xad, 0x2f, 0x8f, 0xe6,
	0x81, 0xb8, 0x4a, 0x26, 0xd6, 0x94, 0x85, 0xed, 0x90, 0xf1, 0xe4, 0x9a, 0xb4, 0x27, 0x0b, 0xc2,
	0x45, 0xfb, 0xf6, 0x7f, 0x79, 0x52, 0x51, 0xeb, 0xe7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0xaf, 0xe7, 0xc0, 0xb0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexClient interface {
	Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error)
	GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Document, error)
	Index(ctx context.Context, opts ...grpc.CallOption) (Index_IndexClient, error)
	Delete(ctx context.Context, opts ...grpc.CallOption) (Index_DeleteClient, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GetIndexConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IndexConfig, error)
	GetIndexStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IndexStats, error)
}

type indexClient struct {
	cc *grpc.ClientConn
}

func NewIndexClient(cc *grpc.ClientConn) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error) {
	out := new(raft.Node)
	err := c.cc.Invoke(ctx, "/index.Index/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error) {
	out := new(raft.Cluster)
	err := c.cc.Invoke(ctx, "/index.Index/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Get(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/index.Index/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Index(ctx context.Context, opts ...grpc.CallOption) (Index_IndexClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Index_serviceDesc.Streams[0], "/index.Index/Index", opts...)
	if err != nil {
		return nil, err
	}
	x := &indexIndexClient{stream}
	return x, nil
}

type Index_IndexClient interface {
	Send(*Document) error
	CloseAndRecv() (*UpdateResult, error)
	grpc.ClientStream
}

type indexIndexClient struct {
	grpc.ClientStream
}

func (x *indexIndexClient) Send(m *Document) error {
	return x.ClientStream.SendMsg(m)
}

func (x *indexIndexClient) CloseAndRecv() (*UpdateResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *indexClient) Delete(ctx context.Context, opts ...grpc.CallOption) (Index_DeleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Index_serviceDesc.Streams[1], "/index.Index/Delete", opts...)
	if err != nil {
		return nil, err
	}
	x := &indexDeleteClient{stream}
	return x, nil
}

type Index_DeleteClient interface {
	Send(*Document) error
	CloseAndRecv() (*UpdateResult, error)
	grpc.ClientStream
}

type indexDeleteClient struct {
	grpc.ClientStream
}

func (x *indexDeleteClient) Send(m *Document) error {
	return x.ClientStream.SendMsg(m)
}

func (x *indexDeleteClient) CloseAndRecv() (*UpdateResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *indexClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetIndexConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IndexConfig, error) {
	out := new(IndexConfig)
	err := c.cc.Invoke(ctx, "/index.Index/GetIndexConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetIndexStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IndexStats, error) {
	out := new(IndexStats)
	err := c.cc.Invoke(ctx, "/index.Index/GetIndexStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
type IndexServer interface {
	Join(context.Context, *raft.Node) (*empty.Empty, error)
	Leave(context.Context, *raft.Node) (*empty.Empty, error)
	GetNode(context.Context, *empty.Empty) (*raft.Node, error)
	GetCluster(context.Context, *empty.Empty) (*raft.Cluster, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *Document) (*Document, error)
	Index(Index_IndexServer) error
	Delete(Index_DeleteServer) error
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	GetIndexConfig(context.Context, *empty.Empty) (*IndexConfig, error)
	GetIndexStats(context.Context, *empty.Empty) (*IndexStats, error)
}

func RegisterIndexServer(s *grpc.Server, srv IndexServer) {
	s.RegisterService(&_Index_serviceDesc, srv)
}

func _Index_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Join(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Leave(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetNode(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetCluster(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Document)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Get(ctx, req.(*Document))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Index_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IndexServer).Index(&indexIndexServer{stream})
}

type Index_IndexServer interface {
	SendAndClose(*UpdateResult) error
	Recv() (*Document, error)
	grpc.ServerStream
}

type indexIndexServer struct {
	grpc.ServerStream
}

func (x *indexIndexServer) SendAndClose(m *UpdateResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *indexIndexServer) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Index_Delete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IndexServer).Delete(&indexDeleteServer{stream})
}

type Index_DeleteServer interface {
	SendAndClose(*UpdateResult) error
	Recv() (*Document, error)
	grpc.ServerStream
}

type indexDeleteServer struct {
	grpc.ServerStream
}

func (x *indexDeleteServer) SendAndClose(m *UpdateResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *indexDeleteServer) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Index_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetIndexConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetIndexConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetIndexConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetIndexConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetIndexStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetIndexStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetIndexStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetIndexStats(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Index_serviceDesc = grpc.ServiceDesc{
	ServiceName: "index.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Index_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Index_Leave_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Index_GetNode_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Index_GetCluster_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Index_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Index_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Index_Search_Handler,
		},
		{
			MethodName: "GetIndexConfig",
			Handler:    _Index_GetIndexConfig_Handler,
		},
		{
			MethodName: "GetIndexStats",
			Handler:    _Index_GetIndexStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Index",
			Handler:       _Index_Index_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Delete",
			Handler:       _Index_Delete_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/index/index.proto",
}
