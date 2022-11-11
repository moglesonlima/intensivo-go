package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	//Verifica se existe algum erro na criação da ordem
	err := order.isValid()
	// Se existir algum erro, retorna o erro
	if err != nil {
		return nil, err
	}
	// Se não existir nenhum erro, retorna a ordem
	return order, nil
}

func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("Invalid ID")
	}
	if o.Price <= 0 {
		return errors.New("Invalid Price")
	}
	if o.Tax <= 0 {
		return errors.New("Invalid Tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	//Verifica se existe algum erro no calculo do preço final
	err := o.isValid()
	if err != nil {
		return err
	}
	return nil
}
