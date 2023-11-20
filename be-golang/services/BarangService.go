package services

import (
	"be-golang/models"
	"log"
	"net/http"
	"time"
)

func GetListBarang(filter models.BarangFilter) (data []models.Barang, message string, status int){
	log.Println("Service : Get List Barang")
	barangModel := models.BarangModel()
	data, err := barangModel.GetListBarang(filter)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
		return
	}

	if len(data) > 0 {
		status = http.StatusOK
		message = "Success"
	} else {
		status = http.StatusNotFound
		message = "Data tidak ditemukan"
	}
	
	return
}

func CreateBarang(request []models.Barang) (data models.Barang, message string, status int){
	log.Println("Service : Create Barang")
	barangModel := models.BarangModel()
	historyModel := models.HistoryModel()

	for _, barang := range request {
		// Check Code
		data, err := barangModel.CheckCodeBarang(barang)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			break
		}
		if data.Code != "" {
			data = models.Barang{}
			status = http.StatusBadRequest
			message = "Code sudah ada"
			break
		}
	
		barang.CreatedAt = time.Now()
		data, err = barangModel.CreateBarang(barang)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			break
		} else {
			status = http.StatusOK
			message = "Sukses tambah data barang"
			go historyModel.CreateHistory(models.History{
				IdBarang: data.Id,
				Difference: data.Total,
				Type: "TAMBAH",
				CreatedAt: time.Now(),
			})
		}
	}

	return
}

func UpdateBarang(request []models.Barang) (data models.Barang, message string, status int){
	log.Println("Service : Update Barang")
	barangModel := models.BarangModel()
	historyModel := models.HistoryModel()

	for _, barang := range request {
		// Find By Id
		data, err := barangModel.FindBarang(barang)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			break
		}
		
		if data.Id == 0 {
			status = http.StatusNotFound
			message = "Data tidak ditemukan"
			break
		}
		
		if barang.Total < 0 {
			status = http.StatusNotFound
			message = "Jumlah barang tidak mencukupi"
			break
		}
	
		// Check Code
		data, err = barangModel.CheckCodeBarang(barang)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			break
		}
		if data.Code != "" && data.Id != barang.Id {
			data = models.Barang{}
			status = http.StatusBadRequest
			message = "Code sudah ada"
			break
		}
	
		barang.UpdatedAt = time.Now()
		data, err = barangModel.UpdateBarang(barang)
		if err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
			break
		} else {
			data = barang
			status = http.StatusOK
			message = "Sukses update data barang"
			tempSelisih := data.Total - barang.Total
			tempType := "TAMBAH"
			if data.Total > barang.Total {
				tempType = "KURANG"
			}
			go historyModel.CreateHistory(models.History{
				IdBarang: data.Id,
				Difference: tempSelisih,
				Type: tempType,
				CreatedAt: time.Now(),
			})
		}
	}
	return
}

func DeleteBarang(request models.Barang) (data models.Barang, message string, status int){
	log.Println("Service : Delete Barang")
	barangModel := models.BarangModel()

	// Find By Id
	data, err := barangModel.FindBarang(request)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
		return
	}
	
	if data.Id == 0 {
		status = http.StatusNotFound
		message = "Data tidak ditemukan"
		return
	}

	data.DeletedAt = time.Now()
	data, err = barangModel.UpdateBarang(data)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
		return
	} else {
		status = http.StatusOK
		message = "Sukses hapus data barang"
		return
	}
}

