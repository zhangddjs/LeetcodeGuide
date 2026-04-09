.PHONY: gen clean test

# Generate today's practice files
gen:
	@cd cmd && python3 daily_gen.py

# Archive generated files
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "traverse_*.py" ! -name "traverse_template.tmpl" ! -name "traverse_test.py" -exec mv {} archived/ \;
	@find . -maxdepth 1 -name "dag_*.py" ! -name "dag_template.tmpl" ! -name "dag_test.py" -exec mv {} archived/ \;
	@find . -maxdepth 1 -name "dijkstra_*.py" ! -name "dijkstra_template.tmpl" ! -name "dijkstra_test.py" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"

# Run tests
test:
	@python3 -m pytest traverse_test.py dag_test.py dijkstra_test.py -v || python3 -m unittest discover -p "*_test.py" -v
