// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/distance_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A distance view with metrics aggregated by the user’s distance from an
// advertiser’s location extensions. Each DistanceBucket includes all
// impressions that fall within its distance and a single impression will
// contribute to the metrics for all DistanceBuckets that include the user’s
// distance.
type DistanceView struct {
	// The resource name of the distance view.
	// Distance view resource names have the form:
	//
	// `customers/{customer_id}/distanceViews/1~{distance_bucket}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Grouping of user distance from location extensions.
	DistanceBucket enums.DistanceBucketEnum_DistanceBucket `protobuf:"varint,2,opt,name=distance_bucket,json=distanceBucket,proto3,enum=google.ads.googleads.v2.enums.DistanceBucketEnum_DistanceBucket" json:"distance_bucket,omitempty"`
	// True if the DistanceBucket is using the metric system, false otherwise.
	MetricSystem         *wrappers.BoolValue `protobuf:"bytes,3,opt,name=metric_system,json=metricSystem,proto3" json:"metric_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DistanceView) Reset()         { *m = DistanceView{} }
func (m *DistanceView) String() string { return proto.CompactTextString(m) }
func (*DistanceView) ProtoMessage()    {}
func (*DistanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_3430a19803c9c189, []int{0}
}

func (m *DistanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistanceView.Unmarshal(m, b)
}
func (m *DistanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistanceView.Marshal(b, m, deterministic)
}
func (m *DistanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistanceView.Merge(m, src)
}
func (m *DistanceView) XXX_Size() int {
	return xxx_messageInfo_DistanceView.Size(m)
}
func (m *DistanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_DistanceView.DiscardUnknown(m)
}

var xxx_messageInfo_DistanceView proto.InternalMessageInfo

func (m *DistanceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DistanceView) GetDistanceBucket() enums.DistanceBucketEnum_DistanceBucket {
	if m != nil {
		return m.DistanceBucket
	}
	return enums.DistanceBucketEnum_UNSPECIFIED
}

func (m *DistanceView) GetMetricSystem() *wrappers.BoolValue {
	if m != nil {
		return m.MetricSystem
	}
	return nil
}

func init() {
	proto.RegisterType((*DistanceView)(nil), "google.ads.googleads.v2.resources.DistanceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/distance_view.proto", fileDescriptor_3430a19803c9c189)
}

var fileDescriptor_3430a19803c9c189 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x0a, 0x17, 0x6e, 0x6e, 0xdb, 0x8b, 0x59, 0x95, 0x22, 0xd2, 0x2a, 0x85, 0xae,
	0x26, 0x90, 0xe2, 0x66, 0x5c, 0x68, 0x82, 0x52, 0x70, 0x21, 0x25, 0x42, 0x16, 0x12, 0x28, 0xd3,
	0xe4, 0x18, 0x06, 0x93, 0x99, 0x90, 0x49, 0x5a, 0x7c, 0x1d, 0x97, 0x3e, 0x8a, 0x2f, 0xe1, 0xde,
	0x87, 0x10, 0x69, 0x26, 0x33, 0xd4, 0x42, 0x75, 0xf7, 0xcf, 0x9c, 0xef, 0xff, 0xcf, 0x39, 0x93,
	0x58, 0xe7, 0x29, 0xe7, 0x69, 0x06, 0x0e, 0x49, 0x84, 0x23, 0xe5, 0x56, 0xad, 0x5d, 0xa7, 0x04,
	0xc1, 0xeb, 0x32, 0x06, 0xe1, 0x24, 0x54, 0x54, 0x84, 0xc5, 0xb0, 0x5c, 0x53, 0xd8, 0xa0, 0xa2,
	0xe4, 0x15, 0xb7, 0xc7, 0x92, 0x45, 0x24, 0x11, 0x48, 0xdb, 0xd0, 0xda, 0x45, 0xda, 0x36, 0x9c,
	0x1d, 0x4a, 0x06, 0x56, 0xe7, 0x3b, 0xa9, 0xab, 0x3a, 0x7e, 0x82, 0x4a, 0xe6, 0x0e, 0x4f, 0x5a,
	0x53, 0x73, 0x5a, 0xd5, 0x8f, 0xce, 0xa6, 0x24, 0x45, 0x01, 0xa5, 0x68, 0xeb, 0xc7, 0x2a, 0xb4,
	0xa0, 0x0e, 0x61, 0x8c, 0x57, 0xa4, 0xa2, 0x9c, 0xb5, 0xd5, 0xd3, 0x77, 0xc3, 0xea, 0x5e, 0xb7,
	0xb9, 0x21, 0x85, 0x8d, 0x7d, 0x66, 0xf5, 0xd4, 0x40, 0x4b, 0x46, 0x72, 0x18, 0x18, 0x23, 0x63,
	0xfa, 0x37, 0xe8, 0xaa, 0xcb, 0x3b, 0x92, 0x83, 0x4d, 0xad, 0xff, 0x7b, 0xc3, 0x0c, 0xcc, 0x91,
	0x31, 0xed, 0xbb, 0x57, 0xe8, 0xd0, 0x96, 0xcd, 0x0a, 0x48, 0xb5, 0xf2, 0x1b, 0xd3, 0x0d, 0xab,
	0xf3, 0xbd, 0xab, 0xa0, 0x9f, 0x7c, 0x3b, 0xdb, 0x97, 0x56, 0x2f, 0x87, 0xaa, 0xa4, 0xf1, 0x52,
	0x3c, 0x8b, 0x0a, 0xf2, 0x41, 0x67, 0x64, 0x4c, 0xff, 0xb9, 0x43, 0xd5, 0x48, 0xad, 0x8d, 0x7c,
	0xce, 0xb3, 0x90, 0x64, 0x35, 0x04, 0x5d, 0x69, 0xb8, 0x6f, 0x78, 0xff, 0xd3, 0xb0, 0x26, 0x31,
	0xcf, 0xd1, 0xaf, 0xcf, 0xef, 0x1f, 0xed, 0x3e, 0xc4, 0x62, 0x9b, 0xbb, 0x30, 0x1e, 0x6e, 0x5b,
	0x5f, 0xca, 0x33, 0xc2, 0x52, 0xc4, 0xcb, 0xd4, 0x49, 0x81, 0x35, 0x5d, 0xd5, 0x37, 0x2a, 0xa8,
	0xf8, 0xe1, 0x67, 0xb8, 0xd0, 0xea, 0xc5, 0xec, 0xcc, 0x3d, 0xef, 0xd5, 0x1c, 0xcf, 0x65, 0xa4,
	0x97, 0x08, 0x24, 0xe5, 0x56, 0x85, 0x2e, 0x0a, 0x14, 0xf9, 0xa6, 0x98, 0xc8, 0x4b, 0x44, 0xa4,
	0x99, 0x28, 0x74, 0x23, 0xcd, 0x7c, 0x98, 0x13, 0x59, 0xc0, 0xd8, 0x4b, 0x04, 0xc6, 0x9a, 0xc2,
	0x38, 0x74, 0x31, 0xd6, 0xdc, 0xea, 0x4f, 0x33, 0xec, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x2f, 0x26, 0x1d, 0xb8, 0x02, 0x00, 0x00,
}
