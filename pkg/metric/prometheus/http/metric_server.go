package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/log"
	gosoPrometheus "github.com/justtrackio/gosoline/pkg/metric/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Settings struct {
	Port    int             `cfg:"port" default:"8092"`
	Path    string          `cfg:"path" default:"/metrics"`
	Timeout TimeoutSettings `cfg:"timeout"`
}

// TimeoutSettings configures IO timeouts.
type TimeoutSettings struct {
	// You need to give at least 1s as timeout.
	// Read timeout is the maximum duration for reading the entire request, including the body.
	Read time.Duration `cfg:"read" default:"60s" validate:"min=1000000000"`
	// Write timeout is the maximum duration before timing out writes of the response.
	Write time.Duration `cfg:"write" default:"60s" validate:"min=1000000000"`
	// Idle timeout is the maximum amount of time to wait for the next request when keep-alives are enabled
	Idle time.Duration `cfg:"idle" default:"60s" validate:"min=1000000000"`
}

type metricsServer struct {
	kernel.EssentialModule
	kernel.ServiceStage

	logger   log.Logger
	server   *http.Server
	listener net.Listener
}

func New() kernel.ModuleFactory {
	return func(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
		settings := &Settings{}
		config.UnmarshalKey("prometheus.api", settings)

		registry, err := gosoPrometheus.ProvideRegistry(ctx, "default")
		if err != nil {
			return nil, err
		}

		return NewWithInterfaces(logger, registry, settings)
	}
}

func NewWithInterfaces(logger log.Logger, registry *prometheus.Registry, s *Settings) (*metricsServer, error) {
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		ReadTimeout:  s.Timeout.Read,
		WriteTimeout: s.Timeout.Write,
		IdleTimeout:  s.Timeout.Idle,
		Handler:      handler,
	}

	handler.Handle(s.Path, promhttp.InstrumentMetricHandler(
		registry, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
	))

	var err error
	var listener net.Listener
	address := server.Addr

	if address == "" {
		address = ":http"
	}

	// open a port for the server already in this step so we can already start accepting connections
	// when this module is later run (see also issue #201)
	if listener, err = net.Listen("tcp", address); err != nil {
		return nil, err
	}

	logger.Info("serving metrics on address %s", listener.Addr().String())

	return &metricsServer{
		logger:   logger,
		server:   server,
		listener: listener,
	}, nil
}

func (a *metricsServer) Run(ctx context.Context) error {
	go a.waitForStop(ctx)

	err := a.server.Serve(a.listener)

	if err != http.ErrServerClosed {
		a.logger.Error("Server closed unexpected: %w", err)

		return err
	}

	return nil
}

func (a *metricsServer) waitForStop(ctx context.Context) {
	<-ctx.Done()
	err := a.server.Close()
	if err != nil {
		a.logger.Error("Server Close: %w", err)
	}

	a.logger.Info("leaving metrics server")
}

func (a *metricsServer) GetPort() (*int, error) {
	if a == nil {
		return nil, errors.New("metricsServer is nil, module is not yet booted")
	}

	if a.listener == nil {
		return nil, errors.New("could not get port. module is not yet booted")
	}

	address := a.listener.Addr().String()
	_, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return nil, fmt.Errorf("could not get port from address %s: %w", address, err)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("can not convert port string to int: %w", err)
	}

	return &port, nil
}
