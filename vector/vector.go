package vector

import (
	"math/big"
	"reflect"
)

type Vector struct {
	vec []interface{}
}

func New(args ...interface{}) *Vector {
	v := new(Vector)
	v.vec = make([]interface{}, len(args))

	for i, e := range args {
		v.vec[i] = e
	}

	return v
}

func WithSize(s int) *Vector {
	if s <= 0 {
		negativeSize()
	}

	v := new(Vector)
	v.vec = make([]interface{}, s)

	return v
}

func FromArray(arr []interface{}) *Vector {
	v := new(Vector)
	v.vec = make([]interface{}, len(arr))

	for i, e := range arr {
		v.vec[i] = e
	}

	return v
}

func (v *Vector) GetAt(i int) interface{} {
	if i < 0 || i >= len(v.vec) {
		indexOutOfRange()
	}

	return v.vec[i]
}

func (v *Vector) SetAt(i int, data interface{}) {
	if i < 0 || i >= len(v.vec) {
		indexOutOfRange()
	}

	v.vec[i] = data
}

// Vection addition
// This method relies on reflect to automatically detect element type, which
// tends to be slow. Using CalcBy method is favored.
func (v *Vector) Add(other *Vector) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		ta := reflect.TypeOf(v.vec[i])
		tb := reflect.TypeOf(other.vec[i])

		if ta != tb {
			return out, unequalType()
		}

		switch ta.String() {
		case "float64":
			a := v.vec[i].(float64)
			b := other.vec[i].(float64)
			out.SetAt(i, a+b)
		case "int":
			a := v.vec[i].(int)
			b := other.vec[i].(int)
			out.SetAt(i, a+b)
		case reflect.TypeOf(big.NewInt(0)).String():
			a := v.vec[i].(*big.Int)
			b := other.vec[i].(*big.Int)
			n := big.NewInt(0)
			n.Add(a, b)
			out.SetAt(i, n)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			a := v.vec[i].(*big.Float)
			b := other.vec[i].(*big.Float)
			n := big.NewFloat(0.0)
			n.Add(a, b)
			out.SetAt(i, n)
		default:
			return out, unknownType()
		}
	}

	return out, nil
}

// Vector algebra delegating to function object.
// This method delegates vector algebra to function object set by users, making
// it faster then these methods relying on reflection.
func (v *Vector) CalcBy(other *Vector, f func(interface{}, interface{}) (interface{}, error)) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		n, err := f(v.vec[i], other.vec[i])
		if err != nil {
			return out, err
		}
		out.SetAt(i, n)
	}

	return out, nil
}