package service

import (
	"Test-Golang-ITMX/model"
	"Test-Golang-ITMX/repository"
)

type CustomersServiceInterface interface {
	GetById(id int) (res model.CustomersResponse, Error error)
	GetAll() (res []model.CustomersResponse, Error error)
	Create(req model.CustomersRequest) (res model.CustomersResponse, Error error)
	Update(id int, req model.CustomersRequest) (res model.CustomersResponse, Error error)
	Delete(id int) (Error error)
}

type Service struct {
	repository repository.CustomersRepositoryInterface
}

func NewCustomersService(repository repository.CustomersRepositoryInterface) CustomersServiceInterface {
	return Service{
		repository: repository,
	}
}
