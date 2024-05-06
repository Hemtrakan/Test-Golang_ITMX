package repository

import (
	"Test-Golang-ITMX/model"
	"github.com/stretchr/testify/mock"
)

type customersRepositoryMock struct {
	mock.Mock
}

func NewCustomersRepositoryMock() *customersRepositoryMock {
	return &customersRepositoryMock{}
}

func (m *customersRepositoryMock) GetById(id int) (res model.Customers, Error error) {
	args := m.Called(id)
	res = args.Get(0).(model.Customers)
	Error = args.Error(1)
	return
}

func (m *customersRepositoryMock) GetAll() (res []model.Customers, Error error) {
	args := m.Called()
	res = args.Get(0).([]model.Customers)
	Error = args.Error(1)
	return
}

func (m *customersRepositoryMock) Create(req model.CustomersRequest) (res model.Customers, Error error) {
	args := m.Called(req)
	res = args.Get(0).(model.Customers)
	Error = args.Error(1)
	return
}

func (m *customersRepositoryMock) Update(id int, req model.CustomersRequest) (res model.Customers, Error error) {
	args := m.Called(id, req)
	res = args.Get(0).(model.Customers)
	Error = args.Error(1)
	return
}

func (m *customersRepositoryMock) Delete(id int) (Error error) {
	args := m.Called(id)
	Error = args.Error(0)
	return
}
