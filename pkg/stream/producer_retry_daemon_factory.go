package stream

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/log"
)

func ProducerRetryDaemonFactory(ctx context.Context, config cfg.Config, logger log.Logger) (map[string]kernel.ModuleFactory, error) {
	modules := map[string]kernel.ModuleFactory{}
	producerRetrySettings := readAllProducerRetrySettings(config)

	for name, settings := range producerRetrySettings {
		daemon, err := ProvideProducerRetryDaemon(ctx, config, logger, settings.Output, RetryMetadata{
			name:           name,
			retryConfigKey: ConfigurableConsumerRetryKey(name),
			retrySettings:  &settings.Retry,
		})
		if err != nil {
			return nil, fmt.Errorf("can not create producer daemon %s: %w", name, err)
		}

		moduleName := fmt.Sprintf("producer-daemon-%s", name)
		modules[moduleName] = func(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
			return daemon, nil
		}
	}

	return modules, nil
}

func readAllProducerRetrySettings(config cfg.Config) map[string]*ProducerSettings {
	producerSettings := make(map[string]*ProducerSettings)
	producerMap := config.GetStringMap("stream.producer", map[string]interface{}{})

	for name := range producerMap {
		settings := readProducerSettings(config, name)
		if !settings.Retry.Enabled {
			continue
		}

		producerSettings[name] = settings
	}

	return producerSettings
}
