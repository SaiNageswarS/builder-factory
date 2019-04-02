// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ComponentMgmt.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Package_RepoProvider int32

const (
	Package_BITBUCKET Package_RepoProvider = 0
	Package_GITHUB    Package_RepoProvider = 1
)

var Package_RepoProvider_name = map[int32]string{
	0: "BITBUCKET",
	1: "GITHUB",
}

var Package_RepoProvider_value = map[string]int32{
	"BITBUCKET": 0,
	"GITHUB":    1,
}

func (x Package_RepoProvider) String() string {
	return proto.EnumName(Package_RepoProvider_name, int32(x))
}

func (Package_RepoProvider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{0, 0}
}

type ComponentDetail_ComponentType int32

const (
	ComponentDetail_FRONTEND_WEB      ComponentDetail_ComponentType = 0
	ComponentDetail_BACKEND_API_REST  ComponentDetail_ComponentType = 1
	ComponentDetail_BACKEND_API_RPC   ComponentDetail_ComponentType = 2
	ComponentDetail_SCHEDULED_JOB     ComponentDetail_ComponentType = 3
	ComponentDetail_WORKFLOW          ComponentDetail_ComponentType = 4
	ComponentDetail_MAPREDUCE_CLUSTER ComponentDetail_ComponentType = 5
)

var ComponentDetail_ComponentType_name = map[int32]string{
	0: "FRONTEND_WEB",
	1: "BACKEND_API_REST",
	2: "BACKEND_API_RPC",
	3: "SCHEDULED_JOB",
	4: "WORKFLOW",
	5: "MAPREDUCE_CLUSTER",
}

var ComponentDetail_ComponentType_value = map[string]int32{
	"FRONTEND_WEB":      0,
	"BACKEND_API_REST":  1,
	"BACKEND_API_RPC":   2,
	"SCHEDULED_JOB":     3,
	"WORKFLOW":          4,
	"MAPREDUCE_CLUSTER": 5,
}

func (x ComponentDetail_ComponentType) String() string {
	return proto.EnumName(ComponentDetail_ComponentType_name, int32(x))
}

func (ComponentDetail_ComponentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{1, 0}
}

type ComponentDetail_PipelineState int32

const (
	ComponentDetail_DEPLOYED       ComponentDetail_PipelineState = 0
	ComponentDetail_NOT_CONFIGURED ComponentDetail_PipelineState = 1
	ComponentDetail_BROKEN         ComponentDetail_PipelineState = 2
)

var ComponentDetail_PipelineState_name = map[int32]string{
	0: "DEPLOYED",
	1: "NOT_CONFIGURED",
	2: "BROKEN",
}

var ComponentDetail_PipelineState_value = map[string]int32{
	"DEPLOYED":       0,
	"NOT_CONFIGURED": 1,
	"BROKEN":         2,
}

func (x ComponentDetail_PipelineState) String() string {
	return proto.EnumName(ComponentDetail_PipelineState_name, int32(x))
}

func (ComponentDetail_PipelineState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{1, 1}
}

type Package struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RepoUrl              string               `protobuf:"bytes,2,opt,name=repoUrl,proto3" json:"repoUrl,omitempty"`
	Language             string               `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	RepoProvider         Package_RepoProvider `protobuf:"varint,4,opt,name=repoProvider,proto3,enum=services.Package_RepoProvider" json:"repoProvider,omitempty"`
	RepoWebUrl           string               `protobuf:"bytes,5,opt,name=repoWebUrl,proto3" json:"repoWebUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{0}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *Package) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Package) GetRepoProvider() Package_RepoProvider {
	if m != nil {
		return m.RepoProvider
	}
	return Package_BITBUCKET
}

func (m *Package) GetRepoWebUrl() string {
	if m != nil {
		return m.RepoWebUrl
	}
	return ""
}

type ComponentDetail struct {
	ComponentName        string                        `protobuf:"bytes,3,opt,name=componentName,proto3" json:"componentName,omitempty"`
	ApplicationName      string                        `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	ComponentType        ComponentDetail_ComponentType `protobuf:"varint,2,opt,name=componentType,proto3,enum=services.ComponentDetail_ComponentType" json:"componentType,omitempty"`
	OrgName              string                        `protobuf:"bytes,4,opt,name=orgName,proto3" json:"orgName,omitempty"`
	Package              *Package                      `protobuf:"bytes,5,opt,name=package,proto3" json:"package,omitempty"`
	PipelineState        ComponentDetail_PipelineState `protobuf:"varint,6,opt,name=pipelineState,proto3,enum=services.ComponentDetail_PipelineState" json:"pipelineState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ComponentDetail) Reset()         { *m = ComponentDetail{} }
func (m *ComponentDetail) String() string { return proto.CompactTextString(m) }
func (*ComponentDetail) ProtoMessage()    {}
func (*ComponentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{1}
}

func (m *ComponentDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentDetail.Unmarshal(m, b)
}
func (m *ComponentDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentDetail.Marshal(b, m, deterministic)
}
func (m *ComponentDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentDetail.Merge(m, src)
}
func (m *ComponentDetail) XXX_Size() int {
	return xxx_messageInfo_ComponentDetail.Size(m)
}
func (m *ComponentDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentDetail proto.InternalMessageInfo

func (m *ComponentDetail) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

func (m *ComponentDetail) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ComponentDetail) GetComponentType() ComponentDetail_ComponentType {
	if m != nil {
		return m.ComponentType
	}
	return ComponentDetail_FRONTEND_WEB
}

func (m *ComponentDetail) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *ComponentDetail) GetPackage() *Package {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *ComponentDetail) GetPipelineState() ComponentDetail_PipelineState {
	if m != nil {
		return m.PipelineState
	}
	return ComponentDetail_DEPLOYED
}

type ComponentCreateResponse struct {
	ComponentName        string   `protobuf:"bytes,1,opt,name=componentName,proto3" json:"componentName,omitempty"`
	ComponentId          int64    `protobuf:"varint,2,opt,name=componentId,proto3" json:"componentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentCreateResponse) Reset()         { *m = ComponentCreateResponse{} }
func (m *ComponentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ComponentCreateResponse) ProtoMessage()    {}
func (*ComponentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{2}
}

func (m *ComponentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentCreateResponse.Unmarshal(m, b)
}
func (m *ComponentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentCreateResponse.Marshal(b, m, deterministic)
}
func (m *ComponentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentCreateResponse.Merge(m, src)
}
func (m *ComponentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ComponentCreateResponse.Size(m)
}
func (m *ComponentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentCreateResponse proto.InternalMessageInfo

func (m *ComponentCreateResponse) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

func (m *ComponentCreateResponse) GetComponentId() int64 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

type GetComponentRequest struct {
	OrgName              string   `protobuf:"bytes,1,opt,name=orgName,proto3" json:"orgName,omitempty"`
	AppName              string   `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetComponentRequest) Reset()         { *m = GetComponentRequest{} }
func (m *GetComponentRequest) String() string { return proto.CompactTextString(m) }
func (*GetComponentRequest) ProtoMessage()    {}
func (*GetComponentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{3}
}

func (m *GetComponentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetComponentRequest.Unmarshal(m, b)
}
func (m *GetComponentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetComponentRequest.Marshal(b, m, deterministic)
}
func (m *GetComponentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetComponentRequest.Merge(m, src)
}
func (m *GetComponentRequest) XXX_Size() int {
	return xxx_messageInfo_GetComponentRequest.Size(m)
}
func (m *GetComponentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetComponentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetComponentRequest proto.InternalMessageInfo

func (m *GetComponentRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *GetComponentRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

type GetComponentResponse struct {
	Components           []*ComponentDetail `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetComponentResponse) Reset()         { *m = GetComponentResponse{} }
func (m *GetComponentResponse) String() string { return proto.CompactTextString(m) }
func (*GetComponentResponse) ProtoMessage()    {}
func (*GetComponentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f308f339edce119a, []int{4}
}

func (m *GetComponentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetComponentResponse.Unmarshal(m, b)
}
func (m *GetComponentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetComponentResponse.Marshal(b, m, deterministic)
}
func (m *GetComponentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetComponentResponse.Merge(m, src)
}
func (m *GetComponentResponse) XXX_Size() int {
	return xxx_messageInfo_GetComponentResponse.Size(m)
}
func (m *GetComponentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetComponentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetComponentResponse proto.InternalMessageInfo

func (m *GetComponentResponse) GetComponents() []*ComponentDetail {
	if m != nil {
		return m.Components
	}
	return nil
}

func init() {
	proto.RegisterEnum("services.Package_RepoProvider", Package_RepoProvider_name, Package_RepoProvider_value)
	proto.RegisterEnum("services.ComponentDetail_ComponentType", ComponentDetail_ComponentType_name, ComponentDetail_ComponentType_value)
	proto.RegisterEnum("services.ComponentDetail_PipelineState", ComponentDetail_PipelineState_name, ComponentDetail_PipelineState_value)
	proto.RegisterType((*Package)(nil), "services.Package")
	proto.RegisterType((*ComponentDetail)(nil), "services.ComponentDetail")
	proto.RegisterType((*ComponentCreateResponse)(nil), "services.ComponentCreateResponse")
	proto.RegisterType((*GetComponentRequest)(nil), "services.GetComponentRequest")
	proto.RegisterType((*GetComponentResponse)(nil), "services.GetComponentResponse")
}

func init() { proto.RegisterFile("ComponentMgmt.proto", fileDescriptor_f308f339edce119a) }

var fileDescriptor_f308f339edce119a = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x81, 0x40, 0x32, 0x81, 0x64, 0x33, 0x49, 0x55, 0x37, 0x52, 0x23, 0x6a, 0x55, 0x2a,
	0x55, 0x25, 0x0e, 0xf4, 0xd4, 0x43, 0x0f, 0xb1, 0xbd, 0x49, 0x5c, 0x88, 0xed, 0x2e, 0xb6, 0x50,
	0x4f, 0xc8, 0x21, 0x2b, 0x64, 0x95, 0xd8, 0x5b, 0xdb, 0x89, 0xd4, 0x7b, 0x8f, 0xfd, 0x27, 0xfd,
	0x57, 0xfd, 0x25, 0xd5, 0x1a, 0x30, 0x76, 0x3e, 0xd4, 0x1b, 0xf3, 0xe6, 0xe3, 0xbd, 0x79, 0xb3,
	0x18, 0x8e, 0x8c, 0xf8, 0x56, 0xc4, 0x11, 0x8f, 0xb2, 0xab, 0xf9, 0x6d, 0xd6, 0x17, 0x49, 0x9c,
	0xc5, 0xb8, 0x93, 0xf2, 0xe4, 0x3e, 0x9c, 0xf1, 0x54, 0xfb, 0xab, 0x40, 0xcb, 0x0d, 0x66, 0xdf,
	0x83, 0x39, 0x47, 0x84, 0x46, 0x14, 0xdc, 0x72, 0x55, 0xe9, 0x2a, 0xbd, 0x5d, 0x96, 0xff, 0x46,
	0x15, 0x5a, 0x09, 0x17, 0xb1, 0x9f, 0x2c, 0xd4, 0xad, 0x1c, 0x5e, 0x87, 0x78, 0x02, 0x3b, 0x8b,
	0x20, 0x9a, 0xdf, 0x05, 0x73, 0xae, 0xd6, 0xf3, 0x54, 0x11, 0xa3, 0x0e, 0x6d, 0x59, 0xe6, 0x26,
	0xf1, 0x7d, 0x78, 0xc3, 0x13, 0xb5, 0xd1, 0x55, 0x7a, 0xfb, 0x83, 0xd3, 0xfe, 0x9a, 0xb6, 0xbf,
	0xa2, 0xec, 0xb3, 0x52, 0x15, 0xab, 0xf4, 0xe0, 0x29, 0x80, 0x8c, 0x27, 0xfc, 0x5a, 0x92, 0x6f,
	0xe7, 0x0c, 0x25, 0x44, 0x7b, 0x0f, 0xed, 0x72, 0x37, 0x76, 0x60, 0x57, 0xb7, 0x3c, 0xdd, 0x37,
	0x86, 0xd4, 0x23, 0x35, 0x04, 0x68, 0x5e, 0x58, 0xde, 0xa5, 0xaf, 0x13, 0x45, 0xfb, 0xdd, 0x80,
	0x83, 0xc2, 0x06, 0x93, 0x67, 0x41, 0xb8, 0xc0, 0xb7, 0xd0, 0x99, 0xad, 0x21, 0x5b, 0x6e, 0xbd,
	0xdc, 0xa1, 0x0a, 0x62, 0x0f, 0x0e, 0x02, 0x21, 0x16, 0xe1, 0x2c, 0xc8, 0xc2, 0x38, 0xb2, 0x37,
	0xee, 0x3c, 0x84, 0xf1, 0xaa, 0x34, 0xcf, 0xfb, 0x29, 0x78, 0x6e, 0xd7, 0xfe, 0xe0, 0xdd, 0x66,
	0xe7, 0x07, 0x0a, 0x36, 0xb1, 0x2c, 0x67, 0xd5, 0x6e, 0xe9, 0x7b, 0x9c, 0xcc, 0x73, 0xc2, 0xc6,
	0xd2, 0xf7, 0x55, 0x88, 0x1f, 0xa0, 0x25, 0x96, 0xee, 0xe5, 0xa6, 0xec, 0x0d, 0x0e, 0x1f, 0xd9,
	0xca, 0xd6, 0x15, 0x52, 0x95, 0x08, 0x05, 0x5f, 0x84, 0x11, 0x1f, 0x67, 0x41, 0xc6, 0xd5, 0xe6,
	0xff, 0x54, 0xb9, 0xe5, 0x72, 0x56, 0xed, 0xd6, 0x7e, 0x29, 0xd0, 0xa9, 0xc8, 0x46, 0x02, 0xed,
	0x73, 0xe6, 0xd8, 0x1e, 0xb5, 0xcd, 0xe9, 0x84, 0xea, 0xa4, 0x86, 0xc7, 0x40, 0xf4, 0x33, 0x63,
	0x28, 0x81, 0x33, 0xd7, 0x9a, 0x32, 0x3a, 0xf6, 0x88, 0x82, 0x47, 0x70, 0x50, 0x41, 0x5d, 0x83,
	0x6c, 0xe1, 0x21, 0x74, 0xc6, 0xc6, 0x25, 0x35, 0xfd, 0x11, 0x35, 0xa7, 0x5f, 0x1c, 0x9d, 0xd4,
	0xb1, 0x0d, 0x3b, 0x13, 0x87, 0x0d, 0xcf, 0x47, 0xce, 0x84, 0x34, 0xf0, 0x05, 0x1c, 0x5e, 0x9d,
	0xb9, 0x8c, 0x9a, 0xbe, 0x41, 0xa7, 0xc6, 0xc8, 0x1f, 0x7b, 0x94, 0x91, 0x6d, 0xed, 0x33, 0x74,
	0x2a, 0x32, 0x65, 0x97, 0x49, 0xdd, 0x91, 0xf3, 0x8d, 0x9a, 0xa4, 0x86, 0x08, 0xfb, 0xb6, 0xe3,
	0x4d, 0x0d, 0xc7, 0x3e, 0xb7, 0x2e, 0x7c, 0x46, 0x4d, 0xa2, 0xc8, 0xe7, 0xa0, 0x33, 0x67, 0x48,
	0x6d, 0xb2, 0xa5, 0x05, 0xf0, 0xb2, 0x58, 0xc2, 0x48, 0xb8, 0xdc, 0x93, 0xa7, 0x22, 0x8e, 0x52,
	0xfe, 0xf8, 0x55, 0x28, 0x4f, 0xbd, 0x8a, 0x2e, 0xec, 0x15, 0x80, 0x75, 0x93, 0x5f, 0xba, 0xce,
	0xca, 0x90, 0x66, 0xc1, 0xd1, 0x05, 0xcf, 0x0a, 0x16, 0xc6, 0x7f, 0xdc, 0xf1, 0x34, 0x2b, 0x5f,
	0x55, 0xa9, 0x5e, 0x55, 0x85, 0x56, 0x20, 0x44, 0x9e, 0x59, 0xfd, 0xcf, 0x56, 0xa1, 0xf6, 0x15,
	0x8e, 0xab, 0xa3, 0x56, 0x52, 0x3f, 0x01, 0x14, 0x8c, 0xa9, 0xaa, 0x74, 0xeb, 0xbd, 0xbd, 0xc1,
	0xab, 0x67, 0xef, 0xca, 0x4a, 0xc5, 0x83, 0x3f, 0xe5, 0x33, 0xca, 0xcf, 0x02, 0x5e, 0x42, 0x73,
	0xe9, 0x04, 0x3e, 0x3f, 0xe2, 0xe4, 0xcd, 0x13, 0xa9, 0xaa, 0x7f, 0x5a, 0x0d, 0x5d, 0xe8, 0x94,
	0xe5, 0xa6, 0xf8, 0x7a, 0xd3, 0xf5, 0x84, 0x25, 0x27, 0xa7, 0xcf, 0xa5, 0xd7, 0x13, 0xaf, 0x9b,
	0xf9, 0x37, 0xeb, 0xe3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf6, 0x68, 0x58, 0xca, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ComponentMgmtClient is the client API for ComponentMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComponentMgmtClient interface {
	Create(ctx context.Context, in *ComponentDetail, opts ...grpc.CallOption) (*ComponentCreateResponse, error)
	GetComponents(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*GetComponentResponse, error)
}

type componentMgmtClient struct {
	cc *grpc.ClientConn
}

func NewComponentMgmtClient(cc *grpc.ClientConn) ComponentMgmtClient {
	return &componentMgmtClient{cc}
}

func (c *componentMgmtClient) Create(ctx context.Context, in *ComponentDetail, opts ...grpc.CallOption) (*ComponentCreateResponse, error) {
	out := new(ComponentCreateResponse)
	err := c.cc.Invoke(ctx, "/services.ComponentMgmt/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentMgmtClient) GetComponents(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*GetComponentResponse, error) {
	out := new(GetComponentResponse)
	err := c.cc.Invoke(ctx, "/services.ComponentMgmt/GetComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentMgmtServer is the server API for ComponentMgmt service.
type ComponentMgmtServer interface {
	Create(context.Context, *ComponentDetail) (*ComponentCreateResponse, error)
	GetComponents(context.Context, *GetComponentRequest) (*GetComponentResponse, error)
}

func RegisterComponentMgmtServer(s *grpc.Server, srv ComponentMgmtServer) {
	s.RegisterService(&_ComponentMgmt_serviceDesc, srv)
}

func _ComponentMgmt_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentMgmtServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ComponentMgmt/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentMgmtServer).Create(ctx, req.(*ComponentDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentMgmt_GetComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentMgmtServer).GetComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ComponentMgmt/GetComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentMgmtServer).GetComponents(ctx, req.(*GetComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComponentMgmt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ComponentMgmt",
	HandlerType: (*ComponentMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ComponentMgmt_Create_Handler,
		},
		{
			MethodName: "GetComponents",
			Handler:    _ComponentMgmt_GetComponents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ComponentMgmt.proto",
}
