package main

import (
	"github.com/dave/vecty"
	"github.com/dave/vecty/examples/bug/actions"
	"github.com/dave/vecty/examples/bug/components"
	"github.com/dave/vecty/examples/bug/dispatcher"
	"github.com/dave/vecty/examples/bug/store"

	"github.com/dave/vecty/examples/bug/store/model"
)

func main() {
	foo := &model.Item{Title: "foo"}
	bar := &model.Item{Title: "bar"}
	baz := &model.Item{Title: "baz"}
	//qux := &model.Item{Title: "qux"}
	store.Items = []*model.Item{
		foo,
		bar,
		baz,
	}
	p := &components.PageView{}
	store.Listeners.Add(p, func() {
		p.Items = store.Items
		vecty.Rerender(p)
	})
	p.Items = store.Items
	vecty.RenderBody(p)

	store.Items = []*model.Item{
		//qux,
		foo,
		bar,
		baz,
	}

	dispatcher.Dispatch(&actions.Action{})
}
