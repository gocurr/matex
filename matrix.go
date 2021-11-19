// matext is the extension for gonum.mat
// usage like jblas in java

package matext

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// BetterPrint printing in a pretty way
func BetterPrint(dense *mat.Dense) {
	fmt.Printf("Matrix :\n%v\n\n", mat.Formatted(dense, mat.Prefix(""), mat.Excerpt(0)))
	rows, cols := dense.Dims()
	fmt.Println("Matrix: rows: ", rows)
	fmt.Println("Matrix: cols: ", cols)
}

// Flatten flattens [][] -> r, c, []
func Flatten(f [][]float64) (r, c int, d []float64) {
	r = len(f)
	if r == 0 {
		panic("bad test: no row")
	}
	c = len(f[0])
	d = make([]float64, 0, r*c)
	for _, row := range f {
		if len(row) != c {
			panic("bad test: ragged input")
		}
		d = append(d, row...)
	}
	return r, c, d
}

// PowByN returns matrix whose elements will be powered by number `n`
func PowByN(dense *mat.Dense, n float64) *mat.Dense {
	return BiForEach(dense, n, math.Pow)
}

// Square return matrix whose elements will be squared
func Square(dense *mat.Dense) *mat.Dense {
	return ForEach(dense, math.Sqrt)
}

// Multiply returns matrix whose elements will be multiplied by `x`
func Multiply(dense *mat.Dense, x float64) *mat.Dense {
	return BiForEach(dense, x, func(a, b float64) float64 {
		return a * b
	})
}

// ForEach returns matrix whose elements will be invoked by `fn` with 1 parameter
func ForEach(dense *mat.Dense, fn func(item float64) float64) *mat.Dense {
	var builder []float64
	r, c := dense.Dims()
	data := dense.RawMatrix().Data
	for i := range data {
		builder = append(builder, fn(data[i]))
	}
	return mat.NewDense(r, c, builder)
}

// BiForEach returns matrix whose elements will be invoked by `fn` with 2 parameters
func BiForEach(dense *mat.Dense, another float64, fn func(a, b float64) float64) *mat.Dense {
	var builder []float64
	r, c := dense.Dims()
	data := dense.RawMatrix().Data
	for i := range data {
		builder = append(builder, fn(data[i], another))
	}
	return mat.NewDense(r, c, builder)
}

// MatrixTruth returns matrix whose elements will be `trued`, 0 -> 0, !0 -> 1
func MatrixTruth(dense *mat.Dense) *mat.Dense {
	return ForEach(dense, func(item float64) float64 {
		if item == 0 {
			return 0
		} else {
			return 1
		}
	})
}

// MaxColumn returns matrix which has 1 row with max column value
func MaxColumn(dense *mat.Dense) *mat.Dense {
	return MatrixColumnForEach(dense, mat.Max)
}

// MinColumn returns matrix which has 1 row with min column value
func MinColumn(dense *mat.Dense) *mat.Dense {
	return MatrixColumnForEach(dense, mat.Min)
}

// ColSum returns matrix which has 1 row with sum column value
func ColSum(dense *mat.Dense) *mat.Dense {
	return MatrixColumnForEach(dense, mat.Sum)
}

// MatrixColumnForEach matrix column for-each function
func MatrixColumnForEach(dense *mat.Dense, fn func(a mat.Matrix) float64) *mat.Dense {
	var builder []float64
	_, c := dense.Dims()
	for i := 0; i < c; i++ {
		colView := dense.ColView(i)
		out := fn(colView)
		builder = append(builder, out)
	}

	return mat.NewDense(1, c, builder)
}
