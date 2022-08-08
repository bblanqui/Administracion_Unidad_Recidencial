package entity

import (
	"time"
    "go.mongodb.org/mongo-driver/bson/primitive"

)

type Usuario struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Nombre_Usuario string `bson:"nombre_usuario" json:"nombre_usuario"`
	Password string `bson:"password,omitempty" json:"password"`
    Role string `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdateAt time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
}

type Usuarios []Usuario




