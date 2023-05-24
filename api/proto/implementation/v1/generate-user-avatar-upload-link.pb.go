// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v1/generate-user-avatar-upload-link.proto

package v1

import (
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

type GenerateUserAvatarUploadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GenerateUserAvatarUploadLinkRequest) Reset() {
	*x = GenerateUserAvatarUploadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_generate_user_avatar_upload_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateUserAvatarUploadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateUserAvatarUploadLinkRequest) ProtoMessage() {}

func (x *GenerateUserAvatarUploadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_generate_user_avatar_upload_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateUserAvatarUploadLinkRequest.ProtoReflect.Descriptor instead.
func (*GenerateUserAvatarUploadLinkRequest) Descriptor() ([]byte, []int) {
	return file_v1_generate_user_avatar_upload_link_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateUserAvatarUploadLinkRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GenerateUserAvatarUploadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link     string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	AvatarId string `protobuf:"bytes,2,opt,name=avatarId,proto3" json:"avatarId,omitempty"`
}

func (x *GenerateUserAvatarUploadLinkResponse) Reset() {
	*x = GenerateUserAvatarUploadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_generate_user_avatar_upload_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateUserAvatarUploadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateUserAvatarUploadLinkResponse) ProtoMessage() {}

func (x *GenerateUserAvatarUploadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_generate_user_avatar_upload_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateUserAvatarUploadLinkResponse.ProtoReflect.Descriptor instead.
func (*GenerateUserAvatarUploadLinkResponse) Descriptor() ([]byte, []int) {
	return file_v1_generate_user_avatar_upload_link_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateUserAvatarUploadLinkResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *GenerateUserAvatarUploadLinkResponse) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

var File_v1_generate_user_avatar_upload_link_proto protoreflect.FileDescriptor

var file_v1_generate_user_avatar_upload_link_proto_rawDesc = []byte{
	0x0a, 0x29, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22,
	0x3d, 0x0a, 0x23, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56,
	0x0a, 0x24, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_generate_user_avatar_upload_link_proto_rawDescOnce sync.Once
	file_v1_generate_user_avatar_upload_link_proto_rawDescData = file_v1_generate_user_avatar_upload_link_proto_rawDesc
)

func file_v1_generate_user_avatar_upload_link_proto_rawDescGZIP() []byte {
	file_v1_generate_user_avatar_upload_link_proto_rawDescOnce.Do(func() {
		file_v1_generate_user_avatar_upload_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_generate_user_avatar_upload_link_proto_rawDescData)
	})
	return file_v1_generate_user_avatar_upload_link_proto_rawDescData
}

var file_v1_generate_user_avatar_upload_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_generate_user_avatar_upload_link_proto_goTypes = []interface{}{
	(*GenerateUserAvatarUploadLinkRequest)(nil),  // 0: v1.GenerateUserAvatarUploadLinkRequest
	(*GenerateUserAvatarUploadLinkResponse)(nil), // 1: v1.GenerateUserAvatarUploadLinkResponse
}
var file_v1_generate_user_avatar_upload_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_generate_user_avatar_upload_link_proto_init() }
func file_v1_generate_user_avatar_upload_link_proto_init() {
	if File_v1_generate_user_avatar_upload_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_generate_user_avatar_upload_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateUserAvatarUploadLinkRequest); i {
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
		file_v1_generate_user_avatar_upload_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateUserAvatarUploadLinkResponse); i {
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
			RawDescriptor: file_v1_generate_user_avatar_upload_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_generate_user_avatar_upload_link_proto_goTypes,
		DependencyIndexes: file_v1_generate_user_avatar_upload_link_proto_depIdxs,
		MessageInfos:      file_v1_generate_user_avatar_upload_link_proto_msgTypes,
	}.Build()
	File_v1_generate_user_avatar_upload_link_proto = out.File
	file_v1_generate_user_avatar_upload_link_proto_rawDesc = nil
	file_v1_generate_user_avatar_upload_link_proto_goTypes = nil
	file_v1_generate_user_avatar_upload_link_proto_depIdxs = nil
}
