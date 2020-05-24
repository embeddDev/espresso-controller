// Code generated by protoc-gen-go. DO NOT EDIT.
// source: espresso.proto

package espressopb

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
	return fileDescriptor_445399412d1702d2, []int{0}
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

type TemperatureHistory struct {
	Samples              []*TemperatureSample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TemperatureHistory) Reset()         { *m = TemperatureHistory{} }
func (m *TemperatureHistory) String() string { return proto.CompactTextString(m) }
func (*TemperatureHistory) ProtoMessage()    {}
func (*TemperatureHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{1}
}

func (m *TemperatureHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureHistory.Unmarshal(m, b)
}
func (m *TemperatureHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureHistory.Marshal(b, m, deterministic)
}
func (m *TemperatureHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureHistory.Merge(m, src)
}
func (m *TemperatureHistory) XXX_Size() int {
	return xxx_messageInfo_TemperatureHistory.Size(m)
}
func (m *TemperatureHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureHistory proto.InternalMessageInfo

func (m *TemperatureHistory) GetSamples() []*TemperatureSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type TemperatureStreamRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemperatureStreamRequest) Reset()         { *m = TemperatureStreamRequest{} }
func (m *TemperatureStreamRequest) String() string { return proto.CompactTextString(m) }
func (*TemperatureStreamRequest) ProtoMessage()    {}
func (*TemperatureStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{2}
}

func (m *TemperatureStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureStreamRequest.Unmarshal(m, b)
}
func (m *TemperatureStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureStreamRequest.Marshal(b, m, deterministic)
}
func (m *TemperatureStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureStreamRequest.Merge(m, src)
}
func (m *TemperatureStreamRequest) XXX_Size() int {
	return xxx_messageInfo_TemperatureStreamRequest.Size(m)
}
func (m *TemperatureStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureStreamRequest proto.InternalMessageInfo

type TemperatureStreamResponse struct {
	// Types that are valid to be assigned to Data:
	//	*TemperatureStreamResponse_History
	//	*TemperatureStreamResponse_Sample
	Data                 isTemperatureStreamResponse_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *TemperatureStreamResponse) Reset()         { *m = TemperatureStreamResponse{} }
func (m *TemperatureStreamResponse) String() string { return proto.CompactTextString(m) }
func (*TemperatureStreamResponse) ProtoMessage()    {}
func (*TemperatureStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_445399412d1702d2, []int{3}
}

func (m *TemperatureStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureStreamResponse.Unmarshal(m, b)
}
func (m *TemperatureStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureStreamResponse.Marshal(b, m, deterministic)
}
func (m *TemperatureStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureStreamResponse.Merge(m, src)
}
func (m *TemperatureStreamResponse) XXX_Size() int {
	return xxx_messageInfo_TemperatureStreamResponse.Size(m)
}
func (m *TemperatureStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureStreamResponse proto.InternalMessageInfo

type isTemperatureStreamResponse_Data interface {
	isTemperatureStreamResponse_Data()
}

type TemperatureStreamResponse_History struct {
	History *TemperatureHistory `protobuf:"bytes,1,opt,name=history,proto3,oneof"`
}

type TemperatureStreamResponse_Sample struct {
	Sample *TemperatureSample `protobuf:"bytes,2,opt,name=sample,proto3,oneof"`
}

func (*TemperatureStreamResponse_History) isTemperatureStreamResponse_Data() {}

func (*TemperatureStreamResponse_Sample) isTemperatureStreamResponse_Data() {}

func (m *TemperatureStreamResponse) GetData() isTemperatureStreamResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TemperatureStreamResponse) GetHistory() *TemperatureHistory {
	if x, ok := m.GetData().(*TemperatureStreamResponse_History); ok {
		return x.History
	}
	return nil
}

func (m *TemperatureStreamResponse) GetSample() *TemperatureSample {
	if x, ok := m.GetData().(*TemperatureStreamResponse_Sample); ok {
		return x.Sample
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TemperatureStreamResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TemperatureStreamResponse_History)(nil),
		(*TemperatureStreamResponse_Sample)(nil),
	}
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
	return fileDescriptor_445399412d1702d2, []int{4}
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
	return fileDescriptor_445399412d1702d2, []int{5}
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
	return fileDescriptor_445399412d1702d2, []int{6}
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
	return fileDescriptor_445399412d1702d2, []int{7}
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
	proto.RegisterType((*TemperatureSample)(nil), "espressopb.TemperatureSample")
	proto.RegisterType((*TemperatureHistory)(nil), "espressopb.TemperatureHistory")
	proto.RegisterType((*TemperatureStreamRequest)(nil), "espressopb.TemperatureStreamRequest")
	proto.RegisterType((*TemperatureStreamResponse)(nil), "espressopb.TemperatureStreamResponse")
	proto.RegisterType((*GetTargetTemperatureRequest)(nil), "espressopb.GetTargetTemperatureRequest")
	proto.RegisterType((*GetTargetTemperatureResponse)(nil), "espressopb.GetTargetTemperatureResponse")
	proto.RegisterType((*SetTargetTemperatureRequest)(nil), "espressopb.SetTargetTemperatureRequest")
	proto.RegisterType((*SetTargetTemperatureResponse)(nil), "espressopb.SetTargetTemperatureResponse")
}

func init() {
	proto.RegisterFile("espresso.proto", fileDescriptor_445399412d1702d2)
}

var fileDescriptor_445399412d1702d2 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0x87, 0xa9, 0x48, 0x31, 0xa7, 0x89, 0x09, 0x13, 0x16, 0xb5, 0x88, 0x36, 0x8d, 0xc6, 0xae,
	0x8a, 0xd6, 0x05, 0x89, 0x2e, 0x0c, 0x24, 0x86, 0x6e, 0xdc, 0xb4, 0xec, 0xcd, 0x34, 0x1c, 0x90,
	0xd8, 0x3a, 0xe3, 0xcc, 0x29, 0x89, 0x6f, 0xe2, 0x5b, 0xdc, 0x57, 0xbc, 0xa1, 0x7f, 0x42, 0x93,
	0x0b, 0x94, 0xcd, 0xbd, 0xcb, 0x99, 0xf3, 0x6b, 0xbf, 0xef, 0xcc, 0x39, 0xf0, 0x12, 0xb5, 0x54,
	0xa8, 0xb5, 0x08, 0xa4, 0x12, 0x24, 0x18, 0x34, 0x67, 0x99, 0x3a, 0x6f, 0x77, 0x42, 0xec, 0x32,
	0x9c, 0x95, 0x95, 0xb4, 0xd8, 0xce, 0x68, 0x9f, 0xa3, 0x26, 0x9e, 0xcb, 0x2a, 0xec, 0x6d, 0x61,
	0xb4, 0xc6, 0x5c, 0xa2, 0xe2, 0x54, 0x28, 0x4c, 0x78, 0x2e, 0x33, 0x64, 0x63, 0x18, 0x1c, 0x78,
	0x56, 0xa0, 0x6d, 0xb8, 0x86, 0x6f, 0xc4, 0xd5, 0x81, 0x7d, 0x05, 0x4b, 0xa4, 0x1a, 0xd5, 0x01,
	0x37, 0x3f, 0x39, 0xd9, 0xcf, 0x5c, 0xc3, 0xb7, 0x42, 0x27, 0xa8, 0x08, 0x41, 0x43, 0x08, 0xd6,
	0x0d, 0x21, 0x86, 0x26, 0xbe, 0x20, 0xef, 0x07, 0xb0, 0x16, 0x27, 0xda, 0x6b, 0x12, 0xea, 0x1f,
	0x9b, 0xc3, 0x50, 0x97, 0x48, 0x6d, 0x1b, 0x6e, 0xdf, 0xb7, 0xc2, 0x69, 0x70, 0x92, 0x0f, 0x1e,
	0x88, 0xc5, 0x4d, 0xda, 0x73, 0xc0, 0x6e, 0x57, 0x49, 0x21, 0xcf, 0x63, 0xfc, 0x5b, 0xa0, 0x26,
	0xef, 0xbf, 0x01, 0xaf, 0xce, 0x14, 0xb5, 0x14, 0x7f, 0x34, 0xb2, 0x2f, 0x30, 0xfc, 0x55, 0xd1,
	0xcb, 0xee, 0xac, 0xf0, 0xcd, 0x05, 0x64, 0xed, 0x18, 0xf5, 0xe2, 0xe6, 0x03, 0x36, 0x07, 0xb3,
	0x12, 0xa8, 0x9b, 0xbf, 0x6e, 0x1b, 0xf5, 0xe2, 0x3a, 0xbe, 0x34, 0xe1, 0xf9, 0x86, 0x13, 0xf7,
	0xa6, 0x30, 0x59, 0x21, 0xad, 0xb9, 0xda, 0x21, 0xb5, 0xf2, 0x8d, 0xb9, 0x86, 0xd7, 0xe7, 0xcb,
	0xb5, 0xbb, 0x0b, 0x16, 0x9d, 0xae, 0xeb, 0xe9, 0xb4, 0xaf, 0xd8, 0x27, 0x30, 0x35, 0xd2, 0x6d,
	0xe3, 0x19, 0x68, 0xa4, 0x05, 0x79, 0xdf, 0x60, 0x92, 0x5c, 0x76, 0xea, 0x66, 0x1e, 0xad, 0x93,
	0xa7, 0xb6, 0x0e, 0xef, 0xfa, 0xf0, 0xe2, 0x7b, 0xfd, 0xf8, 0x0c, 0x61, 0xbc, 0x52, 0xa2, 0x90,
	0x11, 0xf2, 0x4d, 0xcb, 0x80, 0xbd, 0xbb, 0x34, 0x9f, 0xf6, 0xbe, 0x38, 0xef, 0x3b, 0x52, 0x55,
	0x1b, 0x1f, 0x0d, 0x96, 0xc2, 0x68, 0x29, 0xf6, 0x19, 0xaa, 0x47, 0x64, 0xfc, 0x86, 0xf1, 0xb9,
	0x15, 0x60, 0x1f, 0xda, 0x3f, 0xb8, 0xb2, 0x43, 0x8e, 0xdf, 0x1d, 0xac, 0x60, 0x5e, 0xef, 0x08,
	0x4b, 0x3a, 0x61, 0xc9, 0xad, 0xb0, 0xe4, 0x2a, 0x2c, 0x35, 0xcb, 0x61, 0x7e, 0xbe, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x82, 0x5e, 0x24, 0x9b, 0xaf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EspressoClient is the client API for Espresso service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EspressoClient interface {
	GroupHeadTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_GroupHeadTemperatureClient, error)
	BoilerTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_BoilerTemperatureClient, error)
	GetTargetTemperature(ctx context.Context, in *GetTargetTemperatureRequest, opts ...grpc.CallOption) (*GetTargetTemperatureResponse, error)
	SetTargetTemperature(ctx context.Context, in *SetTargetTemperatureRequest, opts ...grpc.CallOption) (*SetTargetTemperatureResponse, error)
}

type espressoClient struct {
	cc grpc.ClientConnInterface
}

func NewEspressoClient(cc grpc.ClientConnInterface) EspressoClient {
	return &espressoClient{cc}
}

func (c *espressoClient) GroupHeadTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_GroupHeadTemperatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Espresso_serviceDesc.Streams[0], "/espressopb.Espresso/GroupHeadTemperature", opts...)
	if err != nil {
		return nil, err
	}
	x := &espressoGroupHeadTemperatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Espresso_GroupHeadTemperatureClient interface {
	Recv() (*TemperatureStreamResponse, error)
	grpc.ClientStream
}

type espressoGroupHeadTemperatureClient struct {
	grpc.ClientStream
}

func (x *espressoGroupHeadTemperatureClient) Recv() (*TemperatureStreamResponse, error) {
	m := new(TemperatureStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *espressoClient) BoilerTemperature(ctx context.Context, in *TemperatureStreamRequest, opts ...grpc.CallOption) (Espresso_BoilerTemperatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Espresso_serviceDesc.Streams[1], "/espressopb.Espresso/BoilerTemperature", opts...)
	if err != nil {
		return nil, err
	}
	x := &espressoBoilerTemperatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Espresso_BoilerTemperatureClient interface {
	Recv() (*TemperatureStreamResponse, error)
	grpc.ClientStream
}

type espressoBoilerTemperatureClient struct {
	grpc.ClientStream
}

func (x *espressoBoilerTemperatureClient) Recv() (*TemperatureStreamResponse, error) {
	m := new(TemperatureStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *espressoClient) GetTargetTemperature(ctx context.Context, in *GetTargetTemperatureRequest, opts ...grpc.CallOption) (*GetTargetTemperatureResponse, error) {
	out := new(GetTargetTemperatureResponse)
	err := c.cc.Invoke(ctx, "/espressopb.Espresso/GetTargetTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *espressoClient) SetTargetTemperature(ctx context.Context, in *SetTargetTemperatureRequest, opts ...grpc.CallOption) (*SetTargetTemperatureResponse, error) {
	out := new(SetTargetTemperatureResponse)
	err := c.cc.Invoke(ctx, "/espressopb.Espresso/SetTargetTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EspressoServer is the server API for Espresso service.
type EspressoServer interface {
	GroupHeadTemperature(*TemperatureStreamRequest, Espresso_GroupHeadTemperatureServer) error
	BoilerTemperature(*TemperatureStreamRequest, Espresso_BoilerTemperatureServer) error
	GetTargetTemperature(context.Context, *GetTargetTemperatureRequest) (*GetTargetTemperatureResponse, error)
	SetTargetTemperature(context.Context, *SetTargetTemperatureRequest) (*SetTargetTemperatureResponse, error)
}

// UnimplementedEspressoServer can be embedded to have forward compatible implementations.
type UnimplementedEspressoServer struct {
}

func (*UnimplementedEspressoServer) GroupHeadTemperature(req *TemperatureStreamRequest, srv Espresso_GroupHeadTemperatureServer) error {
	return status.Errorf(codes.Unimplemented, "method GroupHeadTemperature not implemented")
}
func (*UnimplementedEspressoServer) BoilerTemperature(req *TemperatureStreamRequest, srv Espresso_BoilerTemperatureServer) error {
	return status.Errorf(codes.Unimplemented, "method BoilerTemperature not implemented")
}
func (*UnimplementedEspressoServer) GetTargetTemperature(ctx context.Context, req *GetTargetTemperatureRequest) (*GetTargetTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetTemperature not implemented")
}
func (*UnimplementedEspressoServer) SetTargetTemperature(ctx context.Context, req *SetTargetTemperatureRequest) (*SetTargetTemperatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTargetTemperature not implemented")
}

func RegisterEspressoServer(s *grpc.Server, srv EspressoServer) {
	s.RegisterService(&_Espresso_serviceDesc, srv)
}

func _Espresso_GroupHeadTemperature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TemperatureStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EspressoServer).GroupHeadTemperature(m, &espressoGroupHeadTemperatureServer{stream})
}

type Espresso_GroupHeadTemperatureServer interface {
	Send(*TemperatureStreamResponse) error
	grpc.ServerStream
}

type espressoGroupHeadTemperatureServer struct {
	grpc.ServerStream
}

func (x *espressoGroupHeadTemperatureServer) Send(m *TemperatureStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Espresso_BoilerTemperature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TemperatureStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EspressoServer).BoilerTemperature(m, &espressoBoilerTemperatureServer{stream})
}

type Espresso_BoilerTemperatureServer interface {
	Send(*TemperatureStreamResponse) error
	grpc.ServerStream
}

type espressoBoilerTemperatureServer struct {
	grpc.ServerStream
}

func (x *espressoBoilerTemperatureServer) Send(m *TemperatureStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Espresso_GetTargetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EspressoServer).GetTargetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/espressopb.Espresso/GetTargetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EspressoServer).GetTargetTemperature(ctx, req.(*GetTargetTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Espresso_SetTargetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTargetTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EspressoServer).SetTargetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/espressopb.Espresso/SetTargetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EspressoServer).SetTargetTemperature(ctx, req.(*SetTargetTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Espresso_serviceDesc = grpc.ServiceDesc{
	ServiceName: "espressopb.Espresso",
	HandlerType: (*EspressoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTargetTemperature",
			Handler:    _Espresso_GetTargetTemperature_Handler,
		},
		{
			MethodName: "SetTargetTemperature",
			Handler:    _Espresso_SetTargetTemperature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GroupHeadTemperature",
			Handler:       _Espresso_GroupHeadTemperature_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BoilerTemperature",
			Handler:       _Espresso_BoilerTemperature_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "espresso.proto",
}