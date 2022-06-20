package tgbcr

import (
	"testing"
)

func Test_Dispatch(t *testing.T) {

	var Request Context

	router := New()
	router.Handle("/hello", func(c Context) {})
	router.Dispatch("/hello params", Request)

}
