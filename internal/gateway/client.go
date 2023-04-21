package gateway

import "github.com.br/patricksalmeida/fc3-digitalwallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
