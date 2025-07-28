package admintwiglinter

import (
	"github.com/shyim/go-version"

	"github.com/heycart/heycart-cli/internal/html"
	"github.com/heycart/heycart-cli/internal/validation"
	"github.com/heycart/heycart-cli/internal/verifier/twiglinter"
)

type SkeletonBarFixer struct{}

func init() {
	twiglinter.AddAdministrationFixer(SkeletonBarFixer{})
}

func (s SkeletonBarFixer) Check(nodes []html.Node) []validation.CheckResult {
	var errors []validation.CheckResult
	html.TraverseNode(nodes, func(node *html.ElementNode) {
		if node.Tag == "sw-skeleton-bar" {
			errors = append(errors, validation.CheckResult{
				Message:    "sw-skeleton-bar is removed, use mt-skeleton-bar instead.",
				Severity:   "warn",
				Identifier: "sw-skeleton-bar",
				Line:       node.Line,
			})
		}
	})
	return errors
}

func (s SkeletonBarFixer) Supports(v *version.Version) bool {
	return twiglinter.Shopware67Constraint.Check(v)
}

func (s SkeletonBarFixer) Fix(nodes []html.Node) error {
	html.TraverseNode(nodes, func(node *html.ElementNode) {
		if node.Tag == "sw-skeleton-bar" {
			node.Tag = "mt-skeleton-bar"
		}
	})
	return nil
}
