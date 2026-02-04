# OCI Runtime Spec Editor
## OCI 运行时规范编辑器

A fluent, chainable Go library for creating and modifying [OCI runtime-spec](https://github.com/opencontainers/runtime-spec) configuration files (`config.json`).  
一个流畅、可链式调用的 Go 库，用于创建和修改 [OCI 运行时规范](https://github.com/opencontainers/runtime-spec) 配置文件（`config.json`）。

---

## Features / 特性

- **Fluent API**: Chain method calls for intuitive configuration.  
  **流畅 API**：通过链式调用直观地构建配置。
- **Create from scratch**: Generate a complete `config.json` programmatically.  
  **从零生成**：以编程方式生成完整的 `config.json`。
- **Modify existing specs**: Load and update an existing `config.json`.  
  **修改现有配置**：加载并更新已有的 `config.json`。
- **Extract sub-sections**: Retrieve specific parts (e.g., `linux`, `process`) as structs.  
  **提取子配置块**：获取指定部分（如 `linux`、`process`）为结构体。
- **Pretty-printed output**: Built-in support for formatted JSON.  
  **美化输出**：内置支持格式化 JSON 输出。
- **Error handling**: Short-circuit on error with `.Err` field (like GORM).  
  **错误处理**：出错时自动短路，通过 `.Err` 字段统一检查（类似 GORM）。

---

## Usage / 使用方法

### 1. Generate a new `config.json` / 生成新配置

```go
func main() {
    s := editor.New()

    var result string
    if err := s.SetVersion("2.0.0").
        SetProcessTerminal(true).
        SetProcessUserUid(0).
        SetProcessUserGid(0).
        SetProcessArgs([]string{"sh", "-c", "echo hello world"}).
        SetProcessEnv([]string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"}).
        SetProcessCwd("/").
        SetProcessCapabilitiesBounding([]string{"CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"}).
        SetProcessCapabilitiesEffective([]string{"CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"}).
        SetProcessCapabilitiesPermitted([]string{"CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"}).
        SetProcessCapabilitiesAmbient([]string{"CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"}).
        SetProcessRlimits([]specs.POSIXRlimit{{Type: "RLIMIT_NOFILE", Hard: 1024, Soft: 1024}}).
        SetProcessNoNewPrivileges(true).
        SetRootPath("/rootfs").SetRootReadonly(true).
        SetHostname("oci-editor").
        AddMount(specs.Mount{Type: "proc", Source: "proc", Destination: "/proc"}).
        AddMount(specs.Mount{
            Type:        "tmpfs",
            Source:      "tmpfs",
            Destination: "/dev",
            Options:     []string{"nosuid", "strictatime", "mode=755", "size=65536k"},
        }).
        AddMount(specs.Mount{
            Type:        "devpts",
            Source:      "devpts",
            Destination: "/dev/pts",
            Options:     []string{"nosuid", "noexec", "newinstance", "ptmxmode=0666", "mode=0620", "gid=5"},
        }).
        AddLinuxResourcesDevice(specs.LinuxDeviceCgroup{Allow: false, Access: "rwm"}).
        AddLinuxNameSpace(specs.PIDNamespace, "").
        AddLinuxNameSpace(specs.IPCNamespace, "").
        AddLinuxNameSpace(specs.UTSNamespace, "").
        AddLinuxNameSpace(specs.MountNamespace, "").
        AddLinuxNameSpace(specs.UserNamespace, "").
        AddLinuxNameSpace(specs.NetworkNamespace, "").
        AddLinuxNameSpace(specs.CgroupNamespace, "").
        SetLinuxMaskedPaths([]string{"/proc/kcore", "/sys/firmware"}).
        SetLinuxReadonlyPaths([]string{"/proc/sys", "/proc/sysrq-trigger", "/proc/irq", "/proc/bus"}).
        Pretty(). // Enable pretty-printed JSON
        Generate(&result).Err; err != nil {
        panic(err)
    }

    println(result)
}
```




```json
{
  "ociVersion": "2.0.0",
  "process": {
    "terminal": true,
    "user": {
      "uid": 0,
      "gid": 0
    },
    "args": ["sh", "-c", "echo hello world"],
    "env": ["PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"],
    "cwd": "/",
    "capabilities": {
      "bounding": ["CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"],
      "effective": ["CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"],
      "permitted": ["CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"],
      "ambient": ["CAP_AUDIT_WRITE", "CAP_KILL", "CAP_NET_BIND_SERVICE"]
    },
    "rlimits": [{"type": "RLIMIT_NOFILE", "hard": 1024, "soft": 1024}],
    "noNewPrivileges": true
  },
  "root": {"path": "/rootfs", "readonly": true},
  "hostname": "oci-editor",
  "mounts": [
    {"destination": "/proc", "type": "proc", "source": "proc"},
    {
      "destination": "/dev",
      "type": "tmpfs",
      "source": "tmpfs",
      "options": ["nosuid", "strictatime", "mode=755", "size=65536k"]
    },
    {
      "destination": "/dev/pts",
      "type": "devpts",
      "source": "devpts",
      "options": ["nosuid", "noexec", "newinstance", "ptmxmode=0666", "mode=0620", "gid=5"]
    }
  ],
  "linux": {
    "resources": {
      "devices": [{"allow": false, "access": "rwm"}]
    },
    "namespaces": [
      {"type": "pid"}, {"type": "ipc"}, {"type": "uts"},
      {"type": "mount"}, {"type": "user"}, {"type": "network"}, {"type": "cgroup"}
    ],
    "maskedPaths": ["/proc/kcore", "/sys/firmware"],
    "readonlyPaths": ["/proc/sys", "/proc/sysrq-trigger", "/proc/irq", "/proc/bus"]
  }
}
```


---

### 2. Modify an existing `config.json` / 修改已有配置

```go
func main() {
    const cfg = `{...}` // Your existing config.json content

    s := editor.New()
    var result string

    if err := s.WithContent([]byte(cfg)).
        SetRootPath("/roooooooooooootfs").
        Pretty().
        Generate(&result).Err; err != nil {
        panic(err)
    }

    fmt.Println(result)
}
```




```json
{
  ...
  "root": {
    "path": "/roooooooooooootfs",
    "readonly": true
  },
  "hostname": "runc",
  ...
}
```


---

### 3. Extract a specific section / 提取指定配置块

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/Iristack//editor"
)

func main() {
    s := editor.New()
    linux, err := s.WithContent([]byte(cfg)).GetLinux()
    if err != nil {
        panic(err)
    }

    data, _ := json.MarshalIndent(linux, "", "  ")
    fmt.Println(string(data))
}
```




```json
{
  "resources": {
    "devices": [{"allow": false, "access": "rwm"}]
  },
  "namespaces": [
    {"type": "pid"}, {"type": "network"}, {"type": "ipc"},
    {"type": "uts"}, {"type": "mount"}, {"type": "cgroup"}
  ],
  "maskedPaths": [
    "/proc/acpi", "/proc/asound", "/proc/kcore", "/proc/keys",
    "/proc/latency_stats", "/proc/timer_list", "/proc/timer_stats",
    "/proc/sched_debug", "/sys/firmware", "/proc/scsi"
  ],
  "readonlyPaths": [
    "/proc/bus", "/proc/fs", "/proc/irq", "/proc/sys", "/proc/sysrq-trigger"
  ]
}
```
