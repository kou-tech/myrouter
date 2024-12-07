package myrouter

import "net/http"

type Node struct {
	isRoot bool
	character byte
	children []*Node
	handlers map[string]http.Handler
}

func NewNode(character byte) *Node {
	return &Node{
		character: character,
		children: []*Node{},
		handlers: make(map[string]http.Handler),
	}
}

type Router struct {
	tree *Node
}

func NewRouter() *Router {
	return &Router{
		tree: &Node{
			isRoot: true,
			handlers: nil,
		},
	}
}

func (r *Router) Get(endpoint string, handler http.Handler) {
	r.insert(http.MethodGet, endpoint, handler)
}

func (r *Router) insert(method, endpoint string, handler http.Handler) {
	currentNode := r.tree
	for i := 0; i < len(endpoint); i++ {
		target := endpoint[i]

		nextNode := currentNode.nextChild(target)
		if nextNode == nil {
			node := NewNode(target)
			currentNode.children = append(currentNode.children, node)
			currentNode = node
			continue
		}

		currentNode = nextNode
	}
	currentNode.handlers[method] = handler
}

func (n *Node) nextChild(character byte) *Node {
	for _, child := range n.children {
		if child.character == character {
			return child
		}
	}

	return nil
}

func (r *Router) Search(method, endpoint string) http.Handler {
	currentNode := r.tree
	lcpIndex := 0
	for {
		nextNode := currentNode.nextChild(endpoint[lcpIndex])

		if nextNode == nil {
			return nil
		}

		// 各ノードの文字数は1文字なので、次のノードが見つかったらインデックスを進める
		lcpIndex++
		currentNode = nextNode
		if lcpIndex == len(endpoint) {
			break
		}
	}

	return currentNode.handlers[method]
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.Search(req.Method, req.URL.Path)
	if handler != nil {
		handler.ServeHTTP(w, req)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))

}
