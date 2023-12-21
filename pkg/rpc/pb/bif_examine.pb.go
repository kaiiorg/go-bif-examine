// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: pkg/rpc/pb/bif_examine.proto

package pb

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

// /////////////////////////////////////
// GetAllProjects
// /////////////////////////////////////
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OriginalKeyFileName string `protobuf:"bytes,3,opt,name=original_key_file_name,json=originalKeyFileName,proto3" json:"original_key_file_name,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetOriginalKeyFileName() string {
	if x != nil {
		return x.OriginalKeyFileName
	}
	return ""
}

type GetAllProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllProjectsRequest) Reset() {
	*x = GetAllProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProjectsRequest) ProtoMessage() {}

func (x *GetAllProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProjectsRequest.ProtoReflect.Descriptor instead.
func (*GetAllProjectsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{1}
}

type GetAllProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorDescription string     `protobuf:"bytes,1,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	Projects         []*Project `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *GetAllProjectsResponse) Reset() {
	*x = GetAllProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProjectsResponse) ProtoMessage() {}

func (x *GetAllProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProjectsResponse.ProtoReflect.Descriptor instead.
func (*GetAllProjectsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllProjectsResponse) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

func (x *GetAllProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

// /////////////////////////////////////
// DeleteProject
// /////////////////////////////////////
type DeleteProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteProjectRequest) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type DeleteProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorDescription string `protobuf:"bytes,1,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
}

func (x *DeleteProjectResponse) Reset() {
	*x = DeleteProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectResponse) ProtoMessage() {}

func (x *DeleteProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteProjectResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProjectResponse) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

// /////////////////////////////////////
// UploadKey
// /////////////////////////////////////
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParsedVersion           string   `protobuf:"bytes,1,opt,name=parsed_version,json=parsedVersion,proto3" json:"parsed_version,omitempty"`
	ParsedSignature         string   `protobuf:"bytes,2,opt,name=parsed_signature,json=parsedSignature,proto3" json:"parsed_signature,omitempty"`
	ResourceEntryCount      uint32   `protobuf:"varint,3,opt,name=resource_entry_count,json=resourceEntryCount,proto3" json:"resource_entry_count,omitempty"`
	ResourcesWithAudio      uint32   `protobuf:"varint,5,opt,name=resources_with_audio,json=resourcesWithAudio,proto3" json:"resources_with_audio,omitempty"`
	BifFilesContainingAudio []string `protobuf:"bytes,4,rep,name=bif_files_containing_audio,json=bifFilesContainingAudio,proto3" json:"bif_files_containing_audio,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{5}
}

func (x *Key) GetParsedVersion() string {
	if x != nil {
		return x.ParsedVersion
	}
	return ""
}

func (x *Key) GetParsedSignature() string {
	if x != nil {
		return x.ParsedSignature
	}
	return ""
}

func (x *Key) GetResourceEntryCount() uint32 {
	if x != nil {
		return x.ResourceEntryCount
	}
	return 0
}

func (x *Key) GetResourcesWithAudio() uint32 {
	if x != nil {
		return x.ResourcesWithAudio
	}
	return 0
}

func (x *Key) GetBifFilesContainingAudio() []string {
	if x != nil {
		return x.BifFilesContainingAudio
	}
	return nil
}

type UploadKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	FileName    string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Contents    []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *UploadKeyRequest) Reset() {
	*x = UploadKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadKeyRequest) ProtoMessage() {}

func (x *UploadKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadKeyRequest.ProtoReflect.Descriptor instead.
func (*UploadKeyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{6}
}

func (x *UploadKeyRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *UploadKeyRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadKeyRequest) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type UploadKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorDescription string `protobuf:"bytes,1,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	ProjectId        uint32 `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Key              *Key   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UploadKeyResponse) Reset() {
	*x = UploadKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadKeyResponse) ProtoMessage() {}

func (x *UploadKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadKeyResponse.ProtoReflect.Descriptor instead.
func (*UploadKeyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{7}
}

func (x *UploadKeyResponse) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

func (x *UploadKeyResponse) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UploadKeyResponse) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

// /////////////////////////////////////
// UploadBif
// /////////////////////////////////////
type UploadBifRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	FileName  string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	NameInKey string `protobuf:"bytes,3,opt,name=name_in_key,json=nameInKey,proto3" json:"name_in_key,omitempty"`
	Contents  []byte `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *UploadBifRequest) Reset() {
	*x = UploadBifRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBifRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBifRequest) ProtoMessage() {}

func (x *UploadBifRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBifRequest.ProtoReflect.Descriptor instead.
func (*UploadBifRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{8}
}

func (x *UploadBifRequest) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UploadBifRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadBifRequest) GetNameInKey() string {
	if x != nil {
		return x.NameInKey
	}
	return ""
}

func (x *UploadBifRequest) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type UploadBifResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorDescription string `protobuf:"bytes,1,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
}

func (x *UploadBifResponse) Reset() {
	*x = UploadBifResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBifResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBifResponse) ProtoMessage() {}

func (x *UploadBifResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_pb_bif_examine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBifResponse.ProtoReflect.Descriptor instead.
func (*UploadBifResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP(), []int{9}
}

func (x *UploadBifResponse) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

var File_pkg_rpc_pb_bif_examine_proto protoreflect.FileDescriptor

var file_pkg_rpc_pb_bif_examine_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x69, 0x66,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x62, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x01, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x3b, 0x0a, 0x1a, 0x62, 0x69,
	0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17,
	0x62, 0x69, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x22, 0x6e, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7a, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x40, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x8f, 0x02, 0x0a, 0x0a, 0x42, 0x69, 0x66, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e,
	0x65, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69, 0x66, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x69, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_pb_bif_examine_proto_rawDescOnce sync.Once
	file_pkg_rpc_pb_bif_examine_proto_rawDescData = file_pkg_rpc_pb_bif_examine_proto_rawDesc
)

func file_pkg_rpc_pb_bif_examine_proto_rawDescGZIP() []byte {
	file_pkg_rpc_pb_bif_examine_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_pb_bif_examine_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_pb_bif_examine_proto_rawDescData)
	})
	return file_pkg_rpc_pb_bif_examine_proto_rawDescData
}

var file_pkg_rpc_pb_bif_examine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_rpc_pb_bif_examine_proto_goTypes = []interface{}{
	(*Project)(nil),                // 0: pb.Project
	(*GetAllProjectsRequest)(nil),  // 1: pb.GetAllProjectsRequest
	(*GetAllProjectsResponse)(nil), // 2: pb.GetAllProjectsResponse
	(*DeleteProjectRequest)(nil),   // 3: pb.DeleteProjectRequest
	(*DeleteProjectResponse)(nil),  // 4: pb.DeleteProjectResponse
	(*Key)(nil),                    // 5: pb.Key
	(*UploadKeyRequest)(nil),       // 6: pb.UploadKeyRequest
	(*UploadKeyResponse)(nil),      // 7: pb.UploadKeyResponse
	(*UploadBifRequest)(nil),       // 8: pb.UploadBifRequest
	(*UploadBifResponse)(nil),      // 9: pb.UploadBifResponse
}
var file_pkg_rpc_pb_bif_examine_proto_depIdxs = []int32{
	0, // 0: pb.GetAllProjectsResponse.projects:type_name -> pb.Project
	5, // 1: pb.UploadKeyResponse.key:type_name -> pb.Key
	1, // 2: pb.BifExamine.GetAllProjects:input_type -> pb.GetAllProjectsRequest
	3, // 3: pb.BifExamine.DeleteProject:input_type -> pb.DeleteProjectRequest
	6, // 4: pb.BifExamine.UploadKey:input_type -> pb.UploadKeyRequest
	8, // 5: pb.BifExamine.UploadBif:input_type -> pb.UploadBifRequest
	2, // 6: pb.BifExamine.GetAllProjects:output_type -> pb.GetAllProjectsResponse
	4, // 7: pb.BifExamine.DeleteProject:output_type -> pb.DeleteProjectResponse
	7, // 8: pb.BifExamine.UploadKey:output_type -> pb.UploadKeyResponse
	9, // 9: pb.BifExamine.UploadBif:output_type -> pb.UploadBifResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_rpc_pb_bif_examine_proto_init() }
func file_pkg_rpc_pb_bif_examine_proto_init() {
	if File_pkg_rpc_pb_bif_examine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProjectsRequest); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProjectsResponse); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectRequest); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectResponse); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadKeyRequest); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadKeyResponse); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBifRequest); i {
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
		file_pkg_rpc_pb_bif_examine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBifResponse); i {
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
			RawDescriptor: file_pkg_rpc_pb_bif_examine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_pb_bif_examine_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_pb_bif_examine_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_pb_bif_examine_proto_msgTypes,
	}.Build()
	File_pkg_rpc_pb_bif_examine_proto = out.File
	file_pkg_rpc_pb_bif_examine_proto_rawDesc = nil
	file_pkg_rpc_pb_bif_examine_proto_goTypes = nil
	file_pkg_rpc_pb_bif_examine_proto_depIdxs = nil
}
