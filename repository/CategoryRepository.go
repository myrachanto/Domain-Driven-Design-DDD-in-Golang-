package repository

import (
	"fmt"
	"strconv"

	"github.com/myrachanto/ddd/httperors"
	"github.com/myrachanto/ddd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Categoryrepository ...
var (
	Categoryrepository CategoryInterface = &categoryrepository{}
)

type CategoryInterface interface {
	Create(category *model.Category) *httperors.HttpError
	GetOne(id string) (category *model.Category, errors *httperors.HttpError)
	GetAll(code string) ([]*model.Category, *httperors.HttpError)
	Update(code string, category *model.Category) *httperors.HttpError
	Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError)
}

func New() *categoryrepository {
	return &categoryrepository{}
}

type categoryrepository struct{}

func (r *categoryrepository) Create(category *model.Category) *httperors.HttpError {
	// if err1 := category.Validate(); err1 != nil {
	// 	return err1
	// }
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return e
	}
	code, err1 := r.genecode()
	if err1 != nil {
		return err1
	}
	category.Code = code
	collection := db.Collection("category")
	_, err := collection.InsertOne(ctx, category)
	if err != nil {
		return httperors.NewBadRequestError(fmt.Sprintf("Create Category Failed, %d", err))
	}
	Mongorepo.DbClose(c)
	return nil
}

func (r *categoryrepository) GetOne(id string) (category *model.Category, errors *httperors.HttpError) {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return nil, t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, httperors.NewNotFoundError("primitive issue")
	}
	filter := bson.M{"_id": idPrimitive}
	err = collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		return nil, httperors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	Mongorepo.DbClose(c)
	return category, nil
}

func (r *categoryrepository) GetAll(code string) ([]*model.Category, *httperors.HttpError) {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return nil, t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return nil, e
	}
	categorys := []*model.Category{}
	collection := db.Collection("category")
	filter := bson.M{}
	fmt.Println("Brrrrrrrrrrrrrrrrrrr", filter)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, httperors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var category model.Category
		err := cur.Decode(&category)
		if err != nil {
			return nil, httperors.NewNotFoundError("Error while decoding results!")
		}
		categorys = append(categorys, &category)
	}
	if err := cur.Err(); err != nil {
		return nil, httperors.NewNotFoundError("Error with cursor!")
	}
	Mongorepo.DbClose(c)
	return categorys, nil

}

func (r *categoryrepository) Update(code string, category *model.Category) *httperors.HttpError {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return e
	}
	result, err3 := r.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result)
	if category.Name == "" {
		category.Name = result.Name
	}
	if category.Title == "" {
		category.Title = result.Title
	}
	if category.Description == "" {
		category.Description = result.Description
	}
	if category.Code == "" {
		category.Code = result.Code
	}
	collection := db.Collection("category")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(category)
	update := bson.M{"$set": category}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return httperors.NewNotFoundError("Error updating!")
	}
	Mongorepo.DbClose(c)
	return nil
}
func (r categoryrepository) Delete(id string) (*httperors.HttpSuccess, *httperors.HttpError) {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return nil, t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	Mongorepo.DbClose(c)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (r categoryrepository) genecode() (string, *httperors.HttpError) {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return "", t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return "", e
	}
	collection := db.Collection("category")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil {
		return "", httperors.NewNotFoundError("no results found")
	}
	code := "CategoryCode" + strconv.FormatUint(uint64(co), 10)

	Mongorepo.DbClose(c)
	return code, nil
}
func (r categoryrepository) getuno(code string) (result *model.Category, err *httperors.HttpError) {
	c, t := Mongorepo.Mongoclient()
	if t != nil {
		return nil, t
	}
	db, e := Mongorepo.Mongodb()
	if e != nil {
		return nil, e
	}
	collection := db.Collection("category")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("no results found")
	}
	Mongorepo.DbClose(c)
	return result, nil
}
