package repository

import (
	"Test-Golang-ITMX/model"
	"errors"
	"gorm.io/gorm"
)

func (a Repository) GetById(id int) (res model.Customers, Error error) {
	Error = a.db.Where("id = ?", id).First(&res).Error
	if Error != nil {
		if !errors.Is(Error, gorm.ErrRecordNotFound) {
			Error = Error
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}

func (a Repository) GetAll() (res []model.Customers, Error error) {
	Error = a.db.Find(&res).Error
	return
}

func (a Repository) Create(req model.CustomersRequest) (res model.Customers, Error error) {
	res = model.Customers{
		Name: req.Name,
		Age:  req.Age,
	}
	Error = a.db.Create(&res).Error
	if Error != nil {
		return
	}
	Error = a.db.Take(&res).Error
	if Error != nil {
		return
	}
	return
}

func (a Repository) Update(id int, req model.CustomersRequest) (res model.Customers, Error error) {
	res = model.Customers{
		Id:   id,
		Name: req.Name,
		Age:  req.Age,
	}
	Error = a.db.Where("id = ? ", id).Updates(&res).Error
	if Error != nil {
		return
	}
	Error = a.db.Take(&res).Error
	if Error != nil {
		return
	}

	return
}

func (a Repository) Delete(id int) (Error error) {
	Error = a.db.Where("id = ? ", id).Delete(&model.Customers{}).Error
	if Error != nil {
		return
	}
	return
}
