# MyMathParser

I was inspired by this [GoValuate](https://github.com/Knetic/govaluate) package, that's the basis. I added several basic math functions at once, like: sin, cos, tan, asin, etc. You can find the full list below. You can also add your own function with a description of what it does.

### Functions that already exist
`sin`, `asin`, `sinh`, `asinh`, `cos`, `acos`, `cosh`, `acosh`, `tan`, `atan`, `tanh`, `atanh`, `exp`, `sqrt`, `pow`, `log`, `log10`

### Here the example

```go
package main

import (
	"fmt"
	"mathparse"
)

func main() {
	parser := mathparse.New()

	parser.AddFunction("cube", "cube(x) = x ^ 3", func(args ...interface{}) (interface{}, error) {
		x := args[0].(float64)
		return x * x * x, nil
	})

	expression := "(cos(x) + cube(y)) / (pow(z, 2) + 2)"
	function, err := parser.Parse(expression)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := map[string]interface{}{
		"x": 0.5,
		"y": 2.0,
		"z": 3.4,
	}
	result, err := function(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	fmt.Println(parser.GetFunction("cube"))
}
```

There are the functions `AddFunction(name, definition, func() () {})` and `GetFunction(name)` (it will reveal a definition).