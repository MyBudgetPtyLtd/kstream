package streams

import (
	"context"
	"github.com/MyBudgetPtyLtd/kstream/v2/kafka"
)

type RunnerContext interface {
	context.Context
	ConsumerGroupMeta() (*kafka.GroupMeta, error)
	TopicMeta() kafka.TopicMeta
}

type kRunnerContext struct {
	context.Context
}

func (k *kRunnerContext) ConsumerGroupMeta() (*kafka.GroupMeta, error) {
	panic("implement me")
}

func (k *kRunnerContext) TopicMeta() kafka.TopicMeta {
	panic("implement me")
}
