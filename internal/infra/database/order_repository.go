package database

import (
	"database/sql"
	"errors"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List() ([]*entity.Order, error) {
	rows, err := r.Db.Query("select id, price, tax, final_price from orders")
	if err != nil {
		return nil, err
	}
	var orders []*entity.Order
	for rows.Next() {
		o, err := r.toEntity(rows)
		if err != nil {
			return nil, errors.New("error on fetch orders")
		}
		orders = append(orders, o)
	}
	return orders, nil
}


func (r *OrderRepository) toEntity(row *sql.Rows) (*entity.Order, error) {
	var id string
	var price float64
	var tax float64
	var finalPrice float64
	if err := row.Scan(&id, &price, &tax, &finalPrice); err != nil {
		return nil, err
	}
	return entity.RestoreOrder(id, price, tax, finalPrice)
}
