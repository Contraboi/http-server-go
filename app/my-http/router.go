package myhttp

import (
	"strings"
)

type HandlerFunc func(res *Response, req *Request, ctx *Context)

type Node struct {
	Children  map[string]*Node
	handler   map[string]HandlerFunc
	isDyanmic bool
	param     string
	method    string
}

func NewNode(isDynamic bool) *Node {
	return &Node{
		Children:  make(map[string]*Node),
		isDyanmic: false,
		handler:   make(map[string]HandlerFunc),
	}
}

type Router struct {
	root *Node
}

func NewRouter() *Router {
	return &Router{
		root: NewNode(false),
	}
}

func (r *Router) Insert(path string, method string, handler HandlerFunc) {
	curr := r.root
	split := strings.Split(path, "/")

	for _, part := range split {
		sanitizedPart := strings.TrimSpace(part)
		isDynamic := strings.HasPrefix(part, ":")
		if sanitizedPart == "" {
			sanitizedPart = "/"
		}
		if isDynamic {
			sanitizedPart = ":"
		}

		if _, ok := curr.Children[sanitizedPart]; !ok {
			curr.Children[sanitizedPart] = NewNode(isDynamic)
			if isDynamic {
				curr.Children[sanitizedPart].param = strings.TrimPrefix(part, ":")
			}
		}

		curr.Children[sanitizedPart].handler[method] = handler
		curr.Children[sanitizedPart].method = method
		curr = curr.Children[sanitizedPart]
	}
}

func (r *Router) Search(path string) (*Node, map[string]string) {
	curr := r.root
	split := strings.Split(path, "/")
	params := make(map[string]string)

	for _, part := range split {
		if part == "" {
			part = "/"
		}

		if child, exists := curr.Children[part]; exists {
			curr = child
		} else if child, exists := curr.Children[":"]; exists {
			curr = child
			params[child.param] = part
		} else {
			return nil, nil
		}
	}

	return curr, params
}
