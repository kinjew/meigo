package init

// All the init() function are here.

import (
	conf "meigo/library/config"
	"meigo/library/server"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConfInit func 在main() func 之前加载配置
func ConfInit(ExeDir string) {
	MainDir := getParentDirectory(ExeDir)
	//fmt.Println(MainDir)
	conf.LoadConfigFile(MainDir)
	conf.WatchConf()
	gin.SetMode(server.ServerConf.RunMode)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
