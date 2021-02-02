package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	mgInit "meigo/library/init"
	"meigo/library/log"
	"meigo/library/rateLimit"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/mkevac/debugcharts"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
)

// SourceData 实体
type SourceData struct {
	MainId           int             `gorm:"column:main_id;" json:"main_id" form:"main_id"`
	WxSystemUserId   int             `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" `
	ToolId           int             `gorm:"column:tool_id;" json:"tool_id" form:"tool_id"`
	ToolType         int8            `gorm:"column:tool_type;" json:"tool_type" form:"tool_type"`
	MemberId         int             `gorm:"column:member_id;" json:"member_id" form:"member_id"`
	WxOpenId         string          `gorm:"column:wx_open_id;" json:"wx_open_id" form:"wx_open_id"`
	ClientIp         string          `gorm:"column:client_ip;" json:"client_ip" form:"client_ip"`
	DataGenerationAt int             `gorm:"column:data_generation_at;" json:"data_generation_at" form:"data_generation_at"`
	UserIdentityType int             `gorm:"column:user_identity_type;" json:"user_identity_type" form:"user_identity_type"`
	Type             string          `gorm:"column:type;" json:"type" form:"type"`
	Data             OtherSourceData `gorm:"column:data;" json:"data" form:"data"`
}

// OtherSourceData 实体
type OtherSourceData struct {
	DeviceType  string `gorm:"column:device_type;" json:"device_type" form:"device_type"`
	BrowserType string `gorm:"column:browser_type;" json:"browser_type" form:"browser_type"`
	ChannelId   int    `gorm:"column:channel_id;" json:"channel_id" form:"channel_id"`
}

var ch0, ch1, ch2, ch3, ch4, ch5, ch6, ch7, ch8, ch9 chan string

//var strList = []string{"ch1", "ch2", "ch3", "ch4", "ch5", "ch6", "ch7", "ch8", "ch9"}
var ExeDir string

func main() {

	/*
		timer2 := time.NewTimer(time.Second)
		go func() {
			<-timer2.C
			fmt.Println("Timer 2 expired")
		}()
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("Timer 2 stopped")
		}

		/*
			ticker := time.NewTicker(time.Millisecond * 500)
			go func() {
				for t := range ticker.C {
					fmt.Println("Tick at", t)
				}
			}()
			time.Sleep(time.Second)

	*/
	/*
		for key, _ := range chanList {
			temKey := key
			chanList[temKey] = make(chan string, 5)
			fmt.Println("chanList[temKey]: ", chanList[temKey])
		}
	*/
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	ExeDir = filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	// 配置读取加载
	mgInit.ConfInit(ExeDir)

	//限速设置
	var lr rateLimit.LimitRate
	lr.SetRate(viper.GetInt("const.rateLimiter"))

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

	//初始化channel
	ch0 = make(chan string, 5)
	ch1 = make(chan string, 5)
	ch2 = make(chan string, 5)
	ch3 = make(chan string, 5)
	ch4 = make(chan string, 5)
	ch5 = make(chan string, 5)
	ch6 = make(chan string, 5)
	ch7 = make(chan string, 5)
	ch8 = make(chan string, 5)
	ch9 = make(chan string, 5)
	//chN = make(chan string)
	/*
		ch0 = make(chan string, 5)
		ch1 = make(chan string, 5)
		ch2 = make(chan string, 5)
		ch3 = make(chan string, 5)
		ch4 = make(chan string, 5)
		ch5 = make(chan string, 5)
		ch6 = make(chan string, 5)
		ch7 = make(chan string, 5)
		ch8 = make(chan string, 5)
		ch9 = make(chan string, 5)
		chN = make(chan string, 5)
	*/
	//fmt.Println("redisAddr: ", viper.GetString("redis.addr"))
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})
	var ctx = context.Background()

	//l := rate.NewLimiter(10000, 30000)

	//读取redis数据
	go readRedis(ctx, rdb, lr)

	//阻塞读取channel数据
	for {
		select {
		case val := <-ch0:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch1:
			//fmt.Println("get ch1: ", val)
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch2:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch3:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch4:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch5:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch6:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch7:
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch8:
			//fmt.Println("get ch8: ", val)
			go requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch9:
			go requestOuterApiOnce(val, ctx, rdb)
			/*
				case <-time.After(10 * time.Second):
					fmt.Println("For test env, Time out: ", "100s")
					return

			*/
		}
	}

	//主协程休眠1s，保证调度成功
	//time.Sleep(time.Second)

	fmt.Println("runing: ", "end")
	/*
		defer func() {
			close(ch0)
			close(ch1)
			close(ch2)
			close(ch3)
			close(ch4)
			close(ch5)
			close(ch6)
			close(ch7)
			close(ch8)
			close(ch9)
			close(chN)
		}()
	*/

}

//readRedis 读取redis数据，存入channel
func readRedis(ctx context.Context, rdb *redis.Client, lr rateLimit.LimitRate) {
	for {
		var remainderInt = -1
		var listValueStr = ""
		/*
			n := lr.Limit()
			fmt.Println("nnnn----", n, "----nnnnn")
		*/
		/*
			if l.AllowN(time.Now(), 20000) {
				remainderInt, listValueStr = readRedisOnce(ctx, rdb)
			} else {
				fmt.Println("sleep: ", "500ms")
				time.Sleep(2 * time.Millisecond)
			}
		*/
		if lr.Limit() {
			remainderInt, listValueStr = readRedisOnce(ctx, rdb)
		} else {
			fmt.Println("sleep: ", "50ms")
			time.Sleep(50 * time.Millisecond)
		}

		fmt.Println("remainderInt: ", remainderInt)
		log.Info("remainderInt:", remainderInt)

		if remainderInt == 0 {
			//fmt.Println("ch0: ", ch1)
			ch0 <- listValueStr
		}
		if remainderInt == 1 {
			//fmt.Println("ch1: ", ch1)
			ch1 <- listValueStr
		}
		if remainderInt == 2 {
			//fmt.Println("ch2: ", ch2)
			ch2 <- listValueStr
		}
		if remainderInt == 3 {
			//fmt.Println("ch3: ", ch3)
			ch3 <- listValueStr
		}
		if remainderInt == 4 {
			//fmt.Println("ch4: ", ch4)
			ch4 <- listValueStr
		}
		if remainderInt == 5 {
			//fmt.Println("ch5: ", ch5)
			ch5 <- listValueStr
		}
		if remainderInt == 6 {
			//fmt.Println("ch6: ", ch6)
			ch6 <- listValueStr
		}
		if remainderInt == 7 {
			//fmt.Println("ch7: ", ch7)
			ch7 <- listValueStr
		}
		if remainderInt == 8 {
			//fmt.Println("ch8: ", ch8)
			ch8 <- listValueStr
		}
		if remainderInt == 9 {
			//fmt.Println("ch9: ", ch9)
			ch9 <- listValueStr
		}
		/*
			if remainderInt == -1 {
				chN <- listValueStr
				fmt.Println("chN: ", chN)
			}

		*/
	}
}

//readRedisOnce 读取redis数据
//go 发起http请求 https://www.cnblogs.com/tigerzhouv587/p/11458772.html
//go 发起http请求 https://blog.csdn.net/zangdaiyang1991/article/details/107071529/
func readRedisOnce(ctx context.Context, rdb *redis.Client) (remainder int, listValue string) {
	listValue, err := rdb.RPop(ctx, viper.GetString("redis.source_data_queue")).Result()
	/*
		fmt.Println("listValue: ", listValue)
		log.Info("listValue:", listValue)
	*/
	if err != nil {
		//rdb.RPush(ctx, "list-key", listValue)
		fmt.Println("readRedis-error: ", err)
		log.Error("readRedis-error:", err)
		//panic(err)
	}
	if listValue == "" {
		remainder = -1
		time.Sleep(1 * time.Second)
		return
	}
	//json解析
	jsonStr := []byte(listValue)
	sourceDataObj := SourceData{}
	if err := json.Unmarshal(jsonStr, &sourceDataObj); err != nil {
		fmt.Println("unmarshal err: ", err)
		log.Error("unmarshal err: ", err)
	}
	//求余数
	remainder = sourceDataObj.MainId % 10
	fmt.Println("readRedisData: ", remainder, sourceDataObj)
	//log.Info("readRedisData: ", remainder, sourceDataObj)
	return
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
		client := &http.Client{Timeout: 3 * time.Second}
		resp, err = client.Get(url)
		//请求失败放回错误队列
		//fmt.Println("Response: ", resp.StatusCode, sourceDataJson, resp.Body)
		fmt.Println("tryNum:", tryNum)
		if (err != nil || resp.StatusCode != 200) && tryNum < 2 {
			continue
		} else if (err != nil || resp.StatusCode != 200) && tryNum == 2 {
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
