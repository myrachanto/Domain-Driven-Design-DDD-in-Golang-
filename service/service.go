package service

import (
	"fmt"
	"github.com/myrachanto/asoko/categorymicro/httperors"
	"github.com/myrachanto/asoko/categorymicro/model"
	r "github.com/myrachanto/asoko/categorymicro/repository"
)

var (
	CategoryService categoryService = categoryService{}
	repo = r.ChooseRepo()

)
type Redirectcategory interface{
	Create(category *model.Category) (*model.Category, *httperors.HttpError)
	GetOne(id string) (*model.Category, *httperors.HttpError)
	GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError)
	Update(id string, category *model.Category) (*model.Category, *httperors.HttpError)
	Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError)
}


type categoryService struct {
	respository r.Redirectrepository
}
func NewRedirectService(respository r.Redirectrepository) Redirectcategory{
	return &categoryService{
		respository,
	}
}

func (service categoryService) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	
	fmt.Println(category)
	fmt.Println("--------service--------------")
	category, err1 := repo.Create(category)
	if err1 != nil {
		return nil, err1
	}
	 return category, nil

}

func (service categoryService) GetOne(id string) (*model.Category, *httperors.HttpError) {
	fmt.Println(id)
	category, err1 := repo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return category, nil
}

func (service categoryService) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	categorys, err := repo.GetAll(categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (service categoryService) Update(id string, category *model.Category) (*model.Category, *httperors.HttpError) {
	
	fmt.Println("update1-controller")
	fmt.Println(id)
	category, err1 := repo.Update(id, category)
	if err1 != nil {
		return nil, err1
	}
	
	return category, nil
}
func (service categoryService) Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := repo.Delete(id)
		return success, failure
}
