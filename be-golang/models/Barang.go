package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	Id          int64  `form:"Id" json:"Id" param:"Id" gorm:"column:Id"`
	Code        string  `form:"Code" json:"Code" gorm:"column:Code"`
	Name        string  `form:"Name" json:"Name" gorm:"column:Name"`
	Total       int64  `form:"Total" json:"Total" gorm:"column:Total"`
	Description string  `form:"Description" json:"Description" gorm:"column:Description"`
	IsActive    int  `form:"IsActive" json:"IsActive" gorm:"column:IsActive"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"column:CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"column:UpdatedAt"`
	DeletedAt   time.Time `json:"DeletedAt" gorm:"column:DeletedAt"`
}

type BarangFilter struct {
	Limit  int `json:"Limit"`
	Offset int `json:"Offset"`
	OrderBy string `json:"OrderBy"`
	Search string `json:"Search"`
}

const TableBarang = "barang"

func BarangModel() *Barang {
	return &Barang{}
}

func (*Barang) GetListBarang(filter BarangFilter) (listBarang []Barang, err error) {
	log.Println("Model : Get List Barang")
	tx := DBAsaba.Select("*").Table(TableBarang)
	if filter.Search != "" {
		tx.Where("Code LIKE ?", "%"+filter.Search+"%").Or("Name LIKE ?", "%"+filter.Search+"%")
	}
	if filter.Limit > 0 {
		tx.Limit(filter.Limit)
	} else {
		tx.Limit(10)
	}
	if filter.Offset > 0 {
		tx.Offset(filter.Offset)
	} else {
		tx.Offset(0)
	}
	if filter.OrderBy != "" {
		tx.Order(filter.OrderBy)
	} else {
		tx.Order("Name ASC")
	}
	tx.Find(&listBarang)
	
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		log.Println("Error query Barang: ", tx.Error.Error())
		return []Barang{}, tx.Error
	}

	return listBarang, nil
}

func (*Barang) FindBarang(request Barang) (barang Barang, err error) {
	log.Println("Model : Find Barang")
	tx := DBAsaba.Table(TableBarang).Where("Id = ?", request.Id).Where("DeletedAt IS NULL").Find(&barang)
	if tx.Error != nil {
		log.Println("Error query Barang: ", tx.Error.Error())
		return Barang{}, tx.Error
	}

	if barang.Id == 0 {
		return Barang{}, nil
	}

	return barang, nil
}

func (*Barang) CreateBarang(request Barang) (barang Barang, err error) {
	log.Println("Model : Create Barang")
	tx := DBAsaba.Table(TableBarang).Create(&request)
	if tx.Error != nil {
		log.Println("Error query Barang: ", tx.Error.Error())
		return Barang{}, tx.Error
	}
	
	tx = DBAsaba.Select("*").Table(TableBarang).Last(&barang)
	if tx.Error != nil {
		log.Println("Error query Barang: ", tx.Error.Error())
		return Barang{}, tx.Error
	}

	return barang, nil
}

func (*Barang) CheckCodeBarang(request Barang) (barang Barang, err error) {
	log.Println("Model : Check Code Barang")
	tx := DBAsaba.Table(TableBarang).Where("Code = ?", request.Code).Find(&barang)
	if tx.Error != nil {
		log.Println("Error query Barang: ", tx.Error.Error())
		return Barang{}, tx.Error
	}

	return barang, nil
}

func (*Barang) UpdateBarang(request Barang) (barang Barang, err error) {
	log.Println("Model : Update Barang")
	tx := DBAsaba.Table(TableBarang).Where("Id = ?", request.Id).Save(&request)
	if tx.Error != nil {
		log.Println("Error query Barang: ", tx.Error.Error())
		return Barang{}, tx.Error
	}

	return barang, nil
}