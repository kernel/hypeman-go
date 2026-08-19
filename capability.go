// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hypeman

import (
	"context"
	"net/http"
	"slices"

	"github.com/kernel/hypeman-go/internal/apijson"
	"github.com/kernel/hypeman-go/internal/requestconfig"
	"github.com/kernel/hypeman-go/option"
	"github.com/kernel/hypeman-go/packages/respjson"
)

// CapabilityService contains methods and other services that help with interacting
// with the hypeman API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCapabilityService] method instead.
type CapabilityService struct {
	Options []option.RequestOption
}

// NewCapabilityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCapabilityService(opts ...option.RequestOption) (r CapabilityService) {
	r = CapabilityService{}
	r.Options = opts
	return
}

// Returns machine-readable host capabilities: server and API version, host
// OS/architecture, every runtime available on this host with its per-runtime
// feature IDs, the configured default runtime and whether it is available, guest
// networking model and host gateway, supported image platforms, and stable
// server-level feature IDs.
//
// Runtime-derived values reflect the actual host (for example, snapshot and
// standby support on macOS is gated on the host OS version), so clients can gate
// behavior on capabilities without hard-coding hypervisor knowledge.
func (r *CapabilityService) Get(ctx context.Context, opts ...option.RequestOption) (res *Capabilities, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Capabilities struct {
	DefaultRuntime CapabilitiesDefaultRuntime `json:"default_runtime" api:"required"`
	// Stable server-level feature IDs: API surfaces this server exposes regardless of
	// which runtime backs an instance. Always present: "instances", "images",
	// "builds", "volumes", "ingress", "exec", "logs". Host-conditional: "devices"
	// (device passthrough management, Linux hosts only) and "rosetta-emulation" (Apple
	// Silicon macOS hosts with Rosetta currently installed, per the same availability
	// probe launches enforce). Per-runtime features are reported under each runtimes[]
	// entry.
	Features []string            `json:"features" api:"required"`
	Host     CapabilitiesHost    `json:"host" api:"required"`
	Images   CapabilitiesImages  `json:"images" api:"required"`
	Network  CapabilitiesNetwork `json:"network" api:"required"`
	// Every runtime this server build supports on this host platform, each with its
	// own availability flag and feature IDs. Hosts commonly support several runtimes
	// at once (for example cloud-hypervisor, firecracker, qemu, and qemu-microvm on
	// linux/amd64). A listed runtime is only launchable when its "available" flag is
	// true. Entries are sorted by name.
	Runtimes []CapabilitiesRuntime `json:"runtimes" api:"required"`
	Server   CapabilitiesServer    `json:"server" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultRuntime respjson.Field
		Features       respjson.Field
		Host           respjson.Field
		Images         respjson.Field
		Network        respjson.Field
		Runtimes       respjson.Field
		Server         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Capabilities) RawJSON() string { return r.JSON.raw }
func (r *Capabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesDefaultRuntime struct {
	// Whether the default runtime can launch on this host: it appears in runtimes and
	// its launch prerequisites are met (matches that entry's "available"). When false,
	// launches that rely on the default will fail until the server is reconfigured
	// with an available runtime or the missing prerequisite (for example the QEMU
	// system binary) is installed.
	Available bool `json:"available" api:"required"`
	// Runtime used for launches that do not name one
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesDefaultRuntime) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesDefaultRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesHost struct {
	// Host CPU architecture
	Arch string `json:"arch" api:"required"`
	// Host operating system
	Os string `json:"os" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arch        respjson.Field
		Os          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesHost) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesHost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesImages struct {
	// Image platform selected when a create request omits one
	DefaultPlatform string `json:"default_platform" api:"required"`
	// Image platforms (os/arch) this host can run. On Apple Silicon macOS this
	// includes linux/amd64 only when Rosetta is currently installed — probed via the
	// same Virtualization.framework availability check launches enforce — so a listed
	// platform is launchable right now. Install Rosetta (softwareupdate
	// --install-rosetta) to enable it.
	Platforms []string `json:"platforms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultPlatform respjson.Field
		Platforms       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesImages) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesNetwork struct {
	// Whether direct VM-to-VM traffic is permitted on the default network
	GuestToGuest bool `json:"guest_to_guest" api:"required"`
	// Guest networking model. "bridge" is a Linux bridge with per-VM TAP devices;
	// "nat" is hypervisor-provided NAT (macOS).
	//
	// Any of "bridge", "nat".
	Model CapabilitiesNetworkModel `json:"model" api:"required"`
	// Guest-visible host gateway IP. Guests reach host services (including host
	// ingress) through this address. Omitted when no default network has been resolved
	// on this host yet.
	Gateway string `json:"gateway"`
	// Guest subnet CIDR
	Subnet string `json:"subnet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GuestToGuest respjson.Field
		Model        respjson.Field
		Gateway      respjson.Field
		Subnet       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesNetwork) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Guest networking model. "bridge" is a Linux bridge with per-VM TAP devices;
// "nat" is hypervisor-provided NAT (macOS).
type CapabilitiesNetworkModel string

const (
	CapabilitiesNetworkModelBridge CapabilitiesNetworkModel = "bridge"
	CapabilitiesNetworkModelNat    CapabilitiesNetworkModel = "nat"
)

type CapabilitiesRuntime struct {
	// Whether this runtime's launch prerequisites are currently met on this host.
	// Listed runtimes are supported by this server build on this platform;
	// available=false means a host prerequisite is missing (for example qemu requires
	// a runnable system-installed QEMU binary and the host vhost-vsock device) and
	// launches naming this runtime will fail until it is installed.
	Available bool `json:"available" api:"required"`
	// Stable feature IDs supported by this runtime on this host: "snapshots"
	// (snapshot/restore), "standby" (pause + memory snapshot, with later restore),
	// "fork" (clone an instance from a stopped source; forking a standby or running
	// source restores/creates snapshots and additionally requires "standby"), "pause"
	// (pause/resume), "hotplug-memory" (live memory resize), "balloon-control"
	// (runtime balloon target changes), "vsock" (guest vsock communication),
	// "gpu-passthrough" (GPU/PCI device passthrough), "disk-io-limit" (disk I/O rate
	// limiting), "disk-resize" (live disk resize). Values are host- and
	// configuration-truthful: vz omits snapshots and standby on macOS 13, which lacks
	// Virtualization.framework VM save/restore, while still advertising fork
	// (stopped-source clones need no save/restore there), and cloud-hypervisor reports
	// "disk-resize" only when the configured default version supports it.
	Features []string `json:"features" api:"required"`
	// Runtime identifier
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		Features    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesRuntime) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesServer struct {
	// API contract version (matches the OpenAPI document info version)
	APIVersion string `json:"api_version" api:"required"`
	// Server build version (short git revision, with "-dirty" suffix for uncommitted
	// builds, or "unknown")
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIVersion  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesServer) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
