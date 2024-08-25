// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: internal/proto/service.proto

package proto

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

type CandleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *CandleRequest) Reset() {
	*x = CandleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandleRequest) ProtoMessage() {}

func (x *CandleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandleRequest.ProtoReflect.Descriptor instead.
func (*CandleRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *CandleRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type Candle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Timestamp int64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Open      float64 `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`
	High      float64 `protobuf:"fixed64,4,opt,name=high,proto3" json:"high,omitempty"`
	Low       float64 `protobuf:"fixed64,5,opt,name=low,proto3" json:"low,omitempty"`
	Close     float64 `protobuf:"fixed64,6,opt,name=close,proto3" json:"close,omitempty"`
	Volume    float64 `protobuf:"fixed64,7,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Candle) Reset() {
	*x = Candle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_internal_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Candle) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Candle) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Candle) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Candle) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Candle) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Candle) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Candle) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

var File_internal_proto_service_proto protoreflect.FileDescriptor

var file_internal_proto_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x22, 0xa6, 0x01, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x4c, 0x0a, 0x0e, 0x54, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x74,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_service_proto_rawDescOnce sync.Once
	file_internal_proto_service_proto_rawDescData = file_internal_proto_service_proto_rawDesc
)

func file_internal_proto_service_proto_rawDescGZIP() []byte {
	file_internal_proto_service_proto_rawDescOnce.Do(func() {
		file_internal_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_service_proto_rawDescData)
	})
	return file_internal_proto_service_proto_rawDescData
}

var file_internal_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_service_proto_goTypes = []any{
	(*CandleRequest)(nil), // 0: trading.CandleRequest
	(*Candle)(nil),        // 1: trading.Candle
}
var file_internal_proto_service_proto_depIdxs = []int32{
	0, // 0: trading.TradingService.StreamCandles:input_type -> trading.CandleRequest
	1, // 1: trading.TradingService.StreamCandles:output_type -> trading.Candle
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_service_proto_init() }
func file_internal_proto_service_proto_init() {
	if File_internal_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CandleRequest); i {
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
		file_internal_proto_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Candle); i {
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
			RawDescriptor: file_internal_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_service_proto_goTypes,
		DependencyIndexes: file_internal_proto_service_proto_depIdxs,
		MessageInfos:      file_internal_proto_service_proto_msgTypes,
	}.Build()
	File_internal_proto_service_proto = out.File
	file_internal_proto_service_proto_rawDesc = nil
	file_internal_proto_service_proto_goTypes = nil
	file_internal_proto_service_proto_depIdxs = nil
}