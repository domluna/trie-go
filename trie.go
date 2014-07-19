package trie

// Trie data structure http://en.wikipedia.org/wiki/Trie.
type Trie struct {
	children map[rune]*Trie
	value    interface{}
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
	}
}

func search(t *Trie, key string) *Trie {
	for _, c := range key {
		if _, ok := t.children[c]; !ok {
			t.children[c] = NewTrie()
		}
		t = t.children[c]
	}
	return t
}

// Add adds a value to the trie with the given key.
func (t *Trie) Add(key string, value interface{}) {
	t = search(t, key)
	t.value = value
}

// Get retrieves the value associated with the full string.
//
// If the key is present the tuple (value, true) will be returned.
// If the key is not present the tuple (nil, false) will be returned.
//
// Suppose we add ("one", 55), "one" corresponds with the value 55.
// Here "o" is also in trie, however no value has been associated with "o".
// Therefore if we query for "o" (nil, false) will be returned.
func (t *Trie) Get(key string) (interface{}, bool) {
	t = search(t, key)
	if t.value == nil {
		return nil, false
	}
	return t.value, true
}
