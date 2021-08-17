package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/ddd/controllers"
	"github.com/myrachanto/ddd/service"
	"github.com/myrachanto/ddd/repository"
)

func ApiMicroservice() {
	// repo := repository.New()
	// Service := service.NewRedirectService(&repo)
	// controller := controllers.NewController(Service)
	controller := controllers.NewController(service.NewRedirectService(repository.New()))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 

	// Routes
	e.POST("/categorys", controller.Create)
	e.GET("/categorys", controller.GetAll)
	e.GET("/categorys/:id", controller.GetOne)
	e.PUT("/categorys/:id", controller.Update)
	e.DELETE("/categorys/:id", controller.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
