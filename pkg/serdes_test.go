package openapiart_test

import (
	"fmt"
	"testing"

	. "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

var yaml_config = `a: asdf
b: 22.2
c: 33
`

func TestFromYaml(t *testing.T) {
	api := NewApi()
	c := api.NewPrefixConfig()
	err := c.FromYaml(yaml_config)
	if err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, c.A(), "asdf")
		assert.Equal(t, c.B(), float32(22.2))
		assert.Equal(t, c.C(), int32(33))
		fmt.Println(c.ToYaml())
	}
}

func TestFromJson(t *testing.T) {
	var json_config = `{"a": "asdf", "b": 22.2,	"c": 33}`
	api := NewApi()
	c := api.NewPrefixConfig()
	err := c.FromJson(json_config)
	if err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, c.A(), "asdf")
		assert.Equal(t, c.B(), float32(22.2))
		assert.Equal(t, c.C(), int32(33))
		fmt.Println(c.ToYaml())
	}
}
