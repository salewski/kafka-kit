// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/registry.proto

package registry

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type BrokerRequest struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerRequest) Reset()         { *m = BrokerRequest{} }
func (m *BrokerRequest) String() string { return proto.CompactTextString(m) }
func (*BrokerRequest) ProtoMessage()    {}
func (*BrokerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{0}
}

func (m *BrokerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerRequest.Unmarshal(m, b)
}
func (m *BrokerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerRequest.Marshal(b, m, deterministic)
}
func (m *BrokerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerRequest.Merge(m, src)
}
func (m *BrokerRequest) XXX_Size() int {
	return xxx_messageInfo_BrokerRequest.Size(m)
}
func (m *BrokerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerRequest proto.InternalMessageInfo

func (m *BrokerRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *BrokerRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BrokerResponse struct {
	Brokers              map[uint32]*Broker `protobuf:"bytes,5,rep,name=brokers,proto3" json:"brokers,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ids                  []uint32           `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BrokerResponse) Reset()         { *m = BrokerResponse{} }
func (m *BrokerResponse) String() string { return proto.CompactTextString(m) }
func (*BrokerResponse) ProtoMessage()    {}
func (*BrokerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{1}
}

func (m *BrokerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerResponse.Unmarshal(m, b)
}
func (m *BrokerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerResponse.Marshal(b, m, deterministic)
}
func (m *BrokerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerResponse.Merge(m, src)
}
func (m *BrokerResponse) XXX_Size() int {
	return xxx_messageInfo_BrokerResponse.Size(m)
}
func (m *BrokerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerResponse proto.InternalMessageInfo

func (m *BrokerResponse) GetBrokers() map[uint32]*Broker {
	if m != nil {
		return m.Brokers
	}
	return nil
}

func (m *BrokerResponse) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Broker struct {
	// Registry metadata.
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Broker metadata from ZooKeeper.
	Id                          uint32            `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	ListenerSecurityProtocolMap map[string]string `protobuf:"bytes,6,rep,name=listener_security_protocol_map,json=listenerSecurityProtocolMap,proto3" json:"listener_security_protocol_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Endpoints                   []string          `protobuf:"bytes,7,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Rack                        string            `protobuf:"bytes,8,opt,name=rack,proto3" json:"rack,omitempty"`
	JmxPort                     uint32            `protobuf:"varint,9,opt,name=jmx_port,json=jmxPort,proto3" json:"jmx_port,omitempty"`
	Host                        string            `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`
	Timestamp                   int64             `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Port                        uint32            `protobuf:"varint,12,opt,name=port,proto3" json:"port,omitempty"`
	Version                     uint32            `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}          `json:"-"`
	XXX_unrecognized            []byte            `json:"-"`
	XXX_sizecache               int32             `json:"-"`
}

func (m *Broker) Reset()         { *m = Broker{} }
func (m *Broker) String() string { return proto.CompactTextString(m) }
func (*Broker) ProtoMessage()    {}
func (*Broker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{2}
}

func (m *Broker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Broker.Unmarshal(m, b)
}
func (m *Broker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Broker.Marshal(b, m, deterministic)
}
func (m *Broker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Broker.Merge(m, src)
}
func (m *Broker) XXX_Size() int {
	return xxx_messageInfo_Broker.Size(m)
}
func (m *Broker) XXX_DiscardUnknown() {
	xxx_messageInfo_Broker.DiscardUnknown(m)
}

var xxx_messageInfo_Broker proto.InternalMessageInfo

func (m *Broker) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Broker) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Broker) GetListenerSecurityProtocolMap() map[string]string {
	if m != nil {
		return m.ListenerSecurityProtocolMap
	}
	return nil
}

func (m *Broker) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Broker) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Broker) GetJmxPort() uint32 {
	if m != nil {
		return m.JmxPort
	}
	return 0
}

func (m *Broker) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Broker) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Broker) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Broker) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type TopicRequest struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicRequest) Reset()         { *m = TopicRequest{} }
func (m *TopicRequest) String() string { return proto.CompactTextString(m) }
func (*TopicRequest) ProtoMessage()    {}
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{3}
}

func (m *TopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicRequest.Unmarshal(m, b)
}
func (m *TopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicRequest.Marshal(b, m, deterministic)
}
func (m *TopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRequest.Merge(m, src)
}
func (m *TopicRequest) XXX_Size() int {
	return xxx_messageInfo_TopicRequest.Size(m)
}
func (m *TopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRequest proto.InternalMessageInfo

func (m *TopicRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TopicRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TopicResponse struct {
	Topics               map[string]*Topic `protobuf:"bytes,5,rep,name=topics,proto3" json:"topics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Names                []string          `protobuf:"bytes,6,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TopicResponse) Reset()         { *m = TopicResponse{} }
func (m *TopicResponse) String() string { return proto.CompactTextString(m) }
func (*TopicResponse) ProtoMessage()    {}
func (*TopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{4}
}

func (m *TopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicResponse.Unmarshal(m, b)
}
func (m *TopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicResponse.Marshal(b, m, deterministic)
}
func (m *TopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicResponse.Merge(m, src)
}
func (m *TopicResponse) XXX_Size() int {
	return xxx_messageInfo_TopicResponse.Size(m)
}
func (m *TopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopicResponse proto.InternalMessageInfo

func (m *TopicResponse) GetTopics() map[string]*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *TopicResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Topic struct {
	// Registry metadata.
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Topic metadata from ZooKeeper.
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Partitions           uint32   `protobuf:"varint,6,opt,name=partitions,proto3" json:"partitions,omitempty"`
	Replication          uint32   `protobuf:"varint,7,opt,name=replication,proto3" json:"replication,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_4215e5fe8e6d7e5d, []int{5}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Topic) GetPartitions() uint32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *Topic) GetReplication() uint32 {
	if m != nil {
		return m.Replication
	}
	return 0
}

func init() {
	proto.RegisterType((*BrokerRequest)(nil), "registry.BrokerRequest")
	proto.RegisterType((*BrokerResponse)(nil), "registry.BrokerResponse")
	proto.RegisterMapType((map[uint32]*Broker)(nil), "registry.BrokerResponse.BrokersEntry")
	proto.RegisterType((*Broker)(nil), "registry.Broker")
	proto.RegisterMapType((map[string]string)(nil), "registry.Broker.ListenerSecurityProtocolMapEntry")
	proto.RegisterMapType((map[string]string)(nil), "registry.Broker.TagsEntry")
	proto.RegisterType((*TopicRequest)(nil), "registry.TopicRequest")
	proto.RegisterType((*TopicResponse)(nil), "registry.TopicResponse")
	proto.RegisterMapType((map[string]*Topic)(nil), "registry.TopicResponse.TopicsEntry")
	proto.RegisterType((*Topic)(nil), "registry.Topic")
	proto.RegisterMapType((map[string]string)(nil), "registry.Topic.TagsEntry")
}

func init() { proto.RegisterFile("protos/registry.proto", fileDescriptor_4215e5fe8e6d7e5d) }

var fileDescriptor_4215e5fe8e6d7e5d = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0xe6, 0xc7, 0xe3, 0xa4, 0x2d, 0xc3, 0x4f, 0xb7, 0xa6, 0xaa, 0x2c, 0xa3, 0xa2,
	0x5c, 0x48, 0xd4, 0x54, 0x02, 0x04, 0x07, 0x24, 0x24, 0x54, 0x09, 0x15, 0x54, 0x99, 0x8a, 0x03,
	0x97, 0xc8, 0x4d, 0x56, 0x61, 0xdb, 0xd8, 0x6b, 0x76, 0xb7, 0x51, 0x73, 0xe5, 0x15, 0x78, 0x0f,
	0x24, 0x1e, 0x82, 0x27, 0xe0, 0x15, 0x10, 0x8f, 0xc0, 0x19, 0xed, 0x7a, 0xdd, 0x38, 0x0d, 0x05,
	0x51, 0x6e, 0x33, 0x93, 0xf9, 0xbe, 0x9d, 0xf1, 0xf7, 0x65, 0xe0, 0x76, 0x26, 0xb8, 0xe2, 0xb2,
	0x27, 0xe8, 0x98, 0x49, 0x25, 0x66, 0x5d, 0x93, 0x63, 0xb3, 0xc8, 0xfd, 0xad, 0x31, 0xe7, 0xe3,
	0x09, 0xed, 0xc5, 0x19, 0xeb, 0xc5, 0x69, 0xca, 0x55, 0xac, 0x18, 0x4f, 0x65, 0xde, 0x17, 0xee,
	0x41, 0xfb, 0xb9, 0xe0, 0xa7, 0x54, 0x44, 0xf4, 0xc3, 0x19, 0x95, 0x0a, 0x11, 0x56, 0x54, 0x3c,
	0x96, 0xc4, 0x09, 0xaa, 0x1d, 0x37, 0x32, 0x31, 0xae, 0x42, 0x85, 0x8d, 0x48, 0x25, 0x70, 0x3a,
	0xed, 0xa8, 0xc2, 0x46, 0xe1, 0x17, 0x07, 0x56, 0x0b, 0x94, 0xcc, 0x78, 0x2a, 0x29, 0x3e, 0x83,
	0xc6, 0xb1, 0xa9, 0x48, 0x52, 0x0b, 0xaa, 0x1d, 0xaf, 0xbf, 0xd3, 0xbd, 0x98, 0x68, 0xb1, 0xd5,
	0xa6, 0xf2, 0x45, 0xaa, 0xc4, 0x2c, 0x2a, 0x50, 0xb8, 0x0e, 0x55, 0x36, 0x92, 0xa4, 0x1e, 0x54,
	0x3b, 0xed, 0x48, 0x87, 0xfe, 0x01, 0xb4, 0xca, 0xad, 0xba, 0xe3, 0x94, 0xce, 0x88, 0x63, 0xc6,
	0xd0, 0x21, 0xde, 0x87, 0xda, 0x34, 0x9e, 0x9c, 0x51, 0x33, 0x9a, 0xd7, 0x5f, 0x5f, 0x7a, 0x32,
	0xff, 0xf9, 0x49, 0xe5, 0xb1, 0x13, 0xfe, 0xac, 0x42, 0x3d, 0xaf, 0x62, 0xb7, 0xb4, 0xa2, 0xd7,
	0xf7, 0x2f, 0xa3, 0xba, 0x47, 0xf1, 0xd8, 0x4e, 0x57, 0x5e, 0xbf, 0x56, 0xac, 0x8f, 0x53, 0xd8,
	0x9e, 0x30, 0xa9, 0x68, 0x4a, 0xc5, 0x40, 0xd2, 0xe1, 0x99, 0x60, 0x6a, 0x36, 0x30, 0x9f, 0x73,
	0xc8, 0x27, 0x83, 0x24, 0xce, 0xcc, 0x16, 0x5e, 0x7f, 0x77, 0x89, 0xf9, 0xc0, 0xc2, 0xde, 0x58,
	0xd4, 0xa1, 0x05, 0xbd, 0x8a, 0xb3, 0xfc, 0xc1, 0xbb, 0x93, 0xab, 0x3b, 0x70, 0x0b, 0x5c, 0x9a,
	0x8e, 0x32, 0xce, 0x52, 0x25, 0x49, 0xc3, 0xe8, 0x33, 0x2f, 0x68, 0xe1, 0x44, 0x3c, 0x3c, 0x25,
	0xcd, 0xc0, 0xd1, 0xc2, 0xe9, 0x18, 0x37, 0xa1, 0x79, 0x92, 0x9c, 0x0f, 0x32, 0x2e, 0x14, 0x71,
	0xcd, 0xfc, 0x8d, 0x93, 0xe4, 0xfc, 0x90, 0x0b, 0xa3, 0xf3, 0x7b, 0x2e, 0x15, 0x81, 0xbc, 0x5d,
	0xc7, 0xfa, 0x01, 0xc5, 0x12, 0x2a, 0x55, 0x9c, 0x64, 0xc4, 0x0b, 0x9c, 0x4e, 0x35, 0x9a, 0x17,
	0x34, 0xc2, 0x10, 0xb5, 0x0c, 0x91, 0x89, 0x91, 0x40, 0x63, 0x4a, 0x85, 0x64, 0x3c, 0x25, 0xed,
	0x9c, 0xdf, 0xa6, 0xfe, 0x23, 0x70, 0x2f, 0xbe, 0x63, 0x59, 0x3a, 0x37, 0x97, 0xee, 0x56, 0x59,
	0x3a, 0xb7, 0x24, 0x94, 0xff, 0x1a, 0x82, 0xbf, 0x7d, 0xa6, 0x7f, 0xe1, 0x0b, 0x1f, 0x42, 0xeb,
	0x88, 0x67, 0x6c, 0xf8, 0x27, 0x83, 0x23, 0xac, 0xa4, 0x71, 0x52, 0x80, 0x4d, 0x1c, 0x7e, 0x76,
	0xa0, 0x6d, 0x81, 0xd6, 0xe3, 0x4f, 0xa1, 0xae, 0x74, 0xa1, 0xb0, 0xf8, 0xbd, 0xb9, 0xbe, 0x0b,
	0x8d, 0x79, 0x66, 0x2d, 0x64, 0x21, 0x7a, 0x40, 0x4d, 0x9b, 0x3b, 0xdc, 0x8d, 0xf2, 0xc4, 0x7f,
	0x09, 0x5e, 0xa9, 0xf9, 0x37, 0x7b, 0xed, 0x2c, 0x5a, 0x7c, 0xed, 0xf2, 0x93, 0xa5, 0x45, 0xbf,
	0x3a, 0x50, 0x33, 0x45, 0x7c, 0xb0, 0x60, 0xf0, 0xcd, 0x4b, 0x98, 0x25, 0x7f, 0x17, 0xdb, 0xd7,
	0xe6, 0xdb, 0xe3, 0x36, 0x40, 0x16, 0x0b, 0xc5, 0xcc, 0xad, 0x20, 0x75, 0xa3, 0x6d, 0xa9, 0x82,
	0x01, 0x78, 0x82, 0x66, 0x13, 0x36, 0x34, 0xd7, 0x84, 0x34, 0x4c, 0x43, 0xb9, 0x74, 0x6d, 0x03,
	0xf4, 0x7f, 0x54, 0xa0, 0x19, 0xd9, 0x89, 0xf1, 0x08, 0x60, 0x9f, 0x2a, 0x7b, 0x07, 0x70, 0x63,
	0xf9, 0xa8, 0x18, 0x51, 0x7d, 0x72, 0xd5, 0xb5, 0x09, 0x6f, 0x7e, 0xfc, 0xf6, 0xfd, 0x53, 0xa5,
	0x8d, 0x5e, 0x6f, 0xba, 0xdb, 0x2b, 0x8e, 0xcd, 0x3b, 0xf0, 0xb4, 0xc7, 0xfe, 0x83, 0x96, 0x18,
	0x5a, 0xc4, 0xf5, 0x12, 0x6d, 0x4f, 0xff, 0x79, 0xf1, 0x10, 0xdc, 0x7d, 0xaa, 0x72, 0x55, 0xf1,
	0xce, 0x92, 0x45, 0x72, 0xe2, 0x8d, 0x2b, 0xac, 0x13, 0xa2, 0xe1, 0x6d, 0x21, 0x68, 0x5e, 0x6b,
	0x9d, 0xb7, 0x00, 0x7a, 0xda, 0xeb, 0x52, 0x6e, 0x18, 0xca, 0x1b, 0xb8, 0x36, 0xa7, 0x34, 0x93,
	0x1e, 0xd7, 0xcd, 0xcd, 0xda, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf0, 0x4a, 0x8a, 0x43,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	// GetBrokers returns a BrokerResponse with the brokers field populated
	// with full broker metadata. If the input BrokerRequest.id field is
	// non-nil, a single broker is returned matching the ID specified in the
	// Broker object. Otherwise all brokers are returned, optionally filtered
	// by any provided BrokerRequest.tags parameters.
	GetBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// ListBrokers returns a BrokerResponse with the ids field populated
	// with broker IDs. If the input BrokerRequest.id field is non-nil,
	// a single broker ID is returned matching the ID specified in the
	// Broker object if the broker exists. Otherwise all brokers are returned,
	// optionally filtered by any provided BrokerRequest.tags parameters.
	ListBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error)
	// GetTopics returns a TopicResponse with the topics field populated
	// with full topic metadata. If the input TopicRequest.name field is
	// non-nil, a single topic is returned matching the name specified in the
	// Topic object. Otherwise all topics are returned, optionally filtered
	// by any provided TopicRequest.tags parameters.
	GetTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	// ListTopics returns a TopicResponse with the names field populated
	// with topic names. If the input TopicRequest.name field is non-nil,
	// a single topic name is returned matching the name specified in the
	// Topic object if the topic exists. Otherwise all topics are returned,
	// optionally filtered by any provided TopicRequest.tags parameters.
	ListTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListBrokers(ctx context.Context, in *BrokerRequest, opts ...grpc.CallOption) (*BrokerResponse, error) {
	out := new(BrokerResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListBrokers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTopics(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	// GetBrokers returns a BrokerResponse with the brokers field populated
	// with full broker metadata. If the input BrokerRequest.id field is
	// non-nil, a single broker is returned matching the ID specified in the
	// Broker object. Otherwise all brokers are returned, optionally filtered
	// by any provided BrokerRequest.tags parameters.
	GetBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error)
	// ListBrokers returns a BrokerResponse with the ids field populated
	// with broker IDs. If the input BrokerRequest.id field is non-nil,
	// a single broker ID is returned matching the ID specified in the
	// Broker object if the broker exists. Otherwise all brokers are returned,
	// optionally filtered by any provided BrokerRequest.tags parameters.
	ListBrokers(context.Context, *BrokerRequest) (*BrokerResponse, error)
	// GetTopics returns a TopicResponse with the topics field populated
	// with full topic metadata. If the input TopicRequest.name field is
	// non-nil, a single topic is returned matching the name specified in the
	// Topic object. Otherwise all topics are returned, optionally filtered
	// by any provided TopicRequest.tags parameters.
	GetTopics(context.Context, *TopicRequest) (*TopicResponse, error)
	// ListTopics returns a TopicResponse with the names field populated
	// with topic names. If the input TopicRequest.name field is non-nil,
	// a single topic name is returned matching the name specified in the
	// Topic object if the topic exists. Otherwise all topics are returned,
	// optionally filtered by any provided TopicRequest.tags parameters.
	ListTopics(context.Context, *TopicRequest) (*TopicResponse, error)
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_GetBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetBrokers(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListBrokers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListBrokers(ctx, req.(*BrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetTopics(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTopics(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBrokers",
			Handler:    _Registry_GetBrokers_Handler,
		},
		{
			MethodName: "ListBrokers",
			Handler:    _Registry_ListBrokers_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _Registry_GetTopics_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _Registry_ListTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/registry.proto",
}
