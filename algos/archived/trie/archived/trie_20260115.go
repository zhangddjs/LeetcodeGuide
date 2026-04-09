// Generated on 2026-01-15 07:49:39
// Daily practice file: trie_20260115.go

package trie

type Trie struct {
	IsEndOfWord bool
	Members     map[byte]*Trie
}

// --------------------------------------------------------------

// In a trie, the insert operation involves adding a word character by character. You start at the root and for each character in the word,
// you move down the corresponding child node.
// If a node for that character doesn’t exist, you create a new node.
// Once you reach the end of the word, you mark the final node as a leaf node to indicate that a word ends there.
// This way, you build a path in the trie that represents the word.
// If you insert another word with a common prefix,
// you’ll share nodes for that prefix, which makes tries efficient for storing and searching words.
func (t *Trie) Insert(word string) {
	n, m := t, t.Members
	for i := range word {
		ok := false
		n, ok = m[word[i]]
		if !ok {
			n = NewTrie()
			m[word[i]] = n
		}
		m = n.Members
	}
	n.IsEndOfWord = true
}

// --------------------------------------------------------------

//  1. Start at the Root: The search begins at the trie’s root node.
//  2. Traverse Characters: For each character in the word, you move down the trie by following the corresponding child node.
//     If at any point a character is not found, the search immediately fails and returns false.
//  3. Check the End of the Word: Once all characters of the word are found, you reach the final node.
//  4. Verify the isEndOfWord Flag: At the final node, you check the "isEndOfWord" flag.
//     If it’s set to true, that means the word exists in the trie. If it’s false, then the word is only a prefix and not a complete word.
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
	return n.IsEndOfWord
}

// --------------------------------------------------------------

/*
* The startsWith function in a trie is used to check if there’s any word in the trie that begins with a given prefix.
* The process is similar to the search function, but with a key difference:
* 1. Start at the Root: As with the search function, you begin at the root of the trie.
* 2. Traverse the Prefix: You then follow the path of characters in the prefix.
*    For each character, you move down to the corresponding child node.
* 3. Check for Missing Characters: If at any point a character in the prefix is not found, the function returns false,
*    indicating that no word in the trie starts with that prefix.
* 4. Confirm the Prefix Exists: If you successfully traverse all characters of the prefix,
*    it means that there is at least one word in the trie that begins with that prefix. The function then returns true.
* In essence, the startsWith function only verifies the existence of the prefix and doesn’t require the "isEndOfWord" flag.
* It’s mainly about confirming that the path for the prefix exists in the trie.
 */
func (t *Trie) StartsWith(prefix string) bool {
	n, m := t, t.Members
	for i := range prefix {
		ok := false
		n, ok = m[prefix[i]]
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
