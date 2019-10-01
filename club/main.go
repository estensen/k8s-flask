package main

import (
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	closer, err := initJaeger("slowHello")
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
	}
	defer closer.Close()
	tracer := opentracing.GlobalTracer()

	span := tracer.StartSpan("start-service")
	span.LogKV("started", "yay!")
	span.Finish()

	http.HandleFunc("/club", func(w http.ResponseWriter, r *http.Request) {
		span := tracer.StartSpan("hello")
		span.SetTag("slowHello", "hello")
		defer span.Finish()

		time.Sleep(time.Second)

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		data := []byte("Hello book club!")

		log.Println("Hello")
		span.LogKV("slow-hello", "hello")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	log.Println("club running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initJaeger(service string) (io.Closer, error) {
	// Example from
	// https://github.com/jaegertracing/jaeger-client-go/blob/6c4e7ad7e7cd2960d5e6d918c4f3d56d263bf26d/config/example_test.go#L29
	// Added LocalAgentHostPort to make it play nice with docker-compose
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	jLogger := jaegerlog.StdLogger
	MetricsFactory := prometheus.New()

	closer, err := cfg.InitGlobalTracer(
		service,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(MetricsFactory),
	)
	if err != nil {
		return nil, err
	}

	return closer, err
}
