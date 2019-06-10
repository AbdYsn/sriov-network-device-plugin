// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/dates.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A date range.
type DateRange struct {
	// The start date, in yyyy-mm-dd format. This date is inclusive.
	StartDate *wrappers.StringValue `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format. This date is inclusive.
	EndDate              *wrappers.StringValue `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f82ad1a92cbb16, []int{0}
}

func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (m *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(m, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DateRange) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.ads.googleads.v2.common.DateRange")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/dates.proto", fileDescriptor_50f82ad1a92cbb16)
}

var fileDescriptor_50f82ad1a92cbb16 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0xb1, 0x7f, 0xf8, 0xdb, 0xa8, 0x5b, 0xa6, 0x12, 0x42, 0x28, 0x9e, 0x4a, 0x07, 0x09,
	0xd4, 0xa1, 0xa0, 0x4c, 0x4e, 0x03, 0x59, 0x43, 0x0a, 0x1e, 0x8a, 0xa1, 0xdc, 0x44, 0xaa, 0x30,
	0xd8, 0xba, 0x46, 0x52, 0xd2, 0xb9, 0xaf, 0xd2, 0xb1, 0x8f, 0xd2, 0xf7, 0xe8, 0xd2, 0xa7, 0x28,
	0xb2, 0x6c, 0x6f, 0x2d, 0x9d, 0x7c, 0xf0, 0xfd, 0xce, 0x39, 0x57, 0x97, 0xdc, 0x68, 0x44, 0x5d,
	0x2b, 0x06, 0xd2, 0xb1, 0x28, 0x83, 0x3a, 0x71, 0x76, 0xc0, 0xa6, 0x41, 0xc3, 0x24, 0x78, 0xe5,
	0x68, 0x6b, 0xd1, 0xe3, 0x74, 0x11, 0x01, 0x0a, 0xd2, 0xd1, 0x91, 0xa5, 0x27, 0x4e, 0x23, 0x3b,
	0xeb, 0xe7, 0xac, 0xa3, 0xf7, 0xc7, 0x67, 0xf6, 0x62, 0xa1, 0x6d, 0x95, 0xed, 0xfd, 0xb3, 0xf9,
	0xd0, 0xd5, 0x56, 0x0c, 0x8c, 0x41, 0x0f, 0xbe, 0x42, 0xd3, 0x4f, 0xb3, 0xd7, 0x84, 0x4c, 0xd6,
	0xe0, 0xd5, 0x0e, 0x8c, 0x56, 0xd3, 0x25, 0x21, 0xce, 0x83, 0xf5, 0x4f, 0x61, 0x81, 0xcb, 0xe4,
	0x2a, 0xb9, 0xbe, 0xe0, 0xf3, 0xbe, 0x95, 0x0e, 0x05, 0xf4, 0xc1, 0xdb, 0xca, 0xe8, 0x02, 0xea,
	0xa3, 0xda, 0x4d, 0x3a, 0x3e, 0x24, 0x4c, 0xef, 0xc8, 0xb9, 0x32, 0x32, 0x5a, 0xd3, 0x3f, 0x58,
	0xcf, 0x94, 0x91, 0xc1, 0xb8, 0xfa, 0x4c, 0x48, 0x76, 0xc0, 0x86, 0xfe, 0xfe, 0xd0, 0x15, 0x09,
	0xb0, 0xdb, 0x86, 0xa8, 0x6d, 0xf2, 0xb8, 0xee, 0x69, 0x8d, 0x35, 0x18, 0x4d, 0xd1, 0x6a, 0xa6,
	0x95, 0xe9, 0x8a, 0x86, 0x93, 0xb6, 0x95, 0xfb, 0xe9, 0xc2, 0xcb, 0xf8, 0x79, 0x4b, 0xff, 0x6d,
	0xf2, 0xfc, 0x3d, 0x5d, 0x6c, 0x62, 0x58, 0x2e, 0x1d, 0x8d, 0x32, 0xa8, 0x82, 0xd3, 0xfb, 0x0e,
	0xfb, 0x18, 0x80, 0x32, 0x97, 0xae, 0x1c, 0x81, 0xb2, 0xe0, 0x65, 0x04, 0xbe, 0xd2, 0x2c, 0xfe,
	0x15, 0x22, 0x97, 0x4e, 0x88, 0x11, 0x11, 0xa2, 0xe0, 0x42, 0x44, 0x68, 0xff, 0xbf, 0xdb, 0xee,
	0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x13, 0xd9, 0x53, 0xfe, 0x01, 0x00, 0x00,
}