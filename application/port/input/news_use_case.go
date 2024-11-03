package input

import (
	"github.com/GiovannaK/go-hexagonal-api.git/application/domain"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(
		subject string,
		from string,
	) (*domain.NewsDomain, *rest_err.RestErr)
}
