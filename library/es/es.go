package es

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"meigo/library/log"

	"gopkg.in/olivere/elastic.v5"

	"github.com/spf13/viper"

	_ "gopkg.in/olivere/elastic.v5"
)

var (
	ctx       = context.Background()
	indexName string
	typeName  string
	esUrl     string
	cspAPI    string
)

// ESInit() 初始化ES 读取相关数据
func ESInit() {
	indexName = viper.GetString("ElasticSearch.index")
	typeName = viper.GetString("ElasticSearch.type")
	esUrl = viper.GetString("ElasticSearch.url")
	cspAPI = viper.GetString("outerAPI.csp-api")
}

// Ping 查看ES版本号，以及ES连接是否正常。
func Ping(esURL string) (string, error) {
	info, code, err := ESClient.Ping(esURL).Do(ctx)
	if err != nil {
		log.Error("ping failed.", err)
		return "", err
	}
	ret := fmt.Sprintf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return ret, nil
}

// GetMapping 获取索引mappings 映射信息
func GetMapping(indexName string) (mapping map[string]interface{}, err error) {
	mapping, err = ESClient.GetMapping().Index(indexName).Pretty(true).Do(ctx)
	if err != nil {
		log.Error("GetMapping failed.", err)
		return mapping, err
	}
	return mapping, err
}

// GetFieldMapping 获取某一个字段的类型
func GetFieldMapping(indexName string, typeName string, fieldName string) (fieldType map[string]interface{}, err error) {
	fieldType, err = ESClient.GetFieldMapping().Index(indexName).Type(typeName).Field(fieldName).Do(ctx)
	if err != nil {
		log.Error("get field mapping err: ", err)
		return fieldType, err
	}
	return fieldType, nil
}

// IsIndexExist 判断索引是否存在。
func IsIndexExist(indexName string) (exist bool, err error) {
	exist, err = ESClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Error("IsIndexExist err: ", err)
		return false, err
	}
	return exist, nil
}

// CreateIndex 创建索引。
func CreateIndex(indexName string, mapping string) (cIndex *elastic.IndicesCreateResult, err error) {
	cIndex, err = ESClient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Error("create index err: ", err)
		return cIndex, err
	}
	return cIndex, err
}

// DeleteIndex 删除索引
func DeleteIndex(indexName string) (del *elastic.IndicesDeleteResponse, err error) {
	del, err = ESClient.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		log.Error("delete index err: ", err)
		return del, nil
	}
	return del, nil
}

//  CreateDocumentById 根据传入的任务id，创建文档。
func CreateDocumentById(indexName, typeName, id string, bodyJson interface{}) (ret *elastic.IndexResponse, err error) {
	ret, err = ESClient.Index().Index(indexName).Type(typeName).Id(id).BodyJson(bodyJson).Do(ctx)
	if err != nil {
		log.Error("create doc err: ", err)
		return ret, err
	}
	return ret, err
}

// DeleteDocumentById 通过传入的任务id，删除文档。
func DeleteDocumentById(indexName, typeName, id string) (ret *elastic.DeleteResponse, err error) {
	ret, err = ESClient.Delete().Index(indexName).Type(typeName).Id(id).Do(ctx)
	if err != nil {
		log.Error("del doc err: ", err)
		return ret, err
	}
	return ret, err
}

// UpdateDocById 通过传入的任务id，更新文档。
func UpdateDocById(indexName, typeName, id, field string, value interface{}) (updatedoc *elastic.UpdateResponse, err error) {
	updatedoc, err = ESClient.Update().Index(indexName).Type(typeName).Id(id).Doc(map[string]interface{}{field: value}).Do(ctx)
	if err != nil {
		log.Error("update doc err: ", err)
		return
	}
	return
}

// GetDocById 根据传入的任务id，获取文档。
func GetDocById(indexName, typeName, id string) (doc *elastic.GetResult, err error) {
	doc, err = ESClient.Get().Index(indexName).Type(typeName).Id(id).Do(ctx)
	if err != nil {
		log.Error("get doc err: ", err)
		return
	}
	return
}

// GetAllDoc 获取所有文档，并按照给定字段排序
func GetAllDoc(indexName, typeName string, field string, ascending bool) (doc *elastic.SearchResult, err error) {
	query := elastic.MatchAllQuery{}
	doc, err = ESClient.Search().Index(indexName).Type(typeName).Query(query).Sort(field, ascending).Pretty(true).Do(ctx)
	if err != nil {
		log.Error("get all doc err: ", err)
		return
	}
	return
}

// TermQuery 传入字段名及字段值，查询文档
func TermQuery(indexName, typeName string, from, size int, fieldName string, fieldVal interface{}) (termres *elastic.SearchResult, err error) {
	termq := elastic.NewTermQuery(fieldName, fieldVal)
	termres, err = ESClient.Search().Index(indexName).Type(typeName).Query(termq).From(from).Size(size).Pretty(true).Do(ctx)
	if err != nil {
		log.Error("term query err: ", err)
		return
	}
	return
}

// BulkDoc 批量创建文档。
func BulkDoc(indexName, typeName string, structData []interface{}, docs []byte) (*elastic.BulkResponse, error) {
	if err := json.Unmarshal(docs, &structData); err != nil {
		log.Error("unmarshal err: ", err)
		return nil, err
	}
	bulk := ESClient.Bulk()
	for _, docVal := range structData {
		tmpDoc := elastic.NewBulkIndexRequest().Index(indexName).Type(typeName).Doc(docVal)
		bulk.Add(tmpDoc)
	}
	ret, err := bulk.Refresh("wait_for").Do(ctx)
	if err != nil {
		log.Error("bulk create doc err: ", err)
		return nil, err
	}
	return ret, nil
}

// SearchDoc 搜索符合传入参数条件的文档。
func SearchDoc(indexName, typeName string, bq *elastic.BoolQuery, from, size int, field string, ascending bool) (datArray []map[string]interface{}, err error) {
	var searchResult *elastic.SearchResult

	eQuery := ESClient.Search().
		Index(indexName).
		Type(typeName).
		Query(bq).
		From(from).
		Size(size).
		Sort(field, ascending)

	searchResult, err = eQuery.Pretty(true).Do(ctx)
	if err != nil {
		log.Error("search err: ", err)
		return
	}

	totalHits := searchResult.Hits.TotalHits
	if totalHits == 0 {
		err = errors.New("search num = 0")
		return
	}

	hits := searchResult.Hits.Hits
	datArray = make([]map[string]interface{}, len(hits))

	for i := 0; i < len(hits); i++ {
		hit := searchResult.Hits.Hits[i]
		var dat1 map[string]interface{}
		if err = json.Unmarshal(*hit.Source, &dat1); err != nil {
			log.Error("unmarshal err: ", err)
			return
		}
		datArray[i] = dat1
	}
	return datArray, nil
}
