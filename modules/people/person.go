package people

import (
	"fmt"
	"meigo/library/log"
	peopleMod "meigo/models/people"
	"net/http"
	"os"

	ctxExt "github.com/kinjew/gin-context-ext"
)

var p peopleMod.Person

/*
DeletePerson 删除人员
*/
func DeletePerson(c *ctxExt.Context) {
	id := c.Params.ByName("id")
	if person, err := p.DeletePerson(id); err != nil {
		//错误返回
		w := fmt.Errorf("modules_person:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), person)
	} else {
		c.Success(http.StatusOK, "succ", person)
	}
}

/*
UpdatePerson 更新人员
*/
func UpdatePerson(c *ctxExt.Context) {
	if person, err := p.UpdatePerson(c); err != nil {
		//错误返回
		w := fmt.Errorf("modules_person:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), person)
	} else {
		c.Success(http.StatusOK, "succ", person)
	}
}

/*
CreatePerson 创建人员
*/
func CreatePerson(c *ctxExt.Context) {
	if person, err := p.CreatePerson(c); err != nil {
		//错误返回
		w := fmt.Errorf("modules_person:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), person)
	} else {
		c.Success(http.StatusOK, "succ", person)
	}
}

/*
GetPerson 获取人员
*/
func GetPerson(c *ctxExt.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)
	if person, err := peopleMod.GetPerson(id); err != nil {
		//错误返回
		w := fmt.Errorf("modules_person:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), person)
	} else {
		c.Success(http.StatusOK, "succ", person)
	}
}

/*
GetPeople 获取人员列表
*/
func GetPeople(c *ctxExt.Context) {
	if people, err := p.GetPeople(); err != nil {
		//错误返回
		w := fmt.Errorf("modules_person:%w", err)
		c.Error(http.StatusBadRequest, w.Error(), people)

	} else {
		//sugar.Infow("GetPeople_Output", "people", people, "time", time.Now().Local().String())
		log.Info("GetPeople_Output", people)
		c.Success(http.StatusOK, "succ", people)
	}
}

/*
ConsoleGetPeople 命令行工具获取人员列表
*/
func ConsoleGetPeople() {
	if people, err := p.GetPeople(); err == nil {
		fmt.Println(people)

	}
	os.Exit(0)
}
