package tracing

import (
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func CreateOpenTracingTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	// override config from environment variables: https://github.com/jaegertracing/jaeger-client-go
	cfg, err := cfg.FromEnv()
	if err != nil {
		log.Panic("Couldn't override Jaeger configuration from environment", err)
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaegerlog.StdLogger)) // this StdLogger has debug enabled
	if err != nil {
		log.Panic("Couldn't create Jaeger tracer", err)
	}

	return tracer, closer
}
