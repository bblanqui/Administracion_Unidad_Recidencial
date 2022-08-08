package loginrepository

import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/core/entity"
	"github.com/bblanqui/Administracion_Unidad_Recidencial/infractructure/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UsuarioRepository struct {
	Db         string
	Collection string
	Conexion   *mongo.Client
}

var dbColeccionUsuario = UsuarioRepository{
	Db:         "Usuarios_Admin",
	Collection: "Usuario",
	Conexion:   db.MongoCN,
}

func (user *UsuarioRepository) GetUsuario() []entity.Usuario {
	contexto, cancelar, mongoCliente := db.DbColeccion(dbColeccionUsuario.Db, dbColeccionUsuario.Collection)
	defer cancelar()
	var clientes entity.Usuarios
	var usuario entity.Usuario

	re, _ := mongoCliente.Find(contexto, bson.D{})
	re.Decode(&clientes)

	for re.Next(contexto) {

		if err := re.Decode(&usuario); err != nil {
			log.Fatal(err)
		}
		clientes = append(clientes, usuario)

	}
	if err := re.Err(); err != nil {
		log.Fatal(err)
	}
	return clientes
}

func (user *UsuarioRepository) GetUsuario_id(email string) entity.Usuario {

	contexto, cancelar, mongoCliente := db.DbColeccion(dbColeccionUsuario.Db, dbColeccionUsuario.Collection)
	defer cancelar()

	var usuario entity.Usuario

	mongoCliente.FindOne(contexto, bson.M{"email": email}).Decode(&usuario)

	return usuario
}

func (user *UsuarioRepository) InsertUsuario(Usuario *entity.Usuario) bool {
	return true
}

type IUsuarioRepository interface {
	GetUsuario() []entity.Usuario
	GetUsuario_id(id int) entity.Usuario
	InsertUsuario(Usuario *entity.Usuario) bool
}
