package lp

import (
	"fmt"
	"strings"

	"github.com/willf/pad"
	"gonum.org/v1/gonum/mat"
)

// Tableau represents a simplex tableau (mainly for debugging purposes)
type Tableau struct {
	Vars  []string
	A     mat.Matrix
	Coeff mat.Matrix
	Right mat.Matrix
	B     Basis
}

// NewTableau creates a new simplex tableau
func NewTableau(vars []string, a mat.Matrix, c mat.Matrix, rhs mat.Matrix, b Basis) (tbl Tableau) {
	return Tableau{vars, a, c, rhs, b}
}

// Print prints a simplex tableau for debugging purposes
func (tbl Tableau) Print() {
	rows, cols := tbl.A.Dims()

	var colWidth int = 5
	for _, varName := range tbl.Vars {
		if len(varName) > colWidth {
			colWidth = len(varName)
		}
	}
	colWidth += 2

	// Headers
	for i := 0; i < cols; i++ {
		if tbl.B.IsBasic(i) {
			fmt.Print(pad.Left(fmt.Sprintf("*%v", tbl.Vars[i]), colWidth, " "))
		} else {
			fmt.Print(pad.Left(tbl.Vars[i], colWidth, " "))
		}
	}
	fmt.Print(" │ ")
	fmt.Print(pad.Left("RHS", colWidth, " "))
	fmt.Println()

	// Separator
	fmt.Print(strings.Repeat("─", colWidth*len(tbl.Vars)))
	fmt.Print("─┼─")
	fmt.Print(strings.Repeat("─", colWidth))
	fmt.Println()

	// Coefficients
	for c := 0; c < cols; c++ {
		fmt.Print(pad.Left(fmt.Sprintf("%.3f", tbl.Coeff.At(c, 0)), colWidth, " "))
	}
	fmt.Print(" │ ")
	fmt.Print(pad.Left("z", colWidth, " "))
	fmt.Println()

	// Separator
	fmt.Print(strings.Repeat("─", colWidth*len(tbl.Vars)))
	fmt.Print("─┼─")
	fmt.Print(strings.Repeat("─", colWidth))
	fmt.Println()

	// Body
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Print(pad.Left(fmt.Sprintf("%.3f", tbl.A.At(r, c)), colWidth, " "))
		}
		fmt.Print(" │ ")
		fmt.Print(pad.Left(fmt.Sprintf("%.3f", tbl.Right.At(r, 0)), colWidth, " "))
		fmt.Println()
	}
}
