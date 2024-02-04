package iservice

import (
	"context"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type ImageInterface interface {
	PullImage(ctx context.Context, req dto.Image) error
	ListImage(ctx context.Context) ([]dao.Image, error)
	InspectImage(id string, ctx context.Context) (*dao.InspectImage, error)
	DeleteImage(id string, ctx context.Context) error
	BuildImage(model dto.BuildImage, ctx context.Context) error
}