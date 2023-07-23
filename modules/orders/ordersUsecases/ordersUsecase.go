package ordersUsecases

import (
	"github.com/jamee5e/jame-shop-tutorial/modules/orders"
	"github.com/jamee5e/jame-shop-tutorial/modules/orders/ordersRepositories"
	"github.com/jamee5e/jame-shop-tutorial/modules/products/productsRepositories"
)

type IOrdersUsecase interface {
	FindOneOrder(orderId string) (*orders.Order, error)
}

type ordersUsecase struct {
	ordersRepository   ordersRepositories.IOrdersRepository
	productsRepository productsRepositories.IProductsRepository
}

func OrdersUsecase(ordersRepository ordersRepositories.IOrdersRepository, productsRepository productsRepositories.IProductsRepository) IOrdersUsecase {
	return &ordersUsecase{
		ordersRepository:   ordersRepository,
		productsRepository: productsRepository,
	}
}

func (u *ordersUsecase) FindOneOrder(orderId string) (*orders.Order, error) {
	order, err := u.ordersRepository.FindOneOrder(orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}
