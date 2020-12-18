package marketing_tool_data

import (
	"meigo/library/db"
	"sort"

	"github.com/jinzhu/gorm"
)

// sqlDB 是 *gorm.DB
var sqlDB *gorm.DB

//var err error

//var operatorList = []string{"=", ">", ">=", "<", "<=", "<>"}
//类型one操作符
var operatorTypeOneMap = map[string]string{"eq": "=", "neq": "<>", "gt": ">", "lt": "<", "gte": ">=", "lte": "<="}

//类型two操作符
var operatorTypeTwo = []string{"between", "lbetween", "rbetween", "ibetween"}
var operatorTypeTwoMap = map[string][]string{"between": []string{">=", "<="}, "lbetween": []string{">=", "<"}, "rbetween": []string{">", "<="}, "ibetween": []string{">", "<"}}

//类型three操作符
var operatorTypeThree = []string{"nBetween", "lNBetween", "rNBetween", "iNBetween"}
var operatorTypeThreeMap = map[string][]string{"nBetween": []string{"<=", ">="}, "lNBetween": []string{"<=", ">"}, "rNBetween": []string{"<", ">="}, "iNBetween": []string{"<", ">"}}

var page = "1"
var pageSize = "20"
var totalCount = 0

var orderBy = "id desc"

// InitPersonDB 初始化数据库
func InitMarketingToolDataDB() {
	var err error
	if sqlDB, err = db.ConnDB("marketing_tool_data"); err != nil {
		panic(err)
	}
	//sqlDB.AutoMigrate(&Person{})
}

//判断key是否在map中 isPermittedExpression
func isPermittedExpression(target string, operatorMap map[string]string) bool {
	_, exists := operatorMap[target]
	//fmt.Println(value, exists)
	if exists {
		return true
	}
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

// Close 关闭数据库
func Close() {
	db.Close(sqlDB)
}
