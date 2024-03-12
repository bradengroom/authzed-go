// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: authzed/api/v1/stats_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRelationshipCardinalityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relationship *Relationship `protobuf:"bytes,1,opt,name=relationship,proto3" json:"relationship,omitempty"`
}

func (x *GetRelationshipCardinalityRequest) Reset() {
	*x = GetRelationshipCardinalityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_stats_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationshipCardinalityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationshipCardinalityRequest) ProtoMessage() {}

func (x *GetRelationshipCardinalityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_stats_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationshipCardinalityRequest.ProtoReflect.Descriptor instead.
func (*GetRelationshipCardinalityRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_stats_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRelationshipCardinalityRequest) GetRelationship() *Relationship {
	if x != nil {
		return x.Relationship
	}
	return nil
}

type GetRelationshipCardinalityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cardinality float64 `protobuf:"fixed64,1,opt,name=cardinality,proto3" json:"cardinality,omitempty"`
}

func (x *GetRelationshipCardinalityResponse) Reset() {
	*x = GetRelationshipCardinalityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_stats_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationshipCardinalityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationshipCardinalityResponse) ProtoMessage() {}

func (x *GetRelationshipCardinalityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_stats_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationshipCardinalityResponse.ProtoReflect.Descriptor instead.
func (*GetRelationshipCardinalityResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_stats_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRelationshipCardinalityResponse) GetCardinality() float64 {
	if x != nil {
		return x.Cardinality
	}
	return 0
}

type UpdateRelationshipCardinalityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relationship *Relationship `protobuf:"bytes,1,opt,name=relationship,proto3" json:"relationship,omitempty"`
}

func (x *UpdateRelationshipCardinalityRequest) Reset() {
	*x = UpdateRelationshipCardinalityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_stats_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRelationshipCardinalityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRelationshipCardinalityRequest) ProtoMessage() {}

func (x *UpdateRelationshipCardinalityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_stats_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRelationshipCardinalityRequest.ProtoReflect.Descriptor instead.
func (*UpdateRelationshipCardinalityRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_stats_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRelationshipCardinalityRequest) GetRelationship() *Relationship {
	if x != nil {
		return x.Relationship
	}
	return nil
}

type UpdateRelationshipCardinalityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRelationshipCardinalityResponse) Reset() {
	*x = UpdateRelationshipCardinalityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_stats_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRelationshipCardinalityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRelationshipCardinalityResponse) ProtoMessage() {}

func (x *UpdateRelationshipCardinalityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_stats_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRelationshipCardinalityResponse.ProtoReflect.Descriptor instead.
func (*UpdateRelationshipCardinalityResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_stats_service_proto_rawDescGZIP(), []int{3}
}

var File_authzed_api_v1_stats_service_proto protoreflect.FileDescriptor

var file_authzed_api_v1_stats_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x21, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x22, 0x46, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x68, 0x0a, 0x24,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0x27, 0x0a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x8c, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xb6, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x30, 0x01, 0x12, 0xc2, 0x01, 0x0a, 0x1d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x34, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x43,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x30, 0x01, 0x42, 0x48,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v1_stats_service_proto_rawDescOnce sync.Once
	file_authzed_api_v1_stats_service_proto_rawDescData = file_authzed_api_v1_stats_service_proto_rawDesc
)

func file_authzed_api_v1_stats_service_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_stats_service_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_stats_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_stats_service_proto_rawDescData)
	})
	return file_authzed_api_v1_stats_service_proto_rawDescData
}

var file_authzed_api_v1_stats_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v1_stats_service_proto_goTypes = []interface{}{
	(*GetRelationshipCardinalityRequest)(nil),     // 0: authzed.api.v1.GetRelationshipCardinalityRequest
	(*GetRelationshipCardinalityResponse)(nil),    // 1: authzed.api.v1.GetRelationshipCardinalityResponse
	(*UpdateRelationshipCardinalityRequest)(nil),  // 2: authzed.api.v1.UpdateRelationshipCardinalityRequest
	(*UpdateRelationshipCardinalityResponse)(nil), // 3: authzed.api.v1.UpdateRelationshipCardinalityResponse
	(*Relationship)(nil),                          // 4: authzed.api.v1.Relationship
}
var file_authzed_api_v1_stats_service_proto_depIdxs = []int32{
	4, // 0: authzed.api.v1.GetRelationshipCardinalityRequest.relationship:type_name -> authzed.api.v1.Relationship
	4, // 1: authzed.api.v1.UpdateRelationshipCardinalityRequest.relationship:type_name -> authzed.api.v1.Relationship
	0, // 2: authzed.api.v1.StatsService.GetRelationshipCardinality:input_type -> authzed.api.v1.GetRelationshipCardinalityRequest
	2, // 3: authzed.api.v1.StatsService.UpdateRelationshipCardinality:input_type -> authzed.api.v1.UpdateRelationshipCardinalityRequest
	1, // 4: authzed.api.v1.StatsService.GetRelationshipCardinality:output_type -> authzed.api.v1.GetRelationshipCardinalityResponse
	3, // 5: authzed.api.v1.StatsService.UpdateRelationshipCardinality:output_type -> authzed.api.v1.UpdateRelationshipCardinalityResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_stats_service_proto_init() }
func file_authzed_api_v1_stats_service_proto_init() {
	if File_authzed_api_v1_stats_service_proto != nil {
		return
	}
	file_authzed_api_v1_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1_stats_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationshipCardinalityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authzed_api_v1_stats_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationshipCardinalityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authzed_api_v1_stats_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRelationshipCardinalityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authzed_api_v1_stats_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRelationshipCardinalityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_v1_stats_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v1_stats_service_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_stats_service_proto_depIdxs,
		MessageInfos:      file_authzed_api_v1_stats_service_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_stats_service_proto = out.File
	file_authzed_api_v1_stats_service_proto_rawDesc = nil
	file_authzed_api_v1_stats_service_proto_goTypes = nil
	file_authzed_api_v1_stats_service_proto_depIdxs = nil
}
