package init

// All the init() function are here.

import (
	conf "meigo/library/config"
	"meigo/library/server"

	"github.com/gin-gonic/gin"
)

// ConfInit func 在main() func 之前加载配置
func ConfInit() {
	conf.LoadConfigFile()
	conf.WatchConf()
	gin.SetMode(server.ServerConf.RunMode)
}
