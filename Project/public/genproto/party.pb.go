// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: party.proto

package genproto

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

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyRequest.ProtoReflect.Descriptor instead.
func (*CreatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePartyRequest) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *CreatePartyRequest) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *CreatePartyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetPartyResponse) Reset() {
	*x = GetPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyResponse) ProtoMessage() {}

func (x *GetPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyResponse.ProtoReflect.Descriptor instead.
func (*GetPartyResponse) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPartyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPartyResponse) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *GetPartyResponse) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *GetPartyResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties []*GetPartyResponse `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
}

func (x *GetAllPartyResponse) Reset() {
	*x = GetAllPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPartyResponse) ProtoMessage() {}

func (x *GetAllPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPartyResponse.ProtoReflect.Descriptor instead.
func (*GetAllPartyResponse) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPartyResponse) GetParties() []*GetPartyResponse {
	if x != nil {
		return x.Parties
	}
	return nil
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,2,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,3,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePartyRequest) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *UpdatePartyRequest) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *UpdatePartyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ByIdPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIdPartyRequest) Reset() {
	*x = ByIdPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIdPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIdPartyRequest) ProtoMessage() {}

func (x *ByIdPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIdPartyRequest.ProtoReflect.Descriptor instead.
func (*ByIdPartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{4}
}

func (x *ByIdPartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VoidPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidPartyResponse) Reset() {
	*x = VoidPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidPartyResponse) ProtoMessage() {}

func (x *VoidPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidPartyResponse.ProtoReflect.Descriptor instead.
func (*VoidPartyResponse) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{5}
}

type FilterPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FilterPartyRequest) Reset() {
	*x = FilterPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterPartyRequest) ProtoMessage() {}

func (x *FilterPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterPartyRequest.ProtoReflect.Descriptor instead.
func (*FilterPartyRequest) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{6}
}

func (x *FilterPartyRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FilterPartyRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_party_proto protoreflect.FileDescriptor

var file_party_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x22, 0x0a, 0x10, 0x42, 0x79, 0x49, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xf4, 0x02,
	0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x79, 0x49, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x18, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_proto_rawDescOnce sync.Once
	file_party_proto_rawDescData = file_party_proto_rawDesc
)

func file_party_proto_rawDescGZIP() []byte {
	file_party_proto_rawDescOnce.Do(func() {
		file_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_proto_rawDescData)
	})
	return file_party_proto_rawDescData
}

var file_party_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_party_proto_goTypes = []interface{}{
	(*CreatePartyRequest)(nil),  // 0: voting.CreatePartyRequest
	(*GetPartyResponse)(nil),    // 1: voting.GetPartyResponse
	(*GetAllPartyResponse)(nil), // 2: voting.GetAllPartyResponse
	(*UpdatePartyRequest)(nil),  // 3: voting.UpdatePartyRequest
	(*ByIdPartyRequest)(nil),    // 4: voting.ByIdPartyRequest
	(*VoidPartyResponse)(nil),   // 5: voting.VoidPartyResponse
	(*FilterPartyRequest)(nil),  // 6: voting.FilterPartyRequest
}
var file_party_proto_depIdxs = []int32{
	1, // 0: voting.GetAllPartyResponse.parties:type_name -> voting.GetPartyResponse
	0, // 1: voting.PartyService.CreateParty:input_type -> voting.CreatePartyRequest
	4, // 2: voting.PartyService.GetPartyById:input_type -> voting.ByIdPartyRequest
	6, // 3: voting.PartyService.GetAllParties:input_type -> voting.FilterPartyRequest
	1, // 4: voting.PartyService.UpdateParty:input_type -> voting.GetPartyResponse
	4, // 5: voting.PartyService.DeleteParty:input_type -> voting.ByIdPartyRequest
	5, // 6: voting.PartyService.CreateParty:output_type -> voting.VoidPartyResponse
	1, // 7: voting.PartyService.GetPartyById:output_type -> voting.GetPartyResponse
	2, // 8: voting.PartyService.GetAllParties:output_type -> voting.GetAllPartyResponse
	5, // 9: voting.PartyService.UpdateParty:output_type -> voting.VoidPartyResponse
	5, // 10: voting.PartyService.DeleteParty:output_type -> voting.VoidPartyResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_proto_init() }
func file_party_proto_init() {
	if File_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyRequest); i {
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
		file_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyResponse); i {
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
		file_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPartyResponse); i {
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
		file_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartyRequest); i {
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
		file_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIdPartyRequest); i {
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
		file_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidPartyResponse); i {
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
		file_party_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterPartyRequest); i {
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
			RawDescriptor: file_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_proto_goTypes,
		DependencyIndexes: file_party_proto_depIdxs,
		MessageInfos:      file_party_proto_msgTypes,
	}.Build()
	File_party_proto = out.File
	file_party_proto_rawDesc = nil
	file_party_proto_goTypes = nil
	file_party_proto_depIdxs = nil
}
