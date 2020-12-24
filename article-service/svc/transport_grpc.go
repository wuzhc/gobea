// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: ef2331b7e2
// Version Date: 2020-10-07T23:22:38Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "article"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC ArticleServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.ArticleServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// article

		detail: grpctransport.NewServer(
			endpoints.DetailEndpoint,
			DecodeGRPCDetailRequest,
			EncodeGRPCDetailResponse,
			serverOptions...,
		),
		records: grpctransport.NewServer(
			endpoints.RecordsEndpoint,
			DecodeGRPCRecordsRequest,
			EncodeGRPCRecordsResponse,
			serverOptions...,
		),
		remove: grpctransport.NewServer(
			endpoints.RemoveEndpoint,
			DecodeGRPCRemoveRequest,
			EncodeGRPCRemoveResponse,
			serverOptions...,
		),
		top: grpctransport.NewServer(
			endpoints.TopEndpoint,
			DecodeGRPCTopRequest,
			EncodeGRPCTopResponse,
			serverOptions...,
		),
		recommends: grpctransport.NewServer(
			endpoints.RecommendsEndpoint,
			DecodeGRPCRecommendsRequest,
			EncodeGRPCRecommendsResponse,
			serverOptions...,
		),
		review: grpctransport.NewServer(
			endpoints.ReviewEndpoint,
			DecodeGRPCReviewRequest,
			EncodeGRPCReviewResponse,
			serverOptions...,
		),
		publish: grpctransport.NewServer(
			endpoints.PublishEndpoint,
			DecodeGRPCPublishRequest,
			EncodeGRPCPublishResponse,
			serverOptions...,
		),
		edit: grpctransport.NewServer(
			endpoints.EditEndpoint,
			DecodeGRPCEditRequest,
			EncodeGRPCEditResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the ArticleServer interface
type grpcServer struct {
	detail     grpctransport.Handler
	records    grpctransport.Handler
	remove     grpctransport.Handler
	top        grpctransport.Handler
	recommends grpctransport.Handler
	review     grpctransport.Handler
	publish    grpctransport.Handler
	edit       grpctransport.Handler
}

// Methods for grpcServer to implement ArticleServer interface

func (s *grpcServer) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailResponse, error) {
	_, rep, err := s.detail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DetailResponse), nil
}

func (s *grpcServer) Records(ctx context.Context, req *pb.RecordsRequest) (*pb.RecordsResponse, error) {
	_, rep, err := s.records.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RecordsResponse), nil
}

func (s *grpcServer) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	_, rep, err := s.remove.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RemoveResponse), nil
}

func (s *grpcServer) Top(ctx context.Context, req *pb.TopRequest) (*pb.TopResponse, error) {
	_, rep, err := s.top.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TopResponse), nil
}

func (s *grpcServer) Recommends(ctx context.Context, req *pb.RecommendsRequest) (*pb.RecommendsResponse, error) {
	_, rep, err := s.recommends.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RecommendsResponse), nil
}

func (s *grpcServer) Review(ctx context.Context, req *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	_, rep, err := s.review.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ReviewResponse), nil
}

func (s *grpcServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	_, rep, err := s.publish.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PublishResponse), nil
}

func (s *grpcServer) Edit(ctx context.Context, req *pb.EditRequest) (*pb.EditResponse, error) {
	_, rep, err := s.edit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.EditResponse), nil
}

// Server Decode

// DecodeGRPCDetailRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC detail request to a user-domain detail request. Primarily useful in a server.
func DecodeGRPCDetailRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DetailRequest)
	return req, nil
}

// DecodeGRPCRecordsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC records request to a user-domain records request. Primarily useful in a server.
func DecodeGRPCRecordsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RecordsRequest)
	return req, nil
}

// DecodeGRPCRemoveRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC remove request to a user-domain remove request. Primarily useful in a server.
func DecodeGRPCRemoveRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RemoveRequest)
	return req, nil
}

// DecodeGRPCTopRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC top request to a user-domain top request. Primarily useful in a server.
func DecodeGRPCTopRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.TopRequest)
	return req, nil
}

// DecodeGRPCRecommendsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC recommends request to a user-domain recommends request. Primarily useful in a server.
func DecodeGRPCRecommendsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RecommendsRequest)
	return req, nil
}

// DecodeGRPCReviewRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC review request to a user-domain review request. Primarily useful in a server.
func DecodeGRPCReviewRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ReviewRequest)
	return req, nil
}

// DecodeGRPCPublishRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC publish request to a user-domain publish request. Primarily useful in a server.
func DecodeGRPCPublishRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PublishRequest)
	return req, nil
}

// DecodeGRPCEditRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC edit request to a user-domain edit request. Primarily useful in a server.
func DecodeGRPCEditRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EditRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCDetailResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain detail response to a gRPC detail reply. Primarily useful in a server.
func EncodeGRPCDetailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.DetailResponse)
	return resp, nil
}

// EncodeGRPCRecordsResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain records response to a gRPC records reply. Primarily useful in a server.
func EncodeGRPCRecordsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RecordsResponse)
	return resp, nil
}

// EncodeGRPCRemoveResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain remove response to a gRPC remove reply. Primarily useful in a server.
func EncodeGRPCRemoveResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RemoveResponse)
	return resp, nil
}

// EncodeGRPCTopResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain top response to a gRPC top reply. Primarily useful in a server.
func EncodeGRPCTopResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.TopResponse)
	return resp, nil
}

// EncodeGRPCRecommendsResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain recommends response to a gRPC recommends reply. Primarily useful in a server.
func EncodeGRPCRecommendsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RecommendsResponse)
	return resp, nil
}

// EncodeGRPCReviewResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain review response to a gRPC review reply. Primarily useful in a server.
func EncodeGRPCReviewResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReviewResponse)
	return resp, nil
}

// EncodeGRPCPublishResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain publish response to a gRPC publish reply. Primarily useful in a server.
func EncodeGRPCPublishResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.PublishResponse)
	return resp, nil
}

// EncodeGRPCEditResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain edit response to a gRPC edit reply. Primarily useful in a server.
func EncodeGRPCEditResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.EditResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
