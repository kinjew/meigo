package init

import (
	gqlModel "meigo/gqlgen-todos/graph/model"
	esModel "meigo/library/es"
	marketingToolDataModel "meigo/models/marketing_tool_data"
	wfModel "meigo/models/wf"
	//peopleModel "meigo/models/people"
)

// DBInit init db.
func DBInit() {
	// 业务逻辑，初始化 people 的数据库连接
	//peopleModel.InitPersonDB()

	// 业务逻辑，初始化 MarketingToolData 的数据库连接
	marketingToolDataModel.InitMarketingToolDataDB()

	// 业务逻辑，初始化 wf 的数据库连接
	wfModel.InitWfDB()

	// 业务逻辑，初始化 wf 的数据库连接
	gqlModel.InitTestDB()

	esModel.ESInit()

}

// DBClose close db.
func DBClose() {
	// 业务逻辑，关闭 people 的数据库连接
	//peopleModel.Close()
	marketingToolDataModel.Close()
}
