package iservice

import "github.com/rulanugrh/venus/internal/entity/dto"

type UserInterface interface {
	Login(req dto.User) error
}