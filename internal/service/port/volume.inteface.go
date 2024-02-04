package iservice

import (
	"context"

	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type VolumeInterface interface {
	CreateVolume(req dto.Volume, ctx context.Context) (*dao.Volume, error)
	ListVolume() ([]dao.Volume, error)
	InspectVolume(name string) (*dao.Volume, error)
	DeleteVolume(name string) error
}