package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Session struct {
	gorm.Model
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type Student struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	ClassId int    `json:"class_id"`
}

type Class struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `json:"name"`
	Professor  string `json:"professor"`
	RoomNumber int    `json:"room_number"`
}

type StudentClass struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	ClassName  string `json:"class_name"`
	Professor  string `json:"professor"`
	RoomNumber int    `json:"room_number"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

// type SuccessResponse struct {
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

type SuccessResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}
