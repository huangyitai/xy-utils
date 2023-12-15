package srvx

import "github.com/rs/zerolog/log"

func (s *Starter) job10() {
	summary, err := GetSummary()
	if err != nil {
		log.Err(err).Caller().Msg("[srvx/Starter] GetSummary error")
		return
	}
	if s.enableSummaryLog {
		summary.Log()
	}
	if s.enableReportMetric {
		metricName := s.getMetricName()
		summary.ReportMetric(metricName)
	}
}
