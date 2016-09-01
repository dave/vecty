package main

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/examples/todomvc/components"
	"github.com/davelondon/vecty/examples/todomvc/store"
)

func main() {
	vecty.SetTitle("GopherJS • TodoMVC")
	p := &components.PageView{}
	store.Listeners.Add(p, func() {
		p.ReconcileBody()
	})
	vecty.RenderAsBody(p)
}
