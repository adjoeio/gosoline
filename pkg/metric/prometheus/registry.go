package prometheus

import (
	"context"
	"sync"

	"github.com/justtrackio/gosoline/pkg/appctx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var registryLock sync.Mutex

type registryAppCtxKey string

func ProvideRegistry(ctx context.Context, name string) (*prometheus.Registry, error) {
	registryLock.Lock()
	defer registryLock.Unlock()

	return appctx.Provide(ctx, registryAppCtxKey(name), func() (*prometheus.Registry, error) {
		registry := prometheus.NewRegistry()
		registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		registry.MustRegister(collectors.NewGoCollector())

		return registry, nil
	})
}
