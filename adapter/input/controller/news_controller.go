package controller

import (
	"fmt"
	"net/http"

	"github.com/GiovannaK/go-hexagonal-api.git/adapter/input/model/request"
	"github.com/GiovannaK/go-hexagonal-api.git/application/domain"
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

	logger.Info("GetNews controller called")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to bind query", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println("Parsed From date: ", newsRequest.From)

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),                           
	}

	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
