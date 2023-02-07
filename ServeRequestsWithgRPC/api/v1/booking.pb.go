// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/v1/booking.proto

package booking_v1

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

type GetByUUIDBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *GetByUUIDBookingRequest) Reset() {
	*x = GetByUUIDBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUUIDBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUUIDBookingRequest) ProtoMessage() {}

func (x *GetByUUIDBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUUIDBookingRequest.ProtoReflect.Descriptor instead.
func (*GetByUUIDBookingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *GetByUUIDBookingRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type GetByIDBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetByIDBookingRequest) Reset() {
	*x = GetByIDBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDBookingRequest) ProtoMessage() {}

func (x *GetByIDBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDBookingRequest.ProtoReflect.Descriptor instead.
func (*GetByIDBookingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{1}
}

func (x *GetByIDBookingRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *GetBookingResponse) Reset() {
	*x = GetBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingResponse) ProtoMessage() {}

func (x *GetBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingResponse.ProtoReflect.Descriptor instead.
func (*GetBookingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookingResponse) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type CreateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingRequest.ProtoReflect.Descriptor instead.
func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookingRequest) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type CreateBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *CreateBookingResponse) Reset() {
	*x = CreateBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingResponse) ProtoMessage() {}

func (x *CreateBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookingResponse.ProtoReflect.Descriptor instead.
func (*CreateBookingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBookingResponse) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type UpdateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *UpdateBookingRequest) Reset() {
	*x = UpdateBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingRequest) ProtoMessage() {}

func (x *UpdateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookingRequest) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type UpdateBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking *Booking `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *UpdateBookingResponse) Reset() {
	*x = UpdateBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingResponse) ProtoMessage() {}

func (x *UpdateBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookingResponse) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UUID      string                 `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	FullName  string                 `protobuf:"bytes,4,opt,name=FullName,proto3" json:"FullName,omitempty"`
	StartDate string                 `protobuf:"bytes,5,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate   string                 `protobuf:"bytes,6,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	Active    bool                   `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_api_v1_booking_proto_rawDescGZIP(), []int{7}
}

func (x *Booking) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Booking) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Booking) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Booking) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Booking) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Booking) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Booking) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Booking) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Booking) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_api_v1_booking_proto protoreflect.FileDescriptor

var file_api_v1_booking_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x22, 0x45, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22,
	0x45, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x46, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0xa3,
	0x02, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x32, 0xaf, 0x04, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x55,
	0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x60, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x67, 0x6f, 0x72, 0x2d, 0x62, 0x61, 0x69, 0x62, 0x6f, 0x72,
	0x6f, 0x64, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_booking_proto_rawDescOnce sync.Once
	file_api_v1_booking_proto_rawDescData = file_api_v1_booking_proto_rawDesc
)

func file_api_v1_booking_proto_rawDescGZIP() []byte {
	file_api_v1_booking_proto_rawDescOnce.Do(func() {
		file_api_v1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_booking_proto_rawDescData)
	})
	return file_api_v1_booking_proto_rawDescData
}

var file_api_v1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_booking_proto_goTypes = []interface{}{
	(*GetByUUIDBookingRequest)(nil), // 0: booking.v1.GetByUUIDBookingRequest
	(*GetByIDBookingRequest)(nil),   // 1: booking.v1.GetByIDBookingRequest
	(*GetBookingResponse)(nil),      // 2: booking.v1.GetBookingResponse
	(*CreateBookingRequest)(nil),    // 3: booking.v1.CreateBookingRequest
	(*CreateBookingResponse)(nil),   // 4: booking.v1.CreateBookingResponse
	(*UpdateBookingRequest)(nil),    // 5: booking.v1.UpdateBookingRequest
	(*UpdateBookingResponse)(nil),   // 6: booking.v1.UpdateBookingResponse
	(*Booking)(nil),                 // 7: booking.v1.Booking
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_api_v1_booking_proto_depIdxs = []int32{
	7,  // 0: booking.v1.GetBookingResponse.booking:type_name -> booking.v1.Booking
	7,  // 1: booking.v1.CreateBookingRequest.booking:type_name -> booking.v1.Booking
	7,  // 2: booking.v1.CreateBookingResponse.booking:type_name -> booking.v1.Booking
	7,  // 3: booking.v1.UpdateBookingRequest.booking:type_name -> booking.v1.Booking
	7,  // 4: booking.v1.UpdateBookingResponse.booking:type_name -> booking.v1.Booking
	8,  // 5: booking.v1.Booking.CreatedAt:type_name -> google.protobuf.Timestamp
	8,  // 6: booking.v1.Booking.UpdatedAt:type_name -> google.protobuf.Timestamp
	0,  // 7: booking.v1.BookingService.GetBookingByUUID:input_type -> booking.v1.GetByUUIDBookingRequest
	1,  // 8: booking.v1.BookingService.GetBookingByID:input_type -> booking.v1.GetByIDBookingRequest
	3,  // 9: booking.v1.BookingService.CreateBooking:input_type -> booking.v1.CreateBookingRequest
	5,  // 10: booking.v1.BookingService.UpdateBooking:input_type -> booking.v1.UpdateBookingRequest
	1,  // 11: booking.v1.BookingService.GetBookingStream:input_type -> booking.v1.GetByIDBookingRequest
	3,  // 12: booking.v1.BookingService.CreateBookingStream:input_type -> booking.v1.CreateBookingRequest
	2,  // 13: booking.v1.BookingService.GetBookingByUUID:output_type -> booking.v1.GetBookingResponse
	2,  // 14: booking.v1.BookingService.GetBookingByID:output_type -> booking.v1.GetBookingResponse
	4,  // 15: booking.v1.BookingService.CreateBooking:output_type -> booking.v1.CreateBookingResponse
	6,  // 16: booking.v1.BookingService.UpdateBooking:output_type -> booking.v1.UpdateBookingResponse
	2,  // 17: booking.v1.BookingService.GetBookingStream:output_type -> booking.v1.GetBookingResponse
	4,  // 18: booking.v1.BookingService.CreateBookingStream:output_type -> booking.v1.CreateBookingResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_booking_proto_init() }
func file_api_v1_booking_proto_init() {
	if File_api_v1_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUUIDBookingRequest); i {
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
		file_api_v1_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDBookingRequest); i {
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
		file_api_v1_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingResponse); i {
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
		file_api_v1_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookingRequest); i {
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
		file_api_v1_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookingResponse); i {
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
		file_api_v1_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookingRequest); i {
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
		file_api_v1_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookingResponse); i {
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
		file_api_v1_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
			RawDescriptor: file_api_v1_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_booking_proto_goTypes,
		DependencyIndexes: file_api_v1_booking_proto_depIdxs,
		MessageInfos:      file_api_v1_booking_proto_msgTypes,
	}.Build()
	File_api_v1_booking_proto = out.File
	file_api_v1_booking_proto_rawDesc = nil
	file_api_v1_booking_proto_goTypes = nil
	file_api_v1_booking_proto_depIdxs = nil
}
