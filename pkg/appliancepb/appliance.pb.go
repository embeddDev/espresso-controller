// Code generated by protoc-gen-go. DO NOT EDIT.
// source: appliance.proto

package appliancepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TemperatureSample struct {
	Value                float64              `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	ObservedAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=observed_at,json=observedAt,proto3" json:"observed_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TemperatureSample) Reset()         { *m = TemperatureSample{} }
func (m *TemperatureSample) String() string { return proto.CompactTextString(m) }
func (*TemperatureSample) ProtoMessage()    {}
func (*TemperatureSample) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{0}
}

func (m *TemperatureSample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureSample.Unmarshal(m, b)
}
func (m *TemperatureSample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureSample.Marshal(b, m, deterministic)
}
func (m *TemperatureSample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureSample.Merge(m, src)
}
func (m *TemperatureSample) XXX_Size() int {
	return xxx_messageInfo_TemperatureSample.Size(m)
}
func (m *TemperatureSample) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureSample.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureSample proto.InternalMessageInfo

func (m *TemperatureSample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TemperatureSample) GetObservedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ObservedAt
	}
	return nil
}

type GetBoilerTemperatureHistoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBoilerTemperatureHistoryRequest) Reset()         { *m = GetBoilerTemperatureHistoryRequest{} }
func (m *GetBoilerTemperatureHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetBoilerTemperatureHistoryRequest) ProtoMessage()    {}
func (*GetBoilerTemperatureHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{1}
}

func (m *GetBoilerTemperatureHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBoilerTemperatureHistoryRequest.Unmarshal(m, b)
}
func (m *GetBoilerTemperatureHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBoilerTemperatureHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetBoilerTemperatureHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBoilerTemperatureHistoryRequest.Merge(m, src)
}
func (m *GetBoilerTemperatureHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetBoilerTemperatureHistoryRequest.Size(m)
}
func (m *GetBoilerTemperatureHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBoilerTemperatureHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBoilerTemperatureHistoryRequest proto.InternalMessageInfo

type GetBoilerTemperatureHistoryResponse struct {
	Samples              []*TemperatureSample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetBoilerTemperatureHistoryResponse) Reset()         { *m = GetBoilerTemperatureHistoryResponse{} }
func (m *GetBoilerTemperatureHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetBoilerTemperatureHistoryResponse) ProtoMessage()    {}
func (*GetBoilerTemperatureHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{2}
}

func (m *GetBoilerTemperatureHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBoilerTemperatureHistoryResponse.Unmarshal(m, b)
}
func (m *GetBoilerTemperatureHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBoilerTemperatureHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetBoilerTemperatureHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBoilerTemperatureHistoryResponse.Merge(m, src)
}
func (m *GetBoilerTemperatureHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetBoilerTemperatureHistoryResponse.Size(m)
}
func (m *GetBoilerTemperatureHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBoilerTemperatureHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBoilerTemperatureHistoryResponse proto.InternalMessageInfo

func (m *GetBoilerTemperatureHistoryResponse) GetSamples() []*TemperatureSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type GetCurrentBoilerTemperatureRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentBoilerTemperatureRequest) Reset()         { *m = GetCurrentBoilerTemperatureRequest{} }
func (m *GetCurrentBoilerTemperatureRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentBoilerTemperatureRequest) ProtoMessage()    {}
func (*GetCurrentBoilerTemperatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{3}
}

func (m *GetCurrentBoilerTemperatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentBoilerTemperatureRequest.Unmarshal(m, b)
}
func (m *GetCurrentBoilerTemperatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentBoilerTemperatureRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentBoilerTemperatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentBoilerTemperatureRequest.Merge(m, src)
}
func (m *GetCurrentBoilerTemperatureRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentBoilerTemperatureRequest.Size(m)
}
func (m *GetCurrentBoilerTemperatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentBoilerTemperatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentBoilerTemperatureRequest proto.InternalMessageInfo

type GetCurrentBoilerTemperatureResponse struct {
	Sample               *TemperatureSample `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetCurrentBoilerTemperatureResponse) Reset()         { *m = GetCurrentBoilerTemperatureResponse{} }
func (m *GetCurrentBoilerTemperatureResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentBoilerTemperatureResponse) ProtoMessage()    {}
func (*GetCurrentBoilerTemperatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{4}
}

func (m *GetCurrentBoilerTemperatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentBoilerTemperatureResponse.Unmarshal(m, b)
}
func (m *GetCurrentBoilerTemperatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentBoilerTemperatureResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentBoilerTemperatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentBoilerTemperatureResponse.Merge(m, src)
}
func (m *GetCurrentBoilerTemperatureResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentBoilerTemperatureResponse.Size(m)
}
func (m *GetCurrentBoilerTemperatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentBoilerTemperatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentBoilerTemperatureResponse proto.InternalMessageInfo

func (m *GetCurrentBoilerTemperatureResponse) GetSample() *TemperatureSample {
	if m != nil {
		return m.Sample
	}
	return nil
}

type GetTargetTemperatureRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTargetTemperatureRequest) Reset()         { *m = GetTargetTemperatureRequest{} }
func (m *GetTargetTemperatureRequest) String() string { return proto.CompactTextString(m) }
func (*GetTargetTemperatureRequest) ProtoMessage()    {}
func (*GetTargetTemperatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{5}
}

func (m *GetTargetTemperatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetTemperatureRequest.Unmarshal(m, b)
}
func (m *GetTargetTemperatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetTemperatureRequest.Marshal(b, m, deterministic)
}
func (m *GetTargetTemperatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetTemperatureRequest.Merge(m, src)
}
func (m *GetTargetTemperatureRequest) XXX_Size() int {
	return xxx_messageInfo_GetTargetTemperatureRequest.Size(m)
}
func (m *GetTargetTemperatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetTemperatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetTemperatureRequest proto.InternalMessageInfo

type GetTargetTemperatureResponse struct {
	Temperature          float64              `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	SetAt                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=set_at,json=setAt,proto3" json:"set_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetTargetTemperatureResponse) Reset()         { *m = GetTargetTemperatureResponse{} }
func (m *GetTargetTemperatureResponse) String() string { return proto.CompactTextString(m) }
func (*GetTargetTemperatureResponse) ProtoMessage()    {}
func (*GetTargetTemperatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{6}
}

func (m *GetTargetTemperatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTargetTemperatureResponse.Unmarshal(m, b)
}
func (m *GetTargetTemperatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTargetTemperatureResponse.Marshal(b, m, deterministic)
}
func (m *GetTargetTemperatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTargetTemperatureResponse.Merge(m, src)
}
func (m *GetTargetTemperatureResponse) XXX_Size() int {
	return xxx_messageInfo_GetTargetTemperatureResponse.Size(m)
}
func (m *GetTargetTemperatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTargetTemperatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTargetTemperatureResponse proto.InternalMessageInfo

func (m *GetTargetTemperatureResponse) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *GetTargetTemperatureResponse) GetSetAt() *timestamp.Timestamp {
	if m != nil {
		return m.SetAt
	}
	return nil
}

type SetTargetTemperatureRequest struct {
	Temperature          float64  `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTargetTemperatureRequest) Reset()         { *m = SetTargetTemperatureRequest{} }
func (m *SetTargetTemperatureRequest) String() string { return proto.CompactTextString(m) }
func (*SetTargetTemperatureRequest) ProtoMessage()    {}
func (*SetTargetTemperatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{7}
}

func (m *SetTargetTemperatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTargetTemperatureRequest.Unmarshal(m, b)
}
func (m *SetTargetTemperatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTargetTemperatureRequest.Marshal(b, m, deterministic)
}
func (m *SetTargetTemperatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTargetTemperatureRequest.Merge(m, src)
}
func (m *SetTargetTemperatureRequest) XXX_Size() int {
	return xxx_messageInfo_SetTargetTemperatureRequest.Size(m)
}
func (m *SetTargetTemperatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTargetTemperatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTargetTemperatureRequest proto.InternalMessageInfo

func (m *SetTargetTemperatureRequest) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type SetTargetTemperatureResponse struct {
	Temperature          float64              `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	SetAt                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=set_at,json=setAt,proto3" json:"set_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SetTargetTemperatureResponse) Reset()         { *m = SetTargetTemperatureResponse{} }
func (m *SetTargetTemperatureResponse) String() string { return proto.CompactTextString(m) }
func (*SetTargetTemperatureResponse) ProtoMessage()    {}
func (*SetTargetTemperatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{8}
}

func (m *SetTargetTemperatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTargetTemperatureResponse.Unmarshal(m, b)
}
func (m *SetTargetTemperatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTargetTemperatureResponse.Marshal(b, m, deterministic)
}
func (m *SetTargetTemperatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTargetTemperatureResponse.Merge(m, src)
}
func (m *SetTargetTemperatureResponse) XXX_Size() int {
	return xxx_messageInfo_SetTargetTemperatureResponse.Size(m)
}
func (m *SetTargetTemperatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTargetTemperatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTargetTemperatureResponse proto.InternalMessageInfo

func (m *SetTargetTemperatureResponse) GetTemperature() float64 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *SetTargetTemperatureResponse) GetSetAt() *timestamp.Timestamp {
	if m != nil {
		return m.SetAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TemperatureSample)(nil), "appliancepb.TemperatureSample")
	proto.RegisterType((*GetBoilerTemperatureHistoryRequest)(nil), "appliancepb.GetBoilerTemperatureHistoryRequest")
	proto.RegisterType((*GetBoilerTemperatureHistoryResponse)(nil), "appliancepb.GetBoilerTemperatureHistoryResponse")
	proto.RegisterType((*GetCurrentBoilerTemperatureRequest)(nil), "appliancepb.GetCurrentBoilerTemperatureRequest")
	proto.RegisterType((*GetCurrentBoilerTemperatureResponse)(nil), "appliancepb.GetCurrentBoilerTemperatureResponse")
	proto.RegisterType((*GetTargetTemperatureRequest)(nil), "appliancepb.GetTargetTemperatureRequest")
	proto.RegisterType((*GetTargetTemperatureResponse)(nil), "appliancepb.GetTargetTemperatureResponse")
	proto.RegisterType((*SetTargetTemperatureRequest)(nil), "appliancepb.SetTargetTemperatureRequest")
	proto.RegisterType((*SetTargetTemperatureResponse)(nil), "appliancepb.SetTargetTemperatureResponse")
}

func init() {
	proto.RegisterFile("appliance.proto", fileDescriptor_cb8bec84a49a6ea7)
}

var fileDescriptor_cb8bec84a49a6ea7 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x3d, 0x4f, 0xeb, 0x30,
	0x14, 0x86, 0x9b, 0x5b, 0xb5, 0x57, 0xf7, 0x64, 0xb8, 0xc2, 0xea, 0x50, 0xa5, 0x7c, 0x44, 0x86,
	0x21, 0x2c, 0x29, 0x14, 0x09, 0x21, 0x31, 0xa0, 0xc2, 0x50, 0xe6, 0xa4, 0x2b, 0xaa, 0x1c, 0x38,
	0xad, 0x22, 0x25, 0xb5, 0xb1, 0x9d, 0x4a, 0xac, 0xfc, 0x65, 0xfe, 0x00, 0x22, 0x1f, 0x10, 0x48,
	0xd3, 0x64, 0x62, 0x8c, 0xfd, 0xc6, 0xcf, 0x63, 0x9f, 0x17, 0xfe, 0x33, 0x21, 0xa2, 0x90, 0xad,
	0x1f, 0xd1, 0x15, 0x92, 0x6b, 0x4e, 0xcc, 0xcf, 0x05, 0x11, 0x58, 0x47, 0x2b, 0xce, 0x57, 0x11,
	0x8e, 0xd3, 0xad, 0x20, 0x59, 0x8e, 0x75, 0x18, 0xa3, 0xd2, 0x2c, 0x16, 0x59, 0x9a, 0x2e, 0x61,
	0x6f, 0x8e, 0xb1, 0x40, 0xc9, 0x74, 0x22, 0xd1, 0x67, 0xb1, 0x88, 0x90, 0x0c, 0xa0, 0xb7, 0x61,
	0x51, 0x82, 0x43, 0xc3, 0x36, 0x1c, 0xc3, 0xcb, 0x3e, 0xc8, 0x35, 0x98, 0x3c, 0x50, 0x28, 0x37,
	0xf8, 0xb4, 0x60, 0x7a, 0xf8, 0xc7, 0x36, 0x1c, 0x73, 0x62, 0xb9, 0x19, 0xc1, 0x2d, 0x08, 0xee,
	0xbc, 0x20, 0x78, 0x50, 0xc4, 0xa7, 0x9a, 0x9e, 0x00, 0x9d, 0xa1, 0xbe, 0xe5, 0x61, 0x84, 0xb2,
	0x04, 0xbc, 0x0f, 0x95, 0xe6, 0xf2, 0xc5, 0xc3, 0xe7, 0x04, 0x95, 0xa6, 0x0b, 0x38, 0xde, 0x99,
	0x52, 0x82, 0xaf, 0x15, 0x92, 0x2b, 0xf8, 0xab, 0x52, 0x53, 0x35, 0x34, 0xec, 0xae, 0x63, 0x4e,
	0x0e, 0xdd, 0xd2, 0xa5, 0xdd, 0xca, 0x85, 0xbc, 0x22, 0x9e, 0x6b, 0xdc, 0x25, 0x52, 0xe2, 0xba,
	0xca, 0x29, 0x34, 0x1e, 0x52, 0x8d, 0xfa, 0x54, 0xae, 0x71, 0x09, 0xfd, 0xec, 0xdc, 0xf4, 0x9d,
	0x9a, 0x2d, 0xf2, 0x34, 0x3d, 0x80, 0xd1, 0x0c, 0xf5, 0x9c, 0xc9, 0x15, 0xea, 0x2d, 0x74, 0x05,
	0xfb, 0xdb, 0xb7, 0x73, 0xac, 0x0d, 0xa6, 0xfe, 0x5a, 0xce, 0x67, 0x54, 0x5e, 0x22, 0xe7, 0xd0,
	0x57, 0xa8, 0xdb, 0x0d, 0xa9, 0xa7, 0x50, 0x4f, 0x35, 0xbd, 0x81, 0x91, 0x5f, 0xef, 0xd4, 0xcc,
	0xfc, 0xb0, 0xf6, 0x7f, 0xdb, 0x7a, 0xf2, 0xd6, 0x85, 0x7f, 0xd3, 0xe2, 0xcd, 0xc9, 0xab, 0x91,
	0x3e, 0x6c, 0x5d, 0x7d, 0xc8, 0xf8, 0xdb, 0x7c, 0x9a, 0xeb, 0x68, 0x9d, 0xb5, 0xff, 0x21, 0xbb,
	0x25, 0xed, 0x14, 0x12, 0x75, 0xe5, 0xa9, 0x4a, 0x34, 0x94, 0xb1, 0x2a, 0xd1, 0xd4, 0x4b, 0xda,
	0x21, 0x31, 0x0c, 0xb6, 0x55, 0x88, 0x38, 0x3f, 0xcf, 0xaa, 0x1b, 0xb8, 0x75, 0xda, 0x22, 0x59,
	0xc6, 0xf9, 0xcd, 0x38, 0xbf, 0x35, 0xce, 0xdf, 0x89, 0x0b, 0xfa, 0x69, 0x21, 0x2e, 0xde, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x65, 0xd9, 0x66, 0xca, 0xfb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApplianceClient is the client API for Appliance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplianceClient interface {
	GetBoilerTemperatureHistory(ctx context.Context, in *GetBoilerTemperatureHistoryRequest, opts ...grpc.CallOption) (*GetBoilerTemperatureHistoryResponse, error)
	GetCurrentBoilerTemperature(ctx context.Context, in *GetCurrentBoilerTemperatureRequest, opts ...grpc.CallOption) (*GetCurrentBoilerTemperatureResponse, error)
	GetTargetTemperature(ctx context.Context, in *GetTargetTemperatureRequest, opts ...grpc.CallOption) (*GetTargetTemperatureResponse, error)
	SetTargetTemperature(ctx context.Context, in *SetTargetTemperatureRequest, opts ...grpc.CallOption) (*SetTargetTemperatureResponse, error)
}

type applianceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplianceClient(cc grpc.ClientConnInterface) ApplianceClient {
	return &applianceClient{cc}
}

func (c *applianceClient) GetBoilerTemperatureHistory(ctx context.Context, in *GetBoilerTemperatureHistoryRequest, opts ...grpc.CallOption) (*GetBoilerTemperatureHistoryResponse, error) {
	out := new(GetBoilerTemperatureHistoryResponse)
	err := c.cc.Invoke(ctx, "/appliancepb.Appliance/GetBoilerTemperatureHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applianceClient) GetCurrentBoilerTemperature(ctx context.Context, in *GetCurrentBoilerTemperatureRequest, opts ...grpc.CallOption) (*GetCurrentBoilerTemperatureResponse, error) {
	out := new(GetCurrentBoilerTemperatureResponse)
	err := c.cc.Invoke(ctx, "/appliancepb.Appliance/GetCurrentBoilerTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applianceClient) GetTargetTemperature(ctx context.Context, in *GetTargetTemperatureRequest, opts ...grpc.CallOption) (*GetTargetTemperatureResponse, error) {
	out := new(GetTargetTemperatureResponse)
	err := c.cc.Invoke(ctx, "/appliancepb.Appliance/GetTargetTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applianceClient) SetTargetTemperature(ctx context.Context, in *SetTargetTemperatureRequest, opts ...grpc.CallOption) (*SetTargetTemperatureResponse, error) {
	out := new(SetTargetTemperatureResponse)
	err := c.cc.Invoke(ctx, "/appliancepb.Appliance/SetTargetTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplianceServer is the server API for Appliance service.
type ApplianceServer interface {
	GetBoilerTemperatureHistory(context.Context, *GetBoilerTemperatureHistoryRequest) (*GetBoilerTemperatureHistoryResponse, error)
	GetCurrentBoilerTemperature(context.Context, *GetCurrentBoilerTemperatureRequest) (*GetCurrentBoilerTemperatureResponse, error)
	GetTargetTemperature(context.Context, *GetTargetTemperatureRequest) (*GetTargetTemperatureResponse, error)
	SetTargetTemperature(context.Context, *SetTargetTemperatureRequest) (*SetTargetTemperatureResponse, error)
}

// UnimplementedApplianceServer can be embedded to have forward compatible implementations.
type UnimplementedApplianceServer struct {
}

func (*UnimplementedApplianceServer) GetBoilerTemperatureHistory(ctx context.Context, req *GetBoilerTemperatureHistoryRequest) (*GetBoilerTemperatureHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoilerTemperatureHistory not implemented")
}
func (*UnimplementedApplianceServer) GetCurrentBoilerTemperature(ctx context.Context, req *GetCurrentBoilerTemperatureRequest) (*GetCurrentBoilerTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBoilerTemperature not implemented")
}
func (*UnimplementedApplianceServer) GetTargetTemperature(ctx context.Context, req *GetTargetTemperatureRequest) (*GetTargetTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetTemperature not implemented")
}
func (*UnimplementedApplianceServer) SetTargetTemperature(ctx context.Context, req *SetTargetTemperatureRequest) (*SetTargetTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTargetTemperature not implemented")
}

func RegisterApplianceServer(s *grpc.Server, srv ApplianceServer) {
	s.RegisterService(&_Appliance_serviceDesc, srv)
}

func _Appliance_GetBoilerTemperatureHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoilerTemperatureHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplianceServer).GetBoilerTemperatureHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appliancepb.Appliance/GetBoilerTemperatureHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplianceServer).GetBoilerTemperatureHistory(ctx, req.(*GetBoilerTemperatureHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appliance_GetCurrentBoilerTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentBoilerTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplianceServer).GetCurrentBoilerTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appliancepb.Appliance/GetCurrentBoilerTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplianceServer).GetCurrentBoilerTemperature(ctx, req.(*GetCurrentBoilerTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appliance_GetTargetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplianceServer).GetTargetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appliancepb.Appliance/GetTargetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplianceServer).GetTargetTemperature(ctx, req.(*GetTargetTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Appliance_SetTargetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTargetTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplianceServer).SetTargetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appliancepb.Appliance/SetTargetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplianceServer).SetTargetTemperature(ctx, req.(*SetTargetTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Appliance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appliancepb.Appliance",
	HandlerType: (*ApplianceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBoilerTemperatureHistory",
			Handler:    _Appliance_GetBoilerTemperatureHistory_Handler,
		},
		{
			MethodName: "GetCurrentBoilerTemperature",
			Handler:    _Appliance_GetCurrentBoilerTemperature_Handler,
		},
		{
			MethodName: "GetTargetTemperature",
			Handler:    _Appliance_GetTargetTemperature_Handler,
		},
		{
			MethodName: "SetTargetTemperature",
			Handler:    _Appliance_SetTargetTemperature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appliance.proto",
}
