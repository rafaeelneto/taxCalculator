package database

import (
	"database/sql"
	"testing"

	"github.com/rafaeelneto/taxCalculator/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory")

	suite.NoError(err)

	// create table query
	db.Exec(`CREATE TABLE orders 
	(id (varchar(255) NOT NULL,
	 price FLOAT NOT NULL,
	 tax FLOAT NOT NULL,
	 quantity INTEGER NOT NULL,
	 final_price NOT NULL,
	 PRIMARY KEY (id))`)
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite()  {
	suite.Db.Close()
}

func Test_Suite(t *testing.T) {
	 suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) Test_SavingOrder () {
	order, err := entity.NewOrder("1223", 10.2, 0.1, 10)

	suite.NoError(err)

	repo := NewOrderRepository(suite.Db)
	errOnSave := repo.Save(order)

	suite.NoError(errOnSave)
}