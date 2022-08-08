package loginrepository

import (
	"context"
	"fmt"
	"time"
	"log"


	"github.com/bblanqui/Administracion_Unidad_Recidencial/core/entity"
	"github.com/bblanqui/Administracion_Unidad_Recidencial/infractructure/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
      Db string 
	  Colleccion string
      Conexion *mongo.Client
}
var db_coleccion_usuario = UsuarioRepository{
	 Db:"Usuarios_Admin" ,
	 Colleccion: "Usuario",
	 Conexion: db.MongoCN,
}



func (user *UsuarioRepository) GetUsuario() []entity.Usuario {
     contexto, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	 defer cancel()
	 db :=  db_coleccion_usuario.Conexion.Database(db_coleccion_usuario.Db).Collection(db_coleccion_usuario.Colleccion)
	 
	 var resultado []entity.Usuario
	 var usuario entity.Usuario
	 re,_ := db.Find(contexto, bson.D{})
	 re.Decode(&resultado)

     for re.Next(contexto) {
		
		if err := re.Decode(&usuario); err != nil {
			log.Fatal(err)
		}
		resultado = append(resultado, usuario)
		
		fmt.Println(usuario)
	}
	if err := re.Err(); err != nil {
		log.Fatal(err)
	}
	 return resultado
}    
	 


func (user *UsuarioRepository) GetUsuario_id(id int) entity.Usuario {

	return entity.Usuario{	
    Nombre_Usuario: "id_suario",
	Password : "id_suario",
    Role : "id_suario",
}
}

func (user *UsuarioRepository) InsertUsuario(Usuario *entity.Usuario) bool {
	return true
}



type IUsuarioRepository interface { 

	GetUsuario() []entity.Usuario
	GetUsuario_id(id int) entity.Usuario
	InsertUsuario(Usuario *entity.Usuario) bool
	

}