package service

import (
	"github.com/zhiqinkuang/easyAPI-gin/repository"
	"github.com/zhiqinkuang/easyAPI-gin/util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := repository.Init0(); err != nil {
		os.Exit(1)
	}
	if err := util.InitLogger(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestPublishPost(t *testing.T) {
	//  带有参数的结构体
	type args struct {
		topicId int64
		userId  int64
		content string
	}
	// 测试用例
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发布回帖",
			args: args{
				topicId: 1,
				userId:  2,
				content: "再次回帖",
			},
			wantErr: false,
		},
	}
	// 测试发布函数功能
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PublishPost(tt.args.topicId, tt.args.userId, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
