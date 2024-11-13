package repository

import (
	"github.com/zhiqinkuang/easyAPI-gin/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 定义一个数据表结构体

type Post struct {
	Id         int64     `gorm:"column:id"`
	ParentId   int64     `gorm:"parent_id"`
	UserId     int64     `gorm:"column:user_id"`
	Content    string    `gorm:"column:content"`
	DiggCount  int32     `gorm:"column:digg_count"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Post) TableName() string {
	return "post"
}

type PostDao struct {
}

// 建立一个Dao 对象
var postDao *PostDao
var postOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

// 根据id 查询post id 记录
func (*PostDao) QueryPostById(id int64) (*Post, error) {
	var post Post
	err := db.Where("id = ?", id).Find(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find post by id err:" + err.Error())
		return nil, err
	}
	return &post, nil
}

// 根据parentid 查询
func (*PostDao) QueryPostByParentId(parentId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("parent_id = ?", parentId).Find(&posts).Error
	if err != nil {
		util.Logger.Error("find posts by parent_id err:" + err.Error())
		return nil, err
	}
	return posts, nil
}

// 创建帖子的记录
func (*PostDao) CreatePost(post *Post) error {
	if err := db.Create(post).Error; err != nil {
		util.Logger.Error("insert post err:" + err.Error())
		return err
	}
	return nil
}
