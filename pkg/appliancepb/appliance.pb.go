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

type GetCurrentTemperatureRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentTemperatureRequest) Reset()         { *m = GetCurrentTemperatureRequest{} }
func (m *GetCurrentTemperatureRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentTemperatureRequest) ProtoMessage()    {}
func (*GetCurrentTemperatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{3}
}

func (m *GetCurrentTemperatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentTemperatureRequest.Unmarshal(m, b)
}
func (m *GetCurrentTemperatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentTemperatureRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentTemperatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentTemperatureRequest.Merge(m, src)
}
func (m *GetCurrentTemperatureRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentTemperatureRequest.Size(m)
}
func (m *GetCurrentTemperatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentTemperatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentTemperatureRequest proto.InternalMessageInfo

type GetCurrentTemperatureResponse struct {
	Sample               *TemperatureSample `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetCurrentTemperatureResponse) Reset()         { *m = GetCurrentTemperatureResponse{} }
func (m *GetCurrentTemperatureResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentTemperatureResponse) ProtoMessage()    {}
func (*GetCurrentTemperatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8bec84a49a6ea7, []int{4}
}

func (m *GetCurrentTemperatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentTemperatureResponse.Unmarshal(m, b)
}
func (m *GetCurrentTemperatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentTemperatureResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentTemperatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentTemperatureResponse.Merge(m, src)
}
func (m *GetCurrentTemperatureResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentTemperatureResponse.Size(m)
}
func (m *GetCurrentTemperatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentTemperatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentTemperatureResponse proto.InternalMessageInfo

func (m *GetCurrentTemperatureResponse) GetSample() *TemperatureSample {
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
	proto.RegisterType((*GetCurrentTemperatureRequest)(nil), "appliancepb.GetCurrentTemperatureRequest")
	proto.RegisterType((*GetCurrentTemperatureResponse)(nil), "appliancepb.GetCurrentTemperatureResponse")
	proto.RegisterType((*GetTargetTemperatureRequest)(nil), "appliancepb.GetTargetTemperatureRequest")
	proto.RegisterType((*GetTargetTemperatureResponse)(nil), "appliancepb.GetTargetTemperatureResponse")
	proto.RegisterType((*SetTargetTemperatureRequest)(nil), "appliancepb.SetTargetTemperatureRequest")
	proto.RegisterType((*SetTargetTemperatureResponse)(nil), "appliancepb.SetTargetTemperatureResponse")
}

func init() {
	proto.RegisterFile("appliance.proto", fileDescriptor_cb8bec84a49a6ea7)
}

var fileDescriptor_cb8bec84a49a6ea7 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0xab, 0xd3, 0x40,
	0x10, 0xc7, 0x8d, 0xa5, 0x15, 0x27, 0x07, 0x71, 0xa9, 0x50, 0x52, 0x5b, 0xc3, 0xea, 0x21, 0xf5,
	0x90, 0x6a, 0x05, 0x11, 0x3c, 0x48, 0xf5, 0x50, 0xcf, 0x49, 0xc1, 0x63, 0xd9, 0xe8, 0xb4, 0x04,
	0x92, 0xee, 0xba, 0x3b, 0x29, 0x78, 0xf5, 0x2f, 0xf5, 0x4f, 0x91, 0x97, 0x1f, 0x7d, 0x79, 0x7d,
	0x49, 0x93, 0xd3, 0x3b, 0x66, 0xe7, 0x9b, 0xf9, 0x7c, 0x67, 0xe6, 0x0b, 0xcf, 0x84, 0x52, 0x49,
	0x2c, 0x8e, 0x3f, 0xd1, 0x57, 0x5a, 0x92, 0x64, 0xf6, 0xf9, 0x41, 0x45, 0xce, 0xab, 0x83, 0x94,
	0x87, 0x04, 0x97, 0x79, 0x29, 0xca, 0xf6, 0x4b, 0x8a, 0x53, 0x34, 0x24, 0x52, 0x55, 0xa8, 0xf9,
	0x1e, 0x9e, 0x6f, 0x31, 0x55, 0xa8, 0x05, 0x65, 0x1a, 0x43, 0x91, 0xaa, 0x04, 0xd9, 0x18, 0x86,
	0x27, 0x91, 0x64, 0x38, 0xb1, 0x5c, 0xcb, 0xb3, 0x82, 0xe2, 0x83, 0x7d, 0x06, 0x5b, 0x46, 0x06,
	0xf5, 0x09, 0x7f, 0xed, 0x04, 0x4d, 0x1e, 0xbb, 0x96, 0x67, 0xaf, 0x1c, 0xbf, 0x20, 0xf8, 0x15,
	0xc1, 0xdf, 0x56, 0x84, 0x00, 0x2a, 0xf9, 0x9a, 0xf8, 0x1b, 0xe0, 0x1b, 0xa4, 0xaf, 0x32, 0x4e,
	0x50, 0xd7, 0x80, 0xdf, 0x63, 0x43, 0x52, 0xff, 0x09, 0xf0, 0x77, 0x86, 0x86, 0xf8, 0x0e, 0x5e,
	0x5f, 0x55, 0x19, 0x25, 0x8f, 0x06, 0xd9, 0x27, 0x78, 0x62, 0x72, 0xa7, 0x66, 0x62, 0xb9, 0x03,
	0xcf, 0x5e, 0xcd, 0xfd, 0xda, 0xd0, 0xfe, 0xbd, 0x81, 0x82, 0x4a, 0xce, 0xe7, 0xf0, 0x72, 0x83,
	0xf4, 0x2d, 0xd3, 0x1a, 0x8f, 0x54, 0xd3, 0x55, 0x06, 0x7e, 0xc0, 0xac, 0xa5, 0x5e, 0xa2, 0x3f,
	0xc2, 0xa8, 0xe8, 0x95, 0xef, 0xa6, 0x9b, 0x5c, 0xaa, 0xf9, 0x0c, 0xa6, 0x1b, 0xa4, 0xad, 0xd0,
	0x07, 0x6c, 0xe2, 0x9a, 0xdc, 0x57, 0x43, 0xb9, 0xc4, 0xba, 0x60, 0xd3, 0xed, 0x73, 0x79, 0x97,
	0xfa, 0x13, 0x7b, 0x0f, 0x23, 0x83, 0xd4, 0xef, 0x30, 0x43, 0x83, 0xb4, 0x26, 0xfe, 0x05, 0xa6,
	0x61, 0xbb, 0xa7, 0x6e, 0xe6, 0x8d, 0xeb, 0xf0, 0xa1, 0x5d, 0xaf, 0xfe, 0x0d, 0xe0, 0xe9, 0xba,
	0xda, 0x39, 0xfb, 0x6b, 0xe5, 0x8b, 0x6d, 0x8b, 0x0c, 0x5b, 0xde, 0xb9, 0x4f, 0x77, 0x04, 0x9d,
	0x77, 0xfd, 0x7f, 0x28, 0xa6, 0xe4, 0x8f, 0x98, 0x82, 0x17, 0x8d, 0xa9, 0x61, 0x8b, 0xcb, 0x66,
	0xad, 0xc9, 0x73, 0xde, 0xf6, 0x91, 0x9e, 0x89, 0x29, 0x8c, 0x9b, 0xf2, 0xc2, 0xbc, 0xcb, 0x2e,
	0x6d, 0xd7, 0x75, 0x16, 0x3d, 0x94, 0x75, 0x5c, 0xd8, 0x8d, 0x0b, 0x7b, 0xe3, 0xc2, 0xab, 0xb8,
	0x68, 0x94, 0x5f, 0xff, 0xc3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x08, 0x0b, 0x3d, 0xdc,
	0x04, 0x00, 0x00,
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
	GetCurrentTemperature(ctx context.Context, in *GetCurrentTemperatureRequest, opts ...grpc.CallOption) (*GetCurrentTemperatureResponse, error)
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

func (c *applianceClient) GetCurrentTemperature(ctx context.Context, in *GetCurrentTemperatureRequest, opts ...grpc.CallOption) (*GetCurrentTemperatureResponse, error) {
	out := new(GetCurrentTemperatureResponse)
	err := c.cc.Invoke(ctx, "/appliancepb.Appliance/GetCurrentTemperature", in, out, opts...)
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
	GetCurrentTemperature(context.Context, *GetCurrentTemperatureRequest) (*GetCurrentTemperatureResponse, error)
	GetTargetTemperature(context.Context, *GetTargetTemperatureRequest) (*GetTargetTemperatureResponse, error)
	SetTargetTemperature(context.Context, *SetTargetTemperatureRequest) (*SetTargetTemperatureResponse, error)
}

// UnimplementedApplianceServer can be embedded to have forward compatible implementations.
type UnimplementedApplianceServer struct {
}

func (*UnimplementedApplianceServer) GetBoilerTemperatureHistory(ctx context.Context, req *GetBoilerTemperatureHistoryRequest) (*GetBoilerTemperatureHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoilerTemperatureHistory not implemented")
}
func (*UnimplementedApplianceServer) GetCurrentTemperature(ctx context.Context, req *GetCurrentTemperatureRequest) (*GetCurrentTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTemperature not implemented")
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

func _Appliance_GetCurrentTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplianceServer).GetCurrentTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appliancepb.Appliance/GetCurrentTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplianceServer).GetCurrentTemperature(ctx, req.(*GetCurrentTemperatureRequest))
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
			MethodName: "GetCurrentTemperature",
			Handler:    _Appliance_GetCurrentTemperature_Handler,
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
