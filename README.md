# matex - extension of gonum.mat

The usage of `matex` is like `jblas` (partial implementation) in java.

To download, run:

```bash
go get -u github.com/gocurr/matex
```

Import it in your program as:

```go
import "github.com/gocurr/matex"
```

It requires Go 1.11 or later due to usage of Go Modules.

- Usage:

```go
var array = [][]float64{
    {1, 2, 3},
    {3, 2, 1},
}
r, c, flatten := matext.Flatten(array)
dense := mat.NewDense(r, c, flatten)
matext.BetterPrint(dense)

power := matext.PowByN(dense, 2)
matext.BetterPrint(power)

mul := matext.Multiply(dense, 2)
matext.BetterPrint(mul)

max := matext.MaxColumn(dense)
matext.BetterPrint(max)

min := matext.MinColumn(dense)
matext.BetterPrint(min)

truth := matext.MatrixTruth(dense)
matext.BetterPrint(truth)

square := matext.Square(dense)
matext.BetterPrint(square)

sum := matext.ColSum(dense)
matext.BetterPrint(sum)
```
