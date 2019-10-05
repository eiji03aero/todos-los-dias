package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"

	"github.com/eiji03aero/todos-los-dias/pkg/label"
	"github.com/eiji03aero/todos-los-dias/pkg/label/cockroachdb"
	labelsvc "github.com/eiji03aero/todos-los-dias/pkg/label/implementation"
	"github.com/eiji03aero/todos-los-dias/pkg/label/transport"
	httptransport "github.com/eiji03aero/todos-los-dias/pkg/label/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8081", "HTTP listen address")
	)
	flag.Parse()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "label",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres",
			"postgresql://root@cockroach:26257/defaultdb?sslmode=disable")
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var svc label.Service
	{
		repository, err := cockroachdb.New(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = labelsvc.NewService(repository, logger)
		// svc = middleware.LoggingMiddleware(logger)(svc)
	}
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(svc)
	}
	var h http.Handler
	{
		h = httptransport.NewService(endpoints, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
