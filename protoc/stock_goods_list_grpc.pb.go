// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: stock_goods_list.proto

package protoc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GetGoodsStockListClient is the client API for GetGoodsStockList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetGoodsStockListClient interface {
	GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockListRequest, opts ...grpc.CallOption) (*GetGoodsStockListResponse, error)
}

type getGoodsStockListClient struct {
	cc grpc.ClientConnInterface
}

func NewGetGoodsStockListClient(cc grpc.ClientConnInterface) GetGoodsStockListClient {
	return &getGoodsStockListClient{cc}
}

func (c *getGoodsStockListClient) GetGoodsStockByGoodsIDList(ctx context.Context, in *GetGoodsStockListRequest, opts ...grpc.CallOption) (*GetGoodsStockListResponse, error) {
	out := new(GetGoodsStockListResponse)
	err := c.cc.Invoke(ctx, "/stock_goods.GetGoodsStockList/GetGoodsStockByGoodsIDList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetGoodsStockListServer is the server API for GetGoodsStockList service.
// All implementations must embed UnimplementedGetGoodsStockListServer
// for forward compatibility
type GetGoodsStockListServer interface {
	GetGoodsStockByGoodsIDList(context.Context, *GetGoodsStockListRequest) (*GetGoodsStockListResponse, error)
	mustEmbedUnimplementedGetGoodsStockListServer()
}

// UnimplementedGetGoodsStockListServer must be embedded to have forward compatible implementations.
type UnimplementedGetGoodsStockListServer struct {
}

func (UnimplementedGetGoodsStockListServer) GetGoodsStockByGoodsIDList(context.Context, *GetGoodsStockListRequest) (*GetGoodsStockListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsStockByGoodsIDList not implemented")
}
func (UnimplementedGetGoodsStockListServer) mustEmbedUnimplementedGetGoodsStockListServer() {}

// UnsafeGetGoodsStockListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetGoodsStockListServer will
// result in compilation errors.
type UnsafeGetGoodsStockListServer interface {
	mustEmbedUnimplementedGetGoodsStockListServer()
}

func RegisterGetGoodsStockListServer(s grpc.ServiceRegistrar, srv GetGoodsStockListServer) {
	s.RegisterService(&GetGoodsStockList_ServiceDesc, srv)
}

func _GetGoodsStockList_GetGoodsStockByGoodsIDList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsStockListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGoodsStockListServer).GetGoodsStockByGoodsIDList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_goods.GetGoodsStockList/GetGoodsStockByGoodsIDList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGoodsStockListServer).GetGoodsStockByGoodsIDList(ctx, req.(*GetGoodsStockListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetGoodsStockList_ServiceDesc is the grpc.ServiceDesc for GetGoodsStockList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetGoodsStockList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock_goods.GetGoodsStockList",
	HandlerType: (*GetGoodsStockListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsStockByGoodsIDList",
			Handler:    _GetGoodsStockList_GetGoodsStockByGoodsIDList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock_goods_list.proto",
}