package output

import (
	"github.com/GiovannaK/go-hexagonal-api.git/application/domain"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
