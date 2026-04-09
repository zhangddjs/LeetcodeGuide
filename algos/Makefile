.PHONY: clean gen

# Generate today's practice files for all modules (Python version)
gen:
	@cd unionfind_py && cd cmd && python3 daily_gen.py
	@cd trie_py && cd cmd && python3 daily_gen.py
	@cd sortalgo_py && cd cmd && python3 daily_gen.py
	@cd graph_py && cd cmd && python3 daily_gen.py

# Archive generated files for all modules (Python version)
clean:
	@cd unionfind_py && mkdir -p archived/ && find . -maxdepth 1 -name "unionfind_*.py" ! -name "unionfind_template.tmpl" ! -name "unionfind_test.py" -exec mv {} archived/ \;
	@cd trie_py && mkdir -p archived/ && find . -maxdepth 1 -name "trie_*.py" ! -name "trie_template.tmpl" ! -name "trie_test.py" -exec mv {} archived/ \;
	@cd sortalgo_py && mkdir -p archived/ && find . -maxdepth 1 -name "sortalgo_*.py" ! -name "sortalgo_template.tmpl" ! -name "sortalgo_test.py" -exec mv {} archived/ \;
	@cd graph_py && mkdir -p archived/ && find . -maxdepth 1 -name "traverse_*.py" ! -name "traverse_template.tmpl" ! -name "traverse_test.py" -exec mv {} archived/ \;
	@cd graph_py && find . -maxdepth 1 -name "dag_*.py" ! -name "dag_template.tmpl" ! -name "dag_test.py" -exec mv {} archived/ \;
	@cd graph_py && find . -maxdepth 1 -name "dijkstra_*.py" ! -name "dijkstra_template.tmpl" ! -name "dijkstra_test.py" -exec mv {} archived/ \;
	@echo "📁 Moved all generated files to archived/"

# ============================================================
# OLD GOLANG VERSION (COMMENTED OUT)
# ============================================================
# gen:
#	@cd unionfind && go run cmd/daily_gen.go
#	@cd trie && go run cmd/daily_gen.go
#	@cd sortalgo && go run cmd/daily_gen.go
#	@cd graph && go run cmd/daily_gen.go

# clean:
#	@cd unionfind && mkdir -p archived/ && find . -maxdepth 1 -name "unionfind_*.go" -not -name "unionfind_template.tmpl" -not -name "unionfind_test.go" -exec mv {} archived/ \;
#	@cd trie && mkdir -p archived/ && find . -maxdepth 1 -name "trie_*.go" -not -name "trie_template.tmpl" -not -name "trie_test.go" -exec mv {} archived/ \;
#	@cd sortalgo && mkdir -p archived/ && find . -maxdepth 1 -name "sortalgo_*.go" -not -name "sortalgo_template.tmpl" -not -name "sortalgo_test.go" -exec mv {} archived/ \;
#	@cd graph && mkdir -p archived/ && find . -maxdepth 1 -name "traverse_*.go" -not -name "traverse_template.tmpl" -not -name "traverse_test.go" -exec mv {} archived/ \;
#	@cd graph && find . -maxdepth 1 -name "dag_*.go" -not -name "dag_template.tmpl" -not -name "dag_test.go" -exec mv {} archived/ \;
#	@cd graph && find . -maxdepth 1 -name "dijkstra_*.go" -not -name "dijkstra_template.tmpl" -not -name "dijkstra_test.go" -exec mv {} archived/ \;
#	@echo "📁 Moved all generated files to archived/"
