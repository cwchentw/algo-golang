package matrix

import (
	"math/rand"
	"testing"
)

func TestMatrixNew(t *testing.T) {
	t.Parallel()

	m := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	if !(m.GetAt(0, 0) == 1.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 2.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 3.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 4.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 5.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 6.0) {
		t.Error("Wrong value")
	}
}

func TestMatrixIdentity(t *testing.T) {
	t.Parallel()

	iden := Identity(3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				if iden.GetAt(i, j) != 1 {
					t.Error("Wrong value")
				}
			} else {
				if iden.GetAt(i, j) != 0 {
					t.Error("Wrong value")
				}
			}
		}
	}
}

func TestMatrixT(t *testing.T) {
	t.Parallel()

	m := New(
		[][]float64{
			[]float64{1, 2, 3},
			[]float64{4, 5, 6},
		},
	)

	out := T(m)

	if !(out.GetAt(0, 0) == 1) {
		t.Error("Wrong value")
	}

	if !(out.GetAt(0, 1) == 4) {
		t.Error("Wrong value")
	}

	if !(out.GetAt(1, 0) == 2) {
		t.Error("Wrong value")
	}

	if !(out.GetAt(1, 1) == 5) {
		t.Error("Wrong value")
	}

	if !(out.GetAt(2, 0) == 3) {
		t.Error("Wrong value")
	}

	if !(out.GetAt(2, 1) == 6) {
		t.Error("Wrong value")
	}
}

func TestMatrixAdd(t *testing.T) {
	t.Parallel()

	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{2.0, 3.0, 4.0},
			[]float64{5.0, 6.0, 7.0},
		},
	)

	m := Add(m1, m2)

	if !(m.GetAt(0, 0) == 3.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 5.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 7.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 9.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 11.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 13.0) {
		t.Error("Wrong value")
	}
}

func TestMatrixSub(t *testing.T) {
	t.Parallel()

	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{2.0, 3.0, 4.0},
			[]float64{5.0, 6.0, 7.0},
		},
	)

	m := Sub(m1, m2)

	if !(m.GetAt(0, 0) == -1) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == -1) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == -1) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == -1) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == -1) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == -1) {
		t.Error("Wrong value")
	}
}

func TestMatrixMul(t *testing.T) {
	t.Parallel()

	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{2.0, 3.0, 4.0},
			[]float64{5.0, 6.0, 7.0},
		},
	)

	m := Mul(m1, m2)

	if !(m.GetAt(0, 0) == 2) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 6) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 12) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 20) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 30) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 42) {
		t.Error("Wrong value")
	}
}

func TestMatrixDiv(t *testing.T) {
	t.Parallel()

	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{2.0, 3.0, 4.0},
			[]float64{5.0, 6.0, 7.0},
		},
	)

	m := Div(m1, m2)

	if !(m.GetAt(0, 0) == 0.5) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 2.0/3.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 2) == 0.75) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 0.8) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 5.0/6.0) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 2) == 6.0/7.0) {
		t.Error("Wrong value")
	}
}

func TestMatrixDot(t *testing.T) {
	t.Parallel()

	m1 := New(
		[][]float64{
			[]float64{1.0, 2.0, 3.0},
			[]float64{4.0, 5.0, 6.0},
		},
	)

	m2 := New(
		[][]float64{
			[]float64{1.0, 4.0},
			[]float64{2.0, 5.0},
			[]float64{3.0, 6.0},
		},
	)

	m := Dot(m1, m2)

	if !(m.GetAt(0, 0) == 14) {
		t.Log(m.GetAt(0, 0))
		t.Error("Wrong value")
	}

	if !(m.GetAt(0, 1) == 32) {
		t.Log(m.GetAt(0, 1))
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 0) == 32) {
		t.Error("Wrong value")
	}

	if !(m.GetAt(1, 1) == 77) {
		t.Error("Wrong value")
	}
}

func BenchmarkMatrixAdd(b *testing.B) {
	const SIZE int = 3000
	r := rand.New(rand.NewSource(99))

	m := WithSize(SIZE, SIZE)
	n := WithSize(SIZE, SIZE)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			m.SetAt(i, j, r.Float64())
			n.SetAt(i, j, r.Float64())
		}
	}

	for i := 0; i < b.N; i++ {
		Add(m, n)
	}
}
