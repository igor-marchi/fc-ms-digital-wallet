package createaccount

import (
	"testing"

	"github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
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

func TestCreateAccountUseCase(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.Id).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateAccountUseCase(accountMock, clientMock)

	inputDto := CreateAccountInputDTO{
		clientId: client.Id,
	}

	output, err := useCase.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.id)
	accountMock.AssertExpectations(t)
	clientMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
}
