// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/client_services.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClientRegistry service

type ClientRegistryClient interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error)
	// List OAuth clients. See request message for details.
	ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error)
	UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error)
	DeleteClient(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type clientRegistryClient struct {
	cc *grpc.ClientConn
}

func NewClientRegistryClient(cc *grpc.ClientConn) ClientRegistryClient {
	return &clientRegistryClient{cc}
}

func (c *clientRegistryClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*Clients, error) {
	out := new(Clients)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/ListClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegistryClient) DeleteClient(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientRegistry/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientRegistry service

type ClientRegistryServer interface {
	// Create a new OAuth client. This also sets the given organization or user as
	// first collaborator with all possible rights.
	CreateClient(context.Context, *CreateClientRequest) (*Client, error)
	// Get the OAuth client with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	GetClient(context.Context, *GetClientRequest) (*Client, error)
	// List OAuth clients. See request message for details.
	ListClients(context.Context, *ListClientsRequest) (*Clients, error)
	UpdateClient(context.Context, *UpdateClientRequest) (*Client, error)
	DeleteClient(context.Context, *ClientIdentifiers) (*types.Empty, error)
}

func RegisterClientRegistryServer(s *grpc.Server, srv ClientRegistryServer) {
	s.RegisterService(&_ClientRegistry_serviceDesc, srv)
}

func _ClientRegistry_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).GetClient(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).ListClients(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).UpdateClient(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegistry_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegistryServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientRegistry/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegistryServer).DeleteClient(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientRegistry",
	HandlerType: (*ClientRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientRegistry_CreateClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _ClientRegistry_GetClient_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _ClientRegistry_ListClients_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ClientRegistry_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientRegistry_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}

// Client API for ClientAccess service

type ClientAccessClient interface {
	ListClientRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Set the rights of a collaborator on the OAuth client. Users or organizations
	// are considered to be a collaborator if they have at least one right on the
	// OAuth client.
	SetClientCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ListClientCollaborators(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Collaborators, error)
}

type clientAccessClient struct {
	cc *grpc.ClientConn
}

func NewClientAccessClient(cc *grpc.ClientConn) ClientAccessClient {
	return &clientAccessClient{cc}
}

func (c *clientAccessClient) ListClientRights(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListClientRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) SetClientCollaborator(ctx context.Context, in *SetClientCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/SetClientCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAccessClient) ListClientCollaborators(ctx context.Context, in *ClientIdentifiers, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ClientAccess/ListClientCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientAccess service

type ClientAccessServer interface {
	ListClientRights(context.Context, *ClientIdentifiers) (*Rights, error)
	// Set the rights of a collaborator on the OAuth client. Users or organizations
	// are considered to be a collaborator if they have at least one right on the
	// OAuth client.
	SetClientCollaborator(context.Context, *SetClientCollaboratorRequest) (*types.Empty, error)
	ListClientCollaborators(context.Context, *ClientIdentifiers) (*Collaborators, error)
}

func RegisterClientAccessServer(s *grpc.Server, srv ClientAccessServer) {
	s.RegisterService(&_ClientAccess_serviceDesc, srv)
}

func _ClientAccess_ListClientRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListClientRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListClientRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListClientRights(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_SetClientCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetClientCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).SetClientCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/SetClientCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).SetClientCollaborator(ctx, req.(*SetClientCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAccess_ListClientCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAccessServer).ListClientCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ClientAccess/ListClientCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAccessServer).ListClientCollaborators(ctx, req.(*ClientIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ClientAccess",
	HandlerType: (*ClientAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClientRights",
			Handler:    _ClientAccess_ListClientRights_Handler,
		},
		{
			MethodName: "SetClientCollaborator",
			Handler:    _ClientAccess_SetClientCollaborator_Handler,
		},
		{
			MethodName: "ListClientCollaborators",
			Handler:    _ClientAccess_ListClientCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/client_services.proto",
}

func init() {
	proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_client_services_386ef5bc17b6333c)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/client_services.proto", fileDescriptor_client_services_386ef5bc17b6333c)
}

var fileDescriptor_client_services_386ef5bc17b6333c = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x48, 0x1c, 0x4d,
	0x14, 0x9f, 0xf9, 0x3e, 0x90, 0xef, 0xdb, 0xef, 0x90, 0x8f, 0x21, 0x51, 0x58, 0xf5, 0x25, 0xae,
	0x01, 0x41, 0x74, 0x36, 0x68, 0x13, 0xd2, 0x25, 0xe6, 0x2f, 0x49, 0x65, 0x48, 0x73, 0x8d, 0xec,
	0xad, 0xe3, 0xde, 0xe0, 0xb9, 0x73, 0xd9, 0x99, 0x53, 0x8c, 0x08, 0x92, 0x22, 0xda, 0x04, 0x02,
	0x29, 0x4c, 0x19, 0x52, 0x04, 0x4b, 0x4b, 0x4b, 0x4b, 0x4b, 0x21, 0x04, 0x24, 0x95, 0xbb, 0x9b,
	0xc2, 0xd2, 0xd2, 0x32, 0xdc, 0xec, 0xee, 0xdd, 0xde, 0xdd, 0x1a, 0x0d, 0xa4, 0x9b, 0x9d, 0xdf,
	0x9b, 0xdf, 0xef, 0xbd, 0xdf, 0x7b, 0x6f, 0x8d, 0xf1, 0x9a, 0x08, 0x9c, 0x55, 0xc7, 0x9f, 0x92,
	0xca, 0x71, 0x97, 0x6c, 0xa7, 0xce, 0x6d, 0xb7, 0xc6, 0x99, 0xaf, 0xe6, 0x25, 0x0b, 0x56, 0xb8,
	0xcb, 0x24, 0xad, 0x07, 0x42, 0x09, 0xd2, 0xaf, 0x94, 0x4f, 0xd3, 0x60, 0xba, 0x32, 0x63, 0x4e,
	0x79, 0x5c, 0x55, 0x1b, 0x15, 0xea, 0x8a, 0x65, 0xdb, 0x13, 0x9e, 0xb0, 0x75, 0x58, 0xa5, 0xb1,
	0xa8, 0xbf, 0xf4, 0x87, 0x3e, 0x25, 0xcf, 0xcd, 0x61, 0x4f, 0x08, 0xaf, 0xc6, 0xb4, 0x80, 0xe3,
	0xfb, 0x42, 0x39, 0x8a, 0x0b, 0x3f, 0x25, 0x37, 0x87, 0x52, 0xb4, 0xc5, 0xc1, 0x96, 0xeb, 0x6a,
	0x2d, 0x05, 0xe1, 0xa2, 0x14, 0x53, 0x7c, 0xac, 0x17, 0xe7, 0x0b, 0xcc, 0x57, 0x7c, 0x91, 0xb3,
	0x40, 0x5e, 0x4c, 0x12, 0x70, 0xaf, 0xaa, 0x52, 0x7c, 0x7a, 0xa7, 0xcf, 0xe8, 0x9f, 0xd5, 0xac,
	0x73, 0xcc, 0xe3, 0x52, 0x05, 0x6b, 0xe4, 0x1b, 0x36, 0x4a, 0xb3, 0x01, 0x73, 0x14, 0x4b, 0x00,
	0x32, 0x46, 0x3b, 0x3d, 0xa0, 0x79, 0x74, 0x8e, 0xbd, 0x6a, 0x30, 0xa9, 0xcc, 0x81, 0x9e, 0x20,
	0x0d, 0x5b, 0x5b, 0xf8, 0xcd, 0xd7, 0x1f, 0x1f, 0xfe, 0xda, 0xc4, 0x16, 0xb5, 0x1b, 0x92, 0x05,
	0xd2, 0x5e, 0x77, 0x45, 0xad, 0xe6, 0x54, 0x44, 0xe0, 0x28, 0x11, 0xd0, 0xe6, 0xdd, 0x3c, 0x5f,
	0x90, 0xd9, 0x61, 0x23, 0x2d, 0x52, 0xde, 0xc5, 0x13, 0xe5, 0x67, 0xd6, 0x23, 0x5b, 0x04, 0x9e,
	0xe3, 0xf3, 0xd7, 0x89, 0x6f, 0x5d, 0x8f, 0xf3, 0x98, 0x26, 0xe9, 0xba, 0xc8, 0x93, 0x11, 0xdf,
	0xf8, 0xf7, 0x31, 0x53, 0x69, 0x4d, 0x37, 0xbb, 0xd3, 0x6d, 0x41, 0x97, 0x15, 0x34, 0xae, 0xeb,
	0x19, 0x25, 0x37, 0x32, 0x6e, 0x7b, 0x3d, 0x9d, 0x9c, 0x66, 0x02, 0xad, 0xe3, 0x06, 0xf9, 0x8e,
	0x8d, 0xff, 0x9e, 0x73, 0x99, 0xd2, 0x4a, 0x62, 0x75, 0x13, 0xe6, 0xc0, 0x4c, 0x74, 0xb0, 0x58,
	0x54, 0x5a, 0xef, 0x12, 0x1b, 0xdf, 0x62, 0xf2, 0x4f, 0xa6, 0x5b, 0xbe, 0x4d, 0x7e, 0xd3, 0xd2,
	0xf2, 0x13, 0xf2, 0x87, 0xfc, 0x24, 0xab, 0x46, 0xe9, 0x65, 0x7d, 0xe1, 0x17, 0x33, 0x92, 0x47,
	0x2f, 0xb3, 0x74, 0x42, 0xd7, 0x76, 0xcb, 0xec, 0xb1, 0x94, 0x76, 0x5a, 0xda, 0xec, 0xa2, 0x67,
	0x94, 0x1e, 0xb0, 0x1a, 0x6b, 0x09, 0x8f, 0x16, 0x73, 0x3e, 0x6d, 0x6f, 0x82, 0x39, 0x40, 0x93,
	0x35, 0xa3, 0xd9, 0x9a, 0xd1, 0x87, 0xcd, 0x35, 0xb3, 0x86, 0xb5, 0xec, 0xc0, 0xc4, 0xb5, 0x82,
	0x4e, 0x6e, 0x4c, 0x7f, 0xf9, 0xdb, 0x28, 0x25, 0x5c, 0xf7, 0x5c, 0x97, 0x49, 0x49, 0x02, 0xe3,
	0xff, 0x76, 0xc7, 0xe6, 0xf4, 0x12, 0x5d, 0x4d, 0xbd, 0x2b, 0x24, 0x79, 0x6a, 0x8d, 0x69, 0xf5,
	0x11, 0x32, 0x54, 0xa4, 0x9e, 0x2e, 0x29, 0xd9, 0xc1, 0xc6, 0xf5, 0x17, 0xd9, 0x64, 0xce, 0xe6,
	0x5a, 0x45, 0x26, 0xbb, 0x69, 0x0b, 0xc3, 0xda, 0xce, 0x17, 0x5b, 0x70, 0x47, 0x27, 0x31, 0x6d,
	0x4e, 0x5d, 0x32, 0xcc, 0x76, 0x7e, 0x4c, 0xf4, 0x36, 0x6d, 0x61, 0x63, 0xb0, 0x6d, 0x47, 0x5e,
	0xf3, 0x4a, 0xae, 0x8c, 0xf4, 0x84, 0xe4, 0x19, 0xb2, 0x89, 0x20, 0x56, 0xa1, 0x39, 0x1d, 0xc9,
	0xdc, 0xff, 0x8c, 0x0f, 0x43, 0xc0, 0x47, 0x21, 0xe0, 0xe3, 0x10, 0xd0, 0x49, 0x08, 0xe8, 0x34,
	0x04, 0x74, 0x16, 0x02, 0x3a, 0x0f, 0x01, 0x6f, 0x46, 0x80, 0xb7, 0x23, 0x40, 0xbb, 0x11, 0xe0,
	0xbd, 0x08, 0xd0, 0x7e, 0x04, 0xe8, 0x20, 0x02, 0x74, 0x18, 0x01, 0x3e, 0x8a, 0x00, 0x1f, 0x47,
	0x80, 0x4e, 0x22, 0xc0, 0xa7, 0x11, 0xa0, 0xb3, 0x08, 0xf0, 0x79, 0x04, 0x68, 0x33, 0x06, 0xb4,
	0x1d, 0x03, 0x7e, 0x1f, 0x03, 0xfa, 0x18, 0x03, 0xfe, 0x14, 0x03, 0xda, 0x8d, 0x01, 0xed, 0xc5,
	0x80, 0xf7, 0x63, 0xc0, 0x07, 0x31, 0xe0, 0xf2, 0xa4, 0x27, 0xa8, 0xaa, 0x32, 0x55, 0xe5, 0xbe,
	0x27, 0xa9, 0xcf, 0xd4, 0xaa, 0x08, 0x96, 0xec, 0xce, 0xdf, 0x6d, 0x7d, 0xc9, 0xb3, 0x95, 0xf2,
	0xeb, 0x95, 0x4a, 0x9f, 0x36, 0x7e, 0xe6, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x26, 0xa2,
	0xae, 0x78, 0x06, 0x00, 0x00,
}