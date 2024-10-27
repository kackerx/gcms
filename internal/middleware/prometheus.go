package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "http_requests_total",
		Help:        "Total number of http request",
		ConstLabels: nil,
	}, []string{"method", "path"})

	requestCodeTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "http_requests_code_total",
		Help:        "code of http request",
		ConstLabels: nil,
	}, []string{"method", "path", "code"})

	requestDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "",
		Subsystem: "",
		Name:      "http_requests_duration_seconds",
		Help:      "http requests per seconds",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
		ConstLabels: nil,
		MaxAge:      0,
		AgeBuckets:  0,
		BufCap:      0,
	}, []string{"method", "path"})
)

func init() {
	prometheus.MustRegister(requestCounter, requestDuration, requestCodeTotal)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		requestCounter.WithLabelValues(method, path).Inc()

		c.Next()

		requestDuration.WithLabelValues(method, path).Observe(time.Since(start).Seconds())
		statusCode := c.Writer.Status()
		requestCodeTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
	}
}
