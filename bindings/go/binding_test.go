package tree_sitter_gringo_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-gringo"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gringo.Language())
	if language == nil {
		t.Errorf("Error loading Gringo grammar")
	}
}
