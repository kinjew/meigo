package marketing_tool_data

import (
	"meigo/library/db/common"
	"meigo/library/log"
	"strconv"

	"github.com/jinzhu/gorm"
	ctxExt "github.com/kinjew/gin-context-ext"
)

// GeneralStatAggregationData 实体
type GeneralStatAggregationData struct {
	common.BaseModelV2
	MainId uint `gorm:"column:main_id;" json:"main_id" form:"main_id" binding:"required"`
	//WxSystemUserId              int    `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" binding:"required"`
	//ToolId                      int    `gorm:"column:tool_id;" json:"tool_id" form:"tool_id" binding:"required"`
	//ToolType                    int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type" binding:"required"`
	WxSystemUserId                        uint `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id"`
	ToolId                                uint `gorm:"column:tool_id;" json:"tool_id" form:"tool_id"`
	ToolType                              int8 `gorm:"column:tool_type;" json:"tool_type" form:"tool_type"`
	VisitTotalTimes                       uint `gorm:"column:visit_total_times;" json:"visit_total_times" form:"visit_total_times"`
	VisitTotalNum                         uint `gorm:"column:visit_total_num;" json:"visit_total_num" form:"visit_total_num"`
	BMeetingOnlineEnrollNum               uint `gorm:"column:b_meeting_online_enroll_num;" json:"b_meeting_online_enroll_num" form:"b_meeting_online_enroll_num"`
	BMeetingOnlineEnrollApprovePassedNum  uint `gorm:"column:b_meeting_online_enroll_approve_passed_num;" json:"b_meeting_online_enroll_approve_passed_num" form:"b_meeting_online_enroll_approve_passed_num"`
	BMeetingOfflineEnrollNum              uint `gorm:"column:b_meeting_offline_enroll_num;" json:"b_meeting_offline_enroll_num" form:"b_meeting_offline_enroll_num"`
	BMeetingOfflineEnrollApprovePassedNum uint `gorm:"column:b_meeting_offline_enroll_approve_passed_num;" json:"b_meeting_offline_enroll_approve_passed_num" form:"b_meeting_offline_enroll_approve_passed_num"`
	AMeetingEnrollNum                     uint `gorm:"column:a_meeting_enroll_num;" json:"a_meeting_enroll_num" form:"a_meeting_enroll_num"`
	AMeetingEnrollApprovePassedNum        uint `gorm:"column:a_meeting_enroll_approve_passed_num;" json:"a_meeting_enroll_approve_passed_num" form:"a_meeting_enroll_approve_passed_num"`
	EnrollNum                             uint `gorm:"column:enroll_num;" json:"enroll_num" form:"enroll_num"`
	EnrollApprovePassedNum                uint `gorm:"column:enroll_approve_passed_num;" json:"enroll_approve_passed_num" form:"enroll_approve_passed_num"`
	WatchLiveNum                          uint `gorm:"column:watch_live_num;" json:"watch_live_num" form:"watch_live_num"`
	WatchLiveTimeAverage                  uint `gorm:"column:watch_live_time_average;" json:"watch_live_time_average" form:"watch_live_time_average"`
	OnlyWatchReplayNum                    uint `gorm:"column:only_watch_replay_num;" json:"only_watch_replay_num" form:"only_watch_replay_num"`
	OnlyWatchReplayTimeAverage            uint `gorm:"column:only_watch_replay_time_average;" json:"only_watch_replay_time_average" form:"only_watch_replay_time_average"`
	DownloadTotalTimes                    uint `gorm:"column:download_total_times;" json:"download_total_times" form:"download_total_times"`
	DownloadNum                           uint `gorm:"column:download_num;" json:"download_num" form:"download_num"`
	ShareTotalTimes                       uint `gorm:"column:share_total_times;" json:"share_total_times" form:"share_total_times"`
	ShowNum                               uint `gorm:"column:show_num;" json:"show_num" form:"show_num"`
	FavoriteTotalTimes                    uint `gorm:"column:favorite_total_times;" json:"favorite_total_times" form:"favorite_total_times"`
	FavoriteNum                           uint `gorm:"column:favorite_num;" json:"favorite_num" form:"favorite_num"`
	ViewMaterialTotalTimes                uint `gorm:"column:view_material_total_times;" json:"view_material_total_times" form:"view_material_total_times"`
	ViewMaterialNum                       uint `gorm:"column:view_material_num;" json:"view_material_num" form:"view_material_num"`
	NewFansNum                            uint `gorm:"column:new_fans_num;" json:"new_fans_num" form:"new_fans_num"`
	NewMemberNum                          uint `gorm:"column:new_member_num;" json:"new_member_num" form:"new_member_num"`
	SignNum                               uint `gorm:"column:sign_num;" json:"sign_num" form:"sign_num"`
	JoinNum                               uint `gorm:"column:join_num;" json:"join_num" form:"join_num"`
	JoinRate                              uint `gorm:"column:join_rate;" json:"join_rate" form:"join_rate"`
	TransferToLeadsNum                    uint `gorm:"column:transfer_to_leads_num;" json:"transfer_to_leads_num" form:"transfer_to_leads_num"`
	TransferToLeadsRate                   uint `gorm:"column:transfer_to_leads_rate;" json:"transfer_to_leads_rate" form:"transfer_to_leads_rate"`
	AcceptAsLeadsNum                      uint `gorm:"column:accept_as_leads_num;" json:"accept_as_leads_num" form:"accept_as_leads_num"`
	AcceptAsLeadsRate                     uint `gorm:"column:accept_as_leads_rate;" json:"accept_as_leads_rate" form:"accept_as_leads_rate"`
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
func (gsat *GeneralStatAggregationData) GsatQueryByParams(c *ctxExt.Context) (list []GeneralStatAggregationData, supplementData map[string]interface{}, err error) {
	/*
		err = sqlDB.Find(&list).Error
		return list, err
	*/
	var params GeneralStatAggregationData //请求参数
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

	mapQuery = gsatMapQueryGenerator(params, mapQuery, c)
	/*
		if params.MainId > 0 {
			mapQuery["main_id"] = params.MainId
		} else {
			err = errors.New("main_id不能为空")
			return list, err
		}

	*/

	//创建一个查询
	tx := sqlDB.Table("general_stat_aggregation_data").Where(mapQuery) //Map查询
	/*
		tx = tx.Joins("join action_live_data_5618 on action_live_data_5618.action_data_id = action_data_5618.id").Where("action_live_data_5618.last_live_login_city LIKE ?", "%中国3%")
		err = tx.Select("*").Scan(&list).Error
		return list, err
	*/
	//主表根据操作符查询
	tx, err = gsatOperatorQueryGenerator(params, tx, c)
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
func gsatMapQueryGenerator(params GeneralStatAggregationData, mapQuery map[string]interface{}, c *ctxExt.Context) map[string]interface{} {
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
	//其他查询参数
	/*
		if params.DateTime != "" {
			mapQuery["date_time"] = params.DateTime
		}
	*/
	return mapQuery
}

//gsatOperatorQueryGenerator构造基于操作符的查询
func gsatOperatorQueryGenerator(params GeneralStatAggregationData, tx *gorm.DB, c *ctxExt.Context) (txNew *gorm.DB, err error) {
	/*
		var ViewMaterialTimes = 0
		if params.ViewMaterialTimes != nil {
			ViewMaterialTimes = *params.ViewMaterialTimes
		}
		ViewMaterialTimesOperator := c.Query("view_material_times_operator")
		if ViewMaterialTimesOperator != "" {
			tx, err = operatorQueryAbstract(tx, c, "view_material_times", ViewMaterialTimesOperator, ViewMaterialTimes)
			if err != nil {
				return tx, err
			}
		}
	*/
	return tx, err
}