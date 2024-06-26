// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: reservista/mailer/mailer.proto

package proto_mailer

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

type EmailInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailInput) Reset() {
	*x = EmailInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailInput) ProtoMessage() {}

func (x *EmailInput) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailInput.ProtoReflect.Descriptor instead.
func (*EmailInput) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *EmailInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ContentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ContentInput) Reset() {
	*x = ContentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentInput) ProtoMessage() {}

func (x *ContentInput) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentInput.ProtoReflect.Descriptor instead.
func (*ContentInput) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *ContentInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContentInput) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type QRInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ReservationID string `protobuf:"bytes,2,opt,name=reservationID,proto3" json:"reservationID,omitempty"`
	QRUrlBase     string `protobuf:"bytes,3,opt,name=QRUrlBase,proto3" json:"QRUrlBase,omitempty"`
}

func (x *QRInput) Reset() {
	*x = QRInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRInput) ProtoMessage() {}

func (x *QRInput) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRInput.ProtoReflect.Descriptor instead.
func (*QRInput) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{3}
}

func (x *QRInput) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *QRInput) GetReservationID() string {
	if x != nil {
		return x.ReservationID
	}
	return ""
}

func (x *QRInput) GetQRUrlBase() string {
	if x != nil {
		return x.QRUrlBase
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RestaurantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address         string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Contact         string `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	TableNumber     int32  `protobuf:"varint,4,opt,name=tableNumber,proto3" json:"tableNumber,omitempty"`
	ReservationTime string `protobuf:"bytes,5,opt,name=reservationTime,proto3" json:"reservationTime,omitempty"`
}

func (x *RestaurantInfo) Reset() {
	*x = RestaurantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_mailer_mailer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantInfo) ProtoMessage() {}

func (x *RestaurantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_mailer_mailer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantInfo.ProtoReflect.Descriptor instead.
func (*RestaurantInfo) Descriptor() ([]byte, []int) {
	return file_reservista_mailer_mailer_proto_rawDescGZIP(), []int{5}
}

func (x *RestaurantInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestaurantInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RestaurantInfo) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *RestaurantInfo) GetTableNumber() int32 {
	if x != nil {
		return x.TableNumber
	}
	return 0
}

func (x *RestaurantInfo) GetReservationTime() string {
	if x != nil {
		return x.ReservationTime
	}
	return ""
}

var File_reservista_mailer_mailer_proto protoreflect.FileDescriptor

var file_reservista_mailer_mailer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0a, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3e, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x65, 0x0a, 0x07, 0x51, 0x52, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x51, 0x52, 0x55, 0x72, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x51, 0x52, 0x55, 0x72, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x22, 0x64, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xe9, 0x02, 0x0a, 0x06, 0x4d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x51, 0x52, 0x12, 0x0f, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x51, 0x52, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0x12, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x61, 0x69, 0x64, 0x6f, 0x73, 0x74,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_reservista_mailer_mailer_proto_rawDescOnce sync.Once
	file_reservista_mailer_mailer_proto_rawDescData = file_reservista_mailer_mailer_proto_rawDesc
)

func file_reservista_mailer_mailer_proto_rawDescGZIP() []byte {
	file_reservista_mailer_mailer_proto_rawDescOnce.Do(func() {
		file_reservista_mailer_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservista_mailer_mailer_proto_rawDescData)
	})
	return file_reservista_mailer_mailer_proto_rawDescData
}

var file_reservista_mailer_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reservista_mailer_mailer_proto_goTypes = []interface{}{
	(*EmailInput)(nil),     // 0: mailer.EmailInput
	(*ContentInput)(nil),   // 1: mailer.ContentInput
	(*StatusResponse)(nil), // 2: mailer.StatusResponse
	(*QRInput)(nil),        // 3: mailer.QRInput
	(*UserInfo)(nil),       // 4: mailer.UserInfo
	(*RestaurantInfo)(nil), // 5: mailer.RestaurantInfo
}
var file_reservista_mailer_mailer_proto_depIdxs = []int32{
	1, // 0: mailer.Mailer.SendWelcome:input_type -> mailer.ContentInput
	3, // 1: mailer.Mailer.SendQR:input_type -> mailer.QRInput
	1, // 2: mailer.Mailer.SendAuthCode:input_type -> mailer.ContentInput
	0, // 3: mailer.Mailer.SendReminder:input_type -> mailer.EmailInput
	0, // 4: mailer.Mailer.SendAdvert:input_type -> mailer.EmailInput
	0, // 5: mailer.Mailer.SendResetCode:input_type -> mailer.EmailInput
	2, // 6: mailer.Mailer.SendWelcome:output_type -> mailer.StatusResponse
	2, // 7: mailer.Mailer.SendQR:output_type -> mailer.StatusResponse
	2, // 8: mailer.Mailer.SendAuthCode:output_type -> mailer.StatusResponse
	2, // 9: mailer.Mailer.SendReminder:output_type -> mailer.StatusResponse
	2, // 10: mailer.Mailer.SendAdvert:output_type -> mailer.StatusResponse
	2, // 11: mailer.Mailer.SendResetCode:output_type -> mailer.StatusResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reservista_mailer_mailer_proto_init() }
func file_reservista_mailer_mailer_proto_init() {
	if File_reservista_mailer_mailer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservista_mailer_mailer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailInput); i {
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
		file_reservista_mailer_mailer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentInput); i {
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
		file_reservista_mailer_mailer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_reservista_mailer_mailer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRInput); i {
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
		file_reservista_mailer_mailer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_reservista_mailer_mailer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantInfo); i {
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
			RawDescriptor: file_reservista_mailer_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservista_mailer_mailer_proto_goTypes,
		DependencyIndexes: file_reservista_mailer_mailer_proto_depIdxs,
		MessageInfos:      file_reservista_mailer_mailer_proto_msgTypes,
	}.Build()
	File_reservista_mailer_mailer_proto = out.File
	file_reservista_mailer_mailer_proto_rawDesc = nil
	file_reservista_mailer_mailer_proto_goTypes = nil
	file_reservista_mailer_mailer_proto_depIdxs = nil
}
