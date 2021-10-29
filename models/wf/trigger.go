package wf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"meigo/library/log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
	ctxExt "github.com/kinjew/gin-context-ext"
)

/*
TrigerProcess 触发流程
*/

func TrigerProcess(c *ctxExt.Context) (flag bool, err error) {
	//work flow redis prefix
	//wf_prefix := viper.GetString("redis.wf_prefix")
	//fmt.Println("wf_prefix: ", viper.GetString("redis.wf_prefix"))

	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})
	var ctx = context.Background()

	//flow_id与message不能同时为空
	FlowId := c.Query("flow_id")
	Message := c.Query("message")
	if FlowId == "" || Message == "" {
		log.Error("flow_id or message is not allowed null: ", err)
		return false, fmt.Errorf("flow_id or message is null")
	}
	//flow_id为以逗号分隔的字符串，验证和处理
	FlowIdSlice := strings.Split(FlowId, ",")
	//验证message的数据格式
	var messageObj = make(map[string]string)
	if Message == "" {
		fmt.Println("message is null")
		return false, fmt.Errorf("message is null")
	}
	//json解析
	jsonStr := []byte(Message)
	if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
		fmt.Println("unmarshal err: ", err)
		log.Error("unmarshal err: ", err)

		//从redis中获取消息内容
		messageResult, err := rdb.Get(ctx, Message).Result()
		if err != nil {
			fmt.Println("readRedis-error: ", err)
			log.Error("readRedis-error:", err)
			return false, fmt.Errorf("message's value is null")
		}
		//json解析
		jsonStr := []byte(messageResult)
		if err := json.Unmarshal(jsonStr, &messageObj); err != nil {
			fmt.Println("unmarshal err: ", err)
			log.Error("unmarshal err: ", err)
			return false, fmt.Errorf("message's value is not json")
		}
	}
	//获取节点的yaml内容
	var flowYamls []FlowYaml
	err = sqlDB.Table("flow_yamls").Where("flow_id in (?)", FlowIdSlice).Where("node_id = ?", 0).Select("* ").Scan(&flowYamls).Error
	if err != nil {
		return false, err
	}
	//构造flow_id的map
	/*
		var flowYamlsMap map[int]FlowYaml
		for _, v := range flowYamls {
			flowYamlsMap[v.ID] = v
		}
	*/
	//工作流执行
	for _, item := range flowYamls {
		flag, err := generateYaml(strconv.Itoa(item.FlowId), item, Message)
		if flag == false {
			return false, err
		}
	}
	/*
		argoServerUrl := viper.GetString("const.argoServerUrl")
		fmt.Println("argoServerUrl: ", viper.GetString("const.argoServerUrl"))
		err, ret := Post(argoServerUrl, flowYaml.YamlContent, "application/json")
		if err != nil {
			return false, err
		}
		fmt.Println(ret)
	*/
	// 返回结果
	return true, err
}

func generateYaml(FlowId string, flowYaml FlowYaml, Message string) (flag bool, err error) {
	//构造yaml文件
	//操作文件4种方法，https://studygolang.com/articles/2073
	var randInt = rand.Intn(1000) //生成0-1000之间的随机数
	var fileName = FlowId + "_" + strconv.Itoa(randInt)
	fileName = "/tmp/" + fileName + ".yaml"
	err = ioutil.WriteFile(fileName, []byte(flowYaml.YamlContent), 0666) //写入文件(字节数组)
	if err != nil {
		return false, err
	}
	//提交执行工作流
	//	cmd := exec.Command("/usr/local/bin/argo submit", fileName, "-n argo", "-p message="+Message)
	cmd := exec.Command("argo", "submit", fileName, "-n", "argo", "-p", "message="+Message)
	_, err = cmd.Output()
	/*
		data, err := cmd.Output()
		fmt.Println(string(data))
	*/
	//需要先删除临时文件
	_ = os.Remove(fileName)
	//判断执行结果
	if err != nil {
		return false, err
	}
	return true, err
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) (error, string) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return nil, string(result)
}
