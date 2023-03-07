package entity

import (
	"errors"
)

type Order struct {
	ID string
	Price float64
	Tax float64
	Quantity int
	FinalPrice float64
}

func NewOrder (id string, price float64, tax float64, quantity int) (*Order, error){

	order := &Order{ID: id, Price: price, Tax: tax, Quantity: quantity, FinalPrice: (price+tax)*float64(quantity)}

	err := order.Validate();

	if err != nil {
		return nil, err
		
	}

	return order, nil;
}

func (o *Order) Validate() error {
	switch true {
		case o.ID == "":
			return errors.New("ID is required")
		case o.Price <= 0:
			return errors.New("Price is required")
		case o.Tax <= 0:
			return errors.New("Tax is required")
		case o.Quantity <= 0:
			return errors.New("Quantity is required")
	}

	return nil // null
}