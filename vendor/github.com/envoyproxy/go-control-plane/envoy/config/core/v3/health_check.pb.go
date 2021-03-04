// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/health_check.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0}
}

type HealthCheck struct {
	Timeout               *duration.Duration    `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval              *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter         *duration.Duration    `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter        *duration.Duration    `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent uint32                `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold    *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold      *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort               *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker                isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	NoTrafficInterval            *duration.Duration          `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration          `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration          `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration          `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                      `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	AlwaysLogHealthCheckFailures bool                        `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	TlsOptions                   *HealthCheck_TlsOptions     `protobuf:"bytes,21,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

func (m *HealthCheck) GetTlsOptions() *HealthCheck_TlsOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                             string               `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                             string               `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                             *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                          *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	HiddenEnvoyDeprecatedServiceName string               `protobuf:"bytes,5,opt,name=hidden_envoy_deprecated_service_name,json=hiddenEnvoyDeprecatedServiceName,proto3" json:"hidden_envoy_deprecated_service_name,omitempty"` // Deprecated: Do not use.
	RequestHeadersToAdd              []*HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove           []string             `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	HiddenEnvoyDeprecatedUseHttp2    bool                 `protobuf:"varint,7,opt,name=hidden_envoy_deprecated_use_http2,json=hiddenEnvoyDeprecatedUseHttp2,proto3" json:"hidden_envoy_deprecated_use_http2,omitempty"` // Deprecated: Do not use.
	ExpectedStatuses                 []*v3.Int64Range     `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType                  v3.CodecClientType   `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.v3.CodecClientType" json:"codec_client_type,omitempty"`
	ServiceNameMatcher               *v31.StringMatcher   `protobuf:"bytes,11,opt,name=service_name_matcher,json=serviceNameMatcher,proto3" json:"service_name_matcher,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedServiceName() string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedServiceName
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedUseHttp2() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseHttp2
	}
	return false
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*v3.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() v3.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return v3.CodecClientType_HTTP1
}

func (m *HealthCheck_HttpHealthCheck) GetServiceNameMatcher() *v31.StringMatcher {
	if m != nil {
		return m.ServiceNameMatcher
	}
	return nil
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {
}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_CustomHealthCheck) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig)(nil),
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

type HealthCheck_TlsOptions struct {
	AlpnProtocols        []string `protobuf:"bytes,1,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_TlsOptions) Reset()         { *m = HealthCheck_TlsOptions{} }
func (m *HealthCheck_TlsOptions) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TlsOptions) ProtoMessage()    {}
func (*HealthCheck_TlsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba40dc0efbf9537, []int{0, 6}
}

func (m *HealthCheck_TlsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TlsOptions.Unmarshal(m, b)
}
func (m *HealthCheck_TlsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TlsOptions.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TlsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TlsOptions.Merge(m, src)
}
func (m *HealthCheck_TlsOptions) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TlsOptions.Size(m)
}
func (m *HealthCheck_TlsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TlsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TlsOptions proto.InternalMessageInfo

func (m *HealthCheck_TlsOptions) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.core.v3.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.config.core.v3.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.config.core.v3.HealthCheck.CustomHealthCheck")
	proto.RegisterType((*HealthCheck_TlsOptions)(nil), "envoy.config.core.v3.HealthCheck.TlsOptions")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/health_check.proto", fileDescriptor_2ba40dc0efbf9537)
}

var fileDescriptor_2ba40dc0efbf9537 = []byte{
	// 1527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0x4f, 0x53, 0xe3, 0xc8,
	0x15, 0xb7, 0x8c, 0xc1, 0xf6, 0x33, 0x18, 0xb9, 0x81, 0x45, 0x78, 0x19, 0xd6, 0x33, 0x61, 0x6b,
	0x1c, 0x6a, 0x62, 0xef, 0xc2, 0x66, 0x92, 0xa5, 0x2a, 0x95, 0x20, 0x60, 0xc6, 0x6c, 0x66, 0x18,
	0xaa, 0x31, 0x3b, 0x95, 0xda, 0x54, 0x29, 0x8d, 0xd4, 0xd8, 0xca, 0x0a, 0xb5, 0xd2, 0x6a, 0x7b,
	0xc7, 0xb7, 0x54, 0x4e, 0x53, 0x39, 0xe6, 0x90, 0x43, 0x3e, 0xc2, 0x7c, 0x82, 0x9c, 0x72, 0xcb,
	0x21, 0xd7, 0x1c, 0xf3, 0x09, 0x72, 0xca, 0x21, 0xa7, 0x14, 0xa7, 0x54, 0xb7, 0x64, 0x5b, 0xfe,
	0x03, 0x86, 0xe4, 0x66, 0xbd, 0xf7, 0xfb, 0xfd, 0xde, 0xeb, 0x3f, 0xef, 0xbd, 0x36, 0x3c, 0xa5,
	0x7e, 0x97, 0xf5, 0xea, 0x36, 0xf3, 0xaf, 0xdc, 0x56, 0xdd, 0x66, 0x9c, 0xd6, 0xbb, 0x7b, 0xf5,
	0x36, 0x25, 0x9e, 0x68, 0x5b, 0x76, 0x9b, 0xda, 0xdf, 0xd6, 0x02, 0xce, 0x04, 0x43, 0xab, 0x0a,
	0x58, 0x8b, 0x80, 0x35, 0x09, 0xac, 0x75, 0xf7, 0xca, 0x9f, 0x4c, 0xa5, 0x5f, 0x92, 0x90, 0x46,
	0xb4, 0xf2, 0x93, 0x08, 0x20, 0x7a, 0x01, 0xad, 0x5f, 0x13, 0x61, 0xb7, 0x29, 0x97, 0x88, 0x50,
	0x70, 0xd7, 0x6f, 0xc5, 0x18, 0x23, 0x81, 0x91, 0xc1, 0x85, 0x08, 0x62, 0xcf, 0xc6, 0xa8, 0x87,
	0x13, 0xbf, 0xd5, 0x17, 0xde, 0x68, 0x31, 0xd6, 0xf2, 0x68, 0x5d, 0x7d, 0x5d, 0x76, 0xae, 0xea,
	0xc4, 0xef, 0xc5, 0xae, 0xad, 0x71, 0x97, 0xd3, 0xe1, 0x44, 0xb8, 0xcc, 0x8f, 0xfd, 0x9b, 0xe3,
	0xfe, 0x50, 0xf0, 0x8e, 0x2d, 0x6e, 0x63, 0x7f, 0xc7, 0x49, 0x10, 0x50, 0x1e, 0xc6, 0xfe, 0xc7,
	0x1d, 0x27, 0x20, 0x75, 0xe2, 0xfb, 0x4c, 0x28, 0xd1, 0xb0, 0xde, 0xa5, 0x3c, 0x74, 0x99, 0x3f,
	0x5c, 0xd0, 0xf7, 0xa2, 0xb4, 0x93, 0x18, 0x87, 0x06, 0x9c, 0xda, 0xc9, 0x2c, 0xd6, 0xbb, 0xc4,
	0x73, 0x1d, 0x22, 0x68, 0xbd, 0xff, 0x23, 0x72, 0x3c, 0xf9, 0x47, 0x19, 0x0a, 0x0d, 0x75, 0x00,
	0x87, 0x72, 0xff, 0xd1, 0x4f, 0x21, 0x2b, 0xdc, 0x6b, 0xca, 0x3a, 0xc2, 0xd0, 0x2a, 0x5a, 0xb5,
	0xb0, 0xbb, 0x51, 0x8b, 0x52, 0xac, 0xf5, 0x53, 0xac, 0x1d, 0xc5, 0x0b, 0x34, 0xe1, 0xc6, 0xcc,
	0x7e, 0xd0, 0x32, 0x39, 0x6d, 0x27, 0x85, 0xfb, 0x2c, 0x74, 0x00, 0x39, 0xd7, 0x17, 0x94, 0x77,
	0x89, 0x67, 0xa4, 0x1f, 0xa2, 0x30, 0xa0, 0xa1, 0x9f, 0x41, 0xd1, 0xf5, 0x5d, 0xe1, 0x12, 0xcf,
	0xfa, 0xb5, 0x2b, 0x04, 0xe5, 0xc6, 0xea, 0x0c, 0x21, 0xbc, 0x14, 0x13, 0xbe, 0x52, 0x78, 0x64,
	0xc2, 0x72, 0x5f, 0xad, 0x2f, 0x31, 0x37, 0x4b, 0xa2, 0xd8, 0x67, 0xc4, 0x1a, 0xcf, 0x61, 0x7d,
	0x4c, 0xc3, 0x0a, 0x28, 0xb7, 0xa9, 0x2f, 0x0c, 0x54, 0xd1, 0xaa, 0x4b, 0x78, 0x6d, 0x94, 0x70,
	0x16, 0x39, 0xd1, 0x5b, 0x58, 0xe9, 0xf8, 0xd1, 0x9d, 0xee, 0x59, 0xa2, 0xcd, 0x69, 0xd8, 0x66,
	0x9e, 0x63, 0x64, 0x54, 0xfc, 0xcd, 0x89, 0xf8, 0x17, 0x27, 0xbe, 0xd8, 0xdb, 0xfd, 0x9a, 0x78,
	0x1d, 0x6a, 0xe6, 0x6e, 0xcc, 0xf9, 0xdf, 0x6b, 0x69, 0x5d, 0xc3, 0x68, 0x20, 0xd1, 0xec, 0x2b,
	0xa0, 0x73, 0x28, 0x4d, 0xca, 0xce, 0x3f, 0x48, 0x56, 0x9f, 0x10, 0xfd, 0x11, 0xe4, 0x88, 0x27,
	0xac, 0x80, 0x71, 0x61, 0x2c, 0xcc, 0xd6, 0xc2, 0x59, 0xe2, 0x89, 0x33, 0xc6, 0x05, 0x3a, 0x06,
	0x9d, 0xd3, 0x4e, 0x48, 0x2d, 0x9b, 0xf9, 0x3e, 0xb5, 0xe5, 0x16, 0x1a, 0x59, 0x25, 0x50, 0x9e,
	0x10, 0x30, 0x19, 0xf3, 0x22, 0xfa, 0xb2, 0xe2, 0x1c, 0x0e, 0x28, 0xc8, 0x82, 0x92, 0x2c, 0x41,
	0x2b, 0xd9, 0x04, 0x8c, 0x9c, 0xd2, 0xf9, 0xbc, 0x36, 0xad, 0x0b, 0xd4, 0x12, 0xb7, 0xb5, 0xd6,
	0x10, 0x22, 0x48, 0x7c, 0x37, 0x52, 0x78, 0xb9, 0x3d, 0x6a, 0x42, 0xbf, 0x04, 0x5d, 0xd8, 0x63,
	0xfa, 0x79, 0xa5, 0xff, 0xd9, 0x6c, 0xfd, 0xa6, 0x3d, 0x26, 0x5f, 0x14, 0x23, 0x16, 0x99, 0x7e,
	0x8b, 0x07, 0xf6, 0xa8, 0x7c, 0xe1, 0xbe, 0xe9, 0xbf, 0xe4, 0x81, 0x3d, 0x96, 0x7e, 0x6b, 0xd4,
	0x84, 0x28, 0xac, 0xd8, 0x9d, 0x50, 0xb0, 0xeb, 0xd1, 0x10, 0x4b, 0x2a, 0xc4, 0xde, 0xec, 0x10,
	0x87, 0x8a, 0x3c, 0x1a, 0xa4, 0x64, 0x8f, 0x1b, 0xd1, 0x39, 0xac, 0xf8, 0xcc, 0x12, 0x9c, 0x5c,
	0x5d, 0xb9, 0xb6, 0x35, 0x28, 0xe0, 0xc5, 0x59, 0x05, 0x2c, 0xaf, 0xd6, 0x07, 0x2d, 0xbd, 0x93,
	0xc2, 0x25, 0x9f, 0x35, 0x23, 0xfa, 0x49, 0xbf, 0x8e, 0x31, 0x0c, 0xaf, 0xf1, 0x50, 0xb3, 0xf8,
	0x00, 0xcd, 0x01, 0x7d, 0xa0, 0xf9, 0x0d, 0xac, 0x0f, 0x35, 0xa9, 0xd3, 0xa2, 0x43, 0xe1, 0xe5,
	0xfb, 0x0b, 0xaf, 0x0d, 0x34, 0x8e, 0x9d, 0x16, 0x1d, 0x88, 0xbf, 0x85, 0xb5, 0xe9, 0xd2, 0xfa,
	0xfd, 0xa5, 0x57, 0xa6, 0x09, 0x6f, 0x43, 0x91, 0x76, 0xa9, 0x2f, 0x2c, 0x8f, 0xb5, 0xac, 0x80,
	0x88, 0xb6, 0x51, 0xaa, 0x68, 0xd5, 0x3c, 0x5e, 0x54, 0xd6, 0x57, 0xac, 0x75, 0x46, 0x44, 0x1b,
	0xbd, 0x80, 0x0a, 0xf1, 0xbe, 0x23, 0xbd, 0x50, 0xc1, 0x92, 0xe7, 0x6d, 0x5d, 0x11, 0xd7, 0xeb,
	0x70, 0x1a, 0x1a, 0x2b, 0x15, 0xad, 0x9a, 0xc3, 0x9b, 0x11, 0xee, 0x15, 0x6b, 0x25, 0x0e, 0xf1,
	0x45, 0x8c, 0x41, 0xaf, 0xa1, 0x20, 0xbc, 0xd0, 0x62, 0x81, 0x9a, 0x06, 0xc6, 0x9a, 0x4a, 0xfe,
	0xd9, 0x3d, 0x6e, 0xbb, 0x17, 0xbe, 0x89, 0x38, 0x18, 0xc4, 0xe0, 0x77, 0xf9, 0xb7, 0x1a, 0x64,
	0xcf, 0x48, 0xcf, 0x63, 0xc4, 0x41, 0x8f, 0x20, 0x23, 0xe8, 0xbb, 0x68, 0x36, 0xe4, 0xcd, 0xec,
	0x8d, 0x99, 0xe1, 0xe9, 0x8a, 0xd6, 0x48, 0x61, 0x65, 0x46, 0x06, 0x2c, 0x5c, 0xba, 0x3e, 0xe1,
	0x3d, 0xd5, 0xfa, 0x17, 0x1b, 0x29, 0x1c, 0x7f, 0xef, 0x3f, 0xfb, 0xd3, 0x5f, 0xdf, 0x6f, 0x3d,
	0x85, 0x4f, 0xa3, 0x24, 0x48, 0xe0, 0xd6, 0xba, 0xbb, 0x51, 0x12, 0xc9, 0x0c, 0xe2, 0x30, 0x66,
	0x11, 0xb2, 0x41, 0x1c, 0x71, 0xee, 0x3f, 0xa6, 0x56, 0xfe, 0xcb, 0x02, 0x2c, 0x8f, 0xd5, 0x3a,
	0x42, 0x90, 0x69, 0xb3, 0x30, 0x4e, 0x05, 0xab, 0xdf, 0xe8, 0x63, 0xc8, 0xa8, 0xdd, 0x4d, 0x8f,
	0xa4, 0x87, 0x95, 0x11, 0xfd, 0x04, 0x32, 0x21, 0xf5, 0x9d, 0x78, 0x12, 0x7c, 0x7f, 0xf6, 0x7e,
	0xc4, 0xd9, 0x60, 0x45, 0x43, 0x87, 0x90, 0xe5, 0xd4, 0xa6, 0x6e, 0x97, 0xc6, 0xbd, 0xfc, 0x01,
	0x0a, 0x7d, 0x26, 0xc2, 0xb0, 0xdd, 0x76, 0x1d, 0x87, 0xfa, 0x96, 0xe2, 0x5a, 0xfd, 0x49, 0x4d,
	0x1d, 0x2b, 0xa4, 0xbc, 0xeb, 0xda, 0xd4, 0xf2, 0xc9, 0x35, 0x55, 0x6d, 0x3d, 0x6f, 0xa6, 0x0d,
	0x0d, 0x57, 0x22, 0xfc, 0xb1, 0x84, 0x1f, 0x0d, 0xd0, 0xe7, 0x11, 0xf8, 0x94, 0x5c, 0x53, 0xd4,
	0x82, 0x8f, 0x38, 0xfd, 0x4d, 0x87, 0x86, 0x42, 0xde, 0x19, 0x87, 0xf2, 0xd0, 0x12, 0xcc, 0x22,
	0x8e, 0x63, 0x2c, 0x54, 0xe6, 0xaa, 0x85, 0xdd, 0xa7, 0xb7, 0xe6, 0xe9, 0x50, 0xae, 0xda, 0x72,
	0x74, 0xd2, 0x66, 0xfe, 0xc6, 0x5c, 0xf8, 0x83, 0x36, 0xa7, 0xff, 0x33, 0x8b, 0x57, 0x62, 0xc5,
	0x08, 0x14, 0x36, 0xd9, 0x81, 0xe3, 0xa0, 0x2f, 0x61, 0x63, 0x4a, 0x20, 0x4e, 0xaf, 0x59, 0x97,
	0x1a, 0xb9, 0xca, 0x5c, 0x35, 0x8f, 0x3f, 0x1a, 0xe7, 0x61, 0xe5, 0x45, 0x18, 0x1e, 0xdf, 0xb6,
	0x6e, 0x39, 0x43, 0x64, 0xd3, 0xde, 0x55, 0xe3, 0x23, 0x67, 0xe6, 0xfe, 0xfc, 0xaf, 0x7f, 0xff,
	0x7d, 0x5e, 0x33, 0x34, 0xfc, 0x68, 0xea, 0xd2, 0x2f, 0x42, 0x2a, 0xaf, 0xc2, 0x2e, 0x7a, 0x01,
	0x25, 0xfa, 0x2e, 0xa0, 0xb6, 0xda, 0x3c, 0x41, 0x44, 0x27, 0xa4, 0xa1, 0x91, 0x57, 0x4b, 0xde,
	0x88, 0x97, 0x2c, 0xdf, 0x72, 0x72, 0xad, 0x27, 0xbe, 0x78, 0xfe, 0x05, 0x96, 0x0f, 0x3a, 0xac,
	0xf7, 0x39, 0xe7, 0x31, 0x05, 0x35, 0xa1, 0x64, 0x33, 0x87, 0xda, 0x96, 0xed, 0xb9, 0xb2, 0x46,
	0x25, 0xc9, 0x80, 0x8a, 0x56, 0x2d, 0xee, 0x6e, 0x8d, 0xe9, 0x1c, 0x4a, 0xdc, 0xa1, 0x82, 0x35,
	0x7b, 0x41, 0x34, 0x59, 0x7f, 0xa7, 0x26, 0xeb, 0xb2, 0x3d, 0xea, 0x42, 0x5f, 0xc3, 0x6a, 0xf2,
	0x44, 0xad, 0xf8, 0x3d, 0x1a, 0x0f, 0x87, 0xed, 0xa4, 0x70, 0xec, 0x92, 0x01, 0xce, 0xd5, 0x53,
	0xf5, 0x75, 0x64, 0xc0, 0x28, 0x1c, 0x1e, 0x73, 0x6c, 0xdb, 0xff, 0x42, 0x16, 0x52, 0x1d, 0x7e,
	0x70, 0x77, 0x21, 0x8d, 0x15, 0x4b, 0xf9, 0x6f, 0x1a, 0x14, 0x47, 0x87, 0xd9, 0xa0, 0x1c, 0xb4,
	0xff, 0xbb, 0x1c, 0xd2, 0x6a, 0xcf, 0xff, 0x87, 0x72, 0xd8, 0xdf, 0x93, 0x8b, 0xa9, 0xc1, 0xb3,
	0xbb, 0x17, 0x33, 0x9a, 0x78, 0xf9, 0x1b, 0xd0, 0x31, 0x75, 0xdc, 0x30, 0xb9, 0x18, 0x1d, 0xe6,
	0xbe, 0xa5, 0xbd, 0xb8, 0x17, 0xc8, 0x9f, 0xfb, 0x3f, 0x94, 0xd2, 0x9f, 0x41, 0xed, 0x6e, 0xe9,
	0x71, 0xa1, 0xf2, 0x7b, 0x0d, 0x96, 0xc7, 0xc6, 0x32, 0x7a, 0x0c, 0x8b, 0x23, 0xc5, 0x19, 0x45,
	0x29, 0x24, 0x0e, 0x07, 0x6d, 0x42, 0x9e, 0x74, 0x44, 0x9b, 0x71, 0x57, 0x44, 0xbd, 0x2f, 0x8f,
	0x87, 0x86, 0x7b, 0x9e, 0xd9, 0x58, 0xd8, 0xf2, 0x1f, 0xd3, 0x50, 0x9a, 0x18, 0xdf, 0xb2, 0xc5,
	0x0d, 0x93, 0x48, 0xb4, 0x38, 0x69, 0x44, 0x97, 0xb0, 0x75, 0x5b, 0x99, 0x45, 0xc7, 0x12, 0x3f,
	0xc9, 0xd7, 0x27, 0x26, 0xd9, 0xb9, 0xfa, 0x57, 0x22, 0x3b, 0x4e, 0x23, 0x85, 0x3f, 0x9e, 0x5a,
	0x78, 0x87, 0x4a, 0x01, 0x7d, 0x09, 0x8b, 0xf2, 0xd6, 0x0e, 0x14, 0xa3, 0x76, 0xba, 0x3a, 0xa1,
	0x78, 0xe0, 0xf7, 0x1a, 0x29, 0x5c, 0x50, 0xd8, 0x88, 0xba, 0xff, 0x5c, 0xee, 0xc3, 0xe7, 0x50,
	0xbf, 0x7b, 0x1f, 0x26, 0xd6, 0x6c, 0x2e, 0x41, 0x21, 0x0a, 0xa6, 0x6a, 0xb3, 0xec, 0x00, 0x0c,
	0x47, 0x15, 0xfa, 0x14, 0x8a, 0xc4, 0x0b, 0x7c, 0x4b, 0x05, 0xb6, 0x99, 0x17, 0x1a, 0x9a, 0x6a,
	0x45, 0x4b, 0xd2, 0x7a, 0xd6, 0x37, 0xee, 0xd7, 0x65, 0xec, 0x1d, 0xa8, 0xce, 0xb8, 0x6a, 0x03,
	0xdd, 0xfd, 0x6d, 0x49, 0xf8, 0x04, 0x1e, 0xdd, 0x49, 0x30, 0xd7, 0xa0, 0x98, 0x1c, 0xd4, 0x94,
	0xab, 0x81, 0xf5, 0x55, 0x26, 0x07, 0x7a, 0x61, 0xe7, 0x57, 0xb0, 0x18, 0x61, 0xa3, 0x5e, 0x83,
	0x0a, 0x90, 0xbd, 0x38, 0xfd, 0xf9, 0xe9, 0x9b, 0xb7, 0xa7, 0x7a, 0x4a, 0x7e, 0x34, 0x8e, 0x0f,
	0x5e, 0x35, 0x1b, 0xbf, 0xd0, 0x35, 0xb4, 0x04, 0xf9, 0x8b, 0xd3, 0xfe, 0x67, 0x1a, 0x2d, 0x42,
	0xee, 0x08, 0x1f, 0x9c, 0x9c, 0x9e, 0x9c, 0xbe, 0xd4, 0xe7, 0x24, 0xb2, 0x79, 0xf2, 0xfa, 0xf8,
	0xcd, 0x45, 0x53, 0xcf, 0x28, 0xd7, 0xf1, 0x4b, 0x7c, 0x70, 0x74, 0x7c, 0xa4, 0xcf, 0x9b, 0x3f,
	0x86, 0x27, 0x2e, 0x8b, 0xee, 0x78, 0xc0, 0xd9, 0xbb, 0xde, 0xd4, 0x1a, 0x34, 0xf5, 0x44, 0xc6,
	0x6a, 0x47, 0xce, 0xb4, 0xcb, 0x05, 0xb5, 0x5f, 0x7b, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xf0, 0xf4, 0x1e, 0x94, 0x0f, 0x00, 0x00,
}