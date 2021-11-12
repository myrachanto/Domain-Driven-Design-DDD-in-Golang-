package controllers

import (
	"fmt"
	// "log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

type PostData struct {
	Key string
	Value string
}
var theTest = []struct {
	name string
	url string
	method string
	params []PostData
	expectedStatusCode int
}{
	// {"categorys", "/categorys", "GET", []PostData{}, http.StatusOK},
	{"category", "/categorys/:id", "GET", []PostData{}, http.StatusOK},
}

func TestControllers(t *testing.T){
		// Setup
		e := echo.New()
		fmt.Println("step1")
		for _, f := range theTest {
			if f.method == "GET"{
				req := httptest.NewRequest(http.MethodGet, "/categorys", nil)
				res := httptest.NewRecorder()
				c := e.NewContext(req, res)
				c.SetPath(f.url)
				c.SetParamNames("id")
				c.SetParamValues("123456")
				fmt.Println("step2")
				err := CategoryController.GetOne(c)
				if err != nil {
					assert.NotEqual(t, http.StatusOK, res.Code, "querry passed")

				}
				// if assert.NoError(t, CategoryController.GetOne(c)) {
				// 	fmt.Println("step3")
				// 	assert.Equal(t, http.StatusOK, res.Code)
				// 	// assert.Equal(t, userJSON, rec.Body.String())
				// }
				fmt.Println("step4")
			}
			// }else {

			// }
		}

}