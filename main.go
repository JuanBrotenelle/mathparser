package mathparse

import (
	"github.com/Knetic/govaluate"
	"math"
	"sync"
)

type Parser struct {
	functions   map[string]govaluate.ExpressionFunction
	definitions map[string]string
	mu          sync.RWMutex
}

// New create a new parser
func New() *Parser {
	parser := &Parser{
		functions:   make(map[string]govaluate.ExpressionFunction),
		definitions: make(map[string]string),
	}
	parser.loadFunctions()
	parser.loadDefinitions()
	return parser
}

// loadFunctions there are sin, cos, tan, cot, asin... functions
func (p *Parser) loadFunctions() {
	p.functions = map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"asin": func(args ...interface{}) (interface{}, error) {
			return math.Asin(args[0].(float64)), nil
		},
		"sinh": func(args ...interface{}) (interface{}, error) {
			return math.Sinh(args[0].(float64)), nil
		},
		"asinh": func(args ...interface{}) (interface{}, error) {
			return math.Asinh(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"acos": func(args ...interface{}) (interface{}, error) {
			return math.Acos(args[0].(float64)), nil
		},
		"cosh": func(args ...interface{}) (interface{}, error) {
			return math.Cosh(args[0].(float64)), nil
		},
		"acosh": func(args ...interface{}) (interface{}, error) {
			return math.Acosh(args[0].(float64)), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			return math.Tan(args[0].(float64)), nil
		},
		"atan": func(args ...interface{}) (interface{}, error) {
			return math.Atan(args[0].(float64)), nil
		},
		"tanh": func(args ...interface{}) (interface{}, error) {
			return math.Tanh(args[0].(float64)), nil
		},
		"atanh": func(args ...interface{}) (interface{}, error) {
			return math.Atanh(args[0].(float64)), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"pow": func(args ...interface{}) (interface{}, error) {
			return math.Pow(args[0].(float64), args[1].(float64)), nil
		},
	}
}

func (p *Parser) loadDefinitions() {
	p.definitions = map[string]string{
		"sin":   "sin(x)",
		"asin":  "asin(x)",
		"sinh":  "sinh(x)",
		"asinh": "asinh(x)",
		"cos":   "cos(x)",
		"acos":  "acos(x)",
		"cosh":  "cosh(x)",
		"acosh": "acosh(x)",
		"tan":   "tan(x)",
		"atan":  "atan(x)",
		"tanh":  "tanh(x)",
		"atanh": "atanh(x)",
		"exp":   "exp(x)",
		"sqrt":  "sqrt(x)",
		"pow":   "pow(x)",
		"log":   "log(x)",
		"log10": "log10(x)",
	}
}

// AddFunction add own function
func (p *Parser) AddFunction(name, definition string, function govaluate.ExpressionFunction) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.functions[name] = function

	p.definitions[name] = definition
}

// GetFunction return value will be a function
func (p *Parser) GetFunction(name string) string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.definitions[name]
}

// Parse to parse a string
func (p *Parser) Parse(expressionStr string) (func(params map[string]interface{}) (float64, error), error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expressionStr, p.functions)
	if err != nil {
		return nil, err
	}

	return func(params map[string]interface{}) (float64, error) {
		result, err := expression.Evaluate(params)
		if err != nil {
			return 0, err
		}
		return result.(float64), nil
	}, nil
}
