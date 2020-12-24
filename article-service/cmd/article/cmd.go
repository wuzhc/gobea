package main

import (
	pb "article"
	"article/handlers"
	"article/internal/utils"
	"article/svc"
	"article/svc/server"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
)

func main() {
	utils.InitConfig() //初始化配置
	utils.InitConsul() //初始化注册中心
	utils.InitDB()     //初始化数据库连接

	Run()
}

func Run() {
	httpAddr := viper.GetString("http.addr")
	debugAddr := viper.GetString("debug.addr")
	grpcAddr := viper.GetString("grpc.addr")

	service := handlers.NewService()
	endpoints := server.NewEndpoints(service)

	// Mechanical domain.
	errc := make(chan error)

	// Interrupt handler.
	go handlers.InterruptHandler(errc)

	// Debug listener.
	go func() {
		log.Println("transport", "debug", "addr", debugAddr)

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		errc <- http.ListenAndServe(debugAddr, m)
	}()

	// HTTP transport.
	go func() {
		log.Println("transport", "HTTP", "addr", httpAddr)
		h := svc.MakeHTTPHandler(endpoints)
		r := h.(*mux.Router)
		r.Methods("GET").Path("/check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cotent-type", "application/json")
			w.Write([]byte(`{"status":"ok"}`))
		})

		if err := utils.RegHttpService(); err != nil {
			log.Println(err)
			errc <- err
			return
		}
		errc <- http.ListenAndServe(httpAddr, r)
		if err := utils.UnRegHttpService(); err != nil {
			log.Println(err)
		}
	}()

	// gRPC transport.
	go func() {
		log.Println("transport", "gRPC", "addr", grpcAddr)
		ln, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := svc.MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		pb.RegisterArticleServer(s, srv)
		grpc_health_v1.RegisterHealthServer(s, &server.HealthImpl{Status: grpc_health_v1.HealthCheckResponse_SERVING})

		if err := utils.RegGrpcService(); err != nil {
			log.Println(err)
			errc <- err
			return
		}
		errc <- s.Serve(ln)
		if err := utils.UnRegGrpcService(); err != nil {
			log.Println(err)
		}
	}()

	// Run!
	log.Println("exit", <-errc)
}
