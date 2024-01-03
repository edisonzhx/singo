package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	services := service.CreateVideoService{}
	if err := c.ShouldBind(&services); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := services.Create()
		c.JSON(200, res)
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	services := service.ShowVideoService{}
	res := services.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	services := service.ListVideoService{}
	if err := c.ShouldBind(&services); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := services.List()
		c.JSON(200, res)
	}
}

// UpdateVideo 更新视频接口
func UpdateVideo(c *gin.Context) {
	services := service.UpdateVideoService{}
	if err := c.ShouldBind(&services); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := services.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

// DeleteVideo 删除视频的接口
func DeleteVideo(c *gin.Context) {
	services := service.DeleteVideoService{}
	res := services.Delete(c.Param("id"))
	c.JSON(200, res)
}
