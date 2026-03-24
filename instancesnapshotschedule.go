// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/kernel/hypeman-go/internal/encoding/json"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
)

// InstanceSnapshotScheduleService contains methods and other services that help
// with interacting with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceSnapshotScheduleService] method instead.
type InstanceSnapshotScheduleService struct {
	Options []option.RequestOption
}

// NewInstanceSnapshotScheduleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInstanceSnapshotScheduleService(opts ...option.RequestOption) (r InstanceSnapshotScheduleService) {
	r = InstanceSnapshotScheduleService{}
	r.Options = opts
	return
}

// Scheduled runs automatically choose snapshot behavior from current instance
// state:
//
//   - `Running` or `Standby` source: create a `Standby` snapshot.
//   - `Stopped` source: create a `Stopped` snapshot. For running instances, this
//     includes a brief pause/resume cycle during each capture. The minimum supported
//     interval is `1m`, but larger intervals are recommended for heavier or
//     latency-sensitive workloads. Updating only retention, metadata, or
//     `name_prefix` preserves the next scheduled run; changing `interval`
//     establishes a new cadence.
func (r *InstanceSnapshotScheduleService) Update(ctx context.Context, id string, body InstanceSnapshotScheduleUpdateParams, opts ...option.RequestOption) (res *SnapshotSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/snapshot-schedule", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Delete snapshot schedule for an instance
func (r *InstanceSnapshotScheduleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("instances/%s/snapshot-schedule", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get snapshot schedule for an instance
func (r *InstanceSnapshotScheduleService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SnapshotSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/snapshot-schedule", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type InstanceSnapshotScheduleUpdateParams struct {
	SetSnapshotScheduleRequest SetSnapshotScheduleRequestParam
	paramObj
}

func (r InstanceSnapshotScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetSnapshotScheduleRequest)
}
func (r *InstanceSnapshotScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SetSnapshotScheduleRequest)
}
