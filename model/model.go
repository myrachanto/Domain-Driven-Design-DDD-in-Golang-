package model

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/myrachanto/ddd/httperors"
)

type Category struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Code        string             `json:"code"`
	Base        `json:"base"`
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
func (c Category) ValidateTitle() *httperors.HttpError{
	if c.Title == "" {
		return httperors.NewNotFoundError("Invalid title")
	}
	return nil
}
func (c Category) ValidateName() *httperors.HttpError{
	if c.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	return nil
}

func (c Category) ValidateDescription() *httperors.HttpError{
	if c.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	return nil
}