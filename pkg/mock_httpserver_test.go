package openapiart_test

import (
	"io/ioutil"
	"log"
	"net/http"

	art "github.com/open-traffic-generator/openapiart/pkg"
)

type HttpServer struct {
	Api    art.OpenapiartApi
	Config art.PrefixConfig
}

var (
	httpServer HttpServer = HttpServer{}
)

func StartMockHttpServer() {
	httpServer.Api = art.NewApi()
	httpServer.Config = httpServer.Api.NewPrefixConfig()

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			httpServer.Config.FromJson(string(body))
			response := httpServer.Api.NewSetConfigResponse_StatusCode200()
			response.SetBytes([]byte("Successful set config operation"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		case http.MethodPatch:
			body, _ := ioutil.ReadAll(r.Body)
			request := httpServer.Api.NewPrefixConfig()
			request.FromJson(string(body))
			response := httpServer.Api.NewUpdateConfigResponse_StatusCode200()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		case http.MethodGet:
			config := httpServer.Api.NewPrefixConfig()
			response := httpServer.Api.NewGetConfigResponse_StatusCode200()
			response.PrefixConfig().FromJson(config.ToJson())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		}
	})

	log.Fatal(http.ListenAndServe(":50051", nil))
}

func init() {
	go StartMockHttpServer()
}
