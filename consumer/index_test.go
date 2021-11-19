package consumer

import (
	// "fmt"
	"testing"

	"github.com/myrachanto/ddd/model"
	"github.com/stretchr/testify/assert"
)
var tests = []struct{
	url string
	results interface{}
	code int
	errors interface{}
	isErr bool
}{
	{url:"https://app.basmart.co.ke/front/product/newarrivals",code:200,errors: nil, isErr:true},
	{url:"https://app.basmart.co.ke/front/product/newarrivalss",code:400,errors: nil, isErr:true},
}

func TestValidityotheurl(t *testing.T){
	for _, tt := range tests {
		got,err := Consumer.Retrivedata(tt.url)
		if tt.isErr{
			if err == nil{
				assert.Nil(t, err, "the test passed even with the wrong address there is a response")
			}
		}else{
			if err != nil {
				t.Error("did not expect an error but got one")
			}
		}
		if got.StatusCode == tt.code{
			assert.EqualValues(t, got.StatusCode, tt.code, "success")
		}
		if got.Body != nil{
			assert.NotNil(t, got.Body, "the responsecd results must not be equal to nil")
	   }
	}
	// url := "https://app.basmart.co.ke/front/product/newarrivalss"
	// consumer, err := Consumer.Retrivedata(url)	
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>", consumer.StatusCode)
	// assert.EqualValues(t, consumer.StatusCode, 400, "there was an error because the url is wrong")
	// assert.Nil(t, err, "the test passed even with the wrong address there is a response")
	// assert.NotNil(t, consumer, "the responsecd results must not be equal to nil")

}
var res = []byte{}
var tests2 = []struct{
	url string
	results []byte
	errors string
	isErr bool
}{
	{url:"https://app.basmart.co.ke/front/product/newarrivalss",results: nil ,errors: "Something went wrong with reading responsce body", isErr:true},
	{url:"https://app.basmart.co.ke/front/product/newarrivals",results: res ,errors: "", isErr:true},
}
func TestValidityReading(t *testing.T){
	for _, tt := range tests2 {
		data,_ := Consumer.Retrivedata(tt.url)
	  got, err := Consumer.Readdata(data.Body)
		if tt.isErr{
			if err == nil{
				assert.Nil(t, err, "the error reading is not nil")
			}
		}else{
			if err != nil {
				assert.EqualValues(t, err, tt.errors, "success")
			}
		}
		if got == nil{
			assert.EqualValues(t, got, tt.results, "success")
		}
	}
	// url := "https://app.basmart.co.ke/front/product/newarrivalss"
	// data, _ := Consumer.Retrivedata(url)
	// consumer, err := Consumer.Readdata(data.Body)
	// assert.Nil(t, err, "the error reading is not nil")
	// assert.NotNil(t, consumer, "the consumer reading err is not nil")

}
var prods = []model.Prods{}
var tests3 = []struct{
	url string
	results []model.Prods
	errors string
	isErr bool
}{
	{url:"https://app.basmart.co.ke/front/product/newarrivalss",results: nil ,errors: "Something went wrong Unmasharling!!!", isErr:true},
	{url:"https://app.basmart.co.ke/front/product/newarrivals",results: prods ,errors: "", isErr:true},
}
func TestValidityofjsonBody(t *testing.T){
	for _, tt := range tests3 {
		data, _ := Consumer.Retrivedata(tt.url)
		consumer, _ := Consumer.Readdata(data.Body)
		got, err := Consumer.Marshajson(consumer, &model.Prods{})
			if tt.isErr{
			if err == nil{
				assert.Nil(t, err, "the error reading is not nil")
			}
		}else{
			if err != nil {
				assert.EqualValues(t, err, tt.errors, "success")
			}
		}
		if got == nil{
			assert.EqualValues(t, got, tt.results, "success")
		}
	}
	// url := "https://app.basmart.co.ke/front/product/newarrivalss"
	// data, _ := Consumer.Retrivedata(url)
	// consumer, _ := Consumer.Readdata(data.Body)
	// js, err := Consumer.Marshajson(consumer, &model.Prods{})
	// assert.Nil(t, err, "the error marshalling is not nil")
	// assert.NotNil(t, js, "the consumer marshalling err is not nil")

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