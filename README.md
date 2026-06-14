# zero — a tiny math CLI in Go

`zero` is a command-line calculator built with [Cobra](https://github.com/spf13/cobra).
It supports the usual arithmetic operations as dedicated subcommands and also ships
a full expression evaluator (`calc`) that understands operator precedence and
parentheses.

## Install / build

```sh
git clone https://github.com/RandomMoaz/CommandLine.git
cd CommandLine
go build -o zero .
./zero --help
```

Requires Go 1.24 or newer (see `go.mod`).

## Commands

| Command     | Aliases                   | Description                                   |
| ----------- | ------------------------- | --------------------------------------------- |
| `add`       | `addition`, `plus`        | Add two numbers                               |
| `subtract`  | `sub`, `minus`            | Subtract the second number from the first     |
| `multiply`  | `mul`, `mult`, `times`    | Multiply two numbers                          |
| `divide`    | `div`                     | Divide the first number by the second         |
| `power`     | `pow`, `exp`              | Raise a base to an exponent                   |
| `sqrt`      | `squareroot`              | Square root of a non-negative number          |
| `mod`       | `modulo`, `remainder`     | Remainder of `a / b` (floats allowed)         |
| `factorial` | `fact`, `fac`             | `n!` for non-negative integers up to 170      |
| `calc`      | `eval`, `compute`         | Evaluate a full arithmetic expression         |
| `version`   |                           | Print the CLI version                         |

Run `zero <command> --help` for details on any subcommand.

## Examples

```sh
$ zero add 2 3
2 + 3 = 5

$ zero subtract 10 4
10 - 4 = 6

$ zero multiply 6 7
6 * 7 = 42

$ zero divide 22 7
22 / 7 = 3.142857142857143

$ zero power 2 10
2 ^ 10 = 1024

$ zero sqrt 2
sqrt(2) = 1.4142135623730951

$ zero mod 10 3
10 mod 3 = 1

$ zero factorial 10
10! = 3628800
```

### `calc` — expression evaluator

`calc` evaluates a full arithmetic expression in one go. It supports:

- Operators: `+`, `-`, `*`, `/`, `%`, `^` (power, right-associative), and unary `-`
- Parentheses for grouping
- Integer, decimal, and scientific notation (`1.5e3`) numbers
- Standard operator precedence: `^` > `*` `/` `%` > `+` `-`

```sh
$ zero calc "2 + 3 * 4"
2 + 3 * 4 = 14

$ zero calc "(2 + 3) * 4"
(2 + 3) * 4 = 20

$ zero calc "2 ^ 3 ^ 2"          # right-associative: 2^(3^2) = 2^9
2 ^ 3 ^ 2 = 512

$ zero calc "-5 + 10 / 2"
-5 + 10 / 2 = 0

$ zero calc "2 + 3 * (4 - 1) ^ 2"
2 + 3 * (4 - 1) ^ 2 = 29
```

Quoting the expression is recommended (some shells interpret `*` and `(`),
but unquoted multi-arg forms also work: `zero calc 2 + 3`.

## Error handling

Invalid input produces a clear error on `stderr` and a non-zero exit code,
so `zero` composes cleanly with shell scripts and `&&` chains:

```sh
$ zero add abc 3
error: first value "abc" is not a number
$ echo $?
1

$ zero divide 5 0
error: cannot divide by zero
$ echo $?
1

$ zero sqrt -4
error: cannot take square root of a negative number (-4)
```

## Versioning

The version string is injected at build time via `-ldflags`:

```sh
go build -ldflags "-X go-cl/cmd.Version=v0.2.0" -o zero .
./zero version
# zero v0.2.0
```

Local builds without the flag report `zero dev`.

## Tests

```sh
go test ./...
```

Covers every math helper, the factorial edge cases (negatives, non-integers,
overflow), and the expression parser (precedence, associativity, unary minus,
parens, divide-by-zero, malformed input).

## Project layout

```
.
├── main.go              # entry point
├── go.mod, go.sum
└── cmd/
    ├── root.go          # root cobra command
    ├── math.go          # arithmetic helpers + result formatting
    ├── expr.go          # recursive-descent expression parser
    ├── add.go           # subcommand definitions
    ├── subtract.go
    ├── multiply.go
    ├── divide.go
    ├── power.go
    ├── sqrt.go
    ├── mod.go
    ├── factorial.go
    ├── calc.go
    ├── version.go
    ├── math_test.go
    └── expr_test.go
```

## Notes on this iteration

This version fixes three pre-existing build issues alongside the new features:

1. **File casing.** `Main.GO` and `cmd/root.GO` were renamed to lowercase `.go`.
   The Go toolchain only recognizes the lowercase extension, so the repo did not
   build on Linux or macOS (it appeared to work on Windows only because NTFS is
   case-insensitive).
2. **Import path.** `main.go` now imports `"go-cl/cmd"` to match the module
   name declared in `go.mod`. The previous `"GO-CL/cmd"` did not resolve.
3. **Silent failures on bad input.** The old helpers printed an error and
   returned an empty string, which the caller would then print as the "result."
   Helpers now return `error` values and the CLI exits with a non-zero status
   on any failure.
