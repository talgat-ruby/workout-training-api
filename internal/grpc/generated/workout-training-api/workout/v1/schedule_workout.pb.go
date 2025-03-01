// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: workout-training-api/workout/v1/schedule_workout.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScheduleWorkoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkoutId     string                 `protobuf:"bytes,1,opt,name=workout_id,json=workoutId,proto3" json:"workout_id,omitempty"`
	ScheduledDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scheduled_date,json=scheduledDate,proto3" json:"scheduled_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleWorkoutRequest) Reset() {
	*x = ScheduleWorkoutRequest{}
	mi := &file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleWorkoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleWorkoutRequest) ProtoMessage() {}

func (x *ScheduleWorkoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleWorkoutRequest.ProtoReflect.Descriptor instead.
func (*ScheduleWorkoutRequest) Descriptor() ([]byte, []int) {
	return file_workout_training_api_workout_v1_schedule_workout_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleWorkoutRequest) GetWorkoutId() string {
	if x != nil {
		return x.WorkoutId
	}
	return ""
}

func (x *ScheduleWorkoutRequest) GetScheduledDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledDate
	}
	return nil
}

type ScheduleWorkoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleWorkoutResponse) Reset() {
	*x = ScheduleWorkoutResponse{}
	mi := &file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleWorkoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleWorkoutResponse) ProtoMessage() {}

func (x *ScheduleWorkoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleWorkoutResponse.ProtoReflect.Descriptor instead.
func (*ScheduleWorkoutResponse) Descriptor() ([]byte, []int) {
	return file_workout_training_api_workout_v1_schedule_workout_proto_rawDescGZIP(), []int{1}
}

var File_workout_training_api_workout_v1_schedule_workout_proto protoreflect.FileDescriptor

var file_workout_training_api_workout_v1_schedule_workout_proto_rawDesc = string([]byte{
	0x0a, 0x36, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x16, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x41, 0x0a,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x19, 0x0a, 0x17, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x77,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_workout_training_api_workout_v1_schedule_workout_proto_rawDescOnce sync.Once
	file_workout_training_api_workout_v1_schedule_workout_proto_rawDescData []byte
)

func file_workout_training_api_workout_v1_schedule_workout_proto_rawDescGZIP() []byte {
	file_workout_training_api_workout_v1_schedule_workout_proto_rawDescOnce.Do(func() {
		file_workout_training_api_workout_v1_schedule_workout_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_workout_training_api_workout_v1_schedule_workout_proto_rawDesc), len(file_workout_training_api_workout_v1_schedule_workout_proto_rawDesc)))
	})
	return file_workout_training_api_workout_v1_schedule_workout_proto_rawDescData
}

var file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_workout_training_api_workout_v1_schedule_workout_proto_goTypes = []any{
	(*ScheduleWorkoutRequest)(nil),  // 0: workout_training_api.ping.v1.ScheduleWorkoutRequest
	(*ScheduleWorkoutResponse)(nil), // 1: workout_training_api.ping.v1.ScheduleWorkoutResponse
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
}
var file_workout_training_api_workout_v1_schedule_workout_proto_depIdxs = []int32{
	2, // 0: workout_training_api.ping.v1.ScheduleWorkoutRequest.scheduled_date:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_workout_training_api_workout_v1_schedule_workout_proto_init() }
func file_workout_training_api_workout_v1_schedule_workout_proto_init() {
	if File_workout_training_api_workout_v1_schedule_workout_proto != nil {
		return
	}
	file_workout_training_api_workout_v1_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workout_training_api_workout_v1_schedule_workout_proto_rawDesc), len(file_workout_training_api_workout_v1_schedule_workout_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workout_training_api_workout_v1_schedule_workout_proto_goTypes,
		DependencyIndexes: file_workout_training_api_workout_v1_schedule_workout_proto_depIdxs,
		MessageInfos:      file_workout_training_api_workout_v1_schedule_workout_proto_msgTypes,
	}.Build()
	File_workout_training_api_workout_v1_schedule_workout_proto = out.File
	file_workout_training_api_workout_v1_schedule_workout_proto_goTypes = nil
	file_workout_training_api_workout_v1_schedule_workout_proto_depIdxs = nil
}
