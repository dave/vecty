package components

import (
	"fmt"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/examples/bug/store/model"
	"github.com/dave/vecty/prop"
)

type ItemView struct {
	vecty.Core

	Index int
	Item  *model.Item
}

func (p *ItemView) Render() *vecty.HTML {
	return elem.ListItem(
		elem.Div(
			prop.Class("view"),
			vecty.Text(fmt.Sprint(p.Index, " ", p.Item.Title)),
		),
	)
}
