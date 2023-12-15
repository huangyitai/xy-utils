package srvx

// DefaultMetricName TODO
var DefaultMetricName = "Srv.Default"

// SystemDevMetricName TODO
const (
	SystemDevMetricName  = "Srv.System.Dev"
	SystemPreMetricName  = "Srv.System.Pre"
	SystemProdMetricName = "Srv.System.Prod"

	TagNameEnv      = "env"
	TagNameService  = "service"
	TagNameInstance = "instance"

	TagNameNetInterfaceName = "net_interface"
	TagNameDiskMountedOn    = "disk_mounted_on"
	TagNameDiskDevice       = "disk_device"

	MetricNameCPUUsageFraction     = "cpu_usage_fraction"
	MetricNameCPUQuotaFraction     = "cpu_quota_fraction"
	MetricNameCPUQuotaUsagePercent = "cpu_quota_usage_ratio"
	MetricNameCPUUsagePercent      = "cpu_usage_ratio"

	MetricNameMemoryLimitInBytes     = "memory_limit"
	MetricNameMemoryUsageInBytes     = "memory_usage"
	MetricNameMemoryUsagePercent     = "memory_usage_ratio"
	MetricNameMemoryTotalInBytes     = "memory_total"
	MetricNameMemoryUsedInBytes      = "memory_used"
	MetricNameMemoryAvailableInBytes = "memory_available"
	MetricNameMemoryUsedPercent      = "memory_used_ratio"

	MetricNameNetIOSentBytesPerSecond   = "net_sent_bps"
	MetricNameNetIORecvBytesPerSecond   = "net_recv_bps"
	MetricNameNetIOSentPacketsPerSecond = "net_sent_pps"
	MetricNameNetIORecvPacketsPerSecond = "net_recv_pps"

	MetricNameDiskTotal             = "disk_total"
	MetricNameDiskFree              = "disk_free"
	MetricNameDiskUsed              = "disk_used"
	MetricNameDiskUsedPercent       = "disk_used_ratio"
	MetricNameDiskInodesTotal       = "disk_inode_total"
	MetricNameDiskInodesUsed        = "disk_inode_used"
	MetricNameDiskInodesFree        = "disk_inode_free"
	MetricNameDiskInodesUsedPercent = "disk_inode_used_ratio"

	MetricNameDiskIOReadOpsPerSecond    = "disk_read_iops"
	MetricNameDiskIOWriteOpsPerSecond   = "disk_write_iops"
	MetricNameDiskIOReadBytesPerSecond  = "disk_read_bps"
	MetricNameDiskIOWriteBytesPerSecond = "disk_write_bps"
	MetricNameDiskIOUsagePercent        = "disk_io_usage_ratio"
)
