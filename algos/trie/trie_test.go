package trie

import "testing"

type trieInterface interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

var insertTestCases = []struct {
	name  string
	words []string
}{
	{"empty", []string{}},
	{"single word", []string{"hello"}},
	{"multiple words", []string{"hello", "world", "test"}},
	{"overlapping prefixes", []string{"cat", "car", "card", "care", "careful"}},
	{"duplicate words", []string{"test", "test", "hello", "test"}},
	{"single character", []string{"a", "b", "c"}},
	{"nested words", []string{"a", "ab", "abc", "abcd"}},
}

var searchTestCases = []struct {
	name     string
	words    []string
	search   string
	expected bool
}{
	{"empty trie", []string{}, "hello", false},
	{"exact match", []string{"hello"}, "hello", true},
	{"no match", []string{"hello"}, "world", false},
	{"prefix exists but not word", []string{"hello"}, "hell", false},
	{"word exists with longer word", []string{"hello", "helloworld"}, "hello", true},
	{"case sensitive", []string{"Hello"}, "hello", false},
	{"empty string search", []string{"hello"}, "", false},
	{"single character", []string{"a"}, "a", true},
}

var startsWithTestCases = []struct {
	name     string
	words    []string
	prefix   string
	expected bool
}{
	{"empty trie", []string{}, "hello", false},
	{"exact prefix", []string{"hello"}, "hell", true},
	{"no prefix match", []string{"hello"}, "world", false},
	{"exact word match", []string{"hello"}, "hello", true},
	{"empty prefix", []string{"hello"}, "", true},
	{"multiple words with prefix", []string{"cat", "car", "card"}, "car", true},
	{"single character prefix", []string{"hello"}, "h", true},
	{"case sensitive prefix", []string{"Hello"}, "hello", false},
}

func testTrieInsert(t *testing.T, newTrie func() trieInterface) {
	for _, tc := range insertTestCases {
		t.Run(tc.name, func(t *testing.T) {
			trie := newTrie()
			for _, word := range tc.words {
				trie.Insert(word)
			}
		})
	}
}

func testTrieSearch(t *testing.T, newTrie func() trieInterface) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			trie := newTrie()
			for _, word := range tc.words {
				trie.Insert(word)
			}
			result := trie.Search(tc.search)
			if result != tc.expected {
				t.Errorf("Search(%q) = %v, want %v", tc.search, result, tc.expected)
			}
		})
	}
}

func testTrieStartsWith(t *testing.T, newTrie func() trieInterface) {
	for _, tc := range startsWithTestCases {
		t.Run(tc.name, func(t *testing.T) {
			trie := newTrie()
			for _, word := range tc.words {
				trie.Insert(word)
			}
			result := trie.StartsWith(tc.prefix)
			if result != tc.expected {
				t.Errorf("StartsWith(%q) = %v, want %v", tc.prefix, result, tc.expected)
			}
		})
	}
}

func testTrieComprehensive(t *testing.T, newTrie func() trieInterface) {
	trie := newTrie()

	words := []string{"apple", "app", "apricot", "banana", "band", "bandana"}
	for _, word := range words {
		trie.Insert(word)
	}

	searchTests := map[string]bool{
		"app":    true,
		"apple":  true,
		"appl":   false,
		"banana": true,
		"band":   true,
		"ban":    false,
		"orange": false,
	}

	for word, expected := range searchTests {
		if result := trie.Search(word); result != expected {
			t.Errorf("Search(%q) = %v, want %v", word, result, expected)
		}
	}

	prefixTests := map[string]bool{
		"app":    true,
		"appl":   true,
		"ban":    true,
		"band":   true,
		"orange": false,
		"":       true,
	}

	for prefix, expected := range prefixTests {
		if result := trie.StartsWith(prefix); result != expected {
			t.Errorf("StartsWith(%q) = %v, want %v", prefix, result, expected)
		}
	}
}

func testTrieSequentialOperations(t *testing.T, newTrie func() trieInterface) {
	trie := newTrie()

	// Initial search on empty trie
	if result := trie.Search("hello"); result != false {
		t.Errorf("Search('hello') on empty trie = %v, want false", result)
	}

	// Insert first item
	trie.Insert("hello")

	// Search again after insert
	if result := trie.Search("hello"); result != true {
		t.Errorf("Search('hello') after insert = %v, want true", result)
	}

	// StartsWith test
	if result := trie.StartsWith("hell"); result != true {
		t.Errorf("StartsWith('hell') = %v, want true", result)
	}

	// Insert second item
	trie.Insert("helicopter")

	// StartsWith test again after second insert
	if result := trie.StartsWith("hel"); result != true {
		t.Errorf("StartsWith('hel') after second insert = %v, want true", result)
	}

	// Additional verification
	if result := trie.Search("helicopter"); result != true {
		t.Errorf("Search('helicopter') = %v, want true", result)
	}
	if result := trie.Search("helicopte"); result != false {
		t.Errorf("Search('helicopte') = %v, want false", result)
	}
	if result := trie.StartsWith("helic"); result != true {
		t.Errorf("StartsWith('helic') = %v, want true", result)
	}
}

// Uncomment and implement when trie implementation is available
func TestTrie(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		testTrieInsert(t, func() trieInterface { return NewTrie() })
	})
	t.Run("Search", func(t *testing.T) {
		testTrieSearch(t, func() trieInterface { return NewTrie() })
	})
	t.Run("StartsWith", func(t *testing.T) {
		testTrieStartsWith(t, func() trieInterface { return NewTrie() })
	})
	t.Run("Comprehensive", func(t *testing.T) {
		testTrieComprehensive(t, func() trieInterface { return NewTrie() })
	})
	t.Run("SequentialOperations", func(t *testing.T) {
		testTrieSequentialOperations(t, func() trieInterface { return NewTrie() })
	})
}
