package controllers

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myrachanto/ddd/httperors"
	"github.com/myrachanto/ddd/model"
	s "github.com/myrachanto/ddd/service"
)
 
var (
	CategoryController categoryController = categoryController{}
)
type categoryController struct{
	service s.CategoryserviceInterface
}
type CategoryInterface interface {
	Create(echo.Context) error
	GetAll(echo.Context) error
	GetOne(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
	Getproducts(echo.Context) error
}
func NewController(service s.CategoryserviceInterface)CategoryInterface{
 return &categoryController{
	 service: service,
 }
}
/////////controllers/////////////////
func (controller categoryController) Create(c echo.Context) error {
	category := &model.Category{}
	if err := c.Bind(category); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println("cats................", category)
	err1 := controller.service.Create(category)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifully")
}

func (controller categoryController) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	categorys, err3 := controller.service.GetAll(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, categorys)
} 
func (controller categoryController) Getproducts(c echo.Context) error {
	search := c.QueryParam("search")
	categorys, err3 := controller.service.Getproducts(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, categorys)
} 
func (controller categoryController) GetOne(c echo.Context) error {
	id := c.Param("id")
	category, problem := controller.service.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, category)	
}

func (controller categoryController) Update(c echo.Context) error {
	category :=  &model.Category{}
	if err := c.Bind(category); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := c.Param("id")
	problem := controller.service.Update(id, category)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, "updated successifully")
}

func (controller categoryController) Delete(c echo.Context) error {
	id := c.Param("id")
	success, failure := controller.service.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}