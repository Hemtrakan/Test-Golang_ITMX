package config

import (
	"Test-Golang-ITMX/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

func InitConfig() {
	viper.SetConfigName("configLocal")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func ConnectDataBase() (db *gorm.DB, Error error) {
	db, Error = gorm.Open(sqlite.Open(viper.GetString("database.db.servername")+".db"), &gorm.Config{})
	if Error != nil {
		fmt.Println("Error : ", Error.Error())
		return
	}

	Error = db.AutoMigrate(&model.Customers{})
	return
}
