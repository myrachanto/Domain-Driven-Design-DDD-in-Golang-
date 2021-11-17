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
	Create(category *model.Category) (*model.Category,*httperors.HttpError)
	GetOne(id string) (*model.Category, *httperors.HttpError)
	GetAll() ([]*model.Category, *httperors.HttpError)
	Update(code string, category *model.Category) (*httperors.HttpError)
}
func (mock MockRepository)Create(category *model.Category) (*model.Category,*httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	category, err := result.(*model.Category), args.Error(1)
	if category.Title == "" {
		return nil, httperors.NewNotFoundError("test failed category title empty")
	}
	if err != nil {
		return nil, httperors.NewNotFoundError("test failed")
	}
	return category, nil
}
func (mock MockRepository)GetOne(id string) (*model.Category, *httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	category, err := result.(*model.Category), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("test failed")
	}
	return category,nil
}
func (mock MockRepository)GetAll() ([]*model.Category, *httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	categorys, err := result.([]*model.Category), args.Error(1)
	if err != nil {
		return nil, httperors.NewNotFoundError("test failed")
	}
	return categorys, nil
}
func (mock MockRepository)Update(code string, category *model.Category) (*httperors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	category, err := result.(*model.Category), args.Error(1)
	if category.Title == "" {
		return httperors.NewNotFoundError("test failed category title empty")
	}
	if err != nil {
		return httperors.NewNotFoundError("test failed")
	}
	return nil
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
func TestGetAll(t *testing.T){
	mockRepo := new(MockRepository)
	category := &model.Category{
		Name: "name",
		Title: "tiltes",
		Description: "description",
	}
	//set up expecctations
	mockRepo.On("GetAll").Return([]*model.Category{category} ,nil)
	results, _ := mockRepo.GetAll()
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, "name", results[0].Name)
	assert.Equal(t, "tiltes", results[0].Title)
	assert.Equal(t, "description", results[0].Description)


}
func TestCreate(t *testing.T){
	mockRepo := new(MockRepository)
	category := &model.Category{
		Name: "name",
		Title: "tiltes",
		Description: "description",}
		mockRepo.On("Create").Return(category, nil)
		result, err := mockRepo.Create(category)
		//mock assertion: behavioral
		mockRepo.AssertExpectations(t)
		//data assertion
		assert.Equal(t, "name", result.Name)
		assert.Equal(t, "tiltes", result.Title)
		assert.Equal(t, "description", result.Description)
		assert.Nil(t, err)

}