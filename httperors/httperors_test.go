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
func TestNewSuccessMessage(t *testing.T){
	err := NewSuccessMessage("Deletion was a success")
	assert.NotNil(t,err)
	assert.EqualValues(t, http.StatusOK, err.Code)
	assert.EqualValues(t, "Deletion was a success", err.Message)
	assert.EqualValues(t, "Delete success", err.Error)

}