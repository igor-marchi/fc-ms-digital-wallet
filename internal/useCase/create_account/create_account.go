package createaccount

import (
	"github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"
	"github.com/igor-marchi/fc-ms-digital-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	clientId string
}

type CreateAccountOutputDTO struct {
	id string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	clientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: accountGateway,
		clientGateway:  clientGateway,
	}
}

func (useCase *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := useCase.clientGateway.Get(input.clientId)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	err = useCase.accountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		id: account.Id,
	}, nil
}
