// +build !linux

package cri

// Find the total memory in the guest OS
func getSystemTotalMemory() uint64 {
	return 0
}
