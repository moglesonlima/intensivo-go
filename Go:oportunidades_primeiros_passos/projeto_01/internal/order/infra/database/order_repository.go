package database

import (
	"database/sql"

	"github.com/moglesonlima/gointensivo/internal/order/entity"
)

type OrderRepository struct {
	// db é um ponteiro para um sql.DB
	// O sql.DB é uma struct que representa uma conexão com o banco de dados
	db *sql.DB
}

// NewOrderRepository atua como um construtor para OrderRepository(criando a conexão com o banco de dados)
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	// Aqui você deve implementar a lógica para salvar a ordem no banco de dados
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
