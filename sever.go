package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easyAPI-gin/handler"
	"github.com/zhiqinkuang/easyAPI-gin/repository"
	"github.com/zhiqinkuang/easyAPI-gin/util"
	"net/http"
	"os"
)

func main() {
	// 定义router
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 传入topicId
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := handler.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	// 传入topicId,uid
	r.POST("/community/post/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := handler.PublishPost(uid, topicId, content)
		c.JSON(200, data)
	})

	err := r.Run()
	if err != nil {
		return
	}
}

// 初始化全局变量,出错就返回报错
func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
