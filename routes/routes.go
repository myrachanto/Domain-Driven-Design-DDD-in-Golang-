package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/asoko/categorymicro/controllers"
)

func ApiMicroservice() {

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
	e.POST("/categorys", controllers.CategoryController.Create)
	e.GET("/categorys", controllers.CategoryController.GetAll)
	e.GET("/categorys/:id", controllers.CategoryController.GetOne)
	e.PUT("/categorys/:id", controllers.CategoryController.Update)
	e.DELETE("/categorys/:id", controllers.CategoryController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
