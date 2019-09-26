package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fastdeath/linear-programming/lp"
	"gonum.org/v1/gonum/mat"
)

/*
min Z = -2x - 3y - 4z
s.t.	 3x + 2y + 	z <= 10
		 2x + 5y + 3z <= 15
		 x,y,z >= 0

min  z + 2x + 3y + 4z = 0
s.t.	 3x + 2y +	z +	s = 10
		 2x + 5y + 3z +	t = 15
		 x,y,z,s,t >= 0

c =   (	-2	-3	-4	0	0 )T

A =   |	3	2	1	1	0 |
	  |	2	5	3	0	1 |

b =   (	10	15 )T

z	x	y	z	s	t	b
---------------------------
1	2	3	4	0	0	0
0	3	2	1	1	0	10
0	2	5	3	0	1	15
*/

func printMat(name string, m mat.Matrix) {
	prefix := mat.Prefix(strings.Repeat(" ", utf8.RuneCountInString(name)+3))
	fmt.Printf("%v = %v\n\n", name, mat.Formatted(m, prefix))
}

func printIndexVec(name string, values []string, keyVec mat.Matrix) {
	r, _ := keyVec.Dims()
	fmt.Printf("%v = [ ", name)
	for i := 0; i < r; i++ {
		fmt.Printf("%v", values[int(keyVec.At(i, 0))])
		if i != r-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print(" ]\n")
}

func main() {
	fmt.Printf("using gonum\n")

	vars := []string{"x1", "x2", "x3", "s1", "s2"}
	fmt.Printf("vars = %v\n", vars)

	c := mat.NewVecDense(5, []float64{-2, -3, -4, 0, 0})
	printMat("c", c)

	A := mat.NewDense(2, 5, []float64{
		3, 2, 1, 1, 0,
		2, 5, 3, 0, 1,
	})
	printMat("A", A)

	b := mat.NewVecDense(2, []float64{10, 15})
	printMat("b", b)

	constrCount, constrCols := A.Dims()
	if constrCols != len(vars) {
		panic("columns of A do not match variable count")
	}

	bSize, _ := b.Dims()
	if constrCount != bSize {
		panic("rows of A do not match size of b")
	}

	// basis := mat.NewVecDense(constrCount, []float64{3, 4})
	// printIndexVec("basis", vars, basis)

	// B := lp.NewBasis(A, []int{3, 4})
	B := lp.NewBasis(A, nil)
	fmt.Printf("B = %v\n", B)
	fmt.Println(B.Dims())
	printMat("B", B)

	tbl := lp.NewTableau(vars, A, c, b, B)
	tbl.Print()

	// B.Set(0, 0, 9)
	// printMat("B", B)
	// printMat("A", A)

	// B.SetVar(0, 0)
	// printMat("B", B)

	// diag := mat.NewDiagDense(5, []float64{1, 2, 3, 4, 5})
	// matx := mat.DenseCopyOf(diag)
	// printMat("matx", matx)
	// matx.Set(0, 0, 10)
	// printMat("matx", matx)

	// x := mat.NewVecDense(5, []float64{13, 3, 7, 12, 1})
	// y := mat.NewVecDense(5, []float64{7, 7, 7, 7, 7})
	// x.DivElemVec(x, y)
	// x.MulElemVec(x, y)
	// printMat("x", x)
}
