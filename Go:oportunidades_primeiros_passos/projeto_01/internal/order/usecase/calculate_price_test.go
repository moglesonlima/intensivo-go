package usecase

import (
	"database/sql"
	"testing"

	"github.com/moglesonlima/gointensivo/internal/order/entity"
	"github.com/moglesonlima/gointensivo/internal/order/infra/database"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuit struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

// Antes de exec o teste
func (suite *CalculatePriceUseCaseTestSuit) SetupSuite() {

	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db

	suite.OrderRepository = *database.NewOrderRepository(db)
}

// Depois de exec o teste
func (suite *CalculatePriceUseCaseTestSuit) TearDownTest() {
	// Aqui você deve fechar a conexão com o banco de dados
	suite.Db.Close()
}

// Criação de test
func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuit))
}

func (suite *CalculatePriceUseCaseTestSuit) TestGivenAValidOrder_WhenCalculatePrice_ThenShouldCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	order.CalculateFinalPrice()

	CalculateFinalPriceImput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	CalculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := CalculateFinalPriceUseCase.Execute(CalculateFinalPriceImput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)

}
