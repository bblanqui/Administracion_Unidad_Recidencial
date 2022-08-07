package entity

import "time"

type Usuario struct{
	ID string `json:"id"`
    Nombre_Usuario string `json:"nombre_usuario"`
	Password string `json:"password"`
    Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

type Usuarios []Usuario




