package store

import (
	"github.com/dave/vecty/examples/bug/actions"
	"github.com/dave/vecty/examples/bug/dispatcher"
	"github.com/dave/vecty/examples/bug/store/model"
	"github.com/dave/vecty/storeutil"
)

var (
	Items     []*model.Item
	Listeners = storeutil.NewListenerRegistry()
)

func init() {
	dispatcher.Register(onAction)
}

func onAction(action interface{}) {
	switch action.(type) {
	case *actions.Action:

	default:
		return // don't fire listeners
	}

	//fmt.Println("Firing listners")
	Listeners.Fire()
}
