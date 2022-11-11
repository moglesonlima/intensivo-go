package usecase

import (
	"github.com/moglesonlima/gointensivo/internal/order/entity"
	"github.com/moglesonlima/gointensivo/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	// OrderRepository database.OrderRepository // Mega acoplado
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		//OrderRepository: orderRepository,
		OrderRepository: &orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	// Cria a ordem
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	// Se existir algum erro, retorna o erro
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	// Salva a ordem no banco de dados
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	// Cria o output
	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil

}
