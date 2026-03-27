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
	"github.com/kernel/hypeman-go/shared"
)

func TestInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.New(context.TODO(), hypeman.InstanceNewParams{
		Image: "docker.io/library/alpine:latest",
		Name:  "my-workload-1",
		Cmd:   []string{"echo", "hello"},
		Credentials: map[string]hypeman.InstanceNewParamsCredential{
			"OUTBOUND_OPENAI_KEY": {
				Inject: []hypeman.InstanceNewParamsCredentialInject{{
					As: hypeman.InstanceNewParamsCredentialInjectAs{
						Format: "Bearer ${value}",
						Header: "Authorization",
					},
					Hosts: []string{"api.openai.com", "*.openai.com"},
				}},
				Source: hypeman.InstanceNewParamsCredentialSource{
					Env: "OUTBOUND_OPENAI_KEY",
				},
			},
		},
		Devices:    []string{"l4-gpu"},
		DiskIoBps:  hypeman.String("100MB/s"),
		Entrypoint: []string{"/bin/sh", "-c"},
		Env: map[string]string{
			"PORT":     "3000",
			"NODE_ENV": "production",
		},
		GPU: hypeman.InstanceNewParamsGPU{
			Profile: hypeman.String("L40S-1Q"),
		},
		HotplugSize: hypeman.String("2GB"),
		Hypervisor:  hypeman.InstanceNewParamsHypervisorCloudHypervisor,
		Network: hypeman.InstanceNewParamsNetwork{
			BandwidthDownload: hypeman.String("1Gbps"),
			BandwidthUpload:   hypeman.String("1Gbps"),
			Egress: hypeman.InstanceNewParamsNetworkEgress{
				Enabled: hypeman.Bool(true),
				Enforcement: hypeman.InstanceNewParamsNetworkEgressEnforcement{
					Mode: "all",
				},
			},
			Enabled: hypeman.Bool(true),
		},
		OverlaySize:       hypeman.String("20GB"),
		Size:              hypeman.String("2GB"),
		SkipGuestAgent:    hypeman.Bool(false),
		SkipKernelHeaders: hypeman.Bool(true),
		SnapshotPolicy: hypeman.SnapshotPolicyParam{
			Compression: shared.SnapshotCompressionConfigParam{
				Enabled:   true,
				Algorithm: shared.SnapshotCompressionConfigAlgorithmZstd,
				Level:     hypeman.Int(1),
			},
		},
		Tags: map[string]string{
			"team": "backend",
			"env":  "staging",
		},
		Vcpus: hypeman.Int(2),
		Volumes: []hypeman.VolumeMountParam{{
			MountPath:   "/mnt/data",
			VolumeID:    "vol-abc123",
			Overlay:     hypeman.Bool(true),
			OverlaySize: hypeman.String("1GB"),
			Readonly:    hypeman.Bool(true),
		}},
	})
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Update(
		context.TODO(),
		"id",
		hypeman.InstanceUpdateParams{
			Env: map[string]string{
				"OUTBOUND_OPENAI_KEY": "new-rotated-key-456",
			},
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.List(context.TODO(), hypeman.InstanceListParams{
		State: hypeman.InstanceListParamsStateCreated,
		Tags: map[string]string{
			"team": "backend",
			"env":  "staging",
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

func TestInstanceDelete(t *testing.T) {
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
	err := client.Instances.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceForkWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Fork(
		context.TODO(),
		"id",
		hypeman.InstanceForkParams{
			Name:        "my-workload-1-fork",
			FromRunning: hypeman.Bool(false),
			TargetState: hypeman.InstanceForkParamsTargetStateRunning,
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceGet(t *testing.T) {
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
	_, err := client.Instances.Get(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceRestore(t *testing.T) {
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
	_, err := client.Instances.Restore(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStandbyWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Standby(
		context.TODO(),
		"id",
		hypeman.InstanceStandbyParams{
			Compression: shared.SnapshotCompressionConfigParam{
				Enabled:   true,
				Algorithm: shared.SnapshotCompressionConfigAlgorithmZstd,
				Level:     hypeman.Int(1),
			},
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Start(
		context.TODO(),
		"id",
		hypeman.InstanceStartParams{
			Cmd:        []string{"string"},
			Entrypoint: []string{"string"},
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStatWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Stat(
		context.TODO(),
		"id",
		hypeman.InstanceStatParams{
			Path:        "path",
			FollowLinks: hypeman.Bool(true),
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStats(t *testing.T) {
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
	_, err := client.Instances.Stats(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceStop(t *testing.T) {
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
	_, err := client.Instances.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceWaitWithOptionalParams(t *testing.T) {
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
	_, err := client.Instances.Wait(
		context.TODO(),
		"id",
		hypeman.InstanceWaitParams{
			State:   hypeman.InstanceWaitParamsStateCreated,
			Timeout: hypeman.String("timeout"),
		},
	)
	if err != nil {
		var apierr *hypeman.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
