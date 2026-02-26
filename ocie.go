package ocie

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
)

// Spec runtime-spec wrapper
type Spec struct {
	Err    error       // errMsg
	pretty bool        // json pretty print
	specs  *specs.Spec // origin specs.Spec
}

// New make a new Spec wrapper
func New() *Spec {
	return &Spec{nil, false, &specs.Spec{}}
}

// WithPath load runtime-spec from config.json path
func (s *Spec) WithPath(path string) *Spec {
	if path == "" {
		s.Err = os.ErrInvalid
		return s
	}

	if _, err := os.Stat(path); err != nil {
		s.Err = err
		return s
	}

	data, err := os.ReadFile(path)
	if err != nil {
		s.Err = err
		return s
	}

	var spec specs.Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		s.Err = err
		return s
	}

	s.specs = &spec
	return s
}

// WithContent load runtime-spec from content
func (s *Spec) WithContent(content []byte) *Spec {
	if content == nil {
		s.Err = os.ErrInvalid
		return s
	}

	var spec specs.Spec
	if err := json.Unmarshal(content, &spec); err != nil {
		s.Err = err
		return s
	}

	s.specs = &spec
	return s
}

// SetVersion set ociVersion
func (s *Spec) SetVersion(version string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.Version = version
	return s
}

// SetProcess set process
func (s *Spec) SetProcess(process *specs.Process) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Process = process
	return s
}

// SetProcessTerminal set process.Terminal
func (s *Spec) SetProcessTerminal(b bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Terminal = b
	return s
}

// SetProcessConsoleSize set process.ConsoleSize
func (s *Spec) SetProcessConsoleSize(width, height uint) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.ConsoleSize = &specs.Box{Width: width, Height: height}
	return s
}

// SetProcessUser set process.User
func (s *Spec) SetProcessUser(user specs.User) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User = user
	return s
}

// SetProcessUserUid set process.User.UID
func (s *Spec) SetProcessUserUid(uid uint32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User.UID = uid
	return s
}

// SetProcessUserGid set process.User.GID
func (s *Spec) SetProcessUserGid(gid uint32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User.GID = gid
	return s
}

// SetProcessUserUmask set process.User.Umask
func (s *Spec) SetProcessUserUmask(umask uint32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User.Umask = &umask
	return s
}

// SetProcessUserAdditionalGids set process.User.AdditionalGids
func (s *Spec) SetProcessUserAdditionalGids(additionalGids []uint32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User.AdditionalGids = additionalGids
	return s
}

// SetProcessUserUserName set process.User.Username
func (s *Spec) SetProcessUserUserName(name string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.User.Username = name
	return s
}

// SetProcessArgs set process.Args
func (s *Spec) SetProcessArgs(args []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Args = args
	return s
}

func (s *Spec) AppendProcessArgs(args ...string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Args = append(s.specs.Process.Args, args...)
	return s
}

// SetProcessCommandLine set process.CommandLine
func (s *Spec) SetProcessCommandLine(cmdLine string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.CommandLine = cmdLine
	return s
}

// SetProcessEnv set process.Env
func (s *Spec) SetProcessEnv(env []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Env = env
	return s
}

// SetProcessCwd set process.Cwd
func (s *Spec) SetProcessCwd(cwd string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Cwd = cwd
	return s
}

// SetProcessCapabilities set process.Capabilities
func (s *Spec) SetProcessCapabilities(capabilities *specs.LinuxCapabilities) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Capabilities = capabilities
	return s
}

// SetProcessCapabilitiesBounding set process.Capabilities.Bounding
func (s *Spec) SetProcessCapabilitiesBounding(bounding []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Capabilities == nil {
		s.specs.Process.Capabilities = &specs.LinuxCapabilities{}
	}

	s.specs.Process.Capabilities.Bounding = bounding

	return s
}

// SetProcessCapabilitiesEffective set process.Capabilities.Effective
func (s *Spec) SetProcessCapabilitiesEffective(effective []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()

	if s.specs.Process.Capabilities == nil {
		s.specs.Process.Capabilities = &specs.LinuxCapabilities{}
	}

	s.specs.Process.Capabilities.Effective = effective

	return s
}

// SetProcessCapabilitiesInheritable set process.Capabilities.Inheritable
func (s *Spec) SetProcessCapabilitiesInheritable(inheritable []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}
	s.readyProcess()

	if s.specs.Process.Capabilities == nil {
		s.specs.Process.Capabilities = &specs.LinuxCapabilities{}
	}

	s.specs.Process.Capabilities.Inheritable = inheritable

	return s
}

// SetProcessCapabilitiesPermitted set process.Capabilities.Permitted
func (s *Spec) SetProcessCapabilitiesPermitted(permitted []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}
	s.readyProcess()
	if s.specs.Process.Capabilities == nil {
		s.specs.Process.Capabilities = &specs.LinuxCapabilities{}
	}

	s.specs.Process.Capabilities.Permitted = permitted

	return s
}

// SetProcessCapabilitiesAmbient set process.Capabilities.Ambient
func (s *Spec) SetProcessCapabilitiesAmbient(ambient []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Capabilities == nil {
		s.specs.Process.Capabilities = &specs.LinuxCapabilities{}
	}

	s.specs.Process.Capabilities.Ambient = ambient

	return s
}

// SetProcessRlimits set process.Rlimits
func (s *Spec) SetProcessRlimits(rlimits []specs.POSIXRlimit) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Rlimits = rlimits
	return s
}

// SetProcessNoNewPrivileges set process.NoNewPrivileges
func (s *Spec) SetProcessNoNewPrivileges(noNewPrivileges bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.NoNewPrivileges = noNewPrivileges
	return s
}

// SetProcessApparmorProfile set process.ApparmorProfile
func (s *Spec) SetProcessApparmorProfile(profile string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.ApparmorProfile = profile
	return s
}

// SetProcessOOMScoreAdj set process.OOMScoreAdj
func (s *Spec) SetProcessOOMScoreAdj(oomScoreAdj int) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.OOMScoreAdj = &oomScoreAdj
	return s
}

// SetProcessScheduler set process.Scheduler
func (s *Spec) SetProcessScheduler(scheduler *specs.Scheduler) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Scheduler = scheduler
	return s
}

// SetProcessSchedulerPolicy set process.Scheduler.Policy
func (s *Spec) SetProcessSchedulerPolicy(policy specs.LinuxSchedulerPolicy) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Policy = policy
	return s
}

// SetProcessSchedulerNice set process.Scheduler.Nice
func (s *Spec) SetProcessSchedulerNice(nice int32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Nice = nice
	return s
}

// SetProcessSchedulerPriority set process.Scheduler.Priority
func (s *Spec) SetProcessSchedulerPriority(priority int32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Priority = priority
	return s
}

// SetProcessSchedulerFlags set process.Scheduler.Flags
func (s *Spec) SetProcessSchedulerFlags(flags []specs.LinuxSchedulerFlag) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Flags = flags
	return s
}

// SetProcessSchedulerRuntime set process.Scheduler.Runtime
func (s *Spec) SetProcessSchedulerRuntime(runtime uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Runtime = runtime
	return s
}

// SetProcessSchedulerDeadline set process.Scheduler.Deadline
func (s *Spec) SetProcessSchedulerDeadline(deadline uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Deadline = deadline
	return s
}

// SetProcessSchedulerPeriod set process.Scheduler.Period
func (s *Spec) SetProcessSchedulerPeriod(period uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.Scheduler == nil {
		s.specs.Process.Scheduler = &specs.Scheduler{}
	}

	s.specs.Process.Scheduler.Period = period
	return s
}

// SetProcessSelinuxLabel set process.SelinuxLabel
func (s *Spec) SetProcessSelinuxLabel(label string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.SelinuxLabel = label
	return s
}

// SetProcessIOPriority set process.IOPriority
func (s *Spec) SetProcessIOPriority(priority *specs.LinuxIOPriority) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.IOPriority = priority
	return s
}

// SetProcessIOPriorityClass set process.IOPriority.Class
func (s *Spec) SetProcessIOPriorityClass(class specs.IOPriorityClass) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.IOPriority == nil {
		s.specs.Process.IOPriority = &specs.LinuxIOPriority{}
	}

	s.specs.Process.IOPriority.Class = class
	return s
}

// SetProcessIOPriorityPriority set process.IOPriority.Priority
func (s *Spec) SetProcessIOPriorityPriority(priority int) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.IOPriority == nil {
		s.specs.Process.IOPriority = &specs.LinuxIOPriority{}
	}

	s.specs.Process.IOPriority.Priority = priority
	return s
}

// SetProcessExecCPUAffinity set process.ExecCPUAffinity
func (s *Spec) SetProcessExecCPUAffinity(initial, final string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.ExecCPUAffinity == nil {
		s.specs.Process.ExecCPUAffinity = &specs.CPUAffinity{}
	}

	s.specs.Process.ExecCPUAffinity.Initial = initial
	s.specs.Process.ExecCPUAffinity.Final = final
	return s
}

// SetProcessExecCPUAffinityInitial set process.ExecCPUAffinity.Initial
func (s *Spec) SetProcessExecCPUAffinityInitial(initial string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.ExecCPUAffinity == nil {
		s.specs.Process.ExecCPUAffinity = &specs.CPUAffinity{}
	}

	s.specs.Process.ExecCPUAffinity.Initial = initial
	return s
}

// SetProcessExecCPUAffinityFinal set process.ExecCPUAffinity.Final
func (s *Spec) SetProcessExecCPUAffinityFinal(final string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	if s.specs.Process.ExecCPUAffinity == nil {
		s.specs.Process.ExecCPUAffinity = &specs.CPUAffinity{}
	}

	s.specs.Process.ExecCPUAffinity.Final = final
	return s
}

// SetRoot set root
func (s *Spec) SetRoot(root *specs.Root) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Root = root
	return s
}

// SetRootPath set Root.Path
func (s *Spec) SetRootPath(path string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Root == nil {
		s.specs.Root = &specs.Root{}
	}

	s.specs.Root.Path = path
	return s
}

// SetRootReadonly set Root.Readonly
func (s *Spec) SetRootReadonly(readonly bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Root == nil {
		s.specs.Root = &specs.Root{}
	}

	s.specs.Root.Readonly = readonly
	return s
}

// SetHostname set hostname
func (s *Spec) SetHostname(hostname string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Hostname = hostname
	return s
}

// SetDomainName set domainname
func (s *Spec) SetDomainName(domaiNname string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Domainname = domaiNname
	return s
}

// SetMounts set mounts
func (s *Spec) SetMounts(mounts []specs.Mount) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Mounts = mounts
	return s
}

// AddMount add mount
func (s *Spec) AddMount(mount specs.Mount) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Mounts == nil {
		s.specs.Mounts = []specs.Mount{}
	}

	s.specs.Mounts = append(s.specs.Mounts, mount)
	return s
}

// SetHooks set hooks
func (s *Spec) SetHooks(hooks *specs.Hooks) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Hooks = hooks
	return s
}

// SetHooksPrestart set Hooks.Prestart
// Deprecated: use [Hooks.CreateRuntime], [Hooks.CreateContainer], Prestart will be removed in the future
func (s *Spec) SetHooksPrestart(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.Prestart = hooks
	return s
}

// SetHooksCreateRuntime set Hooks.CreateRuntime
func (s *Spec) SetHooksCreateRuntime(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.CreateRuntime = hooks
	return s
}

// SetHooksCreateContainer set Hooks.CreateContainer
func (s *Spec) SetHooksCreateContainer(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.CreateContainer = hooks
	return s
}

// SetHooksStartContainer set Hooks.StartContainer
func (s *Spec) SetHooksStartContainer(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.StartContainer = hooks
	return s
}

// SetHooksPoststart set Hooks.Poststart
func (s *Spec) SetHooksPoststart(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.Poststart = hooks
	return s
}

// SetHooksPoststop set Hooks.Poststop
func (s *Spec) SetHooksPoststop(hooks []specs.Hook) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Hooks == nil {
		s.specs.Hooks = &specs.Hooks{}
	}

	s.specs.Hooks.Poststop = hooks
	return s
}

// SetAnnotations set annotations
func (s *Spec) SetAnnotations(annotations map[string]string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Annotations = annotations
	return s
}

// AddAnnotation add annotation
func (s *Spec) AddAnnotation(key, value string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Annotations == nil {
		s.specs.Annotations = map[string]string{}
	}

	s.specs.Annotations[key] = value
	return s
}

// SetLinux set linux
func (s *Spec) SetLinux(linux *specs.Linux) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.specs.Linux = linux
	return s
}

// SetLinuxUIDMappings set linux.UIDMappings
func (s *Spec) SetLinuxUIDMappings(uidMappings []specs.LinuxIDMapping) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	s.specs.Linux.UIDMappings = uidMappings
	return s
}

// AddLinuxUIDMapping add linux.UIDMappings
func (s *Spec) AddLinuxUIDMapping(uidMapping specs.LinuxIDMapping) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	if s.specs.Linux.UIDMappings == nil {
		s.specs.Linux.UIDMappings = []specs.LinuxIDMapping{}
	}

	s.specs.Linux.UIDMappings = append(s.specs.Linux.UIDMappings, uidMapping)
	return s
}

// SetLinuxGIDMappings set linux.GIDMappings
func (s *Spec) SetLinuxGIDMappings(gidMappings []specs.LinuxIDMapping) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	s.specs.Linux.GIDMappings = gidMappings
	return s
}

// AddLinuxGIDMapping add linux.GIDMappings
func (s *Spec) AddLinuxGIDMapping(gidMapping specs.LinuxIDMapping) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	if s.specs.Linux.GIDMappings == nil {
		s.specs.Linux.GIDMappings = []specs.LinuxIDMapping{}
	}

	s.specs.Linux.GIDMappings = append(s.specs.Linux.GIDMappings, gidMapping)
	return s
}

// SetLinuxSysctl set linux.Sysctl
func (s *Spec) SetLinuxSysctl(sysctl map[string]string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	s.specs.Linux.Sysctl = sysctl
	return s
}

// AddLinuxSysctl add linux.Sysctl
func (s *Spec) AddLinuxSysctl(key, value string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	if s.specs.Linux.Sysctl == nil {
		s.specs.Linux.Sysctl = map[string]string{}
	}

	s.specs.Linux.Sysctl[key] = value
	return s
}

// SetLinuxResources set linux.Resources
func (s *Spec) SetLinuxResources(resources *specs.LinuxResources) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	s.specs.Linux.Resources = resources
	return s
}

// SetLinuxResourcesDevices set linux.Resources.Devices
func (s *Spec) SetLinuxResourcesDevices(devices []specs.LinuxDeviceCgroup) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	s.specs.Linux.Resources.Devices = devices
	return s
}

// AddLinuxResourcesDevice add linux.Resources.Devices
func (s *Spec) AddLinuxResourcesDevice(device specs.LinuxDeviceCgroup) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	if s.specs.Linux.Resources == nil {
		s.specs.Linux.Resources = &specs.LinuxResources{}
	}

	if s.specs.Linux.Resources.Devices == nil {
		s.specs.Linux.Resources.Devices = []specs.LinuxDeviceCgroup{}
	}

	s.specs.Linux.Resources.Devices = append(s.specs.Linux.Resources.Devices, device)
	return s
}

// SetLinuxResourcesMemory set linux.Resources.Memory
func (s *Spec) SetLinuxResourcesMemory(memory *specs.LinuxMemory) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}

	if s.specs.Linux.Resources == nil {
		s.specs.Linux.Resources = &specs.LinuxResources{}
	}

	s.specs.Linux.Resources.Memory = memory
	return s
}

// SetLinuxResourcesMemoryLimitWithBytes set linux.Resources.Memory.Limit
func (s *Spec) SetLinuxResourcesMemoryLimitWithBytes(limit int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.Limit = &limit
	return s
}

// SetLinuxResourcesMemoryLimitWithMB set linux.Resources.Memory.Limit
func (s *Spec) SetLinuxResourcesMemoryLimitWithMB(limit int64) *Spec {
	return s.SetLinuxResourcesMemoryLimitWithBytes(limit * 1024 * 1024)
}

// SetLinuxResourcesMemoryReservationWithBytes set linux.Resources.Memory.Reservation
func (s *Spec) SetLinuxResourcesMemoryReservationWithBytes(reservation int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.Reservation = &reservation
	return s
}

// SetLinuxResourcesMemoryReservationWithMB set linux.Resources.Memory.Reservation
func (s *Spec) SetLinuxResourcesMemoryReservationWithMB(reservation int64) *Spec {
	return s.SetLinuxResourcesMemoryReservationWithBytes(reservation * 1024 * 1024)
}

// SetLinuxResourcesMemorySwapWithBytes set linux.Resources.Memory.Swap
func (s *Spec) SetLinuxResourcesMemorySwapWithBytes(swap int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.Swap = &swap
	return s
}

// SetLinuxResourcesMemorySwapWithMB set linux.Resources.Memory.Swap
func (s *Spec) SetLinuxResourcesMemorySwapWithMB(swap int64) *Spec {
	return s.SetLinuxResourcesMemorySwapWithBytes(swap * 1024 * 1024)
}

// SetLinuxResourcesMemoryKernelWithBytes set linux.Resources.Memory.Kernel
// Deprecated: kernel-memory limits are not supported in cgroups v2, and were obsoleted
func (s *Spec) SetLinuxResourcesMemoryKernelWithBytes(kernel int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.Kernel = &kernel
	return s
}

// SetLinuxResourcesMemoryKernelWithMB set linux.Resources.Memory.Kernel
// Deprecated: kernel-memory limits are not supported in cgroups v2, and were obsoleted
func (s *Spec) SetLinuxResourcesMemoryKernelWithMB(kernel int64) *Spec {
	return s.SetLinuxResourcesMemoryKernelWithBytes(kernel * 1024 * 1024)
}

// SetLinuxResourcesMemoryKernelTCPWithBytes set linux.Resources.Memory.KernelTCP
func (s *Spec) SetLinuxResourcesMemoryKernelTCPWithBytes(kernelTCP int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.KernelTCP = &kernelTCP
	return s
}

// SetLinuxResourcesMemoryKernelTCPWithMB set linux.Resources.Memory.KernelTCP
func (s *Spec) SetLinuxResourcesMemoryKernelTCPWithMB(kernelTCP int64) *Spec {
	return s.SetLinuxResourcesMemoryKernelTCPWithBytes(kernelTCP * 1024 * 1024)
}

// SetLinuxResourcesMemorySwappiness set linux.Resources.Memory.Swappiness
func (s *Spec) SetLinuxResourcesMemorySwappiness(swappiness uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.Swappiness = &swappiness
	return s
}

// SetLinuxResourcesMemoryDisableOOMKiller set linux.Resources.Memory.DisableOOMKiller
func (s *Spec) SetLinuxResourcesMemoryDisableOOMKiller(disableOOMKiller bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.DisableOOMKiller = &disableOOMKiller
	return s
}

// SetLinuxResourcesMemoryUseHierarchy set linux.Resources.Memory.UseHierarchy
func (s *Spec) SetLinuxResourcesMemoryUseHierarchy(useHierarchy bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.UseHierarchy = &useHierarchy
	return s
}

// SetLinuxResourceMemoryCheckBeforeUpdate set linux.Resources.Memory.CheckBeforeUpdate
func (s *Spec) SetLinuxResourceMemoryCheckBeforeUpdate(checkBeforeUpdate bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.Resources.Memory.CheckBeforeUpdate = &checkBeforeUpdate
	return s
}

// SetLinuxResourcesCpu set linux.Resources.CPU
func (s *Spec) SetLinuxResourcesCpu(cpu *specs.LinuxCPU) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.CPU = cpu
	return s
}

// SetLinuxResourcesCpuShares set linux.Resources.CPU.Shares
func (s *Spec) SetLinuxResourcesCpuShares(shares uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Shares = &shares
	return s
}

// SetLinuxResourcesCpuQuota set linux.Resources.CPU.Quota
func (s *Spec) SetLinuxResourcesCpuQuota(quota int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Quota = &quota
	return s
}

// SetLinuxResourcesCpuBurst set linux.Resources.CPU.Burst
func (s *Spec) SetLinuxResourcesCpuBurst(burst uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Burst = &burst
	return s
}

// SetLinuxResourcesCpuPeriod set linux.Resources.CPU.Period
func (s *Spec) SetLinuxResourcesCpuPeriod(period uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Period = &period
	return s
}

// SetLinuxResourcesCpuRealtimeRuntime set linux.Resources.CPU.RealtimeRuntime
func (s *Spec) SetLinuxResourcesCpuRealtimeRuntime(realtime int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.RealtimeRuntime = &realtime
	return s
}

// SetLinuxResourcesCpuRealtimePeriod set linux.Resources.CPU.RealtimePeriod
func (s *Spec) SetLinuxResourcesCpuRealtimePeriod(realtimePeriod uint64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.RealtimePeriod = &realtimePeriod
	return s
}

// SetLinuxResourcesCpuCpus set linux.Resources.CPU.Cpus
func (s *Spec) SetLinuxResourcesCpuCpus(cpus string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Cpus = cpus
	return s
}

// SetLinuxResourcesCpuMens set linux.Resources.CPU.Mems
func (s *Spec) SetLinuxResourcesCpuMens(mems string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Mems = mems
	return s
}

// SetLinuxResourcesCpuIdle set linux.Resources.CPU.Idle
func (s *Spec) SetLinuxResourcesCpuIdle(idle int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesCPU()
	s.specs.Linux.Resources.CPU.Idle = &idle
	return s
}

// SetLinuxResourcesPids set linux.Resources.Pids
func (s *Spec) SetLinuxResourcesPids(pids *specs.LinuxPids) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.Pids = pids
	return s
}

// SetLinuxResourcesPidsLimit set linux.Resources.Pids.Limit
func (s *Spec) SetLinuxResourcesPidsLimit(limit int64) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesPids()
	s.specs.Linux.Resources.Pids.Limit = &limit
	return s
}

// SetLinuxResourcesBlockIO set linux.Resources.BlockIO
func (s *Spec) SetLinuxResourcesBlockIO(blockIO *specs.LinuxBlockIO) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.BlockIO = blockIO
	return s
}

// SetLinuxResourcesBlockIOWeight set linux.Resources.BlockIO.Weight
func (s *Spec) SetLinuxResourcesBlockIOWeight(weight uint16) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.Weight = &weight
	return s
}

// SetLinuxResourcesBlockIOLeafWeight set linux.Resources.BlockIO.LeafWeight
func (s *Spec) SetLinuxResourcesBlockIOLeafWeight(leafWeight uint16) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.LeafWeight = &leafWeight
	return s
}

// SetLinuxResourcesBlockIOWeightDevice set linux.Resources.BlockIO.WeightDevice
func (s *Spec) SetLinuxResourcesBlockIOWeightDevice(weightDevice []specs.LinuxWeightDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.WeightDevice = weightDevice
	return s
}

// AddLinuxResourcesBlockIOWeightDevice set linux.Resources.BlockIO.WeightDevice
func (s *Spec) AddLinuxResourcesBlockIOWeightDevice(weightDevice specs.LinuxWeightDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIOWeightDevice()
	s.specs.Linux.Resources.BlockIO.WeightDevice = append(s.specs.Linux.Resources.BlockIO.WeightDevice, weightDevice)
	return s
}

// SetLinuxResourcesBlockIOThrottleReadBpsDevice set linux.Resources.BlockIO.ThrottleReadBpsDevice
func (s *Spec) SetLinuxResourcesBlockIOThrottleReadBpsDevice(throttleReadBpsDevice []specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.ThrottleReadBpsDevice = throttleReadBpsDevice
	return s
}

// AddLinuxResourcesBlockIOThrottleReadBpsDevice set linux.Resources.BlockIO.ThrottleReadBpsDevice
func (s *Spec) AddLinuxResourcesBlockIOThrottleReadBpsDevice(throttleReadBpsDevice specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIOThrottleReadBpsDevice()
	s.specs.Linux.Resources.BlockIO.ThrottleReadBpsDevice = append(s.specs.Linux.Resources.BlockIO.ThrottleReadBpsDevice, throttleReadBpsDevice)
	return s
}

// SetLinuxResourcesBlockIOThrottleWriteBpsDevice set linux.Resources.BlockIO.ThrottleWriteBpsDevice
func (s *Spec) SetLinuxResourcesBlockIOThrottleWriteBpsDevice(throttleWriteBpsDevice []specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.ThrottleWriteBpsDevice = throttleWriteBpsDevice
	return s
}

// AddLinuxResourcesBlockIOThrottleWriteBpsDevice set linux.Resources.BlockIO.ThrottleWriteBpsDevice
func (s *Spec) AddLinuxResourcesBlockIOThrottleWriteBpsDevice(throttleWriteBpsDevice specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIOThrottleWriteBpsDevice()
	s.specs.Linux.Resources.BlockIO.ThrottleWriteBpsDevice = append(s.specs.Linux.Resources.BlockIO.ThrottleWriteBpsDevice, throttleWriteBpsDevice)
	return s
}

// SetLinuxResourcesBlockIOThrottleReadIOPSDevice add linux.Resources.BlockIO.ThrottleReadIOPSDevice
func (s *Spec) SetLinuxResourcesBlockIOThrottleReadIOPSDevice(throttleReadIOPSDevice []specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.ThrottleReadIOPSDevice = throttleReadIOPSDevice
	return s
}

// AddLinuxResourcesBlockIOThrottleReadIOPSDevice add linux.Resources.BlockIO.ThrottleReadIOPSDevice
func (s *Spec) AddLinuxResourcesBlockIOThrottleReadIOPSDevice(throttleReadIOPSDevice specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIOThrottleReadIOPSDevice()
	s.specs.Linux.Resources.BlockIO.ThrottleReadIOPSDevice = append(s.specs.Linux.Resources.BlockIO.ThrottleReadIOPSDevice, throttleReadIOPSDevice)
	return s
}

// SetLinuxResourcesBlockIOThrottleWriteIOPSDevice set linux.Resources.BlockIO.ThrottleWriteIOPSDevice
func (s *Spec) SetLinuxResourcesBlockIOThrottleWriteIOPSDevice(throttleWriteIOPSDevice []specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIO()
	s.specs.Linux.Resources.BlockIO.ThrottleWriteIOPSDevice = throttleWriteIOPSDevice
	return s
}

// AddLinuxResourcesBlockIOThrottleWriteIOPSDevice add linux.Resources.BlockIO.ThrottleWriteIOPSDevice
func (s *Spec) AddLinuxResourcesBlockIOThrottleWriteIOPSDevice(throttleWriteIOPSDevice specs.LinuxThrottleDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesBlockIOThrottleWriteIOPSDevice()
	s.specs.Linux.Resources.BlockIO.ThrottleWriteIOPSDevice = append(s.specs.Linux.Resources.BlockIO.ThrottleWriteIOPSDevice, throttleWriteIOPSDevice)
	return s
}

// SetLinuxResourcesHugepageLimits set linux.Resources.HugepageLimits
func (s *Spec) SetLinuxResourcesHugepageLimits(hugepageLimits []specs.LinuxHugepageLimit) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.HugepageLimits = hugepageLimits
	return s
}

// AddLinuxResourcesHugepageLimits add linux.Resources.HugepageLimits
func (s *Spec) AddLinuxResourcesHugepageLimits(hugepageLimit specs.LinuxHugepageLimit) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesHugepageLimits()
	s.specs.Linux.Resources.HugepageLimits = append(s.specs.Linux.Resources.HugepageLimits, hugepageLimit)
	return s
}

// SetLinuxResourcesNetwork set linux.Resources.Network
func (s *Spec) SetLinuxResourcesNetwork(network *specs.LinuxNetwork) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.Network = network
	return s
}

// SetLinuxResourcesNetworkClassID set linux.Resources.Network.ClassID
func (s *Spec) SetLinuxResourcesNetworkClassID(classID uint32) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesNetwork()
	s.specs.Linux.Resources.Network.ClassID = &classID
	return s
}

// SetLinuxResourcesNetworkPriorities set linux.Resources.Network.Priorities
func (s *Spec) SetLinuxResourcesNetworkPriorities(priorities []specs.LinuxInterfacePriority) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesNetwork()
	s.specs.Linux.Resources.Network.Priorities = priorities
	return s
}

// AddLinuxResourcesNetworkPriorities add linux.Resources.Network.Priorities
func (s *Spec) AddLinuxResourcesNetworkPriorities(priority specs.LinuxInterfacePriority) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesNetworkPriorities()
	s.specs.Linux.Resources.Network.Priorities = append(s.specs.Linux.Resources.Network.Priorities, priority)
	return s
}

// SetLinuxResourcesRdma set linux.Resources.Rdma
func (s *Spec) SetLinuxResourcesRdma(rdma map[string]specs.LinuxRdma) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.Rdma = rdma
	return s
}

// AddLinuxResourcesRdma add linux.Resources.Rdma
func (s *Spec) AddLinuxResourcesRdma(key string, rdma specs.LinuxRdma) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesRdma()
	s.specs.Linux.Resources.Rdma[key] = rdma
	return s
}

// SetLinuxResourcesUnified set linux.Resources.Unified
func (s *Spec) SetLinuxResourcesUnified(unified map[string]string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResources()
	s.specs.Linux.Resources.Unified = unified
	return s
}

// AddLinuxResourcesUnified add linux.Resources.Unified
func (s *Spec) AddLinuxResourcesUnified(key string, value string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesUnified()
	s.specs.Linux.Resources.Unified[key] = value
	return s
}

// SetLinuxCgroupsPath set linux.CgroupsPath
func (s *Spec) SetLinuxCgroupsPath(cgroupsPath string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.CgroupsPath = cgroupsPath
	return s
}

// SetLinuxNameSpaces set linux.Namespaces
func (s *Spec) SetLinuxNameSpaces(namespaces []specs.LinuxNamespace) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.Namespaces = namespaces
	return s
}

// AddLinuxNameSpace add linux.Namespaces
func (s *Spec) AddLinuxNameSpace(nsType specs.LinuxNamespaceType, path string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxNameSpaces()
	s.specs.Linux.Namespaces = append(s.specs.Linux.Namespaces, specs.LinuxNamespace{
		Type: nsType,
		Path: path,
	})
	return s
}

// SetLinuxDevices set linux.Devices
func (s *Spec) SetLinuxDevices(devices []specs.LinuxDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.Devices = devices
	return s
}

// AddLinuxDevices add linux.Devices
func (s *Spec) AddLinuxDevices(device specs.LinuxDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxDevices()
	s.specs.Linux.Devices = append(s.specs.Linux.Devices, device)
	return s
}

// SetLinuxNetDevices set linux.NetDevices
func (s *Spec) SetLinuxNetDevices(devices map[string]specs.LinuxNetDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxDevices()
	s.specs.Linux.NetDevices = devices
	return s
}

// AddLinuxNetDevices add linux.NetDevices
func (s *Spec) AddLinuxNetDevices(key string, device specs.LinuxNetDevice) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxNetDevices()
	s.specs.Linux.NetDevices[key] = device
	return s
}

// SetLinuxSeccomp set specs.Linux.Seccomp
func (s *Spec) SetLinuxSeccomp(seccomp *specs.LinuxSeccomp) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.Seccomp = seccomp
	return s
}

// SetLinuxSeccompDefaultAction set specs.Linux.Seccomp.DefaultAction
func (s *Spec) SetLinuxSeccompDefaultAction(action specs.LinuxSeccompAction) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.DefaultAction = action
	return s
}

// SetLinuxSeccompDefaultErrnoRet set specs.Linux.Seccomp.DefaultErrnoRet
func (s *Spec) SetLinuxSeccompDefaultErrnoRet(errnoRet uint) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.DefaultErrnoRet = &errnoRet
	return s
}

// SetLinuxSeccompArchitectures set specs.Linux.Seccomp.Architectures
func (s *Spec) SetLinuxSeccompArchitectures(architectures []specs.Arch) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.Architectures = architectures
	return s
}

// AddLinuxSeccompArchitectures add specs.Linux.Seccomp.Architectures
func (s *Spec) AddLinuxSeccompArchitectures(architecture specs.Arch) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccompArchitectures()
	s.specs.Linux.Seccomp.Architectures = append(s.specs.Linux.Seccomp.Architectures, architecture)
	return s
}

// SetLinuxSeccompFlags set specs.Linux.Seccomp.Flags
func (s *Spec) SetLinuxSeccompFlags(flags []specs.LinuxSeccompFlag) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.Flags = flags
	return s
}

// AddLinuxSeccompFlags add specs.Linux.Seccomp.Flags
func (s *Spec) AddLinuxSeccompFlags(flag specs.LinuxSeccompFlag) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccompFlags()
	s.specs.Linux.Seccomp.Flags = append(s.specs.Linux.Seccomp.Flags, flag)
	return s
}

// SetLinuxSeccompListenerPath set specs.Linux.Seccomp.ListenerPath
func (s *Spec) SetLinuxSeccompListenerPath(listenerPath string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.ListenerPath = listenerPath
	return s
}

// SetLinuxSeccompListenerMetadata set specs.Linux.Seccomp.ListenerMetadata
func (s *Spec) SetLinuxSeccompListenerMetadata(listenerMetadata string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.ListenerMetadata = listenerMetadata
	return s
}

// SetLinuxSeccompSyscalls set specs.Linux.Seccomp.Syscalls
func (s *Spec) SetLinuxSeccompSyscalls(syscalls []specs.LinuxSyscall) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccomp()
	s.specs.Linux.Seccomp.Syscalls = syscalls
	return s
}

// AddLinuxSeccompSyscalls add Linux.Seccomp.Syscalls
func (s *Spec) AddLinuxSeccompSyscalls(syscall specs.LinuxSyscall) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxSeccompSyscalls()
	s.specs.Linux.Seccomp.Syscalls = append(s.specs.Linux.Seccomp.Syscalls, syscall)
	return s
}

// SetLinuxRootfsPropagation set specs.Linux.RootfsPropagation
func (s *Spec) SetLinuxRootfsPropagation(propagation string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.RootfsPropagation = propagation
	return s
}

// SetLinuxMaskedPaths set specs.Linux.MaskedPaths
func (s *Spec) SetLinuxMaskedPaths(maskedPaths []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.MaskedPaths = maskedPaths
	return s
}

// AddLinuxMaskedPaths add specs.Linux.MaskedPaths
func (s *Spec) AddLinuxMaskedPaths(maskedPath string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxMaskedPaths()
	s.specs.Linux.MaskedPaths = append(s.specs.Linux.MaskedPaths, maskedPath)
	return s
}

// SetLinuxReadonlyPaths set specs.Linux.ReadonlyPaths
func (s *Spec) SetLinuxReadonlyPaths(readonlyPaths []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.ReadonlyPaths = readonlyPaths
	return s
}

// AddLinuxReadonlyPaths add specs.Linux.ReadonlyPaths
func (s *Spec) AddLinuxReadonlyPaths(readonlyPath string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxReadonlyPaths()
	s.specs.Linux.ReadonlyPaths = append(s.specs.Linux.ReadonlyPaths, readonlyPath)
	return s
}

// SetLinuxMountLabel set specs.Linux.MountLabel
func (s *Spec) SetLinuxMountLabel(label string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.MountLabel = label
	return s
}

// SetLinuxIntelRdt set specs.Linux.IntelRdt
func (s *Spec) SetLinuxIntelRdt(intelRdt *specs.LinuxIntelRdt) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.IntelRdt = intelRdt
	return s
}

// SetLinuxIntelRdtClosId set specs.Linux.IntelRdt.ClosID
func (s *Spec) SetLinuxIntelRdtClosId(closID string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxIntelRdt()
	s.specs.Linux.IntelRdt.ClosID = closID
	return s
}

// SetLinuxIntelRdtSchemata set specs.Linux.IntelRdt.Schemata
func (s *Spec) SetLinuxIntelRdtSchemata(schemata []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxIntelRdt()
	s.specs.Linux.IntelRdt.Schemata = schemata
	return s
}

// SetLinuxIntelRdtL3CacheSchema set specs.Linux.IntelRdt.L3CacheSchema
func (s *Spec) SetLinuxIntelRdtL3CacheSchema(l3CacheSchema string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxIntelRdt()
	s.specs.Linux.IntelRdt.L3CacheSchema = l3CacheSchema
	return s
}

// SetLinuxIntelRdtMemBwSchema set specs.Linux.IntelRdt.MemBwSchema
func (s *Spec) SetLinuxIntelRdtMemBwSchema(memBwSchema string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxIntelRdt()
	s.specs.Linux.IntelRdt.MemBwSchema = memBwSchema
	return s
}

// SetLinuxIntelRdtEnableMonitoring set specs.Linux.IntelRdt.EnableMonitoring
func (s *Spec) SetLinuxIntelRdtEnableMonitoring(enableMonitoring bool) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxIntelRdt()
	s.specs.Linux.IntelRdt.EnableMonitoring = enableMonitoring
	return s
}

// SetLinuxMemoryPolicy set specs.Linux.MemoryPolicy
func (s *Spec) SetLinuxMemoryPolicy(policy *specs.LinuxMemoryPolicy) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxResourcesMemory()
	s.specs.Linux.MemoryPolicy = policy
	return s
}

// SetLinuxMemoryPolicyMode set specs.Linux.MemoryPolicy.Mode
func (s *Spec) SetLinuxMemoryPolicyMode(mode specs.MemoryPolicyModeType) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxMemoryPolicy()
	s.specs.Linux.MemoryPolicy.Mode = mode
	return s
}

// SetLinuxMemoryPolicyNodes set specs.Linux.MemoryPolicy.Nodes
func (s *Spec) SetLinuxMemoryPolicyNodes(nodes string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxMemoryPolicy()
	s.specs.Linux.MemoryPolicy.Nodes = nodes
	return s
}

// SetLinuxMemoryPolicyFlags set specs.Linux.MemoryPolicy.Flags
func (s *Spec) SetLinuxMemoryPolicyFlags(flags []specs.MemoryPolicyFlagType) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxMemoryPolicy()
	s.specs.Linux.MemoryPolicy.Flags = flags
	return s
}

// AddLinuxMemoryPolicyFlags add specs.Linux.MemoryPolicy.Flags
func (s *Spec) AddLinuxMemoryPolicyFlags(flag specs.MemoryPolicyFlagType) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxMemoryPolicyFlags()
	s.specs.Linux.MemoryPolicy.Flags = append(s.specs.Linux.MemoryPolicy.Flags, flag)
	return s
}

// SetLinuxPersonality set specs.Linux.Personality
func (s *Spec) SetLinuxPersonality(personality *specs.LinuxPersonality) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.Personality = personality
	return s
}

func (s *Spec) AppendProcessEnv(env []string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyProcess()
	s.specs.Process.Env = append(s.specs.Process.Env, env...)
	return s
}

// SetLinuxPersonalityDomain set specs.Linux.Personality.Domain
func (s *Spec) SetLinuxPersonalityDomain(domain specs.LinuxPersonalityDomain) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxPersonality()
	s.specs.Linux.Personality.Domain = domain
	return s
}

// SetLinuxPersonalityFlags set specs.Linux.Personality.Flags
func (s *Spec) SetLinuxPersonalityFlags(flags []specs.LinuxPersonalityFlag) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxPersonality()
	s.specs.Linux.Personality.Flags = flags
	return s
}

// AddLinuxPersonalityFlags add specs.Linux.Personality.Flags
func (s *Spec) AddLinuxPersonalityFlags(flag specs.LinuxPersonalityFlag) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxPersonalityFlags()
	s.specs.Linux.Personality.Flags = append(s.specs.Linux.Personality.Flags, flag)
	return s
}

// SetLinuxTimeOffsets set specs.Linux.TimeOffsets
func (s *Spec) SetLinuxTimeOffsets(timeOffsets map[string]specs.LinuxTimeOffset) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinux()
	s.specs.Linux.TimeOffsets = timeOffsets
	return s
}

// AddLinuxTimeOffset add specs.Linux.TimeOffsets
func (s *Spec) AddLinuxTimeOffset(name string, timeOffset specs.LinuxTimeOffset) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readyLinuxTimeOffsets()
	s.specs.Linux.TimeOffsets[name] = timeOffset
	return s
}

// SetSolaris set specs.Solaris
func (s *Spec) SetSolaris(solaris *specs.Solaris) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.Solaris = solaris
	return s
}

// SetSolarisMilestone set specs.Solaris.Milestone
func (s *Spec) SetSolarisMilestone(milestone string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolaris()
	s.specs.Solaris.Milestone = milestone
	return s
}

// SetSolarisLimitPriv set specs.Solaris.LimitPriv
func (s *Spec) SetSolarisLimitPriv(limitPriv string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolaris()
	s.specs.Solaris.LimitPriv = limitPriv
	return s
}

// SetSolarisMaxShmMemory set specs.Solaris.MaxShmMemory
func (s *Spec) SetSolarisMaxShmMemory(maxShmMemory string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolaris()
	s.specs.Solaris.MaxShmMemory = maxShmMemory
	return s
}

// SetSolarisAnet set specs.Solaris.Anet
func (s *Spec) SetSolarisAnet(anet []specs.SolarisAnet) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolaris()
	s.specs.Solaris.Anet = anet
	return s
}

// SetSolarisCappedCpu set specs.Solaris.CappedCPU
func (s *Spec) SetSolarisCappedCpu(cappedCpu *specs.SolarisCappedCPU) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolarisCappedCpu()
	s.specs.Solaris.CappedCPU = cappedCpu
	return s
}

// SetSolarisCappedCpuNCpus set specs.Solaris.CappedCPU.Ncpus
func (s *Spec) SetSolarisCappedCpuNCpus(nCpus string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolarisCappedCpu()
	s.specs.Solaris.CappedCPU.Ncpus = nCpus
	return s
}

// SetSolarisCappedMemory set specs.Solaris.CappedMemory
func (s *Spec) SetSolarisCappedMemory(cappedMemory *specs.SolarisCappedMemory) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolarisCappedMemory()
	s.specs.Solaris.CappedMemory = cappedMemory
	return s
}

// SetSolarisCappedMemoryPhysical set specs.Solaris.CappedMemory.Physical
func (s *Spec) SetSolarisCappedMemoryPhysical(physical string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolarisCappedMemory()
	s.specs.Solaris.CappedMemory.Physical = physical
	return s
}

// SetSolarisCappedMemorySwap set specs.Solaris.CappedMemory.Swap
func (s *Spec) SetSolarisCappedMemorySwap(swap string) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySolarisCappedMemory()
	s.specs.Solaris.CappedMemory.Swap = swap
	return s
}

// SetWindows set Windows
func (s *Spec) SetWindows(windows *specs.Windows) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.Windows = windows
	return s
}

// SetVM set VM
func (s *Spec) SetVM(vm *specs.VM) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.VM = vm
	return s
}

// SetZOS set ZOS
func (s *Spec) SetZOS(zos *specs.ZOS) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.ZOS = zos
	return s
}

// SetFreeBSD set FreeBSD
func (s *Spec) SetFreeBSD(freeBSD *specs.FreeBSD) *Spec {
	if err := s.ready(); err != nil {
		return s
	}

	s.readySpecs()
	s.specs.FreeBSD = freeBSD
	return s
}

// Pretty set JSON indent for Generate
func (s *Spec) Pretty() *Spec {
	s.pretty = true
	return s
}

// Generate generate json string
func (s *Spec) Generate(buf *string) *Spec {
	var rst []byte
	if s.pretty {
		rst, s.Err = json.MarshalIndent(s.specs, "", "  ")
	} else {
		rst, s.Err = json.Marshal(s.specs)
	}

	*buf = string(rst)
	return s
}

//
// Get
//

func (s *Spec) GetSpec() (*specs.Spec, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs, nil
}

func (s *Spec) GetVersion() (string, error) {
	if s.specs == nil {
		return "", fmt.Errorf("specs is nil")
	}
	return s.specs.Version, nil
}

func (s *Spec) GetProcess() (*specs.Process, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Process, nil // 可能为 nil，由调用者判断
}

func (s *Spec) GetRoot() (*specs.Root, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Root, nil
}

func (s *Spec) GetHostname() (string, error) {
	if s.specs == nil {
		return "", fmt.Errorf("specs is nil")
	}
	return s.specs.Hostname, nil
}

func (s *Spec) GetDomainName() (string, error) {
	if s.specs == nil {
		return "", fmt.Errorf("specs is nil")
	}
	return s.specs.Domainname, nil
}

func (s *Spec) GetMounts() ([]specs.Mount, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Mounts, nil
}

func (s *Spec) GetHooks() (*specs.Hooks, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Hooks, nil
}

func (s *Spec) GetAnnotations() (map[string]string, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Annotations, nil
}

func (s *Spec) GetLinux() (*specs.Linux, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Linux, nil
}

func (s *Spec) GetSolaris() (*specs.Solaris, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Solaris, nil
}

func (s *Spec) GetWindows() (*specs.Windows, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.Windows, nil
}

func (s *Spec) GetVM() (*specs.VM, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.VM, nil
}

func (s *Spec) GetZOS() (*specs.ZOS, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.ZOS, nil
}

func (s *Spec) GetFreeBSD() (*specs.FreeBSD, error) {
	if s.specs == nil {
		return nil, fmt.Errorf("specs is nil")
	}
	return s.specs.FreeBSD, nil
}

// check func's
func (s *Spec) ready() error {
	if s.specs == nil {
		return os.ErrInvalid
	}

	if s.Err != nil {
		return s.Err
	}
	return nil
}

func (s *Spec) readySpecs() {
	if s.specs == nil {
		s.specs = &specs.Spec{}
	}
}

func (s *Spec) readyLinux() {
	s.readySpecs()

	if s.specs.Linux == nil {
		s.specs.Linux = &specs.Linux{}
	}
}

func (s *Spec) readyLinuxResources() {
	s.readyLinux()

	if s.specs.Linux.Resources == nil {
		s.specs.Linux.Resources = &specs.LinuxResources{}
	}
}

func (s *Spec) readyLinuxResourcesMemory() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.Memory == nil {
		s.specs.Linux.Resources.Memory = &specs.LinuxMemory{}
	}
}

func (s *Spec) readyLinuxResourcesCPU() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.CPU == nil {
		s.specs.Linux.Resources.CPU = &specs.LinuxCPU{}
	}
}

func (s *Spec) readyLinuxResourcesPids() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.Pids == nil {
		s.specs.Linux.Resources.Pids = &specs.LinuxPids{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIO() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.BlockIO == nil {
		s.specs.Linux.Resources.BlockIO = &specs.LinuxBlockIO{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIOWeightDevice() {
	s.readyLinuxResourcesBlockIO()

	if s.specs.Linux.Resources.BlockIO.WeightDevice == nil {
		s.specs.Linux.Resources.BlockIO.WeightDevice = []specs.LinuxWeightDevice{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIOThrottleReadBpsDevice() {
	s.readyLinuxResourcesBlockIO()

	if s.specs.Linux.Resources.BlockIO.ThrottleReadBpsDevice == nil {
		s.specs.Linux.Resources.BlockIO.ThrottleReadBpsDevice = []specs.LinuxThrottleDevice{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIOThrottleReadIOPSDevice() {
	s.readyLinuxResourcesBlockIO()

	if s.specs.Linux.Resources.BlockIO.ThrottleReadIOPSDevice == nil {
		s.specs.Linux.Resources.BlockIO.ThrottleReadIOPSDevice = []specs.LinuxThrottleDevice{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIOThrottleWriteBpsDevice() {
	s.readyLinuxResourcesBlockIO()

	if s.specs.Linux.Resources.BlockIO.ThrottleWriteBpsDevice == nil {
		s.specs.Linux.Resources.BlockIO.ThrottleWriteBpsDevice = []specs.LinuxThrottleDevice{}
	}
}

func (s *Spec) readyLinuxResourcesBlockIOThrottleWriteIOPSDevice() {
	s.readyLinuxResourcesBlockIO()

	if s.specs.Linux.Resources.BlockIO.ThrottleWriteIOPSDevice == nil {
		s.specs.Linux.Resources.BlockIO.ThrottleWriteIOPSDevice = []specs.LinuxThrottleDevice{}
	}
}

func (s *Spec) readyLinuxResourcesHugepageLimits() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.HugepageLimits == nil {
		s.specs.Linux.Resources.HugepageLimits = []specs.LinuxHugepageLimit{}
	}
}

func (s *Spec) readyLinuxResourcesNetwork() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.Network == nil {
		s.specs.Linux.Resources.Network = &specs.LinuxNetwork{}
	}
}

func (s *Spec) readyLinuxResourcesNetworkPriorities() {
	s.readyLinuxResourcesNetwork()

	if s.specs.Linux.Resources.Network.Priorities == nil {
		s.specs.Linux.Resources.Network.Priorities = []specs.LinuxInterfacePriority{}
	}
}

func (s *Spec) readyLinuxResourcesRdma() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.Rdma == nil {
		s.specs.Linux.Resources.Rdma = map[string]specs.LinuxRdma{}
	}
}

func (s *Spec) readyLinuxResourcesUnified() {
	s.readyLinuxResources()

	if s.specs.Linux.Resources.Unified == nil {
		s.specs.Linux.Resources.Unified = map[string]string{}
	}
}

func (s *Spec) readyLinuxNameSpaces() {
	s.readyLinux()

	if s.specs.Linux.Namespaces == nil {
		s.specs.Linux.Namespaces = []specs.LinuxNamespace{}
	}
}

func (s *Spec) readyLinuxDevices() {
	s.readyLinux()

	if s.specs.Linux.Devices == nil {
		s.specs.Linux.Devices = []specs.LinuxDevice{}
	}
}

func (s *Spec) readyLinuxNetDevices() {
	s.readyLinux()

	if s.specs.Linux.Devices == nil {
		s.specs.Linux.Devices = []specs.LinuxDevice{}
	}
}

func (s *Spec) readyLinuxSeccomp() {
	s.readyLinux()

	if s.specs.Linux.Seccomp == nil {
		s.specs.Linux.Seccomp = &specs.LinuxSeccomp{}
	}
}

func (s *Spec) readyLinuxSeccompArchitectures() {
	s.readyLinuxSeccomp()

	if s.specs.Linux.Seccomp.Architectures == nil {
		s.specs.Linux.Seccomp.Architectures = []specs.Arch{}
	}
}

func (s *Spec) readyLinuxSeccompSyscalls() {
	s.readyLinuxSeccomp()

	if s.specs.Linux.Seccomp.Syscalls == nil {
		s.specs.Linux.Seccomp.Syscalls = []specs.LinuxSyscall{}
	}
}

func (s *Spec) readyLinuxSeccompFlags() {
	s.readyLinuxSeccomp()

	if s.specs.Linux.Seccomp.Flags == nil {
		s.specs.Linux.Seccomp.Flags = []specs.LinuxSeccompFlag{}
	}
}

func (s *Spec) readyLinuxMaskedPaths() {
	s.readyLinux()

	if s.specs.Linux.MaskedPaths == nil {
		s.specs.Linux.MaskedPaths = []string{}
	}
}

func (s *Spec) readyLinuxReadonlyPaths() {
	s.readyLinux()

	if s.specs.Linux.ReadonlyPaths == nil {
		s.specs.Linux.ReadonlyPaths = []string{}
	}
}

func (s *Spec) readyLinuxIntelRdt() {
	s.readyLinux()

	if s.specs.Linux.IntelRdt == nil {
		s.specs.Linux.IntelRdt = &specs.LinuxIntelRdt{}
	}
}

func (s *Spec) readyLinuxMemoryPolicy() {
	s.readyLinux()

	if s.specs.Linux.MemoryPolicy == nil {
		s.specs.Linux.MemoryPolicy = &specs.LinuxMemoryPolicy{}
	}
}

func (s *Spec) readyLinuxPersonality() {
	s.readyLinux()

	if s.specs.Linux.Personality == nil {
		s.specs.Linux.Personality = &specs.LinuxPersonality{}
	}
}

func (s *Spec) readyLinuxTimeOffsets() {
	s.readyLinux()

	if s.specs.Linux.TimeOffsets == nil {
		s.specs.Linux.TimeOffsets = map[string]specs.LinuxTimeOffset{}
	}
}

func (s *Spec) readyLinuxIntelRdtSchemata() {
	s.readyLinuxIntelRdt()

	if s.specs.Linux.IntelRdt.Schemata == nil {
		s.specs.Linux.IntelRdt.Schemata = []string{}
	}
}

func (s *Spec) readyLinuxMemoryPolicyFlags() {
	s.readyLinuxMemoryPolicy()

	if s.specs.Linux.MemoryPolicy.Flags == nil {
		s.specs.Linux.MemoryPolicy.Flags = []specs.MemoryPolicyFlagType{}
	}
	return
}

func (s *Spec) readyLinuxPersonalityFlags() {
	s.readyLinuxPersonality()

	if s.specs.Linux.Personality.Flags == nil {
		s.specs.Linux.Personality.Flags = []specs.LinuxPersonalityFlag{}
	}
}

func (s *Spec) readySolaris() {
	s.readySpecs()

	if s.specs.Solaris == nil {
		s.specs.Solaris = &specs.Solaris{}
	}
}

func (s *Spec) readySolarisAnet() {
	s.readySolaris()

	if s.specs.Solaris.Anet == nil {
		s.specs.Solaris.Anet = []specs.SolarisAnet{}
	}
}

func (s *Spec) readySolarisCappedCpu() {
	s.readySolaris()

	if s.specs.Solaris.CappedCPU == nil {
		s.specs.Solaris.CappedCPU = &specs.SolarisCappedCPU{}
	}
}

func (s *Spec) readySolarisCappedMemory() {
	s.readySolaris()

	if s.specs.Solaris.CappedMemory == nil {
		s.specs.Solaris.CappedMemory = &specs.SolarisCappedMemory{}
	}
}

func (s *Spec) readyProcess() {
	s.readySpecs()

	if s.specs.Process == nil {
		s.specs.Process = &specs.Process{}
	}
}
