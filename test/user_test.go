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

type UserTest struct {
	suite.Suite
	client util.SuiteInterface
	res    *web.Success
}

func NewUserTest() *UserTest {
	return &UserTest{
		client: util.NewSuiteUtils(&http.Client{}),
		res:    &web.Success{},
	}
}

func (user *UserTest) TestLogin() {
	conf := config.GetConfig()
	input := dto.User{
		Email:    conf.Admin.Email,
		Password: conf.Admin.Password,
	}

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)
	url := fmt.Sprintf("http://%s:%s/api/v1/user/login", conf.Server.Host, conf.Server.Port)
	res, resp, err := user.client.Post(url, byt, user.res)
	user.Equal(http.StatusAccepted, resp.StatusCode)
	user.Equal("success login", res.Message)

}

func TestUserService(t *testing.T) {
	suite.Run(t, NewUserTest())
}
