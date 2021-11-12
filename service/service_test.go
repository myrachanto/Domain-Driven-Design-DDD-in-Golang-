package service

import (
	
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/myrachanto/ddd/model" 
	"github.com/myrachanto/ddd/httperors"
)

type MockRepository struct {
	mock.Mock
}
var (
	category = &model.Category{
		Name: "name of the category",
		Title: "title",
		Description: "description of the category",
	}
)
type CategoryMockInterface interface{
	Create(category *model.Category) (*httperors.HttpError)
	GetOne(id string) (category *model.Category, errors *httperors.HttpError)
	GetAll(code string) ([]*model.Category, *httperors.HttpError)
	Update(code string, category *model.Category) (*httperors.HttpError)
	Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError)
	Getproducts(code string) ([]model.Product, *httperors.HttpError)
}
func (mock MockRepository)Create(category *model.Category) (*model.Category,*httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	blog, err := result.(model.Category), args.Error(1)
	if err != nil {
		return nil, httperors.NewBadRequestError("Something went wrong creating a product")
	}
	return &blog, nil
}
func (mock MockRepository)GetOne(id string) (*model.Category, *httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	category, err := result.(model.Category), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("test failed")
	}
	return &category,nil
}
func TestCategoryNameValidate(t *testing.T){ 
	category = &model.Category{
		Name: "",
		Title: "",
		Description: "",
	}
	err := category.ValidateName()
	expected := "Invalid Name"
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Message)
}
func TestCategoryTitleValidate(t *testing.T){ 
	category = &model.Category{
		Name: "Name",
		Title: "",
		Description: "",
	}
	err := category.ValidateTitle()
	expected := "Invalid title"
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Message)
}
func TestCategoryDescriptionValidate(t *testing.T){ 
	category = &model.Category{
		Name: "name",
		Title: "tilte",
		Description: "",
	}
	err := category.ValidateDescription()
	expected := "Invalid Description"
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Message)
}
func TestCreatedCategory(t *testing.T){
	mockRepo := new(MockRepository)	
	//set up expecctations
	mockRepo.On("Create").Return([]*model.Category{category} ,nil)
	results, _ := mockRepo.Create(category)
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, "name of the category", results.Name)
	assert.Equal(t, "title", results.Title)
	assert.Equal(t, "description of the category", results.Description)

}
func TestCreate(t *testing.T){
	mockRepo := new(MockRepository)	
	//set up expecctations
	mockRepo.On("Create").Return([]*model.Category{category} ,nil)
	_, err := mockRepo.Create(category)
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, err.Message, "Something went wrong creating a product")

}