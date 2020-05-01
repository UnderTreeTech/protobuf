// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app.proto

package app

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AppReq struct {
	Sappkey              string   `protobuf:"bytes,1,opt,name=sappkey,proto3" json:"sappkey" form:"sappkey" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReq) Reset()         { *m = AppReq{} }
func (m *AppReq) String() string { return proto.CompactTextString(m) }
func (*AppReq) ProtoMessage()    {}
func (*AppReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}
func (m *AppReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReq.Merge(m, src)
}
func (m *AppReq) XXX_Size() int {
	return m.Size()
}
func (m *AppReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppReq proto.InternalMessageInfo

func (m *AppReq) GetSappkey() string {
	if m != nil {
		return m.Sappkey
	}
	return ""
}

type AppReply struct {
	Appkey               string   `protobuf:"bytes,1,opt,name=appkey,proto3" json:"app_key"`
	Appsecret            string   `protobuf:"bytes,2,opt,name=appsecret,proto3" json:"app_secret"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReply) Reset()         { *m = AppReply{} }
func (m *AppReply) String() string { return proto.CompactTextString(m) }
func (*AppReply) ProtoMessage()    {}
func (*AppReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}
func (m *AppReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReply.Merge(m, src)
}
func (m *AppReply) XXX_Size() int {
	return m.Size()
}
func (m *AppReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppReply proto.InternalMessageInfo

func (m *AppReply) GetAppkey() string {
	if m != nil {
		return m.Appkey
	}
	return ""
}

func (m *AppReply) GetAppsecret() string {
	if m != nil {
		return m.Appsecret
	}
	return ""
}

type SkipUrlsReply struct {
	Urls                 []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SkipUrlsReply) Reset()         { *m = SkipUrlsReply{} }
func (m *SkipUrlsReply) String() string { return proto.CompactTextString(m) }
func (*SkipUrlsReply) ProtoMessage()    {}
func (*SkipUrlsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}
func (m *SkipUrlsReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SkipUrlsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SkipUrlsReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SkipUrlsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkipUrlsReply.Merge(m, src)
}
func (m *SkipUrlsReply) XXX_Size() int {
	return m.Size()
}
func (m *SkipUrlsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SkipUrlsReply.DiscardUnknown(m)
}

var xxx_messageInfo_SkipUrlsReply proto.InternalMessageInfo

func (m *SkipUrlsReply) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func init() {
	proto.RegisterType((*AppReq)(nil), "service.app.v1.AppReq")
	proto.RegisterType((*AppReply)(nil), "service.app.v1.AppReply")
	proto.RegisterType((*SkipUrlsReply)(nil), "service.app.v1.SkipUrlsReply")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5d, 0x21, 0x08, 0xab, 0x72, 0xd8, 0x44, 0xd2, 0x54, 0x6d, 0xc9, 0x7a, 0xe1, 0x20,
	0x4b, 0xd0, 0xc4, 0x03, 0x27, 0x21, 0x31, 0x1c, 0xbc, 0x98, 0x12, 0x2e, 0x26, 0xc6, 0x14, 0x58,
	0x6a, 0x43, 0x71, 0x87, 0x6d, 0x4b, 0xd2, 0x37, 0xf1, 0x91, 0x38, 0xfa, 0x04, 0x8d, 0xc1, 0x1b,
	0x47, 0x9f, 0xc0, 0x74, 0x5b, 0x42, 0x6a, 0xe4, 0xd2, 0x74, 0xbe, 0xf9, 0xfb, 0xcf, 0xcc, 0x5f,
	0x5c, 0xb1, 0x01, 0x18, 0x48, 0x11, 0x08, 0x52, 0xf5, 0xb9, 0x5c, 0xba, 0x63, 0xce, 0x12, 0xb4,
	0x6c, 0xeb, 0x4d, 0xc7, 0x0d, 0xde, 0xc2, 0x11, 0x1b, 0x8b, 0x79, 0xcb, 0x11, 0x8e, 0x68, 0x29,
	0xd9, 0x28, 0x9c, 0xaa, 0x4a, 0x15, 0xea, 0x2d, 0xfd, 0x5c, 0x3f, 0x77, 0x84, 0x70, 0x3c, 0xbe,
	0x53, 0xf1, 0x39, 0x04, 0x51, 0xda, 0xa4, 0x43, 0x5c, 0xea, 0x02, 0x58, 0x7c, 0x41, 0x1e, 0xf1,
	0x91, 0x6f, 0x03, 0xcc, 0x78, 0xa4, 0xa1, 0x3a, 0x6a, 0x54, 0x7a, 0xed, 0x4d, 0x6c, 0x6e, 0xd1,
	0x4f, 0x6c, 0xd2, 0xa9, 0x90, 0xf3, 0x0e, 0xcd, 0x00, 0xad, 0x2f, 0x6d, 0xcf, 0x9d, 0xd8, 0x01,
	0xef, 0x50, 0xc9, 0x17, 0xa1, 0x2b, 0xf9, 0x84, 0x5a, 0x5b, 0x39, 0x7d, 0xc1, 0x65, 0x65, 0x0b,
	0x5e, 0x44, 0xae, 0x70, 0x29, 0xe7, 0x7b, 0x9c, 0xf8, 0xda, 0x00, 0xaf, 0x33, 0x1e, 0x59, 0x59,
	0x8b, 0x5c, 0xab, 0x83, 0x7d, 0x3e, 0x96, 0x3c, 0xd0, 0x0e, 0x95, 0xae, 0xba, 0x89, 0x4d, 0x9c,
	0xe8, 0x52, 0x6a, 0xed, 0x04, 0xb4, 0x89, 0x4f, 0x07, 0x33, 0x17, 0x86, 0xd2, 0xf3, 0xd3, 0x19,
	0x17, 0xb8, 0x18, 0x4a, 0xcf, 0xd7, 0x50, 0xbd, 0xd0, 0xa8, 0xf4, 0xca, 0x9b, 0xd8, 0x54, 0xb5,
	0xa5, 0x9e, 0x37, 0x2b, 0x84, 0x0b, 0x5d, 0x00, 0x72, 0x87, 0x8b, 0x4f, 0xee, 0xbb, 0x43, 0x6a,
	0x2c, 0x8d, 0x84, 0x6d, 0x23, 0x61, 0x0f, 0x49, 0x24, 0xfa, 0x1e, 0x4e, 0xee, 0xf1, 0x49, 0x9f,
	0x07, 0x5d, 0x80, 0x81, 0x1a, 0x4f, 0x6a, 0x2c, 0xff, 0x47, 0x58, 0x1a, 0xa1, 0xae, 0xfd, 0xcb,
	0x93, 0xfd, 0xfa, 0xb8, 0x9a, 0x39, 0x64, 0x6b, 0xef, 0xf5, 0xb8, 0xfc, 0xcb, 0x73, 0x87, 0xf6,
	0xce, 0x56, 0x6b, 0x03, 0x7d, 0xae, 0x0d, 0xf4, 0xb5, 0x36, 0xd0, 0xc7, 0xb7, 0x71, 0xf0, 0x5c,
	0xb0, 0x01, 0x46, 0x25, 0xb5, 0xf1, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xc8, 0x81,
	0x99, 0x36, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAppSecret(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppReply, error)
	GetAppSkipUrls(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*SkipUrlsReply, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/service.app.v1.App/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetAppSecret(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/service.app.v1.App/GetAppSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetAppSkipUrls(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*SkipUrlsReply, error) {
	out := new(SkipUrlsReply)
	err := c.cc.Invoke(ctx, "/service.app.v1.App/GetAppSkipUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	GetAppSecret(context.Context, *AppReq) (*AppReply, error)
	GetAppSkipUrls(context.Context, *AppReq) (*SkipUrlsReply, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAppServer) GetAppSecret(ctx context.Context, req *AppReq) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSecret not implemented")
}
func (*UnimplementedAppServer) GetAppSkipUrls(ctx context.Context, req *AppReq) (*SkipUrlsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSkipUrls not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.app.v1.App/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetAppSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetAppSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.app.v1.App/GetAppSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetAppSecret(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetAppSkipUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetAppSkipUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.app.v1.App/GetAppSkipUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetAppSkipUrls(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.app.v1.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _App_Ping_Handler,
		},
		{
			MethodName: "GetAppSecret",
			Handler:    _App_GetAppSecret_Handler,
		},
		{
			MethodName: "GetAppSkipUrls",
			Handler:    _App_GetAppSkipUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func (m *AppReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Sappkey) > 0 {
		i -= len(m.Sappkey)
		copy(dAtA[i:], m.Sappkey)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Sappkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AppReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Appsecret) > 0 {
		i -= len(m.Appsecret)
		copy(dAtA[i:], m.Appsecret)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Appsecret)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Appkey) > 0 {
		i -= len(m.Appkey)
		copy(dAtA[i:], m.Appkey)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Appkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SkipUrlsReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SkipUrlsReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SkipUrlsReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Urls) > 0 {
		for iNdEx := len(m.Urls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Urls[iNdEx])
			copy(dAtA[i:], m.Urls[iNdEx])
			i = encodeVarintApp(dAtA, i, uint64(len(m.Urls[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintApp(dAtA []byte, offset int, v uint64) int {
	offset -= sovApp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sappkey)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AppReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Appkey)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.Appsecret)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SkipUrlsReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Urls) > 0 {
		for _, s := range m.Urls {
			l = len(s)
			n += 1 + l + sovApp(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApp(x uint64) (n int) {
	return sovApp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sappkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sappkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Appkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Appkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Appsecret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Appsecret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SkipUrlsReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SkipUrlsReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SkipUrlsReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urls = append(m.Urls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApp = fmt.Errorf("proto: unexpected end of group")
)