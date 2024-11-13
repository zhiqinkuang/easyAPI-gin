package handler

import (
	"github.com/zhiqinkuang/easyAPI-gin/service"
	"strconv"
)

// 返回结构体
type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 根据请求结构体查询数据
// 传入topicId 获得这个page的数据
func QueryPageInfo(topicIdStr string) *PageData {
	//参数转换
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	//获取service层结果
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}

}
