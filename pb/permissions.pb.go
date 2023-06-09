// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pb/permissions.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPermissionsConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*PermissionConfig `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Roles       []*RoleConfig       `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Groups      []*GroupConfig      `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetPermissionsConfigResponse) Reset() {
	*x = GetPermissionsConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionsConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsConfigResponse) ProtoMessage() {}

func (x *GetPermissionsConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsConfigResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionsConfigResponse) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *GetPermissionsConfigResponse) GetPermissions() []*PermissionConfig {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *GetPermissionsConfigResponse) GetRoles() []*RoleConfig {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GetPermissionsConfigResponse) GetGroups() []*GroupConfig {
	if x != nil {
		return x.Groups
	}
	return nil
}

type PermissionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PermissionConfig) Reset() {
	*x = PermissionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionConfig) ProtoMessage() {}

func (x *PermissionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionConfig.ProtoReflect.Descriptor instead.
func (*PermissionConfig) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RoleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description     string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PermissionNames []string `protobuf:"bytes,3,rep,name=permission_names,json=permissionNames,proto3" json:"permission_names,omitempty"`
}

func (x *RoleConfig) Reset() {
	*x = RoleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleConfig) ProtoMessage() {}

func (x *RoleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleConfig.ProtoReflect.Descriptor instead.
func (*RoleConfig) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *RoleConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoleConfig) GetPermissionNames() []string {
	if x != nil {
		return x.PermissionNames
	}
	return nil
}

type GroupConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RoleNames   []string `protobuf:"bytes,3,rep,name=role_names,json=roleNames,proto3" json:"role_names,omitempty"`
}

func (x *GroupConfig) Reset() {
	*x = GroupConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupConfig) ProtoMessage() {}

func (x *GroupConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupConfig.ProtoReflect.Descriptor instead.
func (*GroupConfig) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *GroupConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GroupConfig) GetRoleNames() []string {
	if x != nil {
		return x.RoleNames
	}
	return nil
}

var File_pb_permissions_proto protoreflect.FileDescriptor

var file_pb_permissions_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x24, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22,
	0x48, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0a, 0x52, 0x6f, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x5f, 0x0a, 0x0b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x76, 0x61,
	0x53, 0x6f, 0x66, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_permissions_proto_rawDescOnce sync.Once
	file_pb_permissions_proto_rawDescData = file_pb_permissions_proto_rawDesc
)

func file_pb_permissions_proto_rawDescGZIP() []byte {
	file_pb_permissions_proto_rawDescOnce.Do(func() {
		file_pb_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_permissions_proto_rawDescData)
	})
	return file_pb_permissions_proto_rawDescData
}

var file_pb_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_permissions_proto_goTypes = []interface{}{
	(*GetPermissionsConfigResponse)(nil), // 0: pb.GetPermissionsConfigResponse
	(*PermissionConfig)(nil),             // 1: pb.PermissionConfig
	(*RoleConfig)(nil),                   // 2: pb.RoleConfig
	(*GroupConfig)(nil),                  // 3: pb.GroupConfig
	(*emptypb.Empty)(nil),                // 4: google.protobuf.Empty
}
var file_pb_permissions_proto_depIdxs = []int32{
	1, // 0: pb.GetPermissionsConfigResponse.permissions:type_name -> pb.PermissionConfig
	2, // 1: pb.GetPermissionsConfigResponse.roles:type_name -> pb.RoleConfig
	3, // 2: pb.GetPermissionsConfigResponse.groups:type_name -> pb.GroupConfig
	4, // 3: pb.Permissions.GetPermissionsConfig:input_type -> google.protobuf.Empty
	0, // 4: pb.Permissions.GetPermissionsConfig:output_type -> pb.GetPermissionsConfigResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_permissions_proto_init() }
func file_pb_permissions_proto_init() {
	if File_pb_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionsConfigResponse); i {
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
		file_pb_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionConfig); i {
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
		file_pb_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleConfig); i {
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
		file_pb_permissions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupConfig); i {
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
			RawDescriptor: file_pb_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_permissions_proto_goTypes,
		DependencyIndexes: file_pb_permissions_proto_depIdxs,
		MessageInfos:      file_pb_permissions_proto_msgTypes,
	}.Build()
	File_pb_permissions_proto = out.File
	file_pb_permissions_proto_rawDesc = nil
	file_pb_permissions_proto_goTypes = nil
	file_pb_permissions_proto_depIdxs = nil
}
