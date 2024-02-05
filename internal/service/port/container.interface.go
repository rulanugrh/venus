package iservice

import (
	"context"
	"io"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type ContainerInterface interface {
	Create(req dto.Container, ctx context.Context) (*dao.Container, error)
	ListContainer(ctx context.Context) ([]dao.ListContainer, error)
	InspectContainer(ctx context.Context, id string) (*dao.InspectContainer, error)
	DeleteContaienr(id string, ctx context.Context) error
	ExecContainer(id string, r io.Reader, w io.Writer, ctx context.Context) error
	Logger(name string, ctx context.Context, output io.Writer) error 
}