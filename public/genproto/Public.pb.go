// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Public.proto

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

type CreatePublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday  string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation    string `protobuf:"bytes,6,opt,name=nation,proto3" json:"nation,omitempty"`
	PartyId   string `protobuf:"bytes,7,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *CreatePublicRequest) Reset() {
	*x = CreatePublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublicRequest) ProtoMessage() {}

func (x *CreatePublicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublicRequest.ProtoReflect.Descriptor instead.
func (*CreatePublicRequest) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePublicRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePublicRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreatePublicRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreatePublicRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *CreatePublicRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreatePublicRequest) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *CreatePublicRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type VoidPublicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidPublicResponse) Reset() {
	*x = VoidPublicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidPublicResponse) ProtoMessage() {}

func (x *VoidPublicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidPublicResponse.ProtoReflect.Descriptor instead.
func (*VoidPublicResponse) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{1}
}

type ByIdPublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIdPublicRequest) Reset() {
	*x = ByIdPublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIdPublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIdPublicRequest) ProtoMessage() {}

func (x *ByIdPublicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIdPublicRequest.ProtoReflect.Descriptor instead.
func (*ByIdPublicRequest) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{2}
}

func (x *ByIdPublicRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPublicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday  string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation    string `protobuf:"bytes,6,opt,name=nation,proto3" json:"nation,omitempty"`
	PartyId   string `protobuf:"bytes,7,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *GetPublicResponse) Reset() {
	*x = GetPublicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicResponse) ProtoMessage() {}

func (x *GetPublicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicResponse.ProtoReflect.Descriptor instead.
func (*GetPublicResponse) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{3}
}

func (x *GetPublicResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPublicResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetPublicResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetPublicResponse) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetPublicResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetPublicResponse) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *GetPublicResponse) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type UpdatePublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday  string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation    string `protobuf:"bytes,5,opt,name=nation,proto3" json:"nation,omitempty"`
	PartyId   string `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *UpdatePublicRequest) Reset() {
	*x = UpdatePublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublicRequest) ProtoMessage() {}

func (x *UpdatePublicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublicRequest.ProtoReflect.Descriptor instead.
func (*UpdatePublicRequest) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePublicRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdatePublicRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdatePublicRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdatePublicRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdatePublicRequest) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *UpdatePublicRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type FilterPublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FilterPublicRequest) Reset() {
	*x = FilterPublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterPublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterPublicRequest) ProtoMessage() {}

func (x *FilterPublicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterPublicRequest.ProtoReflect.Descriptor instead.
func (*FilterPublicRequest) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{5}
}

func (x *FilterPublicRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FilterPublicRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAllPublicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publics []*GetPublicResponse `protobuf:"bytes,1,rep,name=publics,proto3" json:"publics,omitempty"`
}

func (x *GetAllPublicResponse) Reset() {
	*x = GetAllPublicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Public_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublicResponse) ProtoMessage() {}

func (x *GetAllPublicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Public_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublicResponse.ProtoReflect.Descriptor instead.
func (*GetAllPublicResponse) Descriptor() ([]byte, []int) {
	return file_Public_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllPublicResponse) GetPublics() []*GetPublicResponse {
	if x != nil {
		return x.Publics
	}
	return nil
}

var File_Public_proto protoreflect.FileDescriptor

var file_Public_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xc8, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x42, 0x79, 0x49, 0x64, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc6, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x43, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x73, 0x32, 0x82, 0x03, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x19,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Public_proto_rawDescOnce sync.Once
	file_Public_proto_rawDescData = file_Public_proto_rawDesc
)

func file_Public_proto_rawDescGZIP() []byte {
	file_Public_proto_rawDescOnce.Do(func() {
		file_Public_proto_rawDescData = protoimpl.X.CompressGZIP(file_Public_proto_rawDescData)
	})
	return file_Public_proto_rawDescData
}

var file_Public_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Public_proto_goTypes = []interface{}{
	(*CreatePublicRequest)(nil),  // 0: voting.CreatePublicRequest
	(*VoidPublicResponse)(nil),   // 1: voting.VoidPublicResponse
	(*ByIdPublicRequest)(nil),    // 2: voting.ByIdPublicRequest
	(*GetPublicResponse)(nil),    // 3: voting.GetPublicResponse
	(*UpdatePublicRequest)(nil),  // 4: voting.UpdatePublicRequest
	(*FilterPublicRequest)(nil),  // 5: voting.FilterPublicRequest
	(*GetAllPublicResponse)(nil), // 6: voting.GetAllPublicResponse
}
var file_Public_proto_depIdxs = []int32{
	3, // 0: voting.GetAllPublicResponse.publics:type_name -> voting.GetPublicResponse
	0, // 1: voting.PublicService.CreatePublic:input_type -> voting.CreatePublicRequest
	2, // 2: voting.PublicService.GetByIdPublic:input_type -> voting.ByIdPublicRequest
	5, // 3: voting.PublicService.GetAllPublic:input_type -> voting.FilterPublicRequest
	3, // 4: voting.PublicService.UpdatePublic:input_type -> voting.GetPublicResponse
	2, // 5: voting.PublicService.DeletePublic:input_type -> voting.ByIdPublicRequest
	1, // 6: voting.PublicService.CreatePublic:output_type -> voting.VoidPublicResponse
	3, // 7: voting.PublicService.GetByIdPublic:output_type -> voting.GetPublicResponse
	6, // 8: voting.PublicService.GetAllPublic:output_type -> voting.GetAllPublicResponse
	1, // 9: voting.PublicService.UpdatePublic:output_type -> voting.VoidPublicResponse
	1, // 10: voting.PublicService.DeletePublic:output_type -> voting.VoidPublicResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Public_proto_init() }
func file_Public_proto_init() {
	if File_Public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePublicRequest); i {
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
		file_Public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidPublicResponse); i {
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
		file_Public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIdPublicRequest); i {
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
		file_Public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicResponse); i {
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
		file_Public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePublicRequest); i {
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
		file_Public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterPublicRequest); i {
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
		file_Public_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublicResponse); i {
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
			RawDescriptor: file_Public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Public_proto_goTypes,
		DependencyIndexes: file_Public_proto_depIdxs,
		MessageInfos:      file_Public_proto_msgTypes,
	}.Build()
	File_Public_proto = out.File
	file_Public_proto_rawDesc = nil
	file_Public_proto_goTypes = nil
	file_Public_proto_depIdxs = nil
}
