package hypeman_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/kernel/hypeman-go"
	"github.com/kernel/hypeman-go/option"
)

func TestImagePathParameterEscapesRepositorySlashes(t *testing.T) {
	const imageName = "docker.io/library/alpine:latest"
	const wantPath = "/images/docker.io%2Flibrary%2Falpine:latest"

	tests := []struct {
		name   string
		method string
		call   func(hypeman.Client) error
	}{
		{
			name:   "get",
			method: http.MethodGet,
			call: func(client hypeman.Client) error {
				_, err := client.Images.Get(context.Background(), imageName)
				return err
			},
		},
		{
			name:   "delete",
			method: http.MethodDelete,
			call: func(client hypeman.Client) error {
				return client.Images.Delete(context.Background(), imageName)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotPath string
			client := hypeman.NewClient(
				option.WithAPIKey("test-key"),
				option.WithHTTPClient(&http.Client{
					Transport: &closureTransport{
						fn: func(req *http.Request) (*http.Response, error) {
							if req.Method != tt.method {
								t.Errorf("method = %s, want %s", req.Method, tt.method)
							}
							gotPath = req.URL.EscapedPath()
							status := http.StatusNoContent
							if tt.method == http.MethodGet {
								status = http.StatusNotFound
							}
							return &http.Response{
								StatusCode: status,
								Header:     make(http.Header),
								Body:       io.NopCloser(strings.NewReader(`{"code":"not_found","message":"not found"}`)),
							}, nil
						},
					},
				}),
			)

			_ = tt.call(client)
			if gotPath != wantPath {
				t.Fatalf("escaped path = %q, want %q", gotPath, wantPath)
			}
		})
	}
}
