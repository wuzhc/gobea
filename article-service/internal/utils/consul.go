package utils

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

var consulClient *consulapi.Client

func InitConsul() {
	config := consulapi.DefaultConfig()
	config.Address = viper.GetString("consul.addr")
	client, err := consulapi.NewClient(config)
	if err != nil {
		panic("new consul client failed, " + err.Error())
	}
	consulClient = client
}

//注册服务到consul
func RegService() error {
	if err := RegHttpService(); err != nil {
		return err
	}
	if err := RegGrpcService(); err != nil {
		return err
	}
	return nil
}

//注册http服务
func RegHttpService() error {
	svcname := viper.GetString("service.name")
	port := viper.GetInt("http.port")
	ip := viper.GetString("http.ip")
	addr := viper.GetString("http.addr")
	tags := viper.GetStringSlice("http.tags")

	reg := new(consulapi.AgentServiceRegistration)
	reg.Address = ip
	reg.ID = addr
	reg.Port = port
	reg.Name = svcname
	reg.Tags = tags
	reg.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/check", ip, port),
		Timeout:                        "3s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "30s",
	}
	err := consulClient.Agent().ServiceRegister(reg)
	if err != nil {
		return fmt.Errorf("%s http service register failed, %s", svcname, err)
	} else {
		return nil
	}
}

//注册grpc服务
func RegGrpcService() error {
	svcname := viper.GetString("service.name")
	port := viper.GetInt("grpc.port")
	ip := viper.GetString("grpc.ip")
	addr := viper.GetString("grpc.addr")
	tags := viper.GetStringSlice("grpc.tags")
	reg := new(consulapi.AgentServiceRegistration)
	reg.Address = ip
	reg.ID = addr
	reg.Port = port
	reg.Name = svcname
	reg.Tags = tags
	reg.Check = &consulapi.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d/%s", ip, port, "grpc.health.v1"),
		Timeout:                        "3s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "30s",
	}
	err := consulClient.Agent().ServiceRegister(reg)
	if err != nil {
		return fmt.Errorf("%s grpc service register failed, %s", svcname, err)
	} else {
		return nil
	}
}

//注销服务
func UnReService() {
	_ = UnRegHttpService()
	_ = UnRegGrpcService()
}

//注销http服务
func UnRegHttpService() error {
	addr := viper.GetString("http.addr")
	return consulClient.Agent().ServiceDeregister(addr)
}

//注销grpc服务
func UnRegGrpcService() error {
	addr := viper.GetString("grpc.addr")
	return consulClient.Agent().ServiceDeregister(addr)
}
