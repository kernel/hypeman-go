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
)

// BuilderService contains methods and other services that help with interacting
// with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuilderService] method instead.
type BuilderService struct {
	Options []option.RequestOption
}

// NewBuilderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuilderService(opts ...option.RequestOption) (r BuilderService) {
	r = BuilderService{}
	r.Options = opts
	return
}

// Creates a builder and its cache disk. One build at a time runs per builder.
func (r *BuilderService) New(ctx context.Context, body BuilderNewParams, opts ...option.RequestOption) (res *Builder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "builders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List builders
func (r *BuilderService) List(ctx context.Context, query BuilderListParams, opts ...option.RequestOption) (res *[]Builder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "builders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a builder and its cache disk.
func (r *BuilderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("builders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get builder details
func (r *BuilderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Builder, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("builders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Resets the builder's cache disk. The builder transitions to pruning, then ready.
// Builder identity is preserved.
func (r *BuilderService) Prune(ctx context.Context, id string, opts ...option.RequestOption) (res *Builder, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("builders/%s/prune", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Builder struct {
	// Builder identifier
	ID string `json:"id" api:"required"`
	// Creation timestamp (RFC3339)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Persistent builder cache disk size in gigabytes. Cannot be changed after
	// creation.
	DiskSizeGB int64 `json:"disk_size_gb" api:"required"`
	// Maximum concurrent builds on this builder. Currently fixed at 1.
	MaxConcurrency int64 `json:"max_concurrency" api:"required"`
	// Point-in-time IDs of queued builds waiting for this builder, oldest first
	QueuedBuilds []string `json:"queued_builds" api:"required"`
	// Builder lifecycle status
	//
	// Any of "ready", "pruning", "deleting", "error".
	Status BuilderStatus `json:"status" api:"required"`
	// Point-in-time ID of the build currently running on this builder
	ActiveBuildID string `json:"active_build_id" api:"nullable"`
	// When a build last ran on this builder
	LastUsedAt time.Time `json:"last_used_at" api:"nullable" format:"date-time"`
	// Optional non-unique display name
	Name string `json:"name"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DiskSizeGB     respjson.Field
		MaxConcurrency respjson.Field
		QueuedBuilds   respjson.Field
		Status         respjson.Field
		ActiveBuildID  respjson.Field
		LastUsedAt     respjson.Field
		Name           respjson.Field
		Tags           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Builder) RawJSON() string { return r.JSON.raw }
func (r *Builder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Builder lifecycle status
type BuilderStatus string

const (
	BuilderStatusReady    BuilderStatus = "ready"
	BuilderStatusPruning  BuilderStatus = "pruning"
	BuilderStatusDeleting BuilderStatus = "deleting"
	BuilderStatusError    BuilderStatus = "error"
)

type BuilderNewParams struct {
	// Optional caller-supplied identifier, auto-generated if not provided
	ID param.Opt[string] `json:"id,omitzero"`
	// Cache disk size in gigabytes. Omit to use the server default.
	DiskSizeGB param.Opt[int64] `json:"disk_size_gb,omitzero"`
	// Optional non-unique display name
	Name param.Opt[string] `json:"name,omitzero"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r BuilderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BuilderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuilderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuilderListParams struct {
	// Filter builders by tag key-value pairs.
	Tags map[string]string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BuilderListParams]'s query parameters as `url.Values`.
func (r BuilderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
