package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("dog")
	trie.Insert("dash")
	found := trie.Search("apple")
	if !found {
		t.Error("incorrect result: expected:  string to be found, output: not found")
	}
	found = trie.Search("dog")
	if !found {
		t.Error("incorrect result: expected:  string to be found, output: not found")
	}
	found = trie.Search("dash")
	if !found {
		t.Error("incorrect result: expected:  string to be found, output: not found")
	}
	found = trie.Search("App")
	if found {
		t.Error("incorrect result: expected:  string to be not found, output: found")
	}

	startsWith := trie.StartsWith("das")
	if !startsWith {
		t.Error("incorrect result: expected:  true, output: false")
	}
}
