// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: grpc/overlay.proto

package grpc

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_overlay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_overlay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_overlay_proto_rawDescGZIP(), []int{0}
}

var File_grpc_overlay_proto protoreflect.FileDescriptor

var file_grpc_overlay_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c,
	0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd7, 0x07, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c,
	0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e,
	0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0c, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e,
	0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x23, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x70,
	0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a,
	0x42, 0x32, 0x43, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x70, 0x65,
	0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x32, 0x43, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d,
	0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x32, 0x43, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c,
	0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x70, 0x65, 0x73,
	0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x32, 0x42, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x32, 0x42, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x70, 0x65,
	0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x32, 0x42, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x32, 0x42, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x32, 0x42, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65,
	0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x32, 0x42, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x12, 0x20, 0x2e, 0x6d, 0x70, 0x65, 0x73,
	0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d, 0x70,
	0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x51, 0x52, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x70,
	0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x70, 0x65,
	0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x27, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x6d, 0x70, 0x65,
	0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x54,
	0x61, 0x78, 0x12, 0x1e, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x78, 0x52,
	0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6d, 0x70, 0x65, 0x73, 0x61, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x69, 0x74, 0x54, 0x61, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_overlay_proto_rawDescOnce sync.Once
	file_grpc_overlay_proto_rawDescData = file_grpc_overlay_proto_rawDesc
)

func file_grpc_overlay_proto_rawDescGZIP() []byte {
	file_grpc_overlay_proto_rawDescOnce.Do(func() {
		file_grpc_overlay_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_overlay_proto_rawDescData)
	})
	return file_grpc_overlay_proto_rawDescData
}

var file_grpc_overlay_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_overlay_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: mpesaoverlay.grpc.Empty
	(*ExpressQueryReq)(nil),       // 1: mpesaoverlay.grpc.ExpressQueryReq
	(*ExpressSimulateReq)(nil),    // 2: mpesaoverlay.grpc.ExpressSimulateReq
	(*B2CPaymentReq)(nil),         // 3: mpesaoverlay.grpc.B2CPaymentReq
	(*AccountBalanceReq)(nil),     // 4: mpesaoverlay.grpc.AccountBalanceReq
	(*C2BRegisterURLReq)(nil),     // 5: mpesaoverlay.grpc.C2BRegisterURLReq
	(*C2BSimulateReq)(nil),        // 6: mpesaoverlay.grpc.C2BSimulateReq
	(*GenerateQRReq)(nil),         // 7: mpesaoverlay.grpc.GenerateQRReq
	(*ReverseReq)(nil),            // 8: mpesaoverlay.grpc.ReverseReq
	(*TransactionStatusReq)(nil),  // 9: mpesaoverlay.grpc.TransactionStatusReq
	(*RemitTaxReq)(nil),           // 10: mpesaoverlay.grpc.RemitTaxReq
	(*TokenResp)(nil),             // 11: mpesaoverlay.grpc.TokenResp
	(*ExpressQueryResp)(nil),      // 12: mpesaoverlay.grpc.ExpressQueryResp
	(*ExpressSimulateResp)(nil),   // 13: mpesaoverlay.grpc.ExpressSimulateResp
	(*B2CPaymentResp)(nil),        // 14: mpesaoverlay.grpc.B2CPaymentResp
	(*AccountBalanceResp)(nil),    // 15: mpesaoverlay.grpc.AccountBalanceResp
	(*C2BRegisterURLResp)(nil),    // 16: mpesaoverlay.grpc.C2BRegisterURLResp
	(*C2BSimulateResp)(nil),       // 17: mpesaoverlay.grpc.C2BSimulateResp
	(*GenerateQRResp)(nil),        // 18: mpesaoverlay.grpc.GenerateQRResp
	(*ReverseResp)(nil),           // 19: mpesaoverlay.grpc.ReverseResp
	(*TransactionStatusResp)(nil), // 20: mpesaoverlay.grpc.TransactionStatusResp
	(*RemitTaxResp)(nil),          // 21: mpesaoverlay.grpc.RemitTaxResp
}
var file_grpc_overlay_proto_depIdxs = []int32{
	0,  // 0: mpesaoverlay.grpc.Service.GetToken:input_type -> mpesaoverlay.grpc.Empty
	1,  // 1: mpesaoverlay.grpc.Service.ExpressQuery:input_type -> mpesaoverlay.grpc.ExpressQueryReq
	2,  // 2: mpesaoverlay.grpc.Service.ExpressSimulate:input_type -> mpesaoverlay.grpc.ExpressSimulateReq
	3,  // 3: mpesaoverlay.grpc.Service.B2CPayment:input_type -> mpesaoverlay.grpc.B2CPaymentReq
	4,  // 4: mpesaoverlay.grpc.Service.AccountBalance:input_type -> mpesaoverlay.grpc.AccountBalanceReq
	5,  // 5: mpesaoverlay.grpc.Service.C2BRegisterURL:input_type -> mpesaoverlay.grpc.C2BRegisterURLReq
	6,  // 6: mpesaoverlay.grpc.Service.C2BSimulate:input_type -> mpesaoverlay.grpc.C2BSimulateReq
	7,  // 7: mpesaoverlay.grpc.Service.GenerateQR:input_type -> mpesaoverlay.grpc.GenerateQRReq
	8,  // 8: mpesaoverlay.grpc.Service.Reverse:input_type -> mpesaoverlay.grpc.ReverseReq
	9,  // 9: mpesaoverlay.grpc.Service.TransactionStatus:input_type -> mpesaoverlay.grpc.TransactionStatusReq
	10, // 10: mpesaoverlay.grpc.Service.RemitTax:input_type -> mpesaoverlay.grpc.RemitTaxReq
	11, // 11: mpesaoverlay.grpc.Service.GetToken:output_type -> mpesaoverlay.grpc.TokenResp
	12, // 12: mpesaoverlay.grpc.Service.ExpressQuery:output_type -> mpesaoverlay.grpc.ExpressQueryResp
	13, // 13: mpesaoverlay.grpc.Service.ExpressSimulate:output_type -> mpesaoverlay.grpc.ExpressSimulateResp
	14, // 14: mpesaoverlay.grpc.Service.B2CPayment:output_type -> mpesaoverlay.grpc.B2CPaymentResp
	15, // 15: mpesaoverlay.grpc.Service.AccountBalance:output_type -> mpesaoverlay.grpc.AccountBalanceResp
	16, // 16: mpesaoverlay.grpc.Service.C2BRegisterURL:output_type -> mpesaoverlay.grpc.C2BRegisterURLResp
	17, // 17: mpesaoverlay.grpc.Service.C2BSimulate:output_type -> mpesaoverlay.grpc.C2BSimulateResp
	18, // 18: mpesaoverlay.grpc.Service.GenerateQR:output_type -> mpesaoverlay.grpc.GenerateQRResp
	19, // 19: mpesaoverlay.grpc.Service.Reverse:output_type -> mpesaoverlay.grpc.ReverseResp
	20, // 20: mpesaoverlay.grpc.Service.TransactionStatus:output_type -> mpesaoverlay.grpc.TransactionStatusResp
	21, // 21: mpesaoverlay.grpc.Service.RemitTax:output_type -> mpesaoverlay.grpc.RemitTaxResp
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_overlay_proto_init() }
func file_grpc_overlay_proto_init() {
	if File_grpc_overlay_proto != nil {
		return
	}
	file_grpc_requests_proto_init()
	file_grpc_responses_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_overlay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpc_overlay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_overlay_proto_goTypes,
		DependencyIndexes: file_grpc_overlay_proto_depIdxs,
		MessageInfos:      file_grpc_overlay_proto_msgTypes,
	}.Build()
	File_grpc_overlay_proto = out.File
	file_grpc_overlay_proto_rawDesc = nil
	file_grpc_overlay_proto_goTypes = nil
	file_grpc_overlay_proto_depIdxs = nil
}