package service

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type volumestruct struct {
	client *docker.Client
}

func NewVolumeService(client *docker.Client) iservice.VolumeInterface {
	return &volumestruct{
		client: client,
	}
}

func(volume *volumestruct) CreateVolume(req dto.Volume, ctx context.Context) (*dao.Volume, error) {
	data, err := volume.client.CreateVolume(docker.CreateVolumeOptions{
		Name: req.Name,
		Driver: req.Driver,
		DriverOpts: req.DriverOpts,
		Labels: req.Labels,
		Context: ctx,
	})

	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	response := dao.Volume{
		Name: data.Name,
		Driver: data.Driver,
		Labels: data.Labels,
		DriverOpts: data.Options,
	}

	return &response, nil
}

func(volume *volumestruct) ListVolume() ([]dao.Volume, error) {
	data, err := volume.client.ListVolumes(docker.ListVolumesOptions{})
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	var response []dao.Volume
	for _, result := range data {
		vlm := dao.Volume{
			Name: result.Name,
			Labels: result.Labels,
			Driver: result.Driver,
			DriverOpts: result.Options,
		}

		response = append(response, vlm)
	}

	return response, nil
}

func(volume *volumestruct) InspectVolume(name string) (*dao.Volume, error) {
	data, err := volume.client.InspectVolume(name)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	response := dao.Volume{
		Name: data.Name,
		Driver: data.Driver,
		DriverOpts: data.Options,
		Labels: data.Labels,
	}

	return &response, nil
}

func(volume *volumestruct) DeleteVolume(name string) error {
	err := volume.client.RemoveVolume(name)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	return nil
}