package cri

import "syscall"

// Find the total memory in the guest OS
func getSystemTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	return uint64(in.Totalram) * uint64(in.Unit)
}
