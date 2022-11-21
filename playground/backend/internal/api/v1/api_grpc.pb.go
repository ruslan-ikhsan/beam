//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/v1/api.proto

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
	// Get the string representation of the pipeline execution graph in DOT format.
	GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error)
	// Get the error of pipeline execution.
	GetRunError(ctx context.Context, in *GetRunErrorRequest, opts ...grpc.CallOption) (*GetRunErrorResponse, error)
	// Get the result of pipeline validation.
	GetValidationOutput(ctx context.Context, in *GetValidationOutputRequest, opts ...grpc.CallOption) (*GetValidationOutputResponse, error)
	// Get the result of pipeline preparation.
	GetPreparationOutput(ctx context.Context, in *GetPreparationOutputRequest, opts ...grpc.CallOption) (*GetPreparationOutputResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(ctx context.Context, in *GetCompileOutputRequest, opts ...grpc.CallOption) (*GetCompileOutputResponse, error)
	// Cancel code processing
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	// Get all precompiled objects from the cloud datastore.
	GetPrecompiledObjects(ctx context.Context, in *GetPrecompiledObjectsRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectsResponse, error)
	// Get precompiled object from the cloud datastore.
	GetPrecompiledObject(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectResponse, error)
	// Get the code of an PrecompiledObject.
	GetPrecompiledObjectCode(ctx context.Context, in *GetPrecompiledObjectCodeRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectCodeResponse, error)
	// Get the precompiled details of an PrecompiledObject.
	GetPrecompiledObjectOutput(ctx context.Context, in *GetPrecompiledObjectOutputRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectOutputResponse, error)
	// Get the logs of an PrecompiledObject.
	GetPrecompiledObjectLogs(ctx context.Context, in *GetPrecompiledObjectLogsRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectLogsResponse, error)
	// Get the graph of an PrecompiledObject.
	GetPrecompiledObjectGraph(ctx context.Context, in *GetPrecompiledObjectGraphRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectGraphResponse, error)
	// Get the default precompile object for the sdk.
	GetDefaultPrecompiledObject(ctx context.Context, in *GetDefaultPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetDefaultPrecompiledObjectResponse, error)
	// Save the snippet required for the sharing.
	SaveSnippet(ctx context.Context, in *SaveSnippetRequest, opts ...grpc.CallOption) (*SaveSnippetResponse, error)
	// Get the snippet of playground.
	GetSnippet(ctx context.Context, in *GetSnippetRequest, opts ...grpc.CallOption) (*GetSnippetResponse, error)
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

func (c *playgroundServiceClient) GetGraph(ctx context.Context, in *GetGraphRequest, opts ...grpc.CallOption) (*GetGraphResponse, error) {
	out := new(GetGraphResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetGraph", in, out, opts...)
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

func (c *playgroundServiceClient) GetValidationOutput(ctx context.Context, in *GetValidationOutputRequest, opts ...grpc.CallOption) (*GetValidationOutputResponse, error) {
	out := new(GetValidationOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetValidationOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPreparationOutput(ctx context.Context, in *GetPreparationOutputRequest, opts ...grpc.CallOption) (*GetPreparationOutputResponse, error) {
	out := new(GetPreparationOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPreparationOutput", in, out, opts...)
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

func (c *playgroundServiceClient) GetPrecompiledObject(ctx context.Context, in *GetPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectResponse, error) {
	out := new(GetPrecompiledObjectResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectCode(ctx context.Context, in *GetPrecompiledObjectCodeRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectCodeResponse, error) {
	out := new(GetPrecompiledObjectCodeResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectOutput(ctx context.Context, in *GetPrecompiledObjectOutputRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectOutputResponse, error) {
	out := new(GetPrecompiledObjectOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectLogs(ctx context.Context, in *GetPrecompiledObjectLogsRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectLogsResponse, error) {
	out := new(GetPrecompiledObjectLogsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetPrecompiledObjectGraph(ctx context.Context, in *GetPrecompiledObjectGraphRequest, opts ...grpc.CallOption) (*GetPrecompiledObjectGraphResponse, error) {
	out := new(GetPrecompiledObjectGraphResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetPrecompiledObjectGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetDefaultPrecompiledObject(ctx context.Context, in *GetDefaultPrecompiledObjectRequest, opts ...grpc.CallOption) (*GetDefaultPrecompiledObjectResponse, error) {
	out := new(GetDefaultPrecompiledObjectResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetDefaultPrecompiledObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) SaveSnippet(ctx context.Context, in *SaveSnippetRequest, opts ...grpc.CallOption) (*SaveSnippetResponse, error) {
	out := new(SaveSnippetResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/SaveSnippet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetSnippet(ctx context.Context, in *GetSnippetRequest, opts ...grpc.CallOption) (*GetSnippetResponse, error) {
	out := new(GetSnippetResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetSnippet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServiceServer is the server API for PlaygroundService service.
// All implementations must embed UnimplementedPlaygroundServiceServer
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
	// Get the string representation of the pipeline execution graph in DOT format.
	GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error)
	// Get the error of pipeline execution.
	GetRunError(context.Context, *GetRunErrorRequest) (*GetRunErrorResponse, error)
	// Get the result of pipeline validation.
	GetValidationOutput(context.Context, *GetValidationOutputRequest) (*GetValidationOutputResponse, error)
	// Get the result of pipeline preparation.
	GetPreparationOutput(context.Context, *GetPreparationOutputRequest) (*GetPreparationOutputResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(context.Context, *GetCompileOutputRequest) (*GetCompileOutputResponse, error)
	// Cancel code processing
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	// Get all precompiled objects from the cloud datastore.
	GetPrecompiledObjects(context.Context, *GetPrecompiledObjectsRequest) (*GetPrecompiledObjectsResponse, error)
	// Get precompiled object from the cloud datastore.
	GetPrecompiledObject(context.Context, *GetPrecompiledObjectRequest) (*GetPrecompiledObjectResponse, error)
	// Get the code of an PrecompiledObject.
	GetPrecompiledObjectCode(context.Context, *GetPrecompiledObjectCodeRequest) (*GetPrecompiledObjectCodeResponse, error)
	// Get the precompiled details of an PrecompiledObject.
	GetPrecompiledObjectOutput(context.Context, *GetPrecompiledObjectOutputRequest) (*GetPrecompiledObjectOutputResponse, error)
	// Get the logs of an PrecompiledObject.
	GetPrecompiledObjectLogs(context.Context, *GetPrecompiledObjectLogsRequest) (*GetPrecompiledObjectLogsResponse, error)
	// Get the graph of an PrecompiledObject.
	GetPrecompiledObjectGraph(context.Context, *GetPrecompiledObjectGraphRequest) (*GetPrecompiledObjectGraphResponse, error)
	// Get the default precompile object for the sdk.
	GetDefaultPrecompiledObject(context.Context, *GetDefaultPrecompiledObjectRequest) (*GetDefaultPrecompiledObjectResponse, error)
	// Save the snippet required for the sharing.
	SaveSnippet(context.Context, *SaveSnippetRequest) (*SaveSnippetResponse, error)
	// Get the snippet of playground.
	GetSnippet(context.Context, *GetSnippetRequest) (*GetSnippetResponse, error)
	mustEmbedUnimplementedPlaygroundServiceServer()
}

// UnimplementedPlaygroundServiceServer must be embedded to have forward compatible implementations.
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
func (UnimplementedPlaygroundServiceServer) GetGraph(context.Context, *GetGraphRequest) (*GetGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetRunError(context.Context, *GetRunErrorRequest) (*GetRunErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunError not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetValidationOutput(context.Context, *GetValidationOutputRequest) (*GetValidationOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidationOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPreparationOutput(context.Context, *GetPreparationOutputRequest) (*GetPreparationOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreparationOutput not implemented")
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
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObject(context.Context, *GetPrecompiledObjectRequest) (*GetPrecompiledObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObject not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectCode(context.Context, *GetPrecompiledObjectCodeRequest) (*GetPrecompiledObjectCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectCode not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectOutput(context.Context, *GetPrecompiledObjectOutputRequest) (*GetPrecompiledObjectOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectLogs(context.Context, *GetPrecompiledObjectLogsRequest) (*GetPrecompiledObjectLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectLogs not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetPrecompiledObjectGraph(context.Context, *GetPrecompiledObjectGraphRequest) (*GetPrecompiledObjectGraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrecompiledObjectGraph not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetDefaultPrecompiledObject(context.Context, *GetDefaultPrecompiledObjectRequest) (*GetDefaultPrecompiledObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultPrecompiledObject not implemented")
}
func (UnimplementedPlaygroundServiceServer) SaveSnippet(context.Context, *SaveSnippetRequest) (*SaveSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSnippet not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetSnippet(context.Context, *GetSnippetRequest) (*GetSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnippet not implemented")
}
func (UnimplementedPlaygroundServiceServer) mustEmbedUnimplementedPlaygroundServiceServer() {}

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

func _PlaygroundService_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetGraph(ctx, req.(*GetGraphRequest))
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

func _PlaygroundService_GetValidationOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidationOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetValidationOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetValidationOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetValidationOutput(ctx, req.(*GetValidationOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPreparationOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreparationOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPreparationOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPreparationOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPreparationOutput(ctx, req.(*GetPreparationOutputRequest))
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

func _PlaygroundService_GetPrecompiledObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObject(ctx, req.(*GetPrecompiledObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectCodeRequest)
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
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectCode(ctx, req.(*GetPrecompiledObjectCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectOutputRequest)
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
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectOutput(ctx, req.(*GetPrecompiledObjectOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObjectLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectLogs(ctx, req.(*GetPrecompiledObjectLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetPrecompiledObjectGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrecompiledObjectGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetPrecompiledObjectGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetPrecompiledObjectGraph(ctx, req.(*GetPrecompiledObjectGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetDefaultPrecompiledObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultPrecompiledObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetDefaultPrecompiledObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetDefaultPrecompiledObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetDefaultPrecompiledObject(ctx, req.(*GetDefaultPrecompiledObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_SaveSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).SaveSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/SaveSnippet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).SaveSnippet(ctx, req.(*SaveSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetSnippet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetSnippet(ctx, req.(*GetSnippetRequest))
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
			MethodName: "GetGraph",
			Handler:    _PlaygroundService_GetGraph_Handler,
		},
		{
			MethodName: "GetRunError",
			Handler:    _PlaygroundService_GetRunError_Handler,
		},
		{
			MethodName: "GetValidationOutput",
			Handler:    _PlaygroundService_GetValidationOutput_Handler,
		},
		{
			MethodName: "GetPreparationOutput",
			Handler:    _PlaygroundService_GetPreparationOutput_Handler,
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
			MethodName: "GetPrecompiledObject",
			Handler:    _PlaygroundService_GetPrecompiledObject_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectCode",
			Handler:    _PlaygroundService_GetPrecompiledObjectCode_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectOutput",
			Handler:    _PlaygroundService_GetPrecompiledObjectOutput_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectLogs",
			Handler:    _PlaygroundService_GetPrecompiledObjectLogs_Handler,
		},
		{
			MethodName: "GetPrecompiledObjectGraph",
			Handler:    _PlaygroundService_GetPrecompiledObjectGraph_Handler,
		},
		{
			MethodName: "GetDefaultPrecompiledObject",
			Handler:    _PlaygroundService_GetDefaultPrecompiledObject_Handler,
		},
		{
			MethodName: "SaveSnippet",
			Handler:    _PlaygroundService_SaveSnippet_Handler,
		},
		{
			MethodName: "GetSnippet",
			Handler:    _PlaygroundService_GetSnippet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}
