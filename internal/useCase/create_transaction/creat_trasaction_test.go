package createtransaction

import (
	"testing"

	"github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase(t *testing.T) {
	client1, _ := entity.NewClient("client1", "client1@client1.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(100)

	client2, _ := entity.NewClient("client2", "client2@client2.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(100)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("Get", account1.Id).Return(account1, nil)
	mockAccount.On("Get", account2.Id).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIdFrom: account1.Id,
		AccountIdTo:   account2.Id,
		Amount:        100,
	}

	useCase := NewCreateTransactionUseCase(mockTransaction, mockAccount)

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.Id)
	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "Get", 2)
}
