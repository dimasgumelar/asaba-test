package controllers

import (
	"be-golang/helpers"
	"be-golang/models"
	"be-golang/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(e echo.Context) error {
	log.Println("Starting Login")
	var response helpers.JSONResponse
	var request models.User

	err := e.Bind(&request)
	if err != nil {
		response.Data = nil
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
	}
	
	data, message, status := services.Login(request)

	if data.Id != 0 {
		response.Data = data
	} else {
		response.Data = nil
	}
	response.Message = message
	response.Status = status

	return e.JSON(response.Status, response)
}