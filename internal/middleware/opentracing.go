package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
)

var (
	report    = reporter.NewReporter("http://localhost:9411/api/v2/spans")
	zipTracer opentracing.Tracer
)

func init() {
	endpoint, err := zipkin.NewEndpoint("gcms", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	// 创建tracer
	tracer, err := zipkin.NewTracer(
		report,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithTraceID128Bit(true),
	)
	if err != nil {
		panic(err)
	}

	zipTracer = zipkintracer.Wrap(tracer)
	opentracing.SetGlobalTracer(zipTracer)
}

// OpentracingMiddleware 上报地址 :9411/api/v2/spans
func OpentracingMiddleware() gin.HandlerFunc {
	// 创建Reporter
	// 创建本地节点
	// 创建Zipkin Tracer

	return ginhttp.Middleware(zipTracer, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	}))
}
