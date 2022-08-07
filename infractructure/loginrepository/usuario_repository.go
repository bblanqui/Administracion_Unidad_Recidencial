package loginrepository

import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/core/entity"
)

type UsuarioRepository struct {
      Conexion string

}

func (user *UsuarioRepository) GetUsuario() []entity.Usuario {
	user.Conexion ="ok"
	p1 := entity.Usuario{
	ID: "id_suario",
    Nombre_Usuario: "id_suario",
	Password : "id_suario",
    Role : "id_suario",
}
p2 := entity.Usuario{
	ID: "id_suario",
    Nombre_Usuario: "id_suario",
	Password : "id_suario",
    Role : "id_suario",
}
	
	 return []entity.Usuario{p1,p2}
	}    
	 


func (user *UsuarioRepository) GetUsuario_id(id int) entity.Usuario {
	user.Conexion ="ok"
	return entity.Usuario{	ID: "id_suario",
    Nombre_Usuario: "id_suario",
	Password : "id_suario",
    Role : "id_suario",
}
}

func (user *UsuarioRepository) InsertUsuario(Usuario *entity.Usuario) bool {
	user.Conexion ="ok"
	return true
}



type IUsuarioRepository interface { 

	GetUsuario() []entity.Usuario
	GetUsuario_id(id int) entity.Usuario
	InsertUsuario(Usuario *entity.Usuario) bool
	

}