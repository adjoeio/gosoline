package kafka_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/kafka"
	"github.com/justtrackio/gosoline/pkg/test/assert"
)

func Test_DefaultTopicNamingStrategy(t *testing.T) {
	fqName := kafka.DefaultTopicNamingStrategy(cfg.AppId{Project: "prog", Environment: "prod", Application: "myapp"}, "topic")
	assert.Equal(t, "prod-topic", fqName)
}

func Test_DefaultGroupIdNamingStrategy_Empty(t *testing.T) {
	fqName := kafka.DefaultGroupIdNamingStrategy(cfg.AppId{Project: "prog", Environment: "prod", Application: "myapp"}, "")
	assert.Equal(t, "myapp", fqName)
}

func Test_DefaultGroupIdNamingStrategy_Specified(t *testing.T) {
	fqName := kafka.DefaultGroupIdNamingStrategy(cfg.AppId{Project: "prog", Environment: "prod", Application: "myapp"}, "group")
	assert.Equal(t, "prod-myapp-group", fqName)
}
