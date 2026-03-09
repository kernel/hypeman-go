// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/kernel/hypeman-go/internal/apijson"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
	"github.com/kernel/hypeman-go/packages/param"
)

// InstanceSnapshotService contains methods and other services that help with
// interacting with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceSnapshotService] method instead.
type InstanceSnapshotService struct {
	Options []option.RequestOption
}

// NewInstanceSnapshotService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstanceSnapshotService(opts ...option.RequestOption) (r InstanceSnapshotService) {
	r = InstanceSnapshotService{}
	r.Options = opts
	return
}

// Create a snapshot for an instance
func (r *InstanceSnapshotService) New(ctx context.Context, id string, body InstanceSnapshotNewParams, opts ...option.RequestOption) (res *Snapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("instances/%s/snapshots", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Restore an instance from a snapshot in-place
func (r *InstanceSnapshotService) Restore(ctx context.Context, snapshotID string, params InstanceSnapshotRestoreParams, opts ...option.RequestOption) (res *Instance, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return
	}
	path := fmt.Sprintf("instances/%s/snapshots/%s/restore", params.ID, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type InstanceSnapshotNewParams struct {
	// Snapshot capture kind
	//
	// Any of "Standby", "Stopped".
	Kind SnapshotKind `json:"kind,omitzero" api:"required"`
	// Optional snapshot name (lowercase letters, digits, and dashes only; cannot start
	// or end with a dash)
	Name param.Opt[string] `json:"name,omitzero"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r InstanceSnapshotNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceSnapshotNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceSnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceSnapshotRestoreParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Optional hypervisor override. Allowed only when restoring from a Stopped
	// snapshot. Standby snapshots must restore with their original hypervisor.
	//
	// Any of "cloud-hypervisor", "firecracker", "qemu", "vz".
	TargetHypervisor InstanceSnapshotRestoreParamsTargetHypervisor `json:"target_hypervisor,omitzero"`
	// Optional final state after restore. Defaults by snapshot kind:
	//
	// - Standby snapshot defaults to Running
	// - Stopped snapshot defaults to Stopped
	//
	// Any of "Stopped", "Standby", "Running".
	TargetState InstanceSnapshotRestoreParamsTargetState `json:"target_state,omitzero"`
	paramObj
}

func (r InstanceSnapshotRestoreParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceSnapshotRestoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceSnapshotRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional hypervisor override. Allowed only when restoring from a Stopped
// snapshot. Standby snapshots must restore with their original hypervisor.
type InstanceSnapshotRestoreParamsTargetHypervisor string

const (
	InstanceSnapshotRestoreParamsTargetHypervisorCloudHypervisor InstanceSnapshotRestoreParamsTargetHypervisor = "cloud-hypervisor"
	InstanceSnapshotRestoreParamsTargetHypervisorFirecracker     InstanceSnapshotRestoreParamsTargetHypervisor = "firecracker"
	InstanceSnapshotRestoreParamsTargetHypervisorQemu            InstanceSnapshotRestoreParamsTargetHypervisor = "qemu"
	InstanceSnapshotRestoreParamsTargetHypervisorVz              InstanceSnapshotRestoreParamsTargetHypervisor = "vz"
)

// Optional final state after restore. Defaults by snapshot kind:
//
// - Standby snapshot defaults to Running
// - Stopped snapshot defaults to Stopped
type InstanceSnapshotRestoreParamsTargetState string

const (
	InstanceSnapshotRestoreParamsTargetStateStopped InstanceSnapshotRestoreParamsTargetState = "Stopped"
	InstanceSnapshotRestoreParamsTargetStateStandby InstanceSnapshotRestoreParamsTargetState = "Standby"
	InstanceSnapshotRestoreParamsTargetStateRunning InstanceSnapshotRestoreParamsTargetState = "Running"
)
