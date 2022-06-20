package tgbcr

import "testing"

func Test_New(t *testing.T) {

	router := New()
	router.Handle("/hello", func(c Context) {})
	router.Dispatch()

}
