package trie

type Trie struct {
	IsEndOfWord bool
	Members     map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	if t == nil {
		return
	}
	members := t.Members
	node, ok := t, false
	for i := 0; i < len(word); i++ {
		node, ok = members[word[i]]
		if !ok {
			node = &Trie{false, make(map[byte]*Trie)}
			members[word[i]] = node
		}
		members = node.Members
	}
	node.IsEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	if t == nil {
		return false
	}
	node, ok := t, false
	for i := range word {
		node, ok = node.Members[word[i]]
		if !ok {
			return false
		}
	}
	return node.IsEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	if t == nil {
		return false
	}
	node, ok := t, false
	for i := range prefix {
		node, ok = node.Members[prefix[i]]
		if !ok {
			return false
		}
	}
	return true
}

func NewTrie() *Trie {
	return &Trie{false, make(map[byte]*Trie)}
}
