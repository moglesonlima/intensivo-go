package database

import (
	"database/sql"
	"testing"

	"github.com/moglesonlima/gointensivo/internal/order/entity"
	"github.com/stretchr/testify/suite"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Trabalhando com testes (suite de testes)
type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// Antes de exec o teste
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	// Aqui você deve criar a conexão com o banco de dados
	// e executar as migrations
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	// cria a tabela orders
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

// Depois de exec o teste
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	// Aqui você deve fechar a conexão com o banco de dados
	suite.Db.Close()
}

// Criação de test
func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAValidOrder_WhenSave_ThenShouldSaveOrder() {
	// Aqui você deve implementar a lógica para salvar a ordem no banco de dados
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	// Aqui você deve implementar a lógica para buscar a ordem no banco de dados
	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	// Aqui você deve implementar a lógica para comparar a ordem salva com a ordem buscada
	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)

}
