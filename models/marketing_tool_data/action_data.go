package marketing_tool_data

import (
	"errors"
	"fmt"
	"html"
	"meigo/library/db/common"
	"meigo/library/log"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	ctxExt "github.com/kinjew/gin-context-ext"
)

// ActionData 实体
type ActionData struct {
	//WxSystemUserId              int    `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" binding:"required"`
	//ToolId                      int    `gorm:"column:tool_id;" json:"tool_id" form:"tool_id" binding:"required"`
	//ToolType                    int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type" binding:"required"`
	//UserIdentityType            int    `gorm:"column:user_identity_type;" json:"user_identity_type" form:"user_identity_type"`
	common.BaseModelV1
	MainId            int    `gorm:"column:main_id;" json:"main_id" form:"main_id" binding:"required"`
	WxSystemUserId    int    `gorm:"column:wx_system_user_id;" json:"wx_system_user_id" form:"wx_system_user_id" `
	ToolId            int    `gorm:"column:tool_id;" json:"tool_id" form:"tool_id"`
	ToolType          int8   `gorm:"column:tool_type;" json:"tool_type" form:"tool_type"`
	MemberId          int    `gorm:"column:member_id;" json:"member_id" form:"member_id"`
	WxOpenId          string `gorm:"column:wx_open_id;" json:"wx_open_id" form:"wx_open_id"`
	ClientIp          string `gorm:"column:client_ip;" json:"client_ip" form:"client_ip"`
	FirstVisitBrowser string `gorm:"column:first_visit_browser;" json:"first_visit_browser" form:"first_visit_browser"`
	FirstVisitClient  int    `gorm:"column:first_visit_client;" json:"first_visit_client" form:"first_visit_client"`
	//FirstVisitEquipment         string `gorm:"column:first_visit_equipment;" json:"first_visit_equipment" form:"first_visit_equipment"`
	FirstVisitChannelId          int    `gorm:"column:first_visit_channel_id;" json:"first_visit_channel_id" form:"first_visit_channel_id"`
	FollowChannelId              int    `gorm:"column:follow_channel_id;" json:"follow_channel_id" form:"follow_channel_id"`
	EnrollChannelId              int    `gorm:"column:enroll_channel_id;" json:"enroll_channel_id" form:"enroll_channel_id"`
	WxFollowInviterId            int    `gorm:"column:wx_follow_inviter_id;" json:"wx_follow_inviter_id" form:"wx_follow_inviter_id"`
	MeetingEnrollInviterId       int    `gorm:"column:meeting_enroll_inviter_id;" json:"meeting_enroll_inviter_id" form:"meeting_enroll_inviter_id"`
	IsEnroll                     *int   `gorm:"column:is_enroll;" json:"is_enroll" form:"is_enroll"`
	EnrollType                   int    `gorm:"column:enroll_type;" json:"enroll_type" form:"enroll_type"`
	EnrollWay                    int    `gorm:"column:enroll_way;" json:"enroll_way" form:"enroll_way"`
	EnrollTime                   int    `gorm:"column:enroll_time;" json:"enroll_time" form:"enroll_time"`
	EnrollMeetingStatus          int    `gorm:"column:enroll_meeting_status;" json:"enroll_meeting_status" form:"enroll_meeting_status"`
	EnrollApproveStatus          int    `gorm:"column:enroll_approve_status;" json:"enroll_approve_status" form:"enroll_approve_status"`
	IsSign                       *int   `gorm:"column:is_sign;" json:"is_sign" form:"is_sign"`
	IsNewFans                    *int   `gorm:"column:is_new_fans;" json:"is_new_fans" form:"is_new_fans"`
	IsNewMember                  *int   `gorm:"column:is_new_member;" json:"is_new_member" form:"is_new_member"`
	PosterInviteFollowNum        *int   `gorm:"column:poster_invite_follow_num;" json:"poster_invite_follow_num" form:"poster_invite_follow_num"`
	PayMoney                     *int   `gorm:"column:pay_money;" json:"pay_money" form:"pay_money"`
	IsWatchLive                  *int   `gorm:"column:is_watch_live;" json:"is_watch_live" form:"is_watch_live"`
	IsWatchReplay                *int   `gorm:"column:is_watch_replay;" json:"is_watch_replay" form:"is_watch_replay"`
	VisitTimes                   *int   `gorm:"column:visit_times;" json:"visit_times" form:"visit_times"`
	DownloadTimes                *int   `gorm:"column:download_times;" json:"download_times" form:"download_times"`
	ShareTimes                   *int   `gorm:"column:share_times;" json:"share_times" form:"share_times"`
	FavoriteTimes                *int   `gorm:"column:favorite_times;" json:"favorite_times" form:"favorite_times"`
	ViewMaterialTimes            *int   `gorm:"column:view_material_times;" json:"view_material_times" form:"view_material_times"`
	SeatNumber                   string `gorm:"column:seat_number;" json:"seat_number" form:"seat_number"`
	IsClickEnrollButton          *int   `gorm:"column:is_click_enroll_button;" json:"is_click_enroll_button" form:"is_click_enroll_button"`
	IsClickPayButton             *int   `gorm:"column:is_click_pay_button;" json:"is_click_pay_button" form:"is_click_pay_button"`
	IsClickGenerateFissionButton *int   `gorm:"column:is_click_generate_fission_button;" json:"is_click_generate_fission_button" form:"is_click_generate_fission_button"`
	IsClickEnterLiveButton       *int   `gorm:"column:is_click_enter_live_button;" json:"is_click_enter_live_button" form:"is_click_enter_live_button"`
	IsClickWatchReplayButton     *int   `gorm:"column:is_click_watch_replay_button;" json:"is_click_watch_replay_button" form:"is_click_watch_replay_button"`
	IsDel                        *int   `gorm:"column:is_del;" json:"is_del" form:"is_del"`
	ActionLiveData
}

// ActionLiveData 实体
type ActionLiveData struct {
	ActionDataId          int    `gorm:"column:action_data_id;" json:"action_data_id" form:"action_data_id"`
	LivePlatformType      int    `gorm:"column:live_platform_type;" json:"live_platform_type" form:"live_platform_type"`
	LastLiveWatchClient   string `gorm:"column:last_live_watch_client;" json:"last_live_watch_client" form:"last_live_watch_client"`
	LastLiveLeaveTime     int    `gorm:"column:last_live_leave_time;" json:"last_live_leave_time" form:"last_live_leave_time"`
	LastLiveLoginCity     string `gorm:"column:last_live_login_city;" json:"last_live_login_city" form:"last_live_login_city"`
	LiveWatchTime         int    `gorm:"column:live_watch_time;" json:"live_watch_time" form:"live_watch_time"`
	LiveWatchTimes        int    `gorm:"column:live_watch_times;" json:"live_watch_times" form:"live_watch_times"`
	FirstLiveEnterTime    int    `gorm:"column:first_live_enter_time;" json:"first_live_enter_time" form:"first_live_enter_time"`
	FirstReplayEnterTime  int    `gorm:"column:first_replay_enter_time;" json:"first_replay_enter_time" form:"first_replay_enter_time"`
	ReplayWatchTime       int    `gorm:"column:replay_watch_time;" json:"replay_watch_time" form:"replay_watch_time"`
	ReplayWatchTimes      int    `gorm:"column:replay_watch_times;" json:"replay_watch_times" form:"replay_watch_times"`
	LastReplayWatchClient string `gorm:"column:last_replay_watch_client;" json:"last_replay_watch_client" form:"last_replay_watch_client"`
	LastReplayLoginCity   string `gorm:"column:last_replay_login_city;" json:"last_replay_login_city" form:"last_replay_login_city"`
	TotalWatchTime        int    `gorm:"column:total_watch_time;" json:"total_watch_time" form:"total_watch_time"`
}

//ActionLiveData 实体表需要返回的有限字段
var ActionLiveDataColumn = "action_data_id,live_platform_type,last_live_watch_client,last_live_leave_time,last_live_login_city,live_watch_time," +
	"live_watch_times,first_live_enter_time,first_replay_enter_time,replay_watch_time,replay_watch_times,last_replay_watch_client,last_replay_login_city,total_watch_time"

/*
var operatorTypeOneMap = []string{"=", ">", ">=", "<", "<=", "<>"}
var page = "1"
var pageSize = "20"
var totalCount = 0
*/
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

func (ad *ActionData) QueryByParams(c *ctxExt.Context) (list []ActionData, supplementData map[string]interface{}, err error) {
	/*
		err = sqlDB.Find(&list).Error
		return list, err
	*/
	var params ActionData //请求参数
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

	mapQuery = mapQueryGenerator(params, mapQuery, c)
	/*
		if params.MainId > 0 {
			mapQuery["main_id"] = params.MainId
		} else {
			err = errors.New("main_id不能为空")
			return list, err
		}

	*/
	//构造分表
	tableSegmentation := "action_data" + "_" + strconv.Itoa(params.MainId)
	liveTableSegmentation := "action_live_data" + "_" + strconv.Itoa(params.MainId)

	//创建一个查询
	tx := sqlDB.Table(tableSegmentation).Where(mapQuery) //Map查询
	/*
		tx = tx.Joins("join action_live_data_5618 on action_live_data_5618.action_data_id = action_data_5618.id").Where("action_live_data_5618.last_live_login_city LIKE ?", "%中国3%")
		err = tx.Select("*").Scan(&list).Error
		return list, err
	*/
	//主表根据操作符查询
	tx, err = operatorQueryGenerator(params, tx, c)
	if err != nil {
		return
	}

	//范围查询，该方法放置于join操作之前
	tx = inQueryGenerator(params, tx, c)

	//表left join操作
	tx = tx.Joins("left join " + liveTableSegmentation + " on " + liveTableSegmentation + ".action_data_id = " + tableSegmentation + ".id")

	tx = joinQueryGenerator(params, liveTableSegmentation, c, tx)

	//join操作的表根据操作符查询
	tx, err = joinOperatorQueryGenerator(params, liveTableSegmentation, tx, c)
	if err != nil {
		return
	}

	//内部or查询
	orClauseSql, err := orQueryGenerator(c, liveTableSegmentation)
	orClauseSql = strings.TrimRight(orClauseSql, "or ")
	if err != nil {
		return
	}

	tx = tx.Where(orClauseSql)

	//执行查询操作
	page = c.DefaultQuery("page", page)
	pageSize = c.DefaultQuery("pageSize", pageSize)
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	//计算偏移
	offSet := (pageInt - 1) * pageSizeInt

	//获取总数
	tx.Count(&totalCount)

	//排序处理
	orderBy = c.DefaultQuery("orderBy", orderBy)

	//err = tx.Select("*").Scan(&list).Error
	//err = tx.Select("*").Order(orderBy).Offset(offSet).Limit(pageSizeInt).Scan(&list).Error
	err = tx.Select(tableSegmentation + ".*, " + ActionLiveDataColumn).Order(orderBy).Offset(offSet).Limit(pageSizeInt).Scan(&list).Error

	supplementData["page"] = pageInt
	supplementData["pageSize"] = pageSizeInt
	supplementData["totalCount"] = totalCount
	/*
		fmt.Println("sD: ", supplementData)
		fmt.Println("list: ", list)
	*/
	return list, supplementData, err
}

//mapQueryGenerator构造mapQuery对象
func mapQueryGenerator(params ActionData, mapQuery map[string]interface{}, c *ctxExt.Context) map[string]interface{} {
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
	if params.MemberId > 0 {
		mapQuery["member_id"] = params.MemberId
	}
	if params.WxOpenId != "" {
		mapQuery["wx_open_id"] = params.WxOpenId
	}
	if params.ClientIp != "" {
		mapQuery["client_ip"] = params.ClientIp
	}
	/*
		if params.UserIdentityType != "" {
			mapQuery["user_identity_type"] = params.UserIdentityType
		}
	*/
	if params.FirstVisitBrowser != "" {
		mapQuery["first_visit_browser"] = params.FirstVisitBrowser
	}
	if params.FirstVisitClient > 0 {
		mapQuery["first_visit_client"] = params.FirstVisitClient
	}
	/*
		if params.FirstVisitEquipment != "" {
			mapQuery["first_visit_equipment"] = params.FirstVisitEquipment
		}
	*/
	if params.FirstVisitChannelId > 0 {
		mapQuery["first_visit_channel_id"] = params.FirstVisitChannelId
	}
	if params.FollowChannelId > 0 {
		mapQuery["follow_channel_id"] = params.FollowChannelId
	}
	if params.EnrollChannelId > 0 {
		mapQuery["enroll_channel_id"] = params.EnrollChannelId
	}
	if params.WxFollowInviterId > 0 {
		mapQuery["wx_follow_inviter_id"] = params.WxFollowInviterId
	}
	if params.MeetingEnrollInviterId > 0 {
		mapQuery["meeting_enroll_inviter_id"] = params.MeetingEnrollInviterId
	}
	//验证是否有参数is_enroll请求进来
	IsEnroll := c.Query("is_enroll")
	if params.IsEnroll != nil && *params.IsEnroll >= 0 && IsEnroll != "" {
		mapQuery["is_enroll"] = *params.IsEnroll
	}
	if params.EnrollType > 0 {
		mapQuery["enroll_type"] = params.EnrollType
	}
	if params.EnrollWay > 0 {
		mapQuery["enroll_way"] = params.EnrollWay
	}
	if params.EnrollMeetingStatus > 0 {
		mapQuery["enroll_meeting_status"] = params.EnrollMeetingStatus
	}
	//验证是否有参数enroll_approve_status请求进来
	EnrollMeetingStatus := c.Query("enroll_approve_status")
	if params.EnrollMeetingStatus >= 0 && EnrollMeetingStatus != "" {
		mapQuery["enroll_approve_status"] = params.EnrollMeetingStatus
	}
	//验证是否有参数is_sign请求进来
	IsSign := c.Query("is_sign")
	if params.IsSign != nil && *params.IsSign >= 0 && IsSign != "" {
		mapQuery["is_sign"] = *params.IsSign
	}
	//验证是否有参数is_new_fans请求进来
	IsNewFans := c.Query("is_new_fans")
	if params.IsNewFans != nil && *params.IsNewFans >= 0 && IsNewFans != "" {
		mapQuery["is_new_fans"] = *params.IsNewFans
	}
	//验证是否有参数is_new_member请求进来
	IsNewMember := c.Query("is_new_member")
	if params.IsNewMember != nil && *params.IsNewMember >= 0 && IsNewMember != "" {
		mapQuery["is_new_member"] = *params.IsNewMember
	}
	//验证是否有参数is_watch_live请求进来
	IsWatchLive := c.Query("is_watch_live")
	if params.IsWatchLive != nil && *params.IsWatchLive >= 0 && IsWatchLive != "" {
		mapQuery["is_watch_live"] = *params.IsWatchLive
	}
	//验证是否有参数is_watch_replay请求进来
	IsWatchReplay := c.Query("is_watch_replay")
	if params.IsWatchReplay != nil && *params.IsWatchReplay >= 0 && IsWatchReplay != "" {
		mapQuery["is_watch_replay"] = *params.IsWatchReplay
	}
	if params.SeatNumber != "" {
		mapQuery["seat_number"] = params.SeatNumber
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsClickEnrollButton := c.Query("is_click_enroll_button")
	if params.IsClickEnrollButton != nil && *params.IsClickEnrollButton >= 0 && IsClickEnrollButton != "" {
		mapQuery["is_click_enroll_button"] = *params.IsClickEnrollButton
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsClickPayButton := c.Query("is_click_pay_button")
	if params.IsClickPayButton != nil && *params.IsClickPayButton >= 0 && IsClickPayButton != "" {
		mapQuery["is_click_pay_button"] = *params.IsClickPayButton
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsClickGenerateFissionButton := c.Query("is_click_generate_fission_button")
	if params.IsClickGenerateFissionButton != nil && *params.IsClickGenerateFissionButton >= 0 && IsClickGenerateFissionButton != "" {
		mapQuery["is_click_generate_fission_button"] = *params.IsClickGenerateFissionButton
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsClickEnterLiveButton := c.Query("is_click_enter_live_button")
	if params.IsClickEnterLiveButton != nil && *params.IsClickEnterLiveButton >= 0 && IsClickEnterLiveButton != "" {
		mapQuery["is_click_enter_live_button"] = *params.IsClickEnterLiveButton
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsClickWatchReplayButton := c.Query("is_click_watch_replay_button")
	if params.IsClickWatchReplayButton != nil && *params.IsClickWatchReplayButton >= 0 && IsClickWatchReplayButton != "" {
		mapQuery["is_click_watch_replay_button"] = *params.IsClickWatchReplayButton
	}
	//验证是否有参数is_click_enroll_button请求进来
	IsDel := c.Query("is_del")
	if params.IsDel != nil && *params.IsDel >= 0 && IsDel != "" {
		mapQuery["is_del"] = *params.IsDel
	}
	return mapQuery
}

//operatorQueryGenerator构造基于操作符的查询
func operatorQueryGenerator(params ActionData, tx *gorm.DB, c *ctxExt.Context) (txNew *gorm.DB, err error) {

	//验证是否有参数pay_money请求进来
	/*
		PayMoney := c.Query("pay_money")
		payMoneyOperator := c.Query("pay_money_operator")
		fmt.Println("PayMoney,payMoneyOperator", PayMoney, payMoneyOperator)
		if params.PayMoney != nil && *params.PayMoney >= 0 && PayMoney != "" && payMoneyOperator != "" {
			if isPermittedExpression(payMoneyOperator, operatorTypeOneMap) {
				tx = tx.Where("pay_money "+payMoneyOperator+"  ?", *params.PayMoney)
			} else if isPermittedExpression(payMoneyOperator, operatorTypeTwo) {
				PayMoneyLow := c.Query("pay_money_low")
				PayMoneyHigh := c.Query("pay_money_high")
				if PayMoneyLow == "" || PayMoneyHigh == "" {
					fmt.Println("here", PayMoneyLow, PayMoneyHigh)
					err = errors.New("pay_money_low或pay_money_high为空")
					return tx, err
				}
				tx = tx.Where("pay_money "+operatorTypeTwoMap[payMoneyOperator][0]+"  ?", *params.PayMoney).Where("pay_money "+operatorTypeTwoMap[payMoneyOperator][1]+"  ?", *params.PayMoney)
			} else {
				err = errors.New("invalid operator")
				return tx, err
			}
		}
		return tx, err

	*/
	/*
		PayMoney := c.Query("pay_money")
		payMoneyOperator := c.Query("pay_money_operator")
		fmt.Println("PayMoney,payMoneyOperator", PayMoney, payMoneyOperator)
		if params.PayMoney != nil && *params.PayMoney >= 0 && PayMoney != "" && payMoneyOperator != "" {
			tx, err = operatorQueryAbstract(tx, c, "pay_money", payMoneyOperator, *params.PayMoney)
			if err != nil {
				return tx, err
			}
		}
	*/

	var PayMoney = 0
	if params.PayMoney != nil {
		PayMoney = *params.PayMoney
	}
	payMoneyOperator := c.Query("pay_money_operator")
	if payMoneyOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "pay_money", payMoneyOperator, PayMoney)
		if err != nil {
			return tx, err
		}
	}

	var PosterInviteFollowNum = 0
	if params.PosterInviteFollowNum != nil {
		PosterInviteFollowNum = *params.PosterInviteFollowNum
	}
	PosterInviteFollowNumOperator := c.Query("poster_invite_follow_num_operator")
	if PosterInviteFollowNumOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "poster_invite_follow_num", PosterInviteFollowNumOperator, PosterInviteFollowNum)
		if err != nil {
			return tx, err
		}
	}

	EnrollTimeOperator := c.Query("enroll_time_operator")
	if EnrollTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "enroll_time", EnrollTimeOperator, params.EnrollTime)
		if err != nil {
			return tx, err
		}
	}

	CreatedAtOperator := c.Query("created_at_operator")
	if CreatedAtOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "created_at", CreatedAtOperator, params.CreatedAt)
		if err != nil {
			return tx, err
		}
	}

	UpdatedAtOperator := c.Query("updated_at_operator")
	if UpdatedAtOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "updated_at", UpdatedAtOperator, params.UpdatedAt)
		if err != nil {
			return tx, err
		}
	}

	var VisitTimes = 0
	if params.VisitTimes != nil {
		VisitTimes = *params.VisitTimes
	}
	VisitTimesOperator := c.Query("visit_times_operator")
	if VisitTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "visit_times", VisitTimesOperator, VisitTimes)
		if err != nil {
			return tx, err
		}
	}

	var DownloadTimes = 0
	if params.DownloadTimes != nil {
		DownloadTimes = *params.DownloadTimes
	}
	DownloadTimesOperator := c.Query("download_times_operator")
	if DownloadTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "download_times", DownloadTimesOperator, DownloadTimes)
		if err != nil {
			return tx, err
		}
	}

	var ShareTimes = 0
	if params.ShareTimes != nil {
		ShareTimes = *params.ShareTimes
	}
	ShareTimesOperator := c.Query("share_times_operator")
	if ShareTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "share_times", ShareTimesOperator, ShareTimes)
		if err != nil {
			return tx, err
		}
	}

	var FavoriteTimes = 0
	if params.FavoriteTimes != nil {
		FavoriteTimes = *params.FavoriteTimes
	}
	FavoriteTimesOperator := c.Query("favorite_times_operator")
	if FavoriteTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, "favorite_times", FavoriteTimesOperator, FavoriteTimes)
		if err != nil {
			return tx, err
		}
	}

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

	return tx, err
}

//join查询
func joinQueryGenerator(params ActionData, liveTableSegmentation string, c *ctxExt.Context, tx *gorm.DB) *gorm.DB {

	//关联表查询
	LivePlatformType := c.Query("live_platform_type")
	if LivePlatformType != "" {
		tx = tx.Where(liveTableSegmentation+".live_platform_type = ?", LivePlatformType)
	}
	LastLiveWatchClient := c.Query("last_live_watch_client")
	if LastLiveWatchClient != "" {
		//fmt.Println("LastLiveLoginCity: ", LastLiveLoginCity)
		tx = tx.Where(liveTableSegmentation+".last_live_watch_client LIKE ?", "%"+LastLiveWatchClient+"%")
	}
	//根据操作符查询
	/*
		LastLiveLeaveTimeOperator := c.Query("last_live_leave_time_operator")
		if params.LastLiveLeaveTime >= 0 && LastLiveLeaveTimeOperator != "" && isPermittedExpression(LastLiveLeaveTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".last_live_leave_time "+LastLiveLeaveTimeOperator+"  ?", params.LastLiveLeaveTime)
		}
	*/
	LastLiveLoginCity := c.Query("last_live_login_city")
	if LastLiveLoginCity != "" {
		//fmt.Println("LastLiveLoginCity: ", LastLiveLoginCity)
		tx = tx.Where(liveTableSegmentation+".last_live_login_city LIKE ?", "%"+LastLiveLoginCity+"%")
	}
	/*
		LiveWatchTimeOperator := c.Query("live_watch_time_operator")
		if params.LiveWatchTime >= 0 && LiveWatchTimeOperator != "" && isPermittedExpression(LiveWatchTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".live_watch_time "+LiveWatchTimeOperator+"  ?", params.LiveWatchTime)
		}
		LiveWatchTimesOperator := c.Query("live_watch_times_operator")
		if params.LiveWatchTimes >= 0 && LiveWatchTimesOperator != "" && isPermittedExpression(LiveWatchTimesOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".live_watch_times "+LiveWatchTimesOperator+"  ?", params.LiveWatchTimes)
		}
		FirstLiveEnterTimeOperator := c.Query("first_live_enter_time_operator")
		if params.FirstLiveEnterTime >= 0 && FirstLiveEnterTimeOperator != "" && isPermittedExpression(FirstLiveEnterTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".first_live_enter_time "+FirstLiveEnterTimeOperator+"  ?", params.FirstLiveEnterTime)
		}
		FirstReplayEnterTimeOperator := c.Query("first_replay_enter_time_operator")
		if params.FirstReplayEnterTime >= 0 && FirstReplayEnterTimeOperator != "" && isPermittedExpression(FirstReplayEnterTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".first_replay_enter_time "+FirstReplayEnterTimeOperator+"  ?", params.FirstReplayEnterTime)
		}
		ReplayWatchTimeOperator := c.Query("replay_watch_time_operator")
		if params.ReplayWatchTime >= 0 && ReplayWatchTimeOperator != "" && isPermittedExpression(ReplayWatchTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".replay_watch_time "+ReplayWatchTimeOperator+"  ?", params.ReplayWatchTime)
		}
		ReplayWatchTimesOperator := c.Query("replay_watch_times_operator")
		if params.ReplayWatchTimes >= 0 && ReplayWatchTimesOperator != "" && isPermittedExpression(ReplayWatchTimesOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".replay_watch_times "+ReplayWatchTimesOperator+"  ?", params.ReplayWatchTimes)
		}
	*/
	LastReplayWatchClient := c.Query("last_replay_watch_client")
	if LastReplayWatchClient != "" {
		//fmt.Println("LastLiveLoginCity: ", LastLiveLoginCity)
		tx = tx.Where(liveTableSegmentation+".last_replay_watch_client LIKE ?", "%"+LastReplayWatchClient+"%")
	}
	LastReplayLoginCity := c.Query("last_replay_login_city")
	if LastReplayLoginCity != "" {
		//fmt.Println("LastLiveLoginCity: ", LastLiveLoginCity)
		tx = tx.Where(liveTableSegmentation+".last_replay_login_city LIKE ?", "%"+LastReplayLoginCity+"%")
	}
	/*
		TotalWatchTimeOperator := c.Query("total_watch_time_operator")
		if params.TotalWatchTime >= 0 && TotalWatchTimeOperator != "" && isPermittedExpression(TotalWatchTimeOperator, operatorTypeOneMap) {
			tx = tx.Where(liveTableSegmentation+".total_watch_time "+TotalWatchTimeOperator+"  ?", params.TotalWatchTime)
		}
	*/
	return tx
}

//inQueryGenerator 构造in查询
func inQueryGenerator(params ActionData, tx *gorm.DB, c *ctxExt.Context) *gorm.DB {
	//member_id_list为member_id以逗号分割的字符串
	MemberIdList := c.Query("member_id_list")
	if MemberIdList != "" {
		MemberIdArr := strings.Split(MemberIdList, ",")
		if len(MemberIdArr) != 0 {
			fmt.Println("MemberIdArr: ", MemberIdArr)
			//fmt.Println("member_id_arr: ", []int{29, 30})
			tx = tx.Where("member_id IN (?) ", MemberIdArr)
		}
	}

	return tx
}

//joinOperatorQueryGenerator，提供join操作的表名，构造基于操作符的查询
func joinOperatorQueryGenerator(params ActionData, liveTableSegmentation string, tx *gorm.DB, c *ctxExt.Context) (txNew *gorm.DB, err error) {

	LastLiveLeaveTimeOperator := c.Query("last_live_leave_time_operator")
	if LastLiveLeaveTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".last_live_leave_time", LastLiveLeaveTimeOperator, params.LastLiveLeaveTime)
		if err != nil {
			return tx, err
		}
	}

	LiveWatchTimeOperator := c.Query("live_watch_time_operator")
	if LiveWatchTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".last_live_leave_time", LiveWatchTimeOperator, params.LiveWatchTime)
		if err != nil {
			return tx, err
		}
	}

	LiveWatchTimesOperator := c.Query("live_watch_times_operator")
	if LiveWatchTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".live_watch_times", LiveWatchTimesOperator, params.LiveWatchTimes)
		if err != nil {
			return tx, err
		}
	}

	FirstLiveEnterTimeOperator := c.Query("first_live_enter_time_operator")
	if FirstLiveEnterTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".first_live_enter_time", FirstLiveEnterTimeOperator, params.FirstLiveEnterTime)
		if err != nil {
			return tx, err
		}
	}

	FirstReplayEnterTimeOperator := c.Query("first_replay_enter_time_operator")
	if FirstReplayEnterTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".first_replay_enter_time", FirstReplayEnterTimeOperator, params.FirstReplayEnterTime)
		if err != nil {
			return tx, err
		}
	}

	ReplayWatchTimeOperator := c.Query("replay_watch_time_operator")
	if ReplayWatchTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".replay_watch_time", ReplayWatchTimeOperator, params.ReplayWatchTime)
		if err != nil {
			return tx, err
		}
	}

	ReplayWatchTimesOperator := c.Query("replay_watch_times_operator")
	if ReplayWatchTimesOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".replay_watch_times", ReplayWatchTimesOperator, params.ReplayWatchTimes)
		if err != nil {
			return tx, err
		}
	}

	TotalWatchTimeOperator := c.Query("total_watch_time_operator")
	if TotalWatchTimeOperator != "" {
		tx, err = operatorQueryAbstract(tx, c, liveTableSegmentation+".total_watch_time", TotalWatchTimeOperator, params.TotalWatchTime)
		if err != nil {
			return tx, err
		}
	}
	return tx, err
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

//orQueryGenerator执行内部or查询
func orQueryGenerator(c *ctxExt.Context, liveTableSegmentation string) (orClauseSql string, err error) {
	orClause := c.QueryMap("orClause")
	//循环map
	for _, value := range orClause {
		var valueArr = strings.Split(value, ",")
		//操作符
		if valueArr[0] == "" {
			err := errors.New("操作符为空")
			return orClauseSql, err
		}
		if isPermittedExpression(valueArr[0], operatorTypeOneMap) {
			if len(valueArr) < 3 {
				err := errors.New("操作值为空")
				return orClauseSql, err
			}
		} else if isStringInSlice(valueArr[0], operatorTypeTwo) || isStringInSlice(valueArr[0], operatorTypeThree) {
			if len(valueArr) < 4 {
				err := errors.New("操作值不足")
				return orClauseSql, err
			}
		}
		//操作字段
		fieldName := valueArr[1]
		//判断字段所在的数据表
		ActionLiveDataColumnArr := strings.Split(ActionLiveDataColumn, "，")
		if isStringInSlice(fieldName, ActionLiveDataColumnArr) {
			fieldName = liveTableSegmentation + "." + fieldName
		}
		//构造条件语句
		orWhere, err := operatorQueryAbstractInner(fieldName, valueArr[0], valueArr)
		if err != nil {
			return orClauseSql, err
		}
		orClauseSql += orWhere

	}
	return orClauseSql, err
}

//operatorQueryAbstract 抽象特殊查询操作符
func operatorQueryAbstractInner(fieldName, operator string, operatorArr []string) (orWhere string, err error) {
	//fmt.Println("operator", operator)
	if isPermittedExpression(operator, operatorTypeOneMap) {
		//fmt.Println("isPermittedExpression", "ok")
		temp := fieldName + " " + operatorTypeOneMap[operator] + " '" + html.EscapeString(operatorArr[2]) + "' or "
		orWhere += temp
		//tx = tx.Where(fieldName+" "+operatorTypeOneMap[operator]+" ?", operatorArr[2])
	} else if isStringInSlice(operator, operatorTypeTwo) {
		//fmt.Println("operatorTypeTwoMap", operatorTypeTwoMap[operator][0], operatorTypeTwoMap[operator][1])
		temp := "(" + fieldName + " " + operatorTypeTwoMap[operator][0] + " '" + html.EscapeString(operatorArr[2]) + "' and " + fieldName + " " + operatorTypeTwoMap[operator][1] + " '" + html.EscapeString(operatorArr[3]) + "') or "
		orWhere += temp
		//tx = tx.Where(fieldName+" "+operatorTypeTwoMap[operator][0]+"  ?", operatorArr[2]).Where(fieldName+" "+operatorTypeTwoMap[operator][1]+"  ?", operatorArr[3])
	} else if isStringInSlice(operator, operatorTypeThree) {
		//fmt.Println("operatorTypeThree", operatorTypeThreeMap[operator][0], operatorTypeThreeMap[operator][1])
		temp := "(" + fieldName + " " + operatorTypeThreeMap[operator][0] + " '" + html.EscapeString(operatorArr[2]) + "' and " + fieldName + " " + operatorTypeThreeMap[operator][1] + " '" + html.EscapeString(operatorArr[3]) + "') or "
		orWhere += temp
		//tx = tx.Where(fieldName+" "+operatorTypeThreeMap[operator][0]+"  ?", operatorArr[2]).Where(fieldName+" "+operatorTypeThreeMap[operator][1]+"  ?", operatorArr[3])
	} else {
		err := errors.New("operatorQueryAbstractInner:invalid operator," + operator)
		return orWhere, err
	}
	return orWhere, err
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
