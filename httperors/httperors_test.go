package httperors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestNewBadRequestError(t *testing.T){
	err := NewBadRequestError("this is the message")
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "bad request", err.Error)
}
func TestNewNotFoundError(t *testing.T){
	err := NewNotFoundError("this is the message")
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusNotFound, err.Code)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "Not Found", err.Error)

}