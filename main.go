package main

import (
	mgInit "meigo/library/init"
	"meigo/library/log"
	Server "meigo/library/server"
	"meigo/routers"
)

func main() {
	// 配置读取加载
	mgInit.ConfInit()

	// 初始化数据库连接
	mgInit.DBInit()
	defer mgInit.DBClose()

	// 初始化路由
	router := routers.InitRouter()

	// 启动服务
	if err := router.Run(Server.ServerConf.Port); err != nil {
		log.Error("err", err)
	}

	/*
		// 平滑启动
		if err := graceup.ListenAndServe(Server.ServerConf.Port, router); err != nil {
			log.Error("err", err)
		}
	*/

}
