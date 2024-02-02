package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rulanugrh/venus/config"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	"github.com/rulanugrh/venus/internal/middleware"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type userstruct struct {
	conf config.App
	validate *validator.Validate
}

func NewUserService(conf config.App) iservice.UserInterface {
	return &userstruct{
		conf: conf,
		validate: validator.New(),
	}
}

func(user *userstruct) Login(req dto.User) error {
	err := middleware.ValidateStruct(user.validate, req)
	if err != nil {
		return web.Error{
			Message: err.Error(),
			Code: 400,
		}
	}

	if req.Email != user.conf.Admin.Email {
		return web.Error{
			Message: "Maaf email anda tidak cocok",
			Code: 401,
		}
	}

	if req.Password != user.conf.Admin.Password {
		return web.Error{
			Message: "Password tidak cocok",
			Code: 401,
		}
	}

	return nil
}
