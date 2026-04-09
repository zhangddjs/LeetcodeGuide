.PHONY: gen clean test

# Generate today's practice file
gen:
	@cd cmd && python3 daily_gen.py

# Archive generated files
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "sortalgo_*.py" ! -name "sortalgo_template.tmpl" ! -name "sortalgo_test.py" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"

# Run tests
test:
	@python3 -m pytest sortalgo_test.py -v || python3 -m unittest sortalgo_test.py -v
