package components

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/examples/todomvc/actions"
	"github.com/davelondon/vecty/examples/todomvc/dispatcher"
	"github.com/davelondon/vecty/examples/todomvc/store"
	"github.com/davelondon/vecty/examples/todomvc/store/model"
	"github.com/davelondon/vecty/prop"
)

type FilterButton struct {
	vecty.Core

	Label  string
	Filter model.FilterState
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) Render() *vecty.HTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}
