package main

import (
	"dept-collector/internal/api"
	"dept-collector/internal/config"
	"dept-collector/internal/pkg/validator"
	"log"

	_ "dept-collector/docs"

	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           DeptCollector
// @version         Alpha
// @description     API for the DebtCollector app
// @host            localhost:8080
// @BasePath        /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("go env loading failed (no .env file) 😥")
	}
	db := config.ConnectDB()
	config.AutoMigrate(db)

	validator.InitCustomValidators()

	router := api.NewRouter(db)
	router.Run()
}
