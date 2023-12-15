package srvx

import (
	"context"

	"github.com/huangyitai/xy-utils/logx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogKeySummary ...
var (
	LogKeySummary = "@summary"
	LogKeyCollect = "@collect"
)

// Log ...
func (s *Summary) Log() {
	s.LogWithContext(context.Background())
}

// LogWithContext ...
func (s *Summary) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	ctx = logx.WithSubLogger(ctx, func(c zerolog.Context) zerolog.Context {
		return c.Str(LogKeyCollect, s.CollectedAt.Format("2006-01-02 15:04:05"))
	})
	s.CPU.LogWithContext(ctx)
	s.Memory.LogWithContext(ctx)
	s.Net.LogWithContext(ctx)
	s.Disk.LogWithContext(ctx)
}

// LogWithContext ...
func (s *CPUSummary) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	log.Ctx(ctx).Info().Str(LogKeySummary, "CPU").
		Float64("fUsageFraction", s.UsageFraction).
		Float64("fQuotaFraction", s.QuotaFraction).
		Float64("fQuotaUsagePercent", s.QuotaUsagePercent).
		Float64("fUsagePercent", s.UsagePercent).
		Msg("[CPUSummary] report")
}

// LogWithContext ...
func (s *MemorySummary) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}

	log.Ctx(ctx).Info().Str(LogKeySummary, "Memory").
		Uint64("iLimitInBytes", s.LimitInBytes).
		Uint64("iUsageInBytes", s.UsageInBytes).
		Float64("fUsagePercent", s.UsagePercent).
		Uint64("iTotalInBytes", s.TotalInBytes).
		Uint64("iUsedInBytes", s.UsedInBytes).
		Uint64("iAvailableInBytes", s.AvailableInBytes).
		Float64("fUsedPercent", s.UsedPercent).
		Msg("[MemorySummary] report")
}

// LogWithContext ...
func (s *NetSummary) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	for _, stat := range s.IOStats {
		stat.LogWithContext(ctx)
	}
}

// LogWithContext ...
func (s *NetIOStats) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}

	log.Ctx(ctx).Info().Str(LogKeySummary, "NetIOStats").
		Str("sInterface", s.Name).
		Float64("fSentBytesPerSecond", s.SentBytesPerSecond).
		Float64("fRecvBytesPerSecond", s.RecvBytesPerSecond).
		Float64("fSentPacketsPerSecond", s.SentPacketsPerSecond).
		Float64("fRecvPacketsPerSecond", s.RecvPacketsPerSecond).
		Msg("[NetIOStats] report")
}

// LogWithContext ...
func (s *DiskSummary) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}

	for _, usage := range s.Usage {
		usage.LogWithContext(ctx)
	}
	for _, stat := range s.IOStats {
		stat.LogWithContext(ctx)
	}
}

// LogWithContext ...
func (u *DiskUsage) LogWithContext(ctx context.Context) {
	if u == nil {
		return
	}

	log.Ctx(ctx).Info().Str(LogKeySummary, "DiskUsage").
		Str("sMountedOn", u.MountedOn).
		Uint64("iTotal", u.Total).
		Uint64("iFree", u.Free).
		Uint64("iUsed", u.Used).
		Float64("fUsedPercent", u.UsedPercent).
		Uint64("iInodesTotal", u.InodesTotal).
		Uint64("iInodesUsed", u.InodesUsed).
		Uint64("iInodesFree", u.InodesFree).
		Float64("fInodesUsedPercent", u.InodesUsedPercent).
		Msg("[DiskUsage] report")
}

// LogWithContext ...
func (s *DiskIOStats) LogWithContext(ctx context.Context) {
	if s == nil {
		return
	}

	log.Ctx(ctx).Info().Str(LogKeySummary, "DiskIOStats").
		Str("sDevice", s.Name).
		Float64("fReadOpsPerSecond", s.ReadOpsPerSecond).
		Float64("fWriteOpsPerSecond", s.WriteOpsPerSecond).
		Float64("fReadBytesPerSecond", s.ReadBytesPerSecond).
		Float64("fWriteBytesPerSecond", s.WriteBytesPerSecond).
		Float64("fUsagePercent", s.UsagePercent).
		Msg("[DiskIOStats] report")
}
