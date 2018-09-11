// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: lorawan-stack/api/organization_services.proto

/*
Package ttnpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ttnpb

import (
	"io"
	"net/http"

	"context"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_OrganizationRegistry_CreateOrganization_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateOrganizationRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["collaborator.user_ids.user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "collaborator.user_ids.user_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "collaborator.user_ids.user_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "collaborator.user_ids.user_id", err)
	}

	msg, err := client.CreateOrganization(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_OrganizationRegistry_GetOrganization_0 = &utilities.DoubleArray{Encoding: map[string]int{"organization_ids": 0, "organization_id": 1}, Base: []int{1, 1, 1, 0}, Check: []int{0, 1, 2, 3}}
)

func request_OrganizationRegistry_GetOrganization_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetOrganizationRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_ids.organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_ids.organization_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "organization_ids.organization_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_ids.organization_id", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_OrganizationRegistry_GetOrganization_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetOrganization(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_OrganizationRegistry_ListOrganizations_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_OrganizationRegistry_ListOrganizations_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListOrganizationsRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_OrganizationRegistry_ListOrganizations_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListOrganizations(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_OrganizationRegistry_ListOrganizations_1 = &utilities.DoubleArray{Encoding: map[string]int{"collaborator": 0, "user_ids": 1, "user_id": 2}, Base: []int{1, 1, 1, 1, 0}, Check: []int{0, 1, 2, 3, 4}}
)

func request_OrganizationRegistry_ListOrganizations_1(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListOrganizationsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["collaborator.user_ids.user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "collaborator.user_ids.user_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "collaborator.user_ids.user_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "collaborator.user_ids.user_id", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_OrganizationRegistry_ListOrganizations_1); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListOrganizations(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationRegistry_UpdateOrganization_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateOrganizationRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization.ids.organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization.ids.organization_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "organization.ids.organization_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization.ids.organization_id", err)
	}

	msg, err := client.UpdateOrganization(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationRegistry_DeleteOrganization_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationRegistryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrganizationIdentifiers
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_id")
	}

	protoReq.OrganizationID, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_id", err)
	}

	msg, err := client.DeleteOrganization(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_ListOrganizationRights_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrganizationIdentifiers
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_id")
	}

	protoReq.OrganizationID, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_id", err)
	}

	msg, err := client.ListOrganizationRights(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_CreateOrganizationAPIKey_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateOrganizationAPIKeyRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_ids.organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_ids.organization_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "organization_ids.organization_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_ids.organization_id", err)
	}

	msg, err := client.CreateOrganizationAPIKey(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_ListOrganizationAPIKeys_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrganizationIdentifiers
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_id")
	}

	protoReq.OrganizationID, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_id", err)
	}

	msg, err := client.ListOrganizationAPIKeys(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_UpdateOrganizationAPIKey_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateOrganizationAPIKeyRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_ids.organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_ids.organization_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "organization_ids.organization_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_ids.organization_id", err)
	}

	val, ok = pathParams["api_key.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "api_key.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "api_key.id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "api_key.id", err)
	}

	msg, err := client.UpdateOrganizationAPIKey(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_SetOrganizationCollaborator_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SetOrganizationCollaboratorRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_ids.organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_ids.organization_id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "organization_ids.organization_id", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_ids.organization_id", err)
	}

	msg, err := client.SetOrganizationCollaborator(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_OrganizationAccess_ListOrganizationCollaborators_0(ctx context.Context, marshaler runtime.Marshaler, client OrganizationAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OrganizationIdentifiers
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["organization_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "organization_id")
	}

	protoReq.OrganizationID, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "organization_id", err)
	}

	msg, err := client.ListOrganizationCollaborators(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterOrganizationRegistryHandlerFromEndpoint is same as RegisterOrganizationRegistryHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterOrganizationRegistryHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterOrganizationRegistryHandler(ctx, mux, conn)
}

// RegisterOrganizationRegistryHandler registers the http handlers for service OrganizationRegistry to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterOrganizationRegistryHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterOrganizationRegistryHandlerClient(ctx, mux, NewOrganizationRegistryClient(conn))
}

// RegisterOrganizationRegistryHandlerClient registers the http handlers for service OrganizationRegistry
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "OrganizationRegistryClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "OrganizationRegistryClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "OrganizationRegistryClient" to call the correct interceptors.
func RegisterOrganizationRegistryHandlerClient(ctx context.Context, mux *runtime.ServeMux, client OrganizationRegistryClient) error {

	mux.Handle("POST", pattern_OrganizationRegistry_CreateOrganization_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_CreateOrganization_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_CreateOrganization_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_OrganizationRegistry_GetOrganization_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_GetOrganization_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_GetOrganization_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_OrganizationRegistry_ListOrganizations_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_ListOrganizations_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_ListOrganizations_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_OrganizationRegistry_ListOrganizations_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_ListOrganizations_1(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_ListOrganizations_1(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_OrganizationRegistry_UpdateOrganization_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_UpdateOrganization_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_UpdateOrganization_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_OrganizationRegistry_DeleteOrganization_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationRegistry_DeleteOrganization_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationRegistry_DeleteOrganization_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_OrganizationRegistry_CreateOrganization_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"users", "collaborator.user_ids.user_id", "organizations"}, ""))

	pattern_OrganizationRegistry_GetOrganization_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"organizations", "organization_ids.organization_id"}, ""))

	pattern_OrganizationRegistry_ListOrganizations_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"organizations"}, ""))

	pattern_OrganizationRegistry_ListOrganizations_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"users", "collaborator.user_ids.user_id", "organizations"}, ""))

	pattern_OrganizationRegistry_UpdateOrganization_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"organizations", "organization.ids.organization_id"}, ""))

	pattern_OrganizationRegistry_DeleteOrganization_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"organizations", "organization_id"}, ""))
)

var (
	forward_OrganizationRegistry_CreateOrganization_0 = runtime.ForwardResponseMessage

	forward_OrganizationRegistry_GetOrganization_0 = runtime.ForwardResponseMessage

	forward_OrganizationRegistry_ListOrganizations_0 = runtime.ForwardResponseMessage

	forward_OrganizationRegistry_ListOrganizations_1 = runtime.ForwardResponseMessage

	forward_OrganizationRegistry_UpdateOrganization_0 = runtime.ForwardResponseMessage

	forward_OrganizationRegistry_DeleteOrganization_0 = runtime.ForwardResponseMessage
)

// RegisterOrganizationAccessHandlerFromEndpoint is same as RegisterOrganizationAccessHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterOrganizationAccessHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterOrganizationAccessHandler(ctx, mux, conn)
}

// RegisterOrganizationAccessHandler registers the http handlers for service OrganizationAccess to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterOrganizationAccessHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterOrganizationAccessHandlerClient(ctx, mux, NewOrganizationAccessClient(conn))
}

// RegisterOrganizationAccessHandlerClient registers the http handlers for service OrganizationAccess
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "OrganizationAccessClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "OrganizationAccessClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "OrganizationAccessClient" to call the correct interceptors.
func RegisterOrganizationAccessHandlerClient(ctx context.Context, mux *runtime.ServeMux, client OrganizationAccessClient) error {

	mux.Handle("GET", pattern_OrganizationAccess_ListOrganizationRights_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_ListOrganizationRights_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_ListOrganizationRights_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_OrganizationAccess_CreateOrganizationAPIKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_CreateOrganizationAPIKey_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_CreateOrganizationAPIKey_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_OrganizationAccess_ListOrganizationAPIKeys_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_ListOrganizationAPIKeys_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_ListOrganizationAPIKeys_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_OrganizationAccess_UpdateOrganizationAPIKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_UpdateOrganizationAPIKey_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_UpdateOrganizationAPIKey_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_OrganizationAccess_SetOrganizationCollaborator_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_SetOrganizationCollaborator_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_SetOrganizationCollaborator_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_OrganizationAccess_ListOrganizationCollaborators_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OrganizationAccess_ListOrganizationCollaborators_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OrganizationAccess_ListOrganizationCollaborators_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_OrganizationAccess_ListOrganizationRights_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"organizations", "organization_id", "rights"}, ""))

	pattern_OrganizationAccess_CreateOrganizationAPIKey_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"organizations", "organization_ids.organization_id", "api-keys"}, ""))

	pattern_OrganizationAccess_ListOrganizationAPIKeys_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"organizations", "organization_id", "api-keys"}, ""))

	pattern_OrganizationAccess_UpdateOrganizationAPIKey_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"organizations", "organization_ids.organization_id", "api-keys", "api_key.id"}, ""))

	pattern_OrganizationAccess_SetOrganizationCollaborator_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"organizations", "organization_ids.organization_id", "collaborators"}, ""))

	pattern_OrganizationAccess_ListOrganizationCollaborators_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2}, []string{"organizations", "organization_id", "collaborators"}, ""))
)

var (
	forward_OrganizationAccess_ListOrganizationRights_0 = runtime.ForwardResponseMessage

	forward_OrganizationAccess_CreateOrganizationAPIKey_0 = runtime.ForwardResponseMessage

	forward_OrganizationAccess_ListOrganizationAPIKeys_0 = runtime.ForwardResponseMessage

	forward_OrganizationAccess_UpdateOrganizationAPIKey_0 = runtime.ForwardResponseMessage

	forward_OrganizationAccess_SetOrganizationCollaborator_0 = runtime.ForwardResponseMessage

	forward_OrganizationAccess_ListOrganizationCollaborators_0 = runtime.ForwardResponseMessage
)