package test

import (
	"bytes"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetRootResponse(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/apitest", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := io.ReadAll(wr.Body)
	r := openapiart.NewCommonResponseSuccess()
	err := r.FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "from GetRootResponse", r.Message())
}

func TestPostRootResponse(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodPost, "/api/apitest", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusBadRequest, wr.Code)

	inputbody := openapiart.NewApiTestInputBody().SetSomeString("this is the input body")
	j, _ := inputbody.ToJson()
	inputbuffer := bytes.NewBuffer([]byte(j))

	req, _ = http.NewRequest(http.MethodPost, "/api/apitest", inputbuffer)
	wr = httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := io.ReadAll(wr.Body)
	r := openapiart.NewCommonResponseSuccess()
	err := r.FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "this is the input body", r.Message())
}

func TestDummyResponseTest(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodDelete, "/api/apitest", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	response, _ := io.ReadAll(wr.Body)
	assert.Equal(t, "text/plain; charset=UTF-8", wr.Header().Get("Content-Type"))
	assert.Equal(t, "\"this is a string response\"", string(response))

}
