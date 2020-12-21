package marketing_tool_data

import (
	"fmt"
	"meigo/library/log"
	mtdMod "meigo/models/marketing_tool_data"
	"net/http"

	"github.com/gin-gonic/gin"

	ctxExt "github.com/kinjew/gin-context-ext"
)

var gsat mtdMod.GeneralStatAggregationData

/*
GsatQueryByParams 获取GeneralStatAggregationData列表
*/
func GsatQueryByParams(c *ctxExt.Context) {

	if list, supplementData, err := gsat.GsatQueryByParams(c); err != nil {
		//错误返回
		w := fmt.Errorf("modules_marketing_tool_data:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), list)

	} else {
		//sugar.Infow("GetPeople_Output", "people", people, "time", time.Now().Local().String())
		log.Info("ActionData_Output", list)
		log.Info("ActionData_Output", supplementData)
		//c.Success(http.StatusOK, "succ", list)
		c.JSON(200, gin.H{
			"ret":            1,
			"code":           200,
			"msg":            "succ",
			"data":           list,
			"supplementData": supplementData,
		})

	}
}
