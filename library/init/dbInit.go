package init

import (
	esModel "meigo/library/es"
	marketingToolDataModel "meigo/models/marketing_tool_data"
	peopleModel "meigo/models/people"
)

// DBInit init db.
func DBInit() {
	// 业务逻辑，初始化 people 的数据库连接
	peopleModel.InitPersonDB()

	// 业务逻辑，初始化 MarketingToolData 的数据库连接
	marketingToolDataModel.InitMarketingToolDataDB()

	esModel.ESInit()

}

// DBClose close db.
func DBClose() {
	// 业务逻辑，关闭 people 的数据库连接
	peopleModel.Close()
	marketingToolDataModel.Close()
}
