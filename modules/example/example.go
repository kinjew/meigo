package example

import (
	"crypto/tls"
	"fmt"
	"meigo/library/log"
	"net"
	"net/http"
	"time"

	"github.com/spf13/viper"

	ctxExt "git.sprucetec.com/meigo/gin-context-ext"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis"
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
*/
func Redis(c *ctxExt.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})

	/*
		哨兵模式客户端
		rdb := redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "crm_api",
			SentinelAddrs: []string{"192.168.2.10:26400", "192.168.2.13:26400"},
		})
		pong_sentinel, err_sentinel := rdb.Ping().Result()
		log.Info("pong_sentinel ", pong_sentinel, err_sentinel)
	*/

	pong, err := client.Ping().Result()
	log.Info("pong ", pong, err)
	// Output: PONG <nil>
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		log.Info("err", err)
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		log.Info("err", err)
		panic(err)
	}
	log.Info("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		log.Info("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		log.Info("key2", val2)
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
