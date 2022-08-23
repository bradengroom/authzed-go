// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: authzed/api/v0/core.proto

package v0

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RelationTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each tupleset specifies keys of a set of relation tuples. The set can
	// include a single tuple key, or all tuples with a given object ID or
	// userset in a namespace, optionally constrained by a relation name.
	//
	// examples:
	// doc:readme#viewer@group:eng#member (fully specified)
	// doc:*#*#group:eng#member (all tuples that this userset relates to)
	// doc:12345#*#* (all tuples with a direct relationship to a document)
	// doc:12345#writer#* (all tuples with direct write relationship with the
	// document) doc:#writer#group:eng#member (all tuples that eng group has write
	// relationship)
	ObjectAndRelation *ObjectAndRelation `protobuf:"bytes,1,opt,name=object_and_relation,json=objectAndRelation,proto3" json:"object_and_relation,omitempty"`
	User              *User              `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RelationTuple) Reset() {
	*x = RelationTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTuple) ProtoMessage() {}

func (x *RelationTuple) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTuple.ProtoReflect.Descriptor instead.
func (*RelationTuple) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_core_proto_rawDescGZIP(), []int{0}
}

func (x *RelationTuple) GetObjectAndRelation() *ObjectAndRelation {
	if x != nil {
		return x.ObjectAndRelation
	}
	return nil
}

func (x *RelationTuple) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ObjectAndRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectId  string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Relation  string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ObjectAndRelation) Reset() {
	*x = ObjectAndRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectAndRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectAndRelation) ProtoMessage() {}

func (x *ObjectAndRelation) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectAndRelation.ProtoReflect.Descriptor instead.
func (*ObjectAndRelation) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_core_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectAndRelation) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectAndRelation) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *ObjectAndRelation) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type RelationReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Relation  string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *RelationReference) Reset() {
	*x = RelationReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationReference) ProtoMessage() {}

func (x *RelationReference) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationReference.ProtoReflect.Descriptor instead.
func (*RelationReference) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_core_proto_rawDescGZIP(), []int{2}
}

func (x *RelationReference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RelationReference) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UserOneof:
	//	*User_Userset
	UserOneof isUser_UserOneof `protobuf_oneof:"user_oneof"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_core_proto_rawDescGZIP(), []int{3}
}

func (m *User) GetUserOneof() isUser_UserOneof {
	if m != nil {
		return m.UserOneof
	}
	return nil
}

func (x *User) GetUserset() *ObjectAndRelation {
	if x, ok := x.GetUserOneof().(*User_Userset); ok {
		return x.Userset
	}
	return nil
}

type isUser_UserOneof interface {
	isUser_UserOneof()
}

type User_Userset struct {
	Userset *ObjectAndRelation `protobuf:"bytes,2,opt,name=userset,proto3,oneof"`
}

func (*User_Userset) isUser_UserOneof() {}

var File_authzed_api_v0_core_proto protoreflect.FileDescriptor

var file_authzed_api_v0_core_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x11, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x30, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x66, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61,
	0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36,
	0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f, 0x29, 0x3f, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xfa, 0x42, 0x32, 0x72, 0x30, 0x28,
	0x80, 0x01, 0x32, 0x2b, 0x5e, 0x28, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5f, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x7c,
	0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x32, 0x37, 0x7d, 0x29, 0x7c, 0x5c, 0x2a, 0x29, 0x24, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x42, 0x2d,
	0x72, 0x2b, 0x28, 0x40, 0x32, 0x27, 0x5e, 0x28, 0x5c, 0x2e, 0x5c, 0x2e, 0x5c, 0x2e, 0x7c, 0x5b,
	0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c,
	0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x66, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61,
	0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36,
	0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f, 0x29, 0x3f, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x42, 0x2d, 0x72, 0x2b, 0x28, 0x40,
	0x32, 0x27, 0x5e, 0x28, 0x5c, 0x2e, 0x5c, 0x2e, 0x5c, 0x2e, 0x7c, 0x5b, 0x61, 0x2d, 0x7a, 0x5d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x74, 0x42, 0x11, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x48, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v0_core_proto_rawDescOnce sync.Once
	file_authzed_api_v0_core_proto_rawDescData = file_authzed_api_v0_core_proto_rawDesc
)

func file_authzed_api_v0_core_proto_rawDescGZIP() []byte {
	file_authzed_api_v0_core_proto_rawDescOnce.Do(func() {
		file_authzed_api_v0_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v0_core_proto_rawDescData)
	})
	return file_authzed_api_v0_core_proto_rawDescData
}

var file_authzed_api_v0_core_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v0_core_proto_goTypes = []interface{}{
	(*RelationTuple)(nil),     // 0: authzed.api.v0.RelationTuple
	(*ObjectAndRelation)(nil), // 1: authzed.api.v0.ObjectAndRelation
	(*RelationReference)(nil), // 2: authzed.api.v0.RelationReference
	(*User)(nil),              // 3: authzed.api.v0.User
}
var file_authzed_api_v0_core_proto_depIdxs = []int32{
	1, // 0: authzed.api.v0.RelationTuple.object_and_relation:type_name -> authzed.api.v0.ObjectAndRelation
	3, // 1: authzed.api.v0.RelationTuple.user:type_name -> authzed.api.v0.User
	1, // 2: authzed.api.v0.User.userset:type_name -> authzed.api.v0.ObjectAndRelation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_authzed_api_v0_core_proto_init() }
func file_authzed_api_v0_core_proto_init() {
	if File_authzed_api_v0_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v0_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTuple); i {
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
		file_authzed_api_v0_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectAndRelation); i {
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
		file_authzed_api_v0_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationReference); i {
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
		file_authzed_api_v0_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
	file_authzed_api_v0_core_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*User_Userset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_v0_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v0_core_proto_goTypes,
		DependencyIndexes: file_authzed_api_v0_core_proto_depIdxs,
		MessageInfos:      file_authzed_api_v0_core_proto_msgTypes,
	}.Build()
	File_authzed_api_v0_core_proto = out.File
	file_authzed_api_v0_core_proto_rawDesc = nil
	file_authzed_api_v0_core_proto_goTypes = nil
	file_authzed_api_v0_core_proto_depIdxs = nil
}
