// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: mqtt/v1/mqtt.proto

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

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_mqtt_v1_mqtt_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushToUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushToUserRequest) Reset() {
	*x = PushToUserRequest{}
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushToUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushToUserRequest) ProtoMessage() {}

func (x *PushToUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushToUserRequest.ProtoReflect.Descriptor instead.
func (*PushToUserRequest) Descriptor() ([]byte, []int) {
	return file_mqtt_v1_mqtt_proto_rawDescGZIP(), []int{1}
}

func (x *PushToUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PushToUserRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushToUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Message []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushToUsersRequest) Reset() {
	*x = PushToUsersRequest{}
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushToUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushToUsersRequest) ProtoMessage() {}

func (x *PushToUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushToUsersRequest.ProtoReflect.Descriptor instead.
func (*PushToUsersRequest) Descriptor() ([]byte, []int) {
	return file_mqtt_v1_mqtt_proto_rawDescGZIP(), []int{2}
}

func (x *PushToUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *PushToUsersRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_mqtt_v1_mqtt_proto_rawDescGZIP(), []int{3}
}

func (x *PushResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PushResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

// EnterGameDelReply 进入游戏转账信息/退出游戏转账信息
type EnterQuitGameDelWdlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId         int64 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Balance        int64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	GamePlatformId int64 `protobuf:"varint,3,opt,name=game_platform_id,json=gamePlatformId,proto3" json:"game_platform_id,omitempty"`
}

func (x *EnterQuitGameDelWdlReply) Reset() {
	*x = EnterQuitGameDelWdlReply{}
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterQuitGameDelWdlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterQuitGameDelWdlReply) ProtoMessage() {}

func (x *EnterQuitGameDelWdlReply) ProtoReflect() protoreflect.Message {
	mi := &file_mqtt_v1_mqtt_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterQuitGameDelWdlReply.ProtoReflect.Descriptor instead.
func (*EnterQuitGameDelWdlReply) Descriptor() ([]byte, []int) {
	return file_mqtt_v1_mqtt_proto_rawDescGZIP(), []int{4}
}

func (x *EnterQuitGameDelWdlReply) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *EnterQuitGameDelWdlReply) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *EnterQuitGameDelWdlReply) GetGamePlatformId() int64 {
	if x != nil {
		return x.GamePlatformId
	}
	return 0
}

var File_mqtt_v1_mqtt_proto protoreflect.FileDescriptor

var file_mqtt_v1_mqtt_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x71, 0x74, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a,
	0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x49,
	0x0a, 0x12, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x77, 0x0a, 0x18, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x51,
	0x75, 0x69, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x57, 0x64, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x67, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x32,
	0xd1, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x6d,
	0x71, 0x74, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6d, 0x71, 0x74, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e,
	0x6d, 0x71, 0x74, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x71, 0x74,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mqtt_v1_mqtt_proto_rawDescOnce sync.Once
	file_mqtt_v1_mqtt_proto_rawDescData = file_mqtt_v1_mqtt_proto_rawDesc
)

func file_mqtt_v1_mqtt_proto_rawDescGZIP() []byte {
	file_mqtt_v1_mqtt_proto_rawDescOnce.Do(func() {
		file_mqtt_v1_mqtt_proto_rawDescData = protoimpl.X.CompressGZIP(file_mqtt_v1_mqtt_proto_rawDescData)
	})
	return file_mqtt_v1_mqtt_proto_rawDescData
}

var file_mqtt_v1_mqtt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mqtt_v1_mqtt_proto_goTypes = []any{
	(*PushRequest)(nil),              // 0: mqtt.v1.PushRequest
	(*PushToUserRequest)(nil),        // 1: mqtt.v1.PushToUserRequest
	(*PushToUsersRequest)(nil),       // 2: mqtt.v1.PushToUsersRequest
	(*PushResponse)(nil),             // 3: mqtt.v1.PushResponse
	(*EnterQuitGameDelWdlReply)(nil), // 4: mqtt.v1.EnterQuitGameDelWdlReply
}
var file_mqtt_v1_mqtt_proto_depIdxs = []int32{
	0, // 0: mqtt.v1.PushService.PushToAll:input_type -> mqtt.v1.PushRequest
	1, // 1: mqtt.v1.PushService.PushToUser:input_type -> mqtt.v1.PushToUserRequest
	2, // 2: mqtt.v1.PushService.PushToUsers:input_type -> mqtt.v1.PushToUsersRequest
	3, // 3: mqtt.v1.PushService.PushToAll:output_type -> mqtt.v1.PushResponse
	3, // 4: mqtt.v1.PushService.PushToUser:output_type -> mqtt.v1.PushResponse
	3, // 5: mqtt.v1.PushService.PushToUsers:output_type -> mqtt.v1.PushResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mqtt_v1_mqtt_proto_init() }
func file_mqtt_v1_mqtt_proto_init() {
	if File_mqtt_v1_mqtt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mqtt_v1_mqtt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mqtt_v1_mqtt_proto_goTypes,
		DependencyIndexes: file_mqtt_v1_mqtt_proto_depIdxs,
		MessageInfos:      file_mqtt_v1_mqtt_proto_msgTypes,
	}.Build()
	File_mqtt_v1_mqtt_proto = out.File
	file_mqtt_v1_mqtt_proto_rawDesc = nil
	file_mqtt_v1_mqtt_proto_goTypes = nil
	file_mqtt_v1_mqtt_proto_depIdxs = nil
}
