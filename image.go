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
	shimjson "github.com/kernel/hypeman-go/internal/encoding/json"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
	"github.com/kernel/hypeman-go/packages/param"
	"github.com/kernel/hypeman-go/packages/respjson"
)

// ImageService contains methods and other services that help with interacting with
// the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r ImageService) {
	r = ImageService{}
	r.Options = opts
	return
}

// Pull and convert OCI image
func (r *ImageService) New(ctx context.Context, body ImageNewParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List images
func (r *ImageService) List(ctx context.Context, query ImageListParams, opts ...option.RequestOption) (res *[]Image, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete image
func (r *ImageService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return err
	}
	path := fmt.Sprintf("images/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get image details
func (r *ImageService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create or update a local image tag
func (r *ImageService) Tag(ctx context.Context, name string, body ImageTagParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/tag", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Image struct {
	// Creation timestamp (RFC3339)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resolved manifest digest
	Digest string `json:"digest" api:"required"`
	// Normalized OCI image reference (tag or digest)
	Name string `json:"name" api:"required"`
	// Build status
	//
	// Any of "pending", "pulling", "converting", "ready", "failed".
	Status ImageStatus `json:"status" api:"required"`
	// CMD from container metadata
	Cmd []string `json:"cmd" api:"nullable"`
	// Entrypoint from container metadata
	Entrypoint []string `json:"entrypoint" api:"nullable"`
	// Environment variables from container metadata
	Env map[string]string `json:"env"`
	// Error message if status is failed
	Error string `json:"error" api:"nullable"`
	// Resolved image platform as os/arch[/variant] (e.g. "linux/amd64")
	Platform string `json:"platform"`
	// Position in build queue (null if not queued)
	QueuePosition int64 `json:"queue_position" api:"nullable"`
	// Disk size in bytes (null until ready)
	SizeBytes int64 `json:"size_bytes" api:"nullable"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags"`
	// Working directory from container metadata
	WorkingDir string `json:"working_dir" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		Digest        respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		Cmd           respjson.Field
		Entrypoint    respjson.Field
		Env           respjson.Field
		Error         respjson.Field
		Platform      respjson.Field
		QueuePosition respjson.Field
		SizeBytes     respjson.Field
		Tags          respjson.Field
		WorkingDir    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Build status
type ImageStatus string

const (
	ImageStatusPending    ImageStatus = "pending"
	ImageStatusPulling    ImageStatus = "pulling"
	ImageStatusConverting ImageStatus = "converting"
	ImageStatusReady      ImageStatus = "ready"
	ImageStatusFailed     ImageStatus = "failed"
)

// The property Target is required.
type TagImageRequestParam struct {
	// Target OCI image reference with a tag. The local tag points to the source image
	// without pulling it again.
	Target string `json:"target" api:"required"`
	paramObj
}

func (r TagImageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TagImageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagImageRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageNewParams struct {
	// OCI image reference (e.g., docker.io/library/nginx:latest)
	Name string `json:"name" api:"required"`
	// Target platform as os/arch[/variant] (e.g. "linux/amd64"), matching Docker
	// --platform. Omit for the host platform. Not a fixed enum: the os/arch[/variant]
	// grammar is validated server-side and invalid values return 400 invalid_platform.
	// Only os "linux" with arch amd64 or arm64 is accepted today.
	Platform param.Opt[string] `json:"platform,omitzero"`
	// Docker-style registry credentials borrowed for one image pull or push request.
	// They remain in memory and are never persisted or logged. When omitted or empty,
	// the server's own registry credentials are used. An interrupted credentialed
	// operation must be retried with fresh credentials.
	Credentials PushCredentialsParam `json:"credentials,omitzero"`
	// User-defined key-value tags.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r ImageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ImageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageListParams struct {
	// Filter images by tag key-value pairs.
	Tags map[string]string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageListParams]'s query parameters as `url.Values`.
func (r ImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageTagParams struct {
	TagImageRequest TagImageRequestParam
	paramObj
}

func (r ImageTagParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TagImageRequest)
}
func (r *ImageTagParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
