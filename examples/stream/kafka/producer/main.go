package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/justtrackio/gosoline/pkg/application"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/kernel"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/justtrackio/gosoline/pkg/mdl"
	"github.com/justtrackio/gosoline/pkg/stream"
	"github.com/justtrackio/gosoline/pkg/uuid"
)

func main() {
	application.RunModule("producer", newOutputModule)
}

type ExampleRecord struct {
	Id        string `json:"id"`
	SoldItems int    `json:"soldItems"`
}

type outputModule struct {
	logger   log.Logger
	uuidGen  uuid.Uuid
	producer stream.Producer
}

func newOutputModule(ctx context.Context, config cfg.Config, logger log.Logger) (kernel.Module, error) {
	var err error

	producer, err := stream.NewProducer(ctx, config, logger, "example_records")
	if err != nil {
		return nil, fmt.Errorf("failed to init Kafka producer for datalake: %w", err)
	}

	modelId := mdl.ModelId{
		Name: "example_records",
	}
	modelId.PadFromConfig(config)

	module := &outputModule{
		logger:   logger,
		uuidGen:  uuid.New(),
		producer: producer,
	}

	return module, nil
}

func (p outputModule) Run(ctx context.Context) error {
	ticker := clock.NewRealTicker(time.Millisecond * 100)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.Chan():
			id := p.uuidGen.NewV4()
			body, err := json.Marshal(ExampleRecord{
				Id:        id,
				SoldItems: rand.Intn(100),
			})
			if err != nil {
				return fmt.Errorf("failed to marshal record: %w", err)
			}

			msg := stream.NewJsonMessage(string(body))
			msgAttributes := map[string]interface{}{stream.AttributeKafkaKey: id}
			p.logger.Info("publishing msg with key: %s", id)
			if err := p.producer.WriteOne(ctx, msg, msgAttributes); err != nil {
				p.logger.Error("failed to publish record with id %s: %s", id, err)
				continue
			}
			p.logger.Info("published msg with key: %s", id)
		}
	}
}
