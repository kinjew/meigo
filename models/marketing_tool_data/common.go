package marketing_tool_data

import (
	"errors"
	"fmt"
	"meigo/library/db"
	"sort"
	"strings"

	ctxExt "github.com/kinjew/gin-context-ext"

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

type OperateFormat struct {
	Operator string      `json:"operator"`
	Field    string      `json:"field"`
	ValueOne interface{} `json:"value_one"`
	ValueTwo interface{} `json:"value_two"`
}

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

//operatorQueryAbstract 抽象特殊查询操作符
func operatorQueryAbstract(tx *gorm.DB, c *ctxExt.Context, fieldName, operator string, operatorValue interface{}) (txNew *gorm.DB, err error) {
	//fmt.Println("operator", operator)
	if isPermittedExpression(operator, operatorTypeOneMap) {
		//fmt.Println("isPermittedExpression", "ok")
		tx = tx.Where(fieldName+" "+operatorTypeOneMap[operator]+" ?", operatorValue)
	} else if isStringInSlice(operator, operatorTypeTwo) {
		//取数据表名称，排除数据库名称
		fieldNameSlice := strings.Split(fieldName, ".")
		fieldNameLow := c.Query(fieldNameSlice[len(fieldNameSlice)-1] + "_low")
		fieldNameHigh := c.Query(fieldNameSlice[len(fieldNameSlice)-1] + "_high")
		//fmt.Println("fieldName,fieldNameLow,fieldNameHigh", fieldName, fieldNameLow, fieldNameHigh)
		if fieldNameLow == "" || fieldNameHigh == "" {
			//fmt.Println("here", fieldNameLow, fieldNameHigh)
			err := errors.New(fieldNameSlice[len(fieldNameSlice)-1] + "_low" + "或" + fieldNameSlice[len(fieldNameSlice)-1] + "_high" + "为空")
			return tx, err
		}
		fmt.Println("operatorTypeTwoMap", operatorTypeTwoMap[operator][0], operatorTypeTwoMap[operator][1])
		tx = tx.Where(fieldName+" "+operatorTypeTwoMap[operator][0]+"  ?", fieldNameLow).Where(fieldName+" "+operatorTypeTwoMap[operator][1]+"  ?", fieldNameHigh)
	} else if isStringInSlice(operator, operatorTypeThree) {
		//取数据表名称，排除数据库名称
		fieldNameSlice := strings.Split(fieldName, ".")
		fieldNameLow := c.Query(fieldNameSlice[len(fieldNameSlice)-1] + "_low")
		fieldNameHigh := c.Query(fieldNameSlice[len(fieldNameSlice)-1] + "_high")
		//fmt.Println("fieldName,fieldNameLow,fieldNameHigh", fieldName, fieldNameLow, fieldNameHigh)
		if fieldNameLow == "" || fieldNameHigh == "" {
			//fmt.Println("here", fieldNameLow, fieldNameHigh)
			err := errors.New(fieldNameSlice[len(fieldNameSlice)-1] + "_low" + "或" + fieldNameSlice[len(fieldNameSlice)-1] + "_high" + "为空")
			return tx, err
		}
		fmt.Println("operatorTypeThree", operatorTypeThreeMap[operator][0], operatorTypeThreeMap[operator][1])
		tx = tx.Where(fieldName+" "+operatorTypeThreeMap[operator][0]+"  ?", fieldNameLow).Where(fieldName+" "+operatorTypeThreeMap[operator][1]+"  ?", fieldNameHigh)
	} else {
		err := errors.New("invalid operator")
		return tx, err
	}
	return tx, err
}

// Close 关闭数据库
func Close() {
	db.Close(sqlDB)
}
