// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package pb_calculator

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

// ===========Subtract===========
type SubtractRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractRequest) Reset()         { *m = SubtractRequest{} }
func (m *SubtractRequest) String() string { return proto.CompactTextString(m) }
func (*SubtractRequest) ProtoMessage()    {}
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *SubtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractRequest.Unmarshal(m, b)
}
func (m *SubtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractRequest.Marshal(b, m, deterministic)
}
func (m *SubtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractRequest.Merge(m, src)
}
func (m *SubtractRequest) XXX_Size() int {
	return xxx_messageInfo_SubtractRequest.Size(m)
}
func (m *SubtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractRequest proto.InternalMessageInfo

func (m *SubtractRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SubtractRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SubtractReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Sub                  int32    `protobuf:"varint,3,opt,name=sub,proto3" json:"sub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractReply) Reset()         { *m = SubtractReply{} }
func (m *SubtractReply) String() string { return proto.CompactTextString(m) }
func (*SubtractReply) ProtoMessage()    {}
func (*SubtractReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SubtractReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractReply.Unmarshal(m, b)
}
func (m *SubtractReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractReply.Marshal(b, m, deterministic)
}
func (m *SubtractReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractReply.Merge(m, src)
}
func (m *SubtractReply) XXX_Size() int {
	return xxx_messageInfo_SubtractReply.Size(m)
}
func (m *SubtractReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractReply proto.InternalMessageInfo

func (m *SubtractReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SubtractReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *SubtractReply) GetSub() int32 {
	if m != nil {
		return m.Sub
	}
	return 0
}

// ===========Multiply===========
type MultiplyRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyRequest) Reset()         { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()    {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *MultiplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyRequest.Unmarshal(m, b)
}
func (m *MultiplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyRequest.Marshal(b, m, deterministic)
}
func (m *MultiplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyRequest.Merge(m, src)
}
func (m *MultiplyRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplyRequest.Size(m)
}
func (m *MultiplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyRequest proto.InternalMessageInfo

func (m *MultiplyRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MultiplyRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type MultiplyReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Mult                 int32    `protobuf:"varint,3,opt,name=mult,proto3" json:"mult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyReply) Reset()         { *m = MultiplyReply{} }
func (m *MultiplyReply) String() string { return proto.CompactTextString(m) }
func (*MultiplyReply) ProtoMessage()    {}
func (*MultiplyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *MultiplyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyReply.Unmarshal(m, b)
}
func (m *MultiplyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyReply.Marshal(b, m, deterministic)
}
func (m *MultiplyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyReply.Merge(m, src)
}
func (m *MultiplyReply) XXX_Size() int {
	return xxx_messageInfo_MultiplyReply.Size(m)
}
func (m *MultiplyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyReply.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyReply proto.InternalMessageInfo

func (m *MultiplyReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MultiplyReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MultiplyReply) GetMult() int32 {
	if m != nil {
		return m.Mult
	}
	return 0
}

// ===========Pi===========
type PiRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PiRequest) Reset()         { *m = PiRequest{} }
func (m *PiRequest) String() string { return proto.CompactTextString(m) }
func (*PiRequest) ProtoMessage()    {}
func (*PiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *PiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PiRequest.Unmarshal(m, b)
}
func (m *PiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PiRequest.Marshal(b, m, deterministic)
}
func (m *PiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PiRequest.Merge(m, src)
}
func (m *PiRequest) XXX_Size() int {
	return xxx_messageInfo_PiRequest.Size(m)
}
func (m *PiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PiRequest proto.InternalMessageInfo

func (m *PiRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type PiReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Pi                   string   `protobuf:"bytes,3,opt,name=pi,proto3" json:"pi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PiReply) Reset()         { *m = PiReply{} }
func (m *PiReply) String() string { return proto.CompactTextString(m) }
func (*PiReply) ProtoMessage()    {}
func (*PiReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{5}
}

func (m *PiReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PiReply.Unmarshal(m, b)
}
func (m *PiReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PiReply.Marshal(b, m, deterministic)
}
func (m *PiReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PiReply.Merge(m, src)
}
func (m *PiReply) XXX_Size() int {
	return xxx_messageInfo_PiReply.Size(m)
}
func (m *PiReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PiReply.DiscardUnknown(m)
}

var xxx_messageInfo_PiReply proto.InternalMessageInfo

func (m *PiReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PiReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PiReply) GetPi() string {
	if m != nil {
		return m.Pi
	}
	return ""
}

func init() {
	proto.RegisterType((*SubtractRequest)(nil), "pb.calculator.SubtractRequest")
	proto.RegisterType((*SubtractReply)(nil), "pb.calculator.SubtractReply")
	proto.RegisterType((*MultiplyRequest)(nil), "pb.calculator.MultiplyRequest")
	proto.RegisterType((*MultiplyReply)(nil), "pb.calculator.MultiplyReply")
	proto.RegisterType((*PiRequest)(nil), "pb.calculator.PiRequest")
	proto.RegisterType((*PiReply)(nil), "pb.calculator.PiReply")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0xb4, 0x01, 0x72, 0x22, 0x50, 0x9d, 0x2a, 0x14, 0x55, 0x08, 0x81, 0x27, 0x16,
	0x32, 0xc0, 0xc2, 0xde, 0x0d, 0x54, 0xa9, 0x0a, 0xbf, 0xc0, 0x36, 0x19, 0x2c, 0x19, 0x6c, 0x5c,
	0x7b, 0xc8, 0x9f, 0xe5, 0xb7, 0x20, 0x3b, 0x49, 0xa3, 0xa6, 0x20, 0xd1, 0xed, 0x5e, 0xfc, 0xee,
	0x7d, 0xd2, 0xbb, 0xc0, 0x5c, 0x30, 0x25, 0xbc, 0x62, 0x4e, 0xdb, 0xd2, 0x58, 0xed, 0x34, 0xe6,
	0x86, 0x97, 0xc3, 0x47, 0xfa, 0x00, 0x97, 0x6f, 0x9e, 0x3b, 0xcb, 0x84, 0xab, 0xea, 0x2f, 0x5f,
	0x6f, 0x1d, 0x9e, 0x03, 0x61, 0x05, 0xb9, 0x25, 0xf7, 0x69, 0x45, 0x58, 0x50, 0xbc, 0x48, 0x5a,
	0xc5, 0xe9, 0x2b, 0xe4, 0x83, 0xdd, 0xa8, 0x06, 0x11, 0x66, 0x42, 0xbf, 0xd7, 0x9d, 0x3f, 0xce,
	0xb8, 0x80, 0xb4, 0xb6, 0x56, 0xdb, 0xb8, 0x96, 0x55, 0xad, 0xc0, 0x39, 0x4c, 0xb7, 0x9e, 0x17,
	0xd3, 0x68, 0x0c, 0x63, 0x60, 0xaf, 0xbd, 0x72, 0xd2, 0xa8, 0xe6, 0x3f, 0xec, 0x35, 0xe4, 0x83,
	0xfd, 0x38, 0x36, 0xc2, 0xec, 0xc3, 0x2b, 0xd7, 0xc1, 0xe3, 0x4c, 0xef, 0x20, 0xdb, 0xc8, 0x9e,
	0xbb, 0x80, 0x54, 0x68, 0xff, 0xe9, 0xba, 0xac, 0x56, 0xd0, 0x15, 0x9c, 0x06, 0xcb, 0x71, 0xac,
	0x0b, 0x48, 0x8c, 0x8c, 0xa4, 0xac, 0x4a, 0x8c, 0x7c, 0xfc, 0x26, 0x00, 0xab, 0x5d, 0xe1, 0xf8,
	0x02, 0x67, 0x7d, 0x83, 0x78, 0x53, 0xee, 0x1d, 0xa3, 0x1c, 0x5d, 0x62, 0x79, 0xfd, 0xe7, 0xbb,
	0x51, 0x0d, 0x9d, 0x84, 0xac, 0xbe, 0x91, 0x83, 0xac, 0x51, 0xb3, 0x07, 0x59, 0x7b, 0x55, 0xd2,
	0x09, 0x3e, 0x43, 0xb2, 0x91, 0x58, 0x8c, 0x5c, 0xbb, 0x86, 0x96, 0x57, 0xbf, 0xbc, 0xc4, 0x4d,
	0x7e, 0x12, 0x7f, 0xac, 0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x03, 0x6f, 0xa3, 0x6c,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	// Subtract subtracts one number for other
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractReply, error)
	// Multiply multiplies too numbers
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error)
	// Pi returns the pi number
	Pi(ctx context.Context, in *PiRequest, opts ...grpc.CallOption) (*PiReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractReply, error) {
	out := new(SubtractReply)
	err := c.cc.Invoke(ctx, "/pb.calculator.Calculator/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error) {
	out := new(MultiplyReply)
	err := c.cc.Invoke(ctx, "/pb.calculator.Calculator/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Pi(ctx context.Context, in *PiRequest, opts ...grpc.CallOption) (*PiReply, error) {
	out := new(PiReply)
	err := c.cc.Invoke(ctx, "/pb.calculator.Calculator/Pi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	// Subtract subtracts one number for other
	Subtract(context.Context, *SubtractRequest) (*SubtractReply, error)
	// Multiply multiplies too numbers
	Multiply(context.Context, *MultiplyRequest) (*MultiplyReply, error)
	// Pi returns the pi number
	Pi(context.Context, *PiRequest) (*PiReply, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Subtract(ctx context.Context, req *SubtractRequest) (*SubtractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (*UnimplementedCalculatorServer) Multiply(ctx context.Context, req *MultiplyRequest) (*MultiplyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalculatorServer) Pi(ctx context.Context, req *PiRequest) (*PiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pi not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.calculator.Calculator/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.calculator.Calculator/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Pi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Pi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.calculator.Calculator/Pi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Pi(ctx, req.(*PiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subtract",
			Handler:    _Calculator_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculator_Multiply_Handler,
		},
		{
			MethodName: "Pi",
			Handler:    _Calculator_Pi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
