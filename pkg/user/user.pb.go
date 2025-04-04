// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateUserPersonalDataIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MiddleName    string                 `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	City          string                 `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Gender        string                 `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Age           int32                  `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserPersonalDataIn) Reset() {
	*x = UpdateUserPersonalDataIn{}
	mi := &file_api_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserPersonalDataIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPersonalDataIn) ProtoMessage() {}

func (x *UpdateUserPersonalDataIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPersonalDataIn.ProtoReflect.Descriptor instead.
func (*UpdateUserPersonalDataIn) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserPersonalDataIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserPersonalDataIn) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UpdateUserPersonalDataOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserPersonalDataOut) Reset() {
	*x = UpdateUserPersonalDataOut{}
	mi := &file_api_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserPersonalDataOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPersonalDataOut) ProtoMessage() {}

func (x *UpdateUserPersonalDataOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPersonalDataOut.ProtoReflect.Descriptor instead.
func (*UpdateUserPersonalDataOut) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserPersonalDataOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_user_proto protoreflect.FileDescriptor

var file_api_user_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xce, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x35, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x60, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x19, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x1a, 0x1a, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_user_proto_rawDescOnce sync.Once
	file_api_user_proto_rawDescData []byte
)

func file_api_user_proto_rawDescGZIP() []byte {
	file_api_user_proto_rawDescOnce.Do(func() {
		file_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)))
	})
	return file_api_user_proto_rawDescData
}

var file_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_user_proto_goTypes = []any{
	(*UpdateUserPersonalDataIn)(nil),  // 0: UpdateUserPersonalDataIn
	(*UpdateUserPersonalDataOut)(nil), // 1: UpdateUserPersonalDataOut
}
var file_api_user_proto_depIdxs = []int32{
	0, // 0: UserService.UpdateUserPersonalData:input_type -> UpdateUserPersonalDataIn
	1, // 1: UserService.UpdateUserPersonalData:output_type -> UpdateUserPersonalDataOut
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_proto_init() }
func file_api_user_proto_init() {
	if File_api_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_proto_goTypes,
		DependencyIndexes: file_api_user_proto_depIdxs,
		MessageInfos:      file_api_user_proto_msgTypes,
	}.Build()
	File_api_user_proto = out.File
	file_api_user_proto_goTypes = nil
	file_api_user_proto_depIdxs = nil
}
