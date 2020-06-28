package repository

import (
	"fmt"
	"log"
	"os"
	"gopkg.in/mgo.v2/bson"
	"github.com/joho/godotenv"
	 mgo "gopkg.in/mgo.v2"
	"github.com/myrachanto/asoko/categorymicro/httperors"
	"github.com/myrachanto/asoko/categorymicro/model"
	s "github.com/myrachanto/asoko/categorymicro/support"
)

var (
	Mongorepository mongorepository = mongorepository{}
)

///curtesy to gorm
type mongorepository struct{}

func GetMongoDB() (*mgo.Database, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Mongo := os.Getenv("MongoDb")
	host := os.Getenv("Mongohost")
	dbName := os.Getenv("MongodbName")

	session, err := mgo.Dial(host)
	if err != nil{
		return nil, err
	}
	db := session.DB(dbName)
	return db, nil
}
func (repository mongorepository) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	db, err := GetMongoDB()
	if err != nil {
		return nil, httperors.NewBadRequestError("Mongo db connection failed")
	}else{ 
		CategoryModel := s.CategoryModel{
			Db: db,
			Collection: "category",
		}
		if err := category.Validate(); err != nil {
			return nil, err
		}		
		fmt.Println(category)
		fmt.Println("------------------------")
		category.Id = bson.NewObjectId()
		err3 := CategoryModel.Create(category)
		if err3 != nil {
			return nil, httperors.NewBadRequestError(fmt.Sprintf("Create category Failed, %d", err3))
		} else {
			return category, nil
		}
	}
}

func (repository mongorepository) GetOne(id string) (*model.Category, *httperors.HttpError) {
	db, err := GetMongoDB()
	if err != nil {
		return nil, httperors.NewBadRequestError("Mongo db connection failed")
	}else{
		CategoryModel := s.CategoryModel{
			Db: db,
			Collection: "category",
		}
		category, err2 := CategoryModel.Find(id)
		if err2 != nil {
			return nil,	httperors.NewNotFoundError("something went wrong")
		}
		if category.Id == "" {
			return nil, httperors.NewNotFoundError("No results found")
		}
		return &category, nil
	}
}

func (repository mongorepository) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	db, err := GetMongoDB()
	if err != nil {
		return nil, httperors.NewNotFoundError("Mongo db connection failed")
	}
	CategoryModel := s.CategoryModel{
		Db: db,
		Collection: "category",
	}
	categorys, err2 := CategoryModel.FindAll()
	if err2 != nil {
		return nil,	httperors.NewNotFoundError("no results found")
	}
	return categorys, nil

}

func (repository mongorepository) Update(id string, category *model.Category) (*model.Category, *httperors.HttpError) {
	db, err := GetMongoDB()
	if err != nil {
		return nil, httperors.NewBadRequestError("Mongo db connection failed")
	}else{
		CategoryModel := s.CategoryModel{
			Db: db,
			Collection: "category",
		}
		ucategory, err2 := CategoryModel.Find(id)
		if err2 != nil {
			return nil, httperors.NewBadRequestError("No category found with that id")
		}
		if category.Name  == "" {
			category.Name = ucategory.Name
		}
		if category.Title  == "" {
			category.Title = ucategory.Title
		}
	
		if category.Description  == "" {
			category.Description = ucategory.Description
		}
		CategoryModel = s.CategoryModel{
			Db: db,
			Collection: "category",
		}
		err3 := CategoryModel.Update(category)
		if err3 != nil {
			return nil, httperors.NewBadRequestError("update failed for the categorywith that id")
		} else {
			return category, nil
		}
	}
}
func (repository mongorepository) Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError) {
	db, err := GetMongoDB()
	if err != nil {
		return nil, httperors.NewBadRequestError("Mongo db connection failed")
	}else{
		CategoryModel := s.CategoryModel{
			Db: db,
			Collection: "category",
		}
		category, err4 := CategoryModel.Find(id)
		if err4 != nil {
			return nil, httperors.NewNotFoundError("could not find categorys with that id" )
		}
		if category.Id == "" {
			return nil, httperors.NewNotFoundError("No results found")
		}
		err2 := CategoryModel.Delete(category)
		if err2 != nil {
			return nil, httperors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err2))
		}else{
			return httperors.NewSuccessMessage("deleted successfully"), nil
		}
	}
}

