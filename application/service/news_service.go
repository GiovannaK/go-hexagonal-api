package service

import (
	"fmt"

	"github.com/GiovannaK/go-hexagonal-api.git/application/domain"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/logger"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/rest_err"
)

type newsService struct{}

func NewNewsService() *newsService {
	return &newsService{}
}

func (ns *newsService) GetNewsService(
	subject string,
	from string,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("GetNewsService called with subject: %s and from: %s", subject, from))

	return nil, nil
}
