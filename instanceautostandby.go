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
