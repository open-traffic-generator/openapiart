
import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func resetRpcLogForTest() {
	if rpcLogInst != nil && rpcLogInst.file != nil && rpcLogInst.file != os.Stdout && rpcLogInst.file != os.Stderr {
		rpcLogInst.file.Close()
	}
	rpcLogInst = nil
	rpcLogOnce = sync.Once{}
}

func TestCheckClientServerVersionCompatibility(t *testing.T) {
	type args struct {
		clientVer string
		serverVer string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: add TCs for build versions (e.g. 0.1.1-200)
		{name: "patch-low-high", args: args{clientVer: "0.2.1", serverVer: "0.2.*"}, wantErr: false},
		{name: "patch-high-low", args: args{clientVer: "0.2.5", serverVer: "0.2.*"}, wantErr: false},
		{name: "minor-low-high", args: args{clientVer: "0.2.5", serverVer: ">0.1.0 <0.3.0"}, wantErr: false},
		{name: "minor-high-low", args: args{clientVer: "0.3.5", serverVer: "^0.2.5"}, wantErr: true},
		{name: "major-low-high", args: args{clientVer: "0.2.5", serverVer: "^1.2.5"}, wantErr: true},
		{name: "major-high-low", args: args{clientVer: "1.2.5", serverVer: "^0.2.5"}, wantErr: true},
		{name: "invalid-valid", args: args{clientVer: "0.2.1.1", serverVer: "0.2.5"}, wantErr: true},
		{name: "valid-invalid", args: args{clientVer: "0.2.1", serverVer: "0.2.5.1"}, wantErr: true},
		{name: "invalid-invalid", args: args{clientVer: "0.2.1.1", serverVer: "0.2.5.1"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkClientServerVersionCompatibility(tt.args.clientVer, tt.args.serverVer, "API"); (err != nil) != tt.wantErr {
				t.Errorf("checkClientServerVersionCompatibility() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRpcLoggerHTTP(t *testing.T) {
	path := filepath.Join(t.TempDir(), "rpc.yaml")
	t.Setenv(rpcLogEnvVar, path)
	resetRpcLogForTest()
	defer resetRpcLogForTest()

	logger := getRpcLog()
	if !logger.enabled() {
		t.Fatal("expected rpc logger to be enabled")
	}
	logger.logHTTP("POST", "api/config", `{"a":"asdf"}`, 200, []byte(`{"ok":true}`), nil)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(content)
	for _, want := range []string{
		"transport: http",
		"method: POST /api/config",
		"a: asdf",
		"status: 200",
		"ok: true",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rpc log to contain %q, got:\n%s", want, got)
		}
	}
}

func TestRpcLoggerHTTPReadError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "rpc.yaml")
	t.Setenv(rpcLogEnvVar, path)
	resetRpcLogForTest()
	defer resetRpcLogForTest()

	getRpcLog().logHTTP("GET", "api/config", "", 200, []byte("partial"), errors.New("unexpected EOF"))

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(content)
	for _, want := range []string{
		"method: GET /api/config",
		"status: 200",
		"<bytes>",
		"7",
		"error: unexpected EOF",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rpc log to contain %q, got:\n%s", want, got)
		}
	}
}

func TestRpcLoggerGRPCError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "rpc.yaml")
	t.Setenv(rpcLogEnvVar, path)
	resetRpcLogForTest()
	defer resetRpcLogForTest()

	getRpcLog().logGRPCError("GetConfig", `{"a":"asdf"}`, status.Error(codes.Unavailable, "server unavailable"))

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(content)
	for _, want := range []string{
		"transport: grpc",
		"method: GetConfig",
		"a: asdf",
		"code: 14",
		"error: server unavailable",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rpc log to contain %q, got:\n%s", want, got)
		}
	}
}
