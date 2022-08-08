package main

import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/api/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.Login(app)
	app.Listen(":4000")

}
