// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: grpc/proto/ghastly.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GhastlyDB_Put_FullMethodName         = "/ghastlydb.GhastlyDB/Put"
	GhastlyDB_Get_FullMethodName         = "/ghastlydb.GhastlyDB/Get"
	GhastlyDB_Delete_FullMethodName      = "/ghastlydb.GhastlyDB/Delete"
	GhastlyDB_Exists_FullMethodName      = "/ghastlydb.GhastlyDB/Exists"
	GhastlyDB_Search_FullMethodName      = "/ghastlydb.GhastlyDB/Search"
	GhastlyDB_BulkPut_FullMethodName     = "/ghastlydb.GhastlyDB/BulkPut"
	GhastlyDB_BulkSearch_FullMethodName  = "/ghastlydb.GhastlyDB/BulkSearch"
	GhastlyDB_HealthCheck_FullMethodName = "/ghastlydb.GhastlyDB/HealthCheck"
	GhastlyDB_GetConfig_FullMethodName   = "/ghastlydb.GhastlyDB/GetConfig"
)

// GhastlyDBClient is the client API for GhastlyDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GhastlyDBClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	BulkPut(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PutRequest, BulkPutResponse], error)
	BulkSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchResponse], error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type ghastlyDBClient struct {
	cc grpc.ClientConnInterface
}

func NewGhastlyDBClient(cc grpc.ClientConnInterface) GhastlyDBClient {
	return &ghastlyDBClient{cc}
}

func (c *ghastlyDBClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_Put_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_Exists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) BulkPut(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PutRequest, BulkPutResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GhastlyDB_ServiceDesc.Streams[0], GhastlyDB_BulkPut_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PutRequest, BulkPutResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GhastlyDB_BulkPutClient = grpc.ClientStreamingClient[PutRequest, BulkPutResponse]

func (c *ghastlyDBClient) BulkSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SearchResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GhastlyDB_ServiceDesc.Streams[1], GhastlyDB_BulkSearch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SearchRequest, SearchResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GhastlyDB_BulkSearchClient = grpc.ServerStreamingClient[SearchResponse]

func (c *ghastlyDBClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ghastlyDBClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, GhastlyDB_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GhastlyDBServer is the server API for GhastlyDB service.
// All implementations must embed UnimplementedGhastlyDBServer
// for forward compatibility.
type GhastlyDBServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	BulkPut(grpc.ClientStreamingServer[PutRequest, BulkPutResponse]) error
	BulkSearch(*SearchRequest, grpc.ServerStreamingServer[SearchResponse]) error
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	mustEmbedUnimplementedGhastlyDBServer()
}

// UnimplementedGhastlyDBServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGhastlyDBServer struct{}

func (UnimplementedGhastlyDBServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedGhastlyDBServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGhastlyDBServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGhastlyDBServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedGhastlyDBServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedGhastlyDBServer) BulkPut(grpc.ClientStreamingServer[PutRequest, BulkPutResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BulkPut not implemented")
}
func (UnimplementedGhastlyDBServer) BulkSearch(*SearchRequest, grpc.ServerStreamingServer[SearchResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BulkSearch not implemented")
}
func (UnimplementedGhastlyDBServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedGhastlyDBServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedGhastlyDBServer) mustEmbedUnimplementedGhastlyDBServer() {}
func (UnimplementedGhastlyDBServer) testEmbeddedByValue()                   {}

// UnsafeGhastlyDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GhastlyDBServer will
// result in compilation errors.
type UnsafeGhastlyDBServer interface {
	mustEmbedUnimplementedGhastlyDBServer()
}

func RegisterGhastlyDBServer(s grpc.ServiceRegistrar, srv GhastlyDBServer) {
	// If the following call pancis, it indicates UnimplementedGhastlyDBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GhastlyDB_ServiceDesc, srv)
}

func _GhastlyDB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_Exists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_BulkPut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GhastlyDBServer).BulkPut(&grpc.GenericServerStream[PutRequest, BulkPutResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GhastlyDB_BulkPutServer = grpc.ClientStreamingServer[PutRequest, BulkPutResponse]

func _GhastlyDB_BulkSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GhastlyDBServer).BulkSearch(m, &grpc.GenericServerStream[SearchRequest, SearchResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GhastlyDB_BulkSearchServer = grpc.ServerStreamingServer[SearchResponse]

func _GhastlyDB_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GhastlyDB_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GhastlyDBServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GhastlyDB_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GhastlyDBServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GhastlyDB_ServiceDesc is the grpc.ServiceDesc for GhastlyDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GhastlyDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ghastlydb.GhastlyDB",
	HandlerType: (*GhastlyDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _GhastlyDB_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GhastlyDB_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GhastlyDB_Delete_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _GhastlyDB_Exists_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _GhastlyDB_Search_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _GhastlyDB_HealthCheck_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _GhastlyDB_GetConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkPut",
			Handler:       _GhastlyDB_BulkPut_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BulkSearch",
			Handler:       _GhastlyDB_BulkSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/proto/ghastly.proto",
}
