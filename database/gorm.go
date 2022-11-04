package database

import (
	"fmt"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDatabase() {
	username := configs.Cfg.DBUsername
	password := configs.Cfg.DBPassword
	address := configs.Cfg.DBAddress
	dbName := configs.Cfg.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, dbName)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	DB = db

	_ = DB.AutoMigrate(&model.User{}, &model.Renter{}, &model.Category{}, &model.Bike{}, &model.Payment{}, &model.Order{}, &model.OrderDetail{}, &model.Review{})
}
