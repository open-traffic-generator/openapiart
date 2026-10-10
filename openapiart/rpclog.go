// RPC request/response logging controlled by the __OPENAPIART_RPC_LOG_ENV__
// environment variable, or OPENAPIART_RPC_LOG as a generic fallback. Set either
// variable to one of:
//
//	"stdout" or "-"  - write to stdout
//	"stderr"         - write to stderr
//	<path>           - append to the named file (created if absent)
//
// Each exchange is written as a YAML document separated by "---".

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ghodss/yaml"
	"google.golang.org/grpc/status"
)

type rpcLogger struct {
	mu   sync.Mutex
	file *os.File
}

const rpcLogEnvVar = "__OPENAPIART_RPC_LOG_ENV__"
const rpcLogFallbackEnvVar = "OPENAPIART_RPC_LOG"

var (
	rpcLogInst *rpcLogger
	rpcLogOnce sync.Once
)

func getRpcLog() *rpcLogger {
	rpcLogOnce.Do(func() {
		rpcLogInst = new(rpcLogger)
		envName, dest := rpcLogDestination()
		if dest == "" {
			return
		}
		switch dest {
		case "-", "stdout":
			rpcLogInst.file = os.Stdout
		case "stderr":
			rpcLogInst.file = os.Stderr
		default:
			f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: failed to open %q: %v\n", envName, dest, err)
			} else {
				rpcLogInst.file = f
			}
		}
	})
	return rpcLogInst
}

func rpcLogDestination() (string, string) {
	for _, envName := range []string{rpcLogEnvVar, rpcLogFallbackEnvVar} {
		if dest := os.Getenv(envName); dest != "" {
			return envName, dest
		}
	}
	return rpcLogEnvVar, ""
}

func (r *rpcLogger) enabled() bool {
	return r.file != nil
}

func jsonToAny(j string) interface{} {
	var v interface{}
	if err := json.Unmarshal([]byte(j), &v); err != nil {
		return j
	}
	return v
}

func (r *rpcLogger) writeEntry(entry map[string]interface{}) {
	entry["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	out, err := yaml.Marshal(entry)
	if err != nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Fprintf(r.file, "---\n%s", string(out))
}

func (r *rpcLogger) logHTTP(httpMethod, urlPath, requestJSON string, statusCode int, responseBody []byte, responseErr error) {
	if !r.enabled() {
		return
	}
	entry := map[string]interface{}{
		"transport": "http",
		"method":    fmt.Sprintf("%s /%s", httpMethod, urlPath),
	}
	if requestJSON != "" {
		entry["request"] = jsonToAny(requestJSON)
	}
	respEntry := map[string]interface{}{"status": statusCode}
	if len(responseBody) > 0 {
		var bodyObj interface{}
		if err := json.Unmarshal(responseBody, &bodyObj); err == nil {
			respEntry["body"] = bodyObj
		} else {
			respEntry["body"] = map[string]interface{}{"<bytes>": len(responseBody)}
		}
	}
	if responseErr != nil {
		respEntry["error"] = map[string]interface{}{"error": responseErr.Error()}
	}
	entry["response"] = respEntry
	r.writeEntry(entry)
}

func (r *rpcLogger) logGRPC(method, requestJSON, responseJSON string) {
	if !r.enabled() {
		return
	}
	entry := map[string]interface{}{
		"transport": "grpc",
		"method":    method,
	}
	if requestJSON != "" {
		entry["request"] = jsonToAny(requestJSON)
	}
	if responseJSON != "" {
		entry["response"] = jsonToAny(responseJSON)
	}
	r.writeEntry(entry)
}

func (r *rpcLogger) logGRPCError(method, requestJSON string, rpcErr error) {
	if !r.enabled() || rpcErr == nil {
		return
	}
	errorEntry := map[string]interface{}{"error": rpcErr.Error()}
	if st, ok := status.FromError(rpcErr); ok {
		errorEntry = map[string]interface{}{
			"code":  int32(st.Code()),
			"error": jsonToAny(st.Message()),
		}
	}
	responseJSON, _ := json.Marshal(map[string]interface{}{"error": errorEntry})
	r.logGRPC(method, requestJSON, string(responseJSON))
}
