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
	return
	/*
		// 平滑启动
		if err := graceup.ListenAndServe(Server.ServerConf.Port, router); err != nil {
			log.Error("err", err)
		}
	*/

}

func maxLenghOfList(list []int) int {
	if len(list) == 1 {
		return 1
	}
	result := 1
	dp := make([]int, len(list))
	dp[0] = 1
	for i := 0; i < len(list); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if list[i] > list[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
