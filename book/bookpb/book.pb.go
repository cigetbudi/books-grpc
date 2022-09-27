// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book/bookpb/book.proto

//kalo beda folder wajib tambahin package dan go_package

package bookpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Book struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	IsRead               bool     `protobuf:"varint,4,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa3eebde80ce761, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetIsRead() bool {
	if m != nil {
		return m.IsRead
	}
	return false
}

type GetBookRequest struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa3eebde80ce761, []int{1}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type GetBookResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *Book    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookResponse) Reset()         { *m = GetBookResponse{} }
func (m *GetBookResponse) String() string { return proto.CompactTextString(m) }
func (*GetBookResponse) ProtoMessage()    {}
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa3eebde80ce761, []int{2}
}

func (m *GetBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookResponse.Unmarshal(m, b)
}
func (m *GetBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookResponse.Marshal(b, m, deterministic)
}
func (m *GetBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookResponse.Merge(m, src)
}
func (m *GetBookResponse) XXX_Size() int {
	return xxx_messageInfo_GetBookResponse.Size(m)
}
func (m *GetBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookResponse proto.InternalMessageInfo

func (m *GetBookResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetBookResponse) GetData() *Book {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "book.Book")
	proto.RegisterType((*GetBookRequest)(nil), "book.GetBookRequest")
	proto.RegisterType((*GetBookResponse)(nil), "book.GetBookResponse")
}

func init() { proto.RegisterFile("book/bookpb/book.proto", fileDescriptor_eaa3eebde80ce761) }

var fileDescriptor_eaa3eebde80ce761 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0xad, 0xdd, 0x3a, 0xca, 0x0a, 0x61, 0x5d, 0x8b, 0x07, 0x59, 0x7a, 0xda, 0xd3,
	0x2a, 0xeb, 0xc5, 0x73, 0x2f, 0xea, 0x35, 0xde, 0x04, 0x91, 0xd4, 0x0c, 0x18, 0x2a, 0x4e, 0x4d,
	0xa6, 0xfe, 0x7e, 0xc9, 0xa4, 0x08, 0xdb, 0x4b, 0x9b, 0xf7, 0x98, 0x79, 0xdf, 0x63, 0x60, 0xdd,
	0x11, 0xf5, 0xb7, 0xf1, 0x33, 0x74, 0xf2, 0xdb, 0x0d, 0x9e, 0x98, 0x54, 0x11, 0xdf, 0xcd, 0x1b,
	0x14, 0x2d, 0x51, 0xaf, 0x96, 0x90, 0x3b, 0x5b, 0x67, 0x9b, 0x6c, 0x7b, 0xaa, 0x73, 0x67, 0xd5,
	0x0a, 0x4e, 0xd8, 0xf1, 0x17, 0xd6, 0xb9, 0x58, 0x49, 0xa8, 0x35, 0x94, 0x66, 0xe4, 0x4f, 0xf2,
	0xf5, 0xb1, 0xd8, 0x93, 0x52, 0x57, 0xb0, 0x70, 0xe1, 0xdd, 0xa3, 0xb1, 0x75, 0xb1, 0xc9, 0xb6,
	0x95, 0x2e, 0x5d, 0xd0, 0x68, 0x6c, 0x73, 0x07, 0xcb, 0x47, 0xe4, 0x48, 0xd0, 0xf8, 0x33, 0x62,
	0x60, 0x75, 0x03, 0x02, 0x16, 0xd4, 0xd9, 0x1e, 0x76, 0xd2, 0x48, 0x06, 0x52, 0xa1, 0x67, 0xb8,
	0xf8, 0xdf, 0x08, 0x03, 0x7d, 0x07, 0xa1, 0x06, 0x36, 0x3c, 0x06, 0x59, 0xaa, 0xf4, 0xa4, 0x62,
	0x94, 0x35, 0x6c, 0xa4, 0xe2, 0x2c, 0x2a, 0xfa, 0xfb, 0x27, 0x38, 0x6f, 0x89, 0xa8, 0x7f, 0x41,
	0xff, 0xeb, 0x3e, 0x50, 0x3d, 0xc0, 0x62, 0x8a, 0x56, 0xab, 0x34, 0x7c, 0xd8, 0xed, 0xfa, 0x72,
	0xe6, 0x26, 0x7e, 0x73, 0xd4, 0x56, 0xaf, 0x65, 0x3a, 0x60, 0x57, 0xca, 0xf1, 0xee, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x0f, 0xc8, 0x36, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BoookServiceClient is the client API for BoookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BoookServiceClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type boookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBoookServiceClient(cc *grpc.ClientConn) BoookServiceClient {
	return &boookServiceClient{cc}
}

func (c *boookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/book.BoookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoookServiceServer is the server API for BoookService service.
type BoookServiceServer interface {
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
}

// UnimplementedBoookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBoookServiceServer struct {
}

func (*UnimplementedBoookServiceServer) GetBook(ctx context.Context, req *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}

func RegisterBoookServiceServer(s *grpc.Server, srv BoookServiceServer) {
	s.RegisterService(&_BoookService_serviceDesc, srv)
}

func _BoookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BoookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BoookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BoookService",
	HandlerType: (*BoookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BoookService_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/bookpb/book.proto",
}