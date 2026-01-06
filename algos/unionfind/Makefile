.PHONY: daily generate clean help

# Generate today's sorting practice file
gen:
	@go run cmd/daily_gen.go

# Archive generated files (keeps template and generators)
clean:
	@mkdir -p archived/
	@find . -maxdepth 1 -name "unionfind_*.go" -not -name "unionfind_template.tmpl" -not -name "unionfind_test.go" -exec mv {} archived/ \;
	@echo "📁 Moved generated unionfind files to archived/"

# Show help
help:
	@echo "Available commands:"
	@echo "  make daily     - Generate today's sorting practice file"
	@echo "  make generate  - Use go generate to create file"
	@echo "  make clean     - Archive all generated unionfind_*.go files to archived/"
	@echo "  make help      - Show this help message"
