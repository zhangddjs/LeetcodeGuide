.PHONY: gen clean test

# Generate today's practice file
gen:
	@cd cmd && python3 daily_gen.py

# Archive generated files
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "trie_*.py" ! -name "trie_template.tmpl" ! -name "trie_test.py" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"

# Run tests
test:
	@python3 -m pytest trie_test.py -v || python3 -m unittest trie_test.py -v
