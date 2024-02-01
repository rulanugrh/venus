package iservice

import (
	"context"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type ContainerInterface interface {
	Create(req dto.Container, ctx context.Context) (*dao.Container, error)
	ListContainer() ([]dao.ListContainer, error)
	InspectContainer(id string) (*dao.InspectContainer, error)
	DeleteContaienr(id string, ctx context.Context) error
}