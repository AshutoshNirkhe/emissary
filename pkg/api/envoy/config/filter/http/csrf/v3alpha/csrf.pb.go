// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v3alpha/csrf.proto

package envoy_config_filter_http_csrf_v3alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v3alpha/core"
	matcher "github.com/datawire/ambassador/pkg/api/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

// CSRF filter config.
type CsrfPolicy struct {
	// Specifies if CSRF is enabled.
	//
	// More information on how this can be controlled via runtime can be found
	// :ref:`here <csrf-runtime>`.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.FractionalPercent.DenominatorType>`.
	FilterEnabled *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies that CSRF policies will be evaluated and tracked, but not enforced.
	// This is intended to be used when filter_enabled is off.
	//
	// More information on how this can be controlled via runtime can be found
	// :ref:`here <csrf-runtime>`.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.FractionalPercent.DenominatorType>`.
	ShadowEnabled *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	// Specifies additional source origins that will be allowed in addition to
	// the destination origin.
	//
	// More information on how this can be configured via runtime can be found
	// :ref:`here <csrf-configuration>`.
	AdditionalOrigins    []*matcher.StringMatcher `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ea04258cfdb2a5, []int{0}
}
func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return m.Size()
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetAdditionalOrigins() []*matcher.StringMatcher {
	if m != nil {
		return m.AdditionalOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v3alpha.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v3alpha/csrf.proto", fileDescriptor_a0ea04258cfdb2a5)
}

var fileDescriptor_a0ea04258cfdb2a5 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xe5, 0x56, 0x30, 0xb8, 0x6a, 0x81, 0x2c, 0x54, 0x1d, 0x4a, 0x8b, 0x84, 0xd4, 0xc9,
	0xae, 0xda, 0x1b, 0x14, 0xc1, 0x86, 0x88, 0x02, 0x12, 0x63, 0xe5, 0x3a, 0x4e, 0xf3, 0x4b, 0xa9,
	0x6d, 0x39, 0xa6, 0x90, 0x2b, 0x70, 0x06, 0x4e, 0xc2, 0xc4, 0xc8, 0xc8, 0x11, 0x50, 0x37, 0x6e,
	0x81, 0x92, 0x3f, 0x25, 0x2b, 0x62, 0x8b, 0x95, 0xf7, 0x7d, 0x7e, 0x7a, 0xa6, 0x53, 0xa5, 0xb7,
	0xa6, 0xe0, 0xd2, 0xe8, 0x04, 0xd6, 0x3c, 0x81, 0xcc, 0x2b, 0xc7, 0x53, 0xef, 0x2d, 0x97, 0xb9,
	0x4b, 0xf8, 0x76, 0x2e, 0x32, 0x9b, 0x8a, 0xea, 0xc0, 0xac, 0x33, 0xde, 0x04, 0x17, 0x15, 0xc1,
	0x90, 0x60, 0x48, 0xb0, 0x92, 0x60, 0x55, 0xa8, 0x26, 0x06, 0x63, 0x14, 0x0b, 0x0b, 0x8d, 0xc4,
	0x38, 0xc5, 0x57, 0x22, 0x57, 0x68, 0x1a, 0x9c, 0x61, 0xc4, 0x17, 0x56, 0xf1, 0x8d, 0xf0, 0x32,
	0x55, 0x8e, 0xe7, 0xde, 0x81, 0x5e, 0xd7, 0x81, 0xd3, 0xad, 0xc8, 0x20, 0x16, 0x5e, 0xf1, 0xfd,
	0x07, 0xfe, 0x38, 0x7f, 0x6d, 0x51, 0x7a, 0x99, 0xbb, 0x24, 0x34, 0x19, 0xc8, 0x22, 0x90, 0xb4,
	0x87, 0x3d, 0x96, 0x4a, 0x8b, 0x55, 0xa6, 0xe2, 0x3e, 0x19, 0x91, 0x49, 0x67, 0x36, 0x65, 0xd8,
	0x55, 0x58, 0xd8, 0xf7, 0x62, 0x65, 0x09, 0x16, 0x3d, 0x6a, 0x0f, 0x1b, 0x75, 0xed, 0x84, 0xf4,
	0x60, 0xb4, 0xc8, 0x42, 0xe5, 0xa4, 0xd2, 0x7e, 0x41, 0xdf, 0xbe, 0xdf, 0xdb, 0x07, 0x2f, 0xa4,
	0x75, 0x4c, 0xa2, 0x2e, 0x3a, 0xaf, 0x50, 0x19, 0x3c, 0xd0, 0x5e, 0x9e, 0x8a, 0xd8, 0x3c, 0xfd,
	0x5e, 0xd2, 0xfa, 0xdf, 0x25, 0x51, 0x17, 0x3d, 0x7b, 0x71, 0x48, 0x03, 0x11, 0xc7, 0x80, 0x99,
	0xa5, 0x71, 0xb0, 0x06, 0x9d, 0xf7, 0xdb, 0xa3, 0xf6, 0xa4, 0x33, 0x1b, 0xd7, 0xf2, 0x72, 0x23,
	0x56, 0x6f, 0xc4, 0xee, 0xaa, 0x8d, 0x6e, 0xf0, 0x14, 0x9d, 0x34, 0xf0, 0x2d, 0xb2, 0x8b, 0xfb,
	0x8f, 0xdd, 0x90, 0x7c, 0xee, 0x86, 0xe4, 0x6b, 0x37, 0x24, 0x74, 0x0e, 0x06, 0x2d, 0xd6, 0x99,
	0xe7, 0x82, 0xfd, 0xe9, 0xf9, 0x16, 0x47, 0xcd, 0xbc, 0x61, 0x39, 0x79, 0x48, 0x56, 0x87, 0xd5,
	0xf6, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xd5, 0x5f, 0xb3, 0x33, 0x02, 0x00, 0x00,
}

func (m *CsrfPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CsrfPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CsrfPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AdditionalOrigins) > 0 {
		for iNdEx := len(m.AdditionalOrigins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdditionalOrigins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCsrf(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ShadowEnabled != nil {
		{
			size, err := m.ShadowEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCsrf(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.FilterEnabled != nil {
		{
			size, err := m.FilterEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCsrf(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCsrf(dAtA []byte, offset int, v uint64) int {
	offset -= sovCsrf(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CsrfPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FilterEnabled != nil {
		l = m.FilterEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if m.ShadowEnabled != nil {
		l = m.ShadowEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if len(m.AdditionalOrigins) > 0 {
		for _, e := range m.AdditionalOrigins {
			l = e.Size()
			n += 1 + l + sovCsrf(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCsrf(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCsrf(x uint64) (n int) {
	return sovCsrf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CsrfPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsrf
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
			return fmt.Errorf("proto: CsrfPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CsrfPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FilterEnabled == nil {
				m.FilterEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.FilterEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowEnabled == nil {
				m.ShadowEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.ShadowEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalOrigins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalOrigins = append(m.AdditionalOrigins, &matcher.StringMatcher{})
			if err := m.AdditionalOrigins[len(m.AdditionalOrigins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCsrf
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCsrf
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
func skipCsrf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
				return 0, ErrInvalidLengthCsrf
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCsrf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCsrf
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
				next, err := skipCsrf(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCsrf
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
	ErrInvalidLengthCsrf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsrf   = fmt.Errorf("proto: integer overflow")
)
