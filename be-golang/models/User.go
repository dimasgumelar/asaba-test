package models

import (
	"log"
	"time"
)

type User struct {
	Id        int64  `json:"Id" gorm:"column:Id"`
	Name      string `json:"Name" gorm:"column:Name"`
	Email     string `json:"Email" gorm:"column:Email"`
	Password  string `json:"Password" gorm:"column:Password"`
	Role      string `json:"Role" gorm:"column:Role"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt" gorm:"column:UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt" gorm:"column:DeletedAt"`
}

const TableUser = "user"

func UserModel() *User {
	return &User{}
}

func (*User) FindUser(request User) (user User, err error) {
	log.Println("Model : Find User")
	tx := DBAsaba.Table(TableUser).Where("Email = ?", request.Email).Where("DeletedAt IS NULL").Find(&user)
	if tx.Error != nil {
		log.Println("Error query User: ", tx.Error.Error())
		return User{}, tx.Error
	}

	if user.Id == 0 {
		return User{}, nil
	}

	return user, nil
}