package repository

import (
	"Test-Golang-ITMX/model"
	"gorm.io/gorm"
)

type CustomersRepositoryInterface interface {
	GetById(id int) (res model.Customers, Error error)
	GetAll() (res []model.Customers, Error error)
	Create(req model.CustomersRequest) (res model.Customers, Error error)
	Update(id int, req model.CustomersRequest) (res model.Customers, Error error)
	Delete(id int) (Error error)
}

type Repository struct {
	db *gorm.DB
}

func NewCustomersRepository(db *gorm.DB) CustomersRepositoryInterface {
	return Repository{
		db: db,
	}
}
