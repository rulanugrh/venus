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

type ContainerTest struct {
	suite.Suite
	client util.SuiteInterface
	res *web.Success
	conf *config.App
}

func NewContainerTest() *ContainerTest {
	return &ContainerTest{
		client: util.NewSuiteUtils(&http.Client{}),
		res: &web.Success{},
		conf: config.GetConfig(),
	}
}

func(container *ContainerTest) TestCreateContainer() {
	req := dto.Container{
		Name: "test-image",
		Platform: "alpine",
		Config: &dto.Config{
			Image: "alpine:latest",
			OpenStdin: true,
			Tty: true,
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	url := fmt.Sprintf("http://%s:%s/api/container/create", container.conf.Server.Host, container.conf.Server.Port)

	res, resp, err := container.client.Post(url, byt, container.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	container.Equal(http.StatusAccepted, resp.StatusCode)
	container.Equal("Success create container", res.Message)
}

func(container *ContainerTest) TestListContainer() {
	url := fmt.Sprintf("http://%s:%s/api/container/find", container.conf.Server.Host, container.conf.Server.Port)

	res, resp, err := container.client.Get(url, container.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	container.Equal(http.StatusAccepted, resp.StatusCode)
	container.Equal("Container ditemukan", res.Message)
}

func TestContainer(t *testing.T) {
	suite.Run(t, NewContainerTest())
}