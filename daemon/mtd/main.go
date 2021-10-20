package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"meigo/library/db"
	"meigo/library/db/common"
	mgInit "meigo/library/init"
	"meigo/library/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/mkevac/debugcharts"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"

	"github.com/jinzhu/gorm"
)

// Node 实体
type Node struct {
	common.BaseModelV1
	ParentId     string `gorm:"column:parent_id;" json:"parent_id" form:"parent_id" binding:"required"`
	FlowId       int    `gorm:"column:flow_id;" json:"flow_id" form:"flow_id" `
	NodeName     string `gorm:"column:node_name;" json:"node_name" form:"node_name"`
	NodeType     string `gorm:"column:node_type;" json:"node_type" form:"node_type"`
	NodeClassify int    `gorm:"column:node_classify;" json:"node_classify" form:"node_classify"`
	Rules        string `gorm:"column:rules;" json:"rules" form:"rules"`
	Styles       string `gorm:"column:styles;" json:"styles" form:"styles"`
	IsRepeat     int    `gorm:"column:is_repeat;" json:"is_repeat" form:"is_repeat"`
	RepeatFreq   string `gorm:"column:repeat_freq;" json:"repeat_freq" form:"repeat_freq"`
	Creator      string `gorm:"column:creator;" json:"creator" form:"creator"`
	Modifier     string `gorm:"column:modifier;" json:"modifier" form:"modifier"`
	IsDel        int    `gorm:"column:is_del;" json:"is_del" form:"is_del"`
}

// FlowExecStates 实体
type FlowExecState struct {
	ID        int    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"` // 主键
	WfUuid    string `gorm:"column:wf_uuid;" json:"wf_uuid" form:"wf_uuid" `
	FlowId    int    `gorm:"column:flow_id;" json:"flow_id" form:"flow_id" `
	NodeId    int    `gorm:"column:node_id;" json:"node_id" form:"node_id" `
	Status    int    `gorm:"column:status;" json:"status" form:"status" `
	Message   string `gorm:"column:message;" json:"message" form:"message" `
	CreatedAt int    `gorm:"column:created_at;" json:"created_at" form:"created_at"` // 创建时间
}

// FlowYaml 实体
type FlowYaml struct {
	common.BaseModelV1
	FlowId      int    `gorm:"column:flow_id;" json:"flow_id" form:"flow_id" `
	NodeId      int    `gorm:"column:node_id;" json:"node_id" form:"node_id" `
	YamlContent string `gorm:"column:yaml_content;" json:"yaml_content" form:"yaml_content"`
	IsDel       int    `gorm:"column:is_del;" json:"is_del" form:"is_del"`
}

// ConditionInfo 实体
type ConditionInfo struct {
	//单位是秒
	Duar   int `gorm:"column:duar;" json:"duar" form:"duar"`
	TimeAt int `gorm:"column:time_at;" json:"time_at" form:"time_at"`
}

// DelayInfo 实体
type DelayInfo struct {
	//单位是秒
	TimerType  string `gorm:"column:timer_type;" json:"timer_type" form:"timer_type"`
	TimerValue int    `gorm:"column:timer_value;" json:"timer_value" form:"timer_value"`
}

// ExecutorInfo 实体
type ExecutorInfo struct {
	//单位是秒
	Duar   int `gorm:"column:duar;" json:"duar" form:"duar"`
	TimeAt int `gorm:"column:time_at;" json:"time_at" form:"time_at"`
}

var ExeDir string
var wf_prefix string

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

	// 初始化数据库连接
	// sqlDB 是 *gorm.DB
	var sqlDB *gorm.DB
	if sqlDB, err = db.ConnDB("test"); err != nil {
		panic(err)
	}

	//读取命令参数
	var message, wfUuid string
	var nodeId int
	//flag.StringVar(&wfId, "wfId", "", "workflow's id")
	flag.StringVar(&message, "message", "", "workflow's input message")
	//传递messageKey，从redis获取值
	//flag.StringVar(&messagekey, "messagekey", "", "workflow's input messagekey")
	flag.StringVar(&wfUuid, "wfUuid", "", "workflow's single exect id")
	flag.IntVar(&nodeId, "nodeId", 0, "workflow's nodeId")
	flag.Parse()
	fmt.Println(wfUuid, nodeId, message)

	//work flow redis prefix
	wf_prefix = viper.GetString("redis.wf_prefix")

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
	if message == "" {
		//除第一个节点外，其他节点的message参数可以为空
		fmt.Println("message is null")
	} else {
		//json解析
		jsonStr := []byte(message)
		if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
			fmt.Println("unmarshal err: ", err)
			log.Error("unmarshal err: ", err)

			//从redis中获取消息内容
			messageResult, err := rdb.Get(ctx, message).Result()
			if err != nil {
				fmt.Println("readRedis-error: ", err)
				log.Error("readRedis-error:", err)
				syscall.Exit(400)
			}
			//json解析
			jsonStr := []byte(messageResult)
			if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
				fmt.Println("unmarshal err: ", err)
				log.Error("unmarshal err: ", err)
				syscall.Exit(400)
			}
		}
	}
	//fmt.Println(messageObj)

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
	ret := run(ctx, rdb, sqlDB, wfUuid, message, nodeId, messageObj)

	//主协程休眠1s，保证调度成功
	//time.Sleep(time.Second)

	if ret == true {
		syscall.Exit(0)
	} else {
		syscall.Exit(200)
	}

	//fmt.Println("runing: ", "end")
}

//run 读取redis数据,执行节点操作
func run(ctx context.Context, rdb *redis.Client, sqlDB *gorm.DB, wfUuid, message string, nodeId int, messageObj map[string]string) bool {
	var inputDataSourceInfo map[string]string
	//获取当前节点信息
	stringValue, err := rdb.Get(ctx, wf_prefix+strconv.Itoa(nodeId)).Result()
	println(wf_prefix+strconv.Itoa(nodeId), stringValue)
	/*
		fmt.Println("listValue: ", listValue)
		log.Info("listValue:", listValue)
	*/
	if err != nil {
		fmt.Println("readRedis-error: ", err)
		log.Error("readRedis-error:", err)
		//panic(err)
		//从数据库获取节点信息，todo

	}
	//json解析
	jsonStr := []byte(stringValue)
	nodeInfoObj := Node{}
	if err := json.Unmarshal(jsonStr, &nodeInfoObj); err != nil {
		fmt.Println("unmarshal err: ", err)
		log.Error("unmarshal err: ", err)
	}
	fmt.Println(nodeInfoObj)
	//处理输入数据源信息
	if len(messageObj) == 0 || nodeInfoObj.ParentId == "" {
		if nodeInfoObj.ParentId == "" {
			//输入信息写入redis hash key
			_, err = rdb.HSet(ctx, wf_prefix+wfUuid, messageObj).Result()
			inputDataSourceInfo = messageObj
		} else {
			//获取数据源信息
			inputDataSourceInfo, err = rdb.HGetAll(ctx, wf_prefix+wfUuid).Result()
			if err != nil {
				fmt.Println("readRedis-error: ", err)
				log.Error("readRedis-error:", err)
				//panic(err)
			}
		}

	} else {
		inputDataSourceInfo = messageObj
	}
	//数据源信息json化
	inputDataSourceInfoStr, _ := json.Marshal(inputDataSourceInfo)
	//解析依赖节点信息
	//if isStringInSlice(nodeInfoObj.nodeType, []string{"condition", "executor"}) {
	if nodeInfoObj.NodeType == "condition" || nodeInfoObj.NodeType == "condition_exclusion" {
		ruleEnginRet := callRuleEngin(nodeInfoObj.Rules, inputDataSourceInfo)
		var status = 0
		if ruleEnginRet == false {
			status = 1
		}
		//更新节点执行状态
		flowExecStateTemp := FlowExecState{WfUuid: wfUuid, FlowId: nodeInfoObj.FlowId, NodeId: nodeId, Status: status, Message: string(inputDataSourceInfoStr), CreatedAt: int(time.Now().Unix())}
		err = sqlDB.Table("flow_exec_states").Create(&flowExecStateTemp).Error
		if err != nil {
			return false
		}
		return ruleEnginRet
	} else if nodeInfoObj.NodeType == "executor" {
		//调用执行服务获取结果，消息中间件
		executorRet := callExecutor(nodeInfoObj.Rules, inputDataSourceInfo)
		//存储数据源信息?,不改变数据源数据
		//如果是同步执行需要更新状态，如果是异步执行，等待状态回调，todo
		//更新执行状态
		var status = 0
		if executorRet == false {
			status = 1
		}
		//更新节点执行状态
		flowExecStateTemp := FlowExecState{WfUuid: wfUuid, FlowId: nodeInfoObj.FlowId, NodeId: nodeId, Status: status, Message: string(inputDataSourceInfoStr), CreatedAt: int(time.Now().Unix())}
		err = sqlDB.Table("flow_exec_states").Create(&flowExecStateTemp).Error
		if err != nil {
			return false
		}
		fmt.Println(9999, nodeInfoObj.IsRepeat, nodeInfoObj.RepeatFreq)
		//重复执行的执行器需要提交到argo cron
		if nodeInfoObj.IsRepeat == 1 {
			flag, cronYaml := generateCronYaml(nodeInfoObj, sqlDB)
			if flag == false {
				return false
			}
			//执行cron工作流
			if ret, _ := doCron(strconv.Itoa(nodeInfoObj.FlowId), cronYaml, message); ret == false {
				return false
			}
		}
		//统一返回
		return executorRet
	} else if nodeInfoObj.NodeType == "delay" {
		fmt.Println(nodeInfoObj.Rules)
		//延迟返回结果
		type Delay struct {
			Constraints DelayInfo ` json:"constraints"`
		}
		var delayObj Delay
		//var delayObj map[string]map[string]string
		//json解码
		err := json.Unmarshal([]byte(nodeInfoObj.Rules), &delayObj)
		if err != nil {
			log.Error("json.Unmarshal-error:", err)
			return false
		}
		fmt.Println(delayObj.Constraints.TimerType, delayObj.Constraints.TimerValue)
		//fmt.Println(delayObj)
		//timerValueInt, _ := strconv.Atoi(delayObj.Constraints.TimerValue)
		//延迟处理
		if delayObj.Constraints.TimerType == "duration" {
			time.Sleep(time.Duration(delayObj.Constraints.TimerValue) * time.Second)
		} else if delayObj.Constraints.TimerType == "fixed_time" {
			//fmt.Println(int64(delayObj.Constraints.TimerValue), time.Now().Unix())
			for int64(delayObj.Constraints.TimerValue) > time.Now().Unix() {
				time.Sleep(1 * time.Second)
			}
		}
		//更新节点执行状态
		flowExecStateTemp := FlowExecState{WfUuid: wfUuid, FlowId: nodeInfoObj.FlowId, NodeId: nodeId, Status: 0, Message: string(inputDataSourceInfoStr), CreatedAt: int(time.Now().Unix())}
		err = sqlDB.Table("flow_exec_states").Create(&flowExecStateTemp).Error
		if err != nil {
			return false
		}
		//不更新数据源
		return true
	}
	return true
}
func generateCronYaml(nodeInfoObj Node, sqlDB *gorm.DB) (bool, string) {
	var err error
	//重复执行的执行器需要提交到argo cron
	var cronTemplate = `
apiVersion: argoproj.io/v1alpha1
kind: CronWorkflow
metadata:
  name: %s
spec:
  schedule: "%s"
  concurrencyPolicy: "Replace"
  startingDeadlineSeconds: 0
  workflowSpec:
    entrypoint: wfServer
    arguments:
      parameters:
        - name: message
          value: '{"test":"hello word"}'
    templates:
    - name: wfServer
      container:
        image: wf:1.0
        command: ["/app/wf-server/bin/wf"]
        args: ["-wfUuid={{workflow.name}}","-nodeId=%s","-message={{workflow.parameters.message}}"]
`
	var cronName = "cron-wf-" + strconv.Itoa(nodeInfoObj.FlowId) + "-" + strconv.Itoa(nodeInfoObj.ID)
	//生成cron yaml文件
	cronYaml := fmt.Sprintf(cronTemplate, cronName, nodeInfoObj.RepeatFreq, strconv.Itoa(nodeInfoObj.ID))
	//cronYaml信息存入数据库

	//存储cron工作流模版
	var flowYaml FlowYaml
	flowYamlTemp := FlowYaml{FlowId: nodeInfoObj.FlowId, NodeId: nodeInfoObj.ID, YamlContent: cronYaml}
	err = sqlDB.Table("flow_yamls").Where("flow_id = ?", nodeInfoObj.FlowId).Where("node_id = ?", nodeInfoObj.ID).Select("* ").Order("id desc").First(&flowYaml).Error //Map查询
	if err == nil && flowYaml.ID > 0 {
		//更新流程内容
		flowYamlTemp.UpdatedAt = int(time.Now().Unix())
		fmt.Println(flowYaml.ID, nodeInfoObj.ID)
		err = sqlDB.Table("flow_yamls").Where("id = ?", flowYaml.ID).Updates(flowYamlTemp).Error
		if err != nil {
			return false, cronYaml
		}
	} else {
		//新建流程内容
		flowYamlTemp.CreatedAt = int(time.Now().Unix())
		err = sqlDB.Table("flow_yamls").Create(&flowYamlTemp).Error
		if err != nil {
			return false, cronYaml
		}
	}
	return true, cronYaml
}

func doCron(FlowId string, cronYaml string, Message string) (flag bool, err error) {
	//构造yaml文件
	//操作文件4种方法，https://studygolang.com/articles/2073
	var randInt = rand.Intn(1000) //生成0-1000之间的随机数
	var fileName = FlowId + "_cron_" + strconv.Itoa(randInt)
	fileName = fileName + ".yaml"
	err = ioutil.WriteFile(fileName, []byte(cronYaml), 0666) //写入文件(字节数组)
	if err != nil {
		return false, err
	}
	//提交执行工作流
	//	cmd := exec.Command("/usr/local/bin/argo submit", fileName, "-n argo", "-p message="+Message)
	cmd := exec.Command("argo", "cron", "create", fileName, "-n", "argo", "-p", "message="+Message)
	_, err = cmd.Output()
	/*
		data, err := cmd.Output()
		fmt.Println(string(data))
	*/
	if err != nil {
		return false, err
	}
	//删除临时文件
	_ = os.Remove(fileName)
	return true, err
}

//调用规则引擎
func callRuleEngin(Rules string, inputDataSourceInfo map[string]string) bool {

	return true
}

//调用执行器
func callExecutor(Rules string, inputDataSourceInfo map[string]string) bool {
	//同步提供返回

	//异步提供回调，更新节点执行状态

	return false
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
