package main

import (
	"fmt"
	mgInit "meigo/library/init"
	"meigo/library/log"
	Server "meigo/library/server"
	"meigo/routers"
	"net/http"
	"os"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
)

var ExeDir string

func main() {

	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	ExeDir = filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	// 配置读取加载
	mgInit.ConfInit(ExeDir)

	// 初始化数据库连接
	mgInit.DBInit()
	defer mgInit.DBClose()

	// 初始化路由
	router := routers.InitRouter()

	//监控
	//https://www.cnblogs.com/52fhy/p/11828448.html
	go func() {
		//提供给负载均衡探活
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))

		})

		//prometheus
		http.Handle("/metrics", promhttp.Handler())

		//pprof, go tool pprof -http=:8081 http://$host:$port/debug/pprof/heap
		http.ListenAndServe(":10108", nil)
	}()

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
