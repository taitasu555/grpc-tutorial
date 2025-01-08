//
//  CommandService用のParam型とResult型を定義したprotoファイル
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: command.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CRUD int32

const (
	CRUD_UNKNOWN CRUD = 0
	CRUD_INSERT  CRUD = 1
	CRUD_UPDATE  CRUD = 2
	CRUD_DELETE  CRUD = 3
)

// Enum value maps for CRUD.
var (
	CRUD_name = map[int32]string{
		0: "UNKNOWN",
		1: "INSERT",
		2: "UPDATE",
		3: "DELETE",
	}
	CRUD_value = map[string]int32{
		"UNKNOWN": 0,
		"INSERT":  1,
		"UPDATE":  2,
		"DELETE":  3,
	}
)

func (x CRUD) Enum() *CRUD {
	p := new(CRUD)
	*p = x
	return p
}

func (x CRUD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CRUD) Descriptor() protoreflect.EnumDescriptor {
	return file_command_proto_enumTypes[0].Descriptor()
}

func (CRUD) Type() protoreflect.EnumType {
	return &file_command_proto_enumTypes[0]
}

func (x CRUD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CRUD.Descriptor instead.
func (CRUD) EnumDescriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

type CategoryUpParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Crud          CRUD                   `protobuf:"varint,1,opt,name=crud,proto3,enum=proto.CRUD" json:"crud,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryUpParam) Reset() {
	*x = CategoryUpParam{}
	mi := &file_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryUpParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpParam) ProtoMessage() {}

func (x *CategoryUpParam) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpParam.ProtoReflect.Descriptor instead.
func (*CategoryUpParam) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryUpParam) GetCrud() CRUD {
	if x != nil {
		return x.Crud
	}
	return CRUD_UNKNOWN
}

func (x *CategoryUpParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryUpParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductUpParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Crud          CRUD                   `protobuf:"varint,1,opt,name=crud,proto3,enum=proto.CRUD" json:"crud,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price         int32                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	CategoryId    string                 `protobuf:"bytes,5,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductUpParam) Reset() {
	*x = ProductUpParam{}
	mi := &file_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductUpParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpParam) ProtoMessage() {}

func (x *ProductUpParam) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpParam.ProtoReflect.Descriptor instead.
func (*ProductUpParam) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (x *ProductUpParam) GetCrud() CRUD {
	if x != nil {
		return x.Crud
	}
	return CRUD_UNKNOWN
}

func (x *ProductUpParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductUpParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductUpParam) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductUpParam) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type CategoryUpResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      *Category              `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Error         *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryUpResult) Reset() {
	*x = CategoryUpResult{}
	mi := &file_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryUpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpResult) ProtoMessage() {}

func (x *CategoryUpResult) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpResult.ProtoReflect.Descriptor instead.
func (*CategoryUpResult) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryUpResult) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *CategoryUpResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CategoryUpResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type ProductUpResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Error         *Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductUpResult) Reset() {
	*x = ProductUpResult{}
	mi := &file_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductUpResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpResult) ProtoMessage() {}

func (x *ProductUpResult) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpResult.ProtoReflect.Descriptor instead.
func (*ProductUpResult) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

func (x *ProductUpResult) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductUpResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ProductUpResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x72, 0x75, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x52, 0x55, 0x44,
	0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x72, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2a, 0x37, 0x0a, 0x04, 0x43, 0x52, 0x55, 0x44, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53,
	0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x32, 0xc2, 0x01,
	0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55,
	0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xbb, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55,
	0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55,
	0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x55, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_command_proto_goTypes = []any{
	(CRUD)(0),                     // 0: proto.CRUD
	(*CategoryUpParam)(nil),       // 1: proto.CategoryUpParam
	(*ProductUpParam)(nil),        // 2: proto.ProductUpParam
	(*CategoryUpResult)(nil),      // 3: proto.CategoryUpResult
	(*ProductUpResult)(nil),       // 4: proto.ProductUpResult
	(*Category)(nil),              // 5: proto.Category
	(*Error)(nil),                 // 6: proto.Error
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*Product)(nil),               // 8: proto.Product
}
var file_command_proto_depIdxs = []int32{
	0,  // 0: proto.CategoryUpParam.crud:type_name -> proto.CRUD
	0,  // 1: proto.ProductUpParam.crud:type_name -> proto.CRUD
	5,  // 2: proto.CategoryUpResult.category:type_name -> proto.Category
	6,  // 3: proto.CategoryUpResult.error:type_name -> proto.Error
	7,  // 4: proto.CategoryUpResult.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 5: proto.ProductUpResult.product:type_name -> proto.Product
	6,  // 6: proto.ProductUpResult.error:type_name -> proto.Error
	7,  // 7: proto.ProductUpResult.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 8: proto.CategoryCommand.Create:input_type -> proto.CategoryUpParam
	1,  // 9: proto.CategoryCommand.Update:input_type -> proto.CategoryUpParam
	1,  // 10: proto.CategoryCommand.Delete:input_type -> proto.CategoryUpParam
	2,  // 11: proto.ProductCommand.Create:input_type -> proto.ProductUpParam
	2,  // 12: proto.ProductCommand.Update:input_type -> proto.ProductUpParam
	2,  // 13: proto.ProductCommand.Delete:input_type -> proto.ProductUpParam
	3,  // 14: proto.CategoryCommand.Create:output_type -> proto.CategoryUpResult
	3,  // 15: proto.CategoryCommand.Update:output_type -> proto.CategoryUpResult
	3,  // 16: proto.CategoryCommand.Delete:output_type -> proto.CategoryUpResult
	4,  // 17: proto.ProductCommand.Create:output_type -> proto.ProductUpResult
	4,  // 18: proto.ProductCommand.Update:output_type -> proto.ProductUpResult
	4,  // 19: proto.ProductCommand.Delete:output_type -> proto.ProductUpResult
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	file_models_proto_init()
	file_error_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		EnumInfos:         file_command_proto_enumTypes,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
