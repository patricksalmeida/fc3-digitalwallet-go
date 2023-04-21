package gateway

import "github.com.br/patricksalmeida/fc3-digitalwallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
