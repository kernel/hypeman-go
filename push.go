// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/kernel/hypeman-go/internal/apijson"
	shimjson "github.com/kernel/hypeman-go/internal/encoding/json"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
	"github.com/kernel/hypeman-go/packages/param"
	"github.com/kernel/hypeman-go/packages/respjson"
)

// PushService contains methods and other services that help with interacting with
// the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPushService] method instead.
type PushService struct {
	Options []option.RequestOption
}

// NewPushService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPushService(opts ...option.RequestOption) (r PushService) {
	r = PushService{}
	r.Options = opts
	return
}

// Creates a push job that exports a hypeman image from the local OCI cache to a
// remote registry (e.g. AWS ECR, Docker Hub). Only images in the ready state can
// be pushed.
func (r *PushService) New(ctx context.Context, body PushNewParams, opts ...option.RequestOption) (res *Push, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pushes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists outbound image push jobs, newest first.
func (r *PushService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Push, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pushes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get push details
func (r *PushService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Push, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("pushes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The properties Image, Target are required.
type CreatePushRequestParam struct {
	// Hypeman image name to push (tag or digest form)
	Image string `json:"image" api:"required"`
	// Full remote reference to push to
	Target string `json:"target" api:"required"`
	// Allow pushing to plain-HTTP registries
	Insecure param.Opt[bool] `json:"insecure,omitzero"`
	// Docker-style registry credentials borrowed for one image pull or push request.
	// They remain in memory and are never persisted or logged. When omitted or empty,
	// the server's own registry credentials are used. An interrupted credentialed
	// operation must be retried with fresh credentials.
	Credentials PushCredentialsParam `json:"credentials,omitzero"`
	paramObj
}

func (r CreatePushRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePushRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePushRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Push struct {
	// Push job identifier
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cached manifest digest being pushed
	Digest string `json:"digest" api:"required"`
	// Hypeman image name (normalized ref)
	Image string `json:"image" api:"required"`
	// Any of "queued", "pushing", "pushed", "failed".
	Status PushStatus `json:"status" api:"required"`
	// Remote reference the image is pushed to
	Target string `json:"target" api:"required"`
	// Total compressed layer bytes pushed
	Bytes       int64     `json:"bytes"`
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error message
	Error string `json:"error" api:"nullable"`
	// Number of layers pushed
	Layers int64 `json:"layers"`
	// Position in the push queue
	QueuePosition int64 `json:"queue_position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Digest        respjson.Field
		Image         respjson.Field
		Status        respjson.Field
		Target        respjson.Field
		Bytes         respjson.Field
		CompletedAt   respjson.Field
		Error         respjson.Field
		Layers        respjson.Field
		QueuePosition respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Push) RawJSON() string { return r.JSON.raw }
func (r *Push) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Docker-style registry credentials borrowed for one image pull or push request.
// They remain in memory and are never persisted or logged. When omitted or empty,
// the server's own registry credentials are used. An interrupted credentialed
// operation must be retried with fresh credentials.
type PushCredentialsParam struct {
	// Registry password or access token
	Password param.Opt[string] `json:"password,omitzero" format:"password"`
	// Bearer token for an Authorization header
	RegistryToken param.Opt[string] `json:"registry_token,omitzero" format:"password"`
	// Registry username
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r PushCredentialsParam) MarshalJSON() (data []byte, err error) {
	type shadow PushCredentialsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PushCredentialsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PushStatus string

const (
	PushStatusQueued  PushStatus = "queued"
	PushStatusPushing PushStatus = "pushing"
	PushStatusPushed  PushStatus = "pushed"
	PushStatusFailed  PushStatus = "failed"
)

type PushNewParams struct {
	CreatePushRequest CreatePushRequestParam
	paramObj
}

func (r PushNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePushRequest)
}
func (r *PushNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
