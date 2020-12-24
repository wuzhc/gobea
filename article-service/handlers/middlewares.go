package handlers

import (
	pb "article"
	"article/internal/utils"
	"article/svc"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/metaverse/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)
	in.WrapAllExcept(handlePanicErr)
	in.DetailEndpoint = rateLimit("detail", in.DetailEndpoint) //获取详情限流
	return in
}

func WrapService(in pb.ArticleServer) pb.ArticleServer {
	return in
}

//panic错误统一处理
func handlePanicErr(in endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case string:
					err = utils.NewErrMsg(t)
				case error:
					err = utils.NewErrMsg(t.Error())
				default:
					err = utils.NewErrMsg("Unknown error")
				}
				//debug.PrintStack()
				return
			}
		}()
		return in(ctx, request)
	}
}

//服务限流
func rateLimit(name string, in endpoint.Endpoint) endpoint.Endpoint {
	limit := rate.NewLimiter(1, 5)
	fmt.Println(name)
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if !limit.Allow() {
			return nil, fmt.Errorf("too many request")
		}
		return in(ctx, request)
	}
}
