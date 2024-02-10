package service

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/config"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	iservice "github.com/rulanugrh/venus/internal/service/port"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type imagestruct struct {
	client *docker.Client
	conf   *config.App
	tracer trace.Tracer
}

func NewImageService(client *docker.Client, tracer trace.Tracer, conf *config.App) iservice.ImageInterface {
	return &imagestruct{
		client: client,
		tracer: tracer,
		conf:   config.GetConfig(),
	}
}

func (image *imagestruct) PullImage(ctx context.Context, req dto.Image) error {
	_, span := image.tracer.Start(ctx, "pullImage")
	defer span.End()

	err := image.client.PullImage(docker.PullImageOptions{
		Repository: req.Repository,
		Platform:   req.Platform,
		Tag:        req.Tag,
		Context:    ctx,
	}, docker.AuthConfiguration{
		Username: image.conf.Docker.Username,
		Password: image.conf.Docker.Password,
		Email:    image.conf.Docker.Email,
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	return nil
}

func (image *imagestruct) ListImage(ctx context.Context) ([]dao.Image, error) {
	_, span := image.tracer.Start(ctx, "listImage")
	defer span.End()

	data, err := image.client.ListImages(docker.ListImagesOptions{
		All:     true,
		Digests: false,
		Context: ctx,
	})

	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	var response []dao.Image
	for _, result := range data {
		img := dao.Image{
			ID:          result.ID,
			Tag:         result.RepoTags,
			Created:     result.Created,
			Size:        result.Size,
			VirtualSize: result.VirtualSize,
			Labels:      result.Labels,
		}

		response = append(response, img)
	}

	return response, nil
}

func (image *imagestruct) InspectImage(id string, ctx context.Context) (*dao.InspectImage, error) {
	_, span := image.tracer.Start(ctx, "inspectImage", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	data, err := image.client.InspectImage(id)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	response := dao.InspectImage{
		ID:            data.ID,
		Created:       data.Created,
		Tag:           data.RepoTags,
		Size:          data.Size,
		VirtualSize:   data.VirtualSize,
		Architecture:  data.Architecture,
		Author:        data.Author,
		Container:     data.Container,
		OS:            data.OS,
		DockerVersion: data.DockerVersion,
	}

	return &response, nil
}

func (image *imagestruct) DeleteImage(id string, ctx context.Context) error {
	_, span := image.tracer.Start(ctx, "deleteImage", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	err := image.client.RemoveImage(id)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	return nil
}

func (image *imagestruct) BuildImage(model dto.BuildImage, ctx context.Context) error {
	_, span := image.tracer.Start(ctx, "buildImage")
	defer span.End()

	err := image.client.BuildImage(docker.BuildImageOptions{
		Dockerfile:   model.Dockerfile,
		Name:         model.Name,
		Remote:       model.Remote,
		InputStream:  model.InputStream,
		OutputStream: model.OutputStream,
		Labels:       model.Labels,
		Context:      ctx,
		Auth: docker.AuthConfiguration{
			Username: image.conf.Docker.Username,
			Email:    image.conf.Docker.Email,
			Password: image.conf.Docker.Password,
		},
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}
	return nil
}
