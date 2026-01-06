.PHONY: daily generate clean help

# Generate today's sorting practice file
gen:
	@go run cmd/daily_gen.go

# Archive generated files (keeps template and generators)
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "traverse_*.go" -not -name "traverse_template.tmpl" -not -name "traverse_test.go" -exec mv {} archived/ \;
	@echo "📁 Moved generated traverse files to archived/"

# Show help
help:
	@echo "Available commands:"
	@echo "  make daily     - Generate today's sorting practice file"
	@echo "  make generate  - Use go generate to create file"
	@echo "  make clean     - Archive all generated traverse_*.go files to archived/"
	@echo "  make help      - Show this help message"
