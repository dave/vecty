package components

import (
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/examples/todomvc/actions"
	"github.com/dave/vecty/examples/todomvc/dispatcher"
	"github.com/dave/vecty/examples/todomvc/store"
	"github.com/dave/vecty/examples/todomvc/store/model"
	"github.com/dave/vecty/prop"
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
