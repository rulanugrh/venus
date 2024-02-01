package service

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"

	iservice "github.com/rulanugrh/venus/internal/service/port"
)
type containerstruct struct {
	client *docker.Client
}

func NewDockerService(client *docker.Client) iservice.ContainerInterface {
	return &containerstruct{
		client: client,
	}
}

func(container *containerstruct) Create(req dto.Container, ctx context.Context) (*dao.Container, error) {
	data, err := container.client.CreateContainer(docker.CreateContainerOptions{
		Name: req.Name,
		Platform: req.Platform,
		Config: &docker.Config{
			Hostname: req.Config.Hostname,
			Domainname: req.Config.Domainname,
			Image: req.Config.Image,
			Tty: req.Config.Tty,
			OpenStdin: req.Config.OpenStdin,
			Env: req.Config.Env,
			ExposedPorts: req.Config.Port,
		},
		Context: ctx,
	})

	if err != nil {
		return nil, web.Error{
			Message: "Tidak bisa membuat container ke docker daemon",
			Code: 400,
		}
	}

	response := dao.Container{
		ID: data.ID,
		Image: data.Image,
		Created: data.Created,
		Path: data.Path,
		HostnamePath: data.HostnamePath,
		HostsPath: data.HostsPath,
		Config: &dao.Config{
			Hostname: data.Config.Hostname,
			Domainname: data.Config.Domainname,
			Tty: data.Config.Tty,
			OpenStdin: data.Config.OpenStdin,
			Env: data.Config.Env,
			Image: data.Image,
			Port: data.Config.ExposedPorts,
		},
	}

	return &response, nil
}

func(container *containerstruct) ListContainer() ([]dao.ListContainer, error) {
	data, err := container.client.ListContainers(docker.ListContainersOptions{
		All: true,
		Limit: 10,	
	})

	if err != nil {
		return nil, web.Error{
			Message: "Tidak bisa melihat semua service di docker daemon",
			Code: 400,
		}
	}

	var response []dao.ListContainer
	var listPorts []dao.Port
	for _, result := range data {
		for _, ports := range result.Ports {
			port := dao.Port{
				PrivatePort: ports.PrivatePort,
				PublicPort: ports.PublicPort,
				IP: ports.IP,
				Type: ports.Type,
			}

			listPorts = append(listPorts, port)

		}

		singleContainer := dao.ListContainer{
			ID: result.ID,
			Image: result.Image,
			Status: result.Status,
			Command: result.Command,
			State: result.State,
			Created: result.Created,
			Ports: listPorts,
		}

		response = append(response, singleContainer)
	}

	return response, nil
}

func(container *containerstruct) InspectContainer(id string) (*dao.InspectContainer, error) {
	data, err := container.client.InspectContainer(id)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	response := dao.InspectContainer{
		ID: data.ID,
		Image: data.Image,
		HostnamePath: data.HostnamePath,
		HostsPath: data.HostsPath,
		Name: data.Name,
		Tty: data.Config.Tty,
		OpenStdin: data.Config.OpenStdin,
		Env: data.Config.Env,
		Port: data.Config.ExposedPorts,
	}

	return &response, nil
}

func(container *containerstruct) DeleteContaienr(id string, ctx context.Context) error {
	err := container.client.RemoveContainer(docker.RemoveContainerOptions{
		Context: ctx,
		ID: id,
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}
	
	return nil
}
