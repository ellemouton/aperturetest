// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pricesrpc/rpc.proto

package pricesrpc

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

type GetPriceRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceRequest) Reset()         { *m = GetPriceRequest{} }
func (m *GetPriceRequest) String() string { return proto.CompactTextString(m) }
func (*GetPriceRequest) ProtoMessage()    {}
func (*GetPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a48c5d96e99c79ae, []int{0}
}

func (m *GetPriceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceRequest.Unmarshal(m, b)
}
func (m *GetPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceRequest.Marshal(b, m, deterministic)
}
func (m *GetPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceRequest.Merge(m, src)
}
func (m *GetPriceRequest) XXX_Size() int {
	return xxx_messageInfo_GetPriceRequest.Size(m)
}
func (m *GetPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceRequest proto.InternalMessageInfo

func (m *GetPriceRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type GetPriceResponse struct {
	Price                int64    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPriceResponse) Reset()         { *m = GetPriceResponse{} }
func (m *GetPriceResponse) String() string { return proto.CompactTextString(m) }
func (*GetPriceResponse) ProtoMessage()    {}
func (*GetPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a48c5d96e99c79ae, []int{1}
}

func (m *GetPriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPriceResponse.Unmarshal(m, b)
}
func (m *GetPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPriceResponse.Marshal(b, m, deterministic)
}
func (m *GetPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPriceResponse.Merge(m, src)
}
func (m *GetPriceResponse) XXX_Size() int {
	return xxx_messageInfo_GetPriceResponse.Size(m)
}
func (m *GetPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPriceResponse proto.InternalMessageInfo

func (m *GetPriceResponse) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPriceRequest)(nil), "pricesrpc.GetPriceRequest")
	proto.RegisterType((*GetPriceResponse)(nil), "pricesrpc.GetPriceResponse")
}

func init() { proto.RegisterFile("pricesrpc/rpc.proto", fileDescriptor_a48c5d96e99c79ae) }

var fileDescriptor_a48c5d96e99c79ae = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x4c,
	0x4e, 0x2d, 0x2e, 0x2a, 0x48, 0xd6, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x84, 0x0b, 0x2a, 0xa9, 0x72, 0xf1, 0xbb, 0xa7, 0x96, 0x04, 0x80, 0xf8, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x05, 0x89, 0x25, 0x19, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x06, 0x97, 0x00, 0x42, 0x59, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x08, 0x17, 0x2b, 0xd8, 0x1c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0xc7,
	0xc8, 0x97, 0x8b, 0x0d, 0xac, 0xac, 0x58, 0xc8, 0x99, 0x8b, 0x03, 0xa6, 0x47, 0x48, 0x4a, 0x0f,
	0x6e, 0xa5, 0x1e, 0x9a, 0x7d, 0x52, 0xd2, 0x58, 0xe5, 0x20, 0x96, 0x38, 0xf1, 0x46, 0x71, 0xeb,
	0xe9, 0xc3, 0xe5, 0x93, 0xd8, 0xc0, 0x1e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xeb,
	0x5e, 0x1e, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PricesClient is the client API for Prices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PricesClient interface {
	GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error)
}

type pricesClient struct {
	cc *grpc.ClientConn
}

func NewPricesClient(cc *grpc.ClientConn) PricesClient {
	return &pricesClient{cc}
}

func (c *pricesClient) GetPrice(ctx context.Context, in *GetPriceRequest, opts ...grpc.CallOption) (*GetPriceResponse, error) {
	out := new(GetPriceResponse)
	err := c.cc.Invoke(ctx, "/pricesrpc.Prices/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricesServer is the server API for Prices service.
type PricesServer interface {
	GetPrice(context.Context, *GetPriceRequest) (*GetPriceResponse, error)
}

// UnimplementedPricesServer can be embedded to have forward compatible implementations.
type UnimplementedPricesServer struct {
}

func (*UnimplementedPricesServer) GetPrice(ctx context.Context, req *GetPriceRequest) (*GetPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}

func RegisterPricesServer(s *grpc.Server, srv PricesServer) {
	s.RegisterService(&_Prices_serviceDesc, srv)
}

func _Prices_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricesrpc.Prices/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServer).GetPrice(ctx, req.(*GetPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricesrpc.Prices",
	HandlerType: (*PricesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrice",
			Handler:    _Prices_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricesrpc/rpc.proto",
}
