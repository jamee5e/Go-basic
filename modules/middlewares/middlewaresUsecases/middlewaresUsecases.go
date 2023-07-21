package middlewaresUsecases

import "github.com/jamee5e/jame-shop-tutorial/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsercase interface {
}
type middlewaresUsercase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewareRepository
}

func MiddlewaresUsercase(middlewaresRepository middlewaresRepositories.IMiddlewareRepository) IMiddlewaresUsercase {
	return &middlewaresUsercase{
		middlewaresRepository: middlewaresRepository,
	}
}
