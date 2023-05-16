package gateway

import "github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
