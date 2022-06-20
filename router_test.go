package tgbcr

import (
	"testing"
)

func Test_Dispatch(t *testing.T) {

	router := New()
	router.Handle("/hello", func() {})
	router.Dispatch("/hello params")

}
