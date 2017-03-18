package components

import (
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/examples/bug/store/model"
)

type PageView struct {
	vecty.Core

	Items []*model.Item
}

func (p *PageView) Render() *vecty.HTML {
	items := vecty.List{}
	for i, item := range p.Items {
		items = append(items, &ItemView{Index: i, Item: item})
	}
	return elem.Body(
		elem.Div(
			vecty.Text("page"),
		),
		elem.Div(
			items,
		),
	)
}
