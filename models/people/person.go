package people

import (
	"meigo/library/db/common"

	"meigo/library/db"

	"meigo/library/log"

	"github.com/jinzhu/gorm"
	ctxExt "github.com/kinjew/gin-context-ext"
)

// Person 实体
type Person struct {
	common.BaseModel
	FirstName string `json:"first_name" form:"first_name" gorm:"type:varchar(50);unique_index:first_name_city"`
	LastName  string `json:"last_name" form:"last_name" gorm:"type:varchar(50);"`
	City      string `json:"city" form:"city" gorm:"type:varchar(50);unique_index:first_name_city"`
}

// sqlDB 是 *gorm.DB
var sqlDB *gorm.DB
var err error

// InitPersonDB 初始化数据库
func InitPersonDB() {

	if sqlDB, err = db.ConnDB("test"); err != nil {
		panic(err)
	}
	sqlDB.AutoMigrate(&Person{})
}

/*
DeletePerson 删除人员
*/
func (p *Person) DeletePerson(id string) (person Person, err error) {
	err = sqlDB.Where("id = ?", id).Delete(&person).Error
	return person, err
}

/*
UpdatePerson 更新人员
*/
func (p *Person) UpdatePerson(c *ctxExt.Context) (person Person, err error) {
	id := c.Params.ByName("id")
	if err := sqlDB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	}
	if err := c.BindJSON(&person); err != nil {
		log.Error("BindJSON err: ", err)
	}
	err = sqlDB.Save(&person).Error
	return person, err
}

/*
CreatePerson 创建人员
*/
func (p *Person) CreatePerson(c *ctxExt.Context) (person Person, err error) {
	if err := c.BindJSON(&person); err != nil {
		log.Error("BindJSON err: ", err)
	}
	err = sqlDB.Create(&person).Error
	//return person, errors.New("test")
	return person, err
}

/*
GetPerson 获取人员
*/
func GetPerson(id string) (person Person, err error) {
	err = sqlDB.Where("id = ?", id).First(&person).Error
	return person, err
}

/*
GetPeople 获取人员列表
*/
func (p *Person) GetPeople() (people []Person, err error) {
	err = sqlDB.Find(&people).Error
	return people, err
}

// Close 关闭数据库
func Close() {
	db.Close(sqlDB)
}
