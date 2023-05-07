package sqlstore

import (
	"Wb-L0/internal/model"
	"database/sql"
)

type SqlStore struct {
	*sql.DB
}

func New(dbUrl string) (*SqlStore, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &SqlStore{db}, nil
}

func (s *SqlStore) GetAll() ([]model.Order, error) {
	q := `SELECT * FROM orders`
	rows, err := s.Query(q)
	defer rows.Close()
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err = rows.Scan(&order.UUID, &order.OrderJson); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *SqlStore) Add(order *model.Order) error {
	q := `INSERT INTO orders VALUES ($1, $2)`
	return s.QueryRow(q, order.UUID, order.OrderJson).Err()
}
