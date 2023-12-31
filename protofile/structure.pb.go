// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.12.4
// source: structure.proto

package protofile

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type OrderType int32

const (
	OrderType_TYPE_UNSPECIFIED OrderType = 0
	OrderType_TYPE_MARKET      OrderType = 1
	OrderType_TYPE_LIMIT       OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_MARKET",
		2: "TYPE_LIMIT",
	}
	OrderType_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_MARKET":      1,
		"TYPE_LIMIT":       2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_structure_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_structure_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_structure_proto_rawDescGZIP(), []int{0}
}

type OrderStatus int32

const (
	OrderStatus_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_STATUS_NEW         OrderStatus = 1
	OrderStatus_STATUS_PART_FILL   OrderStatus = 2
	OrderStatus_STATUS_FILL        OrderStatus = 3
	OrderStatus_STATUS_DONE        OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_NEW",
		2: "STATUS_PART_FILL",
		3: "STATUS_FILL",
		4: "STATUS_DONE",
	}
	OrderStatus_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_NEW":         1,
		"STATUS_PART_FILL":   2,
		"STATUS_FILL":        3,
		"STATUS_DONE":        4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_structure_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_structure_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_structure_proto_rawDescGZIP(), []int{1}
}

type Direction int32

const (
	Direction_DIRECTION_UNSPECIFIED Direction = 0
	Direction_DIRECTION_SELL        Direction = 1
	Direction_DIRECTION_BUY         Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "DIRECTION_SELL",
		2: "DIRECTION_BUY",
	}
	Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"DIRECTION_SELL":        1,
		"DIRECTION_BUY":         2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_structure_proto_enumTypes[2].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_structure_proto_enumTypes[2]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_structure_proto_rawDescGZIP(), []int{2}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string               `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UserId      string               `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Symbol      string               `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price       string               `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Volume      string               `protobuf:"bytes,5,opt,name=volume,proto3" json:"volume,omitempty"`
	OrderType   OrderType            `protobuf:"varint,6,opt,name=orderType,proto3,enum=protofile.OrderType" json:"orderType,omitempty"`
	OrderStatus OrderStatus          `protobuf:"varint,7,opt,name=orderStatus,proto3,enum=protofile.OrderStatus" json:"orderStatus,omitempty"`
	CreateDate  *timestamp.Timestamp `protobuf:"bytes,8,opt,name=createDate,proto3" json:"createDate,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_structure_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CreateOrderRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateOrderRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *CreateOrderRequest) GetOrderType() OrderType {
	if x != nil {
		return x.OrderType
	}
	return OrderType_TYPE_UNSPECIFIED
}

func (x *CreateOrderRequest) GetOrderStatus() OrderStatus {
	if x != nil {
		return x.OrderStatus
	}
	return OrderStatus_STATUS_UNSPECIFIED
}

func (x *CreateOrderRequest) GetCreateDate() *timestamp.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

var File_structure_proto protoreflect.FileDescriptor

var file_structure_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x38, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x2a, 0x42, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x2a, 0x6d, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x04, 0x2a, 0x4d, 0x0a, 0x09, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x02, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_structure_proto_rawDescOnce sync.Once
	file_structure_proto_rawDescData = file_structure_proto_rawDesc
)

func file_structure_proto_rawDescGZIP() []byte {
	file_structure_proto_rawDescOnce.Do(func() {
		file_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_structure_proto_rawDescData)
	})
	return file_structure_proto_rawDescData
}

var file_structure_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_structure_proto_goTypes = []interface{}{
	(OrderType)(0),              // 0: protofile.OrderType
	(OrderStatus)(0),            // 1: protofile.OrderStatus
	(Direction)(0),              // 2: protofile.Direction
	(*CreateOrderRequest)(nil),  // 3: protofile.CreateOrderRequest
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_structure_proto_depIdxs = []int32{
	0, // 0: protofile.CreateOrderRequest.orderType:type_name -> protofile.OrderType
	1, // 1: protofile.CreateOrderRequest.orderStatus:type_name -> protofile.OrderStatus
	4, // 2: protofile.CreateOrderRequest.createDate:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_structure_proto_init() }
func file_structure_proto_init() {
	if File_structure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
			RawDescriptor: file_structure_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_structure_proto_goTypes,
		DependencyIndexes: file_structure_proto_depIdxs,
		EnumInfos:         file_structure_proto_enumTypes,
		MessageInfos:      file_structure_proto_msgTypes,
	}.Build()
	File_structure_proto = out.File
	file_structure_proto_rawDesc = nil
	file_structure_proto_goTypes = nil
	file_structure_proto_depIdxs = nil
}
