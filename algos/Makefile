.PHONY: clean gen

# Generate today's practice files for all modules
gen:
	@cd unionfind && go run cmd/daily_gen.go
	@cd trie && go run cmd/daily_gen.go
	@cd sortalgo && go run cmd/daily_gen.go
	@cd graph && go run cmd/daily_gen.go

# Archive generated files for all modules
clean:
	@cd unionfind && mkdir -p archived/ && find . -maxdepth 1 -name "unionfind_*.go" -not -name "unionfind_template.tmpl" -not -name "unionfind_test.go" -exec mv {} archived/ \;
	@cd trie && mkdir -p archived/ && find . -maxdepth 1 -name "trie_*.go" -not -name "trie_template.tmpl" -not -name "trie_test.go" -exec mv {} archived/ \;
	@cd sortalgo && mkdir -p archived/ && find . -maxdepth 1 -name "sortalgo_*.go" -not -name "sortalgo_template.tmpl" -not -name "sortalgo_test.go" -exec mv {} archived/ \;
	@cd graph && mkdir -p archived/ && find . -maxdepth 1 -name "traverse_*.go" -not -name "traverse_template.tmpl" -not -name "traverse_test.go" -exec mv {} archived/ \;
	@cd graph && find . -maxdepth 1 -name "dag_*.go" -not -name "dag_template.tmpl" -not -name "dag_test.go" -exec mv {} archived/ \;
	@cd graph && find . -maxdepth 1 -name "dijkstra_*.go" -not -name "dijkstra_template.tmpl" -not -name "dijkstra_test.go" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"
