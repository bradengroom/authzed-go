// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: authzed/api/v1/schema_service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// ReadSchemaRequest returns the schema from the database.
type ReadSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadSchemaRequest) Reset() {
	*x = ReadSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_schema_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaRequest) ProtoMessage() {}

func (x *ReadSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_schema_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaRequest.ProtoReflect.Descriptor instead.
func (*ReadSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_schema_service_proto_rawDescGZIP(), []int{0}
}

// ReadSchemaResponse is the resulting data after having read the Object
// Definitions from a Schema.
type ReadSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schema_text is the textual form of the current schema in the system
	SchemaText string `protobuf:"bytes,1,opt,name=schema_text,json=schemaText,proto3" json:"schema_text,omitempty"`
}

func (x *ReadSchemaResponse) Reset() {
	*x = ReadSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_schema_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaResponse) ProtoMessage() {}

func (x *ReadSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_schema_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaResponse.ProtoReflect.Descriptor instead.
func (*ReadSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadSchemaResponse) GetSchemaText() string {
	if x != nil {
		return x.SchemaText
	}
	return ""
}

// WriteSchemaRequest is the required data used to "upsert" the Schema of a
// Permissions System.
type WriteSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Schema containing one or more Object Definitions that will be written
	// to the Permissions System.
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"` // 256KiB
}

func (x *WriteSchemaRequest) Reset() {
	*x = WriteSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_schema_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaRequest) ProtoMessage() {}

func (x *WriteSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_schema_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaRequest.ProtoReflect.Descriptor instead.
func (*WriteSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *WriteSchemaRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

// WriteSchemaResponse is the resulting data after having written a Schema to
// a Permissions System.
type WriteSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteSchemaResponse) Reset() {
	*x = WriteSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_schema_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaResponse) ProtoMessage() {}

func (x *WriteSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_schema_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaResponse.ProtoReflect.Descriptor instead.
func (*WriteSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_schema_service_proto_rawDescGZIP(), []int{3}
}

var File_authzed_api_v1_schema_service_proto protoreflect.FileDescriptor

var file_authzed_api_v1_schema_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x35, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x54, 0x65, 0x78, 0x74, 0x22, 0x37, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x28, 0x80, 0x80, 0x10, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x22, 0x15, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf5, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0a, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x73, 0x0a, 0x0b, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x48, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_authzed_api_v1_schema_service_proto_rawDescOnce sync.Once
	file_authzed_api_v1_schema_service_proto_rawDescData = file_authzed_api_v1_schema_service_proto_rawDesc
)

func file_authzed_api_v1_schema_service_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_schema_service_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_schema_service_proto_rawDescData)
	})
	return file_authzed_api_v1_schema_service_proto_rawDescData
}

var file_authzed_api_v1_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v1_schema_service_proto_goTypes = []interface{}{
	(*ReadSchemaRequest)(nil),   // 0: authzed.api.v1.ReadSchemaRequest
	(*ReadSchemaResponse)(nil),  // 1: authzed.api.v1.ReadSchemaResponse
	(*WriteSchemaRequest)(nil),  // 2: authzed.api.v1.WriteSchemaRequest
	(*WriteSchemaResponse)(nil), // 3: authzed.api.v1.WriteSchemaResponse
}
var file_authzed_api_v1_schema_service_proto_depIdxs = []int32{
	0, // 0: authzed.api.v1.SchemaService.ReadSchema:input_type -> authzed.api.v1.ReadSchemaRequest
	2, // 1: authzed.api.v1.SchemaService.WriteSchema:input_type -> authzed.api.v1.WriteSchemaRequest
	1, // 2: authzed.api.v1.SchemaService.ReadSchema:output_type -> authzed.api.v1.ReadSchemaResponse
	3, // 3: authzed.api.v1.SchemaService.WriteSchema:output_type -> authzed.api.v1.WriteSchemaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_schema_service_proto_init() }
func file_authzed_api_v1_schema_service_proto_init() {
	if File_authzed_api_v1_schema_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1_schema_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaRequest); i {
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
		file_authzed_api_v1_schema_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaResponse); i {
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
		file_authzed_api_v1_schema_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaRequest); i {
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
		file_authzed_api_v1_schema_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaResponse); i {
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
			RawDescriptor: file_authzed_api_v1_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v1_schema_service_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_schema_service_proto_depIdxs,
		MessageInfos:      file_authzed_api_v1_schema_service_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_schema_service_proto = out.File
	file_authzed_api_v1_schema_service_proto_rawDesc = nil
	file_authzed_api_v1_schema_service_proto_goTypes = nil
	file_authzed_api_v1_schema_service_proto_depIdxs = nil
}
