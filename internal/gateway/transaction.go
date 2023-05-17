package gateway

import "github.com/igor-marchi/fc-ms-digital-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
