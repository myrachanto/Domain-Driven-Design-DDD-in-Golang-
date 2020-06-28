package model

import(
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/myrachanto/asoko/categorymicro/httperors"
)

type Category struct {
	Id bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
type Base struct{
	Created_At time.Time `bson:"created_at"`
	Updated_At time.Time `bson:"updated_at"`
	Delete_At *time.Time `bson:"deleted_at"`

}
func (category Category) Validate() *httperors.HttpError{
	if category.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if category.Title == "" {
		return httperors.NewNotFoundError("Invalid title")
	}
	if category.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	return nil
}