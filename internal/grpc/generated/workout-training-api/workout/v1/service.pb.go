// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: workout-training-api/workout/v1/service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_workout_training_api_workout_v1_service_proto protoreflect.FileDescriptor

var file_workout_training_api_workout_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x77,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x77, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xf5, 0x03, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x78, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x12, 0x32, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x32, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x77,
	0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x75, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var file_workout_training_api_workout_v1_service_proto_goTypes = []any{
	(*CreateWorkoutRequest)(nil),  // 0: workout_training_api.ping.v1.CreateWorkoutRequest
	(*UpdateWorkoutRequest)(nil),  // 1: workout_training_api.ping.v1.UpdateWorkoutRequest
	(*DeleteWorkoutRequest)(nil),  // 2: workout_training_api.ping.v1.DeleteWorkoutRequest
	(*ListWorkoutsRequest)(nil),   // 3: workout_training_api.ping.v1.ListWorkoutsRequest
	(*CreateWorkoutResponse)(nil), // 4: workout_training_api.ping.v1.CreateWorkoutResponse
	(*UpdateWorkoutResponse)(nil), // 5: workout_training_api.ping.v1.UpdateWorkoutResponse
	(*WorkoutDeleteResponse)(nil), // 6: workout_training_api.ping.v1.WorkoutDeleteResponse
	(*ListWorkoutsResponse)(nil),  // 7: workout_training_api.ping.v1.ListWorkoutsResponse
}
var file_workout_training_api_workout_v1_service_proto_depIdxs = []int32{
	0, // 0: workout_training_api.ping.v1.WorkoutService.CreateWorkout:input_type -> workout_training_api.ping.v1.CreateWorkoutRequest
	1, // 1: workout_training_api.ping.v1.WorkoutService.UpdateWorkout:input_type -> workout_training_api.ping.v1.UpdateWorkoutRequest
	2, // 2: workout_training_api.ping.v1.WorkoutService.DeleteWorkout:input_type -> workout_training_api.ping.v1.DeleteWorkoutRequest
	3, // 3: workout_training_api.ping.v1.WorkoutService.ListWorkouts:input_type -> workout_training_api.ping.v1.ListWorkoutsRequest
	4, // 4: workout_training_api.ping.v1.WorkoutService.CreateWorkout:output_type -> workout_training_api.ping.v1.CreateWorkoutResponse
	5, // 5: workout_training_api.ping.v1.WorkoutService.UpdateWorkout:output_type -> workout_training_api.ping.v1.UpdateWorkoutResponse
	6, // 6: workout_training_api.ping.v1.WorkoutService.DeleteWorkout:output_type -> workout_training_api.ping.v1.WorkoutDeleteResponse
	7, // 7: workout_training_api.ping.v1.WorkoutService.ListWorkouts:output_type -> workout_training_api.ping.v1.ListWorkoutsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_workout_training_api_workout_v1_service_proto_init() }
func file_workout_training_api_workout_v1_service_proto_init() {
	if File_workout_training_api_workout_v1_service_proto != nil {
		return
	}
	file_workout_training_api_workout_v1_create_workout_proto_init()
	file_workout_training_api_workout_v1_update_workout_proto_init()
	file_workout_training_api_workout_v1_delete_workout_proto_init()
	file_workout_training_api_workout_v1_list_workout_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workout_training_api_workout_v1_service_proto_rawDesc), len(file_workout_training_api_workout_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workout_training_api_workout_v1_service_proto_goTypes,
		DependencyIndexes: file_workout_training_api_workout_v1_service_proto_depIdxs,
	}.Build()
	File_workout_training_api_workout_v1_service_proto = out.File
	file_workout_training_api_workout_v1_service_proto_goTypes = nil
	file_workout_training_api_workout_v1_service_proto_depIdxs = nil
}
