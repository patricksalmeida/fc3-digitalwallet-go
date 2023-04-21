package database

import (
	"database/sql"
	"testing"

	"github.com.br/patricksalmeida/fc3-digitalwallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id string, client_id string, balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id string, account_id_from string, account_id_to string, amount float, created_at date)")

	s.transactionDB = NewTransactionDB(db)

	client1, err := entity.NewClient("Jhon Doe 1", "jhon1@doe.com")
	s.Nil(err)
	s.client1 = client1

	accountFrom := entity.NewAccount(client1)
	accountFrom.Balance = 1000
	s.Nil(err)
	s.accountFrom = accountFrom

	client2, err := entity.NewClient("Jhon Doe 2", "jhon1@doe.com")
	s.Nil(err)
	s.client2 = client2

	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	s.Nil(err)
	s.accountTo = accountTo
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")
}

func (s *TransactionDBTestSuite) TestCreateTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
