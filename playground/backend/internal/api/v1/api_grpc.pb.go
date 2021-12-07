// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package playground

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

// PlaygroundServiceClient is the client API for PlaygroundService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaygroundServiceClient interface {
	// Submit the job for an execution and get the pipeline uuid.
	RunCode(ctx context.Context, in *RunCodeRequest, opts ...grpc.CallOption) (*RunCodeResponse, error)
	// Get the status of pipeline execution.
	CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error)
	// Get the result of pipeline execution.
	GetRunOutput(ctx context.Context, in *GetRunOutputRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error)
	// Get the logs of pipeline execution.
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
	// Get the error of pipeline execution.
	GetRunError(ctx context.Context, in *GetRunErrorRequest, opts ...grpc.CallOption) (*GetRunErrorResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(ctx context.Context, in *GetCompileOutputRequest, opts ...grpc.CallOption) (*GetCompileOutputResponse, error)
	// Cancel code processing
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	// Get all precompiled objects from the cloud storage.
	GetPrecompiledObjects(ctx context.Context, in *GetPrecompiledObjectsRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectsResponse, error)
	// Get the code of an PrecompiledObject.
	GetPrecompiledObjectCode(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectCodeResponse, error)
	// Get the precompiled details of an PrecompiledObject.
	GetPrecompiledObjectOutput(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error)
}

type playgroundServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaygroundServiceClient(cc grpc.ClientConnInterface) PlaygroundServiceClient {
	return &playgroundServiceClient{cc}
}

func (c *playgroundServiceClient) RunCode(ctx context.Context, in *RunCodeRequest, opts ...grpc.CallOption) (*RunCodeResponse, error) {
	out := new(RunCodeResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/RunCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetRunOutput(ctx context.Context, in *GetRunOutputRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error) {
	out := new(GetRunOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetRunOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetRunError(ctx context.Context, in *GetRunErrorRequest, opts ...grpc.CallOption) (*GetRunErrorResponse, error) {
	out := new(GetRunErrorResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetRunError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetCompileOutput(ctx context.Context, in *GetCompileOutputRequest, opts ...grpc.CallOption) (*GetCompileOutputResponse, error) {
	out := new(GetCompileOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetCompileOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjects(ctx context.Context, in *GetPrecompiledObjectsRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectsResponse, error) {
	out := new(GetPrecompiledObjectsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectCode(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectCodeResponse, error) {
	out := new(GetPrecompiledObjectCodeResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectOutput(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error) {
	out := new(GetRunOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServiceServer is the server API for PlaygroundService service.
// All implementations should embed UnimplementedPlaygroundServiceServer
// for forward compatibility
type PlaygroundServiceServer interface {
	// Submit the job for an execution and get the pipeline uuid.
	RunCode(context.Context, *RunCodeRequest) (*RunCodeResponse, error)
	// Get the status of pipeline execution.
	CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error)
	// Get the result of pipeline execution.
	GetRunOutput(context.Context, *GetRunOutputRequest) (*GetRunOutputResponse, error)
	// Get the logs of pipeline execution.
	GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	// Get the error of pipeline execution.
	GetRunError(context.Context, *GetRunErrorRequest) (*GetRunErrorResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(context.Context, *GetCompileOutputRequest) (*GetCompileOutputResponse, error)
	// Cancel code processing
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	// Get all precompiled objects from the cloud storage.
	GetPrecompiledObjects(context.Context, *GetPrecompiledObjectsRequest) (*GetPrecompiledObjectsResponse, error)
	// Get the code of an PrecompiledObject.
	GetPrecompiledObjectCode(context.Context, *GetPrecompiledObjectRequest) (*GetPrecompiledObjectCodeResponse, error)
	// Get the precompiled details of an PrecompiledObject.
	GetPrecompiledObjectOutput(context.Context, *GetPrecompiledObjectRequest) (*GetRunOutputResponse, error)
}

// UnimplementedPlaygroundServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServiceServer struct {
}

func (UnimplementedPlaygroundServiceServer) RunCode(context.Context, *RunCodeRequest) (*RunCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCode not implemented")
}
func (UnimplementedPlaygroundServiceServer) CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetRunOutput(context.Context, *GetRunOutputRequest) (*GetRunOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetRunError(context.Context, *GetRunErrorRequest) (*GetRunErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunError not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetCompileOutput(context.Context, *GetCompileOutputRequest) (*GetCompileOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompileOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjects(context.Context, *GetPrecompiledObjectsRequest) (*GetPrecompiledObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjects not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectCode(context.Context, *GetPrecompiledObjectRequest) (*GetPrecompiledObjectCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectCode not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectOutput(context.Context, *GetPrecompiledObjectRequest) (*GetRunOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectOutput not implemented")
}

// UnsafePlaygroundServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaygroundServiceServer will
// result in compilation errors.
type UnsafePlaygroundServiceServer interface {
	mustEmbedUnimplementedPlaygroundServiceServer()
}

func RegisterPlaygroundServiceServer(s grpc.ServiceRegistrar, srv PlaygroundServiceServer) {
	s.RegisterService(&PlaygroundService_ServiceDesc, srv)
}

func _PlaygroundService_RunCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).RunCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/RunCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).RunCode(ctx, req.(*RunCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).CheckStatus(ctx, req.(*CheckStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetRunOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetRunOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetRunOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetRunOutput(ctx, req.(*GetRunOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetLogs(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetRunError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetRunError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetRunError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetRunError(ctx, req.(*GetRunErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetCompileOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompileOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetCompileOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetCompileOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetCompileOutput(ctx, req.(*GetCompileOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjects(ctx, req.(*GetPrecompiledObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObjectCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectCode(ctx, req.(*GetPrecompiledObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObjectOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectOutput(ctx, req.(*GetPrecompiledObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaygroundService_ServiceDesc is the grpc.ServiceDesc for PlaygroundService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaygroundService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.PlaygroundService",
	HandlerType: (*PlaygroundServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCode",
			Handler:    _PlaygroundService_RunCode_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _PlaygroundService_CheckStatus_Handler,
		},
		{
			MethodName: "GetRunOutput",
			Handler:    _PlaygroundService_GetRunOutput_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _PlaygroundService_GetLogs_Handler,
		},
		{
			MethodName: "GetRunError",
			Handler:    _PlaygroundService_GetRunError_Handler,
		},
		{
			MethodName: "GetCompileOutput",
			Handler:    _PlaygroundService_GetCompileOutput_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _PlaygroundService_Cancel_Handler,
		},
		{
			MethodName: "GetPrecompiledObjects",
			Handler:    _PlaygroundService_GetPrecompiledObjects_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectCode",
			Handler:    _PlaygroundService_GetPrecompiledObjectCode_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectOutput",
			Handler:    _PlaygroundService_GetPrecompiledObjectOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}
