package service

import (
	"context"
	"io"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type containerstruct struct {
	client *docker.Client
	tracer trace.Tracer
}

func NewContainerService(client *docker.Client, tracer trace.Tracer) iservice.ContainerInterface {
	return &containerstruct{
		client: client,
		tracer: tracer,
	}
}

func (container *containerstruct) Create(req dto.Container, ctx context.Context) (*dao.Container, error) {
	_, span := container.tracer.Start(ctx, "createContainer")
	defer span.End()

	data, err := container.client.CreateContainer(docker.CreateContainerOptions{
		Name:     req.Name,
		Platform: req.Platform,
		Config: &docker.Config{
			Hostname:     req.Config.Hostname,
			Domainname:   req.Config.Domainname,
			Image:        req.Config.Image,
			Tty:          req.Config.Tty,
			OpenStdin:    req.Config.OpenStdin,
			Env:          req.Config.Env,
			ExposedPorts: req.Config.Port,
		},
		Context: ctx,
	})

	if err != nil {
		return nil, web.Error{
			Message: "Tidak bisa membuat container ke docker daemon",
			Code:    400,
		}
	}

	errs := container.client.StartContainer(data.ID, &docker.HostConfig{
		Binds:        req.HostConfig.Binds,
		PortBindings: req.HostConfig.PortBinding,
		NetworkMode:  req.HostConfig.NetworkMode,
	})

	if errs != nil {
		return nil, web.Error{
			Message: errs.Error(),
			Code:    400,
		}
	}

	response := dao.Container{
		ID:           data.ID,
		Image:        data.Image,
		Created:      data.Created,
		Path:         data.Path,
		HostnamePath: data.HostnamePath,
		HostsPath:    data.HostsPath,
		Config: &dao.Config{
			Hostname:   data.Config.Hostname,
			Domainname: data.Config.Domainname,
			Tty:        data.Config.Tty,
			OpenStdin:  data.Config.OpenStdin,
			Env:        data.Config.Env,
			Image:      data.Image,
			Port:       data.Config.ExposedPorts,
		},
	}

	return &response, nil
}

func (container *containerstruct) ListContainer(ctx context.Context) ([]dao.ListContainer, error) {
	_, span := container.tracer.Start(ctx, "listContiner")
	defer span.End()

	data, err := container.client.ListContainers(docker.ListContainersOptions{
		All:   true,
		Limit: 10,
	})

	if err != nil {
		return nil, web.Error{
			Message: "Tidak bisa melihat semua service di docker daemon",
			Code:    400,
		}
	}

	var response []dao.ListContainer
	var listPorts []dao.Port
	for _, result := range data {
		for _, ports := range result.Ports {
			port := dao.Port{
				PrivatePort: ports.PrivatePort,
				PublicPort:  ports.PublicPort,
				IP:          ports.IP,
				Type:        ports.Type,
			}

			listPorts = append(listPorts, port)

		}

		singleContainer := dao.ListContainer{
			ID:      result.ID,
			Image:   result.Image,
			Status:  result.Status,
			Command: result.Command,
			State:   result.State,
			Created: result.Created,
			Ports:   listPorts,
		}

		response = append(response, singleContainer)
	}

	return response, nil
}

func (container *containerstruct) InspectContainer(ctx context.Context, id string) (*dao.InspectContainer, error) {
	_, span := container.tracer.Start(ctx, "inspectContainer", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	data, err := container.client.InspectContainer(id)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	response := dao.InspectContainer{
		ID:           data.ID,
		Image:        data.Image,
		HostnamePath: data.HostnamePath,
		HostsPath:    data.HostsPath,
		Name:         data.Name,
		Tty:          data.Config.Tty,
		OpenStdin:    data.Config.OpenStdin,
		Env:          data.Config.Env,
		Port:         data.Config.ExposedPorts,
	}

	return &response, nil
}

func (container *containerstruct) DeleteContaienr(id string, ctx context.Context) error {
	_, span := container.tracer.Start(ctx, "deleteCotainer", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	err := container.client.RemoveContainer(docker.RemoveContainerOptions{
		Context: ctx,
		ID:      id,
	})
	
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    400,
		}
	}

	return nil
}

func(container *containerstruct) ExecContainer(id string, r io.Reader, w io.Writer, ctx context.Context) error {
	_, span := container.tracer.Start(ctx, "execContainer", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	err := container.client.StartExec(id, docker.StartExecOptions{
		Tty: true,
		RawTerminal: true,
		InputStream: r,
		OutputStream: w,
		Context: ctx,
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	return nil
}

func(container *containerstruct) Logger(name string, ctx context.Context, output io.Writer) error {
	err := container.client.Logs(docker.LogsOptions{
		Container: name,
		Context: ctx,
		RawTerminal: true,
		OutputStream: output,
	})

	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	return nil
}