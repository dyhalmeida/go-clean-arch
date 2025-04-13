package usecase

import "github.com/dyhalmeida/go-clean-arch/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return ordersDTO, nil
}
