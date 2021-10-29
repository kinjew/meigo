package wf

import (
	"fmt"
	"meigo/library/log"
	"os/exec"
	"strconv"
	"strings"

	ctxExt "github.com/kinjew/gin-context-ext"
)

/*
CronDeleteProcess 删除cronjob
*/

func CronDeleteProcess(c *ctxExt.Context) (flag bool, err error) {
	//flow_id与node_id不能同时为空
	FlowId := c.Query("flow_id")
	NodeId := c.Query("node_id")
	//fmt.Println(FlowId, NodeId)
	if FlowId == "" && NodeId == "" {
		log.Error("flow_id and noed_id are not allowed both null: ", err)
		return false, fmt.Errorf("flow_id and node_id are both null")
	}
	//获取执行定时任务的节点
	var flowNodes []Node
	if FlowId != "" {
		//flow_id为以逗号分隔的字符串
		FlowIdSlice := strings.Split(FlowId, ",")
		err = sqlDB.Table("flow_nodes").Where("flow_id in (?)", FlowIdSlice).Where("is_repeat = ?", 1).Select("* ").Scan(&flowNodes).Error
		if err != nil {
			return false, err
		}
	} else if NodeId != "" {
		//noed_id为以逗号分隔的字符串
		NodeIdSlice := strings.Split(NodeId, ",")
		err = sqlDB.Table("flow_nodes").Where("id in (?)", NodeIdSlice).Where("is_repeat = ?", 1).Select("* ").Scan(&flowNodes).Error
		if err != nil {
			return false, err
		}
	}
	//删除cron
	for _, item := range flowNodes {
		var cronName = "cron-wf-" + strconv.Itoa(item.FlowId) + "-" + strconv.Itoa(item.ID)
		//提交执行工作流
		//	cmd := exec.Command("/usr/local/bin/argo submit", fileName, "-n argo", "-p message="+Message)
		cmd := exec.Command("argo", "cron", "delete", cronName, "-n", "argo")
		_, err = cmd.Output()
		/*
			data, err := cmd.Output()
			fmt.Println(string(data))
		*/
		//判断执行结果
		if err != nil {
			return false, err
		}
	}
	// 返回结果
	return true, err
}
