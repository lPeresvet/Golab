package app

import (
	"SimpleWebProject/internal/config"
	"SimpleWebProject/internal/repository"
	"SimpleWebProject/internal/service"
	"SimpleWebProject/internal/transport/rest/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func RunApp() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	userHandler.InitRoutes(app)

	config.SetupSwagger(app)

	port := viper.GetString("http.port")

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
