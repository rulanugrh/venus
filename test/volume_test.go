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

type VolumeTest struct {
	suite.Suite
	client util.SuiteInterface
	res    *web.Success
	conf   *config.App
}

func NewVolumeTest() *VolumeTest {
	return &VolumeTest{
		client: util.NewSuiteUtils(&http.Client{}),
		res:    &web.Success{},
		conf:   config.GetConfig(),
	}
}

func (volume *VolumeTest) TestCreateVolume() {
	req := dto.Volume{
		Name: "testing-volume",
		Driver: "local",
		Labels: map[string]string{
			"for": "database",
		},
		DriverOpts: map[string]string{
			"type": "nfs",
		},
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	byt := bytes.NewBuffer(jsonBytes)
	url := fmt.Sprintf("http://%s:%s/api/v1/volume/create", volume.conf.Server.Host, volume.conf.Server.Port)

	res, resp, err := volume.client.Post(url, byt, volume.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	volume.Equal(http.StatusAccepted, resp.StatusCode)
	volume.Equal("Success create volume", res.Message)
}

func(volume *VolumeTest) TestListNetwork() {
	url := fmt.Sprintf("http://%s:%s/api/v1/volume/find", volume.conf.Server.Host, volume.conf.Server.Port)

	res, resp, err := volume.client.Get(url, volume.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	volume.Equal(http.StatusAccepted, resp.StatusCode)
	volume.Equal("volume ditemukan", res.Message)
}

func TestVolume(t *testing.T) {
	suite.Run(t, NewVolumeTest())
}
