package test

import (
	"sort"
	"time"
)

func findHighAccessEmployees(access_times [][]string) []string {
	res := make([]string, 0)
	employee := make(map[string][]string)
	for _, access := range access_times {
		employee[access[0]] = append(employee[access[0]], access[1])
	}
	for k, v := range employee {
		if len(v) < 3 {
			continue
		}
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		cnt := 1
		first, _ := time.Parse("1504", v[0])
		for i := 1; i < len(v); i++ {
			t, _ := time.Parse("1504", v[i])
			if t.Sub(first).Hours() >= 1 {
				if cnt == 2 {
					p, _ := time.Parse("1504", v[i-1])
					first = p
					i--
				} else {
					first = t
				}
				cnt = 1
				continue
			}
			cnt++
			if cnt == 3 {
				res = append(res, k)
				break
			}
		}
	}
	return res
}
