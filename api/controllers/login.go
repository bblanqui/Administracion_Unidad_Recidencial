package controllers

import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/services"
	"github.com/gofiber/fiber/v2"
)

var usuarios = services.ServiceUsuario{}

func GetUsuarios(c *fiber.Ctx) error {
	var usuario = usuarios.GetUsuario()
	return c.Status(fiber.StatusOK).JSON(usuario)

}
func GetUsuario(c *fiber.Ctx) error {
	var usuario = usuarios.GetUsuario_id(c.Params("+"))
	return c.Status(fiber.StatusOK).JSON(usuario)

}
