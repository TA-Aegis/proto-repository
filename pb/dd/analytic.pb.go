// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: dd/analytic.proto

package dd

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

type AnalyticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *AnalyticRequest) Reset() {
	*x = AnalyticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dd_analytic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticRequest) ProtoMessage() {}

func (x *AnalyticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dd_analytic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticRequest.ProtoReflect.Descriptor instead.
func (*AnalyticRequest) Descriptor() ([]byte, []int) {
	return file_dd_analytic_proto_rawDescGZIP(), []int{0}
}

func (x *AnalyticRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type AnalyticData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId         string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TotalUsersInQueue int32                  `protobuf:"varint,2,opt,name=total_users_in_queue,json=totalUsersInQueue,proto3" json:"total_users_in_queue,omitempty"`
	TotalUsersInRoom  int32                  `protobuf:"varint,3,opt,name=total_users_in_room,json=totalUsersInRoom,proto3" json:"total_users_in_room,omitempty"`
	TotalUsers        int32                  `protobuf:"varint,4,opt,name=total_users,json=totalUsers,proto3" json:"total_users,omitempty"`
	Timestamp         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AnalyticData) Reset() {
	*x = AnalyticData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dd_analytic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticData) ProtoMessage() {}

func (x *AnalyticData) ProtoReflect() protoreflect.Message {
	mi := &file_dd_analytic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticData.ProtoReflect.Descriptor instead.
func (*AnalyticData) Descriptor() ([]byte, []int) {
	return file_dd_analytic_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyticData) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AnalyticData) GetTotalUsersInQueue() int32 {
	if x != nil {
		return x.TotalUsersInQueue
	}
	return 0
}

func (x *AnalyticData) GetTotalUsersInRoom() int32 {
	if x != nil {
		return x.TotalUsersInRoom
	}
	return 0
}

func (x *AnalyticData) GetTotalUsers() int32 {
	if x != nil {
		return x.TotalUsers
	}
	return 0
}

func (x *AnalyticData) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_dd_analytic_proto protoreflect.FileDescriptor

var file_dd_analytic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x64, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30,
	0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x12, 0x2d, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x5c, 0x0a, 0x0f, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x62,
	0x2f, 0x64, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dd_analytic_proto_rawDescOnce sync.Once
	file_dd_analytic_proto_rawDescData = file_dd_analytic_proto_rawDesc
)

func file_dd_analytic_proto_rawDescGZIP() []byte {
	file_dd_analytic_proto_rawDescOnce.Do(func() {
		file_dd_analytic_proto_rawDescData = protoimpl.X.CompressGZIP(file_dd_analytic_proto_rawDescData)
	})
	return file_dd_analytic_proto_rawDescData
}

var file_dd_analytic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dd_analytic_proto_goTypes = []interface{}{
	(*AnalyticRequest)(nil),       // 0: analytic.AnalyticRequest
	(*AnalyticData)(nil),          // 1: analytic.AnalyticData
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_dd_analytic_proto_depIdxs = []int32{
	2, // 0: analytic.AnalyticData.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: analytic.AnalyticService.StreamRealtimeData:input_type -> analytic.AnalyticRequest
	1, // 2: analytic.AnalyticService.StreamRealtimeData:output_type -> analytic.AnalyticData
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dd_analytic_proto_init() }
func file_dd_analytic_proto_init() {
	if File_dd_analytic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dd_analytic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticRequest); i {
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
		file_dd_analytic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticData); i {
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
			RawDescriptor: file_dd_analytic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dd_analytic_proto_goTypes,
		DependencyIndexes: file_dd_analytic_proto_depIdxs,
		MessageInfos:      file_dd_analytic_proto_msgTypes,
	}.Build()
	File_dd_analytic_proto = out.File
	file_dd_analytic_proto_rawDesc = nil
	file_dd_analytic_proto_goTypes = nil
	file_dd_analytic_proto_depIdxs = nil
}
