package model

type Customers struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

type CustomersResponse struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type CustomersRequest struct {
	Name string `json:"name,omitempty" validate:"required"`
	Age  int    `json:"age,omitempty" validate:"required"`
}
