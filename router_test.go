package tgbcr

import (
	"testing"
)

func Test_Dispatch(t *testing.T) {

	var request interface{}

	router := New()
	router.Handle("/hello", func(r interface{}) {})
	router.Dispatch("/hello params", request)

}
