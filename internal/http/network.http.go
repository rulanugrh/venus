package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
	ihttp "github.com/rulanugrh/venus/internal/http/port"
	iservice "github.com/rulanugrh/venus/internal/service/port"
)

type networkstruct struct {
	service iservice.NetworkInterface
}

func NewNetworkHandler(service iservice.NetworkInterface) ihttp.NetworkInterface {
	return &networkstruct{
		service: service,
	}
}
func(network *networkstruct) CreateNetwork(ctx *fiber.Ctx) error {
	var model dto.Network
	err := ctx.BodyParser(&model)
	if err != nil {
		response := web.Failure{
			Message: "Gagal binding body",
			Code: 500,
			Error: err.Error(),
		}

		return ctx.Status(500).JSON(response)
	}
	data, errs := network.service.CreateNetwork(model)
	if errs != nil {
		response := web.Failure{
			Message: "Gagal buat network",
			Code: 400,
			Error: errs.Error(),
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "Success buat network",
		Code: 200,
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}

func(network *networkstruct) InspectNetwork(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := network.service.InspectNetwork(id)
	if err != nil {
		response := web.Failure{
			Message: "Gagal inspect network",
			Code: 400,
			Error: err.Error(),
		}

		return ctx.Status(400).JSON(response)
	}

	response := web.Success{
		Message: "Network ditemukan",
		Code: 200,
		Data: data,
	}

	return ctx.Status(200).JSON(response)
}