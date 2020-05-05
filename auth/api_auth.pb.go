// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api_auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type AuthReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_055a4330e6a114fd, []int{0}
}
func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return m.Size()
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

type AuthReply struct {
	Appkey               string   `protobuf:"bytes,1,opt,name=appkey,proto3" json:"app_key"`
	Appsecret            string   `protobuf:"bytes,2,opt,name=appsecret,proto3" json:"app_secret"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_055a4330e6a114fd, []int{1}
}
func (m *AuthReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReply.Merge(m, src)
}
func (m *AuthReply) XXX_Size() int {
	return m.Size()
}
func (m *AuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthReq)(nil), "service.auth.v1.AuthReq")
	proto.RegisterType((*AuthReply)(nil), "service.auth.v1.AuthReply")
}

func init() { proto.RegisterFile("api_auth.proto", fileDescriptor_055a4330e6a114fd) }

var fileDescriptor_055a4330e6a114fd = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2c, 0xc8, 0x8c,
	0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x8b, 0x95, 0x19, 0x4a, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xd5, 0x25, 0x95, 0xa6, 0x81,
	0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xaf, 0xe4, 0xc4, 0xc5, 0xee, 0x58, 0x5a, 0x92, 0x11, 0x94,
	0x5a, 0x28, 0x64, 0xce, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xec, 0xa4, 0xfe,
	0xea, 0x9e, 0x3c, 0x53, 0x66, 0xca, 0xa7, 0x7b, 0xf2, 0xb2, 0x69, 0xf9, 0x45, 0xb9, 0x56, 0x4a,
	0x99, 0x29, 0x4a, 0x0a, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0x56, 0x4a, 0x45, 0xa9,
	0x85, 0xa5, 0x99, 0x45, 0xa9, 0x29, 0x4a, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x71, 0x5c, 0x9c, 0x10,
	0x33, 0x0a, 0x72, 0x2a, 0x85, 0x94, 0xb9, 0xd8, 0x12, 0x0b, 0x0a, 0xb2, 0x53, 0x2b, 0xc1, 0x26,
	0x71, 0x3a, 0x71, 0xbf, 0xba, 0x27, 0xcf, 0x9e, 0x58, 0x50, 0x10, 0x9f, 0x9d, 0x5a, 0x19, 0x04,
	0x95, 0x12, 0xd2, 0xe1, 0xe2, 0x4c, 0x2c, 0x28, 0x28, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x91, 0x60,
	0x02, 0xab, 0xe3, 0x7b, 0x75, 0x4f, 0x9e, 0x0b, 0xa4, 0x0e, 0x22, 0x1a, 0x84, 0x50, 0x60, 0xe4,
	0xc7, 0xc5, 0xee, 0x58, 0x90, 0x09, 0xb2, 0x42, 0xc8, 0x99, 0x8b, 0xdb, 0x3d, 0xb5, 0x04, 0xc4,
	0xf4, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xd0, 0x43, 0xf3, 0xbe, 0x1e, 0xd4, 0x33, 0x52, 0x52, 0x38,
	0x64, 0x0a, 0x72, 0x2a, 0x9d, 0xa4, 0x4e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x01, 0x29, 0x4c,
	0x62, 0x03, 0x07, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x44, 0x13, 0x76, 0x10, 0x68, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiAuthClient is the client API for ApiAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiAuthClient interface {
	GetAuthInfo(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error)
}

type apiAuthClient struct {
	cc *grpc.ClientConn
}

func NewApiAuthClient(cc *grpc.ClientConn) ApiAuthClient {
	return &apiAuthClient{cc}
}

func (c *apiAuthClient) GetAuthInfo(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/service.auth.v1.ApiAuth/GetAuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiAuthServer is the server API for ApiAuth service.
type ApiAuthServer interface {
	GetAuthInfo(context.Context, *AuthReq) (*AuthReply, error)
}

// UnimplementedApiAuthServer can be embedded to have forward compatible implementations.
type UnimplementedApiAuthServer struct {
}

func (*UnimplementedApiAuthServer) GetAuthInfo(ctx context.Context, req *AuthReq) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthInfo not implemented")
}

func RegisterApiAuthServer(s *grpc.Server, srv ApiAuthServer) {
	s.RegisterService(&_ApiAuth_serviceDesc, srv)
}

func _ApiAuth_GetAuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiAuthServer).GetAuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.auth.v1.ApiAuth/GetAuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiAuthServer).GetAuthInfo(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.auth.v1.ApiAuth",
	HandlerType: (*ApiAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthInfo",
			Handler:    _ApiAuth_GetAuthInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_auth.proto",
}

func (m *AuthReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintApiAuth(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuthReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintApiAuth(dAtA, i, uint64(len(m.Appsecret)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Appkey) > 0 {
		i -= len(m.Appkey)
		copy(dAtA[i:], m.Appkey)
		i = encodeVarintApiAuth(dAtA, i, uint64(len(m.Appkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApiAuth(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Appkey)
	if l > 0 {
		n += 1 + l + sovApiAuth(uint64(l))
	}
	l = len(m.Appsecret)
	if l > 0 {
		n += 1 + l + sovApiAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApiAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApiAuth(x uint64) (n int) {
	return sovApiAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiAuth
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
			return fmt.Errorf("proto: AuthReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApiAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiAuth
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
func (m *AuthReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiAuth
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
			return fmt.Errorf("proto: AuthReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Appkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiAuth
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
				return ErrInvalidLengthApiAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuth
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
					return ErrIntOverflowApiAuth
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
				return ErrInvalidLengthApiAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Appsecret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiAuth
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
func skipApiAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiAuth
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
					return 0, ErrIntOverflowApiAuth
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
					return 0, ErrIntOverflowApiAuth
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
				return 0, ErrInvalidLengthApiAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiAuth = fmt.Errorf("proto: unexpected end of group")
)
