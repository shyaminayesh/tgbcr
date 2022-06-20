package tgbcr

import "fmt"

type Router struct{}
type Context struct{}
type Handler func(Context)

func New() *Router {
	return &Router{}
}

func (r *Router) Handle(path string, handle Handler) {

}

func (r *Router) Dispatch() {
	fmt.Println("Dispatch")
}
