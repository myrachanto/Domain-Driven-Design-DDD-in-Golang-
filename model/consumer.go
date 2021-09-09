package model
import (
	"github.com/myrachanto/ddd/httperors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Prods struct {
	Product []Product
}
type Product struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty"`
	Url           string             `json:"url,omitempty"`
	Title         string             `json:"title,omitempty"`
	Description   string             `json:"description,omitempty"`
	Meta          string             `json:"meta,omitempty"`
	Altertag      string             `json:"altertag,omitempty"`
	Footer        string             `json:"footer,omitempty"`
	Code          string             `json:"code,omitempty"`
	Majorcategory string             `json:"majorcat,omitempty"`
	Category      string             `json:"category,omitempty"`
	Newarrivals   string             `json:"newarrivals,omitempty"`
	Subcategory   string             `json:"subcategory,omitempty"`
	Oldprice      float64            `json:"oldprice,omitempty"`
	Newprice      float64            `json:"newprice,omitempty"`
	Buyprice      float64            `json:"buyprice,omitempty"`
	Picture       string             `json:"picture,omitempty"`
	Quantity      float64            `json:"quantity,omitempty"`
	Services      []Service          `json:"services,omitempty"`
	Colors        []Color            `json:"colors,omitempty"`
	Images        []Picture          `json:"images,omitempty"`
	Featured  bool     `json:"featured,omitempty"`
	Promotion bool     `json:"promotion,omitempty"`
	Hotdeals  bool     `json:"hotdeals,omitempty"`
	Base      `json:"base,omitempty"`
}
type Newarrivals struct {
	Product []*Product `json:"product,omitempty"`
}
type Service struct {
	Productcode string `json:"productcode,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       string `json:"price,omitempty"`
}
type Picture struct {
	Productcode string `json:"productcode,omitempty"`
	Name        string `json:"name,omitempty"`
}
type Color struct {
	Productcode string `json:"productcode,omitempty"`
	Name        string `json:"name,omitempty"`
}

func (product Product) Validate() *httperors.HttpError {
	if product.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" {
		return httperors.NewNotFoundError("Invalid title")
	}
	if product.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	return nil
}
