package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/ddd/controllers"
	"github.com/myrachanto/ddd/repository"
	"github.com/myrachanto/ddd/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
type Port struct {
	PORT string `mapstructure:"PORT"`
}
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		
	})
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetReportCaller(true)
	
  }
  func LoadConfig() (port Port, err error) {
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&port)
	return
}
func ApiMicroservice() {
	_,err := repository.Mongorepo.Mongoclient()
	if err != nil{
		log.Panic("Database failed to connect")
	}
	// repo := repository.New() 
	// Service := service.NewRedirectService(&repo)
	// controller := controllers.NewController(Service)
	controller := controllers.NewController(service.NewRedirectService(repository.New()))
	open, err1 := LoadConfig()
	if err1 != nil {
		log.Fatal("cannot load config:", err)
	}
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
	e.Logger.Fatal(e.Start(open.PORT))
}
