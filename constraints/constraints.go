package constraints

type Signed interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a generic type constraint for all unsigned integers.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

// Float is a generic type constraint for all floating point types.
type Float interface{
	~float32 | ~float64
}

// Number is a generic type constraint for all numeric types in Go except Complex types.
type Number interface {
	Integer | Float
}

type Ordered interface {
	Integer	| ~string | Float
}
