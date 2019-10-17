// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/string.proto

package envoy_type_matcher

import (
	fmt "fmt"
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

// Specifies the way to match a string.
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_Regex
	//	*StringMatcher_SafeRegex
	MatchPattern         isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{0}
}
func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}
type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}
type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}
type StringMatcher_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}
type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern()     {}
func (*StringMatcher_Prefix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_Suffix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_Regex) isStringMatcher_MatchPattern()     {}
func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

// Deprecated: Do not use.
func (m *StringMatcher) GetRegex() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_Regex)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc62c75a0f154e3, []int{1}
}
func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.ListStringMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/string.proto", fileDescriptor_1dc62c75a0f154e3) }

var fileDescriptor_1dc62c75a0f154e3 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4f, 0xfa, 0x40,
	0x14, 0xc7, 0x79, 0x2d, 0x25, 0xf0, 0xf8, 0x91, 0xfc, 0xbc, 0x01, 0x1b, 0x86, 0x7a, 0xa2, 0x43,
	0xa7, 0x36, 0xc1, 0xd9, 0xa5, 0x13, 0x26, 0x92, 0x90, 0xba, 0x4b, 0x4e, 0x3c, 0xb0, 0x89, 0xd2,
	0xe6, 0x7a, 0x92, 0xb2, 0x39, 0x3b, 0xfa, 0xd7, 0x18, 0x27, 0x46, 0x47, 0xff, 0x04, 0xd3, 0x8d,
	0x7f, 0xc2, 0x98, 0xbb, 0xab, 0x26, 0xc4, 0x6e, 0x6d, 0xbe, 0x9f, 0xef, 0xe7, 0xbd, 0xf6, 0xe1,
	0x11, 0x5f, 0xad, 0xd3, 0x4d, 0x28, 0x37, 0x19, 0x0f, 0x1f, 0x98, 0x9c, 0xdf, 0x71, 0x11, 0xe6,
	0x52, 0x24, 0xab, 0x65, 0x90, 0x89, 0x54, 0xa6, 0x84, 0x68, 0x20, 0x50, 0x40, 0x50, 0x01, 0x03,
	0xaf, 0xa6, 0x24, 0xf8, 0x92, 0x17, 0xa6, 0x33, 0x38, 0x5c, 0xb3, 0xfb, 0xe4, 0x96, 0x49, 0x1e,
	0xfe, 0x3c, 0x98, 0x60, 0xf8, 0x05, 0xd8, 0xbb, 0xd2, 0xf6, 0x89, 0xa9, 0x91, 0x3e, 0x3a, 0xbc,
	0x60, 0x73, 0xe9, 0x02, 0x05, 0xbf, 0x33, 0x6e, 0xc4, 0xe6, 0x95, 0x9c, 0x60, 0x2b, 0x13, 0x7c,
	0x91, 0x14, 0xae, 0xa5, 0x82, 0xa8, 0xf3, 0xb6, 0xdb, 0xda, 0x4d, 0x61, 0x51, 0x18, 0x37, 0xe2,
	0x2a, 0x52, 0x50, 0xfe, 0xb8, 0x50, 0x90, 0x5d, 0x03, 0x99, 0x88, 0x9c, 0xa2, 0xa3, 0x77, 0x73,
	0x9b, 0x9a, 0xf9, 0xa7, 0x18, 0x47, 0xd8, 0xfe, 0x53, 0xdb, 0x55, 0x98, 0x09, 0xc9, 0x04, 0x31,
	0x67, 0x0b, 0x3e, 0x33, 0xa8, 0x43, 0xc1, 0xef, 0x8e, 0x68, 0xf0, 0xf7, 0xdb, 0x83, 0x58, 0x01,
	0xd5, 0xf6, 0x11, 0x6a, 0xd9, 0x33, 0x58, 0xff, 0x95, 0xaa, 0xa3, 0x0c, 0x3a, 0x8f, 0xfa, 0xd8,
	0xd3, 0x85, 0x59, 0xc6, 0xa4, 0xe4, 0x62, 0x45, 0x9c, 0xd7, 0xdd, 0xd6, 0x86, 0xe1, 0x35, 0x1e,
	0x5c, 0x26, 0xb9, 0xdc, 0xff, 0x07, 0x17, 0xd8, 0xae, 0xb0, 0xdc, 0x05, 0x6a, 0xfb, 0xdd, 0xd1,
	0x71, 0xdd, 0xe4, 0xbd, 0x52, 0x35, 0xfa, 0x05, 0xac, 0x36, 0xc4, 0xbf, 0xf5, 0xe8, 0xfc, 0xbd,
	0xf4, 0xe0, 0xa3, 0xf4, 0xe0, 0xb3, 0xf4, 0x00, 0x69, 0x92, 0x1a, 0x51, 0x26, 0xd2, 0x62, 0x53,
	0xe3, 0x8c, 0xba, 0x46, 0x3a, 0x55, 0xd7, 0x99, 0xc2, 0x4d, 0x4b, 0x9f, 0xe9, 0xec, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x82, 0x4e, 0x93, 0xd0, 0x16, 0x02, 0x00, 0x00,
}

func (m *StringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MatchPattern != nil {
		{
			size := m.MatchPattern.Size()
			i -= size
			if _, err := m.MatchPattern.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *StringMatcher_Exact) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *StringMatcher_Exact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Exact)
	copy(dAtA[i:], m.Exact)
	i = encodeVarintString(dAtA, i, uint64(len(m.Exact)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Prefix) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *StringMatcher_Prefix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Prefix)
	copy(dAtA[i:], m.Prefix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Prefix)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Suffix) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *StringMatcher_Suffix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Suffix)
	copy(dAtA[i:], m.Suffix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Suffix)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Regex) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *StringMatcher_Regex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Regex)
	copy(dAtA[i:], m.Regex)
	i = encodeVarintString(dAtA, i, uint64(len(m.Regex)))
	i--
	dAtA[i] = 0x22
	return len(dAtA) - i, nil
}
func (m *StringMatcher_SafeRegex) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *StringMatcher_SafeRegex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SafeRegex != nil {
		{
			size, err := m.SafeRegex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintString(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *ListStringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListStringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListStringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Patterns) > 0 {
		for iNdEx := len(m.Patterns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Patterns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintString(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintString(dAtA []byte, offset int, v uint64) int {
	offset -= sovString(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StringMatcher_Exact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Exact)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Prefix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Suffix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Suffix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Regex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Regex)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_SafeRegex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SafeRegex != nil {
		l = m.SafeRegex.Size()
		n += 1 + l + sovString(uint64(l))
	}
	return n
}
func (m *ListStringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Patterns) > 0 {
		for _, e := range m.Patterns {
			l = e.Size()
			n += 1 + l + sovString(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovString(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozString(x uint64) (n int) {
	return sovString(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: StringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Exact{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Prefix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Suffix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Regex{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeRegex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegexMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &StringMatcher_SafeRegex{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func (m *ListStringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: ListStringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListStringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patterns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patterns = append(m.Patterns, &StringMatcher{})
			if err := m.Patterns[len(m.Patterns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func skipString(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
				return 0, ErrInvalidLengthString
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthString
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowString
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
				next, err := skipString(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthString
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
	ErrInvalidLengthString = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowString   = fmt.Errorf("proto: integer overflow")
)
