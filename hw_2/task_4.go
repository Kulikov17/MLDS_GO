package main

import (
	"fmt"
	"sort"
)

func hasInTree(tree map[string]string, key string) bool {
	_, ok := tree[key]

	if ok {
		return true
	}

	return false
}

func main() {

	var n int
	var parent, child string

	fmt.Scanf("%d\n", &n)

	var tree = make(map[string]string)
	var heights = make(map[string]int)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &parent, &child)

		tree[parent] = child

		heights[parent] = 0
		heights[child] = 0
	}

	for key, _ := range tree {

		tmp := key

		for hasInTree(tree, tmp) {
			tmp = tree[tmp]
			heights[key] += 1
		}
	}

	keys := make([]string, 0, len(heights))
	for k := range heights {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, heights[k])
	}
}
