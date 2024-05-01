package main

import (
	"Test-Golang-ITMX/config"
	"Test-Golang-ITMX/handler"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitConfig()
	db, _ := config.ConnectDataBase()
	app := fiber.New()

	NewRouter(app, db)
	addr := fmt.Sprintf(":%v", viper.GetString("app.port"))

	go func() {
		err := app.Listen(addr)
		if err != nil {
			fmt.Printf("Error : %v", err)
			return
		}
	}()

	fmt.Printf("start %v success ", viper.GetString("app.name"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

func NewRouter(c *fiber.App, db *gorm.DB) {
	api := c.Group("/api/v1")
	handler.APICustomersHandler(api, db)
}
