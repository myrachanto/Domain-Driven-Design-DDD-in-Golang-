package consumer

import (
	"fmt"
	"testing"

	"github.com/myrachanto/ddd/model"
	"github.com/stretchr/testify/assert"
)

func TestValidityotheurl(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	consumer, err := Consumer.Retrivedata(url)	
	fmt.Println(">>>>>>>>>>>>>>>>>>>>", err)
	// assert.Nil(t, err, "the test passed even with the wrong address there is a response")
	assert.NotNil(t, consumer, "the responsecd results must not be equal to nil")

}
func TestValidityReading(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	data, _ := Consumer.Retrivedata(url)
	consumer, err := Consumer.Readdata(data.Body)
	assert.Nil(t, err, "the error reading is not nil")
	assert.NotNil(t, consumer, "the consumer reading err is not nil")

}
func TestValidityofjsonBody(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	data, _ := Consumer.Retrivedata(url)
	consumer, _ := Consumer.Readdata(data.Body)
	js, err := Consumer.Marshajson(consumer, &model.Prods{})
	assert.Nil(t, err, "the error marshalling is not nil")
	assert.NotNil(t, js, "the consumer marshalling err is not nil")

}
func TestTocsv(t *testing.T){
	file := "test.cv"
	data :=[]model.Product{}
	err := Consumer.Tocsv(data, file)
	assert.Nil(t, err, "could not create the csv file")

}
func TestConvertboolstring(t *testing.T){
str := Consumer.Convertboolstring(false)
assert.Equal(t, str, "false")
}
func TestConverstservices(t *testing.T){

}

func TestConverstimgs(t *testing.T){

}