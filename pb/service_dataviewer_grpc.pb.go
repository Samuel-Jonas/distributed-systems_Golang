// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: service_dataviewer.proto

package pb

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

// DataViewerClient is the client API for DataViewer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataViewerClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type dataViewerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataViewerClient(cc grpc.ClientConnInterface) DataViewerClient {
	return &dataViewerClient{cc}
}

func (c *dataViewerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.DataViewer/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataViewerServer is the server API for DataViewer service.
// All implementations must embed UnimplementedDataViewerServer
// for forward compatibility
type DataViewerServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedDataViewerServer()
}

// UnimplementedDataViewerServer must be embedded to have forward compatible implementations.
type UnimplementedDataViewerServer struct {
}

func (UnimplementedDataViewerServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDataViewerServer) mustEmbedUnimplementedDataViewerServer() {}

// UnsafeDataViewerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataViewerServer will
// result in compilation errors.
type UnsafeDataViewerServer interface {
	mustEmbedUnimplementedDataViewerServer()
}

func RegisterDataViewerServer(s grpc.ServiceRegistrar, srv DataViewerServer) {
	s.RegisterService(&DataViewer_ServiceDesc, srv)
}

func _DataViewer_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataViewerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DataViewer/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataViewerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataViewer_ServiceDesc is the grpc.ServiceDesc for DataViewer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataViewer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DataViewer",
	HandlerType: (*DataViewerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DataViewer_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dataviewer.proto",
}
