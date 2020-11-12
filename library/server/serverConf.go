package server

import (
	"github.com/spf13/viper"
)

// ServerConf 配置信息实体
var ServerConf Server

// Server 为站点配置参数
type Server struct {
	Domain       string
	Port         string
	RunMode      string //RunMode 服务端运行模式，其中有3种，分别为 "release","test","debug"
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

// ReadServerConfig 加载 web 参数
func ReadServerConfig() {
	viper.New()
	viper.SetDefault("server.runMode", "debug")
	viper.SetDefault("server.port", "8000")
	viper.SetDefault("server.domain", "127.0.0.1")

	ServerConf.Domain = viper.GetString("server.domain")
	ServerConf.Port = viper.GetString("server.port")
	ServerConf.RunMode = viper.GetString("server.runMode")
	ServerConf.ReadTimeout = viper.GetInt("server.readTimeout")
	ServerConf.WriteTimeout = viper.GetInt("server.writeTimeout")
	ServerConf.IdleTimeout = viper.GetInt("server.idleTimeout")

	//fmt.Println("server 配置信息：", ServerConf)

}
