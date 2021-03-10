package rest_errors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)
import "errors"
import "github.com/stretchr/testify/assert"

func TestNewInternalServerError(t *testing.T) {
	err:=NewInternalServerError("testing internal server", errors.New("database error"))
	assert.NotNil(t,err)
	assert.NotNil(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "testing internal server", err.Message())
	assert.EqualValues(t, "Message: testing internal server - Status: 500 - Error: internal_server_error - Causes: [database error]", err.Error())
	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewUnauthorizedError(t *testing.T) {
	err:=NewUnauthorizedError("unauthorized access")
	assert.NotNil(t, err)
	assert.NotNil(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "unauthorized access", err.Message())
	assert.EqualValues(t, "Message: unauthorized access - Status: 401 - Error: unauthorized - Causes: []", err.Error())
	assert.Nil(t, err.Causes())
	assert.EqualValues(t, 0,len(err.Causes()))
	assert.EqualValues(t, []interface{}([]interface{}(nil)),err.Causes())
}

func TestNewNotFound(t *testing.T) {
	err:=NewNotFound("not found error")
	assert.NotNil(t, err)
	assert.NotNil(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "not found error", err.Message())
	assert.EqualValues(t, "Message: not found error - Status: 404 - Error: not_found - Causes: []", err.Error())
	assert.Nil(t, err.Causes())
	assert.EqualValues(t, 0,len(err.Causes()))
	assert.EqualValues(t, []interface{}([]interface{}(nil)),err.Causes())
}

func TestNewBadRequestError(t *testing.T) {
	err:=NewBadRequestError("bad request error")
	assert.NotNil(t, err)
	assert.NotNil(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "bad request error", err.Message())
	assert.EqualValues(t, "Message: bad request error - Status: 400 - Error: bad_request - Causes: []", err.Error())
	assert.Nil(t, err.Causes())
	assert.EqualValues(t, 0,len(err.Causes()))
	assert.EqualValues(t, []interface{}([]interface{}(nil)),err.Causes())
}

func TestNewRestErrorFromBytes(t *testing.T) {
	apiError := restErr{ErrMessage: "Test Message",ErrStatus: http.StatusInternalServerError,ErrError: "invalid_json",ErrCauses: []interface{}([]interface{}(nil))}
	errorBytesWithError := new(bytes.Buffer)
	json.NewEncoder(errorBytesWithError).Encode(apiError)
	_, err := NewRestErrorFromBytes(errorBytesWithError.Bytes())
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), "invalid json")
}

func TestNewRestError(t *testing.T) {
	rstError:=NewRestError("test rest error",500,"rest_error",[]interface{}([]interface{}(nil)))

	assert.NotNil(t, rstError)
	assert.EqualValues(t, http.StatusInternalServerError,rstError.Status())
	assert.EqualValues(t,"Message: test rest error - Status: 500 - Error: rest_error - Causes: []",rstError.Error())
	assert.Nil(t, rstError.Causes())
	assert.EqualValues(t, []interface{}([]interface{}(nil)),rstError.Causes())
}