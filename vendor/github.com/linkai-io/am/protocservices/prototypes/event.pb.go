// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prototypes/event.proto

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

type EventData struct {
	NotificationID       int64    `protobuf:"varint,1,opt,name=NotificationID,proto3" json:"NotificationID,omitempty"`
	OrgID                int32    `protobuf:"varint,2,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	GroupID              int32    `protobuf:"varint,3,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	TypeID               int32    `protobuf:"varint,4,opt,name=TypeID,proto3" json:"TypeID,omitempty"`
	EventTimestamp       int64    `protobuf:"varint,5,opt,name=EventTimestamp,proto3" json:"EventTimestamp,omitempty"`
	Data                 []string `protobuf:"bytes,6,rep,name=Data" json:"Data,omitempty"`
	Read                 bool     `protobuf:"varint,7,opt,name=Read,proto3" json:"Read,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventData) Reset()         { *m = EventData{} }
func (m *EventData) String() string { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()    {}
func (*EventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_99ba53dbf15c9669, []int{0}
}
func (m *EventData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventData.Merge(dst, src)
}
func (m *EventData) XXX_Size() int {
	return m.Size()
}
func (m *EventData) XXX_DiscardUnknown() {
	xxx_messageInfo_EventData.DiscardUnknown(m)
}

var xxx_messageInfo_EventData proto.InternalMessageInfo

func (m *EventData) GetNotificationID() int64 {
	if m != nil {
		return m.NotificationID
	}
	return 0
}

func (m *EventData) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *EventData) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *EventData) GetTypeID() int32 {
	if m != nil {
		return m.TypeID
	}
	return 0
}

func (m *EventData) GetEventTimestamp() int64 {
	if m != nil {
		return m.EventTimestamp
	}
	return 0
}

func (m *EventData) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EventData) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

type EventSubscriptions struct {
	TypeID               int32    `protobuf:"varint,1,opt,name=TypeID,proto3" json:"TypeID,omitempty"`
	SubscribedTimestamp  int64    `protobuf:"varint,2,opt,name=SubscribedTimestamp,proto3" json:"SubscribedTimestamp,omitempty"`
	Subscribed           bool     `protobuf:"varint,3,opt,name=Subscribed,proto3" json:"Subscribed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSubscriptions) Reset()         { *m = EventSubscriptions{} }
func (m *EventSubscriptions) String() string { return proto.CompactTextString(m) }
func (*EventSubscriptions) ProtoMessage()    {}
func (*EventSubscriptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_99ba53dbf15c9669, []int{1}
}
func (m *EventSubscriptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSubscriptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSubscriptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventSubscriptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSubscriptions.Merge(dst, src)
}
func (m *EventSubscriptions) XXX_Size() int {
	return m.Size()
}
func (m *EventSubscriptions) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSubscriptions.DiscardUnknown(m)
}

var xxx_messageInfo_EventSubscriptions proto.InternalMessageInfo

func (m *EventSubscriptions) GetTypeID() int32 {
	if m != nil {
		return m.TypeID
	}
	return 0
}

func (m *EventSubscriptions) GetSubscribedTimestamp() int64 {
	if m != nil {
		return m.SubscribedTimestamp
	}
	return 0
}

func (m *EventSubscriptions) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

type UserEventSettings struct {
	WeeklyReportSendDay  int32                 `protobuf:"varint,1,opt,name=WeeklyReportSendDay,proto3" json:"WeeklyReportSendDay,omitempty"`
	ShouldWeeklyEmail    bool                  `protobuf:"varint,2,opt,name=ShouldWeeklyEmail,proto3" json:"ShouldWeeklyEmail,omitempty"`
	DailyReportSendHour  int32                 `protobuf:"varint,3,opt,name=DailyReportSendHour,proto3" json:"DailyReportSendHour,omitempty"`
	ShouldDailyEmail     bool                  `protobuf:"varint,4,opt,name=ShouldDailyEmail,proto3" json:"ShouldDailyEmail,omitempty"`
	UserTimezone         string                `protobuf:"bytes,5,opt,name=UserTimezone,proto3" json:"UserTimezone,omitempty"`
	Subscriptions        []*EventSubscriptions `protobuf:"bytes,6,rep,name=Subscriptions" json:"Subscriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserEventSettings) Reset()         { *m = UserEventSettings{} }
func (m *UserEventSettings) String() string { return proto.CompactTextString(m) }
func (*UserEventSettings) ProtoMessage()    {}
func (*UserEventSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_99ba53dbf15c9669, []int{2}
}
func (m *UserEventSettings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserEventSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserEventSettings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserEventSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEventSettings.Merge(dst, src)
}
func (m *UserEventSettings) XXX_Size() int {
	return m.Size()
}
func (m *UserEventSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEventSettings.DiscardUnknown(m)
}

var xxx_messageInfo_UserEventSettings proto.InternalMessageInfo

func (m *UserEventSettings) GetWeeklyReportSendDay() int32 {
	if m != nil {
		return m.WeeklyReportSendDay
	}
	return 0
}

func (m *UserEventSettings) GetShouldWeeklyEmail() bool {
	if m != nil {
		return m.ShouldWeeklyEmail
	}
	return false
}

func (m *UserEventSettings) GetDailyReportSendHour() int32 {
	if m != nil {
		return m.DailyReportSendHour
	}
	return 0
}

func (m *UserEventSettings) GetShouldDailyEmail() bool {
	if m != nil {
		return m.ShouldDailyEmail
	}
	return false
}

func (m *UserEventSettings) GetUserTimezone() string {
	if m != nil {
		return m.UserTimezone
	}
	return ""
}

func (m *UserEventSettings) GetSubscriptions() []*EventSubscriptions {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type EventFilter struct {
	Start                int64       `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
	Limit                int32       `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Filters              *FilterType `protobuf:"bytes,3,opt,name=Filters" json:"Filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_99ba53dbf15c9669, []int{3}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(dst, src)
}
func (m *EventFilter) XXX_Size() int {
	return m.Size()
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *EventFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *EventFilter) GetFilters() *FilterType {
	if m != nil {
		return m.Filters
	}
	return nil
}

func init() {
	proto.RegisterType((*EventData)(nil), "EventData")
	proto.RegisterType((*EventSubscriptions)(nil), "EventSubscriptions")
	proto.RegisterType((*UserEventSettings)(nil), "UserEventSettings")
	proto.RegisterType((*EventFilter)(nil), "EventFilter")
}
func (m *EventData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NotificationID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.NotificationID))
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.OrgID))
	}
	if m.GroupID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.GroupID))
	}
	if m.TypeID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.TypeID))
	}
	if m.EventTimestamp != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.EventTimestamp))
	}
	if len(m.Data) > 0 {
		for _, s := range m.Data {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Read {
		dAtA[i] = 0x38
		i++
		if m.Read {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EventSubscriptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSubscriptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TypeID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.TypeID))
	}
	if m.SubscribedTimestamp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.SubscribedTimestamp))
	}
	if m.Subscribed {
		dAtA[i] = 0x18
		i++
		if m.Subscribed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UserEventSettings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserEventSettings) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WeeklyReportSendDay != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.WeeklyReportSendDay))
	}
	if m.ShouldWeeklyEmail {
		dAtA[i] = 0x10
		i++
		if m.ShouldWeeklyEmail {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DailyReportSendHour != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.DailyReportSendHour))
	}
	if m.ShouldDailyEmail {
		dAtA[i] = 0x20
		i++
		if m.ShouldDailyEmail {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.UserTimezone) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.UserTimezone)))
		i += copy(dAtA[i:], m.UserTimezone)
	}
	if len(m.Subscriptions) > 0 {
		for _, msg := range m.Subscriptions {
			dAtA[i] = 0x32
			i++
			i = encodeVarintEvent(dAtA, i, uint64(msg.Size()))
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

func (m *EventFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Start))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Limit))
	}
	if m.Filters != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Filters.Size()))
		n1, err := m.Filters.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EventData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotificationID != 0 {
		n += 1 + sovEvent(uint64(m.NotificationID))
	}
	if m.OrgID != 0 {
		n += 1 + sovEvent(uint64(m.OrgID))
	}
	if m.GroupID != 0 {
		n += 1 + sovEvent(uint64(m.GroupID))
	}
	if m.TypeID != 0 {
		n += 1 + sovEvent(uint64(m.TypeID))
	}
	if m.EventTimestamp != 0 {
		n += 1 + sovEvent(uint64(m.EventTimestamp))
	}
	if len(m.Data) > 0 {
		for _, s := range m.Data {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if m.Read {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EventSubscriptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypeID != 0 {
		n += 1 + sovEvent(uint64(m.TypeID))
	}
	if m.SubscribedTimestamp != 0 {
		n += 1 + sovEvent(uint64(m.SubscribedTimestamp))
	}
	if m.Subscribed {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserEventSettings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WeeklyReportSendDay != 0 {
		n += 1 + sovEvent(uint64(m.WeeklyReportSendDay))
	}
	if m.ShouldWeeklyEmail {
		n += 2
	}
	if m.DailyReportSendHour != 0 {
		n += 1 + sovEvent(uint64(m.DailyReportSendHour))
	}
	if m.ShouldDailyEmail {
		n += 2
	}
	l = len(m.UserTimezone)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EventFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovEvent(uint64(m.Start))
	}
	if m.Limit != 0 {
		n += 1 + sovEvent(uint64(m.Limit))
	}
	if m.Filters != nil {
		l = m.Filters.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotificationID", wireType)
			}
			m.NotificationID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotificationID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			m.TypeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TypeID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventTimestamp", wireType)
			}
			m.EventTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Read", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Read = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventSubscriptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventSubscriptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSubscriptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeID", wireType)
			}
			m.TypeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TypeID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscribedTimestamp", wireType)
			}
			m.SubscribedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscribedTimestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *UserEventSettings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: UserEventSettings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserEventSettings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeeklyReportSendDay", wireType)
			}
			m.WeeklyReportSendDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WeeklyReportSendDay |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShouldWeeklyEmail", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ShouldWeeklyEmail = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DailyReportSendHour", wireType)
			}
			m.DailyReportSendHour = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DailyReportSendHour |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShouldDailyEmail", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ShouldDailyEmail = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserTimezone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserTimezone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, &EventSubscriptions{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Filters == nil {
				m.Filters = &FilterType{}
			}
			if err := m.Filters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvent
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
				next, err := skipEvent(dAtA[start:])
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
	ErrInvalidLengthEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prototypes/event.proto", fileDescriptor_event_99ba53dbf15c9669) }

var fileDescriptor_event_99ba53dbf15c9669 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x36, 0xfd, 0xef, 0xa9, 0x8a, 0x9b, 0x95, 0x65, 0x50, 0x28, 0x65, 0x40, 0x29, 0xa2, 0x1d,
	0x5d, 0xaf, 0xbc, 0xd5, 0x59, 0xb5, 0x20, 0x0a, 0xe9, 0x8a, 0xe0, 0x95, 0xe9, 0x34, 0xdb, 0x0d,
	0x3b, 0x33, 0x19, 0x92, 0xcc, 0x42, 0xbd, 0xf0, 0x39, 0x7c, 0x0c, 0x1f, 0x43, 0xf0, 0xc6, 0x47,
	0x90, 0xfa, 0x22, 0x92, 0x93, 0x29, 0xed, 0xd8, 0xde, 0x9d, 0xf3, 0x7d, 0x27, 0xe7, 0xfb, 0x92,
	0x73, 0x02, 0x27, 0x85, 0x56, 0x56, 0xd9, 0x55, 0x21, 0x4c, 0x24, 0xae, 0x45, 0x6e, 0x27, 0x08,
	0xdc, 0xbb, 0xbf, 0x83, 0x5f, 0xc8, 0xd4, 0x0a, 0xed, 0x62, 0x4f, 0x86, 0xbf, 0x08, 0xf4, 0xcf,
	0x5c, 0x71, 0xcc, 0x2d, 0xa7, 0x0f, 0xe1, 0xf6, 0x7b, 0x65, 0xe5, 0x85, 0x4c, 0xb8, 0x95, 0x2a,
	0x9f, 0xc6, 0x01, 0x19, 0x91, 0x71, 0x93, 0xfd, 0x87, 0xd2, 0xbb, 0xd0, 0xfe, 0xa0, 0x97, 0xd3,
	0x38, 0x68, 0x8c, 0xc8, 0xb8, 0xcd, 0x7c, 0x42, 0x03, 0xe8, 0xbe, 0xd1, 0xaa, 0x2c, 0xa6, 0x71,
	0xd0, 0x44, 0x7c, 0x93, 0xd2, 0x13, 0xe8, 0x9c, 0xaf, 0x0a, 0x31, 0x8d, 0x83, 0x16, 0x12, 0x55,
	0xe6, 0xf4, 0x50, 0xfc, 0x5c, 0x66, 0xc2, 0x58, 0x9e, 0x15, 0x41, 0xdb, 0xeb, 0xd5, 0x51, 0x4a,
	0xa1, 0xe5, 0xfc, 0x05, 0x9d, 0x51, 0x73, 0xdc, 0x67, 0x18, 0x3b, 0x8c, 0x09, 0xbe, 0x08, 0xba,
	0x23, 0x32, 0xee, 0x31, 0x8c, 0xc3, 0x6f, 0x40, 0xf1, 0xe4, 0xac, 0x9c, 0x9b, 0x44, 0xcb, 0xc2,
	0xd9, 0x35, 0x3b, 0xea, 0xa4, 0xa6, 0xfe, 0x14, 0x8e, 0xab, 0xc2, 0xb9, 0x58, 0x6c, 0x2d, 0x34,
	0xd0, 0xc2, 0x21, 0x8a, 0x0e, 0x01, 0xb6, 0x30, 0x5e, 0xb2, 0xc7, 0x76, 0x90, 0xf0, 0x47, 0x03,
	0x8e, 0x3e, 0x1a, 0xa1, 0xbd, 0x09, 0x61, 0xad, 0xcc, 0x97, 0xc6, 0xe9, 0x7c, 0x12, 0xe2, 0x2a,
	0x5d, 0x31, 0x51, 0x28, 0x6d, 0x67, 0x22, 0x5f, 0xc4, 0x7c, 0x55, 0x99, 0x39, 0x44, 0xd1, 0xc7,
	0x70, 0x34, 0xbb, 0x54, 0x65, 0xba, 0xf0, 0xe4, 0x59, 0xc6, 0x65, 0x8a, 0xbe, 0x7a, 0x6c, 0x9f,
	0x70, 0xfd, 0x63, 0x2e, 0x77, 0x7b, 0xbc, 0x55, 0xa5, 0xae, 0x66, 0x70, 0x88, 0xa2, 0x8f, 0xe0,
	0x8e, 0x6f, 0x83, 0xa4, 0x6f, 0xdf, 0xc2, 0xf6, 0x7b, 0x38, 0x0d, 0xe1, 0xa6, 0xbb, 0x92, 0x7b,
	0x84, 0xaf, 0x2a, 0x17, 0x38, 0xa1, 0x3e, 0xab, 0x61, 0xf4, 0x05, 0xdc, 0xaa, 0x3d, 0x39, 0x0e,
	0x6a, 0x70, 0x7a, 0x3c, 0xd9, 0x9f, 0x06, 0xab, 0x57, 0x86, 0x5f, 0x60, 0x80, 0x45, 0xaf, 0x71,
	0x33, 0xdd, 0x66, 0xcd, 0x2c, 0xd7, 0xb6, 0x5a, 0x3c, 0x9f, 0x38, 0xf4, 0x9d, 0xcc, 0xa4, 0xdd,
	0xec, 0x1b, 0x26, 0xf4, 0x01, 0x74, 0xfd, 0x29, 0x83, 0x77, 0x1d, 0x9c, 0x0e, 0x26, 0x3e, 0x77,
	0xf3, 0x65, 0x1b, 0xee, 0xe5, 0xab, 0x9f, 0xeb, 0x21, 0xf9, 0xbd, 0x1e, 0x92, 0x3f, 0xeb, 0x21,
	0xf9, 0xfe, 0x77, 0x78, 0xe3, 0xf3, 0xb3, 0xa5, 0xb4, 0x97, 0xe5, 0x7c, 0x92, 0xa8, 0x2c, 0x4a,
	0x65, 0x7e, 0xc5, 0xe5, 0x13, 0xa9, 0x22, 0x9e, 0x45, 0xf8, 0x27, 0x12, 0x23, 0xf4, 0xb5, 0x4c,
	0x84, 0x89, 0xb6, 0x1f, 0x67, 0xde, 0xc1, 0xf8, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0x4c, 0xa0, 0xca, 0x65, 0x03, 0x00, 0x00,
}
