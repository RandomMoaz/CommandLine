package cmd

// Tiny recursive-descent arithmetic expression evaluator.
//
// Grammar (operator precedence encoded by production depth):
//
//   expr   = term   { ("+" | "-") term }
//   term   = factor { ("*" | "/" | "%") factor }
//   factor = unary  [ "^" factor ]              // right-associative
//   unary  = ("+" | "-") unary | primary
//   primary = number | "(" expr ")"
//
// Numbers are float64. Whitespace is ignored.

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type exprParser struct {
	src string
	pos int
}

func (p *exprParser) peek() byte {
	if p.pos >= len(p.src) {
		return 0
	}
	return p.src[p.pos]
}

func (p *exprParser) skipSpace() {
	for p.pos < len(p.src) && unicode.IsSpace(rune(p.src[p.pos])) {
		p.pos++
	}
}

// accept consumes the given byte if it's next (after whitespace) and returns
// true; otherwise leaves the position unchanged and returns false.
func (p *exprParser) accept(c byte) bool {
	p.skipSpace()
	if p.peek() == c {
		p.pos++
		return true
	}
	return false
}

func (p *exprParser) parseExpr() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	for {
		p.skipSpace()
		c := p.peek()
		if c != '+' && c != '-' {
			return left, nil
		}
		p.pos++
		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}
		if c == '+' {
			left += right
		} else {
			left -= right
		}
	}
}

func (p *exprParser) parseTerm() (float64, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}
	for {
		p.skipSpace()
		c := p.peek()
		if c != '*' && c != '/' && c != '%' {
			return left, nil
		}
		p.pos++
		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		switch c {
		case '*':
			left *= right
		case '/':
			if right == 0 {
				return 0, fmt.Errorf("division by zero in expression")
			}
			left /= right
		case '%':
			if right == 0 {
				return 0, fmt.Errorf("modulo by zero in expression")
			}
			left = math.Mod(left, right)
		}
	}
}

func (p *exprParser) parseFactor() (float64, error) {
	base, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	p.skipSpace()
	if p.peek() == '^' {
		p.pos++
		// right-associative: 2^3^2 == 2^(3^2)
		exp, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		return math.Pow(base, exp), nil
	}
	return base, nil
}

func (p *exprParser) parseUnary() (float64, error) {
	p.skipSpace()
	if p.accept('+') {
		return p.parseUnary()
	}
	if p.accept('-') {
		v, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		return -v, nil
	}
	return p.parsePrimary()
}

func (p *exprParser) parsePrimary() (float64, error) {
	p.skipSpace()
	if p.accept('(') {
		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		if !p.accept(')') {
			return 0, fmt.Errorf("missing closing parenthesis at position %d", p.pos)
		}
		return v, nil
	}
	return p.parseNumber()
}

func (p *exprParser) parseNumber() (float64, error) {
	p.skipSpace()
	start := p.pos
	sawDigit := false
	sawDot := false
	for p.pos < len(p.src) {
		c := p.src[p.pos]
		if c >= '0' && c <= '9' {
			sawDigit = true
			p.pos++
			continue
		}
		if c == '.' && !sawDot {
			sawDot = true
			p.pos++
			continue
		}
		// Scientific notation: 1e10, 2.5E-3
		if (c == 'e' || c == 'E') && sawDigit {
			p.pos++
			if p.pos < len(p.src) && (p.src[p.pos] == '+' || p.src[p.pos] == '-') {
				p.pos++
			}
			continue
		}
		break
	}
	if !sawDigit {
		if p.pos >= len(p.src) {
			return 0, fmt.Errorf("unexpected end of expression, wanted a number")
		}
		return 0, fmt.Errorf("unexpected character %q at position %d", p.src[p.pos], p.pos)
	}
	text := p.src[start:p.pos]
	v, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q", text)
	}
	return v, nil
}

// Evaluate parses and evaluates an arithmetic expression and returns its
// result formatted as a string.
func Evaluate(expr string) (string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return "", fmt.Errorf("empty expression")
	}
	p := &exprParser{src: expr}
	v, err := p.parseExpr()
	if err != nil {
		return "", err
	}
	p.skipSpace()
	if p.pos != len(p.src) {
		return "", fmt.Errorf("unexpected trailing input starting at position %d: %q",
			p.pos, p.src[p.pos:])
	}
	return formatResult(v), nil
}
