// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
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
	"github.com/kernel/hypeman-go/shared"
)

// SnapshotService contains methods and other services that help with interacting
// with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapshotService] method instead.
type SnapshotService struct {
	Options []option.RequestOption
}

// NewSnapshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnapshotService(opts ...option.RequestOption) (r SnapshotService) {
	r = SnapshotService{}
	r.Options = opts
	return
}

// List snapshots
func (r *SnapshotService) List(ctx context.Context, query SnapshotListParams, opts ...option.RequestOption) (res *[]Snapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "snapshots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a snapshot
func (r *SnapshotService) Delete(ctx context.Context, snapshotID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return err
	}
	path := fmt.Sprintf("snapshots/%s", snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fork a new instance from a snapshot
func (r *SnapshotService) Fork(ctx context.Context, snapshotID string, body SnapshotForkParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return nil, err
	}
	path := fmt.Sprintf("snapshots/%s/fork", snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get snapshot details
func (r *SnapshotService) Get(ctx context.Context, snapshotID string, opts ...option.RequestOption) (res *Snapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return nil, err
	}
	path := fmt.Sprintf("snapshots/%s", snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Snapshot struct {
	// Auto-generated unique snapshot identifier
	ID string `json:"id" api:"required"`
	// Snapshot creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Snapshot capture kind
	//
	// Any of "Standby", "Stopped".
	Kind SnapshotKind `json:"kind" api:"required"`
	// Total payload size in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Source instance hypervisor at snapshot creation time
	//
	// Any of "cloud-hypervisor", "firecracker", "qemu", "vz".
	SourceHypervisor SnapshotSourceHypervisor `json:"source_hypervisor" api:"required"`
	// Source instance ID at snapshot creation time
	SourceInstanceID string `json:"source_instance_id" api:"required"`
	// Source instance name at snapshot creation time
	SourceInstanceName string `json:"source_instance_name" api:"required"`
	// Compressed memory payload size in bytes
	CompressedSizeBytes int64                            `json:"compressed_size_bytes" api:"nullable"`
	Compression         shared.SnapshotCompressionConfig `json:"compression"`
	// Compression error message when compression_state is error
	CompressionError string `json:"compression_error" api:"nullable"`
	// Compression status of the snapshot payload memory file
	//
	// Any of "none", "compressing", "compressed", "error".
	CompressionState SnapshotCompressionState `json:"compression_state"`
	// Optional human-readable snapshot name (unique per source instance)
	Name string `json:"name" api:"nullable"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags"`
	// Uncompressed memory payload size in bytes
	UncompressedSizeBytes int64 `json:"uncompressed_size_bytes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Kind                  respjson.Field
		SizeBytes             respjson.Field
		SourceHypervisor      respjson.Field
		SourceInstanceID      respjson.Field
		SourceInstanceName    respjson.Field
		CompressedSizeBytes   respjson.Field
		Compression           respjson.Field
		CompressionError      respjson.Field
		CompressionState      respjson.Field
		Name                  respjson.Field
		Tags                  respjson.Field
		UncompressedSizeBytes respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Snapshot) RawJSON() string { return r.JSON.raw }
func (r *Snapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source instance hypervisor at snapshot creation time
type SnapshotSourceHypervisor string

const (
	SnapshotSourceHypervisorCloudHypervisor SnapshotSourceHypervisor = "cloud-hypervisor"
	SnapshotSourceHypervisorFirecracker     SnapshotSourceHypervisor = "firecracker"
	SnapshotSourceHypervisorQemu            SnapshotSourceHypervisor = "qemu"
	SnapshotSourceHypervisorVz              SnapshotSourceHypervisor = "vz"
)

// Compression status of the snapshot payload memory file
type SnapshotCompressionState string

const (
	SnapshotCompressionStateNone        SnapshotCompressionState = "none"
	SnapshotCompressionStateCompressing SnapshotCompressionState = "compressing"
	SnapshotCompressionStateCompressed  SnapshotCompressionState = "compressed"
	SnapshotCompressionStateError       SnapshotCompressionState = "error"
)

// Snapshot capture kind
type SnapshotKind string

const (
	SnapshotKindStandby SnapshotKind = "Standby"
	SnapshotKindStopped SnapshotKind = "Stopped"
)

type SnapshotListParams struct {
	// Filter snapshots by snapshot name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter snapshots by source instance ID
	SourceInstanceID param.Opt[string] `query:"source_instance_id,omitzero" json:"-"`
	// Filter snapshots by kind
	//
	// Any of "Standby", "Stopped".
	Kind SnapshotKind `query:"kind,omitzero" json:"-"`
	// Filter snapshots by tag key-value pairs.
	Tags map[string]string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SnapshotListParams]'s query parameters as `url.Values`.
func (r SnapshotListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SnapshotForkParams struct {
	// Name for the new instance (lowercase letters, digits, and dashes only; cannot
	// start or end with a dash)
	Name string `json:"name" api:"required"`
	// Optional hypervisor override. Allowed only when forking from a Stopped snapshot.
	// Standby snapshots must fork with their original hypervisor.
	//
	// Any of "cloud-hypervisor", "firecracker", "qemu", "vz".
	TargetHypervisor SnapshotForkParamsTargetHypervisor `json:"target_hypervisor,omitzero"`
	// Optional final state for the forked instance. Defaults by snapshot kind:
	//
	// - Standby snapshot defaults to Running
	// - Stopped snapshot defaults to Stopped
	//
	// Any of "Stopped", "Standby", "Running".
	TargetState SnapshotForkParamsTargetState `json:"target_state,omitzero"`
	paramObj
}

func (r SnapshotForkParams) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotForkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotForkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional hypervisor override. Allowed only when forking from a Stopped snapshot.
// Standby snapshots must fork with their original hypervisor.
type SnapshotForkParamsTargetHypervisor string

const (
	SnapshotForkParamsTargetHypervisorCloudHypervisor SnapshotForkParamsTargetHypervisor = "cloud-hypervisor"
	SnapshotForkParamsTargetHypervisorFirecracker     SnapshotForkParamsTargetHypervisor = "firecracker"
	SnapshotForkParamsTargetHypervisorQemu            SnapshotForkParamsTargetHypervisor = "qemu"
	SnapshotForkParamsTargetHypervisorVz              SnapshotForkParamsTargetHypervisor = "vz"
)

// Optional final state for the forked instance. Defaults by snapshot kind:
//
// - Standby snapshot defaults to Running
// - Stopped snapshot defaults to Stopped
type SnapshotForkParamsTargetState string

const (
	SnapshotForkParamsTargetStateStopped SnapshotForkParamsTargetState = "Stopped"
	SnapshotForkParamsTargetStateStandby SnapshotForkParamsTargetState = "Standby"
	SnapshotForkParamsTargetStateRunning SnapshotForkParamsTargetState = "Running"
)
