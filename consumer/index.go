package consumer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/myrachanto/ddd/httperors"
	"github.com/myrachanto/ddd/model"
	log "github.com/sirupsen/logrus"
)
var Consumer consumerinterface = &consumer{}

type consumerinterface interface{
	GetData(string, *model.Prods, string) ([]model.Product, *httperors.HttpError)
	retrivedata(url string)(*http.Response,*httperors.HttpError)
	readdata(f io.Reader)([]byte,*httperors.HttpError)
	marshajson(d []byte, pd *model.Prods,)(*model.Prods,*httperors.HttpError)
	Tocsv(data []model.Product, file string)*httperors.HttpError
	Convertboolstring(sd bool) string
	Converstservices(serv []model.Service) string
	Converstimgs(pics []model.Picture) string
}
type consumer struct {

}
func (c consumer)GetData(Url string, pd *model.Prods, file string) ([]model.Product, *httperors.HttpError){
	resp,e := c.retrivedata(Url)
	if e != nil {
		return nil, e
	}	
	r,e := c.readdata(resp.Body)
	if e != nil {
		return nil, e
	}
	js,e := c.marshajson(r, pd)
	if e != nil {
		return nil, e
	}
	c.Tocsv(js.Product, file)
return pd.Product, nil
}
func (c consumer)retrivedata(url string)(*http.Response,*httperors.HttpError){
	resp, err := http.Get(url)
	log.Info("resp retrived")
	if err != nil {
		log.Warn("resp not found")
		return nil, httperors.NewNotFoundError("not found")
	}
	return resp, nil
}
func (c consumer)readdata(f io.Reader)([]byte,*httperors.HttpError){
	respo, err := ioutil.ReadAll(f)
	log.Info("respo read from resp.body")
	if err != nil {
		log.Warn("respo failed read from resp.body")
		return nil, httperors.NewBadRequestError("Something went wrong with reading responsce body")
	}
	return respo, nil
}
func (c consumer)marshajson(d []byte, pd *model.Prods,)(*model.Prods,*httperors.HttpError){
	err1 := json.Unmarshal([]byte(d), &pd)
	if err1 != nil {
		log.Warn("respo failed json marshal")
		return nil, httperors.NewBadRequestError("Something went wrong UNmasharling!!!")
	}
	return pd, nil
}
func  (c consumer)Tocsv(data []model.Product, file string)*httperors.HttpError{
	if _, err := os.Stat(file); os.IsNotExist(err) {
		csvfile, err := os.Open(file)
	if err != nil{
		log.Warn("resp not found csv")
		httperors.NewNotFoundError("not found")
	}
	defer csvfile.Close()
	csvfilewriter := csv.NewWriter(csvfile)
	 defer csvfilewriter.Flush()
	 headers := []string {
		 "code", "Name", "url","description","meta","title","newarrivals","featured", "hotdeals","services", "images",
	 }
	 csvfilewriter.Write(headers)
	 for _, v := range data {
		 r := make([]string , 0, 1+len(headers))
		 r = append(r, v.Code, v.Name, v.Url, v.Description, v.Meta, v.Title, v.Newarrivals, c.Convertboolstring(v.Featured), c.Convertboolstring(v.Hotdeals), c.Converstservices(v.Services), c.Converstimgs(v.Images),)
		 csvfilewriter.Write(r)
	 }
	 return nil
	  }
	csvfile, err := os.Create("test.csv")
	if err != nil{
		httperors.NewBadRequestError("Something went wrong with creating that file")
	}
	defer csvfile.Close()
	csvfilewriter := csv.NewWriter(csvfile)
	 defer csvfilewriter.Flush()
	 headers := []string {
		 "code", "Name", "url","description","meta","title","newarrivals","featured", "hotdeals","services", "images",
	 }
	 csvfilewriter.Write(headers)
	 for _, v := range data {
		 r := make([]string , 0, 1+len(headers))
		 r = append(r, v.Code, v.Name, v.Url, v.Description, v.Meta, v.Title, v.Newarrivals, c.Convertboolstring(v.Featured), c.Convertboolstring(v.Hotdeals), c.Converstservices(v.Services), c.Converstimgs(v.Images),)
		 csvfilewriter.Write(r)
	 }
	 return nil
}
func  (c consumer)Convertboolstring(sd bool) string{
	if sd == true {
		return "true"
	}
	return "false"
 }
 func (c consumer)Converstservices(serv []model.Service) string{
	 res := ""
	 for _, v := range serv{
		res += v.Name +" - " + v.Price +"|"
	 }
	 return res
 } 
 func (c consumer)Converstimgs(pics []model.Picture) string{
	res := ""
	for _, v := range pics{
	   res += v.Name +"|"
	}
	return res
}

func  (c consumer)GetData1() *httperors.HttpError{
	fmt.Println("fffffffffffffffffffffffffffdddddddddddddd")
// resp := rest.Get("https://app.basmart.co.ke/front/home")
// 	fmt.Println(">>>>>>>>>>>", resp)

// api := gopencils.Api("https://app.basmart.co.ke")
// // Users Resource
// products := api.Res("/front/home")

// 	fmt.Println(">>>>>>>>>>>", products)

resp, err := http.Get("https://app.basmart.co.ke/front/product/newarrivals")
if err != nil {
	return httperors.NewNotFoundError("not found")
}
respo, err := ioutil.ReadAll(resp.Body)
if err != nil {
	return httperors.NewBadRequestError("Something went wrong")
}
// fmt.Println(">>>>>>>>>>>>",string(respo))
	// var prod1 []map[string]interface{}
	var pd model.Prods
	err = json.Unmarshal([]byte(respo), &pd)
	if err != nil {
		fmt.Println("ggggggggggggg",err)
		return httperors.NewBadRequestError("Something went wrong UNmasharling!!!")
	}
		fmt.Println("fffffffffffffff", pd)
	// ps := model.Product{}
	// pss := []model.Product{}
	// ss := model.Service{}
	// sss := []model.Service{}
	// is := model.Picture{}
	// iss := []model.Picture{}
	// for _, val := range prod1 {
	// 	ps.Code  = fmt.Sprintf("%s", val["code"])
	// 	ps.Name  = fmt.Sprintf("%s", val["name"])
	// 	ps.Url  = fmt.Sprintf("%s", val["url"])
	// 	ps.Description  = fmt.Sprintf("%s", val["description"])
	// 	ps.Meta  = fmt.Sprintf("%s", val["meta"])
	// 	ps.Title  = fmt.Sprintf("%s", val["title"])
	// 	ps.Newarrivals  = fmt.Sprintf("%s", val["newarrivals"])
	// 	ps.Featured  = val["featured"].(bool)
	// 	ps.Hotdeals  = val["hotdeals"].(bool)
	// 	serc := fmt.Sprintf("%s", val["services"])
	// 	var servs []map[string]interface{}
	// 	err = json.Unmarshal([]byte(serc), &servs)
	// 	if err != nil {
	// 		return httperors.NewBadRequestError("Something went wrong UNmasharling!!!")
	// 	}
	// 	for _, d := range servs {
	// 		ss.Name  = fmt.Sprintf("%s", d["name"])
	// 		ss.Price  = fmt.Sprintf("%s", d["price"])
	// 		sss = append(sss, ss)
	// 	}
	// 	ps.Services = sss
	// 	imgs := fmt.Sprintf("%s", val["images"])
	// 	var ims []map[string]interface{}
	// 	err = json.Unmarshal([]byte(imgs), &ims)
	// 	if err != nil {
	// 		return httperors.NewBadRequestError("Something went wrong UNmasharling!!!")
	// 	}
	// 	for _, d := range servs {
	// 		is.Name  = fmt.Sprintf("%s", d["name"])
	// 		iss = append(iss, is)
	// 	}
	// 	ps.Services = sss
	// 	ps.Images = iss
	// 	ps.Title  = fmt.Sprintf("%s", val["title"])
	// 	// ps.Services  = fmt.Sprintf("%s", val["services"])
	// 	pss = append(pss, ps)
	// }
	// // fmt.Println(">>>>>", pss)
	// fmt.Println("//////////////////////////////ssdddd/")
	// tocsv(pss)
return nil
}