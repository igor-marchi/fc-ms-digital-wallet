package gateway

import "github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	Get(id string) (*entity.Account, error)
}
