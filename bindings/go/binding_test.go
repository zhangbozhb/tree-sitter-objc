package tree-sitter-objc

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_objc "github.com/tree-sitter/tree-sitter-objc/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_objc.Language())
	if language == nil {
		t.Errorf("Error loading OC grammar")
	}
}
