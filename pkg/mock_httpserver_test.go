package openapiart_test

import (
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/open-traffic-generator/openapiart/pkg"
)

type ResponseWarning struct {
	Warnings []string
}

func StartMockHttpServer() {
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		api := NewApi()
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			request := api.NewPrefixConfig()
			request.FromJson(string(body))
			response := api.NewSetConfigResponse_StatusCode200()
			response.SetBytes([]byte("Successful set config operation"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		case http.MethodPatch:
			body, _ := ioutil.ReadAll(r.Body)
			request := api.NewPrefixConfig()
			request.FromJson(string(body))
			response := api.NewUpdateConfigResponse_StatusCode200()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		case http.MethodGet:
			body, _ := ioutil.ReadAll(r.Body)
			request := api.NewPrefixConfig()
			request.FromJson(string(body))
			response := api.NewGetConfigResponse_StatusCode200()
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
