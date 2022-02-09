package marketing_tool_data

import (
	"fmt"
	"meigo/library/db/common"
	"meigo/library/log"
	"strconv"
	"strings"

	ctxExt "github.com/kinjew/gin-context-ext"
	"gorm.io/gorm"
)

// GeneralStatAggregationData 实体
type GeneralStatAggregationData struct {
	common.BaseModelV1
	MainId uint `gorm:"column:main_id;" json:"main_id" form:"main_id" binding:"required"`
	//WxSystemUserId              int    `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" binding:"required"`
	//ToolId                      int    `gorm:"column:tool_id;" json:"tool_id" form:"tool_id" binding:"required"`
	//ToolType                    int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type" binding:"required"`
	WxSystemUserId uint `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id"`
	ToolId         uint `gorm:"column:tool_id;" json:"tool_id" form:"tool_id"`
	ToolType       int8 `gorm:"column:tool_type;" json:"tool_type" form:"tool_type"`
	//MeetingEndTime         int8 `gorm:"column:meeting_end_time;" json:"meeting_end_time" form:"meeting_end_time"`
	VisitTimes             uint `gorm:"column:visit_times;" json:"visit_times" form:"visit_times"`
	VisitTimesUnidentified uint `gorm:"column:visit_times_unidentified;" json:"visit_times_unidentified" form:"visit_times_unidentified"`
	VisitTimesFans         uint `gorm:"column:visit_times_fans;" json:"visit_times_fans" form:"visit_times_fans"`
	VisitTimesMember       uint `gorm:"column:visit_times_member;" json:"visit_times_member" form:"visit_times_member"`
	VisitTimesEmployee     uint `gorm:"column:visit_times_employee;" json:"visit_times_employee" form:"visit_times_employee"`

	VisitNum             uint `gorm:"column:visit_num;" json:"visit_num" form:"visit_num"`
	VisitNumUnidentified uint `gorm:"column:visit_num_unidentified;" json:"visit_num_unidentified" form:"visit_num_unidentified"`
	VisitNumFans         uint `gorm:"column:visit_num_fans;" json:"visit_num_fans" form:"visit_num_fans"`
	VisitNumMember       uint `gorm:"column:visit_num_member;" json:"visit_num_member" form:"visit_num_member"`
	VisitNumEmployee     uint `gorm:"column:visit_num_employee;" json:"visit_num_employee" form:"visit_num_employee"`

	BMeetingOnlineEnrollNum               uint `gorm:"column:b_meeting_online_enroll_num;" json:"b_meeting_online_enroll_num" form:"b_meeting_online_enroll_num"`
	BMeetingOnlineEnrollApprovePassedNum  uint `gorm:"column:b_meeting_online_enroll_approve_passed_num;" json:"b_meeting_online_enroll_approve_passed_num" form:"b_meeting_online_enroll_approve_passed_num"`
	BMeetingOfflineEnrollNum              uint `gorm:"column:b_meeting_offline_enroll_num;" json:"b_meeting_offline_enroll_num" form:"b_meeting_offline_enroll_num"`
	BMeetingOfflineEnrollApprovePassedNum uint `gorm:"column:b_meeting_offline_enroll_approve_passed_num;" json:"b_meeting_offline_enroll_approve_passed_num" form:"b_meeting_offline_enroll_approve_passed_num"`
	AMeetingEnrollNum                     uint `gorm:"column:a_meeting_enroll_num;" json:"a_meeting_enroll_num" form:"a_meeting_enroll_num"`
	AMeetingEnrollApprovePassedNum        uint `gorm:"column:a_meeting_enroll_approve_passed_num;" json:"a_meeting_enroll_approve_passed_num" form:"a_meeting_enroll_approve_passed_num"`
	GenerateFissionNum                    uint `gorm:"column:generate_fission_num;" json:"generate_fission_num" form:"generate_fission_num"`
	EnrollNum                             uint `gorm:"column:enroll_num;" json:"enroll_num" form:"enroll_num"`
	EnrollNumUnidentified                 uint `gorm:"column:enroll_num_unidentified;" json:"enroll_num_unidentified" form:"enroll_num_unidentified"`
	EnrollNumFans                         uint `gorm:"column:enroll_num_fans;" json:"enroll_num_fans" form:"enroll_num_fans"`
	EnrollNumMember                       uint `gorm:"column:enroll_num_member;" json:"enroll_num_member" form:"enroll_num_member"`
	EnrollNumEmployee                     uint `gorm:"column:enroll_num_employee;" json:"enroll_num_employee" form:"enroll_num_employee"`

	EnrollApprovePassedNum          uint `gorm:"column:enroll_approve_passed_num;" json:"enroll_approve_passed_num" form:"enroll_approve_passed_num"`
	EnrollApprovePassedUnidentified uint `gorm:"column:enroll_approve_passed_unidentified;" json:"enroll_approve_passed_unidentified" form:"enroll_approve_passed_unidentified"`
	EnrollApprovePassedNumFans      uint `gorm:"column:enroll_approve_passed_num_fans;" json:"enroll_approve_passed_num_fans" form:"enroll_approve_passed_num_fans"`
	EnrollApprovePassedNumMember    uint `gorm:"column:enroll_approve_passed_num_member;" json:"enroll_approve_passed_num_member" form:"enroll_approve_passed_num_member"`
	EnrollApprovePassedNumEmployee  uint `gorm:"column:enroll_approve_passed_num_employee;" json:"enroll_approve_passed_num_employee" form:"enroll_approve_passed_num_employee"`

	EnrollApproveNotPassedNum uint `gorm:"column:enroll_approve_not_passed_num;" json:"enroll_approve_not_passed_num" form:"enroll_approve_not_passed_num"`

	ClickEnterLiveTimes         uint `gorm:"column:click_enter_live_times;" json:"click_enter_live_times" form:"click_enter_live_times"`
	ClickEnterLiveTimesFans     uint `gorm:"column:click_enter_live_times_fans;" json:"click_enter_live_times_fans" form:"click_enter_live_times_fans"`
	ClickEnterLiveTimesMember   uint `gorm:"column:click_enter_live_times_member;" json:"click_enter_live_times_member" form:"click_enter_live_times_member"`
	ClickEnterLiveTimesEmployee uint `gorm:"column:click_enter_live_times_employee;" json:"click_enter_live_times_employee" form:"click_enter_live_times_employee"`

	ClickEnterLiveNum         uint `gorm:"column:click_enter_live_num;" json:"click_enter_live_num" form:"click_enter_live_num"`
	ClickEnterLiveNumFans     uint `gorm:"column:click_enter_live_num_fans;" json:"click_enter_live_num_fans" form:"click_enter_live_num_fans"`
	ClickEnterLiveNumMember   uint `gorm:"column:click_enter_live_num_member;" json:"click_enter_live_num_member" form:"click_enter_live_num_member"`
	ClickEnterLiveNumEmployee uint `gorm:"column:click_enter_live_num_employee;" json:"click_enter_live_num_employee" form:"click_enter_live_num_employee"`

	ClickWatchReplayTimes uint `gorm:"column:click_watch_replay_times;" json:"click_watch_replay_times" form:"click_watch_replay_times"`
	ClickWatchReplayNum   uint `gorm:"column:click_watch_replay_num;" json:"click_watch_replay_num" form:"click_watch_replay_num"`

	WatchLiveTime              uint `gorm:"column:watch_live_time;" json:"watch_live_time" form:"watch_live_time"`
	WatchLiveTimes             uint `gorm:"column:watch_live_times;" json:"watch_live_times" form:"watch_live_times"`
	WatchLiveNum               uint `gorm:"column:watch_live_num;" json:"watch_live_num" form:"watch_live_num"`
	WatchLiveTimeAverage       uint `gorm:"column:watch_live_time_average;" json:"watch_live_time_average" form:"watch_live_time_average"`
	OnlyWatchReplayTime        uint `gorm:"column:only_watch_replay_time;" json:"only_watch_replay_time" form:"only_watch_replay_time"`
	OnlyWatchReplayTimes       uint `gorm:"column:only_watch_replay_times;" json:"only_watch_replay_times" form:"only_watch_replay_times"`
	OnlyWatchReplayNum         uint `gorm:"column:only_watch_replay_num;" json:"only_watch_replay_num" form:"only_watch_replay_num"`
	OnlyWatchReplayTimeAverage uint `gorm:"column:only_watch_replay_time_average;" json:"only_watch_replay_time_average" form:"only_watch_replay_time_average"`

	WatchReplayTime        uint `gorm:"column:watch_replay_time;" json:"watch_replay_time" form:"watch_replay_time"`
	WatchReplayTimes       uint `gorm:"column:watch_replay_times;" json:"watch_replay_times" form:"watch_replay_times"`
	WatchReplayNum         uint `gorm:"column:watch_replay_num;" json:"watch_replay_num" form:"watch_replay_num"`
	WatchReplayTimeAverage uint `gorm:"column:watch_replay_time_average;" json:"watch_replay_time_average" form:"watch_replay_time_average"`

	DownloadTimes          uint `gorm:"column:download_times;" json:"download_times" form:"download_times"`
	DownloadNum            uint `gorm:"column:download_num;" json:"download_num" form:"download_num"`
	ShareTimes             uint `gorm:"column:share_times;" json:"share_times" form:"share_times"`
	ShareTimesUnidentified uint `gorm:"column:share_times_unidentified;" json:"share_times_unidentified" form:"share_times_unidentified"`
	ShareTimesFans         uint `gorm:"column:share_times_fans;" json:"share_times_fans" form:"share_times_fans"`
	ShareTimesMember       uint `gorm:"column:share_times_member;" json:"share_times_member" form:"share_times_member"`
	ShareTimesEmployee     uint `gorm:"column:share_times_employee;" json:"share_times_employee" form:"share_times_employee"`

	ShareNum             uint `gorm:"column:share_num;" json:"share_num" form:"share_num"`
	ShareNumUnidentified uint `gorm:"column:share_num_unidentified;" json:"share_num_unidentified" form:"share_num_unidentified"`
	ShareNumFans         uint `gorm:"column:share_num_fans;" json:"share_num_fans" form:"share_num_fans"`
	ShareNumMember       uint `gorm:"column:share_num_member;" json:"share_num_member" form:"share_num_member"`
	ShareNumEmployee     uint `gorm:"column:share_num_employee;" json:"share_num_employee" form:"share_num_employee"`

	//FavoriteTimes         uint `gorm:"column:favorite_times;" json:"favorite_times" form:"favorite_times"`
	//FavoriteTimesMember   uint `gorm:"column:favorite_times_member;" json:"favorite_times_member" form:"favorite_times_member"`
	//FavoriteTimesEmployee uint `gorm:"column:favorite_times_employee;" json:"favorite_times_employee" form:"favorite_times_employee"`

	FavoriteNum         uint `gorm:"column:favorite_num;" json:"favorite_num" form:"favorite_num"`
	FavoriteNumMember   uint `gorm:"column:favorite_num_member;" json:"favorite_num_member" form:"favorite_num_member"`
	FavoriteNumEmployee uint `gorm:"column:favorite_num_employee;" json:"favorite_num_employee" form:"favorite_num_employee"`

	ViewMaterialTimes uint `gorm:"column:view_material_times;" json:"view_material_times" form:"view_material_times"`
	ViewMaterialNum   uint `gorm:"column:view_material_num;" json:"view_material_num" form:"view_material_num"`
	OldFansNum        uint `gorm:"column:old_fans_num;" json:"old_fans_num" form:"old_fans_num"`
	NewFansNum        uint `gorm:"column:new_fans_num;" json:"new_fans_num" form:"new_fans_num"`
	OldMemberNum      uint `gorm:"column:old_member_num;" json:"old_member_num" form:"old_member_num"`
	NewMemberNum      uint `gorm:"column:new_member_num;" json:"new_member_num" form:"new_member_num"`
	SignNum           uint `gorm:"column:sign_num;" json:"sign_num" form:"sign_num"`
	UnsignNum         uint `gorm:"column:unsign_num;" json:"unsign_num" form:"unsign_num"`
	//JoinNum                               uint `gorm:"column:join_num;" json:"join_num" form:"join_num"`
	JoinRate            uint `gorm:"column:join_rate;" json:"join_rate" form:"join_rate"`
	TransferToLeadsNum  uint `gorm:"column:transfer_to_leads_num;" json:"transfer_to_leads_num" form:"transfer_to_leads_num"`
	TransferToLeadsRate uint `gorm:"column:transfer_to_leads_rate;" json:"transfer_to_leads_rate" form:"transfer_to_leads_rate"`
	AcceptAsLeadsNum    uint `gorm:"column:accept_as_leads_num;" json:"accept_as_leads_num" form:"accept_as_leads_num"`
	AcceptAsLeadsRate   uint `gorm:"column:accept_as_leads_rate;" json:"accept_as_leads_rate" form:"accept_as_leads_rate"`

	VisitAndEnrollNum               uint `gorm:"column:visit_and_enroll_num;" json:"visit_and_enroll_num" form:"visit_and_enroll_num"`
	VisitAndEnrollApprovePassedNum  uint `gorm:"column:visit_and_enroll_approve_passed_num;" json:"visit_and_enroll_approve_passed_num" form:"visit_and_enroll_approve_passed_num"`
	VisitAndWatchLiveOrReplayOrSign uint `gorm:"column:visit_and_watch_live_or_replay_or_sign;" json:"visit_and_watch_live_or_replay_or_sign" form:"visit_and_watch_live_or_replay_or_sign"`
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

	//gsat范围查询
	tx = inQueryGeneratorGsat(tx, c)

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

//inQueryGeneratorGsat 构造in查询
func inQueryGeneratorGsat(tx *gorm.DB, c *ctxExt.Context) *gorm.DB {

	//tool_id_list为tool_id以逗号分割的字符串
	ToolIdList := c.Query("tool_id_list")
	if ToolIdList != "" {
		ToolIdArr := strings.Split(ToolIdList, ",")
		if len(ToolIdArr) != 0 {
			fmt.Println("ToolIdArr: ", ToolIdArr)
			tx = tx.Where("tool_id IN (?) ", ToolIdArr)
		}
	}

	return tx
}
