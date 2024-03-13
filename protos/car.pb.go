// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/car.proto

package protos

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CarName   string  `protobuf:"bytes,2,opt,name=car_name,json=carName,proto3" json:"car_name,omitempty"`
	DayRate   float64 `protobuf:"fixed64,3,opt,name=day_rate,json=dayRate,proto3" json:"day_rate,omitempty"`
	MonthRate float64 `protobuf:"fixed64,4,opt,name=month_rate,json=monthRate,proto3" json:"month_rate,omitempty"`
	Image     string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_protos_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_protos_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetCarName() string {
	if x != nil {
		return x.CarName
	}
	return ""
}

func (x *Car) GetDayRate() float64 {
	if x != nil {
		return x.DayRate
	}
	return 0
}

func (x *Car) GetMonthRate() float64 {
	if x != nil {
		return x.MonthRate
	}
	return 0
}

func (x *Car) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type ShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowRequest) Reset() {
	*x = ShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowRequest) ProtoMessage() {}

func (x *ShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowRequest.ProtoReflect.Descriptor instead.
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return file_protos_car_proto_rawDescGZIP(), []int{1}
}

func (x *ShowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarResponse) Reset() {
	*x = CarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarResponse) ProtoMessage() {}

func (x *CarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarResponse.ProtoReflect.Descriptor instead.
func (*CarResponse) Descriptor() ([]byte, []int) {
	return file_protos_car_proto_rawDescGZIP(), []int{2}
}

func (x *CarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

var File_protos_car_proto protoreflect.FileDescriptor

var file_protos_car_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x61, 0x72, 0x22, 0x80, 0x01, 0x0a,
	0x03, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x61, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x64, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x1d, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6e,
	0x74, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x32, 0x44,
	0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x53, 0x68, 0x6f, 0x77, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x61, 0x72, 0x2e,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x75, 0x66, 0x61, 0x6c, 0x64, 0x69, 0x6e, 0x74, 0x61, 0x31, 0x33,
	0x2f, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_car_proto_rawDescOnce sync.Once
	file_protos_car_proto_rawDescData = file_protos_car_proto_rawDesc
)

func file_protos_car_proto_rawDescGZIP() []byte {
	file_protos_car_proto_rawDescOnce.Do(func() {
		file_protos_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_car_proto_rawDescData)
	})
	return file_protos_car_proto_rawDescData
}

var file_protos_car_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_car_proto_goTypes = []interface{}{
	(*Car)(nil),         // 0: rent.car.Car
	(*ShowRequest)(nil), // 1: rent.car.ShowRequest
	(*CarResponse)(nil), // 2: rent.car.CarResponse
}
var file_protos_car_proto_depIdxs = []int32{
	0, // 0: rent.car.CarResponse.car:type_name -> rent.car.Car
	1, // 1: rent.car.CarService.Show:input_type -> rent.car.ShowRequest
	2, // 2: rent.car.CarService.Show:output_type -> rent.car.CarResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_car_proto_init() }
func file_protos_car_proto_init() {
	if File_protos_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_protos_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowRequest); i {
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
		file_protos_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarResponse); i {
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
			RawDescriptor: file_protos_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_car_proto_goTypes,
		DependencyIndexes: file_protos_car_proto_depIdxs,
		MessageInfos:      file_protos_car_proto_msgTypes,
	}.Build()
	File_protos_car_proto = out.File
	file_protos_car_proto_rawDesc = nil
	file_protos_car_proto_goTypes = nil
	file_protos_car_proto_depIdxs = nil
}
