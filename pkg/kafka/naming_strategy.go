package kafka

import (
	"fmt"
	"strings"

	"github.com/justtrackio/gosoline/pkg/cfg"
)

type NamingStrategy func(appId cfg.AppId, name string) string

func NormalizeKafkaName(name string) string {
	return strings.ReplaceAll(name, "_", "-")
}

// FQTopicName returns fully-qualified topic name.
func FQTopicName(appId cfg.AppId, topic string) string {
	return NormalizeKafkaName(topicNamingStrategy(appId, topic))
}

func FQGroupId(appId cfg.AppId, groupId string) string {
	return groupNamingStrategy(appId, groupId)
}

func DefaultTopicNamingStrategy(appId cfg.AppId, name string) string {
	return fmt.Sprintf("%s-%s", appId.Environment, name)
}

func DefaultGroupIdNamingStrategy(appId cfg.AppId, name string) string {
	// Maintain legacy behavior when group id isn't specified.
	if name == "" {
		return appId.Application
	}

	return fmt.Sprintf("%s-%s-%s", appId.Environment, appId.Application, name)
}

func WithTopicNamingStrategy(strategy NamingStrategy) {
	topicNamingStrategy = strategy
}

func WithGroupIdNamingStrategy(strategy NamingStrategy) {
	groupNamingStrategy = strategy
}

var (
	topicNamingStrategy = DefaultTopicNamingStrategy
	groupNamingStrategy = DefaultGroupIdNamingStrategy
)
