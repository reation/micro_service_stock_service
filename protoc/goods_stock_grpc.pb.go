// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: goods_stock.proto

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

// StockOpClient is the client API for StockOp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockOpClient interface {
	GetGoodsStockByGoodsID(ctx context.Context, in *GetGoodsStockRequest, opts ...grpc.CallOption) (*GetGoodsStockResponse, error)
}

type stockOpClient struct {
	cc grpc.ClientConnInterface
}

func NewStockOpClient(cc grpc.ClientConnInterface) StockOpClient {
	return &stockOpClient{cc}
}

func (c *stockOpClient) GetGoodsStockByGoodsID(ctx context.Context, in *GetGoodsStockRequest, opts ...grpc.CallOption) (*GetGoodsStockResponse, error) {
	out := new(GetGoodsStockResponse)
	err := c.cc.Invoke(ctx, "/stock_goods.StockOp/GetGoodsStockByGoodsID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockOpServer is the server API for StockOp service.
// All implementations must embed UnimplementedStockOpServer
// for forward compatibility
type StockOpServer interface {
	GetGoodsStockByGoodsID(context.Context, *GetGoodsStockRequest) (*GetGoodsStockResponse, error)
	mustEmbedUnimplementedStockOpServer()
}

// UnimplementedStockOpServer must be embedded to have forward compatible implementations.
type UnimplementedStockOpServer struct {
}

func (UnimplementedStockOpServer) GetGoodsStockByGoodsID(context.Context, *GetGoodsStockRequest) (*GetGoodsStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsStockByGoodsID not implemented")
}
func (UnimplementedStockOpServer) mustEmbedUnimplementedStockOpServer() {}

// UnsafeStockOpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockOpServer will
// result in compilation errors.
type UnsafeStockOpServer interface {
	mustEmbedUnimplementedStockOpServer()
}

func RegisterStockOpServer(s grpc.ServiceRegistrar, srv StockOpServer) {
	s.RegisterService(&StockOp_ServiceDesc, srv)
}

func _StockOp_GetGoodsStockByGoodsID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockOpServer).GetGoodsStockByGoodsID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_goods.StockOp/GetGoodsStockByGoodsID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockOpServer).GetGoodsStockByGoodsID(ctx, req.(*GetGoodsStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockOp_ServiceDesc is the grpc.ServiceDesc for StockOp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockOp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock_goods.StockOp",
	HandlerType: (*StockOpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsStockByGoodsID",
			Handler:    _StockOp_GetGoodsStockByGoodsID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods_stock.proto",
}
