// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/beacon/db/attestation_container.proto

package db

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	v1alpha1 "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AttestationContainer struct {
	Data                 *v1alpha1.AttestationData             `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	SignaturePairs       []*AttestationContainer_SignaturePair `protobuf:"bytes,2,rep,name=signature_pairs,json=signaturePairs,proto3" json:"signature_pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AttestationContainer) Reset()         { *m = AttestationContainer{} }
func (m *AttestationContainer) String() string { return proto.CompactTextString(m) }
func (*AttestationContainer) ProtoMessage()    {}
func (*AttestationContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_29679516eb4218c9, []int{0}
}
func (m *AttestationContainer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttestationContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttestationContainer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationContainer.Merge(m, src)
}
func (m *AttestationContainer) XXX_Size() int {
	return m.Size()
}
func (m *AttestationContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationContainer.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationContainer proto.InternalMessageInfo

func (m *AttestationContainer) GetData() *v1alpha1.AttestationData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AttestationContainer) GetSignaturePairs() []*AttestationContainer_SignaturePair {
	if m != nil {
		return m.SignaturePairs
	}
	return nil
}

type AttestationContainer_SignaturePair struct {
	AggregationBits      github_com_prysmaticlabs_go_bitfield.Bitlist `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3,casttype=github.com/prysmaticlabs/go-bitfield.Bitlist" json:"aggregation_bits,omitempty"`
	Signature            []byte                                       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *AttestationContainer_SignaturePair) Reset()         { *m = AttestationContainer_SignaturePair{} }
func (m *AttestationContainer_SignaturePair) String() string { return proto.CompactTextString(m) }
func (*AttestationContainer_SignaturePair) ProtoMessage()    {}
func (*AttestationContainer_SignaturePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_29679516eb4218c9, []int{0, 0}
}
func (m *AttestationContainer_SignaturePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttestationContainer_SignaturePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttestationContainer_SignaturePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationContainer_SignaturePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationContainer_SignaturePair.Merge(m, src)
}
func (m *AttestationContainer_SignaturePair) XXX_Size() int {
	return m.Size()
}
func (m *AttestationContainer_SignaturePair) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationContainer_SignaturePair.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationContainer_SignaturePair proto.InternalMessageInfo

func (m *AttestationContainer_SignaturePair) GetAggregationBits() github_com_prysmaticlabs_go_bitfield.Bitlist {
	if m != nil {
		return m.AggregationBits
	}
	return nil
}

func (m *AttestationContainer_SignaturePair) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestationContainer)(nil), "prysm.beacon.db.AttestationContainer")
	proto.RegisterType((*AttestationContainer_SignaturePair)(nil), "prysm.beacon.db.AttestationContainer.SignaturePair")
}

func init() {
	proto.RegisterFile("proto/beacon/db/attestation_container.proto", fileDescriptor_29679516eb4218c9)
}

var fileDescriptor_29679516eb4218c9 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcf, 0x4a, 0xf3, 0x40,
	0x10, 0x27, 0xfd, 0x3e, 0x04, 0x53, 0xb5, 0x12, 0x3c, 0x94, 0x22, 0xb5, 0x88, 0x48, 0x41, 0xbb,
	0x6b, 0xdb, 0x9b, 0x78, 0x31, 0xfa, 0x00, 0x12, 0x6f, 0x2a, 0x94, 0xd9, 0x64, 0xbb, 0x19, 0x48,
	0xb3, 0x61, 0x77, 0x22, 0xf8, 0x0a, 0xbe, 0x93, 0x77, 0x8f, 0x3e, 0x81, 0x48, 0x1f, 0xc3, 0x93,
	0x74, 0x43, 0x6d, 0x2c, 0x7a, 0x9b, 0xf9, 0xf1, 0xfb, 0xb7, 0xb3, 0xfe, 0x49, 0x61, 0x34, 0x69,
	0x2e, 0x24, 0xc4, 0x3a, 0xe7, 0x89, 0xe0, 0x40, 0x24, 0x2d, 0x01, 0xa1, 0xce, 0x27, 0xb1, 0xce,
	0x09, 0x30, 0x97, 0x86, 0x39, 0x56, 0xd0, 0x2a, 0xcc, 0x93, 0x9d, 0xb1, 0x8a, 0xcc, 0x12, 0xd1,
	0x39, 0xaa, 0xd4, 0x92, 0x52, 0xfe, 0x38, 0x84, 0xac, 0x48, 0x61, 0x58, 0x37, 0xa8, 0x64, 0x9d,
	0x81, 0x42, 0x4a, 0x4b, 0xc1, 0x62, 0x3d, 0xe3, 0x4a, 0x2b, 0xcd, 0x1d, 0x2c, 0xca, 0xa9, 0xdb,
	0x2a, 0x8b, 0xc5, 0x54, 0xd1, 0x0f, 0x5f, 0x1a, 0xfe, 0xde, 0xe5, 0xca, 0xe4, 0x6a, 0x59, 0x22,
	0x38, 0xf7, 0xff, 0x27, 0x40, 0xd0, 0xf6, 0x7a, 0x5e, 0xbf, 0x39, 0x3a, 0x66, 0x92, 0x52, 0x69,
	0x64, 0x39, 0x5b, 0x0c, 0x6c, 0x99, 0xcf, 0x6a, 0xd2, 0x6b, 0x20, 0x88, 0x9c, 0x26, 0x78, 0xf0,
	0x5b, 0x16, 0x55, 0x0e, 0x54, 0x1a, 0x39, 0x29, 0x00, 0x8d, 0x6d, 0x37, 0x7a, 0xff, 0xfa, 0xcd,
	0xd1, 0x98, 0xad, 0x3d, 0x8a, 0xfd, 0x96, 0xcd, 0x6e, 0x97, 0xe2, 0x1b, 0x40, 0x13, 0xed, 0xd8,
	0xfa, 0x6a, 0x3b, 0xcf, 0x9e, 0xbf, 0xfd, 0x83, 0x11, 0xdc, 0xfb, 0xbb, 0xa0, 0x94, 0x91, 0xaa,
	0xba, 0xa4, 0x40, 0xb2, 0xae, 0xf7, 0x56, 0x78, 0xf6, 0xf9, 0x7e, 0x70, 0x5a, 0xbb, 0x88, 0x8b,
	0x07, 0xc2, 0x38, 0x03, 0x61, 0xb9, 0xd2, 0x03, 0x81, 0x34, 0x45, 0x99, 0x25, 0x2c, 0x44, 0xca,
	0xd0, 0x52, 0xd4, 0xaa, 0x39, 0x85, 0x48, 0x36, 0xd8, 0xf7, 0x37, 0xbf, 0x0b, 0xb4, 0x1b, 0x0b,
	0xd7, 0x68, 0x05, 0x84, 0x17, 0xaf, 0xf3, 0xae, 0xf7, 0x36, 0xef, 0x7a, 0x1f, 0xf3, 0xae, 0x77,
	0xc7, 0xfe, 0x8c, 0x72, 0x1b, 0x5f, 0xfb, 0x7f, 0xb1, 0xe1, 0x80, 0xf1, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x93, 0x3e, 0x32, 0x19, 0x02, 0x00, 0x00,
}

func (m *AttestationContainer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestationContainer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAttestationContainer(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.SignaturePairs) > 0 {
		for _, msg := range m.SignaturePairs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintAttestationContainer(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AttestationContainer_SignaturePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestationContainer_SignaturePair) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AggregationBits) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAttestationContainer(dAtA, i, uint64(len(m.AggregationBits)))
		i += copy(dAtA[i:], m.AggregationBits)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAttestationContainer(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAttestationContainer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AttestationContainer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovAttestationContainer(uint64(l))
	}
	if len(m.SignaturePairs) > 0 {
		for _, e := range m.SignaturePairs {
			l = e.Size()
			n += 1 + l + sovAttestationContainer(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AttestationContainer_SignaturePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AggregationBits)
	if l > 0 {
		n += 1 + l + sovAttestationContainer(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovAttestationContainer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAttestationContainer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAttestationContainer(x uint64) (n int) {
	return sovAttestationContainer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AttestationContainer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestationContainer
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
			return fmt.Errorf("proto: AttestationContainer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttestationContainer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestationContainer
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
				return ErrInvalidLengthAttestationContainer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &v1alpha1.AttestationData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignaturePairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestationContainer
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
				return ErrInvalidLengthAttestationContainer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignaturePairs = append(m.SignaturePairs, &AttestationContainer_SignaturePair{})
			if err := m.SignaturePairs[len(m.SignaturePairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestationContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAttestationContainer
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
func (m *AttestationContainer_SignaturePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestationContainer
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
			return fmt.Errorf("proto: SignaturePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignaturePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationBits", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestationContainer
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
				return ErrInvalidLengthAttestationContainer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregationBits = append(m.AggregationBits[:0], dAtA[iNdEx:postIndex]...)
			if m.AggregationBits == nil {
				m.AggregationBits = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestationContainer
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
				return ErrInvalidLengthAttestationContainer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestationContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAttestationContainer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAttestationContainer
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
func skipAttestationContainer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestationContainer
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
					return 0, ErrIntOverflowAttestationContainer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAttestationContainer
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
				return 0, ErrInvalidLengthAttestationContainer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAttestationContainer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAttestationContainer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAttestationContainer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAttestationContainer
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAttestationContainer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestationContainer   = fmt.Errorf("proto: integer overflow")
)
