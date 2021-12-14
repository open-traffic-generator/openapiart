package test

import (
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetAllItems(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/serviceb", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := ioutil.ReadAll(wr.Body)
	r := openapiart.NewServiceBItemList()
	err := r.FromJson(string(jsonResponse))
	assert.Nil(t, err)
	items := r.Items().Items()
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "1", items[0].SomeId())
	assert.Equal(t, "2", items[1].SomeId())
}

func TestGetSingleItem(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/serviceb/1", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := ioutil.ReadAll(wr.Body)
	r := openapiart.NewServiceBItem()
	err := r.FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "1", r.SomeId())

	req, _ = http.NewRequest(http.MethodGet, "/api/serviceb/3", nil)
	wr = httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusBadRequest, wr.Code) // missing support for 404

	jsonResponse, _ = ioutil.ReadAll(wr.Body)
	errNew := openapiart.NewCommonResponseError()
	errFromJson := errNew.FromJson(string(jsonResponse))
	assert.Nil(t, errFromJson)
	assert.Equal(t, "not found: id '3'", errNew.Message())
}

func TestGetSingleItemLevel2(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/serviceb/aa/bb", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := ioutil.ReadAll(wr.Body)
	r := openapiart.NewServiceBItem()
	err := r.FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "aa", r.PathId())
	assert.Equal(t, "bb", r.Level2())
}
