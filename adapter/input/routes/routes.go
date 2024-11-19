package routes

import (
	"github.com/GiovannaK/go-hexagonal-api.git/adapter/input/controller"
	newshttp "github.com/GiovannaK/go-hexagonal-api.git/adapter/output/news_http"
	"github.com/GiovannaK/go-hexagonal-api.git/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := newshttp.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)
	r.GET("/news", newsController.GetNews)
}
