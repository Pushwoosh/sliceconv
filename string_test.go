package sliceconv

import (
	"reflect"
	"testing"
)

func Test_example_String(t *testing.T) {
	type protoStatus string
	type sdkStatus string
	const (
		Active  sdkStatus = "active"
		Blocked sdkStatus = "blocked"
	)
	from := []sdkStatus{Active, Blocked}
	to := String[sdkStatus, protoStatus](from)
	t.Log(to)

	if want := []protoStatus{"active", "blocked"}; !reflect.DeepEqual(to, want) {
		t.Errorf("got %v, want %v", to, want)
	}
	if got := String[string, string](nil); got == nil || len(got) != 0 {
		t.Errorf("nil input: got %#v, want empty non-nil slice", got)
	}
}
