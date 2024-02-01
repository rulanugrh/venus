package iservice

import (
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
)

type ImageInterface interface {
	PullImage(req dto.Image) error
	ListImage() ([]dao.Image, error)
	InspectImage(id string) (*dao.InspectImage, error)
	DeleteImage(id string) error
}