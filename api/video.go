package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	services := service.CreatVideoService{}
	if err := c.ShouldBind(&services); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := services.Create()
		c.JSON(200, res)
	}
}
