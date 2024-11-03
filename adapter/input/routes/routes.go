package routes

import (
	"github.com/GiovannaK/go-hexagonal-api.git/adapter/input/controller"
	"github.com/GiovannaK/go-hexagonal-api.git/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsService := service.NewNewsService()
	newsController := controller.NewNewsController(newsService)
	r.GET("/news", newsController.GetNews)
}
