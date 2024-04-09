package usecase

import (
	"github.com/mateus-sousa/fc-clean-architecture/internal/entity"
)

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetAllOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var ordersDTO []OrderOutputDTO
	for _, o := range orders {
		var orderDTO OrderOutputDTO
		orderDTO.ID = o.ID
		orderDTO.Tax = o.Tax
		orderDTO.Price = o.Price
		orderDTO.FinalPrice = o.FinalPrice
		ordersDTO = append(ordersDTO, orderDTO)
	}
	return ordersDTO, nil
}
