package trie

import "testing"

var table1 = map[string]interface{}{
	"on":                2,
	"one":               10,
	"some other string": 55,
}

var table2 = map[string]interface{}{
	"blank1":  2,
	"one1":    10,
	"on1":     10,
	"blank55": 55,
	"o":       0,
}

func TestTrieBasics(t *testing.T) {
	trie := NewTrie()
	for k, v := range table1 {
		trie.Add(k, v)
	}

	for k, e := range table1 {
		v, _ := trie.Get(k)
		if v != e {
			t.Errorf("(Should be in Trie, key %v) expected %v, got %v", k, e, v)
		}
	}

	for k, _ := range table2 {
		_, ok := trie.Get(k)
		if ok {
			t.Errorf("(Should not be in Trie, key %v) expected %v, got %v", k, false, ok)
		}
	}
}

func TestSet(t *testing.T) {
	k := "foo"
	trie := NewTrie()
	trie.Add(k, 5)
	trie.Add(k, 123)

	v, _ := trie.Get(k)
	if v != 123 {
		t.Errorf("(Should be in Trie, key %v) expected %v, got %v", k, 123, v)
	}
}
