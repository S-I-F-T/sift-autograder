package gradables

/*
ResourceLimits Defining resource limits to prevent overuse of limited resources.
*/
type ResourceLimits struct {
	MaxCPUTime      uint `json:"RLIMIT_CPU"`        // Max seconds
	MaxFileSize     uint `json:"RLIMIT_FSIZE"`      // Max bytes
	MaxData         uint `json:"RLIMIT_DATA"`       // Max bytes
	MaxStack        uint `json:"RLIMIT_STACK"`      // Max bytes
	MaxCoreFile     uint `json:"RLIMIT_CORE"`       // Max bytes
	MaxResidentSet  uint `json:"RLIMIT_RSS"`        // Max bytes
	MaxNumProcesses uint `json:"RLIMIT_NPROC"`      // Max threads
	MaxFileNumber   uint `json:"RLIMIT_NOFILE"`     // Max descriptor count
	MaxLockedMemory uint `json:"RLIMIT_MEMLOCK"`    // Max bytes
	MaxAddressSpace uint `json:"RLIMIT_AS"`         // Max bytes
	MaxLocks        uint `json:"RLIMIT_LOCKS"`      // Max locks and leases
	MaxSignalQueue  uint `json:"RLIMIT_SIGPENDING"` // Max num of signals pending
	MaxMessageQueue uint `json:"RLIMIT_MSGQUEUE"`   // Max bytes of messages
	MaxNice         uint `json:"RLIMIT_NICE"`       // Priority Ceiling
	MaxRTPriority   uint `json:"RLIMIT_RTPRIO"`     // Max Real Time Priority
	MaxRTTime       uint `json:"RLIMIT_RTTime"`     // Max Real Time CPU Time
}
