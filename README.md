# Sliceconv
Sliceconv is a go library for converting slices from one type to another.

# Integer

`Integer` converts any integer-based slice to any other integer-based slice. No overflow checks.
```go
func Integer[T1 constraints.Integer, T2 constraints.Integer](from []T1) []T2{}
```

# Float

`Float` converts any float-based slice to any other float-based slice. No overflow checks.
```go
func Float[T1 constraints.Float, T2 constraints.Float](from []T1) []T2 {}
```

# Examples

## Integer
Standard types:
```go
	from := []int32{1, 2, 3}
	to := Integer[int32, int64](from)
```

User defined types:
```go
	type protoPlatform uint32
	type sdkPlatform uint8
	const (
		IOS sdkPlatform = iota + 1
		Android
		Windows
		Web
	)
	from := []sdkPlatform{IOS, Android, Windows}
	to := Integer[sdkPlatform, protoPlatform](from)
```

## Float
```go
    from := []float32{1.1, 2.2, 3.3}
    to := Float[float32, float64](from)
```
