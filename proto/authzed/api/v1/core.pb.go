// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: authzed/api/v1/core.proto

package v1

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

type RelationshipUpdate_Operation int32

const (
	RelationshipUpdate_OPERATION_UNSPECIFIED RelationshipUpdate_Operation = 0
	RelationshipUpdate_OPERATION_CREATE      RelationshipUpdate_Operation = 1
	RelationshipUpdate_OPERATION_TOUCH       RelationshipUpdate_Operation = 2
	RelationshipUpdate_OPERATION_DELETE      RelationshipUpdate_Operation = 3
)

// Enum value maps for RelationshipUpdate_Operation.
var (
	RelationshipUpdate_Operation_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "OPERATION_CREATE",
		2: "OPERATION_TOUCH",
		3: "OPERATION_DELETE",
	}
	RelationshipUpdate_Operation_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"OPERATION_CREATE":      1,
		"OPERATION_TOUCH":       2,
		"OPERATION_DELETE":      3,
	}
)

func (x RelationshipUpdate_Operation) Enum() *RelationshipUpdate_Operation {
	p := new(RelationshipUpdate_Operation)
	*p = x
	return p
}

func (x RelationshipUpdate_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationshipUpdate_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_core_proto_enumTypes[0].Descriptor()
}

func (RelationshipUpdate_Operation) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_core_proto_enumTypes[0]
}

func (x RelationshipUpdate_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationshipUpdate_Operation.Descriptor instead.
func (RelationshipUpdate_Operation) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{4, 0}
}

type AlgebraicSubjectSet_Operation int32

const (
	AlgebraicSubjectSet_OPERATION_UNSPECIFIED  AlgebraicSubjectSet_Operation = 0
	AlgebraicSubjectSet_OPERATION_UNION        AlgebraicSubjectSet_Operation = 1
	AlgebraicSubjectSet_OPERATION_INTERSECTION AlgebraicSubjectSet_Operation = 2
	AlgebraicSubjectSet_OPERATION_EXCLUSION    AlgebraicSubjectSet_Operation = 3
)

// Enum value maps for AlgebraicSubjectSet_Operation.
var (
	AlgebraicSubjectSet_Operation_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "OPERATION_UNION",
		2: "OPERATION_INTERSECTION",
		3: "OPERATION_EXCLUSION",
	}
	AlgebraicSubjectSet_Operation_value = map[string]int32{
		"OPERATION_UNSPECIFIED":  0,
		"OPERATION_UNION":        1,
		"OPERATION_INTERSECTION": 2,
		"OPERATION_EXCLUSION":    3,
	}
)

func (x AlgebraicSubjectSet_Operation) Enum() *AlgebraicSubjectSet_Operation {
	p := new(AlgebraicSubjectSet_Operation)
	*p = x
	return p
}

func (x AlgebraicSubjectSet_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlgebraicSubjectSet_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_core_proto_enumTypes[1].Descriptor()
}

func (AlgebraicSubjectSet_Operation) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_core_proto_enumTypes[1]
}

func (x AlgebraicSubjectSet_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlgebraicSubjectSet_Operation.Descriptor instead.
func (AlgebraicSubjectSet_Operation) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{6, 0}
}

// Relationship specifies how a resource relates to a subject. Relationships
// form the data for the graph over which all permissions questions are
// answered.
type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource is the resource to which the subject is related, in some manner
	Resource *ObjectReference `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// relation is how the resource and subject are related.
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	// subject is the subject to which the resource is related, in some manner.
	Subject *SubjectReference `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *Relationship) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Relationship) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Relationship) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

// SubjectReference is used for referring to the subject portion of a
// Relationship. The relation component is optional and is used for defining a
// sub-relation on the subject, e.g. group:123#members
type SubjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object           *ObjectReference `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	OptionalRelation string           `protobuf:"bytes,2,opt,name=optional_relation,json=optionalRelation,proto3" json:"optional_relation,omitempty"`
}

func (x *SubjectReference) Reset() {
	*x = SubjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectReference) ProtoMessage() {}

func (x *SubjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectReference.ProtoReflect.Descriptor instead.
func (*SubjectReference) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *SubjectReference) GetObject() *ObjectReference {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *SubjectReference) GetOptionalRelation() string {
	if x != nil {
		return x.OptionalRelation
	}
	return ""
}

// ObjectReference is used to refer to a specific object in the system.
type ObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	ObjectId   string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *ObjectReference) Reset() {
	*x = ObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectReference) ProtoMessage() {}

func (x *ObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectReference.ProtoReflect.Descriptor instead.
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectReference) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *ObjectReference) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

// ZedToken is used to provide causality metadata between Write and Check
// requests.
//
// See the authzed.api.v1.Consistency message for more information.
type ZedToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ZedToken) Reset() {
	*x = ZedToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZedToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZedToken) ProtoMessage() {}

func (x *ZedToken) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZedToken.ProtoReflect.Descriptor instead.
func (*ZedToken) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *ZedToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// RelationshipUpdate is used for mutating a single relationship within the
// service.
//
// CREATE will create the relationship only if it doesn't exist, and error
// otherwise.
//
// TOUCH will upsert the relationship, and will not error if it
// already exists.
//
// DELETE will delete the relationship and error if it doesn't
// exist.
type RelationshipUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation    RelationshipUpdate_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=authzed.api.v1.RelationshipUpdate_Operation" json:"operation,omitempty"`
	Relationship *Relationship                `protobuf:"bytes,2,opt,name=relationship,proto3" json:"relationship,omitempty"`
}

func (x *RelationshipUpdate) Reset() {
	*x = RelationshipUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationshipUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationshipUpdate) ProtoMessage() {}

func (x *RelationshipUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationshipUpdate.ProtoReflect.Descriptor instead.
func (*RelationshipUpdate) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{4}
}

func (x *RelationshipUpdate) GetOperation() RelationshipUpdate_Operation {
	if x != nil {
		return x.Operation
	}
	return RelationshipUpdate_OPERATION_UNSPECIFIED
}

func (x *RelationshipUpdate) GetRelationship() *Relationship {
	if x != nil {
		return x.Relationship
	}
	return nil
}

// PermissionRelationshipTree is used for representing a tree of a resource and
// its permission relationships with other objects.
type PermissionRelationshipTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TreeType:
	//	*PermissionRelationshipTree_Intermediate
	//	*PermissionRelationshipTree_Leaf
	TreeType         isPermissionRelationshipTree_TreeType `protobuf_oneof:"tree_type"`
	ExpandedObject   *ObjectReference                      `protobuf:"bytes,3,opt,name=expanded_object,json=expandedObject,proto3" json:"expanded_object,omitempty"`
	ExpandedRelation string                                `protobuf:"bytes,4,opt,name=expanded_relation,json=expandedRelation,proto3" json:"expanded_relation,omitempty"`
}

func (x *PermissionRelationshipTree) Reset() {
	*x = PermissionRelationshipTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionRelationshipTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRelationshipTree) ProtoMessage() {}

func (x *PermissionRelationshipTree) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRelationshipTree.ProtoReflect.Descriptor instead.
func (*PermissionRelationshipTree) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{5}
}

func (m *PermissionRelationshipTree) GetTreeType() isPermissionRelationshipTree_TreeType {
	if m != nil {
		return m.TreeType
	}
	return nil
}

func (x *PermissionRelationshipTree) GetIntermediate() *AlgebraicSubjectSet {
	if x, ok := x.GetTreeType().(*PermissionRelationshipTree_Intermediate); ok {
		return x.Intermediate
	}
	return nil
}

func (x *PermissionRelationshipTree) GetLeaf() *DirectSubjectSet {
	if x, ok := x.GetTreeType().(*PermissionRelationshipTree_Leaf); ok {
		return x.Leaf
	}
	return nil
}

func (x *PermissionRelationshipTree) GetExpandedObject() *ObjectReference {
	if x != nil {
		return x.ExpandedObject
	}
	return nil
}

func (x *PermissionRelationshipTree) GetExpandedRelation() string {
	if x != nil {
		return x.ExpandedRelation
	}
	return ""
}

type isPermissionRelationshipTree_TreeType interface {
	isPermissionRelationshipTree_TreeType()
}

type PermissionRelationshipTree_Intermediate struct {
	Intermediate *AlgebraicSubjectSet `protobuf:"bytes,1,opt,name=intermediate,proto3,oneof"`
}

type PermissionRelationshipTree_Leaf struct {
	Leaf *DirectSubjectSet `protobuf:"bytes,2,opt,name=leaf,proto3,oneof"`
}

func (*PermissionRelationshipTree_Intermediate) isPermissionRelationshipTree_TreeType() {}

func (*PermissionRelationshipTree_Leaf) isPermissionRelationshipTree_TreeType() {}

// AlgebraicSubjectSet is a subject set which is computed based on applying the
// specified operation to the operands according to the algebra of sets.
//
// UNION is a logical set containing the subject members from all operands.
//
// INTERSECTION is a logical set containing only the subject members which are
// present in all operands.
//
// EXCLUSION is a logical set containing only the subject members which are
// present in the first operand, and none of the other operands.
type AlgebraicSubjectSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation AlgebraicSubjectSet_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=authzed.api.v1.AlgebraicSubjectSet_Operation" json:"operation,omitempty"`
	Children  []*PermissionRelationshipTree `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *AlgebraicSubjectSet) Reset() {
	*x = AlgebraicSubjectSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgebraicSubjectSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgebraicSubjectSet) ProtoMessage() {}

func (x *AlgebraicSubjectSet) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgebraicSubjectSet.ProtoReflect.Descriptor instead.
func (*AlgebraicSubjectSet) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{6}
}

func (x *AlgebraicSubjectSet) GetOperation() AlgebraicSubjectSet_Operation {
	if x != nil {
		return x.Operation
	}
	return AlgebraicSubjectSet_OPERATION_UNSPECIFIED
}

func (x *AlgebraicSubjectSet) GetChildren() []*PermissionRelationshipTree {
	if x != nil {
		return x.Children
	}
	return nil
}

// DirectSubjectSet is a subject set which is simply a collection of subjects.
type DirectSubjectSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subjects []*SubjectReference `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty"`
}

func (x *DirectSubjectSet) Reset() {
	*x = DirectSubjectSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectSubjectSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectSubjectSet) ProtoMessage() {}

func (x *DirectSubjectSet) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectSubjectSet.ProtoReflect.Descriptor instead.
func (*DirectSubjectSet) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_core_proto_rawDescGZIP(), []int{7}
}

func (x *DirectSubjectSet) GetSubjects() []*SubjectReference {
	if x != nil {
		return x.Subjects
	}
	return nil
}

var File_authzed_api_v1_core_proto protoreflect.FileDescriptor

var file_authzed_api_v1_core_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27,
	0xfa, 0x42, 0x24, 0x72, 0x22, 0x28, 0x40, 0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x57, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2a, 0xfa, 0x42, 0x27, 0x72, 0x25, 0x28, 0x40, 0x32, 0x21, 0x5e, 0x28, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x24, 0x52, 0x10, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc8,
	0x01, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x69, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80,
	0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x2f, 0x29, 0x3f, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x24, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2d, 0xfa, 0x42, 0x2a, 0x72, 0x28, 0x28, 0x80, 0x01, 0x32, 0x23, 0x5e, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x32, 0x37, 0x7d, 0x24, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x08, 0x5a, 0x65, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0x67, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x55, 0x43, 0x48, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0xa8, 0x02, 0x0a, 0x1a, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x54, 0x72, 0x65, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x67,
	0x65, 0x62, 0x72, 0x61, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x12, 0x36, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x12, 0x48, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x10, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x22, 0x9c, 0x02, 0x0a, 0x13, 0x41, 0x6c, 0x67, 0x65, 0x62, 0x72, 0x61, 0x69, 0x63, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x12, 0x4b, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x67, 0x65, 0x62, 0x72, 0x61, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x70,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x53, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x22, 0x50, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x42, 0x48, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v1_core_proto_rawDescOnce sync.Once
	file_authzed_api_v1_core_proto_rawDescData = file_authzed_api_v1_core_proto_rawDesc
)

func file_authzed_api_v1_core_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_core_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_core_proto_rawDescData)
	})
	return file_authzed_api_v1_core_proto_rawDescData
}

var file_authzed_api_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_authzed_api_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_authzed_api_v1_core_proto_goTypes = []interface{}{
	(RelationshipUpdate_Operation)(0),  // 0: authzed.api.v1.RelationshipUpdate.Operation
	(AlgebraicSubjectSet_Operation)(0), // 1: authzed.api.v1.AlgebraicSubjectSet.Operation
	(*Relationship)(nil),               // 2: authzed.api.v1.Relationship
	(*SubjectReference)(nil),           // 3: authzed.api.v1.SubjectReference
	(*ObjectReference)(nil),            // 4: authzed.api.v1.ObjectReference
	(*ZedToken)(nil),                   // 5: authzed.api.v1.ZedToken
	(*RelationshipUpdate)(nil),         // 6: authzed.api.v1.RelationshipUpdate
	(*PermissionRelationshipTree)(nil), // 7: authzed.api.v1.PermissionRelationshipTree
	(*AlgebraicSubjectSet)(nil),        // 8: authzed.api.v1.AlgebraicSubjectSet
	(*DirectSubjectSet)(nil),           // 9: authzed.api.v1.DirectSubjectSet
}
var file_authzed_api_v1_core_proto_depIdxs = []int32{
	4,  // 0: authzed.api.v1.Relationship.resource:type_name -> authzed.api.v1.ObjectReference
	3,  // 1: authzed.api.v1.Relationship.subject:type_name -> authzed.api.v1.SubjectReference
	4,  // 2: authzed.api.v1.SubjectReference.object:type_name -> authzed.api.v1.ObjectReference
	0,  // 3: authzed.api.v1.RelationshipUpdate.operation:type_name -> authzed.api.v1.RelationshipUpdate.Operation
	2,  // 4: authzed.api.v1.RelationshipUpdate.relationship:type_name -> authzed.api.v1.Relationship
	8,  // 5: authzed.api.v1.PermissionRelationshipTree.intermediate:type_name -> authzed.api.v1.AlgebraicSubjectSet
	9,  // 6: authzed.api.v1.PermissionRelationshipTree.leaf:type_name -> authzed.api.v1.DirectSubjectSet
	4,  // 7: authzed.api.v1.PermissionRelationshipTree.expanded_object:type_name -> authzed.api.v1.ObjectReference
	1,  // 8: authzed.api.v1.AlgebraicSubjectSet.operation:type_name -> authzed.api.v1.AlgebraicSubjectSet.Operation
	7,  // 9: authzed.api.v1.AlgebraicSubjectSet.children:type_name -> authzed.api.v1.PermissionRelationshipTree
	3,  // 10: authzed.api.v1.DirectSubjectSet.subjects:type_name -> authzed.api.v1.SubjectReference
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_core_proto_init() }
func file_authzed_api_v1_core_proto_init() {
	if File_authzed_api_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relationship); i {
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
		file_authzed_api_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectReference); i {
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
		file_authzed_api_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectReference); i {
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
		file_authzed_api_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZedToken); i {
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
		file_authzed_api_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationshipUpdate); i {
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
		file_authzed_api_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionRelationshipTree); i {
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
		file_authzed_api_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgebraicSubjectSet); i {
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
		file_authzed_api_v1_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectSubjectSet); i {
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
	file_authzed_api_v1_core_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*PermissionRelationshipTree_Intermediate)(nil),
		(*PermissionRelationshipTree_Leaf)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_v1_core_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v1_core_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_core_proto_depIdxs,
		EnumInfos:         file_authzed_api_v1_core_proto_enumTypes,
		MessageInfos:      file_authzed_api_v1_core_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_core_proto = out.File
	file_authzed_api_v1_core_proto_rawDesc = nil
	file_authzed_api_v1_core_proto_goTypes = nil
	file_authzed_api_v1_core_proto_depIdxs = nil
}
