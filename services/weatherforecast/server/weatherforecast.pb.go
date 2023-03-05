// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: weatherforecast.proto

package server

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

type WeatherForecastResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *WeatherForecastResult) Reset() {
	*x = WeatherForecastResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherforecast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherForecastResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherForecastResult) ProtoMessage() {}

func (x *WeatherForecastResult) ProtoReflect() protoreflect.Message {
	mi := &file_weatherforecast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherForecastResult.ProtoReflect.Descriptor instead.
func (*WeatherForecastResult) Descriptor() ([]byte, []int) {
	return file_weatherforecast_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherForecastResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetWeatherForecastByCityNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityName string `protobuf:"bytes,1,opt,name=cityName,proto3" json:"cityName,omitempty"`
}

func (x *GetWeatherForecastByCityNameRequest) Reset() {
	*x = GetWeatherForecastByCityNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherforecast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherForecastByCityNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherForecastByCityNameRequest) ProtoMessage() {}

func (x *GetWeatherForecastByCityNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weatherforecast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherForecastByCityNameRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherForecastByCityNameRequest) Descriptor() ([]byte, []int) {
	return file_weatherforecast_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeatherForecastByCityNameRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

type GetWeatherForecastByCityNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *WeatherForecastResult `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetWeatherForecastByCityNameResponse) Reset() {
	*x = GetWeatherForecastByCityNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherforecast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherForecastByCityNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherForecastByCityNameResponse) ProtoMessage() {}

func (x *GetWeatherForecastByCityNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weatherforecast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherForecastByCityNameResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherForecastByCityNameResponse) Descriptor() ([]byte, []int) {
	return file_weatherforecast_proto_rawDescGZIP(), []int{2}
}

func (x *GetWeatherForecastByCityNameResponse) GetData() *WeatherForecastResult {
	if x != nil {
		return x.Data
	}
	return nil
}

type FillWeatherForecastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityName string `protobuf:"bytes,1,opt,name=cityName,proto3" json:"cityName,omitempty"`
}

func (x *FillWeatherForecastRequest) Reset() {
	*x = FillWeatherForecastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weatherforecast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FillWeatherForecastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FillWeatherForecastRequest) ProtoMessage() {}

func (x *FillWeatherForecastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weatherforecast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FillWeatherForecastRequest.ProtoReflect.Descriptor instead.
func (*FillWeatherForecastRequest) Descriptor() ([]byte, []int) {
	return file_weatherforecast_proto_rawDescGZIP(), []int{3}
}

func (x *FillWeatherForecastRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

var File_weatherforecast_proto protoreflect.FileDescriptor

var file_weatherforecast_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x41, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x42, 0x79, 0x43, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x42, 0x79,
	0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x1a,
	0x46, 0x69, 0x6c, 0x6c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74,
	0x42, 0x79, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x42,
	0x79, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x42, 0x79, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13, 0x46, 0x69, 0x6c,
	0x6c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x1b, 0x2e, 0x46, 0x69, 0x6c, 0x6c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_weatherforecast_proto_rawDescOnce sync.Once
	file_weatherforecast_proto_rawDescData = file_weatherforecast_proto_rawDesc
)

func file_weatherforecast_proto_rawDescGZIP() []byte {
	file_weatherforecast_proto_rawDescOnce.Do(func() {
		file_weatherforecast_proto_rawDescData = protoimpl.X.CompressGZIP(file_weatherforecast_proto_rawDescData)
	})
	return file_weatherforecast_proto_rawDescData
}

var file_weatherforecast_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weatherforecast_proto_goTypes = []interface{}{
	(*WeatherForecastResult)(nil),                // 0: WeatherForecastResult
	(*GetWeatherForecastByCityNameRequest)(nil),  // 1: GetWeatherForecastByCityNameRequest
	(*GetWeatherForecastByCityNameResponse)(nil), // 2: GetWeatherForecastByCityNameResponse
	(*FillWeatherForecastRequest)(nil),           // 3: FillWeatherForecastRequest
	(*empty.Empty)(nil),                          // 4: google.protobuf.Empty
}
var file_weatherforecast_proto_depIdxs = []int32{
	0, // 0: GetWeatherForecastByCityNameResponse.data:type_name -> WeatherForecastResult
	1, // 1: WeatherForecast.GetWeatherForecastByCityName:input_type -> GetWeatherForecastByCityNameRequest
	3, // 2: WeatherForecast.FillWeatherForecast:input_type -> FillWeatherForecastRequest
	2, // 3: WeatherForecast.GetWeatherForecastByCityName:output_type -> GetWeatherForecastByCityNameResponse
	4, // 4: WeatherForecast.FillWeatherForecast:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weatherforecast_proto_init() }
func file_weatherforecast_proto_init() {
	if File_weatherforecast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weatherforecast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherForecastResult); i {
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
		file_weatherforecast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherForecastByCityNameRequest); i {
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
		file_weatherforecast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherForecastByCityNameResponse); i {
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
		file_weatherforecast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FillWeatherForecastRequest); i {
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
			RawDescriptor: file_weatherforecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weatherforecast_proto_goTypes,
		DependencyIndexes: file_weatherforecast_proto_depIdxs,
		MessageInfos:      file_weatherforecast_proto_msgTypes,
	}.Build()
	File_weatherforecast_proto = out.File
	file_weatherforecast_proto_rawDesc = nil
	file_weatherforecast_proto_goTypes = nil
	file_weatherforecast_proto_depIdxs = nil
}