### Info 
```
Golang version 1.21
Main Framework
    - Gorm
        : https://gorm.io/index.html
    - Fiber
        : https://docs.gofiber.io/
    - Viper
        : https://github.com/spf13/viper
```

### How To Run Project 

```
 git clone https://github.com/Hemtrakan/Test-Golang_ITMX.git
 cd project Path
 go mod tidy 
 go run main.go
```

### How To Run UnitTest Service 
```
 go test ./service -cover
```

### script to initially create some customer data in the database.
> INSERT INTO customers (name, age) VALUES ('Letter', 25) , ('Mocha',26) , ('Moochi',30);

### Router

```
[Get] /api/v1/customers
[Get] /api/v1/customers/:id
[Post] /api/v1/customers
[Put] /api/v1/customers/:id
[Delete] /api/v1/customers/:id
```




