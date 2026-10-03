// Package newlineafterblock registers the newline-after-block linter as a
// golangci-lint module plugin.
//
// Build a custom golangci-lint binary that includes it by listing this module in
// .custom-gcl.yml and running `golangci-lint custom`; then enable the linter
// under linters.settings.custom in .golangci.yml.
package newlineafterblock

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("newline-after-block", newPlugin)
}

type plugin struct{}

func newPlugin(_ any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

// BuildAnalyzers returns the single analyzer this plugin provides.
func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{New()}, nil
}

// GetLoadMode requires full type information: the analyzer calls
// pass.TypesInfo.TypeOf to detect the error-check-then-defer exception.
func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
