package controller

import (
	"github.com/GiovannaK/go-hexagonal-api.git/adapter/input/model/request"
	"github.com/GiovannaK/go-hexagonal-api.git/application/port/input"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/logger"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *newsController {
	return &newsController{newsUseCase}
}

func (nc *newsController) GetNews(c *gin.Context) {
	//q=tesla&from=2021-08-01&to=2021-08-10&apiKey=""

	logger.Info("GetNews controller called")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to bind query", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain, err := nc.newsUseCase.GetNewsService(newsRequest.Subject, newsRequest.From)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, newsDomain)
}
