// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: stock_goods_list.proto

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

type GetGoodsStockIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int64 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
}

func (x *GetGoodsStockIDList) Reset() {
	*x = GetGoodsStockIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_goods_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsStockIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsStockIDList) ProtoMessage() {}

func (x *GetGoodsStockIDList) ProtoReflect() protoreflect.Message {
	mi := &file_stock_goods_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsStockIDList.ProtoReflect.Descriptor instead.
func (*GetGoodsStockIDList) Descriptor() ([]byte, []int) {
	return file_stock_goods_list_proto_rawDescGZIP(), []int{0}
}

func (x *GetGoodsStockIDList) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

type GetGoodsStockDataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId  int64 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	GoodsNum int64 `protobuf:"varint,2,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
}

func (x *GetGoodsStockDataList) Reset() {
	*x = GetGoodsStockDataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_goods_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsStockDataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsStockDataList) ProtoMessage() {}

func (x *GetGoodsStockDataList) ProtoReflect() protoreflect.Message {
	mi := &file_stock_goods_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsStockDataList.ProtoReflect.Descriptor instead.
func (*GetGoodsStockDataList) Descriptor() ([]byte, []int) {
	return file_stock_goods_list_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsStockDataList) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GetGoodsStockDataList) GetGoodsNum() int64 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

type GetGoodsStockListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsIDList []*GetGoodsStockIDList `protobuf:"bytes,1,rep,name=goodsIDList,proto3" json:"goodsIDList,omitempty"`
}

func (x *GetGoodsStockListRequest) Reset() {
	*x = GetGoodsStockListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_goods_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsStockListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsStockListRequest) ProtoMessage() {}

func (x *GetGoodsStockListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_goods_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsStockListRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsStockListRequest) Descriptor() ([]byte, []int) {
	return file_stock_goods_list_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodsStockListRequest) GetGoodsIDList() []*GetGoodsStockIDList {
	if x != nil {
		return x.GoodsIDList
	}
	return nil
}

type GetGoodsStockListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States         int64                    `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	GoodsStockList []*GetGoodsStockDataList `protobuf:"bytes,2,rep,name=GoodsStockList,proto3" json:"GoodsStockList,omitempty"`
}

func (x *GetGoodsStockListResponse) Reset() {
	*x = GetGoodsStockListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_goods_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsStockListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsStockListResponse) ProtoMessage() {}

func (x *GetGoodsStockListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_goods_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsStockListResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsStockListResponse) Descriptor() ([]byte, []int) {
	return file_stock_goods_list_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoodsStockListResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *GetGoodsStockListResponse) GetGoodsStockList() []*GetGoodsStockDataList {
	if x != nil {
		return x.GoodsStockList
	}
	return nil
}

var File_stock_goods_list_proto protoreflect.FileDescriptor

var file_stock_goods_list_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x22, 0x5e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x44, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x7f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a,
	0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x80, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x6b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x42, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_goods_list_proto_rawDescOnce sync.Once
	file_stock_goods_list_proto_rawDescData = file_stock_goods_list_proto_rawDesc
)

func file_stock_goods_list_proto_rawDescGZIP() []byte {
	file_stock_goods_list_proto_rawDescOnce.Do(func() {
		file_stock_goods_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_goods_list_proto_rawDescData)
	})
	return file_stock_goods_list_proto_rawDescData
}

var file_stock_goods_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stock_goods_list_proto_goTypes = []interface{}{
	(*GetGoodsStockIDList)(nil),       // 0: stock_goods.GetGoodsStockIDList
	(*GetGoodsStockDataList)(nil),     // 1: stock_goods.GetGoodsStockDataList
	(*GetGoodsStockListRequest)(nil),  // 2: stock_goods.GetGoodsStockListRequest
	(*GetGoodsStockListResponse)(nil), // 3: stock_goods.GetGoodsStockListResponse
}
var file_stock_goods_list_proto_depIdxs = []int32{
	0, // 0: stock_goods.GetGoodsStockListRequest.goodsIDList:type_name -> stock_goods.GetGoodsStockIDList
	1, // 1: stock_goods.GetGoodsStockListResponse.GoodsStockList:type_name -> stock_goods.GetGoodsStockDataList
	2, // 2: stock_goods.GetGoodsStockList.GetGoodsStockByGoodsIDList:input_type -> stock_goods.GetGoodsStockListRequest
	3, // 3: stock_goods.GetGoodsStockList.GetGoodsStockByGoodsIDList:output_type -> stock_goods.GetGoodsStockListResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stock_goods_list_proto_init() }
func file_stock_goods_list_proto_init() {
	if File_stock_goods_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_goods_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsStockIDList); i {
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
		file_stock_goods_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsStockDataList); i {
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
		file_stock_goods_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsStockListRequest); i {
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
		file_stock_goods_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsStockListResponse); i {
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
			RawDescriptor: file_stock_goods_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_goods_list_proto_goTypes,
		DependencyIndexes: file_stock_goods_list_proto_depIdxs,
		MessageInfos:      file_stock_goods_list_proto_msgTypes,
	}.Build()
	File_stock_goods_list_proto = out.File
	file_stock_goods_list_proto_rawDesc = nil
	file_stock_goods_list_proto_goTypes = nil
	file_stock_goods_list_proto_depIdxs = nil
}
