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
	"net/http"
	"time"

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

var ch0, ch1, ch2, ch3, ch4, ch5, ch6, ch7, ch8, ch9, chN chan string

var strList = []string{"ch1", "ch2", "ch3", "ch4", "ch5", "ch6", "ch7", "ch8", "ch9"}

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
	// 配置读取加载
	mgInit.ConfInit()

	//初始化channel
	ch0 = make(chan string)
	ch1 = make(chan string)
	ch2 = make(chan string)
	ch3 = make(chan string)
	ch4 = make(chan string)
	ch5 = make(chan string)
	ch6 = make(chan string)
	ch7 = make(chan string)
	ch8 = make(chan string)
	ch9 = make(chan string)
	chN = make(chan string)
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

	//读取redis数据
	go readRedis(ctx, rdb)

	//阻塞读取channel数据
	for {
		select {
		case val := <-ch0:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch1:
			//fmt.Println("get ch1: ", val)
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch2:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch3:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch4:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch5:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch6:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch7:
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch8:
			//fmt.Println("get ch8: ", val)
			requestOuterApiOnce(val, ctx, rdb)
		case val := <-ch9:
			requestOuterApiOnce(val, ctx, rdb)
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

	defer close(ch0)
	defer close(ch1)
	defer close(ch2)
	defer close(ch3)
	defer close(ch4)
	defer close(ch5)
	defer close(ch6)
	defer close(ch7)
	defer close(ch8)
	defer close(ch9)
	defer close(chN)
}

//readRedis 读取redis数据，存入channel
func readRedis(ctx context.Context, rdb *redis.Client) {
	for i := 1; i > 0; i++ {
		fmt.Println("i: ", i)
		log.Info("i:", i)
		remainderInt, listValueStr := readRedisOnce(ctx, rdb)
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
			fmt.Println("ch3: ", ch3)
			ch3 <- listValueStr
		}
		if remainderInt == 4 {
			fmt.Println("ch4: ", ch4)
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
	listValue, err := rdb.RPop(ctx, "list-key").Result()
	fmt.Println("listValue: ", listValue)
	log.Info("listValue:", listValue)
	if err != nil {
		//rdb.RPush(ctx, "list-key", listValue)
		fmt.Println("readRedis-error: ", err)
		log.Info("readRedis-error:", err)
		//panic(err)
	}
	if listValue == "" {
		remainder = -1
		time.Sleep(3 * time.Second)
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
	log.Error("readRedisData: ", remainder, sourceDataObj)
	return
}

//requestOuterApiOnce  请求外部API
func requestOuterApiOnce(sourceDataJson string, ctx context.Context, rdb *redis.Client) {
	fmt.Println("requestOuterApiStart: ", "start")
	log.Error("requestOuterApiStart: ", "start")
	//获取配置数据
	sourceDataProcessUrl := viper.GetString("const.sourceDataProcessUrl")
	url := sourceDataProcessUrl + "?sourceDataJson=" + sourceDataJson
	fmt.Println("requestOuterApiUrl: ", url)
	/*
		url := "http://www.baidu.com" + "?sourceDataJson=" + sourceDataJson
		fmt.Println("requestOuterApiUrl: ", url)

	*/
	//发起get请求
	resp, err := http.Get(url)
	//请求失败返回队列
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("RPush list-key: ", sourceDataJson)
		rdb.RPush(ctx, "list-key", sourceDataJson)
	}
	//defer resp.Body.Close()
	if resp != nil {
		/*
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("responseBody: ", string(body))
		*/
		fmt.Println("requestOuterApiEnd: ", resp.StatusCode, sourceDataJson, resp.Body)
		log.Error("requestOuterApiEnd: ", resp.StatusCode, sourceDataJson, resp.Body)
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
