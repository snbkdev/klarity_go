// 7.9 Пример: вычислитель выражения
package main

import (
	"fmt"
	"math"
	//"net/http"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

type literal float64

type unary struct {
	op rune
	x Expr
}

type binary struct {
	op rune
	x, y Expr
}

type call struct {
	fn string
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}


func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("некорректный унарный оператор %q", u.op)
	}

	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("некорректный бинарный оператор %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("неизвестная функция %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("вызов %s имеет %d вместо %d аргументов", c.fn, len(c.args), arity)
	}

	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

// func parseAndCheck(s string) (Expr, error) {
// 	if s == "" {
// 		return nil, fmt.Errorf("пустое выражение")
// 	}

// 	expr, err := Parse(s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	vars := make(map[Var]bool)
// 	if err := expr.Check(vars); err != nil {
// 		return nil, err
// 	}

// 	for v := range vars {
// 		if v != "x" && v != "y" && v != "r" {
// 			return nil, fmt.Errorf("неизвестная переменная %s", v)
// 		}
// 	}

// 	return expr, nil
// }

// func plot(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	expr, err := parseAndCheck(r.Form.Get("expr"))
// 	if err != nil {
// 		http.Error(w, ""+ err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "image/svg+xml")
// 	surface(w, func(x, y float64) float64 {
// 		r := math.Hypot(x, y)
// 		return expr.Eval(Env{"x": x, "y": y, "r": r})
// 	})
// }