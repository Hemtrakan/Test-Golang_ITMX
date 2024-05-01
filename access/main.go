package access

import (
	"Test-Golang-ITMX/model"
	"gorm.io/gorm"
)

type CustomersAccessInterface interface {
	GetById(id int) (res model.Customers, Error error)
	GetAll() (res []model.Customers, Error error)
	Create(req model.CustomersRequest) (res model.Customers, Error error)
	Update(id int, req model.CustomersRequest) (res model.Customers, Error error)
	Delete(id int) (Error error)
}

type Access struct {
	db *gorm.DB
}

func CustomersAccess(db *gorm.DB) CustomersAccessInterface {
	return Access{
		db: db,
	}
}
