/* Mock HTTP Server

Response Note:
- all returned responses must be <rpcmethod>Response.StatusCode<code>()
- this differs from the grpc server that expects the <rpcmethod>Response
  that is the composite of all status codes for the rpc method response
*/
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
			response := httpServer.Api.NewSetConfigResponse()
			switch httpServer.Config.Response() {
			case PrefixConfigResponse.STATUS_200:
				response.SetStatusCode200([]byte("Successful set config operation"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(response.StatusCode200()))
			case PrefixConfigResponse.STATUS_400:
				response.StatusCode400().SetErrors([]string{"A 400 error has occurred"})
				w.WriteHeader(http.StatusBadRequest)
				resp400, _ := response.StatusCode400().ToJson()
				w.Write([]byte(resp400))
			case PrefixConfigResponse.STATUS_500:
				response.StatusCode500().SetErrors([]string{"A 500 error has occurred"})
				w.WriteHeader(http.StatusInternalServerError)
				resp500, _ := response.StatusCode500().ToJson()
				w.Write([]byte(resp500))
			case PrefixConfigResponse.STATUS_404:
				response.StatusCode404().SetErrors([]string{"Not found error"})
				w.WriteHeader(http.StatusNotFound)
				resp404, _ := response.StatusCode404().ToJson()
				w.Write([]byte(resp404))
			}
		case http.MethodPatch:
			body, _ := ioutil.ReadAll(r.Body)
			request := httpServer.Api.NewPrefixConfig()
			request.FromJson(string(body))
			response := httpServer.Api.NewUpdateConfigResponse()
			conf, _ := httpServer.Config.ToJson()
			response.StatusCode200().FromJson(conf)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp200, _ := response.StatusCode200().ToJson()
			w.Write([]byte(resp200))
		case http.MethodGet:
			response := httpServer.Api.NewGetConfigResponse()
			conf, _ := httpServer.Config.ToJson()
			response.StatusCode200().FromJson(conf)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp200, _ := response.StatusCode200().ToJson()
			w.Write([]byte(resp200))
		}
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			response := httpServer.Api.NewGetMetricsResponse()
			response.StatusCode200().Ports().Add().SetName("p1").SetTxFrames(2000).SetRxFrames(1777)
			response.StatusCode200().Ports().Add().SetName("p2").SetTxFrames(3000).SetRxFrames(2999)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp200, _ := response.StatusCode200().ToJson()
			w.Write([]byte(resp200))
		}
	})

	http.HandleFunc("/warnings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			response := httpServer.Api.NewGetWarningsResponse()
			response.StatusCode200().SetWarnings([]string{"This is your first warning", "Your last warning"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp200, _ := response.StatusCode200().ToJson()
			w.Write([]byte(resp200))
		}
	})

	go func() {
		if err := http.ListenAndServe(httpServer.serverLocation, nil); err != nil {
			log.Fatal("Server failed to serve incoming HTTP request.")
		}
	}()
}
