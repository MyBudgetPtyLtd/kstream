package streams

import (
	"github.com/MyBudgetPtyLtd/kstream/v2/streams/stores"
)

type GlobalStoreBuilderWrapper struct {
	store stores.Store
}

func (s *GlobalStoreBuilderWrapper) Name() string {
	return s.store.Name()
}

func (s *GlobalStoreBuilderWrapper) Build(name string, options ...stores.Option) (stores.Store, error) {
	return s.store, nil
}
