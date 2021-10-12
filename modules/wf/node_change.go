package wf

import (
	"fmt"
	"meigo/library/log"
	wfMod "meigo/models/wf"
	"net/http"

	"github.com/gin-gonic/gin"

	ctxExt "github.com/kinjew/gin-context-ext"
)

//var node wfMod.Node
var node wfMod.Node

/*
QueryByParams 处理Node信息变更问题
*/
func QueryByParams(c *ctxExt.Context) {

	if flag, err := node.ArgoYaml(c); err != nil {
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
