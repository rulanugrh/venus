package iservice

import (
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type NetworkInterface interface {
	CreateNetwork(model dto.Network) (*dao.Network, error)
	InspectNetwork(id string) (*dao.Network, error)
	ListNetworks() ([]dao.Network, error)
}
