package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowVideoService 投稿详情的服务
type ShowVideoService struct {
}

// Show 展示视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(&video),
	}
}
