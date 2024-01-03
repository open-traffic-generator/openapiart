package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetAllItems(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/serviceb", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := io.ReadAll(wr.Body)
	r := goapi.NewServiceAbcItemList()
	err := r.Unmarshal().FromJson(string(jsonResponse))
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

	jsonResponse, _ := io.ReadAll(wr.Body)
	r := goapi.NewServiceAbcItem()
	err := r.Unmarshal().FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "1", r.SomeId())

	req, _ = http.NewRequest(http.MethodGet, "/api/serviceb/3", nil)
	wr = httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusInternalServerError, wr.Code) // missing support for 404

	jsonResponse, _ = io.ReadAll(wr.Body)
	errNew := goapi.NewError()
	errFromJson := errNew.Unmarshal().FromJson(string(jsonResponse))
	assert.Nil(t, errFromJson)
	assert.Equal(t, "not found: id '3'", errNew.Errors()[0])
}

func TestGetSingleItemLevel2(t *testing.T) {
	router := setup()
	req, _ := http.NewRequest(http.MethodGet, "/api/serviceb/aa/bb", nil)
	wr := httptest.NewRecorder()
	router.ServeHTTP(wr, req)
	assert.Equal(t, http.StatusOK, wr.Code)

	jsonResponse, _ := io.ReadAll(wr.Body)
	r := goapi.NewServiceAbcItem()
	err := r.Unmarshal().FromJson(string(jsonResponse))
	assert.Nil(t, err)
	assert.Equal(t, "aa", r.PathId())
	assert.Equal(t, "bb", r.Level2())
}
