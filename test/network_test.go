package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/rulanugrh/venus/config"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	"github.com/rulanugrh/venus/util"
	"github.com/stretchr/testify/suite"
)

type NetworkTest struct {
	suite.Suite
	client util.SuiteInterface
	res    *web.Success
	conf   *config.App
}

func NewNetworkTest() *NetworkTest {
	return &NetworkTest{
		client: util.NewSuiteUtils(&http.Client{}),
		res:    &web.Success{},
		conf:   config.GetConfig(),
	}
}

func (network *NetworkTest) TestGetAllNetwork() {
	url := fmt.Sprintf("http://%s:%s/api/v1/network/find", network.conf.Server.Host, network.conf.Server.Port)
	res, resp, err := network.client.Get(url, network.res)
	if err != nil {
		log.Fatal(err)
	}

	network.Equal(http.StatusAccepted, resp.StatusCode)
	network.Equal("Network ditemukan", res.Message)
}

func (network *NetworkTest) TestInspectNetwork() {
	url := fmt.Sprintf("http://%s:%s/api/v1/network/find", network.conf.Server.Host, network.conf.Server.Port)
	res, resp, err := network.client.Get(url, network.res)

	if err != nil {
		log.Fatal(err)
	}

	network.Equal(http.StatusAccepted, resp.StatusCode)
	network.Equal("Network ditemukan", res.Message)
}

func (network *NetworkTest) TestDeleteNetwork() {
	url := fmt.Sprintf("http://%s:%s/api/v1/network/delete", network.conf.Server.Host, network.conf.Server.Port)
	res, resp, err := network.client.Get(url, network.res)

	if err != nil {
		log.Fatal(err)
	}

	network.Equal(http.StatusAccepted, resp.StatusCode)
	network.Equal("Network berhasil dihapus", res.Message)

}

func (network *NetworkTest) TestCreateNetwork() {
	req := dto.Network{
		Name:       "network-testing",
		Driver:     "bridge",
		Internal:   false,
		EnableIPV6: true,
		Scope:      "forDevelop",
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	url := fmt.Sprintf("http://%s:%s/api/v1/network/create", network.conf.Server.Host, network.conf.Server.Port)

	res, resp, err := network.client.Post(url, byt, network.res)
	if err != nil {
		log.Fatal(err)
	}

	network.Equal(http.StatusAccepted, resp.StatusCode)
	network.Equal("Success buat network", res.Message)
	network.Equal(200, res.Code)
}

func TestNework(t *testing.T) {
	suite.Run(t, NewNetworkTest())
}
