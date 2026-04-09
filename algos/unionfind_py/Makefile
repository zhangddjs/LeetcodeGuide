.PHONY: gen clean test

# Generate today's practice file
gen:
	@cd cmd && python3 daily_gen.py

# Archive generated files
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "unionfind_*.py" ! -name "unionfind_template.tmpl" ! -name "unionfind_test.py" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"

# Run tests
test:
	@python3 -m pytest unionfind_test.py -v || python3 -m unittest unionfind_test.py -v
