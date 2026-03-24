# Shared Params Types

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go/shared#SnapshotCompressionConfigParam">SnapshotCompressionConfigParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go/shared#SnapshotCompressionConfig">SnapshotCompressionConfig</a>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Image">Image</a>

Methods:

- <code title="post /images">client.Images.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageNewParams">ImageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /images">client.Images.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageListParams">ImageListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /images/{name}">client.Images.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /images/{name}">client.Images.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Instances

Params Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SetSnapshotScheduleRequestParam">SetSnapshotScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotPolicyParam">SnapshotPolicyParam</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotScheduleRetentionParam">SnapshotScheduleRetentionParam</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeMountParam">VolumeMountParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceStats">InstanceStats</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#PathInfo">PathInfo</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotPolicy">SnapshotPolicy</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotSchedule">SnapshotSchedule</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotScheduleRetention">SnapshotScheduleRetention</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeMount">VolumeMount</a>

Methods:

- <code title="post /instances">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceNewParams">InstanceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /instances/{id}">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceUpdateParams">InstanceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /instances">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceListParams">InstanceListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /instances/{id}">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /instances/{id}/fork">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Fork">Fork</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceForkParams">InstanceForkParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /instances/{id}">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /instances/{id}/logs">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Logs">Logs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceLogsParams">InstanceLogsParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /instances/{id}/restore">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Restore">Restore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /instances/{id}/standby">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Standby">Standby</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceStandbyParams">InstanceStandbyParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /instances/{id}/start">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceStartParams">InstanceStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /instances/{id}/stat">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Stat">Stat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceStatParams">InstanceStatParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#PathInfo">PathInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /instances/{id}/stats">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Stats">Stats</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceStats">InstanceStats</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /instances/{id}/stop">client.Instances.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Methods:

- <code title="post /instances/{id}/volumes/{volumeId}">client.Instances.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceVolumeService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceVolumeAttachParams">InstanceVolumeAttachParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /instances/{id}/volumes/{volumeId}">client.Instances.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceVolumeService.Detach">Detach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceVolumeDetachParams">InstanceVolumeDetachParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Snapshots

Methods:

- <code title="post /instances/{id}/snapshots">client.Instances.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotNewParams">InstanceSnapshotNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Snapshot">Snapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /instances/{id}/snapshots/{snapshotId}/restore">client.Instances.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotService.Restore">Restore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, snapshotID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotRestoreParams">InstanceSnapshotRestoreParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SnapshotSchedule

Methods:

- <code title="put /instances/{id}/snapshot-schedule">client.Instances.SnapshotSchedule.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotScheduleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotScheduleUpdateParams">InstanceSnapshotScheduleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotSchedule">SnapshotSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /instances/{id}/snapshot-schedule">client.Instances.SnapshotSchedule.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /instances/{id}/snapshot-schedule">client.Instances.SnapshotSchedule.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#InstanceSnapshotScheduleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotSchedule">SnapshotSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Snapshots

Params Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotKind">SnapshotKind</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Snapshot">Snapshot</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotKind">SnapshotKind</a>

Methods:

- <code title="get /snapshots">client.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotListParams">SnapshotListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Snapshot">Snapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /snapshots/{snapshotId}">client.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, snapshotID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /snapshots/{snapshotId}/fork">client.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotService.Fork">Fork</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, snapshotID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotForkParams">SnapshotForkParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Instance">Instance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /snapshots/{snapshotId}">client.Snapshots.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#SnapshotService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, snapshotID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Snapshot">Snapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Volumes

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeAttachment">VolumeAttachment</a>

Methods:

- <code title="post /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeNewParams">VolumeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeListParams">VolumeListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volumes/{id}">client.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /volumes/from-archive">client.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeService.NewFromArchive">NewFromArchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeNewFromArchiveParams">VolumeNewFromArchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes/{id}">client.Volumes.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#AvailableDevice">AvailableDevice</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Device">Device</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceType">DeviceType</a>

Methods:

- <code title="post /devices">client.Devices.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceNewParams">DeviceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /devices/{id}">client.Devices.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /devices">client.Devices.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceListParams">DeviceListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /devices/{id}">client.Devices.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /devices/available">client.Devices.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DeviceService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#AvailableDevice">AvailableDevice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ingresses

Params Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressMatchParam">IngressMatchParam</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressRuleParam">IngressRuleParam</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressTargetParam">IngressTargetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Ingress">Ingress</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressMatch">IngressMatch</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressRule">IngressRule</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressTarget">IngressTarget</a>

Methods:

- <code title="post /ingresses">client.Ingresses.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressNewParams">IngressNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Ingress">Ingress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ingresses">client.Ingresses.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressListParams">IngressListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Ingress">Ingress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /ingresses/{id}">client.Ingresses.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /ingresses/{id}">client.Ingresses.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#IngressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Ingress">Ingress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Resources

Params Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#MemoryReclaimRequestParam">MemoryReclaimRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#DiskBreakdown">DiskBreakdown</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#GPUProfile">GPUProfile</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#GPUResourceStatus">GPUResourceStatus</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#MemoryReclaimAction">MemoryReclaimAction</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#MemoryReclaimResponse">MemoryReclaimResponse</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#PassthroughDevice">PassthroughDevice</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ResourceAllocation">ResourceAllocation</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ResourceStatus">ResourceStatus</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Resources">Resources</a>

Methods:

- <code title="get /resources">client.Resources.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ResourceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Resources">Resources</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /resources/memory/reclaim">client.Resources.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ResourceService.ReclaimMemory">ReclaimMemory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#ResourceReclaimMemoryParams">ResourceReclaimMemoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#MemoryReclaimResponse">MemoryReclaimResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Build">Build</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildEvent">BuildEvent</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildProvenance">BuildProvenance</a>
- <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildStatus">BuildStatus</a>

Methods:

- <code title="post /builds">client.Builds.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildNewParams">BuildNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /builds">client.Builds.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildListParams">BuildListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /builds/{id}">client.Builds.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /builds/{id}/events">client.Builds.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildService.Events">Events</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildEventsParams">BuildEventsParams</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildEvent">BuildEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /builds/{id}">client.Builds.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/kernel/hypeman-go">hypeman</a>.<a href="https://pkg.go.dev/github.com/kernel/hypeman-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
