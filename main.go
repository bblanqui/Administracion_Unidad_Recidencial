package main
import(
  
  "github.com/gofiber/fiber/v2"
  "github.com/bblanqui/Administracion_Unidad_Recidencial/api/routers"
  "github.com/bblanqui/Administracion_Unidad_Recidencial/infractructure/db"
  "fmt"
)


func main() {
 if db.ChequeoConnection(){
    fmt.Println("Starting")
 }else{
  fmt.Println("error")
 }

 app := fiber.New()
 routers.Login(app)
 
 app.Listen(":4000")



}