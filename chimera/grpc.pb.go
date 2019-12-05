// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package chimera

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xd5, 0x40,
	0x14, 0x86, 0x03, 0xe1, 0xb6, 0xe4, 0x04, 0x05, 0x0e, 0x5c, 0x42, 0x1a, 0x57, 0x5d, 0xba, 0x68,
	0x13, 0xc0, 0xc4, 0x98, 0xb8, 0x80, 0x4b, 0x34, 0x6e, 0x4c, 0xd3, 0xcb, 0x75, 0xe1, 0x6e, 0x68,
	0x0f, 0xed, 0xe4, 0x96, 0x99, 0xb1, 0x33, 0x45, 0xd9, 0xfa, 0x0a, 0xbe, 0x85, 0xef, 0xe0, 0x53,
	0xf8, 0x0a, 0x3e, 0x88, 0x99, 0xce, 0xd0, 0x5c, 0x2f, 0xd2, 0xd5, 0xe4, 0xff, 0xce, 0xf9, 0xff,
	0x49, 0xfe, 0x03, 0x50, 0xb5, 0xaa, 0x48, 0x54, 0x2b, 0x8d, 0xc4, 0x49, 0xff, 0x44, 0x2f, 0x2a,
	0x29, 0xab, 0x86, 0x52, 0xa6, 0x78, 0xca, 0x84, 0x90, 0x86, 0x19, 0x2e, 0x85, 0x76, 0x4b, 0xd1,
	0xf3, 0x5b, 0xd2, 0x9a, 0x55, 0xe4, 0xf5, 0xc9, 0xaf, 0x09, 0x04, 0xb3, 0x86, 0x93, 0x30, 0x78,
	0x09, 0xe1, 0x27, 0x6a, 0x35, 0x97, 0x02, 0xf7, 0xdd, 0x34, 0x39, 0x57, 0xdc, 0xa3, 0xe8, 0x31,
	0x8a, 0x0f, 0xbe, 0xff, 0xfe, 0xf3, 0x63, 0xf3, 0x59, 0xbc, 0x9d, 0xde, 0x39, 0xf2, 0x66, 0xe3,
	0x25, 0x7e, 0x84, 0x70, 0xa1, 0x8a, 0x9a, 0x8a, 0x25, 0x1e, 0x79, 0xcb, 0x42, 0xcd, 0xac, 0xce,
	0x49, 0x2b, 0x29, 0x34, 0x45, 0x4f, 0xf0, 0x95, 0xbc, 0xce, 0x25, 0xd8, 0xbc, 0x14, 0xb6, 0xe6,
	0x24, 0x4a, 0x44, 0x6f, 0xb2, 0x22, 0xa7, 0x2f, 0x1d, 0x69, 0x13, 0x1d, 0xfc, 0xc3, 0x5c, 0x0a,
	0xbe, 0x86, 0x30, 0xa7, 0x82, 0xf8, 0x1d, 0xe1, 0xd4, 0xcf, 0xbd, 0x7e, 0xb0, 0x1d, 0xad, 0x63,
	0xef, 0x3c, 0x83, 0xed, 0xb9, 0x91, 0x2d, 0xe5, 0xec, 0x2b, 0x0e, 0xd1, 0x3d, 0x18, 0xfb, 0xef,
	0x2d, 0xec, 0x58, 0x3d, 0xe7, 0x95, 0xa0, 0xf2, 0xea, 0x1b, 0x1e, 0xaf, 0x2c, 0x39, 0x38, 0x6a,
	0x3f, 0x81, 0x9d, 0xf7, 0x64, 0x32, 0xd6, 0x9a, 0xfb, 0x0f, 0xe2, 0x46, 0xe2, 0x9e, 0x5f, 0x1a,
	0x48, 0xf4, 0x88, 0xe0, 0x15, 0xec, 0x2e, 0x54, 0xc9, 0x0c, 0x8d, 0xd9, 0x8e, 0xd7, 0xc9, 0xd0,
	0xf3, 0xb4, 0xef, 0x79, 0x37, 0x86, 0x54, 0xd9, 0x19, 0x17, 0x37, 0xd2, 0x36, 0xfd, 0x0e, 0xb6,
	0xb2, 0x4e, 0xd7, 0x43, 0xd3, 0x56, 0x64, 0xec, 0xbe, 0x91, 0xac, 0x1c, 0x09, 0xdb, 0xeb, 0xc3,
	0x20, 0x9e, 0xa4, 0xaa, 0xd3, 0xb5, 0xcd, 0x39, 0x83, 0xe0, 0x92, 0x1a, 0x32, 0x84, 0x87, 0xde,
	0xe5, 0xe4, 0x43, 0x0d, 0xff, 0xa5, 0xf8, 0x0a, 0x82, 0x9c, 0xb4, 0xbd, 0xf4, 0xe1, 0x70, 0x1e,
	0xbd, 0x72, 0xeb, 0xe9, 0x1a, 0x75, 0xdf, 0x5f, 0x9c, 0x02, 0x72, 0x99, 0x5c, 0x37, 0xcb, 0xa4,
	0xa8, 0xf9, 0x2d, 0xb5, 0x2c, 0x61, 0x8a, 0x67, 0x1b, 0x9f, 0x43, 0x2f, 0x7f, 0x6e, 0xee, 0x5f,
	0x34, 0x4b, 0x2e, 0x93, 0x99, 0x1f, 0x9f, 0x2b, 0x7e, 0x1d, 0xf4, 0x51, 0xa7, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0x5d, 0x0a, 0x4d, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	// Get version of the server running
	Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error)
	// Checks if the server is running
	Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error)
	// Used to store a transaction in the Chimera node
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Retrieve a transaction from the Chimera node
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
	// Used to store a transaction locally
	StoreRaw(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Used to send signed transaction to the Chimera node
	SendSignedTx(ctx context.Context, in *SendSignedRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Used to get party info from Chimera node
	GetPartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfo, error)
	// Update the partyinfo with response from all the nodes connnected to the node
	UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error)
	// Propogate payload to remote node
	Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteRequest, error)
	Resend(ctx context.Context, in *ResendRequest, opts ...grpc.CallOption) (*ResendResponse, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error) {
	out := new(ApiVersion)
	err := c.cc.Invoke(ctx, "/proto.Client/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error) {
	out := new(UpCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/Upcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) StoreRaw(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/StoreRaw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) SendSignedTx(ctx context.Context, in *SendSignedRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/SendSignedTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) GetPartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfo, error) {
	out := new(PartyInfo)
	err := c.cc.Invoke(ctx, "/proto.Client/GetPartyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/UpdatePartyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteRequest, error) {
	out := new(DeleteRequest)
	err := c.cc.Invoke(ctx, "/proto.Client/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Resend(ctx context.Context, in *ResendRequest, opts ...grpc.CallOption) (*ResendResponse, error) {
	out := new(ResendResponse)
	err := c.cc.Invoke(ctx, "/proto.Client/Resend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	// Get version of the server running
	Version(context.Context, *ApiVersion) (*ApiVersion, error)
	// Checks if the server is running
	Upcheck(context.Context, *UpCheckResponse) (*UpCheckResponse, error)
	// Used to store a transaction in the Chimera node
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Retrieve a transaction from the Chimera node
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	// Used to store a transaction locally
	StoreRaw(context.Context, *StoreRequest) (*SendResponse, error)
	// Used to send signed transaction to the Chimera node
	SendSignedTx(context.Context, *SendSignedRequest) (*SendResponse, error)
	// Used to get party info from Chimera node
	GetPartyInfo(context.Context, *PartyInfo) (*PartyInfo, error)
	// Update the partyinfo with response from all the nodes connnected to the node
	UpdatePartyInfo(context.Context, *PartyInfo) (*PartyInfoResponse, error)
	// Propogate payload to remote node
	Push(context.Context, *PushPayload) (*PartyInfoResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteRequest, error)
	Resend(context.Context, *ResendRequest) (*ResendResponse, error)
}

// UnimplementedClientServer can be embedded to have forward compatible implementations.
type UnimplementedClientServer struct {
}

func (*UnimplementedClientServer) Version(ctx context.Context, req *ApiVersion) (*ApiVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedClientServer) Upcheck(ctx context.Context, req *UpCheckResponse) (*UpCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upcheck not implemented")
}
func (*UnimplementedClientServer) Send(ctx context.Context, req *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedClientServer) Receive(ctx context.Context, req *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (*UnimplementedClientServer) StoreRaw(ctx context.Context, req *StoreRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRaw not implemented")
}
func (*UnimplementedClientServer) SendSignedTx(ctx context.Context, req *SendSignedRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSignedTx not implemented")
}
func (*UnimplementedClientServer) GetPartyInfo(ctx context.Context, req *PartyInfo) (*PartyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyInfo not implemented")
}
func (*UnimplementedClientServer) UpdatePartyInfo(ctx context.Context, req *PartyInfo) (*PartyInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartyInfo not implemented")
}
func (*UnimplementedClientServer) Push(ctx context.Context, req *PushPayload) (*PartyInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedClientServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedClientServer) Resend(ctx context.Context, req *ResendRequest) (*ResendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resend not implemented")
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Version(ctx, req.(*ApiVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Upcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpCheckResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Upcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Upcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Upcheck(ctx, req.(*UpCheckResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_StoreRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).StoreRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/StoreRaw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).StoreRaw(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_SendSignedTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSignedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).SendSignedTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/SendSignedTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).SendSignedTx(ctx, req.(*SendSignedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_GetPartyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).GetPartyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/GetPartyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).GetPartyInfo(ctx, req.(*PartyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UpdatePartyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UpdatePartyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/UpdatePartyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UpdatePartyInfo(ctx, req.(*PartyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Push(ctx, req.(*PushPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Resend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Resend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Client/Resend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Resend(ctx, req.(*ResendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Client_Version_Handler,
		},
		{
			MethodName: "Upcheck",
			Handler:    _Client_Upcheck_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Client_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Client_Receive_Handler,
		},
		{
			MethodName: "StoreRaw",
			Handler:    _Client_StoreRaw_Handler,
		},
		{
			MethodName: "SendSignedTx",
			Handler:    _Client_SendSignedTx_Handler,
		},
		{
			MethodName: "GetPartyInfo",
			Handler:    _Client_GetPartyInfo_Handler,
		},
		{
			MethodName: "UpdatePartyInfo",
			Handler:    _Client_UpdatePartyInfo_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Client_Push_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Client_Delete_Handler,
		},
		{
			MethodName: "Resend",
			Handler:    _Client_Resend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
