package main

import (
	"fmt"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

type MyCustomHandlerSettings struct {
	Channel log.Channel `cfg:"channel"`
}

type MyCustomHandler struct {
	channel log.Channel
}

func (h *MyCustomHandler) Channels() []log.Channel {
	return []log.Channel{h.channel}
}

func (h *MyCustomHandler) Level() int {
	return log.PriorityInfo
}

func (h *MyCustomHandler) Log(timestamp time.Time, _ int, msg string, _ []interface{}, _ error, _ log.Data) error {
	fmt.Printf("%s happenend at %s", msg, timestamp.Format(time.RFC822))
	return nil
}

func MyCustomHandlerFactory(config cfg.Config, name string) (log.Handler, error) {
	settings := &MyCustomHandlerSettings{}
	log.UnmarshalHandlerSettingsFromConfig(config, name, settings)

	return &MyCustomHandler{
		channel: settings.Channel,
	}, nil
}

func main() {
	log.AddHandlerFactory("my-custom-handler", MyCustomHandlerFactory)
}
