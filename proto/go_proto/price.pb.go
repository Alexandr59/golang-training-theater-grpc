// protoc -I . price.proto --go_out=plugins=grpc:.
// protoc -I . price.proto --grpc-gateway_out . --go_out=plugins=grpc:.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: price.proto

package go_proto

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccountId     int64 `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	SectorId      int64 `protobuf:"varint,3,opt,name=SectorId,proto3" json:"SectorId,omitempty"`
	PerformanceId int64 `protobuf:"varint,4,opt,name=PerformanceId,proto3" json:"PerformanceId,omitempty"`
	Price         int64 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PriceRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *PriceRequest) GetSectorId() int64 {
	if x != nil {
		return x.SectorId
	}
	return 0
}

func (x *PriceRequest) GetPerformanceId() int64 {
	if x != nil {
		return x.PerformanceId
	}
	return 0
}

func (x *PriceRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccountId     int64 `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	SectorId      int64 `protobuf:"varint,3,opt,name=SectorId,proto3" json:"SectorId,omitempty"`
	PerformanceId int64 `protobuf:"varint,4,opt,name=PerformanceId,proto3" json:"PerformanceId,omitempty"`
	Price         int64 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PriceResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *PriceResponse) GetSectorId() int64 {
	if x != nil {
		return x.SectorId
	}
	return 0
}

func (x *PriceResponse) GetPerformanceId() int64 {
	if x != nil {
		return x.PerformanceId
	}
	return 0
}

func (x *PriceResponse) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type IdPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdPriceRequest) Reset() {
	*x = IdPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdPriceRequest) ProtoMessage() {}

func (x *IdPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdPriceRequest.ProtoReflect.Descriptor instead.
func (*IdPriceRequest) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{2}
}

func (x *IdPriceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdPriceResponse) Reset() {
	*x = IdPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdPriceResponse) ProtoMessage() {}

func (x *IdPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdPriceResponse.ProtoReflect.Descriptor instead.
func (*IdPriceResponse) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{3}
}

func (x *IdPriceResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StatusPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusPriceResponse) Reset() {
	*x = StatusPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusPriceResponse) ProtoMessage() {}

func (x *StatusPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusPriceResponse.ProtoReflect.Descriptor instead.
func (*StatusPriceResponse) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{4}
}

func (x *StatusPriceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_price_proto protoreflect.FileDescriptor

var file_price_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x20, 0x0a, 0x0e, 0x49, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf1, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5c,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x1a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_price_proto_rawDescOnce sync.Once
	file_price_proto_rawDescData = file_price_proto_rawDesc
)

func file_price_proto_rawDescGZIP() []byte {
	file_price_proto_rawDescOnce.Do(func() {
		file_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_proto_rawDescData)
	})
	return file_price_proto_rawDescData
}

var file_price_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_price_proto_goTypes = []interface{}{
	(*PriceRequest)(nil),        // 0: price.PriceRequest
	(*PriceResponse)(nil),       // 1: price.PriceResponse
	(*IdPriceRequest)(nil),      // 2: price.IdPriceRequest
	(*IdPriceResponse)(nil),     // 3: price.IdPriceResponse
	(*StatusPriceResponse)(nil), // 4: price.StatusPriceResponse
}
var file_price_proto_depIdxs = []int32{
	0, // 0: price.PriceService.CreatePrice:input_type -> price.PriceRequest
	2, // 1: price.PriceService.DeletePrice:input_type -> price.IdPriceRequest
	0, // 2: price.PriceService.UpdatePrice:input_type -> price.PriceRequest
	2, // 3: price.PriceService.GetPrice:input_type -> price.IdPriceRequest
	3, // 4: price.PriceService.CreatePrice:output_type -> price.IdPriceResponse
	4, // 5: price.PriceService.DeletePrice:output_type -> price.StatusPriceResponse
	4, // 6: price.PriceService.UpdatePrice:output_type -> price.StatusPriceResponse
	1, // 7: price.PriceService.GetPrice:output_type -> price.PriceResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_price_proto_init() }
func file_price_proto_init() {
	if File_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceResponse); i {
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
		file_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdPriceRequest); i {
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
		file_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdPriceResponse); i {
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
		file_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusPriceResponse); i {
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
			RawDescriptor: file_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_price_proto_goTypes,
		DependencyIndexes: file_price_proto_depIdxs,
		MessageInfos:      file_price_proto_msgTypes,
	}.Build()
	File_price_proto = out.File
	file_price_proto_rawDesc = nil
	file_price_proto_goTypes = nil
	file_price_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PriceServiceClient interface {
	CreatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*IdPriceResponse, error)
	DeletePrice(ctx context.Context, in *IdPriceRequest, opts ...grpc.CallOption) (*StatusPriceResponse, error)
	UpdatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*StatusPriceResponse, error)
	GetPrice(ctx context.Context, in *IdPriceRequest, opts ...grpc.CallOption) (*PriceResponse, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) CreatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*IdPriceResponse, error) {
	out := new(IdPriceResponse)
	err := c.cc.Invoke(ctx, "/price.PriceService/CreatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) DeletePrice(ctx context.Context, in *IdPriceRequest, opts ...grpc.CallOption) (*StatusPriceResponse, error) {
	out := new(StatusPriceResponse)
	err := c.cc.Invoke(ctx, "/price.PriceService/DeletePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) UpdatePrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*StatusPriceResponse, error) {
	out := new(StatusPriceResponse)
	err := c.cc.Invoke(ctx, "/price.PriceService/UpdatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetPrice(ctx context.Context, in *IdPriceRequest, opts ...grpc.CallOption) (*PriceResponse, error) {
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, "/price.PriceService/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
type PriceServiceServer interface {
	CreatePrice(context.Context, *PriceRequest) (*IdPriceResponse, error)
	DeletePrice(context.Context, *IdPriceRequest) (*StatusPriceResponse, error)
	UpdatePrice(context.Context, *PriceRequest) (*StatusPriceResponse, error)
	GetPrice(context.Context, *IdPriceRequest) (*PriceResponse, error)
}

// UnimplementedPriceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPriceServiceServer struct {
}

func (*UnimplementedPriceServiceServer) CreatePrice(context.Context, *PriceRequest) (*IdPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrice not implemented")
}
func (*UnimplementedPriceServiceServer) DeletePrice(context.Context, *IdPriceRequest) (*StatusPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePrice not implemented")
}
func (*UnimplementedPriceServiceServer) UpdatePrice(context.Context, *PriceRequest) (*StatusPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrice not implemented")
}
func (*UnimplementedPriceServiceServer) GetPrice(context.Context, *IdPriceRequest) (*PriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}

func RegisterPriceServiceServer(s *grpc.Server, srv PriceServiceServer) {
	s.RegisterService(&_PriceService_serviceDesc, srv)
}

func _PriceService_CreatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).CreatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price.PriceService/CreatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).CreatePrice(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_DeletePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).DeletePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price.PriceService/DeletePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).DeletePrice(ctx, req.(*IdPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_UpdatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).UpdatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price.PriceService/UpdatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).UpdatePrice(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price.PriceService/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetPrice(ctx, req.(*IdPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PriceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "price.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrice",
			Handler:    _PriceService_CreatePrice_Handler,
		},
		{
			MethodName: "DeletePrice",
			Handler:    _PriceService_DeletePrice_Handler,
		},
		{
			MethodName: "UpdatePrice",
			Handler:    _PriceService_UpdatePrice_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _PriceService_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price.proto",
}