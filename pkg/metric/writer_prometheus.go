package metric

import (
	"context"
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/justtrackio/gosoline/pkg/appctx"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

const (
	UnitPromCounter   StandardUnit = "prom-counter"
	UnitPromGauge     StandardUnit = "prom-gauge"
	UnitPromHistogram StandardUnit = "prom-histogram"
	UnitPromSummary   StandardUnit = "prom-summary"
)

type (
	registryAppCtxKey   string
	promMetricFactory   func(*Datum) prometheus.Metric
	promMetricPersister func(metric prometheus.Metric, data *Datum)
)

func ProvideRegistry(ctx context.Context, name string) (*prometheus.Registry, error) {
	return appctx.Provide(ctx, registryAppCtxKey(name), func() (*prometheus.Registry, error) {
		registry := prometheus.NewRegistry()
		registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		registry.MustRegister(collectors.NewGoCollector())

		return registry, nil
	})
}

type promWriter struct {
	logger    log.Logger
	registry  *prometheus.Registry
	namespace string
}

func NewPromWriter(ctx context.Context, config cfg.Config, logger log.Logger) (*promWriter, error) {
	settings := &PromSettings{}
	config.UnmarshalKey(promSettingsKey, settings)

	appId := cfg.GetAppIdFromConfig(config)
	namespace := promNSNamingStrategy(appId)

	registry, err := ProvideRegistry(ctx, "default")
	if err != nil {
		return nil, err
	}

	return NewPromWriterWithInterfaces(logger, registry, namespace), nil
}

func NewPromWriterWithInterfaces(logger log.Logger, registry *prometheus.Registry, namespace string) *promWriter {
	return &promWriter{
		logger:    logger.WithChannel("metrics"),
		registry:  registry,
		namespace: namespace,
	}
}

func (w *promWriter) GetPriority() int {
	return PriorityLow
}

func (w *promWriter) Write(batch Data) {
	if len(batch) == 0 {
		return
	}

	for i := range batch {
		amendFromDefault(batch[i])

		if batch[i].Priority < w.GetPriority() {
			continue
		}
		w.promMetricFromDatum(batch[i])
	}

	w.logger.Debug("written %d metric data sets to prometheus", len(batch))
}

func (w *promWriter) WriteOne(data *Datum) {
	w.Write(Data{data})
}

func (w *promWriter) promMetricFromDatum(data *Datum) {
	switch data.Unit {
	case UnitCount:
		fallthrough
	case UnitPromCounter:
		w.promCounter(data)
	case UnitPromSummary:
		fallthrough
	case UnitMilliseconds:
		fallthrough
	case UnitSeconds:
		w.promSummary(data)
	case UnitPromHistogram:
		w.promHistogram(data)
	default:
		w.promGauge(data)
	}
}

func (w *promWriter) promMetric(data *Datum, metricFactory promMetricFactory, metricPersister promMetricPersister) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = fmt.Errorf("metric.prom: failed to write metric: %w", e)
			default:
				err = fmt.Errorf("metric.prom: failed to write metric: %v", e)
			}
			w.logger.WithFields(log.Fields{"error": err}).Error("failed to write metric")
		}
	}()
	metric := metricFactory(data)
	promMetric := metric.(prometheus.Metric)
	metricPersister(promMetric, data)
}

func (w *promWriter) promCounter(data *Datum) {
	counterFactory := func(datum *Datum) prometheus.Metric {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Namespace:   w.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
		if err := w.registry.Register(counter); err != nil {
			are := &prometheus.AlreadyRegisteredError{}
			if errors.As(err, are) {
				counter = are.ExistingCollector.(prometheus.Counter)
			} else {
				panic(err)
			}
		}
		return counter
	}

	counterPersister := func(metric prometheus.Metric, datum *Datum) {
		counterMetric := metric.(prometheus.Counter)
		counterMetric.Add(data.Value)
	}

	w.promMetric(data, counterFactory, counterPersister)
}

func (w *promWriter) promGauge(data *Datum) {
	gaugeFactory := func(datum *Datum) prometheus.Metric {
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace:   w.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
		if err := w.registry.Register(gauge); err != nil {
			are := &prometheus.AlreadyRegisteredError{}
			if errors.As(err, are) {
				gauge = are.ExistingCollector.(prometheus.Gauge)
			} else {
				panic(err)
			}
		}
		return gauge
	}

	gaugePersister := func(metric prometheus.Metric, datum *Datum) {
		gaugeMetric := metric.(prometheus.Gauge)
		gaugeMetric.Set(data.Value)
	}

	w.promMetric(data, gaugeFactory, gaugePersister)
}

func (w *promWriter) promSummary(data *Datum) {
	summaryFactory := func(datum *Datum) prometheus.Metric {
		summary := prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:   w.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
		if err := w.registry.Register(summary); err != nil {
			are := &prometheus.AlreadyRegisteredError{}
			if errors.As(err, are) {
				summary = are.ExistingCollector.(prometheus.Summary)
			} else {
				panic(err)
			}
		}
		return summary
	}

	summaryPersister := func(metric prometheus.Metric, datum *Datum) {
		summaryMetric := metric.(prometheus.Summary)
		summaryMetric.Observe(data.Value)
	}

	w.promMetric(data, summaryFactory, summaryPersister)
}

func (w *promWriter) promHistogram(data *Datum) {
	histogramFactory := func(datum *Datum) prometheus.Metric {
		histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace:   w.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
		if err := w.registry.Register(histogram); err != nil {
			are := &prometheus.AlreadyRegisteredError{}
			if errors.As(err, are) {
				histogram = are.ExistingCollector.(prometheus.Histogram)
			} else {
				panic(err)
			}
		}
		return histogram
	}

	histogramPersister := func(metric prometheus.Metric, datum *Datum) {
		histogramMetric := metric.(prometheus.Histogram)
		histogramMetric.Observe(data.Value)
	}

	w.promMetric(data, histogramFactory, histogramPersister)
}

func (w *promWriter) buildHelp(data *Datum) string {
	return fmt.Sprintf("unit: %s", data.Unit)
}
