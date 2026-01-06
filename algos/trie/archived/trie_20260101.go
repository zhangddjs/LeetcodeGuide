package trie

type Trie struct {
	IsEndOfWord bool
	Members     map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	n, m := t, t.Members
	for i := range word {
		ok := false
		n, ok = m[word[i]]
		if !ok {
			m[word[i]] = &Trie{false, make(map[byte]*Trie)}
		}
		n = m[word[i]]
		m = n.Members
	}
	n.IsEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	n, m := t, t.Members
	for i := range word {
		_, ok := m[word[i]]
		if !ok {
			return false
		}
		n, m = m[word[i]], m[word[i]].Members
	}
	return n.IsEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	m := t.Members
	for i := range prefix {
		_, ok := m[prefix[i]]
		if !ok {
			return false
		}
		m = m[prefix[i]].Members
	}
	return true
}

func NewTrie() *Trie {
	return &Trie{false, make(map[byte]*Trie)}
}
