// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.3
// source: blood.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Blood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Blood) Reset() {
	*x = Blood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blood) ProtoMessage() {}

func (x *Blood) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blood.ProtoReflect.Descriptor instead.
func (*Blood) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{0}
}

func (x *Blood) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blood) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Blood) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Blood) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Blood) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AddBloodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddBloodRequest) Reset() {
	*x = AddBloodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBloodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBloodRequest) ProtoMessage() {}

func (x *AddBloodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBloodRequest.ProtoReflect.Descriptor instead.
func (*AddBloodRequest) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{1}
}

func (x *AddBloodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddBloodRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddBloodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blood *Blood `protobuf:"bytes,1,opt,name=Blood,proto3" json:"Blood,omitempty"`
}

func (x *AddBloodResponse) Reset() {
	*x = AddBloodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBloodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBloodResponse) ProtoMessage() {}

func (x *AddBloodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBloodResponse.ProtoReflect.Descriptor instead.
func (*AddBloodResponse) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{2}
}

func (x *AddBloodResponse) GetBlood() *Blood {
	if x != nil {
		return x.Blood
	}
	return nil
}

type GetBloodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBloodRequest) Reset() {
	*x = GetBloodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBloodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBloodRequest) ProtoMessage() {}

func (x *GetBloodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBloodRequest.ProtoReflect.Descriptor instead.
func (*GetBloodRequest) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{3}
}

type GetBloodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bloods []*Blood `protobuf:"bytes,1,rep,name=Bloods,proto3" json:"Bloods,omitempty"`
}

func (x *GetBloodResponse) Reset() {
	*x = GetBloodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBloodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBloodResponse) ProtoMessage() {}

func (x *GetBloodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBloodResponse.ProtoReflect.Descriptor instead.
func (*GetBloodResponse) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{4}
}

func (x *GetBloodResponse) GetBloods() []*Blood {
	if x != nil {
		return x.Bloods
	}
	return nil
}

type GetBloodByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBloodByIDRequest) Reset() {
	*x = GetBloodByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBloodByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBloodByIDRequest) ProtoMessage() {}

func (x *GetBloodByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBloodByIDRequest.ProtoReflect.Descriptor instead.
func (*GetBloodByIDRequest) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{5}
}

func (x *GetBloodByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBloodByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blood *Blood `protobuf:"bytes,1,opt,name=Blood,proto3" json:"Blood,omitempty"`
}

func (x *GetBloodByIDResponse) Reset() {
	*x = GetBloodByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blood_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBloodByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBloodByIDResponse) ProtoMessage() {}

func (x *GetBloodByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blood_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBloodByIDResponse.ProtoReflect.Descriptor instead.
func (*GetBloodByIDResponse) Descriptor() ([]byte, []int) {
	return file_blood_proto_rawDescGZIP(), []int{6}
}

func (x *GetBloodByIDResponse) GetBlood() *Blood {
	if x != nil {
		return x.Blood
	}
	return nil
}

var File_blood_proto protoreflect.FileDescriptor

var file_blood_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x47, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x73,
	0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x32, 0xd8, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x6f, 0x64,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blood_proto_rawDescOnce sync.Once
	file_blood_proto_rawDescData = file_blood_proto_rawDesc
)

func file_blood_proto_rawDescGZIP() []byte {
	file_blood_proto_rawDescOnce.Do(func() {
		file_blood_proto_rawDescData = protoimpl.X.CompressGZIP(file_blood_proto_rawDescData)
	})
	return file_blood_proto_rawDescData
}

var file_blood_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_blood_proto_goTypes = []interface{}{
	(*Blood)(nil),                // 0: proto.Blood
	(*AddBloodRequest)(nil),      // 1: proto.AddBloodRequest
	(*AddBloodResponse)(nil),     // 2: proto.AddBloodResponse
	(*GetBloodRequest)(nil),      // 3: proto.GetBloodRequest
	(*GetBloodResponse)(nil),     // 4: proto.GetBloodResponse
	(*GetBloodByIDRequest)(nil),  // 5: proto.GetBloodByIDRequest
	(*GetBloodByIDResponse)(nil), // 6: proto.GetBloodByIDResponse
}
var file_blood_proto_depIdxs = []int32{
	0, // 0: proto.AddBloodResponse.Blood:type_name -> proto.Blood
	0, // 1: proto.GetBloodResponse.Bloods:type_name -> proto.Blood
	0, // 2: proto.GetBloodByIDResponse.Blood:type_name -> proto.Blood
	1, // 3: proto.BloodService.AddBlood:input_type -> proto.AddBloodRequest
	3, // 4: proto.BloodService.GetBloods:input_type -> proto.GetBloodRequest
	5, // 5: proto.BloodService.GetBloodByID:input_type -> proto.GetBloodByIDRequest
	2, // 6: proto.BloodService.AddBlood:output_type -> proto.AddBloodResponse
	4, // 7: proto.BloodService.GetBloods:output_type -> proto.GetBloodResponse
	6, // 8: proto.BloodService.GetBloodByID:output_type -> proto.GetBloodByIDResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blood_proto_init() }
func file_blood_proto_init() {
	if File_blood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blood); i {
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
		file_blood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBloodRequest); i {
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
		file_blood_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBloodResponse); i {
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
		file_blood_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBloodRequest); i {
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
		file_blood_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBloodResponse); i {
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
		file_blood_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBloodByIDRequest); i {
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
		file_blood_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBloodByIDResponse); i {
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
			RawDescriptor: file_blood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blood_proto_goTypes,
		DependencyIndexes: file_blood_proto_depIdxs,
		MessageInfos:      file_blood_proto_msgTypes,
	}.Build()
	File_blood_proto = out.File
	file_blood_proto_rawDesc = nil
	file_blood_proto_goTypes = nil
	file_blood_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BloodServiceClient is the client API for BloodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BloodServiceClient interface {
	AddBlood(ctx context.Context, in *AddBloodRequest, opts ...grpc.CallOption) (*AddBloodResponse, error)
	GetBloods(ctx context.Context, in *GetBloodRequest, opts ...grpc.CallOption) (*GetBloodResponse, error)
	GetBloodByID(ctx context.Context, in *GetBloodByIDRequest, opts ...grpc.CallOption) (*GetBloodByIDResponse, error)
}

type bloodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBloodServiceClient(cc grpc.ClientConnInterface) BloodServiceClient {
	return &bloodServiceClient{cc}
}

func (c *bloodServiceClient) AddBlood(ctx context.Context, in *AddBloodRequest, opts ...grpc.CallOption) (*AddBloodResponse, error) {
	out := new(AddBloodResponse)
	err := c.cc.Invoke(ctx, "/proto.BloodService/AddBlood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloodServiceClient) GetBloods(ctx context.Context, in *GetBloodRequest, opts ...grpc.CallOption) (*GetBloodResponse, error) {
	out := new(GetBloodResponse)
	err := c.cc.Invoke(ctx, "/proto.BloodService/GetBloods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloodServiceClient) GetBloodByID(ctx context.Context, in *GetBloodByIDRequest, opts ...grpc.CallOption) (*GetBloodByIDResponse, error) {
	out := new(GetBloodByIDResponse)
	err := c.cc.Invoke(ctx, "/proto.BloodService/GetBloodByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BloodServiceServer is the server API for BloodService service.
type BloodServiceServer interface {
	AddBlood(context.Context, *AddBloodRequest) (*AddBloodResponse, error)
	GetBloods(context.Context, *GetBloodRequest) (*GetBloodResponse, error)
	GetBloodByID(context.Context, *GetBloodByIDRequest) (*GetBloodByIDResponse, error)
}

// UnimplementedBloodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBloodServiceServer struct {
}

func (*UnimplementedBloodServiceServer) AddBlood(context.Context, *AddBloodRequest) (*AddBloodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlood not implemented")
}
func (*UnimplementedBloodServiceServer) GetBloods(context.Context, *GetBloodRequest) (*GetBloodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBloods not implemented")
}
func (*UnimplementedBloodServiceServer) GetBloodByID(context.Context, *GetBloodByIDRequest) (*GetBloodByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBloodByID not implemented")
}

func RegisterBloodServiceServer(s *grpc.Server, srv BloodServiceServer) {
	s.RegisterService(&_BloodService_serviceDesc, srv)
}

func _BloodService_AddBlood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBloodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloodServiceServer).AddBlood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BloodService/AddBlood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloodServiceServer).AddBlood(ctx, req.(*AddBloodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloodService_GetBloods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBloodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloodServiceServer).GetBloods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BloodService/GetBloods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloodServiceServer).GetBloods(ctx, req.(*GetBloodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloodService_GetBloodByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBloodByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloodServiceServer).GetBloodByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BloodService/GetBloodByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloodServiceServer).GetBloodByID(ctx, req.(*GetBloodByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BloodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BloodService",
	HandlerType: (*BloodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlood",
			Handler:    _BloodService_AddBlood_Handler,
		},
		{
			MethodName: "GetBloods",
			Handler:    _BloodService_GetBloods_Handler,
		},
		{
			MethodName: "GetBloodByID",
			Handler:    _BloodService_GetBloodByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blood.proto",
}