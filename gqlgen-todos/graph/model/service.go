package model

import (
	"github.com/gin-gonic/gin"
)

/*
CreateTodo 创建todo
*/
func (t *Todo) CreateTodo(c *gin.Context, input NewTodo) (todo *Todo, err error) {

	todo = &Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID, // fix this line
	}
	err = sqlDB.Create(todo).Error
	//return person, errors.New("test")
	return todo, err
}

/*
FindTodo 查找todo
*/
func (t *Todo) FindTodo(c *gin.Context) (todo []*Todo, err error) {

	var todos []*Todo
	err = sqlDB.Table("todo").Scan(&todos).Error
	//return person, errors.New("test")
	return todos, err
}
