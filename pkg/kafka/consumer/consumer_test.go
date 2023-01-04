package consumer_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/justtrackio/gosoline/pkg/kafka/consumer"
	"github.com/justtrackio/gosoline/pkg/kafka/consumer/mocks"
	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConsumer_Manager_Batch_Commit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var (
		managerBatches = 0
		manager        = &mocks.OffsetManager{}
	)

	manager.On("Start", mock.Anything).Return(
		func(ctx context.Context) error {
			<-ctx.Done()
			return ctx.Err()
		})
	manager.On("Batch", mock.Anything).Return(
		func(ctx context.Context) []kafka.Message {
			defer func() {
				time.Sleep(time.Millisecond)
				managerBatches += 1
			}()

			return []kafka.Message{OnFetch(ctx, managerBatches)}
		},
		func(ctx context.Context) error {
			return nil
		},
	)
	manager.On("Commit", mock.Anything, kafka.Message{Offset: 1, Partition: 1}).Return(nil)
	manager.On("Commit", mock.Anything, kafka.Message{Offset: 2, Partition: 2}).Return(errors.New("reader: failed"))
	defer manager.AssertExpectations(t)

	con, err := consumer.NewConsumerWithInterfaces(
		&consumer.Settings{
			BatchSize:    100,
			BatchTimeout: time.Second,
		},
		logMocks.NewLoggerMockedAll(),
		manager,
	)
	assert.Nil(t, err)

	_ = con.Run(ctx)

	assert.Equal(t, kafka.Message{Offset: 1, Partition: 1}, <-con.Data())
	assert.Equal(t, kafka.Message{Offset: 2, Partition: 2}, <-con.Data())
	assert.Equal(t, kafka.Message{Offset: 3, Partition: 3}, <-con.Data())

	// Should succeed.
	assert.Nil(t, con.Commit(ctx, kafka.Message{Offset: 1, Partition: 1}))
	// Should not succeed.
	assert.NotNil(t, con.Commit(ctx, kafka.Message{Offset: 2, Partition: 2}))
}

func TestConsumer_Manager_Batch_Graceful_Exit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var (
		manager    = &mocks.OffsetManager{}
		managerErr = errors.New("manager: failed")
	)

	manager.On("Start", mock.Anything).Return(func(ctx context.Context) error {
		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}

		return managerErr
	})
	manager.On("Batch", mock.Anything).Return(
		func(ctx context.Context) []kafka.Message {
			select {
			case <-time.After(time.Hour):
			case <-ctx.Done():
				return []kafka.Message{}
			}

			return []kafka.Message{}
		},
		func(ctx context.Context) error {
			return nil
		},
	)
	defer manager.AssertExpectations(t)

	con, err := consumer.NewConsumerWithInterfaces(
		&consumer.Settings{
			BatchSize:    1000,
			BatchTimeout: time.Second,
		},
		logMocks.NewLoggerMockedAll(),
		manager,
	)
	assert.Nil(t, err)

	err = con.Run(ctx)
	assert.ErrorIs(t, err, managerErr)
}
