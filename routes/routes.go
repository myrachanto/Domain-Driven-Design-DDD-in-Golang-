package routes

import (
	// "log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/ddd/controllers"
	"github.com/myrachanto/ddd/repository"
	"github.com/myrachanto/ddd/service"
	log "github.com/sirupsen/logrus"
)
  
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		
	})
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetReportCaller(true)
	
  }
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
	log.Info("app initialized")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 

	// Routes
	e.POST("/categorys", controller.Create)
	e.GET("/categorys", controller.GetAll)
	e.GET("/categorys/:id", controller.GetOne)
	e.GET("/products", controller.Getproducts)
	e.PUT("/categorys/:id", controller.Update)
	e.DELETE("/categorys/:id", controller.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
