package middlewaresUsecases

import (
	"github.com/jamee5e/jame-shop-tutorial/modules/middlewares"
	"github.com/jamee5e/jame-shop-tutorial/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsercase interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
}
type middlewaresUsercase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewareRepository
}

func MiddlewaresUsercase(middlewaresRepository middlewaresRepositories.IMiddlewareRepository) IMiddlewaresUsercase {
	return &middlewaresUsercase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsercase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (u *middlewaresUsercase) FindRole() ([]*middlewares.Role, error) {
	roles, err := u.middlewaresRepository.FindRole()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
