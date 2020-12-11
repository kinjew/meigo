package marketing_tool_data

import (
	"meigo/library/db"
	"sort"

	"github.com/jinzhu/gorm"
)

// sqlDB 是 *gorm.DB
var sqlDB *gorm.DB
var err error

var operatorList = []string{"=", ">", ">=", "<", "<=", "<>"}
var page = "1"
var pageSize = "20"
var totalCount = 0

var orderBy = "id desc"

// InitPersonDB 初始化数据库
func InitMarketingToolDataDB() {

	if sqlDB, err = db.ConnDB("marketing_tool_data"); err != nil {
		panic(err)
	}
	//sqlDB.AutoMigrate(&Person{})
}

//判断操作符是否在切片中
func isPermittedOperator(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// Close 关闭数据库
func Close() {
	db.Close(sqlDB)
}
