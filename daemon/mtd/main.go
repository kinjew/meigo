package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	mgInit "meigo/library/init"
	"meigo/library/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/mkevac/debugcharts"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

// NodeInfo 实体
type NodeInfo struct {
	NodeId        int    `gorm:"column:node_id;" json:"node_id" form:"node_id"`
	ParentNodeIds string `gorm:"column:parent_node_ids;" json:"parent_node_ids" form:"parent_node_ids" `
	NodeName      string `gorm:"column:node_name;" json:"node_name" form:"node_name"`
	NodeType      string `gorm:"column:node_type;" json:"node_type" form:"node_type"`
	IsFirstNode   int    `gorm:"column:is_first_node;" json:"is_first_node" form:"is_first_node"`
	ConditionInfo ConditionInfo
	DelayInfo     DelayInfo
	ExecutorInfo  ExecutorInfo
}

// DelayInfo 实体
type ConditionInfo struct {
	//单位是秒
	Duar   int `gorm:"column:duar;" json:"duar" form:"duar"`
	TimeAt int `gorm:"column:time_at;" json:"time_at" form:"time_at"`
}

// DelayInfo 实体
type DelayInfo struct {
	//单位是秒
	Duar   int `gorm:"column:duar;" json:"duar" form:"duar"`
	TimeAt int `gorm:"column:time_at;" json:"time_at" form:"time_at"`
}

// DelayInfo 实体
type ExecutorInfo struct {
	//单位是秒
	Duar   int `gorm:"column:duar;" json:"duar" form:"duar"`
	TimeAt int `gorm:"column:time_at;" json:"time_at" form:"time_at"`
}

var ExeDir string

/*
./bin/mtd -wfUuid='wew'  -nodeId='234' -message='{"title":"json在线解析（简 版） -JSON在线解析","json.url":"https://www.sojson.com/simple_json.html","keywords":"json在线解析","功能":["JSON美化","JSON数据类型显示","JSON数组显示角标","高亮显示","错误提示",{"备注":["www.sojson.com","json.la"]}],"加入我们":{"qq群":"259217951"}}'
*/

func main() {
	//获取执行目录
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	ExeDir = filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	// 配置读取加载
	mgInit.ConfInit(ExeDir)

	//读取命令参数
	var message, wfUuid, messagekey string
	var nodeId uint
	//flag.StringVar(&wfId, "wfId", "", "workflow's id")
	flag.StringVar(&message, "message", "", "workflow's input message")
	//传递messageKey，从redis获取值
	flag.StringVar(&messagekey, "messagekey", "", "workflow's input messagekey")
	flag.StringVar(&wfUuid, "wfUuid", "", "workflow's single exect id")
	flag.UintVar(&nodeId, "nodeId", 0, "workflow's nodeId")
	flag.Parse()
	fmt.Println(wfUuid, nodeId, message)

	//fmt.Println("redisAddr: ", viper.GetString("redis.addr"))
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})
	var ctx = context.Background()

	var messageObj = make(map[string]string)

	if message != "" {
		//json解析
		jsonStr := []byte(message)
		if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
			fmt.Println("unmarshal err: ", err)
			log.Error("unmarshal err: ", err)
		}
	} else if messagekey != "" {
		//从redis中获取消息内容
		message, err = rdb.Get(ctx, "wf_node_"+messagekey).Result()
		if err != nil {
			fmt.Println("readRedis-error: ", err)
			log.Error("readRedis-error:", err)
			//panic(err)
		}
		//json解析
		jsonStr := []byte(message)
		if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
			fmt.Println("unmarshal err: ", err)
			log.Error("unmarshal err: ", err)
		}
	}
	fmt.Println(messageObj)

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
		http.ListenAndServe(":10109", nil)
	}()

	//执行核心程序
	ret := run(ctx, rdb, wfUuid, nodeId, messageObj)

	//主协程休眠1s，保证调度成功
	//time.Sleep(time.Second)

	if ret == true {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

	//fmt.Println("runing: ", "end")
}

//run 读取redis数据,执行节点操作
func run(ctx context.Context, rdb *redis.Client, wfUuid string, nodeId uint, messageObj map[string]string) bool {
	var inputDataSourceInfo map[string]string
	//获取当前节点信息
	stringValue, err := rdb.Get(ctx, "wf_node_"+string(nodeId)).Result()
	/*
		fmt.Println("listValue: ", listValue)
		log.Info("listValue:", listValue)
	*/
	if err != nil {
		fmt.Println("readRedis-error: ", err)
		log.Error("readRedis-error:", err)
		//panic(err)
	}
	//json解析
	jsonStr := []byte(stringValue)
	nodeInfoObj := NodeInfo{}
	if err := json.Unmarshal(jsonStr, &nodeInfoObj); err != nil {
		fmt.Println("unmarshal err: ", err)
		log.Error("unmarshal err: ", err)
	}
	//处理输入数据源信息
	if nodeInfoObj.IsFirstNode > 0 {
		//输入信息写入redis hash key
		_, err = rdb.HSet(ctx, "wf_node_"+wfUuid, messageObj).Result()
		inputDataSourceInfo = messageObj
	} else {
		//获取数据源信息
		inputDataSourceInfo, err = rdb.HGetAll(ctx, "wf_node_"+wfUuid).Result()
		if err != nil {
			fmt.Println("readRedis-error: ", err)
			log.Error("readRedis-error:", err)
			//panic(err)
		}
	}
	//解析依赖节点信息
	//if isStringInSlice(nodeInfoObj.nodeType, []string{"condition", "executor"}) {
	if nodeInfoObj.NodeType == "condition" {
		ruleEnginRet := callRuleEngin(nodeInfoObj.ConditionInfo, inputDataSourceInfo)
		//存储数据源信息?,不改变数据源数据
		return ruleEnginRet

	} else if nodeInfoObj.NodeType == "executor" {
		//调用执行服务获取结果，消息中间件
		executorRet := callExecutor(nodeInfoObj.ExecutorInfo, inputDataSourceInfo)
		//存储数据源信息?,不改变数据源数据
		return executorRet

	} else if nodeInfoObj.NodeType == "delay" {
		//延迟返回结果
		if nodeInfoObj.DelayInfo.Duar != 0 {
			time.Sleep(time.Duration(nodeInfoObj.DelayInfo.Duar) * time.Second)
		} else {
			for int64(nodeInfoObj.DelayInfo.TimeAt) < time.Now().Unix() {
				time.Sleep(1 * time.Second)
			}
		}
		//不更新数据源
		return true
	}
	return true
}

//调用规则引擎
func callRuleEngin(ConditionInfo ConditionInfo, inputDataSourceInfo map[string]string) bool {

	return true
}

//调用执行器
func callExecutor(ExecutorInfo ExecutorInfo, inputDataSourceInfo map[string]string) bool {

	return true
}

//判断操作符是否在切片中
func isStringInSlice(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

//requestOuterApiOnce  请求外部API
func requestOuterApiOnce(sourceDataJson string, ctx context.Context, rdb *redis.Client) {
	fmt.Println("requestOuterApiStart: ", "start")
	//log.Error("requestOuterApiStart: ", "start")
	//获取配置数据
	sourceDataProcessUrl := viper.GetString("const.sourceDataProcessUrl")
	url := sourceDataProcessUrl + "?source_data=" + sourceDataJson
	fmt.Println("requestOuterApiUrl: ", url)
	/*
		url := "http://www.baidu.com" + "?sourceDataJson=" + sourceDataJson
		fmt.Println("requestOuterApiUrl: ", url)

	*/
	//定义resp
	var resp *http.Response
	var err error
	for tryNum := 0; tryNum < 3; tryNum++ {
		//发起get请求
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err = client.Get(url)
		//请求失败放回错误队列
		//fmt.Println("Response: ", resp.StatusCode, sourceDataJson, resp.Body)
		fmt.Println("tryNum:", tryNum)
		//获取http请求值
		var apiRetObj map[string]interface{}
		if resp != nil {
			apiRetObj = Transformation(resp)
		}
		//获取返回的code
		retCode, _ := apiRetObj["code"]
		//fmt.Println("apiRetObj: ", apiRetObj)
		//fmt.Println("retCode: ", retCode.(float64))
		//		if (err != nil || resp.StatusCode != 200 || retCode != 0) && tryNum < 2 {
		if (retCode.(float64) != 0 || resp == nil || err != nil || resp.StatusCode != 200) && tryNum < 2 {
			fmt.Println("casetryNum: ", tryNum, err, resp.StatusCode, retCode)
			continue
		} else if (retCode.(float64) != 0 || resp == nil || err != nil || resp.StatusCode != 200) && tryNum == 2 {
			//如果不是重跑错误队列，则加入错入队列
			if viper.GetString("redis.source_data_queue") != viper.GetString("redis.source_data_error_queue") {
				fmt.Println("LPush redis.source_data_error_queue: ", sourceDataJson)
				log.Info("LPush: "+viper.GetString("redis.source_data_error_queue"), sourceDataJson)
				rdb.LPush(ctx, viper.GetString("redis.source_data_error_queue"), sourceDataJson)
			}
		} else {
			break
		}
	}
	//fmt.Println("Response22: ", resp.StatusCode, sourceDataJson, resp.Body)
	//defer resp.Body.Close()
	if resp != nil {
		/*
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("responseBody: ", string(body))
		*/
		fmt.Println("requestOuterApiEnd: ", resp.StatusCode, sourceDataJson, resp.Body)
		//log.Info("requestOuterApiEnd: ", resp.StatusCode, sourceDataJson, resp.Body)
		resp.Body.Close()
	} else {
		return
	}
	//return
}

/*
https://github.com/valyala/fasthttp
golang使用fasthttp 发起http请求 https://www.jianshu.com/p/1f546747cb09
*/
/*
备用函数
*/
// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Transformation(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}
	return result
}
