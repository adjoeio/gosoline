package producer

import (
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"

	"github.com/justtrackio/gosoline/pkg/kafka"
	"github.com/justtrackio/gosoline/pkg/kafka/connection"
)

type Settings struct {
	ConnectionName string `cfg:"connection" validate:"required"`
	Topic          string `cfg:"topic" validate:"required"`
	// FQTopic is the fully-qualified topic name (with prefixes).
	FQTopic      string
	BatchSize    int           `cfg:"batch_size"`
	BatchTimeout time.Duration `cfg:"idle_timeout"`
	AsyncWrites  bool          `cfg:"async_writes"`
	Balancer     string        `cfg:"balancer" default:"default"`
	Retries      int           `cfg:"retries" default:"5"`
	WriteTimeout time.Duration `cfg:"write_timeout" default:"30s"`
	// Enables debug logs for segmentio/kafka module.
	DebugLogs  bool `cfg:"debug_logs" default:"false"`
	connection *connection.Settings
}

func (s *Settings) Connection() *connection.Settings {
	return s.connection
}

func (s *Settings) WithConnection(conn *connection.Settings) *Settings {
	s.connection = conn
	return s
}

func ParseSettings(c cfg.Config, key string) *Settings {
	settings := &Settings{}
	c.UnmarshalKey(key, settings)

	settings.connection = connection.ParseSettings(c, settings.ConnectionName)
	settings.FQTopic = kafka.FQTopicName(cfg.GetAppIdFromConfig(c), settings.Topic)

	return settings
}
