// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services/gateway/proto/gateway.proto

package proto

import (
	proto "github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_gateway_proto_gateway_proto protoreflect.FileDescriptor

var file_services_gateway_proto_gateway_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x75, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x6f, 0x72, 0x69, 0x2f, 0x6d, 0x69, 0x63,
	0x6f, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_services_gateway_proto_gateway_proto_goTypes = []interface{}{
	(*proto.CreateItemRequest)(nil),  // 0: catalog.CreateItemRequest
	(*proto.GetItemRequest)(nil),     // 1: catalog.GetItemRequest
	(*proto.CreateItemResponse)(nil), // 2: catalog.CreateItemResponse
	(*proto.GetItemResponse)(nil),    // 3: catalog.GetItemResponse
}
var file_services_gateway_proto_gateway_proto_depIdxs = []int32{
	0, // 0: gateway.GatewayService.CreateItem:input_type -> catalog.CreateItemRequest
	1, // 1: gateway.GatewayService.GetItem:input_type -> catalog.GetItemRequest
	2, // 2: gateway.GatewayService.CreateItem:output_type -> catalog.CreateItemResponse
	3, // 3: gateway.GatewayService.GetItem:output_type -> catalog.GetItemResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_gateway_proto_gateway_proto_init() }
func file_services_gateway_proto_gateway_proto_init() {
	if File_services_gateway_proto_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_gateway_proto_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_gateway_proto_gateway_proto_goTypes,
		DependencyIndexes: file_services_gateway_proto_gateway_proto_depIdxs,
	}.Build()
	File_services_gateway_proto_gateway_proto = out.File
	file_services_gateway_proto_gateway_proto_rawDesc = nil
	file_services_gateway_proto_gateway_proto_goTypes = nil
	file_services_gateway_proto_gateway_proto_depIdxs = nil
}
