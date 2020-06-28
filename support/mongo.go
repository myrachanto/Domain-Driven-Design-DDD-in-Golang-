package models
import ( 
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/myrachanto/asoko/categorymicro/model"
) 

type CategoryModel struct {
	Db *mgo.Database
	Collection string
}
func (categoryModel CategoryModel) FindAll() (categorys []model.Category, err error){
	err = categoryModel.Db.C(categoryModel.Collection).Find(bson.M{}).All(&categorys)
	return
}

func (categoryModel CategoryModel) Find(id string) (category model.Category, err error){
	err = categoryModel.Db.C(categoryModel.Collection).FindId(bson.ObjectIdHex(id)).One(&category)
	return
}
func (categoryModel CategoryModel) Create(category *model.Category) error {
	err := categoryModel.Db.C(categoryModel.Collection).Insert(&category)
	return err
}

func (categoryModel CategoryModel) Update(category *model.Category) error {
	err := categoryModel.Db.C(categoryModel.Collection).UpdateId(category.Id, &category)
	return err
}
func (categoryModel CategoryModel) Delete(category model.Category) error {
	err := categoryModel.Db.C(categoryModel.Collection).Remove(category)
	return err
}