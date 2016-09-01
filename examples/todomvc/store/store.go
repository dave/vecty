package store

import (
	"github.com/davelondon/vecty/examples/todomvc/actions"
	"github.com/davelondon/vecty/examples/todomvc/dispatcher"
	"github.com/davelondon/vecty/storeutil"
)

var (
	t         string
	c         int
	Listeners = storeutil.NewListenerRegistry()
)

func init() {
	dispatcher.Register(onAction)
}

func Count() int {
	return c
}

func Type() string {
	return t
}

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.Increment:
		c++

	case *actions.Change:
		t = a.Type

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}
