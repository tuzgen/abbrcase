package main

import (
	"flag"
	"go/ast"
	"regexp"
	"strings"

	"github.com/tuzgen/abbrcase/config"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer)
}

var analyzer = &analysis.Analyzer{
	Name:  "abbrcase",
	Doc:   "check for capital abbreviation abuse",
	Flags: options(),
	Run:   run,
}

func options() flag.FlagSet {
	options := flag.NewFlagSet("", flag.ExitOnError)

	options.String("abbrs", "", "comma separated list of abbreviations to include in analysis")

	return *options
}

func run(pass *analysis.Pass) (interface{}, error) {
	// This would've worked perfectly but the negative lookahead is not supported in golang
	// See: https://swtch.com/~rsc/regexp/regexp3.html
	// regex := regexp.MustCompile(`([A-Z]+(?![a-z])|[A-Z][a-z]+|[a-z]+)`)

	// Fortunately, our use case does not need crazy regex shenanigans.
	regex := regexp.MustCompile(`([A-Z]+[a-z]*|[a-z]+)`)
	cfg := config.WithOptions(
		config.WithAbbrs(
			pass.Analyzer.Flags.Lookup("abbrs").Value.String(),
		),
	)

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {

			if identifier, ok := node.(*ast.Ident); ok {
				allMatches := regex.FindAll([]byte(identifier.String()), 10)
				for _, match := range allMatches {
					for _, abbr := range cfg.Abbrs {
						if cfg.Violates(string(match)) {
							pass.Reportf(identifier.Pos(), "use all caps abbreviations: %s should be %s", match, strings.ToUpper(abbr))
						}
					}
				}

			}
			return true
		})
	}
	return nil, nil
}
