package init

import (
	esModel "meigo/library/es"
	peopleModel "meigo/models/people"
)

// DBInit init db.
func DBInit() {
	// 业务逻辑，初始化 people 的数据库连接
	peopleModel.InitPersonDB()

	esModel.ESInit()

}

// DBClose close db.
func DBClose() {
	// 业务逻辑，关闭 people 的数据库连接
	peopleModel.Close()
}
