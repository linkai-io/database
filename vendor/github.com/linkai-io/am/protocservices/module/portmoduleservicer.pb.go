// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: module/portmoduleservicer.proto

package module // import "github.com/linkai-io/am/protocservices/module"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototypes "github.com/linkai-io/am/protocservices/prototypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type AnalyzeWithPortsRequest struct {
	UserContext          *prototypes.UserContext `protobuf:"bytes,1,opt,name=UserContext" json:"UserContext,omitempty"`
	Address              *prototypes.AddressData `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	Ports                *prototypes.PortResults `protobuf:"bytes,3,opt,name=Ports" json:"Ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AnalyzeWithPortsRequest) Reset()         { *m = AnalyzeWithPortsRequest{} }
func (m *AnalyzeWithPortsRequest) String() string { return proto.CompactTextString(m) }
func (*AnalyzeWithPortsRequest) ProtoMessage()    {}
func (*AnalyzeWithPortsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_portmoduleservicer_654aed1b81872684, []int{0}
}
func (m *AnalyzeWithPortsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnalyzeWithPortsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnalyzeWithPortsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AnalyzeWithPortsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeWithPortsRequest.Merge(dst, src)
}
func (m *AnalyzeWithPortsRequest) XXX_Size() int {
	return m.Size()
}
func (m *AnalyzeWithPortsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeWithPortsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeWithPortsRequest proto.InternalMessageInfo

func (m *AnalyzeWithPortsRequest) GetUserContext() *prototypes.UserContext {
	if m != nil {
		return m.UserContext
	}
	return nil
}

func (m *AnalyzeWithPortsRequest) GetAddress() *prototypes.AddressData {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AnalyzeWithPortsRequest) GetPorts() *prototypes.PortResults {
	if m != nil {
		return m.Ports
	}
	return nil
}

type AnalyzedWithPortsResponse struct {
	Original             *prototypes.AddressData            `protobuf:"bytes,1,opt,name=Original" json:"Original,omitempty"`
	Addresses            map[string]*prototypes.AddressData `protobuf:"bytes,2,rep,name=Addresses" json:"Addresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Results              *prototypes.Bag                    `protobuf:"bytes,3,opt,name=Results" json:"Results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *AnalyzedWithPortsResponse) Reset()         { *m = AnalyzedWithPortsResponse{} }
func (m *AnalyzedWithPortsResponse) String() string { return proto.CompactTextString(m) }
func (*AnalyzedWithPortsResponse) ProtoMessage()    {}
func (*AnalyzedWithPortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_portmoduleservicer_654aed1b81872684, []int{1}
}
func (m *AnalyzedWithPortsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnalyzedWithPortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnalyzedWithPortsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AnalyzedWithPortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzedWithPortsResponse.Merge(dst, src)
}
func (m *AnalyzedWithPortsResponse) XXX_Size() int {
	return m.Size()
}
func (m *AnalyzedWithPortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzedWithPortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzedWithPortsResponse proto.InternalMessageInfo

func (m *AnalyzedWithPortsResponse) GetOriginal() *prototypes.AddressData {
	if m != nil {
		return m.Original
	}
	return nil
}

func (m *AnalyzedWithPortsResponse) GetAddresses() map[string]*prototypes.AddressData {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *AnalyzedWithPortsResponse) GetResults() *prototypes.Bag {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*AnalyzeWithPortsRequest)(nil), "module.AnalyzeWithPortsRequest")
	proto.RegisterType((*AnalyzedWithPortsResponse)(nil), "module.AnalyzedWithPortsResponse")
	proto.RegisterMapType((map[string]*prototypes.AddressData)(nil), "module.AnalyzedWithPortsResponse.AddressesEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortModuleClient is the client API for PortModule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortModuleClient interface {
	AnalyzeWithPorts(ctx context.Context, in *AnalyzeWithPortsRequest, opts ...grpc.CallOption) (*AnalyzedWithPortsResponse, error)
}

type portModuleClient struct {
	cc *grpc.ClientConn
}

func NewPortModuleClient(cc *grpc.ClientConn) PortModuleClient {
	return &portModuleClient{cc}
}

func (c *portModuleClient) AnalyzeWithPorts(ctx context.Context, in *AnalyzeWithPortsRequest, opts ...grpc.CallOption) (*AnalyzedWithPortsResponse, error) {
	out := new(AnalyzedWithPortsResponse)
	err := c.cc.Invoke(ctx, "/module.PortModule/AnalyzeWithPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortModuleServer is the server API for PortModule service.
type PortModuleServer interface {
	AnalyzeWithPorts(context.Context, *AnalyzeWithPortsRequest) (*AnalyzedWithPortsResponse, error)
}

func RegisterPortModuleServer(s *grpc.Server, srv PortModuleServer) {
	s.RegisterService(&_PortModule_serviceDesc, srv)
}

func _PortModule_AnalyzeWithPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeWithPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortModuleServer).AnalyzeWithPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/module.PortModule/AnalyzeWithPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortModuleServer).AnalyzeWithPorts(ctx, req.(*AnalyzeWithPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortModule_serviceDesc = grpc.ServiceDesc{
	ServiceName: "module.PortModule",
	HandlerType: (*PortModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyzeWithPorts",
			Handler:    _PortModule_AnalyzeWithPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/portmoduleservicer.proto",
}

func (m *AnalyzeWithPortsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnalyzeWithPortsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserContext != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPortmoduleservicer(dAtA, i, uint64(m.UserContext.Size()))
		n1, err := m.UserContext.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Address != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPortmoduleservicer(dAtA, i, uint64(m.Address.Size()))
		n2, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Ports != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPortmoduleservicer(dAtA, i, uint64(m.Ports.Size()))
		n3, err := m.Ports.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AnalyzedWithPortsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnalyzedWithPortsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Original != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPortmoduleservicer(dAtA, i, uint64(m.Original.Size()))
		n4, err := m.Original.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Addresses) > 0 {
		for k, _ := range m.Addresses {
			dAtA[i] = 0x12
			i++
			v := m.Addresses[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovPortmoduleservicer(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovPortmoduleservicer(uint64(len(k))) + msgSize
			i = encodeVarintPortmoduleservicer(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPortmoduleservicer(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintPortmoduleservicer(dAtA, i, uint64(v.Size()))
				n5, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n5
			}
		}
	}
	if m.Results != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPortmoduleservicer(dAtA, i, uint64(m.Results.Size()))
		n6, err := m.Results.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPortmoduleservicer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AnalyzeWithPortsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserContext != nil {
		l = m.UserContext.Size()
		n += 1 + l + sovPortmoduleservicer(uint64(l))
	}
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovPortmoduleservicer(uint64(l))
	}
	if m.Ports != nil {
		l = m.Ports.Size()
		n += 1 + l + sovPortmoduleservicer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AnalyzedWithPortsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Original != nil {
		l = m.Original.Size()
		n += 1 + l + sovPortmoduleservicer(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for k, v := range m.Addresses {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovPortmoduleservicer(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovPortmoduleservicer(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovPortmoduleservicer(uint64(mapEntrySize))
		}
	}
	if m.Results != nil {
		l = m.Results.Size()
		n += 1 + l + sovPortmoduleservicer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPortmoduleservicer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPortmoduleservicer(x uint64) (n int) {
	return sovPortmoduleservicer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnalyzeWithPortsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPortmoduleservicer
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
			return fmt.Errorf("proto: AnalyzeWithPortsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnalyzeWithPortsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserContext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserContext == nil {
				m.UserContext = &prototypes.UserContext{}
			}
			if err := m.UserContext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Address == nil {
				m.Address = &prototypes.AddressData{}
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ports == nil {
				m.Ports = &prototypes.PortResults{}
			}
			if err := m.Ports.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPortmoduleservicer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPortmoduleservicer
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
func (m *AnalyzedWithPortsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPortmoduleservicer
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
			return fmt.Errorf("proto: AnalyzedWithPortsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnalyzedWithPortsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Original", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Original == nil {
				m.Original = &prototypes.AddressData{}
			}
			if err := m.Original.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Addresses == nil {
				m.Addresses = make(map[string]*prototypes.AddressData)
			}
			var mapkey string
			var mapvalue *prototypes.AddressData
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPortmoduleservicer
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPortmoduleservicer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPortmoduleservicer
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPortmoduleservicer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthPortmoduleservicer
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthPortmoduleservicer
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &prototypes.AddressData{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPortmoduleservicer(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPortmoduleservicer
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Addresses[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortmoduleservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPortmoduleservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Results == nil {
				m.Results = &prototypes.Bag{}
			}
			if err := m.Results.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPortmoduleservicer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPortmoduleservicer
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
func skipPortmoduleservicer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPortmoduleservicer
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
					return 0, ErrIntOverflowPortmoduleservicer
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
					return 0, ErrIntOverflowPortmoduleservicer
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
				return 0, ErrInvalidLengthPortmoduleservicer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPortmoduleservicer
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
				next, err := skipPortmoduleservicer(dAtA[start:])
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
	ErrInvalidLengthPortmoduleservicer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPortmoduleservicer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("module/portmoduleservicer.proto", fileDescriptor_portmoduleservicer_654aed1b81872684)
}

var fileDescriptor_portmoduleservicer_654aed1b81872684 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0xc6, 0x37, 0xed, 0x6e, 0xbb, 0x75, 0x57, 0xab, 0xca, 0xda, 0xd5, 0xa6, 0x39, 0xa4, 0x4b,
	0x0f, 0xa8, 0x97, 0x3a, 0xa8, 0x5c, 0x10, 0x17, 0xd4, 0x02, 0x17, 0x24, 0xfe, 0xc8, 0x12, 0x20,
	0x71, 0x73, 0x13, 0x2b, 0xb5, 0x9a, 0xc6, 0xc1, 0x76, 0x2a, 0xc2, 0x8d, 0x87, 0x40, 0xe2, 0x91,
	0x38, 0xf2, 0x08, 0xa8, 0xbc, 0x08, 0x4a, 0xec, 0x96, 0x50, 0x81, 0xb8, 0x79, 0xbe, 0x6f, 0x66,
	0xfc, 0x1b, 0x7b, 0x40, 0x67, 0xc6, 0x83, 0x34, 0xa2, 0x5e, 0xc2, 0x85, 0xd2, 0x47, 0x49, 0xc5,
	0x9c, 0xf9, 0x54, 0xa0, 0x44, 0x70, 0xc5, 0x61, 0x4d, 0xab, 0x8e, 0x5d, 0x84, 0x2a, 0x4b, 0xa8,
	0xf4, 0x48, 0x10, 0x08, 0x2a, 0xa5, 0xce, 0x70, 0xfe, 0x96, 0x9c, 0x54, 0x2e, 0x0b, 0x9d, 0x76,
	0x49, 0xce, 0xbb, 0x4b, 0x9f, 0xc4, 0xc6, 0xfa, 0x53, 0xb2, 0xc6, 0x24, 0xd4, 0x6a, 0xf7, 0xde,
	0x02, 0xff, 0x86, 0x31, 0x89, 0xb2, 0x5b, 0x7a, 0xc9, 0xd4, 0xe4, 0x2c, 0xaf, 0xc1, 0xf4, 0x3a,
	0xa5, 0x52, 0x41, 0x04, 0x9a, 0xe7, 0x92, 0x8a, 0x7d, 0x1e, 0x2b, 0x7a, 0xa3, 0x6c, 0xeb, 0xbf,
	0xd5, 0x6b, 0x0e, 0x7e, 0xa1, 0x92, 0x86, 0xcb, 0x09, 0x70, 0x13, 0xd4, 0x87, 0x1a, 0xd2, 0xae,
	0x98, 0x5c, 0x13, 0x1f, 0x10, 0x45, 0xf0, 0xd2, 0x84, 0x5d, 0xf0, 0xa3, 0xb8, 0xc7, 0xae, 0x9a,
	0xac, 0x3c, 0xc2, 0x54, 0xa6, 0x91, 0x92, 0x58, 0x5b, 0xdd, 0xbb, 0x0a, 0x68, 0x1b, 0xae, 0xa0,
	0x04, 0x26, 0x13, 0x1e, 0x4b, 0x0a, 0x7b, 0xe0, 0xe7, 0xa9, 0x60, 0x21, 0x8b, 0x49, 0xb4, 0xc2,
	0x2a, 0x5f, 0xb5, 0x72, 0xe1, 0x09, 0x68, 0x18, 0x83, 0xe6, 0x54, 0xd5, 0x5e, 0x73, 0xb0, 0x85,
	0xf4, 0xeb, 0xa2, 0x4f, 0xfb, 0xa3, 0x55, 0xc9, 0x61, 0xac, 0x44, 0x86, 0xdf, 0x5a, 0x40, 0x17,
	0xd4, 0x0d, 0xa9, 0xa1, 0xff, 0x8e, 0x46, 0x24, 0xc4, 0x4b, 0xd1, 0x39, 0x02, 0xbf, 0xdf, 0x17,
	0xc3, 0x16, 0xa8, 0x4e, 0x69, 0x56, 0x60, 0x36, 0x70, 0x7e, 0xcc, 0xe7, 0x9f, 0x93, 0x28, 0xa5,
	0x1f, 0xbe, 0x92, 0xb6, 0x76, 0x2b, 0x3b, 0xd6, 0x20, 0x00, 0x20, 0xc7, 0x3a, 0x2e, 0x68, 0xe1,
	0x05, 0x68, 0xad, 0x7f, 0x14, 0xec, 0xac, 0x8d, 0xb2, 0xfe, 0x85, 0xce, 0xc6, 0x97, 0xb3, 0x8e,
	0xf6, 0x1e, 0x17, 0xae, 0xf5, 0xb4, 0x70, 0xad, 0xe7, 0x85, 0x6b, 0x3d, 0xbc, 0xb8, 0xdf, 0xae,
	0xfa, 0x21, 0x53, 0x93, 0x74, 0x8c, 0x7c, 0x3e, 0xf3, 0x22, 0x16, 0x4f, 0x09, 0xeb, 0x33, 0xee,
	0x91, 0x99, 0x57, 0xac, 0x8c, 0x6f, 0x76, 0x55, 0x7a, 0xba, 0xf5, 0xb8, 0x56, 0xc8, 0xdb, 0xaf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xe1, 0xff, 0x6b, 0xd6, 0x02, 0x00, 0x00,
}
