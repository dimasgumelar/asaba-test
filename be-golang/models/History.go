package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type History struct {
	Id          int64  `form:"Id" json:"Id" param:"Id" gorm:"column:Id"`
	IdBarang    int64  `form:"IdBarang" json:"IdBarang" param:"IdBarang" gorm:"column:IdBarang"`
	Difference  int64  `form:"Difference" json:"Difference" gorm:"column:Difference"`
	Type 		string  `form:"Type" json:"Type" gorm:"column:Type"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"column:CreatedAt"`
	DeletedAt   time.Time `json:"DeletedAt" gorm:"column:DeletedAt"`
}

type HistoryFilter struct {
	Limit  int `json:"Limit"`
	Offset int `json:"Offset"`
	OrderBy string `json:"OrderBy"`
	Search string `json:"Search"`
}

const TableHistory = "history"

func HistoryModel() *History {
	return &History{}
}

func (*History) GetListHistory(filter HistoryFilter) (listHistory []History, err error) {
	log.Println("Model : Get List History")
	tx := DBAsaba.Select("history.*, barang.Code, barang.Name").Table(TableHistory)
	tx.Joins("JOIN barang ON barang.Id = history.IdBarang")

	if filter.Search != "" {
		tx.Where("barang.Code LIKE ?", "%"+filter.Search+"%").Or("barang.Name LIKE ?", "%"+filter.Search+"%")
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
		tx.Order("barang.Name ASC")
	}
	tx.Find(&listHistory)
	
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		log.Println("Error query History: ", tx.Error.Error())
		return []History{}, tx.Error
	}

	return listHistory, nil
}

func (*History) FindHistory(request History) (barang History, err error) {
	log.Println("Model : Find History")
	tx := DBAsaba.Table(TableHistory).Where("Id = ?", request.Id).Where("DeletedAt IS NULL").Find(&barang)
	if tx.Error != nil {
		log.Println("Error query History: ", tx.Error.Error())
		return History{}, tx.Error
	}

	if barang.Id == 0 {
		return History{}, nil
	}

	return barang, nil
}

func (*History) CreateHistory(request History) (barang History, err error) {
	log.Println("Model : Create History")
	tx := DBAsaba.Table(TableHistory).Create(&request)
	if tx.Error != nil {
		log.Println("Error query History: ", tx.Error.Error())
		return History{}, tx.Error
	}
	
	tx = DBAsaba.Select("*").Table(TableHistory).Order("Id ASC").Last(&barang)
	if tx.Error != nil {
		log.Println("Error query History: ", tx.Error.Error())
		return History{}, tx.Error
	}

	return barang, nil
}

func (*History) UpdateHistory(request History) (barang History, err error) {
	log.Println("Model : Update History")
	tx := DBAsaba.Table(TableHistory).Where("Id = ?", request.Id).Save(&request)
	if tx.Error != nil {
		log.Println("Error query History: ", tx.Error.Error())
		return History{}, tx.Error
	}

	return barang, nil
}