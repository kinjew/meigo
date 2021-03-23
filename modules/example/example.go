package example

import (
	"context"
	"crypto/tls"
	"fmt"
	"meigo/library/log"
	"net"
	"net/http"
	"time"

	"github.com/spf13/viper"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v8"
	ctxExt "github.com/kinjew/gin-context-ext"
	"gopkg.in/go-playground/validator.v8"
)

/*
UploadSingle
curl -X POST http://localhost:8000/example/upload-single \
  -F "file=@/Users/danderui/Desktop/db.php" \
  -H "Content-Type: multipart/form-data"
*/
func UploadSingle(c *ctxExt.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Info(file.Filename)

	// 上传文件至指定目录
	err := c.SaveUploadedFile(file, "/tmp/"+file.Filename)
	if err != nil {
		log.Info("err", err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

/*
UploadMulitple
curl -X POST http://localhost:8000/example/upload-multiple \
  -F "upload[]=@/Users/danderui/Desktop/db.php" \
  -F "upload[]=@/Users/danderui/Desktop/effective-go-zh-en.pdf" \
  -H "Content-Type: multipart/form-data"
*/
func UploadMultiple(c *ctxExt.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Info(file.Filename)

		// 上传文件至指定目录
		if err := c.SaveUploadedFile(file, "/tmp/"+file.Filename); err != nil {
			log.Error("SaveUploadedFile err: ", err)
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

/*
HandleGo
curl -X POST http://localhost:8000/example/upload-single \
  -F "file=@/Users/danderui/Desktop/db.php" \
  -H "Content-Type: multipart/form-data"
当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本
*/
func HandleGo(c *ctxExt.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Info("Done! in path " + cCp.Request.URL.Path)
	}()
}

/*
ValidBookable
curl "localhost:8000/example/valid-bookable?check_in=2019-12-08&check_out=2019-12-09
自定义验证器
*/
func ValidBookable(c *ctxExt.Context) {
	// Booking 包含绑定和验证的数据。
	type Booking struct {
		CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
		CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("bookabledate", bookableDate); err != nil {
			log.Error("RegisterValidation err: ", err)
		}
	}

	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/*
ExampleNewRedisClient https://github.com/go-redis/redis
适用于v6
*/

func Redis(c *ctxExt.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})
	var ctx = context.Background()

	errList := rdb.LPush(ctx, "data_collection_stat_queue", "{\"tool_id\":1,\"tool_type\":1,\"main_id\":5618,\"wx_system_user_id\":36,\"member_id\":20010187,\"wx_open_id\":\"ofK1R1gBLvNqt8Uyn7pp4VVghgC4\",\"client_ip\":\"119.57.93.91\",\"data_generation_at\":1607060185,\"user_identity_type\":1,\"type\":\"visit\",\"data\":{\"device_type\":\"mobile\",\"browser_type\":\"wechat\",\"channel_id\":0}}").Err()
	if errList != nil {
		panic(errList)
	}

	errList1 := rdb.LPush(ctx, "data_collection_stat_queue", "{\"tool_id\":18053,\"tool_type\":1,\"main_id\":5371,\"wx_system_user_id\":36,\"member_id\":20010187,\"wx_open_id\":\"ofK1R1gBLvNqt8Uyn7pp4VVghgC4\",\"client_ip\":\"119.57.93.91\",\"data_generation_at\":1607060185,\"user_identity_type\":1,\"type\":\"visit\",\"data\":{\"device_type\":\"mobile\",\"browser_type\":\"wechat\",\"channel_id\":0}}").Err()
	if errList1 != nil {
		panic(errList1)
	}

	errList2 := rdb.LPush(ctx, "data_collection_stat_queue", "{\"tool_id\":18053,\"tool_type\":1,\"main_id\":5301,\"wx_system_user_id\":36,\"member_id\":20010187,\"wx_open_id\":\"ofK1R1gBLvNqt8Uyn7pp4VVghgC4\",\"client_ip\":\"119.57.93.91\",\"data_generation_at\":1607060185,\"user_identity_type\":1,\"type\":\"visit\",\"data\":{\"device_type\":\"mobile\",\"browser_type\":\"wechat\",\"channel_id\":0}}").Err()
	if errList2 != nil {
		panic(errList2)
	}

	/*
		listValue1, errList_new := rdb.RPop(ctx, "list-key").Result()
		if errList_new != nil {
			panic(errList_new)
			rdb.RPush(ctx, "list-key", "888")

		}
		listValue2, errList_new := rdb.RPop(ctx, "list-key").Result()
		listValue3, errList_new := rdb.RPop(ctx, "list-key").Result()

		fmt.Println("listValue1", listValue1)
		fmt.Println("listValue2", listValue2)
		fmt.Println("listValue3", listValue3)
	*/
	/*
		rdb.RPush(ctx, "list-key", "888")
		rdb.RPush(ctx, "list-key", "999")
		listValue9, errList_new := rdb.RPop(ctx, "list-key").Result()
		fmt.Println("listValue9", listValue9)
		listValue8, errList_new := rdb.RPop(ctx, "list-key").Result()
		fmt.Println("listValue8", listValue8)
	*/
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	c.Success(http.StatusOK, "succ", "test passed")
}

/*
ExampleNewEsClient https://github.com/elastic/go-elasticsearch
*/
func ExampleNewEsClient() {

	// This example demonstrates how to configure the client's Transport.
	//
	// NOTE: These values are for illustrative purposes only, and not suitable
	//       for any production use. The default transport is sufficient.
	//
	cfg := elasticsearch6.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
		//Logger: &estransport.TextLogger{}, // The logger object.
	}

	es, err := elasticsearch6.NewClient(cfg)

	//es, err := elasticsearch6.NewDefaultClient()
	if err != nil {
		log.Error("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Error("Error getting response: %s", err)
	}

	log.Info("res", res)
}
