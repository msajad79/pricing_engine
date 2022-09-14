package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/msajad79/pricing_engine/config"

	"github.com/msajad79/pricing_engine/controllers"
	"github.com/msajad79/pricing_engine/models"
	"github.com/msajad79/pricing_engine/validators"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnect() {
	dsn := "root:sajad@tcp(127.0.0.1:3306)/test?"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	config.DB = db

	//migratrations
	models.MigrateModels()
}

func main() {
	server := gin.Default()
	DatabaseConnect()

	//ValidatorConfig()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		config.Val = v
		validators.ValidateRegister()
	}

	route_add := server.Group("/add")
	{
		route_add.POST("/rules", controllers.AddRulesController)
	}

	server.Run("localhost:8080")
}
