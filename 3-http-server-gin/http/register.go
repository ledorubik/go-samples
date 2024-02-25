package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterHttpEndpoints(logger *logrus.Logger, router *gin.Engine) {
	h := NewHandler(logger)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			echo := v1.Group("/echo")
			{
				echo.POST("", h.Post)
			}
		}
	}
}
