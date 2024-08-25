// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: comm.proto

package safe_remotefscomms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	RemoteFSComms_GetSizeComm_FullMethodName      = "/safe_remotefscomms.RemoteFSComms/GetSizeComm"
	RemoteFSComms_GetLengthComm_FullMethodName    = "/safe_remotefscomms.RemoteFSComms/GetLengthComm"
	RemoteFSComms_GetInfoComm_FullMethodName      = "/safe_remotefscomms.RemoteFSComms/GetInfoComm"
	RemoteFSComms_SetInfoComm_FullMethodName      = "/safe_remotefscomms.RemoteFSComms/SetInfoComm"
	RemoteFSComms_DeleteComm_FullMethodName       = "/safe_remotefscomms.RemoteFSComms/DeleteComm"
	RemoteFSComms_MkDirComm_FullMethodName        = "/safe_remotefscomms.RemoteFSComms/MkDirComm"
	RemoteFSComms_CreateFileComm_FullMethodName   = "/safe_remotefscomms.RemoteFSComms/CreateFileComm"
	RemoteFSComms_RmDirComm_FullMethodName        = "/safe_remotefscomms.RemoteFSComms/RmDirComm"
	RemoteFSComms_SyncFileComm_FullMethodName     = "/safe_remotefscomms.RemoteFSComms/SyncFileComm"
	RemoteFSComms_WriteAtComm_FullMethodName      = "/safe_remotefscomms.RemoteFSComms/WriteAtComm"
	RemoteFSComms_ReadAtComm_FullMethodName       = "/safe_remotefscomms.RemoteFSComms/ReadAtComm"
	RemoteFSComms_RequestReadComm_FullMethodName  = "/safe_remotefscomms.RemoteFSComms/RequestReadComm"
	RemoteFSComms_RequestWriteComm_FullMethodName = "/safe_remotefscomms.RemoteFSComms/RequestWriteComm"
	RemoteFSComms_AckReadComm_FullMethodName      = "/safe_remotefscomms.RemoteFSComms/AckReadComm"
	RemoteFSComms_AckWriteComm_FullMethodName     = "/safe_remotefscomms.RemoteFSComms/AckWriteComm"
)

// RemoteFSCommsClient is the client API for RemoteFSComms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteFSCommsClient interface {
	GetSizeComm(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UintMsg, error)
	GetLengthComm(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UintMsg, error)
	GetInfoComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*FileInfoMsg, error)
	SetInfoComm(ctx context.Context, in *SetInfoParamsMsg, opts ...grpc.CallOption) (*Empty, error)
	DeleteComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error)
	MkDirComm(ctx context.Context, in *MkInodeMsg, opts ...grpc.CallOption) (*UintMsg, error)
	CreateFileComm(ctx context.Context, in *MkInodeMsg, opts ...grpc.CallOption) (*UintMsg, error)
	RmDirComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error)
	SyncFileComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error)
	WriteAtComm(ctx context.Context, in *WriteAtMsg, opts ...grpc.CallOption) (*UintMsg, error)
	ReadAtComm(ctx context.Context, in *ReadAtMsg, opts ...grpc.CallOption) (*ReadAtResponseMsg, error)
	RequestReadComm(ctx context.Context, in *RequestResourceMsg, opts ...grpc.CallOption) (*RequestResponseMsg, error)
	RequestWriteComm(ctx context.Context, in *RequestResourceMsg, opts ...grpc.CallOption) (*RequestResponseMsg, error)
	AckReadComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*AckResponseMsg, error)
	AckWriteComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*AckResponseMsg, error)
}

type remoteFSCommsClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteFSCommsClient(cc grpc.ClientConnInterface) RemoteFSCommsClient {
	return &remoteFSCommsClient{cc}
}

func (c *remoteFSCommsClient) GetSizeComm(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UintMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UintMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_GetSizeComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) GetLengthComm(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UintMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UintMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_GetLengthComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) GetInfoComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*FileInfoMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileInfoMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_GetInfoComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) SetInfoComm(ctx context.Context, in *SetInfoParamsMsg, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RemoteFSComms_SetInfoComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) DeleteComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RemoteFSComms_DeleteComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) MkDirComm(ctx context.Context, in *MkInodeMsg, opts ...grpc.CallOption) (*UintMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UintMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_MkDirComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) CreateFileComm(ctx context.Context, in *MkInodeMsg, opts ...grpc.CallOption) (*UintMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UintMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_CreateFileComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) RmDirComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RemoteFSComms_RmDirComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) SyncFileComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RemoteFSComms_SyncFileComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) WriteAtComm(ctx context.Context, in *WriteAtMsg, opts ...grpc.CallOption) (*UintMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UintMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_WriteAtComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) ReadAtComm(ctx context.Context, in *ReadAtMsg, opts ...grpc.CallOption) (*ReadAtResponseMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadAtResponseMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_ReadAtComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) RequestReadComm(ctx context.Context, in *RequestResourceMsg, opts ...grpc.CallOption) (*RequestResponseMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestResponseMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_RequestReadComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) RequestWriteComm(ctx context.Context, in *RequestResourceMsg, opts ...grpc.CallOption) (*RequestResponseMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestResponseMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_RequestWriteComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) AckReadComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*AckResponseMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AckResponseMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_AckReadComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteFSCommsClient) AckWriteComm(ctx context.Context, in *UintMsg, opts ...grpc.CallOption) (*AckResponseMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AckResponseMsg)
	err := c.cc.Invoke(ctx, RemoteFSComms_AckWriteComm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteFSCommsServer is the server API for RemoteFSComms service.
// All implementations must embed UnimplementedRemoteFSCommsServer
// for forward compatibility
type RemoteFSCommsServer interface {
	GetSizeComm(context.Context, *Empty) (*UintMsg, error)
	GetLengthComm(context.Context, *Empty) (*UintMsg, error)
	GetInfoComm(context.Context, *UintMsg) (*FileInfoMsg, error)
	SetInfoComm(context.Context, *SetInfoParamsMsg) (*Empty, error)
	DeleteComm(context.Context, *UintMsg) (*Empty, error)
	MkDirComm(context.Context, *MkInodeMsg) (*UintMsg, error)
	CreateFileComm(context.Context, *MkInodeMsg) (*UintMsg, error)
	RmDirComm(context.Context, *UintMsg) (*Empty, error)
	SyncFileComm(context.Context, *UintMsg) (*Empty, error)
	WriteAtComm(context.Context, *WriteAtMsg) (*UintMsg, error)
	ReadAtComm(context.Context, *ReadAtMsg) (*ReadAtResponseMsg, error)
	RequestReadComm(context.Context, *RequestResourceMsg) (*RequestResponseMsg, error)
	RequestWriteComm(context.Context, *RequestResourceMsg) (*RequestResponseMsg, error)
	AckReadComm(context.Context, *UintMsg) (*AckResponseMsg, error)
	AckWriteComm(context.Context, *UintMsg) (*AckResponseMsg, error)
	mustEmbedUnimplementedRemoteFSCommsServer()
}

// UnimplementedRemoteFSCommsServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteFSCommsServer struct {
}

func (UnimplementedRemoteFSCommsServer) GetSizeComm(context.Context, *Empty) (*UintMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSizeComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) GetLengthComm(context.Context, *Empty) (*UintMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLengthComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) GetInfoComm(context.Context, *UintMsg) (*FileInfoMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) SetInfoComm(context.Context, *SetInfoParamsMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfoComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) DeleteComm(context.Context, *UintMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) MkDirComm(context.Context, *MkInodeMsg) (*UintMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MkDirComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) CreateFileComm(context.Context, *MkInodeMsg) (*UintMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) RmDirComm(context.Context, *UintMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RmDirComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) SyncFileComm(context.Context, *UintMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncFileComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) WriteAtComm(context.Context, *WriteAtMsg) (*UintMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteAtComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) ReadAtComm(context.Context, *ReadAtMsg) (*ReadAtResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAtComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) RequestReadComm(context.Context, *RequestResourceMsg) (*RequestResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestReadComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) RequestWriteComm(context.Context, *RequestResourceMsg) (*RequestResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestWriteComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) AckReadComm(context.Context, *UintMsg) (*AckResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckReadComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) AckWriteComm(context.Context, *UintMsg) (*AckResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckWriteComm not implemented")
}
func (UnimplementedRemoteFSCommsServer) mustEmbedUnimplementedRemoteFSCommsServer() {}

// UnsafeRemoteFSCommsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteFSCommsServer will
// result in compilation errors.
type UnsafeRemoteFSCommsServer interface {
	mustEmbedUnimplementedRemoteFSCommsServer()
}

func RegisterRemoteFSCommsServer(s grpc.ServiceRegistrar, srv RemoteFSCommsServer) {
	s.RegisterService(&RemoteFSComms_ServiceDesc, srv)
}

func _RemoteFSComms_GetSizeComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).GetSizeComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_GetSizeComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).GetSizeComm(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_GetLengthComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).GetLengthComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_GetLengthComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).GetLengthComm(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_GetInfoComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).GetInfoComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_GetInfoComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).GetInfoComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_SetInfoComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInfoParamsMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).SetInfoComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_SetInfoComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).SetInfoComm(ctx, req.(*SetInfoParamsMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_DeleteComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).DeleteComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_DeleteComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).DeleteComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_MkDirComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkInodeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).MkDirComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_MkDirComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).MkDirComm(ctx, req.(*MkInodeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_CreateFileComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkInodeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).CreateFileComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_CreateFileComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).CreateFileComm(ctx, req.(*MkInodeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_RmDirComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).RmDirComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_RmDirComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).RmDirComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_SyncFileComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).SyncFileComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_SyncFileComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).SyncFileComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_WriteAtComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteAtMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).WriteAtComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_WriteAtComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).WriteAtComm(ctx, req.(*WriteAtMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_ReadAtComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAtMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).ReadAtComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_ReadAtComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).ReadAtComm(ctx, req.(*ReadAtMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_RequestReadComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResourceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).RequestReadComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_RequestReadComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).RequestReadComm(ctx, req.(*RequestResourceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_RequestWriteComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResourceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).RequestWriteComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_RequestWriteComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).RequestWriteComm(ctx, req.(*RequestResourceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_AckReadComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).AckReadComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_AckReadComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).AckReadComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteFSComms_AckWriteComm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UintMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteFSCommsServer).AckWriteComm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteFSComms_AckWriteComm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteFSCommsServer).AckWriteComm(ctx, req.(*UintMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteFSComms_ServiceDesc is the grpc.ServiceDesc for RemoteFSComms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteFSComms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "safe_remotefscomms.RemoteFSComms",
	HandlerType: (*RemoteFSCommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSizeComm",
			Handler:    _RemoteFSComms_GetSizeComm_Handler,
		},
		{
			MethodName: "GetLengthComm",
			Handler:    _RemoteFSComms_GetLengthComm_Handler,
		},
		{
			MethodName: "GetInfoComm",
			Handler:    _RemoteFSComms_GetInfoComm_Handler,
		},
		{
			MethodName: "SetInfoComm",
			Handler:    _RemoteFSComms_SetInfoComm_Handler,
		},
		{
			MethodName: "DeleteComm",
			Handler:    _RemoteFSComms_DeleteComm_Handler,
		},
		{
			MethodName: "MkDirComm",
			Handler:    _RemoteFSComms_MkDirComm_Handler,
		},
		{
			MethodName: "CreateFileComm",
			Handler:    _RemoteFSComms_CreateFileComm_Handler,
		},
		{
			MethodName: "RmDirComm",
			Handler:    _RemoteFSComms_RmDirComm_Handler,
		},
		{
			MethodName: "SyncFileComm",
			Handler:    _RemoteFSComms_SyncFileComm_Handler,
		},
		{
			MethodName: "WriteAtComm",
			Handler:    _RemoteFSComms_WriteAtComm_Handler,
		},
		{
			MethodName: "ReadAtComm",
			Handler:    _RemoteFSComms_ReadAtComm_Handler,
		},
		{
			MethodName: "RequestReadComm",
			Handler:    _RemoteFSComms_RequestReadComm_Handler,
		},
		{
			MethodName: "RequestWriteComm",
			Handler:    _RemoteFSComms_RequestWriteComm_Handler,
		},
		{
			MethodName: "AckReadComm",
			Handler:    _RemoteFSComms_AckReadComm_Handler,
		},
		{
			MethodName: "AckWriteComm",
			Handler:    _RemoteFSComms_AckWriteComm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comm.proto",
}
