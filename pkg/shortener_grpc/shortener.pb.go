// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/shortener.proto

package shortener_grpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type AddURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongURL string `protobuf:"bytes,1,opt,name=longURL,proto3" json:"longURL,omitempty"`
}

func (x *AddURLRequest) Reset() {
	*x = AddURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddURLRequest) ProtoMessage() {}

func (x *AddURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddURLRequest.ProtoReflect.Descriptor instead.
func (*AddURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *AddURLRequest) GetLongURL() string {
	if x != nil {
		return x.LongURL
	}
	return ""
}

type AddURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddURLResponse) Reset() {
	*x = AddURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddURLResponse) ProtoMessage() {}

func (x *AddURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddURLResponse.ProtoReflect.Descriptor instead.
func (*AddURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *AddURLResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

func (x *AddURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetURLRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongURL string `protobuf:"bytes,1,opt,name=longURL,proto3" json:"longURL,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetURLResponse) GetLongURL() string {
	if x != nil {
		return x.LongURL
	}
	return ""
}

func (x *GetURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type BatchURLItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	LongURL       string `protobuf:"bytes,2,opt,name=longURL,proto3" json:"longURL,omitempty"`
}

func (x *BatchURLItem) Reset() {
	*x = BatchURLItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchURLItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchURLItem) ProtoMessage() {}

func (x *BatchURLItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchURLItem.ProtoReflect.Descriptor instead.
func (*BatchURLItem) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *BatchURLItem) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *BatchURLItem) GetLongURL() string {
	if x != nil {
		return x.LongURL
	}
	return ""
}

type AddBatchURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchURL []*BatchURLItem `protobuf:"bytes,1,rep,name=BatchURL,proto3" json:"BatchURL,omitempty"`
}

func (x *AddBatchURLRequest) Reset() {
	*x = AddBatchURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBatchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBatchURLRequest) ProtoMessage() {}

func (x *AddBatchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBatchURLRequest.ProtoReflect.Descriptor instead.
func (*AddBatchURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *AddBatchURLRequest) GetBatchURL() []*BatchURLItem {
	if x != nil {
		return x.BatchURL
	}
	return nil
}

type AddBatchURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchURL []*BatchURLItem `protobuf:"bytes,1,rep,name=BatchURL,proto3" json:"BatchURL,omitempty"`
	Error    string          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddBatchURLResponse) Reset() {
	*x = AddBatchURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBatchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBatchURLResponse) ProtoMessage() {}

func (x *AddBatchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBatchURLResponse.ProtoReflect.Descriptor instead.
func (*AddBatchURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *AddBatchURLResponse) GetBatchURL() []*BatchURLItem {
	if x != nil {
		return x.BatchURL
	}
	return nil
}

func (x *AddBatchURLResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type URLComplex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL    string `protobuf:"bytes,1,opt,name=ShortURL,proto3" json:"ShortURL,omitempty"`
	OriginalURL string `protobuf:"bytes,2,opt,name=OriginalURL,proto3" json:"OriginalURL,omitempty"`
}

func (x *URLComplex) Reset() {
	*x = URLComplex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLComplex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLComplex) ProtoMessage() {}

func (x *URLComplex) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLComplex.ProtoReflect.Descriptor instead.
func (*URLComplex) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *URLComplex) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

func (x *URLComplex) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

type GetURLsByUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLs  []*URLComplex `protobuf:"bytes,1,rep,name=URLs,proto3" json:"URLs,omitempty"`
	Error string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetURLsByUserResponse) Reset() {
	*x = GetURLsByUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLsByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLsByUserResponse) ProtoMessage() {}

func (x *GetURLsByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLsByUserResponse.ProtoReflect.Descriptor instead.
func (*GetURLsByUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *GetURLsByUserResponse) GetURLs() []*URLComplex {
	if x != nil {
		return x.URLs
	}
	return nil
}

func (x *GetURLsByUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL []string `protobuf:"bytes,1,rep,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *DeleteURLsRequest) Reset() {
	*x = DeleteURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteURLsRequest) ProtoMessage() {}

func (x *DeleteURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteURLsRequest.ProtoReflect.Descriptor instead.
func (*DeleteURLsRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteURLsRequest) GetShortURL() []string {
	if x != nil {
		return x.ShortURL
	}
	return nil
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLS  int32 `protobuf:"varint,1,opt,name=URLS,proto3" json:"URLS,omitempty"`
	Users int32 `protobuf:"varint,2,opt,name=Users,proto3" json:"Users,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *GetStatsResponse) GetURLS() int32 {
	if x != nil {
		return x.URLS
	}
	return 0
}

func (x *GetStatsResponse) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

var File_proto_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x29, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x22, 0x42, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x40, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4e, 0x0a,
	0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x52, 0x4c, 0x22, 0x49, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x22, 0x60, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x0a, 0x55, 0x52,
	0x4c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x4c, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x2f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x53, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x32,
	0xaf, 0x03, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x3d, 0x0a,
	0x06, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52,
	0x4c, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_proto_rawDescData = file_proto_shortener_proto_rawDesc
)

func file_proto_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_proto_rawDescData)
	})
	return file_proto_shortener_proto_rawDescData
}

var file_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_shortener_proto_goTypes = []interface{}{
	(*AddURLRequest)(nil),         // 0: shortener.AddURLRequest
	(*AddURLResponse)(nil),        // 1: shortener.AddURLResponse
	(*GetURLRequest)(nil),         // 2: shortener.GetURLRequest
	(*GetURLResponse)(nil),        // 3: shortener.GetURLResponse
	(*BatchURLItem)(nil),          // 4: shortener.BatchURLItem
	(*AddBatchURLRequest)(nil),    // 5: shortener.AddBatchURLRequest
	(*AddBatchURLResponse)(nil),   // 6: shortener.AddBatchURLResponse
	(*URLComplex)(nil),            // 7: shortener.URLComplex
	(*GetURLsByUserResponse)(nil), // 8: shortener.GetURLsByUserResponse
	(*DeleteURLsRequest)(nil),     // 9: shortener.DeleteURLsRequest
	(*GetStatsResponse)(nil),      // 10: shortener.GetStatsResponse
	(*empty.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_proto_shortener_proto_depIdxs = []int32{
	4,  // 0: shortener.AddBatchURLRequest.BatchURL:type_name -> shortener.BatchURLItem
	4,  // 1: shortener.AddBatchURLResponse.BatchURL:type_name -> shortener.BatchURLItem
	7,  // 2: shortener.GetURLsByUserResponse.URLs:type_name -> shortener.URLComplex
	0,  // 3: shortener.Shortener.AddURL:input_type -> shortener.AddURLRequest
	2,  // 4: shortener.Shortener.GetURLByShort:input_type -> shortener.GetURLRequest
	5,  // 5: shortener.Shortener.AddBatchURLs:input_type -> shortener.AddBatchURLRequest
	11, // 6: shortener.Shortener.GetURLsByUser:input_type -> google.protobuf.Empty
	9,  // 7: shortener.Shortener.DeleteURLs:input_type -> shortener.DeleteURLsRequest
	11, // 8: shortener.Shortener.GetStats:input_type -> google.protobuf.Empty
	1,  // 9: shortener.Shortener.AddURL:output_type -> shortener.AddURLResponse
	3,  // 10: shortener.Shortener.GetURLByShort:output_type -> shortener.GetURLResponse
	6,  // 11: shortener.Shortener.AddBatchURLs:output_type -> shortener.AddBatchURLResponse
	8,  // 12: shortener.Shortener.GetURLsByUser:output_type -> shortener.GetURLsByUserResponse
	11, // 13: shortener.Shortener.DeleteURLs:output_type -> google.protobuf.Empty
	10, // 14: shortener.Shortener.GetStats:output_type -> shortener.GetStatsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_shortener_proto_init() }
func file_proto_shortener_proto_init() {
	if File_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchURLItem); i {
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
		file_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBatchURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBatchURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLComplex); i {
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
		file_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLsByUserResponse); i {
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
		file_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteURLsRequest); i {
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
		file_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
			RawDescriptor: file_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_proto_depIdxs,
		MessageInfos:      file_proto_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_proto = out.File
	file_proto_shortener_proto_rawDesc = nil
	file_proto_shortener_proto_goTypes = nil
	file_proto_shortener_proto_depIdxs = nil
}
