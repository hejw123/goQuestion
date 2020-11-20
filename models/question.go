package models

import (
	"question/mysql"
)

type Question struct {
	Id       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Question string `json:"question" gorm:"column:question"`
	Type     int    `json:"type" gorm:"column:type"`
	Class    int    `json:"class" gorm:"class"`
}

func (q *Question) TableName() string {
	return "question"
}

/** 根据课时查询 */
func GetClassQuestion(class , t int) ([]Question, error) {
	Q := []Question{}
	d := mysql.DB.Where(" `class` = ? and `type` = ? ", class, t).Find(&Q)
	return Q, d.Error
}

/** 根据id 修改对应的状态 */
func UpdateQuestion(id, t int) {
	mysql.DB.Table("question").Where(" `id` = ? ", id).Update(" type = ? ", t)
}

/** 添加问题*/
func InsertQuestion(q *Question) {
	mysql.DB.Table("question").Create(q)
}
