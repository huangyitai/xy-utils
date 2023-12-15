package srvx

import (
	"context"

	"github.com/huangyitai/xy-utils/metricx"
	"github.com/rs/zerolog/log"
)

// ReportMetric 报告上报监控
func (s *Summary) ReportMetric(name string) {
	ctx := metricx.NewByName(name).
		Tag(TagNameEnv, GetEnvironment()).
		Tag(TagNameService, GetService()).
		Tag(TagNameInstance, GetInstance()).
		WithContext(context.Background())
	s.ReportMetricWithContext(ctx)
}

// ReportMetricWithContext 报告根据context上报
func (s *Summary) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	s.CPU.ReportMetricWithContext(ctx)
	s.Memory.ReportMetricWithContext(ctx)
	s.Net.ReportMetricWithContext(ctx)
	s.Disk.ReportMetricWithContext(ctx)
}

// ReportMetricWithContext 报告根据context上报
func (s *CPUSummary) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	err := metricx.Ctx(ctx).Report().
		MaxCounter(MetricNameCPUUsageFraction, s.UsageFraction).
		MaxCounter(MetricNameCPUQuotaFraction, s.QuotaFraction).
		MaxCounter(MetricNameCPUQuotaUsagePercent, s.QuotaUsagePercent).
		MaxCounter(MetricNameCPUUsagePercent, s.UsagePercent).
		Send()
	if err != nil {
		log.Err(err).Msg("[Summary][CPUSummary] Report fail")
	}
}

// ReportMetricWithContext ...
func (s *MemorySummary) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	err := metricx.Ctx(ctx).Report().
		MaxCounter(MetricNameMemoryLimitInBytes, float64(s.LimitInBytes)).
		MaxCounter(MetricNameMemoryUsageInBytes, float64(s.UsageInBytes)).
		MaxCounter(MetricNameMemoryUsagePercent, s.UsagePercent).
		MaxCounter(MetricNameMemoryTotalInBytes, float64(s.TotalInBytes)).
		MaxCounter(MetricNameMemoryUsedInBytes, float64(s.UsedInBytes)).
		MaxCounter(MetricNameMemoryAvailableInBytes, float64(s.AvailableInBytes)).
		MaxCounter(MetricNameMemoryUsedPercent, s.UsedPercent).
		Send()
	if err != nil {
		log.Err(err).Msg("[Summary][MemorySummary] Report fail")
	}
}

// ReportMetricWithContext ...
func (s *NetSummary) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	for _, stat := range s.IOStats {
		stat.ReportMetricWithContext(ctx)
	}
}

// ReportMetricWithContext ...
func (s *NetIOStats) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	err := metricx.Ctx(ctx).Report().
		Tag(TagNameNetInterfaceName, s.Name).
		MaxCounter(MetricNameNetIOSentBytesPerSecond, s.SentBytesPerSecond).
		MaxCounter(MetricNameNetIORecvBytesPerSecond, s.RecvBytesPerSecond).
		MaxCounter(MetricNameNetIOSentPacketsPerSecond, s.SentPacketsPerSecond).
		MaxCounter(MetricNameNetIORecvPacketsPerSecond, s.RecvPacketsPerSecond).
		Send()
	if err != nil {
		log.Err(err).Msg("[Summary][NetIOStats] Report fail")
	}
}

// ReportMetricWithContext ...
func (s *DiskSummary) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	for _, usage := range s.Usage {
		usage.ReportMetricWithContext(ctx)
	}
	for _, stat := range s.IOStats {
		stat.ReportMetricWithContext(ctx)
	}
}

// ReportMetricWithContext ...
func (u *DiskUsage) ReportMetricWithContext(ctx context.Context) {
	if u == nil {
		return
	}
	err := metricx.Ctx(ctx).Report().
		Tag(TagNameDiskMountedOn, u.MountedOn).
		MaxCounter(MetricNameDiskTotal, float64(u.Total)).
		MaxCounter(MetricNameDiskFree, float64(u.Free)).
		MaxCounter(MetricNameDiskUsed, float64(u.Used)).
		MaxCounter(MetricNameDiskUsedPercent, u.UsedPercent).
		MaxCounter(MetricNameDiskInodesTotal, float64(u.InodesTotal)).
		MaxCounter(MetricNameDiskInodesUsed, float64(u.InodesUsed)).
		MaxCounter(MetricNameDiskInodesFree, float64(u.InodesFree)).
		MaxCounter(MetricNameDiskInodesUsedPercent, u.InodesUsedPercent).
		Send()
	if err != nil {
		log.Err(err).Msg("[Summary][DiskUsage] Report fail")
	}
}

// ReportMetricWithContext ...
func (s *DiskIOStats) ReportMetricWithContext(ctx context.Context) {
	if s == nil {
		return
	}
	err := metricx.Ctx(ctx).Report().
		Tag(TagNameDiskDevice, s.Name).
		MaxCounter(MetricNameDiskIOReadOpsPerSecond, s.ReadOpsPerSecond).
		MaxCounter(MetricNameDiskIOWriteOpsPerSecond, s.WriteOpsPerSecond).
		MaxCounter(MetricNameDiskIOReadBytesPerSecond, s.ReadBytesPerSecond).
		MaxCounter(MetricNameDiskIOWriteBytesPerSecond, s.WriteBytesPerSecond).
		MaxCounter(MetricNameDiskIOUsagePercent, s.UsagePercent).
		Send()
	if err != nil {
		log.Err(err).Msg("[Summary][DiskIOStats] Report fail")
	}
}
