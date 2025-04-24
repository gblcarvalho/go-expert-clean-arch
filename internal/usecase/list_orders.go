package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

func (uc *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	output := ListOrdersOutputDTO{}
	orders, err := uc.OrderRepository.List()
	if err != nil {
		return output, err
	}

	var ordersDTOs []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersDTOs = append(ordersDTOs, dto)
	}
	output.Orders = ordersDTOs

	return output, nil
}
