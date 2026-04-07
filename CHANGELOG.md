# Changelog

## 0.17.0 (2026-04-07)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/kernel/hypeman-go/compare/v0.16.0...v0.17.0)

### Features

* Add Linux auto-standby controller and E2E coverage ([211a8a5](https://github.com/kernel/hypeman-go/commit/211a8a5c99cd759999d46d69f290717fb6eeabf1))
* Add scheduled instance snapshots with retention cleanup ([f0d4d52](https://github.com/kernel/hypeman-go/commit/f0d4d52bf563ca828cba5f31f6a0c913d4dba4f8))
* Add waitForState endpoint for blocking state transitions ([d6cea17](https://github.com/kernel/hypeman-go/commit/d6cea176033845f13c0dcbe79c0a8150c5a87d94))
* **internal:** support comma format in multipart form encoding ([e6abba9](https://github.com/kernel/hypeman-go/commit/e6abba95be9eebf3a437b3ca371cd5f30795da02))


### Bug Fixes

* fix issue with unmarshaling in some cases ([1d74ffd](https://github.com/kernel/hypeman-go/commit/1d74ffdcf803144714c530ed72e6d65fba2a9b54))
* prevent duplicate ? in query params ([6bc4417](https://github.com/kernel/hypeman-go/commit/6bc44177b7bd0d7b03702664d573cc160cc11a75))


### Chores

* **ci:** skip lint on metadata-only changes ([e3c8b1f](https://github.com/kernel/hypeman-go/commit/e3c8b1f0943858b91f196db4889a27f0d626a6e4))
* **ci:** support opting out of skipping builds on metadata-only commits ([c5320fd](https://github.com/kernel/hypeman-go/commit/c5320fdaa5ad7a7e223590fd2664ee4fdb35fa6f))
* **client:** fix multipart serialisation of Default() fields ([0abcf13](https://github.com/kernel/hypeman-go/commit/0abcf13ab96198afe17152d6653309ad5aecbb97))
* **internal:** support default value struct tag ([e9226b4](https://github.com/kernel/hypeman-go/commit/e9226b48b553d83aad20295d320ca976f97dc0b1))
* **internal:** update gitignore ([3dbe64e](https://github.com/kernel/hypeman-go/commit/3dbe64e2bdb18744c58d30a8a9df44107fed6a61))
* remove unnecessary error check for url parsing ([cd07c8f](https://github.com/kernel/hypeman-go/commit/cd07c8ffd89e1fc31ca8769be974d6830f38cb07))
* update docs for api:"required" ([45f88d9](https://github.com/kernel/hypeman-go/commit/45f88d9c8507e0c6aa729d52a69d58041d0dd3db))

## 0.16.0 (2026-03-23)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/kernel/hypeman-go/compare/v0.15.0...v0.16.0)

### Features

* add active ballooning reclaim controller ([b8ecb54](https://github.com/kernel/hypeman-go/commit/b8ecb541b8237d1283a05422eed6e30ff2b0536f))
* Add always-on /metrics endpoint with dual pull/push telemetry ([0b3751a](https://github.com/kernel/hypeman-go/commit/0b3751a37068dca3771ad04fc1af876c5e52384f))
* Add optional snapshot compression defaults and standby integration ([b6d9ab3](https://github.com/kernel/hypeman-go/commit/b6d9ab3d0b4ffdd7d56153464387ed9e603d8d0a))
* add optional VM egress MITM proxy with mock-secret header rewriting ([e8b721c](https://github.com/kernel/hypeman-go/commit/e8b721cbc709af739827929839e443981d636233))
* Add strict metadata tags across mutable resources ([8b5543e](https://github.com/kernel/hypeman-go/commit/8b5543e6f8fed7062d28c37fa35ba031e14c6f79))
* Rename tag fields from metadata to tags ([2f8e29e](https://github.com/kernel/hypeman-go/commit/2f8e29e033b44a8a492ec9a3bdcbc44cabf2eb11))
* Snapshot ([c4a0fbb](https://github.com/kernel/hypeman-go/commit/c4a0fbb34ee2b325229f6a02ae9cc3c57346788d))
* support updating egress proxy secret envs for key rotation ([96b3209](https://github.com/kernel/hypeman-go/commit/96b32092d707499c6a962d89a3530b413be4ccab))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([0d9d654](https://github.com/kernel/hypeman-go/commit/0d9d654f8074387de55f5befe709477c6905deaa))
* **internal:** codegen related update ([e6a6702](https://github.com/kernel/hypeman-go/commit/e6a6702ba1fcf355cf64a60fd016393e5474e9bc))
* **internal:** minor cleanup ([1917d6d](https://github.com/kernel/hypeman-go/commit/1917d6de179a6a36358598cf3b1f6d4ebbe7f034))
* **internal:** tweak CI branches ([6dc7e68](https://github.com/kernel/hypeman-go/commit/6dc7e68725e65b75bd7c2aa07d4427909ad2365c))
* **internal:** use explicit returns ([923db74](https://github.com/kernel/hypeman-go/commit/923db74ba5750233d429a5e98dc6cb5fb0e806ba))
* **internal:** use explicit returns in more places ([16131b8](https://github.com/kernel/hypeman-go/commit/16131b802aa3f5a2b32d298d179966a440f53af5))
* update placeholder string ([bea84ac](https://github.com/kernel/hypeman-go/commit/bea84ac19f1178f378cc6ba8bad07a26848b5492))

## 0.15.0 (2026-03-04)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/kernel/hypeman-go/compare/v0.14.0...v0.15.0)

### Features

* Add fork operation to stainless config ([5ab52b7](https://github.com/kernel/hypeman-go/commit/5ab52b7c74031eb7e874615aa42c00090cb00b51))


### Chores

* **internal:** codegen related update ([83363a5](https://github.com/kernel/hypeman-go/commit/83363a5275d88e131d0736516fea215faaab84a7))

## 0.14.0 (2026-03-02)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/kernel/hypeman-go/compare/v0.13.0...v0.14.0)

### Features

* Add cloud hypervisor VM forking helpers ([bfac548](https://github.com/kernel/hypeman-go/commit/bfac548bac8dab6cbe47dd39f0f8cc7fd53b0339))
* add Firecracker hypervisor support ([251f4b9](https://github.com/kernel/hypeman-go/commit/251f4b953c665f32305f393196974c8cb4fe7ec3))


### Chores

* **stainless:** update SDKs for PR [#116](https://github.com/kernel/hypeman-go/issues/116) ([560f851](https://github.com/kernel/hypeman-go/commit/560f8512392ffda9d3f478494b593bda363b1c86))

## 0.13.0 (2026-02-26)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/kernel/hypeman-go/compare/v0.12.0...v0.13.0)

### Features

* wire up memory_mb and cpus in builds API ([ed09f5a](https://github.com/kernel/hypeman-go/commit/ed09f5a0bc3c60e7ccf6da1101d23311107c7118))

## 0.12.0 (2026-02-26)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/kernel/hypeman-go/compare/v0.11.0...v0.12.0)

### Features

* add metadata and state filtering to GET /instances ([8149c9f](https://github.com/kernel/hypeman-go/commit/8149c9fe5e9b36aa5709767e7f3986d6778bd432))
* Disable default hotplug memory allocation ([4c65d5c](https://github.com/kernel/hypeman-go/commit/4c65d5c271ac3a620da549b47ece553e1860aaf6))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([daa2281](https://github.com/kernel/hypeman-go/commit/daa2281e6e9833ae4dba2b9b6870014ceb0b2fff))
* send query params for NewFromArchive ([a8c45a6](https://github.com/kernel/hypeman-go/commit/a8c45a69e83c96137c772d999a115972f0e6a003))


### Chores

* **internal:** move custom custom `json` tags to `api` ([d04f6ed](https://github.com/kernel/hypeman-go/commit/d04f6ed70c95357f323aa4e76b3a6ad8ebd12ec3))
* **internal:** remove mock server code ([b511676](https://github.com/kernel/hypeman-go/commit/b51167627fc0cd0f947633cb8694f4ee0756c268))
* update mock server docs ([d2ae478](https://github.com/kernel/hypeman-go/commit/d2ae478d46fb67550a3b35d4261218f9368709f0))

## 0.11.0 (2026-02-15)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/kernel/hypeman-go/compare/v0.10.0...v0.11.0)

### Features

* Add image_name parameter to builds ([36ea383](https://github.com/kernel/hypeman-go/commit/36ea383988ce81b907a893220bbc983e8143a1d2))

## 0.10.0 (2026-02-13)

Full Changelog: [v0.9.8...v0.10.0](https://github.com/kernel/hypeman-go/compare/v0.9.8...v0.10.0)

### Features

* Add metadata field to instances ([8ce4014](https://github.com/kernel/hypeman-go/commit/8ce40145c0f6db8e928efa9c2909521a6d452579))
* Better stop behavior ([12cfb4e](https://github.com/kernel/hypeman-go/commit/12cfb4e4ddfc198c65b09efeb73c8db0fd609f5f))

## 0.9.8 (2026-02-11)

Full Changelog: [v0.9.7...v0.9.8](https://github.com/kernel/hypeman-go/compare/v0.9.7...v0.9.8)

## 0.9.7 (2026-02-11)

Full Changelog: [v0.9.6...v0.9.7](https://github.com/kernel/hypeman-go/compare/v0.9.6...v0.9.7)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([e693834](https://github.com/kernel/hypeman-go/commit/e693834704b3541d4a5f260b547026bae8a19b1b))


### Refactors

* cross-platform foundation for macOS support ([8adc4f3](https://github.com/kernel/hypeman-go/commit/8adc4f38026abee34ad85c15509e90f47644a0d0))

## 0.9.6 (2026-01-30)

Full Changelog: [v0.9.0...v0.9.6](https://github.com/kernel/hypeman-go/compare/v0.9.0...v0.9.6)

### Features

* add boot time optimizations for faster VM startup ([3992761](https://github.com/kernel/hypeman-go/commit/3992761e3ad8ebb0cc22fb7408199b068e9d8013))
* Add to stainless config new API endpoints ([de008e8](https://github.com/kernel/hypeman-go/commit/de008e89fadbaedde6554181618fb03c71b49465))
* **api:** manual updates ([f60e600](https://github.com/kernel/hypeman-go/commit/f60e60015bb9ce18c7083963d9ecd11c980de495))
* **builds:** implement two-tier build cache with per-repo token scopes ([0e29d03](https://github.com/kernel/hypeman-go/commit/0e29d03d94cf50a0d0e83c323f7ed9f2e15f3e61))
* **client:** add a convenient param.SetJSON helper ([7fea166](https://github.com/kernel/hypeman-go/commit/7fea1660f3d17d8a35f5d2f6aa352b553785624b))
* Use resources module for input validation ([af678e8](https://github.com/kernel/hypeman-go/commit/af678e8c794307a6bd47476acff3ca42a7a52546))

## 0.9.0 (2026-01-05)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/kernel/hypeman-go/compare/v0.8.0...v0.9.0)

### Features

* QEMU support ([d708091](https://github.com/kernel/hypeman-go/commit/d70809169d136df3f1efbf961f2a90084e1f9fa5))
* Resource accounting ([4141287](https://github.com/kernel/hypeman-go/commit/414128770e8137ed2a40d404f0f4ac06ea1a0731))

## 0.8.0 (2025-12-23)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/kernel/hypeman-go/compare/v0.7.0...v0.8.0)

### Features

* add hypeman cp for file copy to/from running VMs ([49ea898](https://github.com/kernel/hypeman-go/commit/49ea89852eed5e0893febc4c68d295a0d1a8bfe5))
* **encoder:** support bracket encoding form-data object members ([8ab31e8](https://github.com/kernel/hypeman-go/commit/8ab31e89c70baa967842c1c160d0b49db44b089a))
* gpu passthrough ([067a01b](https://github.com/kernel/hypeman-go/commit/067a01b4ac06e82c2db6b165127144afa18a691d))


### Bug Fixes

* skip usage tests that don't work with Prism ([d62b246](https://github.com/kernel/hypeman-go/commit/d62b2466715247e7d083ab7ef33040e5da036bd8))


### Chores

* add float64 to valid types for RegisterFieldValidator ([b4666fd](https://github.com/kernel/hypeman-go/commit/b4666fd1bfcdd17b0a4d4bf88541670cd40c8b1c))

## 0.7.0 (2025-12-11)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/kernel/hypeman-go/compare/v0.6.0...v0.7.0)

### Features

* Operational logs over API: hypeman.log, vmm.log ([ec614f5](https://github.com/kernel/hypeman-go/commit/ec614f5bdc0e110f31cec905d6deb7f1d460305b))
* Support TLS for ingress ([973a5d8](https://github.com/kernel/hypeman-go/commit/973a5d8b65601e70801ed4570f76980d01c92198))


### Bug Fixes

* incorrect reporting of Stopped, add better error reporting ([dc27cbd](https://github.com/kernel/hypeman-go/commit/dc27cbdc7985c1db74b19501f1eb7a5da6442041))

## 0.6.0 (2025-12-06)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/kernel/hypeman-go/compare/v0.5.0...v0.6.0)

### Features

* Start and Stop VM ([b992228](https://github.com/kernel/hypeman-go/commit/b99222818b197010ba324c2e2477047e5bf13802))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([0d87152](https://github.com/kernel/hypeman-go/commit/0d8715273698dab9bb6c276352a13605ddd272a5))
* rename param to avoid collision ([f1ec9d5](https://github.com/kernel/hypeman-go/commit/f1ec9d52e3f5f6c8398bdded04a4ed9cfbd8151b))


### Chores

* elide duplicate aliases ([9be276f](https://github.com/kernel/hypeman-go/commit/9be276faa6d683ddffe3a21c969b44f13acface0))
* **internal:** codegen related update ([f3de06d](https://github.com/kernel/hypeman-go/commit/f3de06d220faf866b70829862cd1b76ee4e8fbf8))

## 0.5.0 (2025-12-05)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/kernel/hypeman-go/compare/v0.4.0...v0.5.0)

### Features

* add Push and PushImage functions for OCI registry push ([7417cc8](https://github.com/kernel/hypeman-go/commit/7417cc8a56c7d11c535ac7ab9a7b3d21d80bd2b4))
* Ingress ([c751d1a](https://github.com/kernel/hypeman-go/commit/c751d1a6bba5ca619c03f833f27251c6d3b855a7))
* Initialize volume with data ([32d4047](https://github.com/kernel/hypeman-go/commit/32d404746df0a3e9d83e7651105e6c6daa16476f))
* try to fix name collision in codegen ([8173a73](https://github.com/kernel/hypeman-go/commit/8173a73d0317d35870d5a3cec8f3fdec56fcf362))
* Volume readonly multi-attach ([bac3fd2](https://github.com/kernel/hypeman-go/commit/bac3fd2cee3325dc3d1b31e6077ad1f1ce13340c))
* Volumes ([099f9b8](https://github.com/kernel/hypeman-go/commit/099f9b8a2553087e117c8c8a9731900081d713f0))

## 0.4.0 (2025-11-26)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/kernel/hypeman-go/compare/v0.3.0...v0.4.0)

### Features

* Generate log streaming ([f444c22](https://github.com/kernel/hypeman-go/commit/f444c22bd9eb0ad06e66b3ca167171ddec2836e4))

## 0.3.0 (2025-11-26)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/kernel/hypeman-go/compare/v0.2.0...v0.3.0)

### Features

* Remove exec from openapi spec ([ee8d1bb](https://github.com/kernel/hypeman-go/commit/ee8d1bb586a130c0b6629603ca4edb489f671889))

## 0.2.0 (2025-11-26)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/kernel/hypeman-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** add exec ([f3992ff](https://github.com/kernel/hypeman-go/commit/f3992ffe807e7006a25ae2211cd5cb25fb599bff))

## 0.1.0 (2025-11-26)

Full Changelog: [v0.0.3...v0.1.0](https://github.com/kernel/hypeman-go/compare/v0.0.3...v0.1.0)

### Features

* Network manager ([7864aba](https://github.com/kernel/hypeman-go/commit/7864abadad29bcfbb61d2c35a7135ef2407d6c47))

## 0.0.3 (2025-11-19)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/kernel/hypeman-go/compare/v0.0.2...v0.0.3)

### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([ac1a646](https://github.com/kernel/hypeman-go/commit/ac1a64697c333aecdc6a463fe760b99635ba8b72))

## 0.0.2 (2025-11-11)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/kernel/hypeman-go/compare/v0.0.1...v0.0.2)

### Chores

* update SDK settings ([ecdeb35](https://github.com/kernel/hypeman-go/commit/ecdeb354a1d6a82a1d2afc1742ca02b25eb3218f))
