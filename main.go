package main
import(
  
  "github.com/gofiber/fiber/v2"
  "github.com/bblanqui/Administracion_Unidad_Recidencial/api/routers"
)


func main() {

 app := fiber.New()
 routers.Login(app)
 
 app.Listen(":4000")



}