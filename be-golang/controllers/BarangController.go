package controllers

import (
	"be-golang/helpers"
	"be-golang/models"
	"be-golang/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListBarang(e echo.Context) error {
	log.Println("Controller : Get List Barang")
	var response helpers.JSONResponse
	var filter models.BarangFilter

	data, message, status := services.GetListBarang(filter)
	
	response.Data = data
	response.Message = message
	response.Status = status

	log.Println("Response : ", response)

	return e.JSON(response.Status, response)
}

func CreateBarang(e echo.Context) error {
	log.Println("Controller : Create Barang")
	var response helpers.JSONResponse
	var request []models.Barang

	err := e.Bind(&request)
	if err != nil {
		response.Data = nil
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
	}

	data, message, status := services.CreateBarang(request)
	
	if data.Code != "" {
		response.Data = data
	} else {
		response.Data = nil
	}
	response.Message = message
	response.Status = status

	log.Println("Response : ", response)

	return e.JSON(response.Status, response)
}

func UpdateBarang(e echo.Context) error {
	log.Println("Controller : Update Barang")
	var response helpers.JSONResponse
	var request []models.Barang
	
	err := e.Bind(&request)
	if err != nil {
		response.Data = nil
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
	}

	data, message, status := services.UpdateBarang(request)
	
	if data.Code != "" {
		response.Data = data
	} else {
		response.Data = nil
	}
	response.Message = message
	response.Status = status

	log.Println("Response : ", response)

	return e.JSON(response.Status, response)
}

func DeleteBarang(e echo.Context) error {
	log.Println("Controller : Delete Barang")
	var response helpers.JSONResponse
	var request models.Barang
	
	err := e.Bind(&request)
	if err != nil {
		response.Data = nil
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
	}

	data, message, status := services.DeleteBarang(request)
	
	if data.Code != "" {
		response.Data = data
	} else {
		response.Data = nil
	}
	response.Message = message
	response.Status = status

	log.Println("Response : ", response)

	return e.JSON(response.Status, response)
}

