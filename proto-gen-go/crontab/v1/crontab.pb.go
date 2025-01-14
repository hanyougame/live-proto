// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.19.4
// source: crontab/v1/crontab.proto

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

type EmptyRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRes) Reset() {
	*x = EmptyRes{}
	mi := &file_crontab_v1_crontab_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRes) ProtoMessage() {}

func (x *EmptyRes) ProtoReflect() protoreflect.Message {
	mi := &file_crontab_v1_crontab_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRes.ProtoReflect.Descriptor instead.
func (*EmptyRes) Descriptor() ([]byte, []int) {
	return file_crontab_v1_crontab_proto_rawDescGZIP(), []int{0}
}

type AddReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`     // 任务唯一标识
	Spec          string                 `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`       // 执行时间
	Type          int32                  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`      // 执行类型 0 正常执行 1 上次没执行完跳过 2 上次没执行完延后执行 3 只执行一次
	Handler       string                 `protobuf:"bytes,4,opt,name=handler,proto3" json:"handler,omitempty"` // 使用的方法
	Params        string                 `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	mi := &file_crontab_v1_crontab_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_crontab_v1_crontab_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_crontab_v1_crontab_proto_rawDescGZIP(), []int{1}
}

func (x *AddReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AddReq) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *AddReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AddReq) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *AddReq) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

type AddRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CrontabId     int64                  `protobuf:"varint,1,opt,name=crontab_id,json=crontabId,proto3" json:"crontab_id,omitempty"` // 定时任务id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRes) Reset() {
	*x = AddRes{}
	mi := &file_crontab_v1_crontab_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRes) ProtoMessage() {}

func (x *AddRes) ProtoReflect() protoreflect.Message {
	mi := &file_crontab_v1_crontab_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRes.ProtoReflect.Descriptor instead.
func (*AddRes) Descriptor() ([]byte, []int) {
	return file_crontab_v1_crontab_proto_rawDescGZIP(), []int{2}
}

func (x *AddRes) GetCrontabId() int64 {
	if x != nil {
		return x.CrontabId
	}
	return 0
}

type DelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`                           // 定时任务唯一标识
	CrontabId     int64                  `protobuf:"varint,2,opt,name=crontab_id,json=crontabId,proto3" json:"crontab_id,omitempty"` // 定时任务id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelReq) Reset() {
	*x = DelReq{}
	mi := &file_crontab_v1_crontab_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReq) ProtoMessage() {}

func (x *DelReq) ProtoReflect() protoreflect.Message {
	mi := &file_crontab_v1_crontab_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReq.ProtoReflect.Descriptor instead.
func (*DelReq) Descriptor() ([]byte, []int) {
	return file_crontab_v1_crontab_proto_rawDescGZIP(), []int{3}
}

func (x *DelReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DelReq) GetCrontabId() int64 {
	if x != nil {
		return x.CrontabId
	}
	return 0
}

var File_crontab_v1_crontab_proto protoreflect.FileDescriptor

var file_crontab_v1_crontab_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f,
	0x6e, 0x74, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22,
	0x78, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x06, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x49,
	0x64, 0x32, 0x6b, 0x0a, 0x15, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62,
	0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_crontab_v1_crontab_proto_rawDescOnce sync.Once
	file_crontab_v1_crontab_proto_rawDescData = file_crontab_v1_crontab_proto_rawDesc
)

func file_crontab_v1_crontab_proto_rawDescGZIP() []byte {
	file_crontab_v1_crontab_proto_rawDescOnce.Do(func() {
		file_crontab_v1_crontab_proto_rawDescData = protoimpl.X.CompressGZIP(file_crontab_v1_crontab_proto_rawDescData)
	})
	return file_crontab_v1_crontab_proto_rawDescData
}

var file_crontab_v1_crontab_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crontab_v1_crontab_proto_goTypes = []any{
	(*EmptyRes)(nil), // 0: user.v1.EmptyRes
	(*AddReq)(nil),   // 1: user.v1.AddReq
	(*AddRes)(nil),   // 2: user.v1.AddRes
	(*DelReq)(nil),   // 3: user.v1.DelReq
}
var file_crontab_v1_crontab_proto_depIdxs = []int32{
	1, // 0: user.v1.LiveCrontabRpcService.Add:input_type -> user.v1.AddReq
	3, // 1: user.v1.LiveCrontabRpcService.Del:input_type -> user.v1.DelReq
	2, // 2: user.v1.LiveCrontabRpcService.Add:output_type -> user.v1.AddRes
	0, // 3: user.v1.LiveCrontabRpcService.Del:output_type -> user.v1.EmptyRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crontab_v1_crontab_proto_init() }
func file_crontab_v1_crontab_proto_init() {
	if File_crontab_v1_crontab_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crontab_v1_crontab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crontab_v1_crontab_proto_goTypes,
		DependencyIndexes: file_crontab_v1_crontab_proto_depIdxs,
		MessageInfos:      file_crontab_v1_crontab_proto_msgTypes,
	}.Build()
	File_crontab_v1_crontab_proto = out.File
	file_crontab_v1_crontab_proto_rawDesc = nil
	file_crontab_v1_crontab_proto_goTypes = nil
	file_crontab_v1_crontab_proto_depIdxs = nil
}
