package service

import (
	"fmt"
	"github.com/myrachanto/ddd/httperors"
	"github.com/myrachanto/ddd/model"
	r "github.com/myrachanto/ddd/repository"
)

var (
	CategoryService CategoryserviceInterface = &categoryService{}
	// repo = r.ChooseRepo()

)
type CategoryserviceInterface interface{
	Create(category *model.Category) (*httperors.HttpError)
	GetOne(id string) (category *model.Category, errors *httperors.HttpError)
	GetAll(code string) ([]*model.Category, *httperors.HttpError)
	Update(code string, category *model.Category) (*httperors.HttpError)
	Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError)
}


type categoryService struct {
	respository r.CategoryInterface
}
func NewRedirectService(respository r.CategoryInterface) CategoryserviceInterface{
	return &categoryService{
		respository,
	}
}

func (service categoryService) Create(category *model.Category) (*httperors.HttpError) {
	
	fmt.Println(category)
	fmt.Println("--------service--------------")
	err1 := service.respository.Create(category)
	if err1 != nil {
		return err1
	}
	 return  nil

}

func (service categoryService) GetOne(id string) (*model.Category, *httperors.HttpError) {
	fmt.Println(id)
	category, err1 := service.respository.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return category, nil
}

func (service categoryService) GetAll(code string) ([]*model.Category, *httperors.HttpError) {
	categorys, err := service.respository.GetAll(code)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (service categoryService) Update(id string, category *model.Category) (*httperors.HttpError) {
	
	fmt.Println("update1-controller")
	fmt.Println(id)
	err1 := service.respository.Update(id, category)
	if err1 != nil {
		return err1
	}
	
	return nil
}
func (service categoryService) Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := service.respository.Delete(id)
		return success, failure
}
