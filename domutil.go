package vecty

import "github.com/gopherjs/gopherjs/js"

func removeNode(node *js.Object) {
	node.Get("parentNode").Call("removeChild", node)
}

func replaceNode(newNode, oldNode *js.Object) {
	if newNode == oldNode {
		return
	}
	if oldNode.Get("parentNode") == (*js.Object)(nil) {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
