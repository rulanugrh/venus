package service

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type imagestruct struct {
	client *docker.Client
}

func NewImageService(client *docker.Client) iservice.ImageInterface{
	return &imagestruct{
		client: client,
	}
}

func(image *imagestruct) PullImage(req dto.Image) error {
	err := image.client.PullImage(docker.PullImageOptions{
		Repository: req.Repository,
		Platform: req.Platform,
		Tag: req.Tag,
	}, docker.AuthConfiguration{
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	return nil
}

func(image *imagestruct) ListImage() ([]dao.Image, error) {
	data, err := image.client.ListImages(docker.ListImagesOptions{
		All: true,
	})

	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	var response []dao.Image
	for _, result := range data {
		img := dao.Image{
			ID: result.ID,
			Tag: result.RepoTags,
			Created: result.Created,
			Size: result.Size,
			VirtualSize: result.VirtualSize,
			Labels: result.Labels,
		}

		response = append(response, img)
	}

	return response, nil
}

func(image *imagestruct) InspectImage(id string) (*dao.InspectImage, error) {
	data, err := image.client.InspectImage(id)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	response := dao.InspectImage{
		ID: data.ID,
		Created: data.Created,
		Tag: data.RepoTags,
		Size: data.Size,
		VirtualSize: data.VirtualSize,
		Architecture: data.Architecture,
		Author: data.Author,
		Container: data.Container,
		OS: data.OS,
		DockerVersion: data.DockerVersion,
	}

	return &response, nil
}

func(image *imagestruct) DeleteImage(id string) error {
	err := image.client.RemoveImage(id)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	return nil
}