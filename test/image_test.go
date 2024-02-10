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

type ImageTest struct {
	suite.Suite
	client util.SuiteInterface
	res *web.Success
	conf *config.App
}

func NewImageTest() *ImageTest {
	return &ImageTest{
		client: util.NewSuiteUtils(&http.Client{}),
		res: &web.Success{},
		conf: config.GetConfig(),
	}
}

func(image *ImageTest) TestPullImage() {
	req := dto.Image{
		Repository: "alpine",
		Tag: "latest",
		Platform: "alpine",
		Email: image.conf.Docker.Email,
		Username: image.conf.Docker.Username,
		Password: image.conf.Docker.Password,
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	url := fmt.Sprintf("http://%s:%s/api/image/create", image.conf.Server.Host, image.conf.Server.Port)

	res, resp, err := image.client.Post(url, byt, image.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	image.Equal(http.StatusAccepted, resp.StatusCode)
	image.Equal("Berhasil pull image", res.Message)	
}

func(image *ImageTest) TestListImage() {
	url := fmt.Sprintf("http://%s:%s/api/image/find", image.conf.Server.Host, image.conf.Server.Port)

	res, resp, err := image.client.Get(url, image.res, Token.(string))
	if err != nil {
		log.Fatal(err)
	}

	image.Equal(http.StatusAccepted, resp.StatusCode)
	image.Equal("Image ditemukan", res.Message)	
}

func TestImage(t *testing.T) {
	suite.Run(t, NewImageTest())
}