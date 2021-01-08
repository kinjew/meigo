package marketing_tool_data

import (
	"meigo/library/db/common"
	"meigo/library/log"
	"strconv"

	"github.com/jinzhu/gorm"
	ctxExt "github.com/kinjew/gin-context-ext"
)

// GeneralStatData 实体
type GeneralStatData struct {
	common.BaseModelV1
	MainId uint `gorm:"column:main_id;" json:"main_id" form:"main_id" binding:"required"`
	//WxSystemUserId              int    `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" binding:"required"`
	//ToolId                      int    `gorm:"column:tool_id;" json:"tool_id" form:"tool_id" binding:"required"`
	//ToolType                    int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type" binding:"required"`
	WxSystemUserId            uint   `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id"`
	ToolId                    uint   `gorm:"column:tool_id;" json:"tool_id" form:"tool_id"`
	ToolType                  int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type"`
	DateTime                  string `gorm:"column:date_time;" json:"date_time" form:"date_time"`
	VisitTimes                uint   `gorm:"column:visit_times;" json:"visit_times" form:"visit_times"`
	VisitNum                  uint   `gorm:"column:visit_num;" json:"visit_num" form:"visit_num"`
	EnrollNum                 uint   `gorm:"column:enroll_num;" json:"enroll_num" form:"enroll_num"`
	EnrollApprovePassedNum    uint   `gorm:"column:enroll_approve_passed_num;" json:"enroll_approve_passed_num" form:"enroll_approve_passed_num"`
	EnrollApproveNotPassedNum uint   `gorm:"column:enroll_approve_not_passed_num;" json:"enroll_approve_not_passed_num" form:"enroll_approve_not_passed_num"`

	ClickEnterLiveTimes   uint `gorm:"column:click_enter_live_times;" json:"click_enter_live_times" form:"click_enter_live_times"`
	ClickEnterLiveNum     uint `gorm:"column:click_enter_live_num;" json:"click_enter_live_num" form:"click_enter_live_num"`
	ClickWatchReplayTimes uint `gorm:"column:click_watch_replay_times;" json:"click_watch_replay_times" form:"click_watch_replay_times"`
	ClickWatchReplayNum   uint `gorm:"column:click_watch_replay_num;" json:"click_watch_replay_num" form:"click_watch_replay_num"`
	DownloadTimes         uint `gorm:"column:download_times;" json:"download_times" form:"download_times"`
	DownloadNum           uint `gorm:"column:download_num;" json:"download_num" form:"download_num"`
	ShareTimes            uint `gorm:"column:share_times;" json:"share_times" form:"share_times"`
	ShareNum              uint `gorm:"column:share_num;" json:"share_num" form:"share_num"`
	//FavoriteTimes         uint `gorm:"column:favorite_times;" json:"favorite_times" form:"favorite_times"`
	FavoriteNum       uint `gorm:"column:favorite_num;" json:"favorite_num" form:"favorite_num"`
	ViewMaterialTimes uint `gorm:"column:view_material_times;" json:"view_material_times" form:"view_material_times"`
	ViewMaterialNum   uint `gorm:"column:view_material_num;" json:"view_material_num" form:"view_material_num"`
	NewFansNum        uint `gorm:"column:new_fans_num;" json:"new_fans_num" form:"new_fans_num"`
	NewMemberNum      uint `gorm:"column:new_member_num;" json:"new_member_num" form:"new_member_num"`
	/*
		SignNum                uint   `gorm:"column:sign_num;" json:"sign_num" form:"sign_num"`
		JoinNum                uint   `gorm:"column:join_num;" json:"join_num" form:"join_num"`
		JoinRate               uint   `gorm:"column:join_rate;" json:"join_rate" form:"join_rate"`
		TransferToLeadsNum     uint   `gorm:"column:transfer_to_leads_num;" json:"transfer_to_leads_num" form:"transfer_to_leads_num"`
		TransferToLeadsRate    uint   `gorm:"column:transfer_to_leads_rate;" json:"transfer_to_leads_rate" form:"transfer_to_leads_rate"`
		AcceptAsLeadsNum       uint   `gorm:"column:accept_as_leads_num;" json:"accept_as_leads_num" form:"accept_as_leads_num"`
		AcceptAsLeadsRate      uint   `gorm:"column:accept_as_leads_rate;" json:"accept_as_leads_rate" form:"accept_as_leads_rate"`
	*/
}

/*
type supplementData struct {
	pageInt     int
	pageSizeInt int
	totalCount  int
}
*/
/*
func init() {
	InitMarketingToolDataDB()
}
*/
/*
GetPeople 获取人员列表
*/
/*
func (p *ActionData) QueryByParams(c *ctxExt.Context) (people []ActionData, err error) {
	err = sqlDB.Find(&people).Error
	return people, err
}
*/

/*
QueryByParams 获取行为数据
*/
func (gst *GeneralStatData) GstQueryByParams(c *ctxExt.Context) (list []GeneralStatData, supplementData map[string]interface{}, err error) {
	/*
		err = sqlDB.Find(&list).Error
		return list, err
	*/
	var params GeneralStatData //请求参数
	supplementData = make(map[string]interface{})

	if err := c.ShouldBindQuery(&params); err != nil {
		log.Error("BindJSON err: ", err)
		return list, supplementData, err
	}
	/*
		fmt.Println("params: ", params)
		sqlDB.Find(&list)
		fmt.Println("list: ", list)
		os.Exit(200)
	*/
	mapQuery := make(map[string]interface{})

	mapQuery = gstMapQueryGenerator(params, mapQuery, c)
	/*
		if params.MainId > 0 {
			mapQuery["main_id"] = params.MainId
		} else {
			err = errors.New("main_id不能为空")
			return list, err
		}

	*/

	//创建一个查询
	tx := sqlDB.Table("general_stat_data").Where(mapQuery) //Map查询
	/*
		tx = tx.Joins("join action_live_data_5618 on action_live_data_5618.action_data_id = action_data_5618.id").Where("action_live_data_5618.last_live_login_city LIKE ?", "%中国3%")
		err = tx.Select("*").Scan(&list).Error
		return list, err
	*/
	//主表根据操作符查询
	tx, err = gstOperatorQueryGenerator(params, tx, c)
	if err != nil {
		return
	}

	//执行查询操作
	page = c.DefaultQuery("page", page)
	pageSize = c.DefaultQuery("pageSize", pageSize)
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	//计算偏移
	offSet := (pageInt - 1) * pageSizeInt

	//获取总数
	tx.Count(&totalCount)

	//err = tx.Select("*").Scan(&list).Error
	err = tx.Select("*").Offset(offSet).Limit(pageSizeInt).Scan(&list).Error

	supplementData["page"] = pageInt
	supplementData["pageSize"] = pageSizeInt
	supplementData["totalCount"] = totalCount
	/*
		fmt.Println("sD: ", supplementData)
		fmt.Println("list: ", list)
	*/
	return list, supplementData, err
}

//构造mapQuery对象
func gstMapQueryGenerator(params GeneralStatData, mapQuery map[string]interface{}, c *ctxExt.Context) map[string]interface{} {
	//return mapQuery
	if params.MainId > 0 {
		mapQuery["main_id"] = params.MainId
	}
	if params.WxSystemUserId > 0 {
		mapQuery["wx_system_user_id"] = params.WxSystemUserId
	}
	if params.ToolId > 0 {
		mapQuery["tool_id"] = params.ToolId
	}
	if params.ToolType > 0 {
		mapQuery["tool_type"] = params.ToolType
	}
	/*
		if params.DateTime != "" {
			mapQuery["date_time"] = params.DateTime
		}
	*/
	return mapQuery
}

//gstOperatorQueryGenerator构造基于操作符的查询
func gstOperatorQueryGenerator(params GeneralStatData, tx *gorm.DB, c *ctxExt.Context) (txNew *gorm.DB, err error) {

	DateTimeOperator := c.Query("date_time_operator")
	if DateTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "date_time", DateTimeOperator, params.DateTime)
		if err != nil {
			return tx, err
		}
	}

	return tx, err
}
