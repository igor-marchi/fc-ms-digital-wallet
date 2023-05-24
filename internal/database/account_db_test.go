package database

import (
	"database/sql"
	"testing"

	"github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE client (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE account (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Jhon", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE client")
	s.db.Exec("DROP TABLE account")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.db.Exec("INSERT INTO client (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.Id, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.AccountDB.Get(account.Id)
	s.Nil(err)
	s.Equal(account.Id, accountDB.Id)
	s.Equal(account.Client.Id, accountDB.Client.Id)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.Id, accountDB.Client.Id)
}
