# Python 版本算法默写框架

这是一套完全对应 Golang 版本的 Python 算法默写框架。

## 目录结构

```
algos/
├── Makefile                 # 主 Makefile (Python 版本)
├── unionfind_py/            # 并查集模块
│   ├── cmd/daily_gen.py    # 生成器脚本
│   ├── unionfind_template.tmpl  # 模板文件
│   ├── unionfind_test.py   # 测试文件
│   ├── archived/           # 归档目录
│   └── Makefile
├── trie_py/                # 字典树模块
│   ├── cmd/daily_gen.py
│   ├── trie_template.tmpl
│   ├── trie_test.py
│   ├── archived/
│   └── Makefile
├── sortalgo_py/            # 排序算法模块
│   ├── cmd/daily_gen.py
│   ├── sortalgo_template.tmpl
│   ├── sortalgo_test.py
│   ├── archived/
│   └── Makefile
└── graph_py/               # 图算法模块
    ├── cmd/daily_gen.py
    ├── traverse_template.tmpl
    ├── dag_template.tmpl
    ├── dijkstra_template.tmpl
    ├── traverse_test.py
    ├── dag_test.py
    ├── dijkstra_test.py
    ├── archived/
    └── Makefile
```

## 使用方法

### 生成今天的练习文件

```bash
# 生成所有模块的练习文件
make gen

# 或者单独生成某个模块
cd unionfind_py && make gen
cd trie_py && make gen
cd sortalgo_py && make gen
cd graph_py && make gen
```

### 归档旧的练习文件

```bash
# 归档所有模块的练习文件
make clean

# 或者单独归档某个模块
cd unionfind_py && make clean
cd trie_py && make clean
cd sortalgo_py && make clean
cd graph_py && make clean
```

### 运行测试

```bash
# 运行单个模块的测试
cd unionfind_py && make test
cd trie_py && make test
cd sortalgo_py && make test
cd graph_py && make test
```

## 模块说明

### 1. UnionFind (并查集)
- `find(x)` - 查找元素 x 所属的集合
- `union(x, y)` - 合并 x 和 y 所在的集合
- `connected(x, y)` - 判断 x 和 y 是否在同一集合
- `count()` - 返回集合的数量

### 2. Trie (字典树)
- `insert(word)` - 插入单词
- `search(word)` - 搜索完整单词
- `starts_with(prefix)` - 判断是否存在以 prefix 开头的单词

### 3. SortAlgo (排序算法)
- `heapsort(arr)` - 堆排序
- `quicksort(arr)` - 快速排序
- `mergesort(arr)` - 归并排序
- `insertsort(arr)` - 插入排序
- `bubblesort(arr)` - 冒泡排序

### 4. Graph (图算法)

#### Traverse (图遍历)
- `add_edge(from_node, to_node)` - 添加边
- `dfs(start)` - 从起点开始深度优先搜索
- `dfs_whole()` - 遍历整个图（包括不连通的部分）
- `has_path(from_node, to_node)` - 判断是否存在路径
- `shortest_path(from_node, to_node)` - 查找最短路径（BFS）

#### DAG (有向无环图)
- `add_edge(from_node, to_node)` - 添加边
- `add_weighted_edge(from_node, to_node, weight)` - 添加带权重的边
- `topological_sort()` - 拓扑排序（DFS）
- `topological_sort_bfs()` - 拓扑排序（BFS/Kahn算法）
- `is_dag()` - 判断是否为 DAG
- `shortest_path(start)` - DAG 最短路径

#### Dijkstra (最短路径)
- `add_edge(from_node, to_node)` - 添加边
- `add_weighted_edge(from_node, to_node, weight)` - 添加带权重的边
- `shortest_path(start)` - 计算从起点到所有点的最短路径
- `shortest_path_between(from_node, to_node)` - 计算两点间最短路径

## 工作流程

1. **每天练习前**：运行 `make gen` 生成当天的练习文件
2. **完成实现**：在生成的文件中实现算法
3. **运行测试**：使用 `make test` 验证实现是否正确
4. **定期归档**：使用 `make clean` 将完成的练习文件归档

## 测试说明

测试文件会自动加载最新生成的练习文件。每个模块都包含完整的测试用例，包括：
- 边界情况测试
- 基本功能测试
- 复杂场景测试
- 性能测试（大数据量）

## 注意事项

1. 生成的文件名格式为 `{module}_YYYYMMDD.py`，例如 `unionfind_20260409.py`
2. 同一天只能生成一次，重复运行会提示文件已存在
3. 归档操作会将生成的练习文件移动到 `archived/` 目录
4. 模板文件（`*_template.tmpl`）和测试文件（`*_test.py`）不会被归档
5. 所有排序算法都是原地排序（in-place），会修改原数组

## 与 Golang 版本的对应关系

| Golang | Python | 说明 |
|--------|--------|------|
| `Makefile` (Golang 版本) | `Makefile` (Python 版本) | 主 Makefile |
| `unionfind/` | `unionfind_py/` | 并查集模块 |
| `trie/` | `trie_py/` | 字典树模块 |
| `sortalgo/` | `sortalgo_py/` | 排序算法模块 |
| `graph/` | `graph_py/` | 图算法模块 |
| `*.go` | `*.py` | 代码文件 |
| `go run cmd/daily_gen.go` | `python3 cmd/daily_gen.py` | 生成命令 |

函数命名遵循 Python 风格（snake_case），例如：
- `AddEdge` → `add_edge`
- `StartsWith` → `starts_with`
- `TopologicalSort` → `topological_sort`
