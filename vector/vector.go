package vector

import (
	"math/big"
	"reflect"
)

// Vector class.
// Internally, this class holds an array to empty interface, i.e. arbitrary type.
type Vector struct {
	vec []interface{}
}

// Create a Vector with some arguments.
func New(args ...interface{}) *Vector {
	v := new(Vector)
	v.vec = make([]interface{}, len(args))

	for i, e := range args {
		v.vec[i] = e
	}

	return v
}

// Create an empty Vector with specific size.
func WithSize(s int) *Vector {
	if s <= 0 {
		negativeSize()
	}

	v := new(Vector)
	v.vec = make([]interface{}, s)

	return v
}

// Create a Vector from an array.
func FromArray(arr []interface{}) *Vector {
	v := new(Vector)
	v.vec = make([]interface{}, len(arr))

	for i, e := range arr {
		v.vec[i] = e
	}

	return v
}

func (v *Vector) Len() int {
	return len(v.vec)
}

// Get specific element by indexing.
func (v *Vector) GetAt(i int) interface{} {
	if i < 0 || i >= len(v.vec) {
		indexOutOfRange()
	}

	return v.vec[i]
}

// Set specific element by specific index.
func (v *Vector) SetAt(i int, data interface{}) {
	if i < 0 || i >= len(v.vec) {
		indexOutOfRange()
	}

	v.vec[i] = data
}

// Vector addition.
// This method relies on reflect to automatically detect element type, which
// tends to be slow. Using CalcBy method is favored.
func (v *Vector) Add(other *Vector) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		ta := reflect.TypeOf(v.vec[i]).String()
		tb := reflect.TypeOf(other.vec[i]).String()

		if !(ta == "float64" || ta == "int") &&
			!(tb == "float64" || tb == "int") {
			if ta != tb {
				return out, unequalType()
			}
		}

		switch ta {
		case "float64":
			switch tb {
			case "float64":
				a := v.vec[i].(float64)
				b := other.vec[i].(float64)
				out.SetAt(i, a+b)
			case "int":
				a := v.vec[i].(float64)
				b := other.vec[i].(int)
				out.SetAt(i, a+float64(b))
			default:
				return out, unknownType()
			}
		case "int":
			switch tb {
			case "float64":
				a := v.vec[i].(int)
				b := other.vec[i].(float64)
				out.SetAt(i, float64(a)+b)
			case "int":
				a := v.vec[i].(int)
				b := other.vec[i].(int)
				out.SetAt(i, a+b)
			default:
				return out, unknownType()
			}
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

// Vector subtraction.
// This method relies on reflect to automatically detect element type, which
// tends to be slow. Using CalcBy method is favored.
func (v *Vector) Sub(other *Vector) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		ta := reflect.TypeOf(v.vec[i]).String()
		tb := reflect.TypeOf(other.vec[i]).String()

		if !(ta == "float64" || ta == "int") &&
			!(tb == "float64" || tb == "int") {
			if ta != tb {
				return out, unequalType()
			}
		}

		switch ta {
		case "float64":
			switch tb {
			case "float64":
				a := v.vec[i].(float64)
				b := other.vec[i].(float64)
				out.SetAt(i, a-b)
			case "int":
				a := v.vec[i].(float64)
				b := other.vec[i].(int)
				out.SetAt(i, a-float64(b))
			default:
				return out, unknownType()
			}
		case "int":
			switch tb {
			case "float64":
				a := v.vec[i].(int)
				b := other.vec[i].(float64)
				out.SetAt(i, float64(a)-b)
			case "int":
				a := v.vec[i].(int)
				b := other.vec[i].(int)
				out.SetAt(i, a-b)
			default:
				return out, unknownType()
			}
		case reflect.TypeOf(big.NewInt(0)).String():
			a := v.vec[i].(*big.Int)
			b := other.vec[i].(*big.Int)
			n := big.NewInt(0)
			n.Sub(a, b)
			out.SetAt(i, n)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			a := v.vec[i].(*big.Float)
			b := other.vec[i].(*big.Float)
			n := big.NewFloat(0.0)
			n.Sub(a, b)
			out.SetAt(i, n)
		default:
			return out, unknownType()
		}
	}

	return out, nil
}

// Vector multiplication.
// This method relies on reflect to automatically detect element type, which
// tends to be slow. Using CalcBy method is favored.
func (v *Vector) Mul(other *Vector) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		ta := reflect.TypeOf(v.vec[i]).String()
		tb := reflect.TypeOf(other.vec[i]).String()

		if !(ta == "float64" || ta == "int") &&
			!(tb == "float64" || tb == "int") {
			if ta != tb {
				return out, unequalType()
			}
		}

		switch ta {
		case "float64":
			switch tb {
			case "float64":
				a := v.vec[i].(float64)
				b := other.vec[i].(float64)
				out.SetAt(i, a*b)
			case "int":
				a := v.vec[i].(float64)
				b := other.vec[i].(int)
				out.SetAt(i, a*float64(b))
			default:
				return out, unknownType()
			}
		case "int":
			switch tb {
			case "float64":
				a := v.vec[i].(int)
				b := other.vec[i].(float64)
				out.SetAt(i, float64(a)*b)
			case "int":
				a := v.vec[i].(int)
				b := other.vec[i].(int)
				out.SetAt(i, a*b)
			default:
				return out, unknownType()
			}
		case reflect.TypeOf(big.NewInt(0)).String():
			a := v.vec[i].(*big.Int)
			b := other.vec[i].(*big.Int)
			n := big.NewInt(0)
			n.Mul(a, b)
			out.SetAt(i, n)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			a := v.vec[i].(*big.Float)
			b := other.vec[i].(*big.Float)
			n := big.NewFloat(0.0)
			n.Mul(a, b)
			out.SetAt(i, n)
		default:
			return out, unknownType()
		}
	}

	return out, nil
}

// Vector division.
// This method relies on reflect to automatically detect element type, which
// tends to be slow. Using CalcBy method is favored.
func (v *Vector) Div(other *Vector) (*Vector, error) {
	_len := len(v.vec)
	if _len != len(other.vec) {
		unequalLength()
	}

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		ta := reflect.TypeOf(v.vec[i]).String()
		tb := reflect.TypeOf(other.vec[i]).String()

		if !(ta == "float64" || ta == "int") &&
			!(tb == "float64" || tb == "int") {
			if ta != tb {
				return out, unequalType()
			}
		}

		switch ta {
		case "float64":
			switch tb {
			case "float64":
				a := v.vec[i].(float64)
				b := other.vec[i].(float64)
				out.SetAt(i, a/b)
			case "int":
				a := v.vec[i].(float64)
				b := other.vec[i].(int)
				out.SetAt(i, a/float64(b))
			default:
				return out, unknownType()
			}
		case "int":
			switch tb {
			case "float64":
				a := v.vec[i].(int)
				b := other.vec[i].(float64)
				out.SetAt(i, float64(a)/b)
			case "int":
				a := v.vec[i].(int)
				b := other.vec[i].(int)
				out.SetAt(i, a/b)
			default:
				return out, unknownType()
			}
		case reflect.TypeOf(big.NewInt(0)).String():
			a := v.vec[i].(*big.Int)
			b := other.vec[i].(*big.Int)
			n := big.NewInt(0)
			n.Div(a, b)
			out.SetAt(i, n)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			a := v.vec[i].(*big.Float)
			b := other.vec[i].(*big.Float)
			n := big.NewFloat(0.0)
			n.Quo(a, b)
			out.SetAt(i, n)
		default:
			return out, unknownType()
		}
	}

	return out, nil
}

// Vector transformation delegating to function object.
// This method delegates vector transformation to function object set by users.
func (v *Vector) Map(f func(interface{}) (interface{}, error)) (*Vector, error) {
	_len := len(v.vec)

	out := WithSize(_len)

	for i := 0; i < _len; i++ {
		n, err := f(v.vec[i])
		if err != nil {
			return out, err
		}
		out.SetAt(i, n)
	}

	return out, nil
}

// Vector algebra delegating to function object.
// This method delegates vector algebra to function object set by users, making
// it faster then these methods relying on reflection.
func (v *Vector) Apply(other *Vector, f func(interface{}, interface{}) (interface{}, error)) (*Vector, error) {
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

// Vector to scalars reduction.
// This method delegates vector reduction to function object set by users.
func (v *Vector) Reduce(f func(interface{}, interface{}) (interface{}, error)) (interface{}, error) {
	_len := v.Len()
	if _len == 0 {
		return nil, nil
	} else if _len == 1 {
		return v.vec[0], nil
	}

	n := v.vec[0]
	for i := 1; i < _len; i++ {
		m, err := f(n, v.vec[i])
		n = m
		if err != nil {
			return nil, err
		}
	}

	return n, nil
}
