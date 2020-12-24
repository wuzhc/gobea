// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: ef2331b7e2
// Version Date: 2020-10-07T23:22:38Z

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "article"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	DetailEndpoint     endpoint.Endpoint
	RecordsEndpoint    endpoint.Endpoint
	RemoveEndpoint     endpoint.Endpoint
	TopEndpoint        endpoint.Endpoint
	RecommendsEndpoint endpoint.Endpoint
	ReviewEndpoint     endpoint.Endpoint
	PublishEndpoint    endpoint.Endpoint
	EditEndpoint       endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Detail(ctx context.Context, in *pb.DetailRequest) (*pb.DetailResponse, error) {
	response, err := e.DetailEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.DetailResponse), nil
}

func (e Endpoints) Records(ctx context.Context, in *pb.RecordsRequest) (*pb.RecordsResponse, error) {
	response, err := e.RecordsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RecordsResponse), nil
}

func (e Endpoints) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	response, err := e.RemoveEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RemoveResponse), nil
}

func (e Endpoints) Top(ctx context.Context, in *pb.TopRequest) (*pb.TopResponse, error) {
	response, err := e.TopEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.TopResponse), nil
}

func (e Endpoints) Recommends(ctx context.Context, in *pb.RecommendsRequest) (*pb.RecommendsResponse, error) {
	response, err := e.RecommendsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RecommendsResponse), nil
}

func (e Endpoints) Review(ctx context.Context, in *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	response, err := e.ReviewEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ReviewResponse), nil
}

func (e Endpoints) Publish(ctx context.Context, in *pb.PublishRequest) (*pb.PublishResponse, error) {
	response, err := e.PublishEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PublishResponse), nil
}

func (e Endpoints) Edit(ctx context.Context, in *pb.EditRequest) (*pb.EditResponse, error) {
	response, err := e.EditEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.EditResponse), nil
}

// Make Endpoints

func MakeDetailEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DetailRequest)
		v, err := s.Detail(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeRecordsEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RecordsRequest)
		v, err := s.Records(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeRemoveEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RemoveRequest)
		v, err := s.Remove(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeTopEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.TopRequest)
		v, err := s.Top(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeRecommendsEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RecommendsRequest)
		v, err := s.Recommends(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeReviewEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ReviewRequest)
		v, err := s.Review(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakePublishEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.PublishRequest)
		v, err := s.Publish(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeEditEndpoint(s pb.ArticleServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EditRequest)
		v, err := s.Edit(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Detail":     {},
		"Records":    {},
		"Remove":     {},
		"Top":        {},
		"Recommends": {},
		"Review":     {},
		"Publish":    {},
		"Edit":       {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Detail" {
			e.DetailEndpoint = middleware(e.DetailEndpoint)
		}
		if inc == "Records" {
			e.RecordsEndpoint = middleware(e.RecordsEndpoint)
		}
		if inc == "Remove" {
			e.RemoveEndpoint = middleware(e.RemoveEndpoint)
		}
		if inc == "Top" {
			e.TopEndpoint = middleware(e.TopEndpoint)
		}
		if inc == "Recommends" {
			e.RecommendsEndpoint = middleware(e.RecommendsEndpoint)
		}
		if inc == "Review" {
			e.ReviewEndpoint = middleware(e.ReviewEndpoint)
		}
		if inc == "Publish" {
			e.PublishEndpoint = middleware(e.PublishEndpoint)
		}
		if inc == "Edit" {
			e.EditEndpoint = middleware(e.EditEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Detail":     {},
		"Records":    {},
		"Remove":     {},
		"Top":        {},
		"Recommends": {},
		"Review":     {},
		"Publish":    {},
		"Edit":       {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Detail" {
			e.DetailEndpoint = middleware("Detail", e.DetailEndpoint)
		}
		if inc == "Records" {
			e.RecordsEndpoint = middleware("Records", e.RecordsEndpoint)
		}
		if inc == "Remove" {
			e.RemoveEndpoint = middleware("Remove", e.RemoveEndpoint)
		}
		if inc == "Top" {
			e.TopEndpoint = middleware("Top", e.TopEndpoint)
		}
		if inc == "Recommends" {
			e.RecommendsEndpoint = middleware("Recommends", e.RecommendsEndpoint)
		}
		if inc == "Review" {
			e.ReviewEndpoint = middleware("Review", e.ReviewEndpoint)
		}
		if inc == "Publish" {
			e.PublishEndpoint = middleware("Publish", e.PublishEndpoint)
		}
		if inc == "Edit" {
			e.EditEndpoint = middleware("Edit", e.EditEndpoint)
		}
	}
}
