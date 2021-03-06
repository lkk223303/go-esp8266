// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: temp.proto

package pb

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

type TempEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId   int32   `protobuf:"varint,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	EventId    int32   `protobuf:"varint,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	Humidity   float32 `protobuf:"fixed32,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	TempCel    float32 `protobuf:"fixed32,4,opt,name=tempCel,proto3" json:"tempCel,omitempty"`
	HeatIdxCel float32 `protobuf:"fixed32,5,opt,name=heatIdxCel,proto3" json:"heatIdxCel,omitempty"`
}

func (x *TempEvent) Reset() {
	*x = TempEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempEvent) ProtoMessage() {}

func (x *TempEvent) ProtoReflect() protoreflect.Message {
	mi := &file_temp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TempEvent.ProtoReflect.Descriptor instead.
func (*TempEvent) Descriptor() ([]byte, []int) {
	return file_temp_proto_rawDescGZIP(), []int{0}
}

func (x *TempEvent) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *TempEvent) GetEventId() int32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *TempEvent) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *TempEvent) GetTempCel() float32 {
	if x != nil {
		return x.TempCel
	}
	return 0
}

func (x *TempEvent) GetHeatIdxCel() float32 {
	if x != nil {
		return x.HeatIdxCel
	}
	return 0
}

var File_temp_proto protoreflect.FileDescriptor

var file_temp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x97, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65,
	0x61, 0x74, 0x49, 0x64, 0x78, 0x43, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x68, 0x65, 0x61, 0x74, 0x49, 0x64, 0x78, 0x43, 0x65, 0x6c, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temp_proto_rawDescOnce sync.Once
	file_temp_proto_rawDescData = file_temp_proto_rawDesc
)

func file_temp_proto_rawDescGZIP() []byte {
	file_temp_proto_rawDescOnce.Do(func() {
		file_temp_proto_rawDescData = protoimpl.X.CompressGZIP(file_temp_proto_rawDescData)
	})
	return file_temp_proto_rawDescData
}

var file_temp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temp_proto_goTypes = []interface{}{
	(*TempEvent)(nil), // 0: pb.TempEvent
}
var file_temp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temp_proto_init() }
func file_temp_proto_init() {
	if File_temp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TempEvent); i {
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
			RawDescriptor: file_temp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temp_proto_goTypes,
		DependencyIndexes: file_temp_proto_depIdxs,
		MessageInfos:      file_temp_proto_msgTypes,
	}.Build()
	File_temp_proto = out.File
	file_temp_proto_rawDesc = nil
	file_temp_proto_goTypes = nil
	file_temp_proto_depIdxs = nil
}
