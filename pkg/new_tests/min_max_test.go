package openapiart_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestBoundaryValSchema(t *testing.T) {

	// This test checks the values in boundary val schema.
	// Objective is validation should not a problem.

	config := goapi.NewTestConfig()
	rVal := config.NativeFeatures().BoundaryVal()
	rVal.SetIntVal(40)
	rVal.SetNumVal(5.67)
	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestBoundaryArrayValSchema(t *testing.T) {

	// This test checks the values in boundary val schema.
	// Objective is validation should not a problem.

	config := goapi.NewTestConfig()
	rVal := config.NativeFeatures().BoundaryValArray()
	rVal.SetIntVals([]int32{40, 50})
	rVal.SetNumVals([]float32{5.67, 6.78})
	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)

}

func TestErrorForMinCheck(t *testing.T) {

	// This test basically verifies if proper errors are raised by SDk for min and max values

	config := goapi.NewTestConfig()
	rVal := config.NativeFeatures().BoundaryVal()

	// check max errors
	rVal.SetIntVal(300)
	rVal.SetNumVal(500.678)
	_, err := config.Marshal().ToYaml()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "5 <= BoundaryVal.IntVal <= 100 but Got 300")
	assert.Contains(t, err.Error(), "5.0 <= BoundaryVal.NumVal <= 100.0 but Got 500.678009")

	//check min errors
	rVal.SetIntVal(2)
	rVal.SetNumVal(1.23)
	_, err = config.Marshal().ToYaml()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "5 <= BoundaryVal.IntVal <= 100 but Got 2")
	assert.Contains(t, err.Error(), "5.0 <= BoundaryVal.NumVal <= 100.0 but Got 1.23")

}