Trie
====

Trie data structure

## Example

```go

import trie "github.com/domluna/trie-go"

func main() {
  t := trie.New()
  t.Add("Hello", 22)
  v, ok := t.Get("Hello") // v = 22, ok = true
  v, ok := t.Get("Helloa") // v = nil, ok = false
}
```

View the [docs](https://godoc.org/github.com/domluna/trie-go)
