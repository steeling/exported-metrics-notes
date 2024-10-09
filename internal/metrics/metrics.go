package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// RequestCount tracks the number of requests received.
	RequestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:      "proxy_filter_requests_total",
			Help:      "Total number of HTTP requests",
			Namespace: "proxy_filter",
		},
	)
	// ProxyErrorCount tracks the number of proxy errors encountered.
	ProxyErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "proxy_filter_errors_total",
			Help:      "Total number of vmagent-proxy errors",
			Namespace: "proxy_filter",
		},
		[]string{"code"},
	)
)

func init() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ProxyErrorCount)
}
