package mock

import (
	"github.com/MyBudgetPtyLtd/kstream/v2/backend"
	"github.com/MyBudgetPtyLtd/kstream/v2/backend/pebble"
	"os"
	"time"
)

func NewMockBackend(name string, _ time.Duration) backend.Backend {
	conf := pebble.NewConfig()
	tmp, err := os.MkdirTemp(``, `*`)
	if err != nil {
		panic(err)
	}
	conf.Dir = tmp
	b, err := pebble.NewPebbleBackend(name, conf)
	if err != nil {
		panic(err)
	}

	return b
}
