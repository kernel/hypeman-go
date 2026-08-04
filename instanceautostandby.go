// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
)

// InstanceAutoStandbyService contains methods and other services that help with
// interacting with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceAutoStandbyService] method instead.
type InstanceAutoStandbyService struct {
	Options []option.RequestOption
}

// NewInstanceAutoStandbyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstanceAutoStandbyService(opts ...option.RequestOption) (r InstanceAutoStandbyService) {
	r = InstanceAutoStandbyService{}
	r.Options = opts
	return
}

// Places a hold that prevents the auto-standby controller from putting the
// instance into standby before `hold_until`, and cancels any queued auto-standby
// attempt.
//
// Each hold replaces the instance's previous hold, so `hold_until` always reflects
// the most recent call. Holding again after the policy's `idle_timeout` is
// shortened moves `hold_until` earlier.
//
// Callers may use this before opening a connection to a candidate-idle instance: a
// 200 means it is safe to connect until `hold_until`; a 409 means the instance is
// in standby (or irrevocably entering it) and must be restored first.
//
// Instances where auto-standby is disabled, unconfigured, or unsupported return
// 200 with their current status because no auto-standby will occur.
func (r *InstanceAutoStandbyService) Hold(ctx context.Context, id string, opts ...option.RequestOption) (res *AutoStandbyStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/auto-standby/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get auto-standby diagnostic status
func (r *InstanceAutoStandbyService) Status(ctx context.Context, id string, opts ...option.RequestOption) (res *AutoStandbyStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("instances/%s/auto-standby/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
