package repository
import (
	"log"
	"os"
	mgo "gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
	"github.com/myrachanto/asoko/categorymicro/httperors"
	"github.com/myrachanto/asoko/categorymicro/model"
)

type Redirectrepository interface{
	Create(category *model.Category) (*model.Category, *httperors.HttpError)
	GetOne(id string) (*model.Category, *httperors.HttpError)
	GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError)
	Update(id string, category *model.Category) (*model.Category, *httperors.HttpError)
	Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError)
}

/////////////////////////////////////////////////////////////////////////////////////
////////////////figure how to switch repositories automatically//////////////////////////////////
func ChooseRepo() (repository Redirectrepository) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	switch os.Getenv("DbType2") {
	case "Mongo":
		_, err1 := NewMongoRepository()
		if err1 != nil {
			log.Fatal(err1)
		}
		repository = Mongorepository
		return repository
	// case "mysql":
	// 	_, err1 := NewGormRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	// 	repository = Sqlrepository
	// 	// model.CheckMongo(gorm)
	// 	return repository
	
	// case "postgress":
	// 	repository, err1 := NewMongoRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	// 	return repository
	// case "redis":
	// 	repository, err1 := NewMongoRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	}
	return
	
}
func NewMongoRepository()(Redirectrepository, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Mongo := os.Getenv("MongoDb")
	host := os.Getenv("Mongohost")

	_, err = mgo.Dial(host)
	if err != nil{
		return nil, err
	}
	return Mongorepository, nil
}