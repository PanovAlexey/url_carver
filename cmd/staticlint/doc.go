/*
Package main of staticlint implements different code linters:
- Standard static packet analyzers (golang.org/x/tools/go/analysis/passes)
- All SA linters and a few others from staticcheck
- Couple of other linter packages: github.com/go-critic/go-critic and github.com/gostaticanalysis/nilerr
- Custom linter (AnalyzerProhibitExitInMain) that checks if main package code has os.Exit function calling

This multichecker can be run from the root of the project via:

go run cmd/staticlint/main.go ./...
*/
package main
