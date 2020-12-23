package config

import (
	database "meigo/library/db"
	Log "meigo/library/log"
	"meigo/library/redis"
	"meigo/library/server"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置参数
type Config struct {
	Server *ServerConf
	Mysql  *MysqlConf
	Redis  *RedisConf
}

// MysqlConf 提供Mysql 模型实体
type MysqlConf struct {
	MysqlModel *database.MySQL
}

// ServerConf 提供 Server 模型实体
type ServerConf struct {
	ServerModel *server.Server
}

// RedisConf 提供 Redis 模型实体
type RedisConf struct {
	RedisModel *redis.RedisConf
}

// LoadConfigFile 读取并加载 conf/conf.yaml(toml) 配置文件
func LoadConfigFile(MainDir string) {
	//从控制台读取标识
	/*
		var cDir string
		flag.StringVar(&cDir, "cDir", "/data/www/go_meigo/conf/", "config directory")
		flag.Parse()

		//conf := viper.New()
		viper.AddConfigPath(cDir)*/
	//viper.AddConfigPath("/data/www/go_meigo/conf/")
	//viper.AddConfigPath("/Users/danderui/test/meigo/conf/")
	//switch runtime.GOOS { case "darwin": case "windows": case "linux": }
	confDir := MainDir + "/conf/"
	//mac本地调试使用
	if runtime.GOOS == "darwin" {
		confDir = "/Users/danderui/test/meigo/conf/"
	}
	viper.AddConfigPath(confDir)
	viper.SetConfigName("conf")
	viper.SetConfigType("toml") // toml 或者 yaml
	// viper.AutomaticEnv()            // 读取匹配的环境变量
	if err := viper.ReadInConfig(); err != nil {
		Log.Error("viper readInConfig() error: ", err)
	}

	ReadConf()

}

// WatchConf 监控配置文件变化并热加载程序
func WatchConf() {
	// 监控配置文件变化并热加载程序
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// viper配置发生变化了 执行响应的操作

		ReadConf()
		Log.Info("Config file changed: %s", e.Name)
	})
}

// ReadConf 加载配置
func ReadConf() {
	// 加载配置
	server.ReadServerConfig()

	//Redis.ReadRedisConfig()
	Log.ZapInit()
}
