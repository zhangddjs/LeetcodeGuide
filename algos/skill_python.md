# Python 技巧

## 统计list里每个元素的数量

Counter

```py
nums: List[int]
for v, c in Counter(nums).items()
```

## 任一

any

```py
any(c in ps or all(c % p for p in ps) for v, c in Counter(nums).items() if c > 1)
```


