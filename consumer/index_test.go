package consumer

import (
	"fmt"
	"testing"

	"github.com/myrachanto/ddd/model"
	"github.com/stretchr/testify/assert"
)

func TestValidityotheurl(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	consumer, err := Consumer.retrivedata(url)	
	fmt.Println(">>>>>>>>>>>>>>>",err.Message)
	assert.Equal(t, err.Message, "not found")
	assert.Equal(t, consumer, nil)

}
func TestValidityReading(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	data, _ := Consumer.retrivedata(url)
	consumer, err := Consumer.readdata(data.Body)
	assert.Equal(t, err.Message, "Something went wrong with reading responsce body")
	assert.Equal(t, consumer, nil)

}
func TestValidityofjsonBody(t *testing.T){
	url := "https://app.basmart.co.ke/front/product/newarrivalss"
	data, _ := Consumer.retrivedata(url)
	consumer, _ := Consumer.readdata(data.Body)
	js, err := Consumer.marshajson(consumer, &model.Prods{})
	assert.Equal(t, err.Message, "Something went wrong UNmasharling!!!")
	assert.Equal(t, js, nil)

}
func TestTocsv(t *testing.T){
	file := "test.cv"
	data :=[]model.Product{}
	err := Consumer.Tocsv(data, file)
	assert.Equal(t, err.Message, "error creating csv file")

}
func TestConvertboolstring(t *testing.T){
str := Consumer.Convertboolstring(false)
assert.Equal(t, str, "false")
}
func TestConverstservices(t *testing.T){

}

func TestConverstimgs(t *testing.T){

}