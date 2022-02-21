package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node interface {
	String() string
	eval() float64
}

type operator struct {
	operand1, operand2 node
	symbol             string
	operation          func(o1, o2 float64) float64
}

func (o operator) String() string {
	return fmt.Sprintf("%s %s %s", o.operand1.String(), o.symbol, o.operand2.String())
}

func (o operator) eval() float64 {
	return o.operation(o.operand1.eval(), o.operand2.eval())
}

type number struct {
	value float64
}

func (n number) String() string {
	return strconv.FormatFloat(n.value, 'g', -1, 64)
}

func (n number) eval() float64 {
	return n.value
}

func parseNumber(s string) number {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("error: expected number, got \"%s\"\n", s)
		os.Exit(1)
	}
	return number{
		value: v,
	}
}

func splitPreserve(s string, chars string) (string, byte, string) {
	idx := strings.IndexAny(s, chars)
	if idx == -1 {
		return s, 0, ""
	}
	return s[:idx], s[idx], s[idx+1:]
}

func parseProduct(s string) node {
	s1, op, s2 := splitPreserve(s, "*/")
	if s1 == "" {
		fmt.Printf("error: cannot parse operation: %s\n", s)
		os.Exit(2)
	}
	if s2 == "" {
		return parseNumber(s)
	}
	return operator{
		operand1: parseNumber(s1),
		operand2: parseProduct(s2),
		symbol:   string(op),
		operation: map[byte]func(o1, o2 float64) float64{
			'*': func(o1, o2 float64) float64 {
				return o1 * o2
			},
			'/': func(o1, o2 float64) float64 {
				return o1 / o2
			},
		}[op],
	}
}

func parseSum(s string) node {
	s1, op, s2 := splitPreserve(s, "+-")
	if s1 == "" {
		fmt.Printf("error: cannot parse operation: %s\n", s)
		os.Exit(2)
	}
	if s2 == "" {
		return parseProduct(s)
	}
	return operator{
		operand1: parseProduct(s1),
		operand2: parseSum(s2),
		symbol:   string(op),
		operation: map[byte]func(o1, o2 float64) float64{
			'+': func(o1, o2 float64) float64 {
				return o1 + o2
			},
			'-': func(o1, o2 float64) float64 {
				return o1 - o2
			},
		}[op],
	}
}

func parse(s string) node {
	return parseSum(s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	s := strings.ReplaceAll(string(line), " ", "")
	expr := parse(s)
	fmt.Printf("%s = %g", expr.String(), expr.eval())
}
