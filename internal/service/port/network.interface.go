package iservice

import (
	"context"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type NetworkInterface interface {
	CreateNetwork(model dto.Network, ctx context.Context) (*dao.Network, error)
	InspectNetwork(id string, ctx context.Context) (*dao.Network, error)
	ListNetworks(ctx context.Context) ([]dao.Network, error)
	DeleteNetwork(id string, ctx context.Context) error
}
