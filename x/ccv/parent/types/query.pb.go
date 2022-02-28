// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/parent/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/interchain-security/x/ccv/child/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryChildGenesisRequest struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryChildGenesisRequest) Reset()         { *m = QueryChildGenesisRequest{} }
func (m *QueryChildGenesisRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChildGenesisRequest) ProtoMessage()    {}
func (*QueryChildGenesisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c3e3e81de812da, []int{0}
}
func (m *QueryChildGenesisRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChildGenesisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChildGenesisRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChildGenesisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChildGenesisRequest.Merge(m, src)
}
func (m *QueryChildGenesisRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChildGenesisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChildGenesisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChildGenesisRequest proto.InternalMessageInfo

func (m *QueryChildGenesisRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type QueryChildGenesisResponse struct {
	GenesisState types.GenesisState `protobuf:"bytes,1,opt,name=genesis_state,json=genesisState,proto3" json:"genesis_state"`
}

func (m *QueryChildGenesisResponse) Reset()         { *m = QueryChildGenesisResponse{} }
func (m *QueryChildGenesisResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChildGenesisResponse) ProtoMessage()    {}
func (*QueryChildGenesisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c3e3e81de812da, []int{1}
}
func (m *QueryChildGenesisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChildGenesisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChildGenesisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChildGenesisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChildGenesisResponse.Merge(m, src)
}
func (m *QueryChildGenesisResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChildGenesisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChildGenesisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChildGenesisResponse proto.InternalMessageInfo

func (m *QueryChildGenesisResponse) GetGenesisState() types.GenesisState {
	if m != nil {
		return m.GenesisState
	}
	return types.GenesisState{}
}

func init() {
	proto.RegisterType((*QueryChildGenesisRequest)(nil), "interchain_security.ccv.parent.v1.QueryChildGenesisRequest")
	proto.RegisterType((*QueryChildGenesisResponse)(nil), "interchain_security.ccv.parent.v1.QueryChildGenesisResponse")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/parent/v1/query.proto", fileDescriptor_84c3e3e81de812da)
}

var fileDescriptor_84c3e3e81de812da = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x4b, 0xfb, 0x30,
	0x14, 0x6f, 0xbe, 0x7c, 0xfd, 0x15, 0xe7, 0xa5, 0x78, 0xd8, 0x86, 0x54, 0xdd, 0xc9, 0xcb, 0x12,
	0x36, 0x11, 0x86, 0x7a, 0x90, 0x79, 0x10, 0x8f, 0xd6, 0x93, 0x5e, 0x46, 0x97, 0x85, 0x2c, 0xb0,
	0x25, 0x5d, 0x93, 0x06, 0x87, 0x78, 0xf1, 0x2f, 0x10, 0xfc, 0xa7, 0x06, 0x5e, 0x06, 0x22, 0x78,
	0x12, 0xd9, 0xfc, 0x43, 0xa4, 0x69, 0x95, 0x1e, 0x56, 0x06, 0xde, 0x5e, 0xdf, 0x7b, 0x9f, 0x5f,
	0xaf, 0x81, 0x75, 0x2e, 0x34, 0x8d, 0x48, 0x3f, 0xe0, 0xa2, 0xa3, 0x28, 0x89, 0x23, 0xae, 0xc7,
	0x98, 0x10, 0x83, 0xc3, 0x20, 0xa2, 0x42, 0x63, 0xd3, 0xc0, 0xa3, 0x98, 0x46, 0x63, 0x14, 0x46,
	0x52, 0x4b, 0x77, 0x7f, 0xc1, 0x3a, 0x22, 0xc4, 0xa0, 0x74, 0x1d, 0x99, 0x46, 0x75, 0x87, 0x49,
	0xc9, 0x06, 0x14, 0x07, 0x21, 0xc7, 0x81, 0x10, 0x52, 0x07, 0x9a, 0x4b, 0xa1, 0x52, 0x82, 0xea,
	0x36, 0x93, 0x4c, 0xda, 0x12, 0x27, 0x55, 0xd6, 0x45, 0x45, 0x2e, 0x48, 0x9f, 0x0f, 0x7a, 0x89,
	0x09, 0x46, 0x05, 0x55, 0x3c, 0x63, 0xa9, 0x1d, 0xc1, 0xf2, 0x55, 0xe2, 0xea, 0x3c, 0x19, 0x5f,
	0xa4, 0x23, 0x9f, 0x8e, 0x62, 0xaa, 0xb4, 0x5b, 0x81, 0xeb, 0x29, 0x11, 0xef, 0x95, 0xc1, 0x1e,
	0x38, 0xd8, 0xf0, 0xd7, 0xec, 0xf7, 0x65, 0xaf, 0x66, 0x60, 0x65, 0x01, 0x4c, 0x85, 0x52, 0x28,
	0xea, 0xde, 0xc0, 0xad, 0x4c, 0xa4, 0xa3, 0x74, 0xa0, 0xa9, 0x05, 0x6f, 0x36, 0x11, 0x2a, 0x8a,
	0x6c, 0xbd, 0x21, 0xd3, 0x40, 0x19, 0xd3, 0x75, 0x82, 0x6a, 0xff, 0x9f, 0x7c, 0xec, 0x3a, 0x7e,
	0x89, 0xe5, 0x7a, 0xcd, 0x37, 0x00, 0x57, 0xac, 0xb0, 0xfb, 0x02, 0x60, 0x29, 0xaf, 0xee, 0x9e,
	0xa0, 0xa5, 0x17, 0x45, 0x45, 0x51, 0xab, 0xa7, 0x7f, 0x03, 0xa7, 0x81, 0x6b, 0x67, 0x8f, 0xaf,
	0x5f, 0xcf, 0xff, 0x8e, 0xdd, 0x16, 0x5e, 0xf2, 0x06, 0x6c, 0xd0, 0x4e, 0x96, 0x08, 0xdf, 0xff,
	0x5c, 0xf7, 0xa1, 0xed, 0x4f, 0x66, 0x1e, 0x98, 0xce, 0x3c, 0xf0, 0x39, 0xf3, 0xc0, 0xd3, 0xdc,
	0x73, 0xa6, 0x73, 0xcf, 0x79, 0x9f, 0x7b, 0xce, 0x6d, 0x8b, 0x71, 0xdd, 0x8f, 0xbb, 0x88, 0xc8,
	0x21, 0x26, 0x52, 0x0d, 0xa5, 0xca, 0x89, 0xd4, 0x7f, 0x45, 0xee, 0xf2, 0x32, 0x7a, 0x1c, 0x52,
	0xd5, 0x5d, 0xb5, 0x7f, 0xf8, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x11, 0xb4, 0x5f, 0x9a, 0x99,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ChildGenesis queries the genesis state needed to start a child chain whose proposal
	// has been accepted
	ChildGenesis(ctx context.Context, in *QueryChildGenesisRequest, opts ...grpc.CallOption) (*QueryChildGenesisResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ChildGenesis(ctx context.Context, in *QueryChildGenesisRequest, opts ...grpc.CallOption) (*QueryChildGenesisResponse, error) {
	out := new(QueryChildGenesisResponse)
	err := c.cc.Invoke(ctx, "/interchain_security.ccv.parent.v1.Query/ChildGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ChildGenesis queries the genesis state needed to start a child chain whose proposal
	// has been accepted
	ChildGenesis(context.Context, *QueryChildGenesisRequest) (*QueryChildGenesisResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ChildGenesis(ctx context.Context, req *QueryChildGenesisRequest) (*QueryChildGenesisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChildGenesis not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ChildGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChildGenesisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChildGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchain_security.ccv.parent.v1.Query/ChildGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChildGenesis(ctx, req.(*QueryChildGenesisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interchain_security.ccv.parent.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChildGenesis",
			Handler:    _Query_ChildGenesis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchain_security/ccv/parent/v1/query.proto",
}

func (m *QueryChildGenesisRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChildGenesisRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChildGenesisRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChildGenesisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChildGenesisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChildGenesisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.GenesisState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryChildGenesisRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryChildGenesisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GenesisState.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryChildGenesisRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChildGenesisRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChildGenesisRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChildGenesisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChildGenesisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChildGenesisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)