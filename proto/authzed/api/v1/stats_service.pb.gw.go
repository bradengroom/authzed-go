// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: authzed/api/v1/stats_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_StatsService_GetRelationshipCardinality_0(ctx context.Context, marshaler runtime.Marshaler, client StatsServiceClient, req *http.Request, pathParams map[string]string) (StatsService_GetRelationshipCardinalityClient, runtime.ServerMetadata, error) {
	var protoReq GetRelationshipCardinalityRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.GetRelationshipCardinality(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_StatsService_UpdateRelationshipCardinality_0(ctx context.Context, marshaler runtime.Marshaler, client StatsServiceClient, req *http.Request, pathParams map[string]string) (StatsService_UpdateRelationshipCardinalityClient, runtime.ServerMetadata, error) {
	var protoReq UpdateRelationshipCardinalityRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.UpdateRelationshipCardinality(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterStatsServiceHandlerServer registers the http handlers for service StatsService to "mux".
// UnaryRPC     :call StatsServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterStatsServiceHandlerFromEndpoint instead.
func RegisterStatsServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server StatsServiceServer) error {

	mux.Handle("POST", pattern_StatsService_GetRelationshipCardinality_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_StatsService_UpdateRelationshipCardinality_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterStatsServiceHandlerFromEndpoint is same as RegisterStatsServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterStatsServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
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

	return RegisterStatsServiceHandler(ctx, mux, conn)
}

// RegisterStatsServiceHandler registers the http handlers for service StatsService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterStatsServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterStatsServiceHandlerClient(ctx, mux, NewStatsServiceClient(conn))
}

// RegisterStatsServiceHandlerClient registers the http handlers for service StatsService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "StatsServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "StatsServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "StatsServiceClient" to call the correct interceptors.
func RegisterStatsServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client StatsServiceClient) error {

	mux.Handle("POST", pattern_StatsService_GetRelationshipCardinality_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/authzed.api.v1.StatsService/GetRelationshipCardinality", runtime.WithHTTPPathPattern("/v1/stats/getrelationshipcardinality"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StatsService_GetRelationshipCardinality_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StatsService_GetRelationshipCardinality_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_StatsService_UpdateRelationshipCardinality_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/authzed.api.v1.StatsService/UpdateRelationshipCardinality", runtime.WithHTTPPathPattern("/v1/stats/updaterelationshipcardinality"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StatsService_UpdateRelationshipCardinality_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StatsService_UpdateRelationshipCardinality_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_StatsService_GetRelationshipCardinality_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "stats", "getrelationshipcardinality"}, ""))

	pattern_StatsService_UpdateRelationshipCardinality_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "stats", "updaterelationshipcardinality"}, ""))
)

var (
	forward_StatsService_GetRelationshipCardinality_0 = runtime.ForwardResponseStream

	forward_StatsService_UpdateRelationshipCardinality_0 = runtime.ForwardResponseStream
)