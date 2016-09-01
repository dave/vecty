package components

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/examples/todomvc/actions"
	"github.com/davelondon/vecty/examples/todomvc/dispatcher"
)

type PageView struct {
	vecty.Composite
}

// Apply implements the vecty.Markup interface.
func (p *PageView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*PageView); ok {
		p.Body = oldComp.Body
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) render() vecty.Component {
	return elem.Div(
		&TestView{},
		elem.Button(
			event.Click(func(ev *vecty.Event) {
				dispatcher.Dispatch(&actions.Increment{})
			}),
			vecty.Text("Increment"),
		),
		elem.Button(
			event.Click(func(ev *vecty.Event) {
				dispatcher.Dispatch(&actions.Change{"div"})
			}),
			vecty.Text("Div"),
		),
		elem.Button(
			event.Click(func(ev *vecty.Event) {
				dispatcher.Dispatch(&actions.Change{"p"})
			}),
			vecty.Text("Paragraph"),
		),
	)
}
