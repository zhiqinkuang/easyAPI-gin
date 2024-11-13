package handler

import (
	"github.com/zhiqinkuang/easyAPI-gin/service"
	"strconv"
)

// 发布post
// string  uid topicId 和 具体内容
func PublishPost(uidStr, topicIdStr, content string) *PageData {
	//参数转换,将输入的 string 类型转换为int类型
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	topic, _ := strconv.ParseInt(topicIdStr, 10, 64)
	//获取service层结果
	postId, err := service.PublishPost(topic, uid, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}

}
