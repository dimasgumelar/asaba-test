package services

import (
	"be-golang/helpers"
	"be-golang/models"
	"log"
	"net/http"
)

func Login(request models.User) (data models.User, message string, status int){
	log.Println("Service : Login")
	userModel := models.UserModel()

	data, err := userModel.FindUser(request)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
		return
	}

	if data.Id == 0 {
		status = http.StatusNotFound
		message = "User tidak ditemukan"
		return
	} else {
		if helpers.CheckPasswordHash(request.Password, data.Password) {
			status = http.StatusOK
			message = "Success"
			return
		} else {
			status = http.StatusUnauthorized
			message = "Password salah"
			return
		}
	}
}
