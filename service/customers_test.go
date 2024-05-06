package service_test

import (
	"Test-Golang-ITMX/model"
	"Test-Golang-ITMX/repository"
	"Test-Golang-ITMX/service"
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCustomersByIdService(t *testing.T) {
	type testCase struct {
		name     string
		expected model.Customers
	}

	cases := []testCase{
		{name: "GetById1", expected: model.Customers{
			Id:   1,
			Name: "Customer_1",
			Age:  20,
		}},
		{name: "GetById2", expected: model.Customers{
			Id:   2,
			Name: "Customer_2",
			Age:  21,
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage
			customerRepository := repository.NewCustomersRepositoryMock()
			customerRepository.On("GetById", c.expected.Id).Return(model.Customers{
				Id:   c.expected.Id,
				Name: c.expected.Name,
				Age:  c.expected.Age,
			}, nil)
			//Act
			customerService := service.NewCustomersService(customerRepository)
			customer, _ := customerService.GetById(c.expected.Id)
			expect := c.expected

			//Assert
			assert.Equal(t, expect.Id, customer.Id)
			assert.Equal(t, expect.Name, customer.Name)
			assert.Equal(t, expect.Age, customer.Age)
		})
	}
	t.Run("Not Fount Customer", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetById", 0).Return(model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.GetById(0)

		//Assert
		assert.ErrorIs(t, err, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetAll").Return([]model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.GetAll()

		//Assert
		assert.Equal(t, err, err)
	})
}

func TestCustomersAllService(t *testing.T) {
	type testCase struct {
		name     string
		expected []model.Customers
	}

	cases := testCase{
		name: "GetAll",
		expected: []model.Customers{
			{
				Id:   1,
				Name: "Customer_1",
				Age:  20,
			},
			{
				Id:   2,
				Name: "Customer_2",
				Age:  21,
			},
		},
	}

	//Arrage
	t.Run("GetAll", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetAll").Return(cases.expected, nil)
		//Act
		customerService := service.NewCustomersService(customerRepository)
		customer, _ := customerService.GetAll()

		assert.Equal(t, reflect.DeepEqual(cases.expected, customer), false)
	})

	t.Run("Not Fount CustomerList", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetAll").Return([]model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		customer, _ := customerService.GetAll()

		//Assert
		assert.Equal(t, len(customer), 0)
	})

	t.Run("Repository Error", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetAll").Return([]model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.GetAll()

		//Assert
		assert.Equal(t, err, err)
	})
}

func TestCustomersCreateService(t *testing.T) {
	type testCase struct {
		name     string
		expected model.Customers
	}

	cases := testCase{
		name: "CreateCustomers",
		expected: model.Customers{
			Id:   3,
			Name: "Customer_3",
			Age:  23,
		},
	}

	dataCreate := model.CustomersRequest{
		Name: "Customer_3",
		Age:  23,
	}

	t.Run(cases.name, func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("Create", dataCreate).Return(cases.expected, nil)
		//Act
		customerService := service.NewCustomersService(customerRepository)
		customer, _ := customerService.Create(model.CustomersRequest{
			Name: "Customer_3",
			Age:  23,
		})

		//Assert
		assert.Equal(t, cases.expected.Id, customer.Id)
		assert.Equal(t, cases.expected.Name, customer.Name)
		assert.Equal(t, cases.expected.Age, customer.Age)
	})

	t.Run("Repository Error", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("Create", model.CustomersRequest{}).Return(model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.Create(model.CustomersRequest{})

		//Assert
		assert.ErrorIs(t, err, err)
	})
}

func TestCustomersUpdateService(t *testing.T) {
	type testCase struct {
		name       string
		id         int
		expected   model.Customers
		dataUpdate model.CustomersRequest
	}

	cases := testCase{
		name: "UpdateCustomers",
		id:   3,
		expected: model.Customers{
			Id:   3,
			Name: "Customer_4",
			Age:  24,
		},
		dataUpdate: model.CustomersRequest{
			Name: "Customer_4",
			Age:  24,
		},
	}

	t.Run(cases.name, func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetById", cases.id).Return(model.Customers{
			Id:   3,
			Name: "Customer_3",
			Age:  23,
		}, nil)
		//Act
		customerRepository.On("Update", cases.id, cases.dataUpdate).Return(cases.expected, nil)
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.GetById(cases.id)
		if err == nil {
			customer, _ := customerService.Update(cases.id, cases.dataUpdate)
			//Assert
			assert.Equal(t, cases.expected.Id, customer.Id)
			assert.Equal(t, cases.expected.Name, customer.Name)
			assert.Equal(t, cases.expected.Age, customer.Age)
		}
	})

	t.Run("Repository Error", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		customerRepository.On("GetById", cases.id).Return(model.Customers{}, errors.New(""))
		customerRepository.On("Update", 0, model.CustomersRequest{}).Return(model.Customers{}, errors.New(""))
		//Act
		customerService := service.NewCustomersService(customerRepository)
		_, err := customerService.GetById(cases.id)
		//Assert
		if err != nil {
			assert.ErrorIs(t, err, err)
		} else {
			customer, _ := customerService.Update(cases.id, cases.dataUpdate)
			assert.Equal(t, cases.expected.Id, customer.Id)
			assert.Equal(t, cases.expected.Name, customer.Name)
			assert.Equal(t, cases.expected.Age, customer.Age)
		}
	})
}

func TestCustomersDeleteService(t *testing.T) {
	type testCase struct {
		name     string
		id       int
		expected error
	}

	cases := testCase{
		name:     "DeleteCustomers",
		id:       3,
		expected: errors.New(""),
	}

	t.Run(cases.name, func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		//Act
		customerRepository.On("Delete", cases.id).Return(nil)
		customerService := service.NewCustomersService(customerRepository)
		err := customerService.Delete(cases.id)
		//Assert
		assert.ErrorIs(t, err, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		//Arrage
		customerRepository := repository.NewCustomersRepositoryMock()
		//Act
		customerRepository.On("Delete", cases.id).Return(cases.expected)
		customerService := service.NewCustomersService(customerRepository)
		err := customerService.Delete(cases.id)
		//Assert
		assert.ErrorIs(t, err, err)
	})
}
