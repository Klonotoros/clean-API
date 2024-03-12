package main

import (
	"clean-API/internal/client"
	"clean-API/internal/controller"
	"clean-API/internal/database"
	"clean-API/internal/dto"
	"clean-API/internal/repository"
	"clean-API/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := dto.Config{SigningSecret: os.Getenv("SECRET_KEY")}

	db := database.InitDB()

	server := gin.Default()

	clients := client.NewClients(config)
	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config, clients)
	controllers := controller.NewControllers(services, config)

	controllers.Route(server)
	err = server.Run(":8080")

	if err != nil {
		logrus.Info(err)
	}

}
