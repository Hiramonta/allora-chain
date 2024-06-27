// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/types.proto

package types

import (
	fmt "fmt"
	github_com_allora_network_allora_chain_math "github.com/allora-network/allora-chain/math"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type SimpleCursorPaginationRequest struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *SimpleCursorPaginationRequest) Reset()         { *m = SimpleCursorPaginationRequest{} }
func (m *SimpleCursorPaginationRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleCursorPaginationRequest) ProtoMessage()    {}
func (*SimpleCursorPaginationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32ba9f7333c0b525, []int{0}
}
func (m *SimpleCursorPaginationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleCursorPaginationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleCursorPaginationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleCursorPaginationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleCursorPaginationRequest.Merge(m, src)
}
func (m *SimpleCursorPaginationRequest) XXX_Size() int {
	return m.Size()
}
func (m *SimpleCursorPaginationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleCursorPaginationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleCursorPaginationRequest proto.InternalMessageInfo

func (m *SimpleCursorPaginationRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SimpleCursorPaginationRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type SimpleCursorPaginationResponse struct {
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
}

func (m *SimpleCursorPaginationResponse) Reset()         { *m = SimpleCursorPaginationResponse{} }
func (m *SimpleCursorPaginationResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleCursorPaginationResponse) ProtoMessage()    {}
func (*SimpleCursorPaginationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32ba9f7333c0b525, []int{1}
}
func (m *SimpleCursorPaginationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleCursorPaginationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleCursorPaginationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleCursorPaginationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleCursorPaginationResponse.Merge(m, src)
}
func (m *SimpleCursorPaginationResponse) XXX_Size() int {
	return m.Size()
}
func (m *SimpleCursorPaginationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleCursorPaginationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleCursorPaginationResponse proto.InternalMessageInfo

func (m *SimpleCursorPaginationResponse) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

type ListeningCoefficient struct {
	Coefficient github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,1,opt,name=coefficient,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"coefficient"`
}

func (m *ListeningCoefficient) Reset()         { *m = ListeningCoefficient{} }
func (m *ListeningCoefficient) String() string { return proto.CompactTextString(m) }
func (*ListeningCoefficient) ProtoMessage()    {}
func (*ListeningCoefficient) Descriptor() ([]byte, []int) {
	return fileDescriptor_32ba9f7333c0b525, []int{2}
}
func (m *ListeningCoefficient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListeningCoefficient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListeningCoefficient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListeningCoefficient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListeningCoefficient.Merge(m, src)
}
func (m *ListeningCoefficient) XXX_Size() int {
	return m.Size()
}
func (m *ListeningCoefficient) XXX_DiscardUnknown() {
	xxx_messageInfo_ListeningCoefficient.DiscardUnknown(m)
}

var xxx_messageInfo_ListeningCoefficient proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SimpleCursorPaginationRequest)(nil), "emissions.v1.SimpleCursorPaginationRequest")
	proto.RegisterType((*SimpleCursorPaginationResponse)(nil), "emissions.v1.SimpleCursorPaginationResponse")
	proto.RegisterType((*ListeningCoefficient)(nil), "emissions.v1.ListeningCoefficient")
}

func init() { proto.RegisterFile("emissions/v1/types.proto", fileDescriptor_32ba9f7333c0b525) }

var fileDescriptor_32ba9f7333c0b525 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4f, 0x32, 0x41,
	0x10, 0x87, 0xef, 0xde, 0xd7, 0xbf, 0x2b, 0x85, 0x5e, 0x28, 0x80, 0xc4, 0x85, 0x50, 0xd1, 0x78,
	0x1b, 0x62, 0xa1, 0x89, 0x1d, 0x98, 0x58, 0x68, 0x61, 0xce, 0x4a, 0x1b, 0xb2, 0x5c, 0x86, 0x63,
	0xc2, 0xed, 0xce, 0x71, 0xbb, 0x20, 0x7c, 0x0b, 0x3f, 0x16, 0x25, 0xa5, 0xb1, 0x20, 0x06, 0xbe,
	0x88, 0xe1, 0x4e, 0x3d, 0x1b, 0x13, 0x9b, 0xcd, 0x3c, 0x33, 0xb3, 0x4f, 0x31, 0x3f, 0x56, 0x01,
	0x85, 0xc6, 0x20, 0x69, 0x23, 0xa6, 0x6d, 0x61, 0xe7, 0x09, 0x18, 0x3f, 0x49, 0xc9, 0x92, 0x57,
	0xfa, 0x9e, 0xf8, 0xd3, 0x76, 0xad, 0x1a, 0x92, 0x51, 0x64, 0x7a, 0xd9, 0x4c, 0xe4, 0x90, 0x2f,
	0xd6, 0x4e, 0xa4, 0x42, 0x4d, 0x22, 0x7b, 0x3f, 0x5b, 0xe5, 0x88, 0x22, 0xca, 0x57, 0xb7, 0x55,
	0xde, 0x6d, 0xde, 0xb0, 0xd3, 0x07, 0x54, 0x49, 0x0c, 0xdd, 0x49, 0x6a, 0x28, 0xbd, 0x97, 0x11,
	0x6a, 0x69, 0x91, 0x74, 0x00, 0xe3, 0x09, 0x18, 0xeb, 0x1d, 0xb3, 0xff, 0x23, 0x98, 0x57, 0xdc,
	0x86, 0xdb, 0x2a, 0x05, 0xdb, 0xd2, 0x2b, 0xb3, 0xdd, 0x18, 0x15, 0xda, 0xca, 0xbf, 0x86, 0xdb,
	0xda, 0x09, 0x72, 0x68, 0x5e, 0x31, 0xfe, 0x9b, 0xc8, 0x24, 0xa4, 0x0d, 0x78, 0x55, 0x76, 0xa0,
	0x61, 0x66, 0x7b, 0x85, 0x6e, 0x7f, 0xcb, 0xb7, 0x30, 0x6f, 0x8e, 0x59, 0xf9, 0x0e, 0x8d, 0x05,
	0x8d, 0x3a, 0xea, 0x12, 0x0c, 0x06, 0x18, 0x22, 0x68, 0xeb, 0x3d, 0xb2, 0xa3, 0xb0, 0xc0, 0xec,
	0xd7, 0x61, 0xe7, 0x62, 0xb1, 0xaa, 0x3b, 0x6f, 0xab, 0xba, 0x88, 0xd0, 0x0e, 0x27, 0x7d, 0x3f,
	0x24, 0x25, 0x64, 0x1c, 0x53, 0x2a, 0xcf, 0x34, 0xd8, 0x67, 0x4a, 0x47, 0x5f, 0x18, 0x0e, 0x25,
	0x6a, 0xa1, 0xa4, 0x1d, 0xfa, 0xd7, 0x10, 0x06, 0x3f, 0x5d, 0x9d, 0x60, 0xb1, 0xe6, 0xee, 0x72,
	0xcd, 0xdd, 0xf7, 0x35, 0x77, 0x5f, 0x36, 0xdc, 0x59, 0x6e, 0xb8, 0xf3, 0xba, 0xe1, 0xce, 0xd3,
	0xe5, 0x1f, 0xbd, 0x33, 0x51, 0xe4, 0x94, 0x85, 0xd4, 0xdf, 0xcb, 0x6e, 0x7a, 0xfe, 0x11, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x62, 0xab, 0xb0, 0xc1, 0x01, 0x00, 0x00,
}

func (m *SimpleCursorPaginationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleCursorPaginationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleCursorPaginationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimpleCursorPaginationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleCursorPaginationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleCursorPaginationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextKey) > 0 {
		i -= len(m.NextKey)
		copy(dAtA[i:], m.NextKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListeningCoefficient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListeningCoefficient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListeningCoefficient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Coefficient.Size()
		i -= size
		if _, err := m.Coefficient.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SimpleCursorPaginationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovTypes(uint64(m.Limit))
	}
	return n
}

func (m *SimpleCursorPaginationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NextKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ListeningCoefficient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coefficient.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SimpleCursorPaginationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SimpleCursorPaginationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleCursorPaginationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SimpleCursorPaginationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SimpleCursorPaginationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleCursorPaginationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextKey = append(m.NextKey[:0], dAtA[iNdEx:postIndex]...)
			if m.NextKey == nil {
				m.NextKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ListeningCoefficient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ListeningCoefficient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListeningCoefficient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coefficient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coefficient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
