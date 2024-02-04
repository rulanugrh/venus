package service

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	iservice "github.com/rulanugrh/venus/internal/service/port"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type servicestruct struct {
	client *docker.Client
	tracer trace.Tracer
}

func NewNetworkService(client *docker.Client, tracer trace.Tracer) iservice.NetworkInterface {
	return &servicestruct{
		client: client,
		tracer: tracer,
	}
}

func (network *servicestruct) CreateNetwork(model dto.Network, ctx context.Context) (*dao.Network, error) {
	_, span := network.tracer.Start(ctx, "createNetwork")
	defer span.End()

	data, err := network.client.CreateNetwork(docker.CreateNetworkOptions{
		Name:           model.Name,
		Driver:         model.Driver,
		Scope:          model.Scope,
		Internal:       model.Internal,
		EnableIPv6:     model.EnableIPV6,
		CheckDuplicate: true,
	})
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	response := dao.Network{
		Name:       data.Name,
		ID:         data.ID,
		Driver:     data.Driver,
		Scope:      data.Scope,
		Internal:   data.Internal,
		EnableIPV6: data.EnableIPv6,
		Container:  data.Containers,
	}

	return &response, nil
}

func (network *servicestruct) InspectNetwork(id string, ctx context.Context) (*dao.Network, error) {
	_, span := network.tracer.Start(ctx, "inspectNetwork", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	data, err := network.client.NetworkInfo(id)
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	response := dao.Network{
		ID:         data.ID,
		Name:       data.Name,
		Driver:     data.Driver,
		EnableIPV6: data.EnableIPv6,
		Internal:   data.Internal,
		Container:  data.Containers,
	}

	return &response, nil
}

func (network *servicestruct) ListNetworks(ctx context.Context) ([]dao.Network, error) {
	_, span := network.tracer.Start(ctx, "listNetwork")
	defer span.End()

	data, err := network.client.ListNetworks()
	if err != nil {
		return nil, web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	var response []dao.Network
	for _, result := range data {
		net := dao.Network{
			ID:         result.ID,
			Name:       result.Name,
			Scope:      result.Scope,
			Driver:     result.Driver,
			EnableIPV6: result.EnableIPv6,
			Internal:   result.Internal,
		}

		response = append(response, net)
	}

	return response, nil
}

func (network *servicestruct) DeleteNetwork(id string, ctx context.Context) error {
	_, span := network.tracer.Start(ctx, "deleteNetwork", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	err := network.client.RemoveNetwork(id)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	return nil
}
