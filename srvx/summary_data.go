package srvx

import "time"

// Summary ...
type Summary struct {
	CollectedAt time.Time
	CPU         *CPUSummary
	Memory      *MemorySummary
	Net         *NetSummary
	Disk        *DiskSummary
}

// CPUSummary ...
type CPUSummary struct {
	UsageFraction     float64
	QuotaFraction     float64
	QuotaUsagePercent float64
	UsagePercent      float64
}

// MemorySummary ...
type MemorySummary struct {
	LimitInBytes     uint64
	UsageInBytes     uint64
	UsagePercent     float64
	TotalInBytes     uint64
	UsedInBytes      uint64
	AvailableInBytes uint64
	UsedPercent      float64
}

// NetIOStats ...
type NetIOStats struct {
	Name                 string
	SentBytesPerSecond   float64
	RecvBytesPerSecond   float64
	SentPacketsPerSecond float64
	RecvPacketsPerSecond float64
}

// NetSummary ...
type NetSummary struct {
	IOStats []*NetIOStats
}

// DiskSummary ...
type DiskSummary struct {
	Usage   []*DiskUsage
	IOStats []*DiskIOStats
}

// DiskUsage ...
type DiskUsage struct {
	MountedOn         string
	Total             uint64
	Free              uint64
	Used              uint64
	UsedPercent       float64
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
}

// DiskIOStats ...
type DiskIOStats struct {
	Name                string
	ReadOpsPerSecond    float64
	WriteOpsPerSecond   float64
	ReadBytesPerSecond  float64
	WriteBytesPerSecond float64
	UsagePercent        float64
}
