// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/kernel/hypeman-go/internal/apijson"
	"github.com/kernel/hypeman-go/internal/apiquery"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
	"github.com/kernel/hypeman-go/packages/param"
	"github.com/kernel/hypeman-go/packages/respjson"
	"github.com/kernel/hypeman-go/packages/ssestream"
	"github.com/kernel/hypeman-go/shared"
)

// InstanceService contains methods and other services that help with interacting
// with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options          []option.RequestOption
	Volumes          InstanceVolumeService
	Snapshots        InstanceSnapshotService
	SnapshotSchedule InstanceSnapshotScheduleService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r InstanceService) {
	r = InstanceService{}
	r.Options = opts
	r.Volumes = NewInstanceVolumeService(opts...)
	r.Snapshots = NewInstanceSnapshotService(opts...)
	r.SnapshotSchedule = NewInstanceSnapshotScheduleService(opts...)
	return
}

// Create and start instance
func (r *InstanceService) New(ctx context.Context, body InstanceNewParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update mutable properties of a running instance. Currently supports updating
// only the environment variables referenced by existing credential policies,
// enabling secret/key rotation without instance restart.
func (r *InstanceService) Update(ctx context.Context, id string, body InstanceUpdateParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List instances
func (r *InstanceService) List(ctx context.Context, query InstanceListParams, opts ...option.RequestOption) (res *[]Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stop and delete instance
func (r *InstanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fork an instance from stopped, standby, or running (with from_running=true)
func (r *InstanceService) Fork(ctx context.Context, id string, body InstanceForkParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/fork", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get instance details
func (r *InstanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Streams instance logs as Server-Sent Events. Use the `source` parameter to
// select which log to stream:
//
// - `app` (default): Guest application logs (serial console)
// - `vmm`: Cloud Hypervisor VMM logs
// - `hypeman`: Hypeman operations log
//
// Returns the last N lines (controlled by `tail` parameter), then optionally
// continues streaming new lines if `follow=true`.
func (r *InstanceService) LogsStreaming(ctx context.Context, id string, query InstanceLogsParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return ssestream.NewStream[string](nil, err)
	}
	path := fmt.Sprintf("instances/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

// Restore instance from standby
func (r *InstanceService) Restore(ctx context.Context, id string, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/restore", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Put instance in standby (pause, snapshot, delete VMM)
func (r *InstanceService) Standby(ctx context.Context, id string, body InstanceStandbyParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/standby", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Start a stopped instance
func (r *InstanceService) Start(ctx context.Context, id string, body InstanceStartParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/start", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns information about a path in the guest filesystem. Useful for checking if
// a path exists, its type, and permissions before performing file operations.
func (r *InstanceService) Stat(ctx context.Context, id string, query InstanceStatParams, opts ...option.RequestOption) (res *PathInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/stat", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns real-time resource utilization statistics for a running VM instance.
// Metrics are collected from /proc/<pid>/stat and /proc/<pid>/statm for CPU and
// memory, and from TAP interface statistics for network I/O.
func (r *InstanceService) Stats(ctx context.Context, id string, opts ...option.RequestOption) (res *InstanceStats, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/stats", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stop instance (graceful shutdown)
func (r *InstanceService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/stop", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Blocks until the instance reaches the specified target state, the timeout
// expires, or the instance enters a terminal/error state. Useful for avoiding
// client-side polling when waiting for state transitions (e.g. waiting for an
// instance to become Running).
func (r *InstanceService) Wait(ctx context.Context, id string, query InstanceWaitParams, opts ...option.RequestOption) (res *WaitForStateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/wait", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Instance struct {
	// Auto-generated unique identifier (CUID2 format)
	ID string `json:"id" api:"required"`
	// Creation timestamp (RFC3339)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// OCI image reference
	Image string `json:"image" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Instance state:
	//
	// - Created: VMM created but not started (Cloud Hypervisor native)
	// - Initializing: VM is running while guest init is still in progress
	// - Running: Guest program has started and instance is ready
	// - Paused: VM is paused (Cloud Hypervisor native)
	// - Shutdown: VM shut down but VMM exists (Cloud Hypervisor native)
	// - Stopped: No VMM running, no snapshot exists
	// - Standby: No VMM running, snapshot exists (can be restored)
	// - Unknown: Failed to determine state (see state_error for details)
	//
	// Any of "Created", "Initializing", "Running", "Paused", "Shutdown", "Stopped",
	// "Standby", "Unknown".
	State InstanceState `json:"state" api:"required"`
	// Disk I/O rate limit (human-readable, e.g., "100MB/s")
	DiskIoBps string `json:"disk_io_bps"`
	// Environment variables
	Env map[string]string `json:"env"`
	// App exit code (null if VM hasn't exited)
	ExitCode int64 `json:"exit_code" api:"nullable"`
	// Human-readable description of exit (e.g., "command not found", "killed by signal
	// 9 (SIGKILL) - OOM")
	ExitMessage string `json:"exit_message"`
	// GPU information attached to the instance
	GPU InstanceGPU `json:"gpu"`
	// Whether a snapshot exists for this instance
	HasSnapshot bool `json:"has_snapshot"`
	// Hotplug memory size (human-readable)
	HotplugSize string `json:"hotplug_size"`
	// Hypervisor running this instance
	//
	// Any of "cloud-hypervisor", "firecracker", "qemu", "vz".
	Hypervisor InstanceHypervisor `json:"hypervisor"`
	// Network configuration of the instance
	Network InstanceNetwork `json:"network"`
	// Writable overlay disk size (human-readable)
	OverlaySize string `json:"overlay_size"`
	// Base memory size (human-readable)
	Size           string         `json:"size"`
	SnapshotPolicy SnapshotPolicy `json:"snapshot_policy"`
	// Start timestamp (RFC3339)
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Error message if state couldn't be determined (only set when state is Unknown)
	StateError string `json:"state_error" api:"nullable"`
	// Stop timestamp (RFC3339)
	StoppedAt time.Time `json:"stopped_at" api:"nullable" format:"date-time"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags"`
	// Number of virtual CPUs
	Vcpus int64 `json:"vcpus"`
	// Volumes attached to the instance
	Volumes []VolumeMount `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Image          respjson.Field
		Name           respjson.Field
		State          respjson.Field
		DiskIoBps      respjson.Field
		Env            respjson.Field
		ExitCode       respjson.Field
		ExitMessage    respjson.Field
		GPU            respjson.Field
		HasSnapshot    respjson.Field
		HotplugSize    respjson.Field
		Hypervisor     respjson.Field
		Network        respjson.Field
		OverlaySize    respjson.Field
		Size           respjson.Field
		SnapshotPolicy respjson.Field
		StartedAt      respjson.Field
		StateError     respjson.Field
		StoppedAt      respjson.Field
		Tags           respjson.Field
		Vcpus          respjson.Field
		Volumes        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Instance) RawJSON() string { return r.JSON.raw }
func (r *Instance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance state:
//
// - Created: VMM created but not started (Cloud Hypervisor native)
// - Initializing: VM is running while guest init is still in progress
// - Running: Guest program has started and instance is ready
// - Paused: VM is paused (Cloud Hypervisor native)
// - Shutdown: VM shut down but VMM exists (Cloud Hypervisor native)
// - Stopped: No VMM running, no snapshot exists
// - Standby: No VMM running, snapshot exists (can be restored)
// - Unknown: Failed to determine state (see state_error for details)
type InstanceState string

const (
	InstanceStateCreated      InstanceState = "Created"
	InstanceStateInitializing InstanceState = "Initializing"
	InstanceStateRunning      InstanceState = "Running"
	InstanceStatePaused       InstanceState = "Paused"
	InstanceStateShutdown     InstanceState = "Shutdown"
	InstanceStateStopped      InstanceState = "Stopped"
	InstanceStateStandby      InstanceState = "Standby"
	InstanceStateUnknown      InstanceState = "Unknown"
)

// GPU information attached to the instance
type InstanceGPU struct {
	// mdev device UUID
	MdevUuid string `json:"mdev_uuid"`
	// vGPU profile name
	Profile string `json:"profile"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MdevUuid    respjson.Field
		Profile     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceGPU) RawJSON() string { return r.JSON.raw }
func (r *InstanceGPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hypervisor running this instance
type InstanceHypervisor string

const (
	InstanceHypervisorCloudHypervisor InstanceHypervisor = "cloud-hypervisor"
	InstanceHypervisorFirecracker     InstanceHypervisor = "firecracker"
	InstanceHypervisorQemu            InstanceHypervisor = "qemu"
	InstanceHypervisorVz              InstanceHypervisor = "vz"
)

// Network configuration of the instance
type InstanceNetwork struct {
	// Download bandwidth limit (human-readable, e.g., "1Gbps", "125MB/s")
	BandwidthDownload string `json:"bandwidth_download"`
	// Upload bandwidth limit (human-readable, e.g., "1Gbps", "125MB/s")
	BandwidthUpload string `json:"bandwidth_upload"`
	// Whether instance is attached to the default network
	Enabled bool `json:"enabled"`
	// Assigned IP address (null if no network)
	IP string `json:"ip" api:"nullable"`
	// Assigned MAC address (null if no network)
	Mac string `json:"mac" api:"nullable"`
	// Network name (always "default" when enabled)
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthDownload respjson.Field
		BandwidthUpload   respjson.Field
		Enabled           respjson.Field
		IP                respjson.Field
		Mac               respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceNetwork) RawJSON() string { return r.JSON.raw }
func (r *InstanceNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Real-time resource utilization statistics for a VM instance
type InstanceStats struct {
	// Total memory allocated to the VM (Size + HotplugSize) in bytes
	AllocatedMemoryBytes int64 `json:"allocated_memory_bytes" api:"required"`
	// Number of vCPUs allocated to the VM
	AllocatedVcpus int64 `json:"allocated_vcpus" api:"required"`
	// Total CPU time consumed by the VM hypervisor process in seconds
	CPUSeconds float64 `json:"cpu_seconds" api:"required"`
	// Instance identifier
	InstanceID string `json:"instance_id" api:"required"`
	// Instance name
	InstanceName string `json:"instance_name" api:"required"`
	// Resident Set Size - actual physical memory used by the VM in bytes
	MemoryRssBytes int64 `json:"memory_rss_bytes" api:"required"`
	// Virtual Memory Size - total virtual memory allocated in bytes
	MemoryVmsBytes int64 `json:"memory_vms_bytes" api:"required"`
	// Total network bytes received by the VM (from TAP interface)
	NetworkRxBytes int64 `json:"network_rx_bytes" api:"required"`
	// Total network bytes transmitted by the VM (from TAP interface)
	NetworkTxBytes int64 `json:"network_tx_bytes" api:"required"`
	// Memory utilization ratio (RSS / allocated memory). Only present when
	// allocated_memory_bytes > 0.
	MemoryUtilizationRatio float64 `json:"memory_utilization_ratio" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllocatedMemoryBytes   respjson.Field
		AllocatedVcpus         respjson.Field
		CPUSeconds             respjson.Field
		InstanceID             respjson.Field
		InstanceName           respjson.Field
		MemoryRssBytes         respjson.Field
		MemoryVmsBytes         respjson.Field
		NetworkRxBytes         respjson.Field
		NetworkTxBytes         respjson.Field
		MemoryUtilizationRatio respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceStats) RawJSON() string { return r.JSON.raw }
func (r *InstanceStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PathInfo struct {
	// Whether the path exists
	Exists bool `json:"exists" api:"required"`
	// Error message if stat failed (e.g., permission denied). Only set when exists is
	// false due to an error rather than the path not existing.
	Error string `json:"error" api:"nullable"`
	// True if this is a directory
	IsDir bool `json:"is_dir"`
	// True if this is a regular file
	IsFile bool `json:"is_file"`
	// True if this is a symbolic link (only set when follow_links=false)
	IsSymlink bool `json:"is_symlink"`
	// Symlink target path (only set when is_symlink=true)
	LinkTarget string `json:"link_target" api:"nullable"`
	// File mode (Unix permissions)
	Mode int64 `json:"mode"`
	// File size in bytes
	Size int64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exists      respjson.Field
		Error       respjson.Field
		IsDir       respjson.Field
		IsFile      respjson.Field
		IsSymlink   respjson.Field
		LinkTarget  respjson.Field
		Mode        respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PathInfo) RawJSON() string { return r.JSON.raw }
func (r *PathInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Interval, Retention are required.
type SetSnapshotScheduleRequestParam struct {
	// Snapshot interval (Go duration format, minimum 1m).
	Interval string `json:"interval" api:"required"`
	// At least one of max_count or max_age must be provided.
	Retention SnapshotScheduleRetentionParam `json:"retention,omitzero" api:"required"`
	// Optional prefix for auto-generated scheduled snapshot names (max 47 chars).
	NamePrefix param.Opt[string] `json:"name_prefix,omitzero"`
	// User-defined key-value tags.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r SetSnapshotScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SetSnapshotScheduleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetSnapshotScheduleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotPolicy struct {
	Compression shared.SnapshotCompressionConfig `json:"compression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compression respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotPolicy) RawJSON() string { return r.JSON.raw }
func (r *SnapshotPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SnapshotPolicy to a SnapshotPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SnapshotPolicyParam.Overrides()
func (r SnapshotPolicy) ToParam() SnapshotPolicyParam {
	return param.Override[SnapshotPolicyParam](json.RawMessage(r.RawJSON()))
}

type SnapshotPolicyParam struct {
	Compression shared.SnapshotCompressionConfigParam `json:"compression,omitzero"`
	paramObj
}

func (r SnapshotPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotSchedule struct {
	// Schedule creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Source instance ID.
	InstanceID string `json:"instance_id" api:"required"`
	// Snapshot interval (Go duration format).
	Interval string `json:"interval" api:"required"`
	// Next scheduled run time.
	NextRunAt time.Time `json:"next_run_at" api:"required" format:"date-time"`
	// Automatic cleanup policy for scheduled snapshots.
	Retention SnapshotScheduleRetention `json:"retention" api:"required"`
	// Schedule update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Last schedule run error, if any.
	LastError string `json:"last_error" api:"nullable"`
	// Last schedule execution time.
	LastRunAt time.Time `json:"last_run_at" api:"nullable" format:"date-time"`
	// Snapshot ID produced by the last successful run.
	LastSnapshotID string `json:"last_snapshot_id" api:"nullable"`
	// User-defined key-value tags.
	Metadata map[string]string `json:"metadata"`
	// Optional prefix used for generated scheduled snapshot names.
	NamePrefix string `json:"name_prefix" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		InstanceID     respjson.Field
		Interval       respjson.Field
		NextRunAt      respjson.Field
		Retention      respjson.Field
		UpdatedAt      respjson.Field
		LastError      respjson.Field
		LastRunAt      respjson.Field
		LastSnapshotID respjson.Field
		Metadata       respjson.Field
		NamePrefix     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotSchedule) RawJSON() string { return r.JSON.raw }
func (r *SnapshotSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic cleanup policy for scheduled snapshots.
type SnapshotScheduleRetention struct {
	// Delete scheduled snapshots older than this duration (Go duration format).
	MaxAge string `json:"max_age"`
	// Keep at most this many scheduled snapshots for the instance (0 disables
	// count-based cleanup).
	MaxCount int64 `json:"max_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAge      respjson.Field
		MaxCount    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotScheduleRetention) RawJSON() string { return r.JSON.raw }
func (r *SnapshotScheduleRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SnapshotScheduleRetention to a
// SnapshotScheduleRetentionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SnapshotScheduleRetentionParam.Overrides()
func (r SnapshotScheduleRetention) ToParam() SnapshotScheduleRetentionParam {
	return param.Override[SnapshotScheduleRetentionParam](json.RawMessage(r.RawJSON()))
}

// Automatic cleanup policy for scheduled snapshots.
type SnapshotScheduleRetentionParam struct {
	// Delete scheduled snapshots older than this duration (Go duration format).
	MaxAge param.Opt[string] `json:"max_age,omitzero"`
	// Keep at most this many scheduled snapshots for the instance (0 disables
	// count-based cleanup).
	MaxCount param.Opt[int64] `json:"max_count,omitzero"`
	paramObj
}

func (r SnapshotScheduleRetentionParam) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotScheduleRetentionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotScheduleRetentionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeMount struct {
	// Path where volume is mounted in the guest
	MountPath string `json:"mount_path" api:"required"`
	// Volume identifier
	VolumeID string `json:"volume_id" api:"required"`
	// Create per-instance overlay for writes (requires readonly=true)
	Overlay bool `json:"overlay"`
	// Max overlay size as human-readable string (e.g., "1GB"). Required if
	// overlay=true.
	OverlaySize string `json:"overlay_size"`
	// Whether volume is mounted read-only
	Readonly bool `json:"readonly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MountPath   respjson.Field
		VolumeID    respjson.Field
		Overlay     respjson.Field
		OverlaySize respjson.Field
		Readonly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeMount) RawJSON() string { return r.JSON.raw }
func (r *VolumeMount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeMount to a VolumeMountParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeMountParam.Overrides()
func (r VolumeMount) ToParam() VolumeMountParam {
	return param.Override[VolumeMountParam](json.RawMessage(r.RawJSON()))
}

// The properties MountPath, VolumeID are required.
type VolumeMountParam struct {
	// Path where volume is mounted in the guest
	MountPath string `json:"mount_path" api:"required"`
	// Volume identifier
	VolumeID string `json:"volume_id" api:"required"`
	// Create per-instance overlay for writes (requires readonly=true)
	Overlay param.Opt[bool] `json:"overlay,omitzero"`
	// Max overlay size as human-readable string (e.g., "1GB"). Required if
	// overlay=true.
	OverlaySize param.Opt[string] `json:"overlay_size,omitzero"`
	// Whether volume is mounted read-only
	Readonly param.Opt[bool] `json:"readonly,omitzero"`
	paramObj
}

func (r VolumeMountParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeMountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeMountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaitForStateResponse struct {
	// Current instance state when the wait completed
	//
	// Any of "Created", "Initializing", "Running", "Paused", "Shutdown", "Stopped",
	// "Standby", "Unknown".
	State WaitForStateResponseState `json:"state" api:"required"`
	// Whether the timeout expired before the target state was reached
	TimedOut bool `json:"timed_out" api:"required"`
	// Error message when derived state is Unknown
	StateError string `json:"state_error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		State       respjson.Field
		TimedOut    respjson.Field
		StateError  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaitForStateResponse) RawJSON() string { return r.JSON.raw }
func (r *WaitForStateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current instance state when the wait completed
type WaitForStateResponseState string

const (
	WaitForStateResponseStateCreated      WaitForStateResponseState = "Created"
	WaitForStateResponseStateInitializing WaitForStateResponseState = "Initializing"
	WaitForStateResponseStateRunning      WaitForStateResponseState = "Running"
	WaitForStateResponseStatePaused       WaitForStateResponseState = "Paused"
	WaitForStateResponseStateShutdown     WaitForStateResponseState = "Shutdown"
	WaitForStateResponseStateStopped      WaitForStateResponseState = "Stopped"
	WaitForStateResponseStateStandby      WaitForStateResponseState = "Standby"
	WaitForStateResponseStateUnknown      WaitForStateResponseState = "Unknown"
)

type InstanceNewParams struct {
	// OCI image reference
	Image string `json:"image" api:"required"`
	// Human-readable name (lowercase letters, digits, and dashes only; cannot start or
	// end with a dash)
	Name string `json:"name" api:"required"`
	// Disk I/O rate limit (e.g., "100MB/s", "500MB/s"). Defaults to proportional share
	// based on CPU allocation if configured.
	DiskIoBps param.Opt[string] `json:"disk_io_bps,omitzero"`
	// Additional memory for hotplug (human-readable format like "3GB", "1G"). Omit to
	// disable hotplug memory.
	HotplugSize param.Opt[string] `json:"hotplug_size,omitzero"`
	// Writable overlay disk size (human-readable format like "10GB", "50G")
	OverlaySize param.Opt[string] `json:"overlay_size,omitzero"`
	// Base memory size (human-readable format like "1GB", "512MB", "2G")
	Size param.Opt[string] `json:"size,omitzero"`
	// Skip guest-agent installation during boot. When true, the exec and stat APIs
	// will not work for this instance. The instance will still run, but remote command
	// execution will be unavailable.
	SkipGuestAgent param.Opt[bool] `json:"skip_guest_agent,omitzero"`
	// Skip kernel headers installation during boot for faster startup. When true, DKMS
	// (Dynamic Kernel Module Support) will not work, preventing compilation of
	// out-of-tree kernel modules (e.g., NVIDIA vGPU drivers). Recommended for
	// workloads that don't need kernel module compilation.
	SkipKernelHeaders param.Opt[bool] `json:"skip_kernel_headers,omitzero"`
	// Number of virtual CPUs
	Vcpus param.Opt[int64] `json:"vcpus,omitzero"`
	// Override image CMD (like docker run <image> <command>). Omit to use image
	// default.
	Cmd []string `json:"cmd,omitzero"`
	// Host-managed credential brokering policies keyed by guest-visible env var name.
	// Those guest env vars receive mock placeholder values, while the real values
	// remain host-scoped in the request `env` map and are only materialized on the
	// mediated egress path according to each credential's `source` and `inject` rules.
	Credentials map[string]InstanceNewParamsCredential `json:"credentials,omitzero"`
	// Device IDs or names to attach for GPU/PCI passthrough
	Devices []string `json:"devices,omitzero"`
	// Override image entrypoint (like docker run --entrypoint). Omit to use image
	// default.
	Entrypoint []string `json:"entrypoint,omitzero"`
	// Environment variables
	Env map[string]string `json:"env,omitzero"`
	// GPU configuration for the instance
	GPU InstanceNewParamsGPU `json:"gpu,omitzero"`
	// Hypervisor to use for this instance. Defaults to server configuration.
	//
	// Any of "cloud-hypervisor", "firecracker", "qemu", "vz".
	Hypervisor InstanceNewParamsHypervisor `json:"hypervisor,omitzero"`
	// Network configuration for the instance
	Network InstanceNewParamsNetwork `json:"network,omitzero"`
	// Snapshot compression policy for this instance. Controls compression settings
	// applied when creating snapshots or entering standby.
	SnapshotPolicy SnapshotPolicyParam `json:"snapshot_policy,omitzero"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags,omitzero"`
	// Volumes to attach to the instance at creation time
	Volumes []VolumeMountParam `json:"volumes,omitzero"`
	paramObj
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Inject, Source are required.
type InstanceNewParamsCredential struct {
	Inject []InstanceNewParamsCredentialInject `json:"inject,omitzero" api:"required"`
	Source InstanceNewParamsCredentialSource   `json:"source,omitzero" api:"required"`
	paramObj
}

func (r InstanceNewParamsCredential) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property As is required.
type InstanceNewParamsCredentialInject struct {
	// Current v1 transform shape. Header templating is supported now; other transform
	// types (for example request signing) can be added in future revisions.
	As InstanceNewParamsCredentialInjectAs `json:"as,omitzero" api:"required"`
	// Optional destination host patterns (`api.example.com`, `*.example.com`). Omit to
	// allow injection on all destinations.
	Hosts []string `json:"hosts,omitzero"`
	paramObj
}

func (r InstanceNewParamsCredentialInject) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsCredentialInject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsCredentialInject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current v1 transform shape. Header templating is supported now; other transform
// types (for example request signing) can be added in future revisions.
//
// The properties Format, Header are required.
type InstanceNewParamsCredentialInjectAs struct {
	// Template that must include `${value}`.
	Format string `json:"format" api:"required"`
	// Header name to set/mutate for matching outbound requests.
	Header string `json:"header" api:"required"`
	paramObj
}

func (r InstanceNewParamsCredentialInjectAs) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsCredentialInjectAs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsCredentialInjectAs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Env is required.
type InstanceNewParamsCredentialSource struct {
	// Name of the real credential in the request `env` map. The guest-visible env var
	// key can receive a mock placeholder, while the mediated egress path resolves that
	// placeholder back to this real value only on the host.
	Env string `json:"env" api:"required"`
	paramObj
}

func (r InstanceNewParamsCredentialSource) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsCredentialSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsCredentialSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU configuration for the instance
type InstanceNewParamsGPU struct {
	// vGPU profile name (e.g., "L40S-1Q"). Only used in vGPU mode.
	Profile param.Opt[string] `json:"profile,omitzero"`
	paramObj
}

func (r InstanceNewParamsGPU) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsGPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsGPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hypervisor to use for this instance. Defaults to server configuration.
type InstanceNewParamsHypervisor string

const (
	InstanceNewParamsHypervisorCloudHypervisor InstanceNewParamsHypervisor = "cloud-hypervisor"
	InstanceNewParamsHypervisorFirecracker     InstanceNewParamsHypervisor = "firecracker"
	InstanceNewParamsHypervisorQemu            InstanceNewParamsHypervisor = "qemu"
	InstanceNewParamsHypervisorVz              InstanceNewParamsHypervisor = "vz"
)

// Network configuration for the instance
type InstanceNewParamsNetwork struct {
	// Download bandwidth limit (external→VM, e.g., "1Gbps", "125MB/s"). Defaults to
	// proportional share based on CPU allocation.
	BandwidthDownload param.Opt[string] `json:"bandwidth_download,omitzero"`
	// Upload bandwidth limit (VM→external, e.g., "1Gbps", "125MB/s"). Defaults to
	// proportional share based on CPU allocation.
	BandwidthUpload param.Opt[string] `json:"bandwidth_upload,omitzero"`
	// Whether to attach instance to the default network
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Host-mediated outbound network policy. Omit this object, or set
	// `enabled: false`, to preserve normal direct outbound networking when
	// `network.enabled` is true.
	Egress InstanceNewParamsNetworkEgress `json:"egress,omitzero"`
	paramObj
}

func (r InstanceNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Host-mediated outbound network policy. Omit this object, or set
// `enabled: false`, to preserve normal direct outbound networking when
// `network.enabled` is true.
type InstanceNewParamsNetworkEgress struct {
	// Whether to enable the mediated egress path. When false or omitted, the instance
	// keeps normal direct outbound networking and host-managed credential rewriting is
	// disabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Egress enforcement policy applied when mediation is enabled.
	Enforcement InstanceNewParamsNetworkEgressEnforcement `json:"enforcement,omitzero"`
	paramObj
}

func (r InstanceNewParamsNetworkEgress) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsNetworkEgress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsNetworkEgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Egress enforcement policy applied when mediation is enabled.
type InstanceNewParamsNetworkEgressEnforcement struct {
	// `all` (default) rejects direct non-mediated TCP egress from the VM, while
	// `http_https_only` rejects direct egress only on TCP ports 80 and 443.
	//
	// Any of "all", "http_https_only".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r InstanceNewParamsNetworkEgressEnforcement) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsNetworkEgressEnforcement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsNetworkEgressEnforcement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceNewParamsNetworkEgressEnforcement](
		"mode", "all", "http_https_only",
	)
}

type InstanceUpdateParams struct {
	// Environment variables to update (merged with existing). Only keys referenced by
	// the instance's existing credential `source.env` bindings are accepted. Use this
	// to rotate real credential values without restarting the VM.
	Env map[string]string `json:"env,omitzero"`
	paramObj
}

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceListParams struct {
	// Filter instances by state (e.g., Running, Stopped)
	//
	// Any of "Created", "Initializing", "Running", "Paused", "Shutdown", "Stopped",
	// "Standby", "Unknown".
	State InstanceListParamsState `query:"state,omitzero" json:"-"`
	// Filter instances by tag key-value pairs. Uses deepObject style:
	// ?tags[team]=backend&tags[env]=staging Multiple entries are ANDed together. All
	// specified key-value pairs must match.
	Tags map[string]string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter instances by state (e.g., Running, Stopped)
type InstanceListParamsState string

const (
	InstanceListParamsStateCreated      InstanceListParamsState = "Created"
	InstanceListParamsStateInitializing InstanceListParamsState = "Initializing"
	InstanceListParamsStateRunning      InstanceListParamsState = "Running"
	InstanceListParamsStatePaused       InstanceListParamsState = "Paused"
	InstanceListParamsStateShutdown     InstanceListParamsState = "Shutdown"
	InstanceListParamsStateStopped      InstanceListParamsState = "Stopped"
	InstanceListParamsStateStandby      InstanceListParamsState = "Standby"
	InstanceListParamsStateUnknown      InstanceListParamsState = "Unknown"
)

type InstanceForkParams struct {
	// Name for the forked instance (lowercase letters, digits, and dashes only; cannot
	// start or end with a dash)
	Name string `json:"name" api:"required"`
	// Allow forking from a running source instance. When true and source is Running,
	// the source is put into standby, forked, then restored back to Running.
	FromRunning param.Opt[bool] `json:"from_running,omitzero"`
	// Optional final state for the forked instance. Default is the source instance
	// state at fork time. For example, forking from Running defaults the fork result
	// to Running.
	//
	// Any of "Stopped", "Standby", "Running".
	TargetState InstanceForkParamsTargetState `json:"target_state,omitzero"`
	paramObj
}

func (r InstanceForkParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceForkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceForkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional final state for the forked instance. Default is the source instance
// state at fork time. For example, forking from Running defaults the fork result
// to Running.
type InstanceForkParamsTargetState string

const (
	InstanceForkParamsTargetStateStopped InstanceForkParamsTargetState = "Stopped"
	InstanceForkParamsTargetStateStandby InstanceForkParamsTargetState = "Standby"
	InstanceForkParamsTargetStateRunning InstanceForkParamsTargetState = "Running"
)

type InstanceLogsParams struct {
	// Continue streaming new lines after initial output
	Follow param.Opt[bool] `query:"follow,omitzero" json:"-"`
	// Number of lines to return from end
	Tail param.Opt[int64] `query:"tail,omitzero" json:"-"`
	// Log source to stream:
	//
	// - app: Guest application logs (serial console output)
	// - vmm: Cloud Hypervisor VMM logs (hypervisor stdout+stderr)
	// - hypeman: Hypeman operations log (actions taken on this instance)
	//
	// Any of "app", "vmm", "hypeman".
	Source InstanceLogsParamsSource `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceLogsParams]'s query parameters as `url.Values`.
func (r InstanceLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Log source to stream:
//
// - app: Guest application logs (serial console output)
// - vmm: Cloud Hypervisor VMM logs (hypervisor stdout+stderr)
// - hypeman: Hypeman operations log (actions taken on this instance)
type InstanceLogsParamsSource string

const (
	InstanceLogsParamsSourceApp     InstanceLogsParamsSource = "app"
	InstanceLogsParamsSourceVmm     InstanceLogsParamsSource = "vmm"
	InstanceLogsParamsSourceHypeman InstanceLogsParamsSource = "hypeman"
)

type InstanceStandbyParams struct {
	Compression shared.SnapshotCompressionConfigParam `json:"compression,omitzero"`
	paramObj
}

func (r InstanceStandbyParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceStandbyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceStandbyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceStartParams struct {
	// Override image CMD for this run. Omit to keep previous value.
	Cmd []string `json:"cmd,omitzero"`
	// Override image entrypoint for this run. Omit to keep previous value.
	Entrypoint []string `json:"entrypoint,omitzero"`
	paramObj
}

func (r InstanceStartParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceStatParams struct {
	// Path to stat in the guest filesystem
	Path string `query:"path" api:"required" json:"-"`
	// Follow symbolic links (like stat vs lstat)
	FollowLinks param.Opt[bool] `query:"follow_links,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceStatParams]'s query parameters as `url.Values`.
func (r InstanceStatParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstanceWaitParams struct {
	// Target state to wait for
	//
	// Any of "Created", "Initializing", "Running", "Paused", "Shutdown", "Stopped",
	// "Standby", "Unknown".
	State InstanceWaitParamsState `query:"state,omitzero" api:"required" json:"-"`
	// Maximum duration to wait (Go duration format, e.g. "30s", "2m"). Capped at 5
	// minutes. Defaults to 60 seconds.
	Timeout param.Opt[string] `query:"timeout,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceWaitParams]'s query parameters as `url.Values`.
func (r InstanceWaitParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Target state to wait for
type InstanceWaitParamsState string

const (
	InstanceWaitParamsStateCreated      InstanceWaitParamsState = "Created"
	InstanceWaitParamsStateInitializing InstanceWaitParamsState = "Initializing"
	InstanceWaitParamsStateRunning      InstanceWaitParamsState = "Running"
	InstanceWaitParamsStatePaused       InstanceWaitParamsState = "Paused"
	InstanceWaitParamsStateShutdown     InstanceWaitParamsState = "Shutdown"
	InstanceWaitParamsStateStopped      InstanceWaitParamsState = "Stopped"
	InstanceWaitParamsStateStandby      InstanceWaitParamsState = "Standby"
	InstanceWaitParamsStateUnknown      InstanceWaitParamsState = "Unknown"
)
