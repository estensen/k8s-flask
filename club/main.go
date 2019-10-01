package main

import (
	"github.com/uber/jaeger-lib/metrics/prometheus"
	//"github.com/uber/jaeger-client-go"
	//"github.com/uber/jaeger-lib/metrics/prometheus"
	"io"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	closer, err := initJaeger()
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

func initJaeger() (io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, err
	}
	cfg.ServiceName = "club"

	jLogger := jaegerlog.StdLogger
	MetricsFactory := prometheus.New()

	closer, err := cfg.InitGlobalTracer(
		cfg.ServiceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(MetricsFactory),
	)
	if err != nil {
		return nil, err
	}

	return closer, err
}
