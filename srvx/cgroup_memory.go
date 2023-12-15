package srvx

import (
	"fmt"
	"strings"

	"github.com/huangyitai/xy-utils/xxx"
	"github.com/spf13/cast"
)

// GetCgroupMemoryLimitInBytes ...
func GetCgroupMemoryLimitInBytes() (uint64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	if err != nil {
		return 0, err
	}

	if len(strs) == 0 {
		return 0, fmt.Errorf("GetCgroupMemoryLimitInBytes file is empty")
	}

	return cast.ToUint64E(strs[0])
}

// GetCgroupMemoryUsageInBytes ...
func GetCgroupMemoryUsageInBytes() (uint64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/memory/memory.usage_in_bytes")
	if err != nil {
		return 0, err
	}

	if len(strs) == 0 {
		return 0, fmt.Errorf("GetCgroupMemoryLimitInBytes file is empty")
	}

	return cast.ToUint64E(strs[0])
}

// GetCgroupMemoryWorkingSetInBytes 获取内存工作集大小（k8s OOM判据）= memory.usage_in_bytes - total_inactive_file (>=0)
func GetCgroupMemoryWorkingSetInBytes() (uint64, error) {
	usage, err := GetCgroupMemoryUsageInBytes()
	if err != nil {
		return 0, err
	}
	totalInactiveFile, err := GetCgroupMemoryTotalInactiveFileInBytes()
	if err != nil {
		return 0, err
	}
	if usage < totalInactiveFile {
		return 0, nil
	}
	return usage - totalInactiveFile, nil
}

// GetCgroupMemoryTotalInactiveFileInBytes ...
func GetCgroupMemoryTotalInactiveFileInBytes() (uint64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/memory/memory.stat")
	if err != nil {
		return 0, err
	}

	retStr := "0"
	for _, str := range strs {
		if !strings.HasPrefix(str, "total_inactive_file") {
			continue
		}
		splits := strings.Split(str, " ")
		if len(splits) != 2 {
			return 0, fmt.Errorf("GetMemoryWorkingSetInBytes total_inactive_file is empty")
		}
		retStr = splits[1]
	}
	return cast.ToUint64E(retStr)
}
