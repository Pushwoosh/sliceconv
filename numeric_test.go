package sliceconv

import (
	"testing"
)

func Test_example_Integer(t *testing.T) {
	from := []int32{1, 2, 3}
	to := Integer[int32, int64](from)
	t.Log(to)
}

func Test_example_Float(t *testing.T) {
	from := []float32{1.0, 2.0, 3.0}
	to := Float[float32, float64](from)
	t.Log(to)
}

func Test_example_Platform(t *testing.T) {
	type protoPlatform uint32
	type sdkPlatform uint8
	const (
		UNKNOWN sdkPlatform = iota
		IOS
		Android
		IE
		Email
	)
	from := []sdkPlatform{IOS, Android, IE, Email}
	to := Integer[sdkPlatform, protoPlatform](from)
	t.Log(to)
}
