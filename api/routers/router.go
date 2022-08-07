package routers

import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App) {

	app.Get("/", controllers.GetUsuarios)
	app.Get("/usuario", controllers.GetUsuario)

}
