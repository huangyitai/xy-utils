package srvx

import (
	"context"
	"time"

	"github.com/huangyitai/xy-utils/dox"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// GetSummary ...
func GetSummary() (*Summary, error) {
	return GetSummaryWithDuration(2 * time.Second)
}

// GetSummaryWithDuration ...
func GetSummaryWithDuration(dur time.Duration) (*Summary, error) {
	return GetSummaryWithDurations(dur, dur, dur)
}

// GetSummaryWithDurations ...
func GetSummaryWithDurations(cpuDur time.Duration, netDur time.Duration, diskDur time.Duration) (*Summary, error) {
	res := &Summary{}
	err := dox.RunWithContext(func(ctx context.Context) error {
		s, err := GetCPUSummary(cpuDur)
		if err != nil {
			return err
		}
		res.CPU = s
		return nil
	}).While(func(ctx context.Context) error {
		s, err := GetNetSummary(netDur)
		if err != nil {
			return err
		}
		res.Net = s
		return nil
	}).While(func(ctx context.Context) error {
		s, err := GetDiskSummary(diskDur)
		if err != nil {
			return err
		}
		res.Disk = s
		return nil
	}).Then(func(ctx context.Context) error {
		s, err := GetMemorySummary()
		if err != nil {
			return err
		}
		res.Memory = s
		return nil
	})(context.Background())
	res.CollectedAt = time.Now()
	return res, err
}

// GetCPUSummary ...
func GetCPUSummary(dur time.Duration) (*CPUSummary, error) {
	res := &CPUSummary{}
	if dur <= 0 {
		return res, nil
	}

	err := dox.RunWithContext(func(ctx context.Context) error {
		usage, err := GetCgroupCPUUsageFraction(dur)
		if err != nil {
			return err
		}
		res.UsageFraction = usage
		return nil
	}).While(func(ctx context.Context) error {
		pers, err := cpu.Percent(time.Second, false)
		if err != nil {
			return err
		}
		res.UsagePercent = pers[0]
		return nil
	})(context.Background())
	if err != nil {
		return res, err
	}

	quota, err := GetCgroupCPUQuotaFraction()
	if err != nil {
		return res, err
	}
	res.QuotaFraction = quota

	if quota != 0 {
		res.QuotaUsagePercent = res.UsageFraction / res.QuotaFraction * 100
	}
	return res, nil
}

// GetMemorySummary ...
func GetMemorySummary() (*MemorySummary, error) {
	res := &MemorySummary{}
	limit, err := GetCgroupMemoryLimitInBytes()
	if err != nil {
		return res, err
	}

	usage, err := GetCgroupMemoryWorkingSetInBytes()
	if err != nil {
		return res, err
	}
	res.UsageInBytes = usage

	if limit != 1<<63-1 {
		res.UsagePercent = float64(usage) / float64(limit) * 100
		res.LimitInBytes = limit
	}

	vms, err := mem.VirtualMemory()
	if err != nil {
		return res, err
	}
	res.TotalInBytes = vms.Total
	res.UsedInBytes = vms.Used
	res.AvailableInBytes = vms.Available
	res.UsedPercent = vms.UsedPercent
	return res, nil
}

// GetNetSummary ...
func GetNetSummary(dur time.Duration) (*NetSummary, error) {
	res := &NetSummary{}
	if dur <= 0 {
		return res, nil
	}

	startCnt, err := net.IOCounters(true)
	if err != nil {
		return res, nil
	}
	time.Sleep(dur)
	endCnt, err := net.IOCounters(true)
	cntMap := map[string]*net.IOCountersStat{}
	for i, start := range startCnt {
		cntMap[start.Name] = &startCnt[i]
	}
	for _, end := range endCnt {
		start := cntMap[end.Name]
		if start == nil {
			continue
		}
		stat := &NetIOStats{Name: end.Name}
		stat.SentBytesPerSecond = float64(end.BytesSent-start.BytesSent) / dur.Seconds()
		stat.RecvBytesPerSecond = float64(end.BytesRecv-start.BytesRecv) / dur.Seconds()
		stat.SentPacketsPerSecond = float64(end.PacketsSent-start.PacketsSent) / dur.Seconds()
		stat.RecvPacketsPerSecond = float64(end.PacketsRecv-start.PacketsRecv) / dur.Seconds()
		res.IOStats = append(res.IOStats, stat)
	}
	return res, nil
}

// GetDiskSummary ...
func GetDiskSummary(dur time.Duration) (*DiskSummary, error) {
	res := &DiskSummary{}
	if dur > 0 {
		startCnt, err := disk.IOCounters()
		if err != nil {
			return res, err
		}
		time.Sleep(dur)
		endCnt, err := disk.IOCounters()
		if err != nil {
			return res, err
		}
		for name, end := range endCnt {
			start, ok := startCnt[name]
			if !ok {
				continue
			}
			stat := &DiskIOStats{Name: name}
			stat.ReadBytesPerSecond = float64(end.ReadBytes-start.ReadBytes) / dur.Seconds()
			stat.WriteBytesPerSecond = float64(end.WriteBytes-start.WriteBytes) / dur.Seconds()
			stat.ReadOpsPerSecond = float64(end.ReadCount-start.ReadCount) / dur.Seconds()
			stat.WriteOpsPerSecond = float64(end.WriteCount-start.WriteCount) / dur.Seconds()
			stat.UsagePercent = float64(end.IoTime-start.IoTime) / float64(dur.Milliseconds()) * 100
			res.IOStats = append(res.IOStats, stat)
		}
	}

	parts, err := disk.Partitions(false)
	if err != nil {
		return res, err
	}
	u, err := disk.Usage("/")
	if err != nil {
		return res, err
	}
	res.Usage = append(res.Usage, calcDiskUsage(u))
	for _, part := range parts {
		if part.Mountpoint == "/" {
			continue
		}
		u, err := disk.Usage(part.Mountpoint)
		if err != nil {
			continue
		}
		res.Usage = append(res.Usage, calcDiskUsage(u))
	}
	return res, nil
}

func calcDiskUsage(stat *disk.UsageStat) *DiskUsage {
	return &DiskUsage{
		MountedOn:         stat.Path,
		Total:             stat.Total,
		Free:              stat.Free,
		Used:              stat.Used,
		UsedPercent:       stat.UsedPercent,
		InodesTotal:       stat.InodesTotal,
		InodesUsed:        stat.InodesUsed,
		InodesFree:        stat.InodesFree,
		InodesUsedPercent: stat.InodesUsedPercent,
	}
}
