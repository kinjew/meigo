package wf

import (
	"fmt"
	"meigo/library/log"
	"net/http"

	wfMod "meigo/models/wf"

	"github.com/gin-gonic/gin"
	ctxExt "github.com/kinjew/gin-context-ext"
)

/*
TriggerProcess 处理流程触发
*/
func TriggerProcess(c *ctxExt.Context) {

	if flag, err := wfMod.TrigerProcess(c); err != nil {
		//错误返回
		w := fmt.Errorf("modules_wf:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), flag)

	} else {
		//sugar.Infow("GetPeople_Output", "people", people, "time", time.Now().Local().String())
		log.Info("Node_Output", flag)
		//c.Success(http.StatusOK, "succ", list)
		c.JSON(200, gin.H{
			"ret":  1,
			"code": 200,
			"msg":  "succ",
			"data": flag,
		})

	}
}
