package es

import (
	"meigo/library/log"

	"github.com/spf13/viper"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
)

// ESClient 带配置信息的了解句柄
var ESClient *elastic.Client
var sniff = viper.GetBool("ElasticSearch.sniff")
var healthcheck = viper.GetBool("ElasticSearch.healthcheck")

var cfg = config.Config{
	URL:         viper.GetString("ElasticSearch.url"),
	Index:       viper.GetString("ElasticSearch.index"),
	Username:    viper.GetString("ElasticSearch.username"),
	Password:    viper.GetString("ElasticSearch.password"),
	Shards:      viper.GetInt("ElasticSearch.shards"),
	Replicas:    viper.GetInt("ElasticSearch.replicas"),
	Sniff:       &sniff,
	Healthcheck: &healthcheck,
	Infolog:     viper.GetString("ElasticSearch.infolog"),
	Errorlog:    viper.GetString("ElasticSearch.errorlog"),
	Tracelog:    viper.GetString("ElasticSearch.tracelog"),
}

// ReadESConfig 读取配置文件中ES配置，并新建一个ES连接
func ReadESConfig() {
	esClient, err := elastic.NewClientFromConfig(&cfg)

	if err != nil {
		panic(err)
	}
	log.Info("conn es succ", esClient)
	ESClient = esClient

}
