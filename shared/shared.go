// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/kernel/hypeman-go/internal/apijson"
	"github.com/kernel/hypeman-go/packages/param"
	"github.com/kernel/hypeman-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type SnapshotCompressionConfig struct {
	// Enable snapshot memory compression
	Enabled bool `json:"enabled" api:"required"`
	// Compression algorithm (defaults to zstd when enabled). Ignored when enabled is
	// false.
	//
	// Any of "zstd", "lz4".
	Algorithm SnapshotCompressionConfigAlgorithm `json:"algorithm"`
	// Compression level. Allowed ranges are zstd=1-19 and lz4=0-9. When omitted, zstd
	// defaults to 1 and lz4 defaults to 0. Ignored when enabled is false.
	Level int64 `json:"level"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Algorithm   respjson.Field
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotCompressionConfig) RawJSON() string { return r.JSON.raw }
func (r *SnapshotCompressionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SnapshotCompressionConfig to a
// SnapshotCompressionConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SnapshotCompressionConfigParam.Overrides()
func (r SnapshotCompressionConfig) ToParam() SnapshotCompressionConfigParam {
	return param.Override[SnapshotCompressionConfigParam](json.RawMessage(r.RawJSON()))
}

// Compression algorithm (defaults to zstd when enabled). Ignored when enabled is
// false.
type SnapshotCompressionConfigAlgorithm string

const (
	SnapshotCompressionConfigAlgorithmZstd SnapshotCompressionConfigAlgorithm = "zstd"
	SnapshotCompressionConfigAlgorithmLz4  SnapshotCompressionConfigAlgorithm = "lz4"
)

// The property Enabled is required.
type SnapshotCompressionConfigParam struct {
	// Enable snapshot memory compression
	Enabled bool `json:"enabled" api:"required"`
	// Compression level. Allowed ranges are zstd=1-19 and lz4=0-9. When omitted, zstd
	// defaults to 1 and lz4 defaults to 0. Ignored when enabled is false.
	Level param.Opt[int64] `json:"level,omitzero"`
	// Compression algorithm (defaults to zstd when enabled). Ignored when enabled is
	// false.
	//
	// Any of "zstd", "lz4".
	Algorithm SnapshotCompressionConfigAlgorithm `json:"algorithm,omitzero"`
	paramObj
}

func (r SnapshotCompressionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SnapshotCompressionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SnapshotCompressionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
