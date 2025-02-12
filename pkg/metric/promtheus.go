package metric

import (
	"context"
	"fmt"
	"sync"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/kernel/common"
	"github.com/justtrackio/gosoline/pkg/log"
)

type singleMetricsModule struct {
	sync.Mutex
	logger log.Logger

	channel *metricChannel
	writer  Writer
	enabled bool
}

func NewSingleMetricsModule(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
	var metricWriter Writer
	var err error

	settings := getMetricSettings(config)

	channel := ProviderMetricChannel()
	channel.enabled = settings.Enabled
	channel.logger = logger.WithChannel("metrics")

	if !settings.Enabled {
		return NewMetricDaemonWithInterfaces(logger, channel, newNoopWriter(), settings)
	}

	metricWriter, err = NewMetricWriter(ctx, config, logger, settings.Writer)
	if err != nil {
		return nil, fmt.Errorf("can not create metric writer: %w", err)
	}

	return NewSingleMetricsModuleWithInterfaces(logger, channel, metricWriter, settings.Enabled)
}

func NewSingleMetricsModuleWithInterfaces(logger log.Logger, channel *metricChannel, writer Writer, enabled bool) (kernel.Module, error) {
	return &singleMetricsModule{
		logger:  logger.WithChannel("metrics"),
		channel: channel,
		writer:  writer,
		enabled: enabled,
	}, nil
}

func (p *singleMetricsModule) IsEssential() bool {
	return false
}

func (p *singleMetricsModule) IsBackground() bool {
	return true
}

func (p *singleMetricsModule) GetStage() int {
	return common.StageEssential
}

func (p *singleMetricsModule) Run(ctx context.Context) error {
	if !p.enabled {
		p.logger.Info("metrics not enabled")

		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case data := <-p.channel.c:
			p.writer.Write(data)
		}
	}
}
