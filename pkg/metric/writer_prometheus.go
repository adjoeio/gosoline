package metric

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/justtrackio/gosoline/pkg/mdl"
	gosoPrometheus "github.com/justtrackio/gosoline/pkg/metric/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	UnitPromCounter   StandardUnit = "prom-counter"
	UnitPromGauge     StandardUnit = "prom-gauge"
	UnitPromHistogram StandardUnit = "prom-histogram"
	UnitPromSummary   StandardUnit = "prom-summary"
)

type promMetricFactory func(*promDatum) prometheus.Metric
type promMetricPersister func(metric prometheus.Metric, data *promDatum)

type promSettings struct {
	// MetricLimit is used to avoid having metrics for which the name is programmatically generated (or have large number
	// of possible dimensions) which could lead in a memory leak.
	MetricLimit int64 `cfg:"metric_limit" default:"10000"`
}

type promDatum struct {
	*Datum
	namespace string
}

type promWriter struct {
	logger      log.Logger
	clock       clock.Clock
	promMetrics sync.Map
	registry    *prometheus.Registry
	namespace   string
	metricLimit int64
	metrics     *int64
}

func NewPromWriter(ctx context.Context, config cfg.Config, logger log.Logger) (*promWriter, error) {
	settings := &promSettings{}
	config.UnmarshalKey("metric.prometheus", settings)

	appId := cfg.GetAppIdFromConfig(config)
	namespace := fmt.Sprintf("%s:%s:%s", appId.Project, appId.Environment, appId.Family)

	registry, err := gosoPrometheus.ProvideRegistry(ctx, "default")
	if err != nil {
		return nil, err
	}

	return NewPromWriterWithInterfaces(logger, registry, namespace, settings.MetricLimit), nil
}

func NewPromWriterWithInterfaces(logger log.Logger, registry *prometheus.Registry, namespace string, metricLimit int64) *promWriter {
	return &promWriter{
		logger:      logger.WithChannel("metrics"),
		registry:    registry,
		namespace:   namespace,
		metricLimit: metricLimit,
		metrics:     mdl.Int64(0),
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
		if batch[i].Priority < w.GetPriority() {
			continue
		}
		w.promMetricFromDatum(&promDatum{
			Datum:     batch[i],
			namespace: w.namespace,
		})
	}

	w.logger.Debug("written %d metric data sets to prometheus", len(batch))
}

func (w *promWriter) WriteOne(data *Datum) {
	w.Write(Data{data})
}

func (w *promWriter) promMetricFromDatum(data *promDatum) {
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

func (w *promWriter) promMetric(data *promDatum, metricFactory promMetricFactory, metricPersister promMetricPersister) {
	metric, ok := w.promMetrics.Load(data.Id())
	if !ok {
		var err error
		metric, err = w.addMetric(metricFactory, data)
		if err != nil {
			return
		}
	}

	promMetric := metric.(prometheus.Metric)
	metricPersister(promMetric, data)
}

func (w *promWriter) promCounter(data *promDatum) {
	counterFactory := func(datum *promDatum) prometheus.Metric {
		return promauto.With(w.registry).NewCounter(prometheus.CounterOpts{
			Namespace:   data.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
	}

	counterPersister := func(metric prometheus.Metric, datum *promDatum) {
		counterMetric := metric.(prometheus.Counter)
		counterMetric.Add(data.Value)
	}

	w.promMetric(data, counterFactory, counterPersister)
}

func (w *promWriter) promGauge(data *promDatum) {
	gaugeFactory := func(datum *promDatum) prometheus.Metric {
		return promauto.With(w.registry).NewGauge(prometheus.GaugeOpts{
			Namespace:   data.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
	}

	gaugePersister := func(metric prometheus.Metric, datum *promDatum) {
		gaugeMetric := metric.(prometheus.Gauge)
		gaugeMetric.Set(data.Value)
	}

	w.promMetric(data, gaugeFactory, gaugePersister)
}

func (w *promWriter) promSummary(data *promDatum) {
	summaryFactory := func(datum *promDatum) prometheus.Metric {
		return promauto.With(w.registry).NewSummary(prometheus.SummaryOpts{
			Namespace:   data.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
	}

	summaryPersister := func(metric prometheus.Metric, datum *promDatum) {
		summaryMetric := metric.(prometheus.Summary)
		summaryMetric.Observe(data.Value)
	}

	w.promMetric(data, summaryFactory, summaryPersister)
}

func (w *promWriter) promHistogram(data *promDatum) {
	histogramFactory := func(datum *promDatum) prometheus.Metric {
		return promauto.With(w.registry).NewHistogram(prometheus.HistogramOpts{
			Namespace:   data.namespace,
			Name:        data.MetricName,
			Help:        w.buildHelp(data),
			ConstLabels: prometheus.Labels(data.Dimensions),
		})
	}

	histogramPersister := func(metric prometheus.Metric, datum *promDatum) {
		histogramMetric := metric.(prometheus.Histogram)
		histogramMetric.Observe(data.Value)
	}

	w.promMetric(data, histogramFactory, histogramPersister)
}

func (w *promWriter) addMetric(metricFactory promMetricFactory, data *promDatum) (prometheus.Metric, error) {
	if atomic.LoadInt64(w.metrics) >= w.metricLimit {
		w.logger.Error("fail to write metric due to exceeding limit")
		return nil, errors.New("metric limit exceeded")
	}

	metric := metricFactory(data)
	w.promMetrics.Store(data.Id(), metric)
	atomic.AddInt64(w.metrics, 1)

	return metric, nil
}

func (w *promWriter) buildHelp(data *promDatum) string {
	return fmt.Sprintf("unit: %s", data.Unit)
}
