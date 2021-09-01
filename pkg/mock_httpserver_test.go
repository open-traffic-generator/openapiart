package openapiart_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/open-traffic-generator/openapiart/pkg"
)

type HttpServer struct {
	serverLocation string
	Location       string
	Api            OpenapiartApi
	Config         PrefixConfig
}

var (
	httpServer HttpServer = HttpServer{
		serverLocation: "127.0.0.1:50051",
	}
)

func StartMockHttpServer() {
	httpServer.Location = fmt.Sprintf("http://%s", httpServer.serverLocation)
	httpServer.Api = NewApi()
	httpServer.Config = httpServer.Api.NewPrefixConfig()

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			httpServer.Config.FromJson(string(body))
			w.Header().Set("Content-Type", "application/json")
			switch httpServer.Config.Response() {
			case PrefixConfigResponse.STATUS_200:
				response := httpServer.Api.NewSetConfigResponse_StatusCode200()
				response.SetBytes([]byte("Successful set config operation"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(response.ToJson()))
			case PrefixConfigResponse.STATUS_400:
				response := httpServer.Api.NewSetConfigResponse_StatusCode400()
				response.ErrorDetails().SetErrors([]string{"A 400 error has occurred"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(response.ToJson()))
			case PrefixConfigResponse.STATUS_500:
				response := httpServer.Api.NewSetConfigResponse_StatusCode500()
				response.Error().SetErrors([]string{"A 500 error has occurred"})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(response.ToJson()))
			}
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

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			response := httpServer.Api.NewGetMetricsResponse_StatusCode200()
			response.Metrics().Ports().Add().SetName("p1").SetTxFrames(2000).SetRxFrames(1777)
			response.Metrics().Ports().Add().SetName("p2").SetTxFrames(3000).SetRxFrames(2999)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		}
	})

	http.HandleFunc("/warnings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			response := httpServer.Api.NewGetWarningsResponse_StatusCode200()
			response.WarningDetails().SetWarnings([]string{"Warning number 1", "Your last warning"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		}
	})

	go func() {
		if err := http.ListenAndServe(httpServer.serverLocation, nil); err != nil {
			log.Fatal("Server failed to serve incoming HTTP request.")
		}
	}()
}
