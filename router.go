package tgbcr

import (
	"strings"
)

type Router struct {
	Collection struct {
		Childs []Child
	}
}

type Context struct {
	Request interface{}
}

type Handler func(c interface{})
type Child struct {
	Path    string
	Handler Handler
}

func New() *Router {
	return &Router{}
}

func (r *Router) Handle(path string, handler Handler) {

	if path == "" {
		panic("Path must not be empty")
	}

	if path[0] != '/' {
		panic("Path must start with \"/\"")
	}

	if len(path) < 2 {
		panic("Path must have at least single charactor to match")
	}

	r.Collection.Childs = append(r.Collection.Childs, Child{
		Path:    path,
		Handler: handler,
	})

}

func (r *Router) Dispatch(text string, request interface{}) {

	context := &Context{
		Request: request,
	}

	/**
	* Sanitize and extract the command string from
	* the provided text string
	 */
	path := strings.Fields(text)[0]
	for _, e := range r.Collection.Childs {
		if e.Path == path {
			e.Handler(context)
			return
		}
	}

}
