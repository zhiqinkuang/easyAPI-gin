package repository

import (
	"github.com/zhiqinkuang/easyAPI-gin/util"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao struct {
}

// 建立全局变量
var topicDao *TopicDao
var topicOnce sync.Once

// 单例创建Dao对象
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

// 根据 usermap
func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	// 结构体对象
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		util.Logger.Error("find topic by id err:" + err.Error())
		return nil, err
	}
	return &topic, nil
}
