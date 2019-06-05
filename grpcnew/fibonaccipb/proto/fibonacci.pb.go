// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fibonacci.proto

package proto

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

type Request struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff1ee5882f276bd, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FibonacciResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff1ee5882f276bd, []int{1}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "grpcnew.Request")
	proto.RegisterType((*FibonacciResponse)(nil), "grpcnew.FibonacciResponse")
}

func init() { proto.RegisterFile("fibonacci.proto", fileDescriptor_fff1ee5882f276bd) }

var fileDescriptor_fff1ee5882f276bd = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcb, 0x4c, 0xca,
	0xcf, 0x4b, 0x4c, 0x4e, 0xce, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2f, 0x2a,
	0x48, 0xce, 0x4b, 0x2d, 0x57, 0x52, 0xe4, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe3, 0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d,
	0x82, 0xf2, 0x94, 0xb4, 0xb9, 0x04, 0xdd, 0x60, 0xda, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x41, 0x8a, 0x8b, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x60, 0x8a, 0x21, 0x3c, 0xa3, 0x30, 0x2e,
	0x41, 0xe7, 0xc4, 0x9c, 0xe4, 0xd2, 0x9c, 0xc4, 0x92, 0xfc, 0xa2, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0x21, 0x47, 0x2e, 0x7e, 0xb8, 0x09, 0xc1, 0xa9, 0x45, 0x99, 0xa9, 0xc5, 0x42, 0x02,
	0x7a, 0x50, 0x17, 0xe8, 0x41, 0xad, 0x97, 0x92, 0x82, 0x8b, 0x60, 0xd8, 0xa6, 0xc4, 0xe0, 0xc4,
	0x1b, 0xc5, 0x0d, 0xf7, 0x43, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x1b, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x5b, 0xb1, 0xd9, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	FibonacciSeries(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FibonacciResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) FibonacciSeries(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FibonacciResponse, error) {
	out := new(FibonacciResponse)
	err := c.cc.Invoke(ctx, "/grpcnew.CalculatorService/FibonacciSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	FibonacciSeries(context.Context, *Request) (*FibonacciResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) FibonacciSeries(ctx context.Context, req *Request) (*FibonacciResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FibonacciSeries not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_FibonacciSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).FibonacciSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcnew.CalculatorService/FibonacciSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).FibonacciSeries(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcnew.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FibonacciSeries",
			Handler:    _CalculatorService_FibonacciSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fibonacci.proto",
}