package gateway

import "github.com.br/patricksalmeida/fc3-digitalwallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
