package lp

import "gonum.org/v1/gonum/mat"

// Check interfaces are correctly implemented
var (
	basis *BasisView

	_ Basis       = basis
	_ mat.Matrix  = basis
	_ mat.Mutable = basis
)

// Basis interface for the simplex method
type Basis interface {
	SetVar(i, v int)
	IsBasic(v int) bool

	mat.Mutable
}

// BasisView implementation as view of Dense matrix
type BasisView struct {
	A    *mat.Dense
	vars []int
}

// NewBasis initializes a new LP basis
func NewBasis(A *mat.Dense, vars []int) (b BasisView) {
	if vars == nil {
		r, _ := A.Dims()
		vars = make([]int, r)
	}
	return BasisView{A, vars}
}

// Dims returns the dimensions of a Matrix.
func (b BasisView) Dims() (r, c int) {
	r, _ = b.A.Dims()
	c = len(b.vars)
	return
}

// At returns the value of a matrix element at row i, column j.
// It will panic if i or j are out of bounds for the matrix.
func (b BasisView) At(i, j int) float64 {
	return b.A.At(i, b.vars[j])
}

// T returns the transpose of the Matrix. Whether T returns a copy of the
// underlying data is implementation dependent.
// This method may be implemented using the Transpose type, which
// provides an implicit matrix transpose.
func (b BasisView) T() mat.Matrix {
	return mat.DenseCopyOf(b).T()
}

// Set alters the matrix element at row i, column j to v.
// It will panic if i or j are out of bounds for the matrix.
func (b BasisView) Set(i, j int, v float64) {
	b.A.Set(i, b.vars[j], v)
}

// SetVar swaps the basic variable at element i with v.
func (b BasisView) SetVar(i, v int) {
	b.vars[i] = v
}

// IsBasic returns true of the passed index denotes a basic variable
func (b BasisView) IsBasic(v int) bool {
	for i := 0; i < len(b.vars); i++ {
		if b.vars[i] == v {
			return true
		}
	}
	return false
}
