// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package manager

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	// Version returns the version information of the Manager.
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionInfo2, error)
	// ArriveAsClient establishes a session between a client and the Manager.
	ArriveAsClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*SessionInfo, error)
	// ArriveAsAgent establishes a session between an agent and the Manager.
	ArriveAsAgent(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (*SessionInfo, error)
	// Remain indicates that the session is still valid, and potentially
	// updates the auth token for the session.
	Remain(ctx context.Context, in *RemainRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Depart terminates a session.
	Depart(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (*empty.Empty, error)
	// WatchAgents notifies a client of the set of known Agents.
	//
	// A session ID is required; if no session ID is given then the call
	// returns immediately, having not deliverd any snapshots.
	WatchAgents(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (Manager_WatchAgentsClient, error)
	// WatchIntercepts notifies a client or agent of the set of intercepts
	// relevant to that client or agent.
	//
	// If a session ID is given, then only intercepts associated with
	// that session are watched.  If no session ID is given, then all
	// intercepts are watched.
	WatchIntercepts(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (Manager_WatchInterceptsClient, error)
	// CreateIntercept lets a client create an intercept.  It will be
	// created in the "WATING" disposition, and it will remain in that
	// state until until the Agent (the app-sidecar) calls
	// ReviewIntercept() to transition it to the "ACTIVE" disposition
	// (or one of the error dispositions).
	CreateIntercept(ctx context.Context, in *CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptInfo, error)
	// RemoveIntercept lets a client remove an intercept.
	RemoveIntercept(ctx context.Context, in *RemoveInterceptRequest2, opts ...grpc.CallOption) (*empty.Empty, error)
	// ReviewIntercept lets an agent approve or reject an intercept by
	// changing the disposition from "WATING" to "ACTIVE" or to an
	// error, and setting a human-readable status message.
	ReviewIntercept(ctx context.Context, in *ReviewInterceptRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionInfo2, error) {
	out := new(VersionInfo2)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ArriveAsClient(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*SessionInfo, error) {
	out := new(SessionInfo)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/ArriveAsClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ArriveAsAgent(ctx context.Context, in *AgentInfo, opts ...grpc.CallOption) (*SessionInfo, error) {
	out := new(SessionInfo)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/ArriveAsAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Remain(ctx context.Context, in *RemainRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/Remain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Depart(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/Depart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) WatchAgents(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (Manager_WatchAgentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Manager_serviceDesc.Streams[0], "/telepresence.manager.Manager/WatchAgents", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerWatchAgentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_WatchAgentsClient interface {
	Recv() (*AgentInfoSnapshot, error)
	grpc.ClientStream
}

type managerWatchAgentsClient struct {
	grpc.ClientStream
}

func (x *managerWatchAgentsClient) Recv() (*AgentInfoSnapshot, error) {
	m := new(AgentInfoSnapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) WatchIntercepts(ctx context.Context, in *SessionInfo, opts ...grpc.CallOption) (Manager_WatchInterceptsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Manager_serviceDesc.Streams[1], "/telepresence.manager.Manager/WatchIntercepts", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerWatchInterceptsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_WatchInterceptsClient interface {
	Recv() (*InterceptInfoSnapshot, error)
	grpc.ClientStream
}

type managerWatchInterceptsClient struct {
	grpc.ClientStream
}

func (x *managerWatchInterceptsClient) Recv() (*InterceptInfoSnapshot, error) {
	m := new(InterceptInfoSnapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) CreateIntercept(ctx context.Context, in *CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptInfo, error) {
	out := new(InterceptInfo)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/CreateIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) RemoveIntercept(ctx context.Context, in *RemoveInterceptRequest2, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/RemoveIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ReviewIntercept(ctx context.Context, in *ReviewInterceptRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.manager.Manager/ReviewIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	// Version returns the version information of the Manager.
	Version(context.Context, *empty.Empty) (*VersionInfo2, error)
	// ArriveAsClient establishes a session between a client and the Manager.
	ArriveAsClient(context.Context, *ClientInfo) (*SessionInfo, error)
	// ArriveAsAgent establishes a session between an agent and the Manager.
	ArriveAsAgent(context.Context, *AgentInfo) (*SessionInfo, error)
	// Remain indicates that the session is still valid, and potentially
	// updates the auth token for the session.
	Remain(context.Context, *RemainRequest) (*empty.Empty, error)
	// Depart terminates a session.
	Depart(context.Context, *SessionInfo) (*empty.Empty, error)
	// WatchAgents notifies a client of the set of known Agents.
	//
	// A session ID is required; if no session ID is given then the call
	// returns immediately, having not deliverd any snapshots.
	WatchAgents(*SessionInfo, Manager_WatchAgentsServer) error
	// WatchIntercepts notifies a client or agent of the set of intercepts
	// relevant to that client or agent.
	//
	// If a session ID is given, then only intercepts associated with
	// that session are watched.  If no session ID is given, then all
	// intercepts are watched.
	WatchIntercepts(*SessionInfo, Manager_WatchInterceptsServer) error
	// CreateIntercept lets a client create an intercept.  It will be
	// created in the "WATING" disposition, and it will remain in that
	// state until until the Agent (the app-sidecar) calls
	// ReviewIntercept() to transition it to the "ACTIVE" disposition
	// (or one of the error dispositions).
	CreateIntercept(context.Context, *CreateInterceptRequest) (*InterceptInfo, error)
	// RemoveIntercept lets a client remove an intercept.
	RemoveIntercept(context.Context, *RemoveInterceptRequest2) (*empty.Empty, error)
	// ReviewIntercept lets an agent approve or reject an intercept by
	// changing the disposition from "WATING" to "ACTIVE" or to an
	// error, and setting a human-readable status message.
	ReviewIntercept(context.Context, *ReviewInterceptRequest) (*empty.Empty, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) Version(context.Context, *empty.Empty) (*VersionInfo2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedManagerServer) ArriveAsClient(context.Context, *ClientInfo) (*SessionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArriveAsClient not implemented")
}
func (UnimplementedManagerServer) ArriveAsAgent(context.Context, *AgentInfo) (*SessionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArriveAsAgent not implemented")
}
func (UnimplementedManagerServer) Remain(context.Context, *RemainRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remain not implemented")
}
func (UnimplementedManagerServer) Depart(context.Context, *SessionInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Depart not implemented")
}
func (UnimplementedManagerServer) WatchAgents(*SessionInfo, Manager_WatchAgentsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAgents not implemented")
}
func (UnimplementedManagerServer) WatchIntercepts(*SessionInfo, Manager_WatchInterceptsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchIntercepts not implemented")
}
func (UnimplementedManagerServer) CreateIntercept(context.Context, *CreateInterceptRequest) (*InterceptInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIntercept not implemented")
}
func (UnimplementedManagerServer) RemoveIntercept(context.Context, *RemoveInterceptRequest2) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIntercept not implemented")
}
func (UnimplementedManagerServer) ReviewIntercept(context.Context, *ReviewInterceptRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewIntercept not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ArriveAsClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ArriveAsClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/ArriveAsClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ArriveAsClient(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ArriveAsAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ArriveAsAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/ArriveAsAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ArriveAsAgent(ctx, req.(*AgentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Remain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Remain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/Remain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Remain(ctx, req.(*RemainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Depart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Depart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/Depart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Depart(ctx, req.(*SessionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_WatchAgents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SessionInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).WatchAgents(m, &managerWatchAgentsServer{stream})
}

type Manager_WatchAgentsServer interface {
	Send(*AgentInfoSnapshot) error
	grpc.ServerStream
}

type managerWatchAgentsServer struct {
	grpc.ServerStream
}

func (x *managerWatchAgentsServer) Send(m *AgentInfoSnapshot) error {
	return x.ServerStream.SendMsg(m)
}

func _Manager_WatchIntercepts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SessionInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).WatchIntercepts(m, &managerWatchInterceptsServer{stream})
}

type Manager_WatchInterceptsServer interface {
	Send(*InterceptInfoSnapshot) error
	grpc.ServerStream
}

type managerWatchInterceptsServer struct {
	grpc.ServerStream
}

func (x *managerWatchInterceptsServer) Send(m *InterceptInfoSnapshot) error {
	return x.ServerStream.SendMsg(m)
}

func _Manager_CreateIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInterceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/CreateIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateIntercept(ctx, req.(*CreateInterceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_RemoveIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInterceptRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).RemoveIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/RemoveIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).RemoveIntercept(ctx, req.(*RemoveInterceptRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ReviewIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewInterceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ReviewIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager.Manager/ReviewIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ReviewIntercept(ctx, req.(*ReviewInterceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Manager_Version_Handler,
		},
		{
			MethodName: "ArriveAsClient",
			Handler:    _Manager_ArriveAsClient_Handler,
		},
		{
			MethodName: "ArriveAsAgent",
			Handler:    _Manager_ArriveAsAgent_Handler,
		},
		{
			MethodName: "Remain",
			Handler:    _Manager_Remain_Handler,
		},
		{
			MethodName: "Depart",
			Handler:    _Manager_Depart_Handler,
		},
		{
			MethodName: "CreateIntercept",
			Handler:    _Manager_CreateIntercept_Handler,
		},
		{
			MethodName: "RemoveIntercept",
			Handler:    _Manager_RemoveIntercept_Handler,
		},
		{
			MethodName: "ReviewIntercept",
			Handler:    _Manager_ReviewIntercept_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchAgents",
			Handler:       _Manager_WatchAgents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchIntercepts",
			Handler:       _Manager_WatchIntercepts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/manager/manager.proto",
}
