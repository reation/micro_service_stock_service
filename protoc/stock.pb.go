// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: stock.proto

package protoc

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

type OrderGoodsID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId  int64 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	GoodsNum int64 `protobuf:"varint,2,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
}

func (x *OrderGoodsID) Reset() {
	*x = OrderGoodsID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGoodsID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGoodsID) ProtoMessage() {}

func (x *OrderGoodsID) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGoodsID.ProtoReflect.Descriptor instead.
func (*OrderGoodsID) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *OrderGoodsID) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *OrderGoodsID) GetGoodsNum() int64 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

type OrderReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int64 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Demand  int64 `protobuf:"varint,3,opt,name=demand,proto3" json:"demand,omitempty"`
	Stock   int64 `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Deal    int64 `protobuf:"varint,5,opt,name=deal,proto3" json:"deal,omitempty"`
}

func (x *OrderReturn) Reset() {
	*x = OrderReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReturn) ProtoMessage() {}

func (x *OrderReturn) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReturn.ProtoReflect.Descriptor instead.
func (*OrderReturn) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{1}
}

func (x *OrderReturn) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *OrderReturn) GetDemand() int64 {
	if x != nil {
		return x.Demand
	}
	return 0
}

func (x *OrderReturn) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *OrderReturn) GetDeal() int64 {
	if x != nil {
		return x.Deal
	}
	return 0
}

type OrderReduceStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsID []*OrderGoodsID `protobuf:"bytes,1,rep,name=GoodsID,proto3" json:"GoodsID,omitempty"`
	OrderID int64           `protobuf:"varint,2,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	IsAll   int64           `protobuf:"varint,3,opt,name=IsAll,proto3" json:"IsAll,omitempty"`
}

func (x *OrderReduceStockRequest) Reset() {
	*x = OrderReduceStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReduceStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReduceStockRequest) ProtoMessage() {}

func (x *OrderReduceStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReduceStockRequest.ProtoReflect.Descriptor instead.
func (*OrderReduceStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{2}
}

func (x *OrderReduceStockRequest) GetGoodsID() []*OrderGoodsID {
	if x != nil {
		return x.GoodsID
	}
	return nil
}

func (x *OrderReduceStockRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *OrderReduceStockRequest) GetIsAll() int64 {
	if x != nil {
		return x.IsAll
	}
	return 0
}

type OrderReduceStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States      int64          `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	GoodsStates []*OrderReturn `protobuf:"bytes,2,rep,name=goods_states,json=goodsStates,proto3" json:"goods_states,omitempty"`
}

func (x *OrderReduceStockResponse) Reset() {
	*x = OrderReduceStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReduceStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReduceStockResponse) ProtoMessage() {}

func (x *OrderReduceStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReduceStockResponse.ProtoReflect.Descriptor instead.
func (*OrderReduceStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{3}
}

func (x *OrderReduceStockResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *OrderReduceStockResponse) GetGoodsStates() []*OrderReturn {
	if x != nil {
		return x.GoodsStates
	}
	return nil
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x46, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x22, 0x6a, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x22, 0x78, 0x0a, 0x17, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x44, 0x52, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x73, 0x41, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x73, 0x41,
	0x6c, 0x6c, 0x22, 0x69, 0x0a, 0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x32, 0x70, 0x0a,
	0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x10, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stock_proto_goTypes = []interface{}{
	(*OrderGoodsID)(nil),             // 0: stock.OrderGoodsID
	(*OrderReturn)(nil),              // 1: stock.OrderReturn
	(*OrderReduceStockRequest)(nil),  // 2: stock.OrderReduceStockRequest
	(*OrderReduceStockResponse)(nil), // 3: stock.OrderReduceStockResponse
}
var file_stock_proto_depIdxs = []int32{
	0, // 0: stock.OrderReduceStockRequest.GoodsID:type_name -> stock.OrderGoodsID
	1, // 1: stock.OrderReduceStockResponse.goods_states:type_name -> stock.OrderReturn
	2, // 2: stock.OrderReduceStockOperation.OrderReduceStock:input_type -> stock.OrderReduceStockRequest
	3, // 3: stock.OrderReduceStockOperation.OrderReduceStock:output_type -> stock.OrderReduceStockResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGoodsID); i {
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
		file_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReturn); i {
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
		file_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReduceStockRequest); i {
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
		file_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReduceStockResponse); i {
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
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}
