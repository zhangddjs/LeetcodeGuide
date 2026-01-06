// Generated on 2026-01-02 11:16:34
// Daily practice file: trie_20260102.go

package trie

type Trie struct {
	IsEndWord bool
	Members   map[byte]*Trie
}

// --------------------------------------------------------------

func (t *Trie) Insert(word string) {
	n, m := t, t.Members
	for i := range word {
		ok := false
		n, ok = m[word[i]]
		if !ok {
			m[word[i]] = &Trie{false, make(map[byte]*Trie)}
		}
		n, m = m[word[i]], m[word[i]].Members
	}
	n.IsEndWord = true
}

// --------------------------------------------------------------

func (t *Trie) Search(word string) bool {
	n, m := t, t.Members
	for i := range word {
		ok := false
		n, ok = m[word[i]]
		if !ok {
			return false
		}
		m = n.Members
	}
	return n.IsEndWord
}

// --------------------------------------------------------------

func (t *Trie) StartsWith(prefix string) bool {
	m := t.Members
	for i := range prefix {
		n, ok := m[prefix[i]]
		if !ok {
			return false
		}
		m = n.Members
	}
	return true
}

// --------------------------------------------------------------

func NewTrie() *Trie {
	return &Trie{false, make(map[byte]*Trie)}
}
