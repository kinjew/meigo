package db

import (
	"fmt"

	"meigo/library/log"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql 驱动加载
	"github.com/spf13/viper"
)

/*
// Gorm gorm配置参数
type Gorm struct {
	// TODO
		Debug        bool
		DBType       string
		MaxLifetime  int
		MaxOpenConns int
		MaxIdleConns int
		TablePrefix  string
		DSN          string
}
*/

// MySQL mysql配置参数
type MySQL struct {
	DBType     string
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	Parameters string
}

// ReadMysqlConfig 加载 myosql 参数
func ReadMysqlConfig(dbname string) (mysqlConf *MySQL) {
	if dbname == "" {
		log.Error("dbname is null.")
		return nil
	}
	dbname = fmt.Sprintf("mysql.%s", dbname)
	mysqlConf = new(MySQL)
	err := viper.UnmarshalKey(dbname, &mysqlConf)
	if err != nil {
		panic(err)
	}
	//fmt.Println("ReadMysqlConfig 中的 mysql 配置信息: ", MysqlConf)
	return
}
