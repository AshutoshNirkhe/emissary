// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v3/certs.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Proto representation of certificate details. Admin endpoint uses this wrapper for `/certs` to
// display certificate information. See :ref:`/certs <operations_admin_interface_certs>` for more
// information.
type Certificates struct {
	// List of certificates known to an Envoy.
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{0}
}
func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return m.Size()
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	// Details of CA certificate.
	CaCert []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Details of Certificate Chain
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{1}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

// [#next-free-field: 7]
type CertificateDetails struct {
	// Path of the certificate.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Certificate Serial Number.
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// List of Subject Alternate names.
	SubjectAltNames []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	// Minimum of days until expiration of certificate and it's chain.
	DaysUntilExpiration uint64 `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	// Indicates the time from which the certificate is valid.
	ValidFrom *types.Timestamp `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	// Indicates the time at which the certificate expires.
	ExpirationTime       *types.Timestamp `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{2}
}
func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return m.Size()
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *types.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *types.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Subject Alternate Name.
	//
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	//	*SubjectAlternateName_IpAddress
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{3}
}
func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return m.Size()
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof" json:"dns,omitempty"`
}
type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
}
type SubjectAlternateName_IpAddress struct {
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3,oneof" json:"ip_address,omitempty"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name()       {}
func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name()       {}
func (*SubjectAlternateName_IpAddress) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *SubjectAlternateName) GetIpAddress() string {
	if x, ok := m.GetName().(*SubjectAlternateName_IpAddress); ok {
		return x.IpAddress
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
		(*SubjectAlternateName_IpAddress)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v3.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v3.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v3.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v3.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v3/certs.proto", fileDescriptor_51228a64c985b2dc) }

var fileDescriptor_51228a64c985b2dc = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x8e, 0x12, 0x4d,
	0x14, 0xfd, 0x6a, 0xe0, 0xc3, 0x70, 0x07, 0x67, 0xb4, 0xd4, 0xa4, 0x83, 0x0a, 0x4c, 0x6b, 0xc6,
	0x8e, 0x89, 0xdd, 0x09, 0xac, 0xc4, 0x85, 0x02, 0x6a, 0x5c, 0x4d, 0x48, 0xab, 0xeb, 0x4e, 0xd1,
	0x5d, 0x40, 0x99, 0xee, 0xaa, 0x4e, 0x55, 0x35, 0x81, 0x9d, 0x71, 0xe5, 0x23, 0x18, 0xdf, 0xc3,
	0x8d, 0x7b, 0x93, 0x59, 0xea, 0x1b, 0x18, 0x9e, 0xc4, 0x54, 0x37, 0x33, 0x03, 0x33, 0x44, 0xdc,
	0x55, 0x9d, 0x73, 0xcf, 0xfd, 0x39, 0xf7, 0x42, 0x9d, 0xf2, 0x99, 0x58, 0x78, 0x24, 0x4a, 0x18,
	0xf7, 0x66, 0x1d, 0x2f, 0xa4, 0x52, 0x2b, 0x37, 0x95, 0x42, 0x0b, 0x7c, 0x90, 0x73, 0x6e, 0xce,
	0xb9, 0xb3, 0x4e, 0xbd, 0x39, 0x11, 0x62, 0x12, 0x53, 0x2f, 0x67, 0x47, 0xd9, 0xd8, 0xd3, 0x2c,
	0xa1, 0x4a, 0x93, 0x24, 0x2d, 0x04, 0xf5, 0xfb, 0x59, 0x94, 0x12, 0x8f, 0x70, 0x2e, 0x34, 0xd1,
	0x4c, 0x70, 0xe5, 0x29, 0x4d, 0x74, 0xb6, 0xca, 0x57, 0x3f, 0xba, 0x42, 0xcf, 0xa8, 0x54, 0x4c,
	0x70, 0xc6, 0x27, 0x45, 0x88, 0x3d, 0x87, 0xda, 0x80, 0x4a, 0xcd, 0xc6, 0x2c, 0x24, 0x9a, 0x2a,
	0xfc, 0x1c, 0x6a, 0xe1, 0xda, 0xdf, 0x42, 0xad, 0x92, 0xb3, 0xdf, 0xbe, 0xeb, 0x6e, 0x76, 0xe6,
	0xae, 0x69, 0xfc, 0x0d, 0x41, 0xf7, 0xd1, 0xd7, 0x1f, 0x9f, 0x1b, 0x36, 0xb4, 0x36, 0x04, 0x6d,
	0x12, 0xa7, 0x53, 0xb2, 0xae, 0x52, 0xf6, 0x37, 0x04, 0xfb, 0x6b, 0x00, 0x7e, 0x06, 0xd7, 0x42,
	0x12, 0x98, 0x5c, 0xab, 0xa2, 0xf6, 0x5f, 0x8a, 0xbe, 0xa4, 0x9a, 0xb0, 0x58, 0xf9, 0x95, 0x90,
	0x18, 0x14, 0xf7, 0x00, 0x8c, 0x32, 0x08, 0xa7, 0x84, 0x71, 0x6b, 0xef, 0x9f, 0xf5, 0x55, 0xa3,
	0x1a, 0x18, 0x51, 0xf7, 0xd8, 0x34, 0x7e, 0x04, 0xcd, 0x1d, 0x8d, 0xdb, 0x9f, 0x4a, 0x80, 0xaf,
	0x66, 0xc2, 0x18, 0xca, 0x29, 0xd1, 0x53, 0x0b, 0xb5, 0x90, 0x53, 0xf5, 0xf3, 0x37, 0x7e, 0x00,
	0xd7, 0x15, 0x95, 0x8c, 0xc4, 0x01, 0xcf, 0x92, 0x11, 0x95, 0xd6, 0x5e, 0x4e, 0xd6, 0x0a, 0xf0,
	0x24, 0xc7, 0xf0, 0x10, 0x6e, 0xaa, 0x6c, 0xf4, 0x81, 0x86, 0x3a, 0x20, 0xb1, 0x0e, 0x38, 0x49,
	0xa8, 0xb2, 0x4a, 0xf9, 0x04, 0x0f, 0x2f, 0x4f, 0xf0, 0xb6, 0x08, 0xec, 0xc5, 0x9a, 0x4a, 0x4e,
	0x34, 0x3d, 0x21, 0x09, 0xf5, 0x0f, 0xd5, 0x39, 0x6a, 0xfe, 0x0a, 0xb7, 0xe1, 0x4e, 0x44, 0x16,
	0x2a, 0xc8, 0xb8, 0x66, 0x71, 0x40, 0xe7, 0x29, 0x93, 0xf9, 0xfa, 0xad, 0x72, 0x0b, 0x39, 0x65,
	0xff, 0x96, 0x21, 0xdf, 0x1b, 0xee, 0xd5, 0x39, 0x85, 0x9f, 0x02, 0xcc, 0x48, 0xcc, 0xa2, 0x60,
	0x2c, 0x45, 0x62, 0xfd, 0xdf, 0x42, 0xce, 0x7e, 0xbb, 0xee, 0x16, 0xf7, 0xe7, 0x9e, 0xdd, 0x9f,
	0xfb, 0xee, 0xec, 0xfe, 0xfc, 0x6a, 0x1e, 0xfd, 0x5a, 0x8a, 0x04, 0x0f, 0xe0, 0xf0, 0xa2, 0x46,
	0x60, 0x4e, 0xd4, 0xaa, 0xec, 0xd4, 0x1f, 0x5c, 0x48, 0x0c, 0xd8, 0x7d, 0x62, 0xdc, 0x77, 0xe0,
	0x78, 0x87, 0xfb, 0x2b, 0xb7, 0xed, 0x2f, 0x08, 0x6e, 0x6f, 0x33, 0x03, 0x63, 0x28, 0x45, 0x5c,
	0x15, 0x5b, 0x78, 0xf3, 0x9f, 0x6f, 0x3e, 0x06, 0xcb, 0x24, 0x2b, 0xcc, 0x37, 0x58, 0x26, 0x19,
	0x6e, 0x02, 0xb0, 0x34, 0x20, 0x51, 0x24, 0xa9, 0x32, 0x76, 0x17, 0x54, 0x95, 0xa5, 0xbd, 0x02,
	0xea, 0x7a, 0xa6, 0xa1, 0xc7, 0xe0, 0x6c, 0x6b, 0x68, 0x5b, 0xe5, 0x7e, 0x05, 0xca, 0x66, 0x77,
	0xfd, 0x17, 0xa7, 0xcb, 0x06, 0xfa, 0xb9, 0x6c, 0xa0, 0xdf, 0xcb, 0x06, 0xfa, 0xfe, 0xf1, 0xf4,
	0x57, 0x65, 0xef, 0x06, 0x82, 0x7b, 0x4c, 0x14, 0xcb, 0x4c, 0xa5, 0x98, 0x2f, 0x2e, 0xed, 0xb5,
	0x0f, 0x66, 0x44, 0x35, 0x34, 0x36, 0x0d, 0xd1, 0xa8, 0x92, 0xfb, 0xd5, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0xdf, 0xee, 0x52, 0x2b, 0x04, 0x00, 0x00,
}

func (m *Certificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCerts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CertChain) > 0 {
		for iNdEx := len(m.CertChain) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CertChain[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCerts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CaCert) > 0 {
		for iNdEx := len(m.CaCert) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CaCert[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCerts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CertificateDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ExpirationTime != nil {
		{
			size, err := m.ExpirationTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCerts(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.ValidFrom != nil {
		{
			size, err := m.ValidFrom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCerts(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.DaysUntilExpiration != 0 {
		i = encodeVarintCerts(dAtA, i, uint64(m.DaysUntilExpiration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SubjectAltNames) > 0 {
		for iNdEx := len(m.SubjectAltNames) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubjectAltNames[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCerts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintCerts(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintCerts(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubjectAlternateName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubjectAlternateName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubjectAlternateName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Name != nil {
		{
			size := m.Name.Size()
			i -= size
			if _, err := m.Name.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubjectAlternateName_Dns) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubjectAlternateName_Dns) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Dns)
	copy(dAtA[i:], m.Dns)
	i = encodeVarintCerts(dAtA, i, uint64(len(m.Dns)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *SubjectAlternateName_Uri) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubjectAlternateName_Uri) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Uri)
	copy(dAtA[i:], m.Uri)
	i = encodeVarintCerts(dAtA, i, uint64(len(m.Uri)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *SubjectAlternateName_IpAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubjectAlternateName_IpAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.IpAddress)
	copy(dAtA[i:], m.IpAddress)
	i = encodeVarintCerts(dAtA, i, uint64(len(m.IpAddress)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func encodeVarintCerts(dAtA []byte, offset int, v uint64) int {
	offset -= sovCerts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Certificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CaCert) > 0 {
		for _, e := range m.CaCert {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if len(m.CertChain) > 0 {
		for _, e := range m.CertChain {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CertificateDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCerts(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovCerts(uint64(l))
	}
	if len(m.SubjectAltNames) > 0 {
		for _, e := range m.SubjectAltNames {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.DaysUntilExpiration != 0 {
		n += 1 + sovCerts(uint64(m.DaysUntilExpiration))
	}
	if m.ValidFrom != nil {
		l = m.ValidFrom.Size()
		n += 1 + l + sovCerts(uint64(l))
	}
	if m.ExpirationTime != nil {
		l = m.ExpirationTime.Size()
		n += 1 + l + sovCerts(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubjectAlternateName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != nil {
		n += m.Name.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubjectAlternateName_Dns) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dns)
	n += 1 + l + sovCerts(uint64(l))
	return n
}
func (m *SubjectAlternateName_Uri) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	n += 1 + l + sovCerts(uint64(l))
	return n
}
func (m *SubjectAlternateName_IpAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IpAddress)
	n += 1 + l + sovCerts(uint64(l))
	return n
}

func sovCerts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCerts(x uint64) (n int) {
	return sovCerts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: Certificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, &Certificate{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCerts
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
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaCert", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaCert = append(m.CaCert, &CertificateDetails{})
			if err := m.CaCert[len(m.CaCert)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertChain = append(m.CertChain, &CertificateDetails{})
			if err := m.CertChain[len(m.CertChain)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCerts
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
func (m *CertificateDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: CertificateDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAltNames", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectAltNames = append(m.SubjectAltNames, &SubjectAlternateName{})
			if err := m.SubjectAltNames[len(m.SubjectAltNames)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaysUntilExpiration", wireType)
			}
			m.DaysUntilExpiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DaysUntilExpiration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidFrom == nil {
				m.ValidFrom = &types.Timestamp{}
			}
			if err := m.ValidFrom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = &types.Timestamp{}
			}
			if err := m.ExpirationTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCerts
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
func (m *SubjectAlternateName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: SubjectAlternateName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubjectAlternateName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = &SubjectAlternateName_Dns{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = &SubjectAlternateName_Uri{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCerts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = &SubjectAlternateName_IpAddress{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCerts
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
func skipCerts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCerts
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
					return 0, ErrIntOverflowCerts
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
					return 0, ErrIntOverflowCerts
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
				return 0, ErrInvalidLengthCerts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCerts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCerts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCerts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCerts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCerts = fmt.Errorf("proto: unexpected end of group")
)