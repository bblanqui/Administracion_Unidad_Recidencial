package services
import (
	"github.com/bblanqui/Administracion_Unidad_Recidencial/core/entity"
	"github.com/bblanqui/Administracion_Unidad_Recidencial/infractructure/loginrepository"
)

var  Repositorio = loginrepository.UsuarioRepository{}

type ServiceUsuario struct {
      Conexion string
     
}

func (user *ServiceUsuario) GetUsuario() []entity.Usuario {
	
	
	 return Repositorio.GetUsuario()
	}    
	 


func (user *ServiceUsuario) GetUsuario_id(id int) entity.Usuario {
	return Repositorio.GetUsuario_id(id)
}

func (user *ServiceUsuario) InsertUsuario(Usuario *entity.Usuario) bool {
	
	return Repositorio.InsertUsuario(Usuario)
}



type IUsuarioService interface { 

	GetUsuario() []entity.Usuario
	GetUsuario_id(id int) entity.Usuario
	InsertUsuario(Usuario *entity.Usuario) bool
	

}