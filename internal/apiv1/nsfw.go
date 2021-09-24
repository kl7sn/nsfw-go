package apiv1

import (
	"github.com/gin-gonic/gin"

	"github.com/kl7sn/nsfw-go/internal/service"
)

func Nsfw(c *gin.Context) {
	url := c.Query("url") // 图片 url
	_, err := service.Download(url)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	nsfw := service.NewNsfwService()
	nsfw.Build("./pkg/data/open_nsfw-weights.npy", service.TENSOR)
	c.JSON(200, "Hello EGO")
}
