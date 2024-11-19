package service

import (
	"fmt"

	"github.com/GiovannaK/go-hexagonal-api.git/application/domain"
	"github.com/GiovannaK/go-hexagonal-api.git/application/port/output"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/logger"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(
	new domain.NewsReqDomain,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("GetNewsService called with subject: %s and from: %s", new.Subject, new.From))

	newsDomainResponse, err := ns.newsPort.GetNewsPort(new)
	return newsDomainResponse, err
}
