package database

import (
	"api-go/internal/entity"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
