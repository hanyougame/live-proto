// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.29.3
// source: manage/v1/manage.proto

package v1

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

type ManageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManageReq) Reset() {
	*x = ManageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_v1_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageReq) ProtoMessage() {}

func (x *ManageReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_v1_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageReq.ProtoReflect.Descriptor instead.
func (*ManageReq) Descriptor() ([]byte, []int) {
	return file_manage_v1_manage_proto_rawDescGZIP(), []int{0}
}

type ManageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManageReply) Reset() {
	*x = ManageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_v1_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageReply) ProtoMessage() {}

func (x *ManageReply) ProtoReflect() protoreflect.Message {
	mi := &file_manage_v1_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageReply.ProtoReflect.Descriptor instead.
func (*ManageReply) Descriptor() ([]byte, []int) {
	return file_manage_v1_manage_proto_rawDescGZIP(), []int{1}
}

type CurrencyDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyId        int64  `protobuf:"varint,1,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`                     // 自增主键，唯一标识每条记录
	Rank              int64  `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`                                                   // 排名，表示货币在列表中的排序位置
	CurrencyName      string `protobuf:"bytes,3,opt,name=currency_name,json=currencyName,proto3" json:"currency_name,omitempty"`                // 货币名称，例如“印度卢比”
	CountryName       string `protobuf:"bytes,4,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`                   // 所属国家，例如“印度”，允许为空（如“波场”）
	CurrencyCode      string `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`                // 货币代码，例如“INR”
	CurrencyFlagUrl   string `protobuf:"bytes,6,opt,name=currency_flag_url,json=currencyFlagUrl,proto3" json:"currency_flag_url,omitempty"`     // 货币国旗图标的 URL，保存图片的路径或链接
	ThousandSeparator string `protobuf:"bytes,7,opt,name=thousand_separator,json=thousandSeparator,proto3" json:"thousand_separator,omitempty"` // 千分位符号，例如“,”，表示格式化显示
	CurrencySymbol    string `protobuf:"bytes,8,opt,name=currency_symbol,json=currencySymbol,proto3" json:"currency_symbol,omitempty"`          // 货币符号，例如“₹”
	ExchangeRate      string `protobuf:"bytes,9,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`                // 货币兑换比例，默认值为 1:1
	CurrencyType      int64  `protobuf:"varint,10,opt,name=currency_type,json=currencyType,proto3" json:"currency_type,omitempty"`              // 货币类型，例如“法定货币”或“数字货币”
	IsEnabled         int64  `protobuf:"varint,11,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`                       // 是否启用，默认启用
	IsVisible         int64  `protobuf:"varint,12,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`                       // 大厅显示开关，表示是否在用户界面显示
}

func (x *CurrencyDetailReply) Reset() {
	*x = CurrencyDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_v1_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyDetailReply) ProtoMessage() {}

func (x *CurrencyDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_manage_v1_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyDetailReply.ProtoReflect.Descriptor instead.
func (*CurrencyDetailReply) Descriptor() ([]byte, []int) {
	return file_manage_v1_manage_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyDetailReply) GetCurrencyId() int64 {
	if x != nil {
		return x.CurrencyId
	}
	return 0
}

func (x *CurrencyDetailReply) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *CurrencyDetailReply) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

func (x *CurrencyDetailReply) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *CurrencyDetailReply) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *CurrencyDetailReply) GetCurrencyFlagUrl() string {
	if x != nil {
		return x.CurrencyFlagUrl
	}
	return ""
}

func (x *CurrencyDetailReply) GetThousandSeparator() string {
	if x != nil {
		return x.ThousandSeparator
	}
	return ""
}

func (x *CurrencyDetailReply) GetCurrencySymbol() string {
	if x != nil {
		return x.CurrencySymbol
	}
	return ""
}

func (x *CurrencyDetailReply) GetExchangeRate() string {
	if x != nil {
		return x.ExchangeRate
	}
	return ""
}

func (x *CurrencyDetailReply) GetCurrencyType() int64 {
	if x != nil {
		return x.CurrencyType
	}
	return 0
}

func (x *CurrencyDetailReply) GetIsEnabled() int64 {
	if x != nil {
		return x.IsEnabled
	}
	return 0
}

func (x *CurrencyDetailReply) GetIsVisible() int64 {
	if x != nil {
		return x.IsVisible
	}
	return 0
}

type CurrencyDetailReplyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*CurrencyDetailReply `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CurrencyDetailReplyList) Reset() {
	*x = CurrencyDetailReplyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_v1_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyDetailReplyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyDetailReplyList) ProtoMessage() {}

func (x *CurrencyDetailReplyList) ProtoReflect() protoreflect.Message {
	mi := &file_manage_v1_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyDetailReplyList.ProtoReflect.Descriptor instead.
func (*CurrencyDetailReplyList) Descriptor() ([]byte, []int) {
	return file_manage_v1_manage_proto_rawDescGZIP(), []int{3}
}

func (x *CurrencyDetailReplyList) GetList() []*CurrencyDetailReply {
	if x != nil {
		return x.List
	}
	return nil
}

type GetCurrInfoByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 自增主键，唯一标识每条记录
}

func (x *GetCurrInfoByIDReq) Reset() {
	*x = GetCurrInfoByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_v1_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrInfoByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrInfoByIDReq) ProtoMessage() {}

func (x *GetCurrInfoByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_v1_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrInfoByIDReq.ProtoReflect.Descriptor instead.
func (*GetCurrInfoByIDReq) Descriptor() ([]byte, []int) {
	return file_manage_v1_manage_proto_rawDescGZIP(), []int{4}
}

func (x *GetCurrInfoByIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_manage_v1_manage_proto protoreflect.FileDescriptor

var file_manage_v1_manage_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x0b, 0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x22, 0x0d, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0xc3, 0x03, 0x0a, 0x13, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x6c,
	0x61, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x61, 0x6e,
	0x64, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb8, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x76, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x54,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x22, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manage_v1_manage_proto_rawDescOnce sync.Once
	file_manage_v1_manage_proto_rawDescData = file_manage_v1_manage_proto_rawDesc
)

func file_manage_v1_manage_proto_rawDescGZIP() []byte {
	file_manage_v1_manage_proto_rawDescOnce.Do(func() {
		file_manage_v1_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_manage_v1_manage_proto_rawDescData)
	})
	return file_manage_v1_manage_proto_rawDescData
}

var file_manage_v1_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_manage_v1_manage_proto_goTypes = []interface{}{
	(*ManageReq)(nil),               // 0: manage.v1.ManageReq
	(*ManageReply)(nil),             // 1: manage.v1.ManageReply
	(*CurrencyDetailReply)(nil),     // 2: manage.v1.CurrencyDetailReply
	(*CurrencyDetailReplyList)(nil), // 3: manage.v1.CurrencyDetailReplyList
	(*GetCurrInfoByIDReq)(nil),      // 4: manage.v1.GetCurrInfoByIDReq
}
var file_manage_v1_manage_proto_depIdxs = []int32{
	2, // 0: manage.v1.CurrencyDetailReplyList.list:type_name -> manage.v1.CurrencyDetailReply
	0, // 1: manage.v1.LiveManageRpcService.GetLegalTenderInfo:input_type -> manage.v1.ManageReq
	4, // 2: manage.v1.LiveManageRpcService.GetCurrInfoById:input_type -> manage.v1.GetCurrInfoByIDReq
	3, // 3: manage.v1.LiveManageRpcService.GetLegalTenderInfo:output_type -> manage.v1.CurrencyDetailReplyList
	2, // 4: manage.v1.LiveManageRpcService.GetCurrInfoById:output_type -> manage.v1.CurrencyDetailReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manage_v1_manage_proto_init() }
func file_manage_v1_manage_proto_init() {
	if File_manage_v1_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manage_v1_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageReq); i {
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
		file_manage_v1_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageReply); i {
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
		file_manage_v1_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyDetailReply); i {
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
		file_manage_v1_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyDetailReplyList); i {
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
		file_manage_v1_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrInfoByIDReq); i {
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
			RawDescriptor: file_manage_v1_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manage_v1_manage_proto_goTypes,
		DependencyIndexes: file_manage_v1_manage_proto_depIdxs,
		MessageInfos:      file_manage_v1_manage_proto_msgTypes,
	}.Build()
	File_manage_v1_manage_proto = out.File
	file_manage_v1_manage_proto_rawDesc = nil
	file_manage_v1_manage_proto_goTypes = nil
	file_manage_v1_manage_proto_depIdxs = nil
}
