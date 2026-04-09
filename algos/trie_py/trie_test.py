import unittest
import sys
import os
import importlib.util
from typing import Protocol


class TrieProtocol(Protocol):
    def insert(self, word: str) -> None: ...
    def search(self, word: str) -> bool: ...
    def starts_with(self, prefix: str) -> bool: ...


def load_trie_module():
    """Dynamically load the most recent trie_*.py file"""
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    files = [f for f in os.listdir(parent_dir) if f.startswith('trie_') and f.endswith('.py')]

    if not files:
        raise FileNotFoundError("No trie_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(parent_dir, latest_file)

    spec = importlib.util.spec_from_file_location("trie_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module.Trie


class TestTrie(unittest.TestCase):

    def setUp(self):
        self.Trie = load_trie_module()

    def test_insert_empty(self):
        trie = self.Trie()
        # Just verify it doesn't crash
        pass

    def test_insert_single_word(self):
        trie = self.Trie()
        trie.insert("hello")

    def test_insert_multiple_words(self):
        trie = self.Trie()
        for word in ["hello", "world", "test"]:
            trie.insert(word)

    def test_insert_overlapping_prefixes(self):
        trie = self.Trie()
        for word in ["cat", "car", "card", "care", "careful"]:
            trie.insert(word)

    def test_insert_duplicate_words(self):
        trie = self.Trie()
        for word in ["test", "test", "hello", "test"]:
            trie.insert(word)

    def test_insert_single_character(self):
        trie = self.Trie()
        for word in ["a", "b", "c"]:
            trie.insert(word)

    def test_insert_nested_words(self):
        trie = self.Trie()
        for word in ["a", "ab", "abc", "abcd"]:
            trie.insert(word)

    def test_search_empty_trie(self):
        trie = self.Trie()
        self.assertFalse(trie.search("hello"))

    def test_search_exact_match(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertTrue(trie.search("hello"))

    def test_search_no_match(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertFalse(trie.search("world"))

    def test_search_prefix_exists_but_not_word(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertFalse(trie.search("hell"))

    def test_search_word_exists_with_longer_word(self):
        trie = self.Trie()
        trie.insert("hello")
        trie.insert("helloworld")
        self.assertTrue(trie.search("hello"))

    def test_search_case_sensitive(self):
        trie = self.Trie()
        trie.insert("Hello")
        self.assertFalse(trie.search("hello"))

    def test_search_empty_string(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertFalse(trie.search(""))

    def test_search_single_character(self):
        trie = self.Trie()
        trie.insert("a")
        self.assertTrue(trie.search("a"))

    def test_starts_with_empty_trie(self):
        trie = self.Trie()
        self.assertFalse(trie.starts_with("hello"))

    def test_starts_with_exact_prefix(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertTrue(trie.starts_with("hell"))

    def test_starts_with_no_prefix_match(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertFalse(trie.starts_with("world"))

    def test_starts_with_exact_word_match(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertTrue(trie.starts_with("hello"))

    def test_starts_with_empty_prefix(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertTrue(trie.starts_with(""))

    def test_starts_with_multiple_words_with_prefix(self):
        trie = self.Trie()
        for word in ["cat", "car", "card"]:
            trie.insert(word)
        self.assertTrue(trie.starts_with("car"))

    def test_starts_with_multiple_words_no_prefix(self):
        trie = self.Trie()
        for word in ["cat", "car", "card"]:
            trie.insert(word)
        self.assertFalse(trie.starts_with("cad"))

    def test_starts_with_single_character_prefix(self):
        trie = self.Trie()
        trie.insert("hello")
        self.assertTrue(trie.starts_with("h"))

    def test_starts_with_case_sensitive_prefix(self):
        trie = self.Trie()
        trie.insert("Hello")
        self.assertFalse(trie.starts_with("hello"))

    def test_comprehensive(self):
        trie = self.Trie()

        words = ["apple", "app", "apricot", "banana", "band", "bandana"]
        for word in words:
            trie.insert(word)

        search_tests = {
            "app": True,
            "apple": True,
            "appl": False,
            "banana": True,
            "band": True,
            "ban": False,
            "orange": False,
        }

        for word, expected in search_tests.items():
            result = trie.search(word)
            self.assertEqual(result, expected,
                           f"Search('{word}') = {result}, want {expected}")

        prefix_tests = {
            "app": True,
            "appl": True,
            "ban": True,
            "band": True,
            "orange": False,
            "": True,
        }

        for prefix, expected in prefix_tests.items():
            result = trie.starts_with(prefix)
            self.assertEqual(result, expected,
                           f"StartsWith('{prefix}') = {result}, want {expected}")

    def test_sequential_operations(self):
        trie = self.Trie()

        # Initial search on empty trie
        self.assertFalse(trie.search("hello"), "Search('hello') on empty trie should be False")

        # Insert first item
        trie.insert("hello")

        # Search again after insert
        self.assertTrue(trie.search("hello"), "Search('hello') after insert should be True")

        # StartsWith test
        self.assertTrue(trie.starts_with("hell"), "StartsWith('hell') should be True")

        # Insert second item
        trie.insert("helicopter")

        # StartsWith test again after second insert
        self.assertTrue(trie.starts_with("hel"), "StartsWith('hel') after second insert should be True")

        # Additional verification
        self.assertTrue(trie.search("helicopter"), "Search('helicopter') should be True")
        self.assertFalse(trie.search("helicopte"), "Search('helicopte') should be False")
        self.assertTrue(trie.starts_with("helic"), "StartsWith('helic') should be True")


if __name__ == '__main__':
    unittest.main()
