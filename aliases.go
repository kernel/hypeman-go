// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"github.com/kernel/hypeman-go/internal/apierror"
	"github.com/kernel/hypeman-go/packages/param"
	"github.com/kernel/hypeman-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type SnapshotCompressionConfig = shared.SnapshotCompressionConfig

// Compression algorithm (defaults to zstd when enabled). Ignored when enabled is
// false.
//
// This is an alias to an internal type.
type SnapshotCompressionConfigAlgorithm = shared.SnapshotCompressionConfigAlgorithm

// Equals "zstd"
const SnapshotCompressionConfigAlgorithmZstd = shared.SnapshotCompressionConfigAlgorithmZstd

// Equals "lz4"
const SnapshotCompressionConfigAlgorithmLz4 = shared.SnapshotCompressionConfigAlgorithmLz4

// This is an alias to an internal type.
type SnapshotCompressionConfigParam = shared.SnapshotCompressionConfigParam
