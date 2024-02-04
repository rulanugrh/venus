package iservice

import (
	"context"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type VolumeInterface interface {
	CreateVolume(req dto.Volume, ctx context.Context) (*dao.Volume, error)
	ListVolume(ctx context.Context) ([]dao.Volume, error)
	InspectVolume(name string, ctx context.Context) (*dao.Volume, error)
	DeleteVolume(name string, ctx context.Context) error
}