// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/kernel/hypeman-go"
	"github.com/kernel/hypeman-go/internal/testutil"
	"github.com/kernel/hypeman-go/option"
)

func TestPushNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hypeman.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pushes.New(context.TODO(), hypeman.PushNewParams{
		CreatePushRequest: hypeman.CreatePushRequestParam{
			Image:  "docker.io/library/alpine:latest",
			Target: "123456789.dkr.ecr.us-east-1.amazonaws.com/myapp:v1",
			Credentials: hypeman.PushCredentialsParam{
				Password:      hypeman.String("password"),
				RegistryToken: hypeman.String("registry_token"),
				Username:      hypeman.String("username"),
			},
			Insecure: hypeman.Bool(true),
		},
	})
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPushList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hypeman.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pushes.List(context.TODO())
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPushGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hypeman.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pushes.Get(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
