// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prototypes/ctrecord.proto

package prototypes // import "github.com/linkai-io/am/protocservices/prototypes"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CTRecord struct {
	CertificateID        int64    `protobuf:"varint,1,opt,name=CertificateID,proto3" json:"CertificateID,omitempty"`
	InsertedTime         int64    `protobuf:"varint,2,opt,name=InsertedTime,proto3" json:"InsertedTime,omitempty"`
	CertHash             string   `protobuf:"bytes,3,opt,name=CertHash,proto3" json:"CertHash,omitempty"`
	SerialNumber         string   `protobuf:"bytes,4,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	NotBefore            int64    `protobuf:"varint,5,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	NotAfter             int64    `protobuf:"varint,6,opt,name=NotAfter,proto3" json:"NotAfter,omitempty"`
	Country              string   `protobuf:"bytes,7,opt,name=Country,proto3" json:"Country,omitempty"`
	Organization         string   `protobuf:"bytes,8,opt,name=Organization,proto3" json:"Organization,omitempty"`
	OrganizationalUnit   string   `protobuf:"bytes,9,opt,name=OrganizationalUnit,proto3" json:"OrganizationalUnit,omitempty"`
	CommonName           string   `protobuf:"bytes,10,opt,name=CommonName,proto3" json:"CommonName,omitempty"`
	VerifiedDNSNames     string   `protobuf:"bytes,11,opt,name=VerifiedDNSNames,proto3" json:"VerifiedDNSNames,omitempty"`
	UnverifiedDNSNames   string   `protobuf:"bytes,12,opt,name=UnverifiedDNSNames,proto3" json:"UnverifiedDNSNames,omitempty"`
	IPAddresses          string   `protobuf:"bytes,13,opt,name=IPAddresses,proto3" json:"IPAddresses,omitempty"`
	EmailAddresses       string   `protobuf:"bytes,14,opt,name=EmailAddresses,proto3" json:"EmailAddresses,omitempty"`
	ETLD                 string   `protobuf:"bytes,15,opt,name=ETLD,proto3" json:"ETLD,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CTRecord) Reset()         { *m = CTRecord{} }
func (m *CTRecord) String() string { return proto.CompactTextString(m) }
func (*CTRecord) ProtoMessage()    {}
func (*CTRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ctrecord_acaee2e3ab70fdfc, []int{0}
}
func (m *CTRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CTRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CTRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CTRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CTRecord.Merge(dst, src)
}
func (m *CTRecord) XXX_Size() int {
	return m.Size()
}
func (m *CTRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CTRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CTRecord proto.InternalMessageInfo

func (m *CTRecord) GetCertificateID() int64 {
	if m != nil {
		return m.CertificateID
	}
	return 0
}

func (m *CTRecord) GetInsertedTime() int64 {
	if m != nil {
		return m.InsertedTime
	}
	return 0
}

func (m *CTRecord) GetCertHash() string {
	if m != nil {
		return m.CertHash
	}
	return ""
}

func (m *CTRecord) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CTRecord) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *CTRecord) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *CTRecord) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *CTRecord) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *CTRecord) GetOrganizationalUnit() string {
	if m != nil {
		return m.OrganizationalUnit
	}
	return ""
}

func (m *CTRecord) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *CTRecord) GetVerifiedDNSNames() string {
	if m != nil {
		return m.VerifiedDNSNames
	}
	return ""
}

func (m *CTRecord) GetUnverifiedDNSNames() string {
	if m != nil {
		return m.UnverifiedDNSNames
	}
	return ""
}

func (m *CTRecord) GetIPAddresses() string {
	if m != nil {
		return m.IPAddresses
	}
	return ""
}

func (m *CTRecord) GetEmailAddresses() string {
	if m != nil {
		return m.EmailAddresses
	}
	return ""
}

func (m *CTRecord) GetETLD() string {
	if m != nil {
		return m.ETLD
	}
	return ""
}

type CTSubdomain struct {
	SubdomainID          int64    `protobuf:"varint,1,opt,name=SubdomainID,proto3" json:"SubdomainID,omitempty"`
	InsertedTime         int64    `protobuf:"varint,2,opt,name=InsertedTime,proto3" json:"InsertedTime,omitempty"`
	CommonName           string   `protobuf:"bytes,3,opt,name=CommonName,proto3" json:"CommonName,omitempty"`
	ETLD                 string   `protobuf:"bytes,4,opt,name=ETLD,proto3" json:"ETLD,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CTSubdomain) Reset()         { *m = CTSubdomain{} }
func (m *CTSubdomain) String() string { return proto.CompactTextString(m) }
func (*CTSubdomain) ProtoMessage()    {}
func (*CTSubdomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_ctrecord_acaee2e3ab70fdfc, []int{1}
}
func (m *CTSubdomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CTSubdomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CTSubdomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CTSubdomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CTSubdomain.Merge(dst, src)
}
func (m *CTSubdomain) XXX_Size() int {
	return m.Size()
}
func (m *CTSubdomain) XXX_DiscardUnknown() {
	xxx_messageInfo_CTSubdomain.DiscardUnknown(m)
}

var xxx_messageInfo_CTSubdomain proto.InternalMessageInfo

func (m *CTSubdomain) GetSubdomainID() int64 {
	if m != nil {
		return m.SubdomainID
	}
	return 0
}

func (m *CTSubdomain) GetInsertedTime() int64 {
	if m != nil {
		return m.InsertedTime
	}
	return 0
}

func (m *CTSubdomain) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *CTSubdomain) GetETLD() string {
	if m != nil {
		return m.ETLD
	}
	return ""
}

type CTETLD struct {
	EtldId               int32    `protobuf:"varint,1,opt,name=etld_id,json=etldId,proto3" json:"etld_id,omitempty"`
	Etld                 string   `protobuf:"bytes,2,opt,name=etld,proto3" json:"etld,omitempty"`
	QueryTimestamp       int64    `protobuf:"varint,3,opt,name=query_timestamp,json=queryTimestamp,proto3" json:"query_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CTETLD) Reset()         { *m = CTETLD{} }
func (m *CTETLD) String() string { return proto.CompactTextString(m) }
func (*CTETLD) ProtoMessage()    {}
func (*CTETLD) Descriptor() ([]byte, []int) {
	return fileDescriptor_ctrecord_acaee2e3ab70fdfc, []int{2}
}
func (m *CTETLD) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CTETLD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CTETLD.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CTETLD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CTETLD.Merge(dst, src)
}
func (m *CTETLD) XXX_Size() int {
	return m.Size()
}
func (m *CTETLD) XXX_DiscardUnknown() {
	xxx_messageInfo_CTETLD.DiscardUnknown(m)
}

var xxx_messageInfo_CTETLD proto.InternalMessageInfo

func (m *CTETLD) GetEtldId() int32 {
	if m != nil {
		return m.EtldId
	}
	return 0
}

func (m *CTETLD) GetEtld() string {
	if m != nil {
		return m.Etld
	}
	return ""
}

func (m *CTETLD) GetQueryTimestamp() int64 {
	if m != nil {
		return m.QueryTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*CTRecord)(nil), "CTRecord")
	proto.RegisterType((*CTSubdomain)(nil), "CTSubdomain")
	proto.RegisterType((*CTETLD)(nil), "CTETLD")
}
func (m *CTRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CTRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CertificateID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.CertificateID))
	}
	if m.InsertedTime != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.InsertedTime))
	}
	if len(m.CertHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.CertHash)))
		i += copy(dAtA[i:], m.CertHash)
	}
	if len(m.SerialNumber) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.SerialNumber)))
		i += copy(dAtA[i:], m.SerialNumber)
	}
	if m.NotBefore != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.NotBefore))
	}
	if m.NotAfter != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.NotAfter))
	}
	if len(m.Country) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.Country)))
		i += copy(dAtA[i:], m.Country)
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	if len(m.OrganizationalUnit) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.OrganizationalUnit)))
		i += copy(dAtA[i:], m.OrganizationalUnit)
	}
	if len(m.CommonName) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.CommonName)))
		i += copy(dAtA[i:], m.CommonName)
	}
	if len(m.VerifiedDNSNames) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.VerifiedDNSNames)))
		i += copy(dAtA[i:], m.VerifiedDNSNames)
	}
	if len(m.UnverifiedDNSNames) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.UnverifiedDNSNames)))
		i += copy(dAtA[i:], m.UnverifiedDNSNames)
	}
	if len(m.IPAddresses) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.IPAddresses)))
		i += copy(dAtA[i:], m.IPAddresses)
	}
	if len(m.EmailAddresses) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.EmailAddresses)))
		i += copy(dAtA[i:], m.EmailAddresses)
	}
	if len(m.ETLD) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.ETLD)))
		i += copy(dAtA[i:], m.ETLD)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CTSubdomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CTSubdomain) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SubdomainID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.SubdomainID))
	}
	if m.InsertedTime != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.InsertedTime))
	}
	if len(m.CommonName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.CommonName)))
		i += copy(dAtA[i:], m.CommonName)
	}
	if len(m.ETLD) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.ETLD)))
		i += copy(dAtA[i:], m.ETLD)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CTETLD) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CTETLD) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EtldId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.EtldId))
	}
	if len(m.Etld) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(len(m.Etld)))
		i += copy(dAtA[i:], m.Etld)
	}
	if m.QueryTimestamp != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCtrecord(dAtA, i, uint64(m.QueryTimestamp))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCtrecord(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CTRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CertificateID != 0 {
		n += 1 + sovCtrecord(uint64(m.CertificateID))
	}
	if m.InsertedTime != 0 {
		n += 1 + sovCtrecord(uint64(m.InsertedTime))
	}
	l = len(m.CertHash)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	if m.NotBefore != 0 {
		n += 1 + sovCtrecord(uint64(m.NotBefore))
	}
	if m.NotAfter != 0 {
		n += 1 + sovCtrecord(uint64(m.NotAfter))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.OrganizationalUnit)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.CommonName)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.VerifiedDNSNames)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.UnverifiedDNSNames)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.IPAddresses)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.EmailAddresses)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.ETLD)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CTSubdomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubdomainID != 0 {
		n += 1 + sovCtrecord(uint64(m.SubdomainID))
	}
	if m.InsertedTime != 0 {
		n += 1 + sovCtrecord(uint64(m.InsertedTime))
	}
	l = len(m.CommonName)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	l = len(m.ETLD)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CTETLD) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EtldId != 0 {
		n += 1 + sovCtrecord(uint64(m.EtldId))
	}
	l = len(m.Etld)
	if l > 0 {
		n += 1 + l + sovCtrecord(uint64(l))
	}
	if m.QueryTimestamp != 0 {
		n += 1 + sovCtrecord(uint64(m.QueryTimestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCtrecord(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCtrecord(x uint64) (n int) {
	return sovCtrecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CTRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCtrecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CTRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CTRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateID", wireType)
			}
			m.CertificateID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CertificateID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertedTime", wireType)
			}
			m.InsertedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InsertedTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotBefore", wireType)
			}
			m.NotBefore = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotBefore |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotAfter", wireType)
			}
			m.NotAfter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotAfter |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationalUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationalUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommonName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedDNSNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifiedDNSNames = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnverifiedDNSNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnverifiedDNSNames = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddresses = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmailAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmailAddresses = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ETLD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ETLD = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCtrecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCtrecord
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
func (m *CTSubdomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCtrecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CTSubdomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CTSubdomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainID", wireType)
			}
			m.SubdomainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubdomainID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertedTime", wireType)
			}
			m.InsertedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InsertedTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommonName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ETLD", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ETLD = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCtrecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCtrecord
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
func (m *CTETLD) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCtrecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CTETLD: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CTETLD: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtldId", wireType)
			}
			m.EtldId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EtldId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Etld", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCtrecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Etld = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryTimestamp", wireType)
			}
			m.QueryTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCtrecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCtrecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCtrecord
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
func skipCtrecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCtrecord
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
					return 0, ErrIntOverflowCtrecord
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
					return 0, ErrIntOverflowCtrecord
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCtrecord
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCtrecord
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
				next, err := skipCtrecord(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthCtrecord = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCtrecord   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prototypes/ctrecord.proto", fileDescriptor_ctrecord_acaee2e3ab70fdfc) }

var fileDescriptor_ctrecord_acaee2e3ab70fdfc = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0x1d, 0xdb, 0xed, 0x9f, 0xd3, 0xdd, 0xae, 0xe4, 0xc6, 0x28, 0x52, 0x4a, 0x11, 0x5d,
	0x04, 0x5b, 0xc4, 0x27, 0xd8, 0x9d, 0x2e, 0x58, 0x90, 0x51, 0xa6, 0xb3, 0x5e, 0x78, 0xe1, 0x92,
	0xce, 0x9c, 0xee, 0x06, 0x27, 0x93, 0x9a, 0x64, 0x16, 0xea, 0x0b, 0xf8, 0x06, 0xe2, 0x23, 0x79,
	0xe9, 0x23, 0x48, 0x7d, 0x11, 0xc9, 0x59, 0xda, 0x4e, 0xb7, 0x5e, 0x78, 0x77, 0xbe, 0xdf, 0xf7,
	0x25, 0x5f, 0x08, 0x09, 0x3c, 0x5a, 0x18, 0xed, 0xb4, 0x5b, 0x2e, 0xd0, 0x8e, 0x52, 0x67, 0x30,
	0xd5, 0x26, 0x1b, 0x12, 0x1b, 0x7c, 0xaf, 0x43, 0x2b, 0x4c, 0x62, 0x42, 0xec, 0x29, 0x1c, 0x85,
	0x68, 0x9c, 0x9c, 0xcb, 0x54, 0x38, 0x9c, 0x8c, 0x79, 0xd0, 0x0f, 0x4e, 0x6a, 0xf1, 0x2e, 0x64,
	0x03, 0x38, 0x9c, 0x14, 0x16, 0x8d, 0xc3, 0x2c, 0x91, 0x0a, 0xf9, 0x7d, 0x0a, 0xed, 0x30, 0xf6,
	0x18, 0x5a, 0x7e, 0xd1, 0x1b, 0x61, 0xaf, 0x79, 0xad, 0x1f, 0x9c, 0xb4, 0xe3, 0x8d, 0xf6, 0xeb,
	0xa7, 0x68, 0xa4, 0xc8, 0xa3, 0x52, 0xcd, 0xd0, 0xf0, 0x3a, 0xf9, 0x3b, 0x8c, 0x3d, 0x81, 0x76,
	0xa4, 0xdd, 0x19, 0xce, 0xb5, 0x41, 0x7e, 0x40, 0x05, 0x5b, 0xe0, 0x77, 0x8f, 0xb4, 0x3b, 0x9d,
	0x3b, 0x34, 0xbc, 0x41, 0xe6, 0x46, 0x33, 0x0e, 0xcd, 0x50, 0x97, 0x85, 0x33, 0x4b, 0xde, 0xa4,
	0x8d, 0xd7, 0xd2, 0xf7, 0xbe, 0x33, 0x57, 0xa2, 0x90, 0x5f, 0x85, 0x93, 0xba, 0xe0, 0xad, 0xdb,
	0xde, 0x2a, 0x63, 0x43, 0x60, 0x55, 0x2d, 0xf2, 0x8b, 0x42, 0x3a, 0xde, 0xa6, 0xe4, 0x3f, 0x1c,
	0xd6, 0x03, 0x08, 0xb5, 0x52, 0xba, 0x88, 0x84, 0x42, 0x0e, 0x94, 0xab, 0x10, 0xf6, 0x02, 0x1e,
	0x7c, 0x40, 0x23, 0xe7, 0x12, 0xb3, 0x71, 0x34, 0xf5, 0xc8, 0xf2, 0x0e, 0xa5, 0xf6, 0xb8, 0xef,
	0xbe, 0x28, 0x6e, 0xee, 0xa6, 0x0f, 0x6f, 0xbb, 0xf7, 0x1d, 0xd6, 0x87, 0xce, 0xe4, 0xfd, 0x69,
	0x96, 0x19, 0xb4, 0x16, 0x2d, 0x3f, 0xa2, 0x60, 0x15, 0xb1, 0x67, 0xd0, 0x3d, 0x57, 0x42, 0xe6,
	0xdb, 0x50, 0x97, 0x42, 0x77, 0x28, 0x63, 0x50, 0x3f, 0x4f, 0xde, 0x8e, 0xf9, 0x31, 0xb9, 0x34,
	0x0f, 0xbe, 0x05, 0xd0, 0x09, 0x93, 0x69, 0x39, 0xcb, 0xb4, 0x12, 0xb2, 0xf0, 0x6d, 0x1b, 0xb1,
	0x79, 0x19, 0x55, 0xf4, 0x5f, 0xef, 0x62, 0xf7, 0xbe, 0x6a, 0x7b, 0xf7, 0xb5, 0x3e, 0x49, 0xbd,
	0x72, 0x92, 0x4f, 0xd0, 0x08, 0x13, 0x3f, 0xb1, 0x87, 0xd0, 0x44, 0x97, 0x67, 0x97, 0x32, 0xa3,
	0xfe, 0x83, 0xb8, 0xe1, 0xe5, 0x24, 0xf3, 0xcb, 0xfc, 0x44, 0x95, 0xed, 0x98, 0x66, 0xf6, 0x1c,
	0x8e, 0xbf, 0x94, 0x68, 0x96, 0x97, 0x4e, 0x2a, 0xb4, 0x4e, 0xa8, 0x05, 0xf5, 0xd5, 0xe2, 0x2e,
	0xe1, 0x64, 0x4d, 0xcf, 0xc2, 0x9f, 0xab, 0x5e, 0xf0, 0x6b, 0xd5, 0x0b, 0x7e, 0xaf, 0x7a, 0xc1,
	0x8f, 0x3f, 0xbd, 0x7b, 0x1f, 0x5f, 0x5d, 0x49, 0x77, 0x5d, 0xce, 0x86, 0xa9, 0x56, 0xa3, 0x5c,
	0x16, 0x9f, 0x85, 0x7c, 0x29, 0xf5, 0x48, 0xa8, 0x11, 0xfd, 0x99, 0xd4, 0xa2, 0xb9, 0x91, 0x29,
	0xda, 0xd1, 0xf6, 0x5b, 0xcd, 0x1a, 0x34, 0xbf, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x63,
	0x5c, 0x5c, 0x6b, 0x03, 0x00, 0x00,
}
