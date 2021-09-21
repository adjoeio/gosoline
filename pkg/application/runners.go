package application

import (
	"time"

	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/exec"
	"github.com/justtrackio/gosoline/pkg/mdlsub"
	"github.com/justtrackio/gosoline/pkg/stream"
)

func RunApiServer(definer apiserver.Definer, options ...Option) {
	options = append(options, WithExecBackoffSettings(&exec.BackoffSettings{
		InitialInterval: time.Millisecond * 100,
		MaxElapsedTime:  time.Second * 10,
		MaxInterval:     time.Second,
	}))

	app := Default(options...)
	app.Add("api", apiserver.New(definer))
	app.Run()
}

func RunConsumer(callback stream.ConsumerCallbackFactory, options ...Option) {
	RunConsumers(stream.ConsumerCallbackMap{
		"default": callback,
	})
}

func RunConsumers(consumers stream.ConsumerCallbackMap, options ...Option) {
	factory := stream.NewConsumerFactory(consumers)
	options = append(options, WithExecBackoffInfinite)

	app := Default(options...)
	app.AddFactory(factory)
	app.Run()
}

func RunMdlSubscriber(transformers mdlsub.TransformerMapTypeVersionFactories, options ...Option) {
	subs := mdlsub.NewSubscriberFactory(transformers)

	options = append(options, WithExecBackoffInfinite)

	app := Default(options...)
	app.AddFactory(subs)
	app.Run()
}
