package openapiart_test

import (
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
)

func TestLogging(t *testing.T) {
	api := openapiart.NewApi()
	openapiart.Logger.Info().Msg("Start configuring test")
	api.NewConfig()
}
