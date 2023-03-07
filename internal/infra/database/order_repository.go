package database

import (
	"database/sql"
	"github.com/rafaeelneto/taxCalculator/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository (db *sql.DB) *OrderRepository{
	return &OrderRepository{Db: db}
}

func (oR *OrderRepository) Save(order *entity.Order) error {
	_, err := oR.Db.Exec("Insert into orders (id, price, tax, quantity) value (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.Quantity)

	if err != nil{
		return err
	}
	return nil
}

func (oR *OrderRepository) GetTotal(order *entity.Order) (int, error) {
	var total int
	err := oR.Db.QueryRow("Select count (*) from orders").Scan(&total)

	if err != nil{
		return 0, err
	}

	return total, nil


}