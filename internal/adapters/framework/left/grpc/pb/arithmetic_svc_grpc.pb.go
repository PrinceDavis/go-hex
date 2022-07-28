// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: arithmetic_svc.proto

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

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetSubstraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetAddition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetSubstraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetSubstraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetMultiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations should embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	GetAddition(context.Context, *OperationParameters) (*Answer, error)
	GetSubstraction(context.Context, *OperationParameters) (*Answer, error)
	GetMultiplication(context.Context, *OperationParameters) (*Answer, error)
	GetDivision(context.Context, *OperationParameters) (*Answer, error)
}

// UnimplementedArithmeticServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) GetAddition(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddition not implemented")
}
func (UnimplementedArithmeticServiceServer) GetSubstraction(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubstraction not implemented")
}
func (UnimplementedArithmeticServiceServer) GetMultiplication(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiplication not implemented")
}
func (UnimplementedArithmeticServiceServer) GetDivision(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivision not implemented")
}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_GetAddition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetAddition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetAddition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetAddition(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetSubstraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetSubstraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetSubstraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetSubstraction(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetMultiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetMultiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetMultiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetMultiplication(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetDivision(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddition",
			Handler:    _ArithmeticService_GetAddition_Handler,
		},
		{
			MethodName: "GetSubstraction",
			Handler:    _ArithmeticService_GetSubstraction_Handler,
		},
		{
			MethodName: "GetMultiplication",
			Handler:    _ArithmeticService_GetMultiplication_Handler,
		},
		{
			MethodName: "GetDivision",
			Handler:    _ArithmeticService_GetDivision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_svc.proto",
}
