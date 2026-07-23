package customerio

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type responseClientFunc func(*http.Request) (*http.Response, error)

func (f responseClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestDoHTTPResponseBodyLimit(t *testing.T) {
	tests := []struct {
		name    string
		size    int64
		wantErr error
	}{
		{name: "maximum accepted size", size: maxResponseBodySize},
		{name: "oversized response", size: maxResponseBodySize + 1, wantErr: ErrResponseTooLarge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := responseClientFunc(func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(strings.Repeat("a", int(tt.size)))),
				}, nil
			})

			body, statusCode, err := doHTTP(
				context.Background(),
				client,
				http.MethodGet,
				"https://example.com",
				"test",
				nil,
				func(*http.Request) {},
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("doHTTP() error = %v, want %v", err, tt.wantErr)
			}
			if statusCode != http.StatusOK {
				t.Fatalf("doHTTP() status = %d, want %d", statusCode, http.StatusOK)
			}
			if tt.wantErr == nil && int64(len(body)) != tt.size {
				t.Fatalf("doHTTP() body size = %d, want %d", len(body), tt.size)
			}
			if tt.wantErr != nil && body != nil {
				t.Fatalf("doHTTP() body = %d bytes, want nil", len(body))
			}
		})
	}
}
