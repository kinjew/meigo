package db

import (
	"fmt"

	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"

	"meigo/library/log"
	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql" // 不使用
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is *gorm.DB
var DB *gorm.DB

// ConnDB 创建数据库连接, dbname 为数据库名
func ConnDB(dbname string) (db *gorm.DB, err error) {
	if dbname == "" {
		log.Error("dbname is null.")
		return nil, err
	}
	msqlConf := ReadMysqlConfig(dbname)
	db, err = InitDB(msqlConf)
	if err != nil {
		log.Error("connect db failed.", err)
		return nil, err
	}

	return db, err
}

// InitDB used for cli
func InitDB(MysqlConf *MySQL) (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		MysqlConf.User,
		MysqlConf.Password,
		MysqlConf.Host,
		MysqlConf.Port,
		MysqlConf.DBName,
		MysqlConf.Parameters,
	)

	//fmt.Println("print:" + MysqlConf.DBType)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config, // DSN data source name
		DefaultStringSize:         256,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "it_", // 表名前缀，`Article` 的表名应该是 `it_articles`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	})
	if err != nil {
		log.Error("Database connection failed. Database name: %s", MysqlConf.DBName, err)
		return nil, err
	}

	// set for db connection
	//setupDB(db)

	return db, err
}

/*
func setupDB(db *gorm.DB) {

	// 表名单数形式 （默认复数形式）
	db.SingularTable(true)

	db.LogMode(true)

	logger := &MyLogger{}
	db.SetLogger(logger)

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(10)
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Duration(7200) * time.Second)
}

*/

// Close close db
func Close(db *gorm.DB) {
	/*
		err := db.Close()
		if err != nil {
			log.Error("DB close error", err)
		}*/
}

//MyLogger 是一个用户类型
type MyLogger struct {
}

/*
Print 创建实现Print方法的接口，个性化日志记录
第一个参数为 level，表示这个是个什么请求（有sql和log两种类型）
第二个参数为打印sql的代码行号
第三个参数是执行时间戳
第四个参数是sql语句
第五个参数是如果有预处理，请求参数
第六个参数是这个sql影响的行数。
*/
func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level        = values[0]
		source       = values[1]
		excTime      = values[2]
		sql          = values[3]
		rparams      = values[4]
		affectedRows = values[5]
	)
	if level == "sql" { //level可以是sql或者log
		log.Info("gorm sql log:", level, sql, excTime, source, rparams, affectedRows)
	} else {
		log.Info("gorm sql log:", values)
	}
}
