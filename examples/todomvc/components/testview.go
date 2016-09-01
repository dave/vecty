package components

import (
	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/examples/todomvc/store"
)

type TestView struct {
	vecty.Composite
}

// Apply implements the vecty.Markup interface.
func (p *TestView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *TestView) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*TestView); ok {
		p.Body = oldComp.Body
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *TestView) render() vecty.Component {
	if store.Type() == "p" {
		return elem.Paragraph(vecty.Text(fmt.Sprintf("P: Count = %d", store.Count())))
	} else {
		return elem.Div(vecty.Text(fmt.Sprintf("DIV: Count = %d", store.Count())))
	}
}
