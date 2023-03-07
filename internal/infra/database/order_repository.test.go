package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory")

	suite.NoError(err)

	// create table query

	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite()  {
	suite.Db.Close()
}

func Test_Suite(t *testing.T) {
	 suite.Run(t, new(OrderRepositoryTestSuite))
}