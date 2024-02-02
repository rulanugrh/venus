package service

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rulanugrh/venus/internal/entity/dao"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type servicestruct struct {
	client *docker.Client
}

func NewNetworkService(client *docker.Client) iservice.NetworkInterface {
	return &servicestruct{
		client: client,
	}
}

func (network *servicestruct) CreateNetwork(model dto.Network) (*dao.Network, error) {
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

func (network *servicestruct) InspectNetwork(id string) (*dao.Network, error) {
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

func (network *servicestruct) ListNetworks() ([]dao.Network, error) {
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

func (network *servicestruct) DeleteNetwork(id string) error {
	err := network.client.RemoveNetwork(id)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code:    500,
		}
	}

	return nil
}
