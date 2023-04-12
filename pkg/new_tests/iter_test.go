package openapiart_test

import (
	"fmt"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

var strlenValues = []string{"200", "300"}
var integer641Values = []int64{2132433546, 3892433546}
var integer62Values = []int64{5645336, 989645336}

func TestIterAdd(t *testing.T) {

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().IterObject().Add().SetStrLen("200").SetInteger641(2132433546).SetInteger642(5645336)
	config.NativeFeatures().IterObject().Add().SetStrLen("300").SetInteger641(3892433546).SetInteger642(989645336)

	assert.Equal(t, len(config.NativeFeatures().IterObject().Items()), 2)
	for idx, iterObj := range config.NativeFeatures().IterObject().Items() {
		assert.Equal(t, strlenValues[idx], iterObj.StrLen())
		assert.Equal(t, integer641Values[idx], iterObj.Integer641())
		assert.Equal(t, integer62Values[idx], iterObj.Integer642())
	}
}

func TestAppend(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().IterObject().Add().SetStrLen("200").SetInteger641(2132433546).SetInteger642(5645336)
	itr := config.NativeFeatures().IterObject().Append(openapiart.NewMixedObject().SetStrLen("300").SetInteger641(3892433546).SetInteger642(989645336))

	assert.Equal(t, len(itr.Items()), 2)
	for idx, iterObj := range config.NativeFeatures().IterObject().Items() {
		assert.Equal(t, strlenValues[idx], iterObj.StrLen())
		assert.Equal(t, integer641Values[idx], iterObj.Integer641())
		assert.Equal(t, integer62Values[idx], iterObj.Integer642())
	}
}

func TestClear(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().IterObject().Add().SetStrLen("200").SetInteger641(2132433546).SetInteger642(5645336)
	config.NativeFeatures().IterObject().Add().SetStrLen("300").SetInteger641(3892433546).SetInteger642(989645336)

	assert.Equal(t, len(config.NativeFeatures().IterObject().Items()), 2)
	config.NativeFeatures().IterObject().Clear()
	assert.Equal(t, len(config.NativeFeatures().IterObject().Items()), 0)
}

func TestSet(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			errValue := "runtime error: index out of range [3] with length 2"
			assert.Equal(t, errValue, fmt.Sprintf("%v", err))
		}
	}()
	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().IterObject().Add().SetStrLen("200").SetInteger641(2132433546).SetInteger642(5645336)
	config.NativeFeatures().IterObject().Add()
	itr := config.NativeFeatures().IterObject().Set(1, openapiart.NewMixedObject().SetStrLen("300").SetInteger641(3892433546).SetInteger642(989645336))

	assert.Equal(t, strlenValues[1], itr.Items()[1].StrLen())
	assert.Equal(t, len(itr.Items()), 2)

	config.NativeFeatures().IterObject().Set(3, openapiart.NewMixedObject().SetStrLen("400").SetInteger641(4789678546).SetInteger642(4567645336))
}
