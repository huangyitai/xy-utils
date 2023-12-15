package srvx

import (
	"fmt"
	"time"

	"github.com/huangyitai/xy-utils/xxx"
	"github.com/spf13/cast"
)

// GetCgroupCPUAcctUsage ...
func GetCgroupCPUAcctUsage() (float64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/cpuacct/cpuacct.usage")
	if err != nil {
		return 0, err
	}

	if len(strs) == 0 {
		return 0, fmt.Errorf("GetCgroupCPUAcctUsage file is empty")
	}

	return cast.ToFloat64E(strs[0])
}

// GetCgroupCPUUsageFraction ...
func GetCgroupCPUUsageFraction(dur time.Duration) (float64, error) {
	usage1, err := GetCgroupCPUAcctUsage()
	if err != nil {
		return 0, err
	}
	start := time.Now()
	time.Sleep(dur)
	usage2, err := GetCgroupCPUAcctUsage()
	if err != nil {
		return 0, err
	}
	total := time.Since(start)
	if total == 0 {
		return 0, fmt.Errorf("duration is too short")
	}
	return (usage2 - usage1) / float64(total), nil
}

// GetCgroupCPUCfsPeriodUs ...
func GetCgroupCPUCfsPeriodUs() (int64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/cpuacct/cpu.cfs_period_us")
	if err != nil {
		return 0, err
	}

	if len(strs) == 0 {
		return 0, fmt.Errorf("GetCgroupCPUCfsPeriodUs file is empty")
	}

	return cast.ToInt64E(strs[0])
}

// GetCgroupCPUCfsQuotaUs ...
func GetCgroupCPUCfsQuotaUs() (int64, error) {
	strs, err := xxx.ReadLines("/sys/fs/cgroup/cpuacct/cpu.cfs_quota_us")
	if err != nil {
		return 0, err
	}

	if len(strs) == 0 {
		return 0, fmt.Errorf("GetCgroupCPUCfsQuotaUs file is empty")
	}

	return cast.ToInt64E(strs[0])
}

// GetCgroupCPUQuotaFraction ...
func GetCgroupCPUQuotaFraction() (float64, error) {
	period, err := GetCgroupCPUCfsPeriodUs()
	if err != nil {
		return 0, err
	}
	quota, err := GetCgroupCPUCfsQuotaUs()
	if err != nil {
		return 0, err
	}

	if quota == -1 {
		return 0, nil
	}

	return float64(quota) / float64(period), nil
}
