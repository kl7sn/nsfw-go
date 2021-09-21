package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/kl7sn/nsfw-go/internal/service"
)

func Nsfw(c *gin.Context) {

	url := c.Query("url") // 图片 url

	imageBytes, err := service.Download(url)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}

	c.JSON(200, "Hello EGO")
}
