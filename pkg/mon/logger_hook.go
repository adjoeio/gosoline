package mon

import (
	"context"
)

type LoggerHook interface {
	Fire(level string, msg string, logErr error, fields Fields, tags Tags, configValues ConfigValues, context context.Context, ecsMetadata EcsMetadata)
}
