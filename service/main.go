package service

import (
	"Test-Golang-ITMX/access"
	"Test-Golang-ITMX/model"
)

type CustomersServiceInterface interface {
	GetById(id int) (res model.CustomersResponse, Error error)
	GetAll() (res []model.CustomersResponse, Error error)
	Create(req model.CustomersRequest) (res model.CustomersResponse, Error error)
	Update(id int, req model.CustomersRequest) (res model.CustomersResponse, Error error)
	Delete(id int) (Error error)
}

type Service struct {
	access access.CustomersAccessInterface
}

func CustomersService(access access.CustomersAccessInterface) CustomersServiceInterface {
	return Service{
		access: access,
	}
}
