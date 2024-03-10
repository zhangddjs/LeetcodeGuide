// WA

func shortestSubstrings(arr []string) []string {
	root := &TrieNode{make(map[byte]*TrieNode, 26)}
	trie := &Trie{root, make(map[byte][]*TrieNode)}
	for _, e := range arr {
		insert(e, trie)
	}

	res := make([]string, 0, len(arr))
	for _, e := range arr {
		minLen := len(e)
		lenStrs := make(map[int][]string)
		for i := 0; i < len(e); i++ {
			for j := i + 1; j < len(e); j++ {
				sub := e[i:j]
				if !isSubOfAnyOther(sub, trie) {
					minLen = min(minLen, len(sub))
					lenStrs[minLen] = append(lenStrs[minLen], sub)
				}
			}
		}
		if len(lenStrs[minLen]) == 0 {
			res = append(res, "")
		} else {
			sort.Strings(lenStrs[minLen])
			res = append(res, lenStrs[minLen][0])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Trie struct {
	Root             *TrieNode
	NodeMapOfInitial map[byte][]*TrieNode
}

type TrieNode struct {
	Children map[byte]*TrieNode
}

func insert(key string, trie *Trie) {
	byts := []byte(key)
	node := trie.Root
	for i := 0; i < len(byts); i++ {
		_, ok := node.Children[byts[i]]
		if !ok {
			node.Children[byts[i]] = &TrieNode{make(map[byte]*TrieNode, 26)}
		}
		node = node.Children[byts[i]]
		trie.NodeMapOfInitial[byts[i]] = append(trie.NodeMapOfInitial[byts[i]], node)
	}
}

func isSubOfAnyOther(key string, trie *Trie) bool {
	byts := []byte(key)
	initial := byts[0]
	nodes, ok := trie.NodeMapOfInitial[initial]
	if !ok {
		return false
	}
	for _, n := range nodes {
		if search(key[1:], n) {
			return true
		}
	}
	return false
}

func search(key string, root *TrieNode) bool {
	node := root
	byts := []byte(key)
	for i := 0; i < len(byts); i++ {
		_, ok := node.Children[byts[i]]
		if !ok {
			return false
		}
		node = node.Children[byts[i]]
	}
	return true
}