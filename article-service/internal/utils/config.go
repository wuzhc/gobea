package utils

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic("os.getwd error, " + err.Error())
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("viper set config error, file is not found" + viper.ConfigFileUsed())
		} else {
			panic("viper set config error, " + err.Error())
		}
	}

	//优先级 配置优先级flag>yaml>default
	pflag.String("debug.addr", ":15060", "Debug and metrics listen address")
	pflag.String("http.addr", ":15050", "HTTP listen address")
	pflag.String("grpc.addr", ":15040", "gRPC (HTTP) listen address")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic("viper bind flag error, " + err.Error())
	}
}
