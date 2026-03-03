package assert

import (
	"reflect"
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got T, want T, msg ...string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		msg := strings.Join(msg, ": ")
		t.Logf("%s\ngot:\n%#v\nwant:\n%#v", msg, got, want)
		t.Fail()
	}
}
