// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v1/user-service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_user_service_proto protoreflect.FileDescriptor

var file_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x16, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x2d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2d, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x04, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x71, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x27, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6b, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f,
	0x6b, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_v1_user_service_proto_goTypes = []interface{}{
	(*GetUserInfoRequest)(nil),                   // 0: v1.GetUserInfoRequest
	(*SetUserNameRequest)(nil),                   // 1: v1.SetUserNameRequest
	(*SetUserDescriptionRequest)(nil),            // 2: v1.SetUserDescriptionRequest
	(*GenerateUserAvatarUploadLinkRequest)(nil),  // 3: v1.GenerateUserAvatarUploadLinkRequest
	(*ConfirmUserAvatarUploadingRequest)(nil),    // 4: v1.ConfirmUserAvatarUploadingRequest
	(*DeleteUserAvatarRequest)(nil),              // 5: v1.DeleteUserAvatarRequest
	(*GetUserInfoResponse)(nil),                  // 6: v1.GetUserInfoResponse
	(*SetUserNameResponse)(nil),                  // 7: v1.SetUserNameResponse
	(*SetUserDescriptionResponse)(nil),           // 8: v1.SetUserDescriptionResponse
	(*GenerateUserAvatarUploadLinkResponse)(nil), // 9: v1.GenerateUserAvatarUploadLinkResponse
	(*ConfirmUserAvatarUploadingResponse)(nil),   // 10: v1.ConfirmUserAvatarUploadingResponse
	(*DeleteUserAvatarResponse)(nil),             // 11: v1.DeleteUserAvatarResponse
}
var file_v1_user_service_proto_depIdxs = []int32{
	0,  // 0: v1.UserService.GetUserInfo:input_type -> v1.GetUserInfoRequest
	1,  // 1: v1.UserService.SetUserName:input_type -> v1.SetUserNameRequest
	2,  // 2: v1.UserService.SetUserDescription:input_type -> v1.SetUserDescriptionRequest
	3,  // 3: v1.UserService.GenerateUserAvatarUploadLink:input_type -> v1.GenerateUserAvatarUploadLinkRequest
	4,  // 4: v1.UserService.ConfirmUserAvatarUploading:input_type -> v1.ConfirmUserAvatarUploadingRequest
	5,  // 5: v1.UserService.DeleteUserAvatar:input_type -> v1.DeleteUserAvatarRequest
	6,  // 6: v1.UserService.GetUserInfo:output_type -> v1.GetUserInfoResponse
	7,  // 7: v1.UserService.SetUserName:output_type -> v1.SetUserNameResponse
	8,  // 8: v1.UserService.SetUserDescription:output_type -> v1.SetUserDescriptionResponse
	9,  // 9: v1.UserService.GenerateUserAvatarUploadLink:output_type -> v1.GenerateUserAvatarUploadLinkResponse
	10, // 10: v1.UserService.ConfirmUserAvatarUploading:output_type -> v1.ConfirmUserAvatarUploadingResponse
	11, // 11: v1.UserService.DeleteUserAvatar:output_type -> v1.DeleteUserAvatarResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_user_service_proto_init() }
func file_v1_user_service_proto_init() {
	if File_v1_user_service_proto != nil {
		return
	}
	file_v1_get_user_info_proto_init()
	file_v1_set_user_name_proto_init()
	file_v1_set_user_description_proto_init()
	file_v1_generate_user_avatar_upload_link_proto_init()
	file_v1_confirm_user_avatar_uploading_proto_init()
	file_v1_delete_user_avatar_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_service_proto_goTypes,
		DependencyIndexes: file_v1_user_service_proto_depIdxs,
	}.Build()
	File_v1_user_service_proto = out.File
	file_v1_user_service_proto_rawDesc = nil
	file_v1_user_service_proto_goTypes = nil
	file_v1_user_service_proto_depIdxs = nil
}
