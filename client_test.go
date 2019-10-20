package openpay

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper func to create mock server
func testServer(res interface{}, code int) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(res)
	}))
	return ts
}

func testClient(res interface{}, code int) *Client {
	c, _ := NewClient("valid_key", "valid_id", nil)
	c.setAPIHost(testServer(res, code).URL + "/")
	return c
}

func TestNewClient(t *testing.T) {
	type args struct {
		key        string
		merchantID string
		options    *Options
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantHost string
	}{
		{
			name: "OK",
			args: args{
				key:        "test_key",
				merchantID: "test_id",
				options:    nil,
			},
			wantErr:  false,
			wantHost: testAPI,
		},
		{
			name: "OK prod",
			args: args{
				key:        "test_key",
				merchantID: "test_id",
				options: func() *Options {
					opts := defaultOptions()
					opts.UseProduction = true
					return opts
				}(),
			},
			wantErr:  false,
			wantHost: liveAPI,
		},
		{
			name: "Empty key",
			args: args{
				key:        "",
				merchantID: "test_id",
				options:    nil,
			},
			wantErr:  true,
			wantHost: testAPI,
		},
		{
			name: "Empty id",
			args: args{
				key:        "test_key",
				merchantID: "",
				options:    nil,
			},
			wantErr:  true,
			wantHost: testAPI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.key, tt.args.merchantID, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr { // validate host only if client is not null (no error expected)
				if got.apiHost != tt.wantHost {
					t.Errorf("NewClient().apiHost = %v, wantHost %v", got.apiHost, tt.wantHost)
					return
				}
			}
		})
	}
}
