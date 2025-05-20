package main

import "fmt"

func main() {
	nums1 := []int{2, 2, 2}
	nums2 := []int{2, 2}
	res := make([]int, 0)

	seen := make(map[int]bool)

	for _, v := range nums1 {
		seen[v] = true
	}

	for _, v := range nums2 {
		if ok := seen[v]; ok {
			delete(seen, v)
			res = append(res, v)
		}
	}

	fmt.Println(res)

}
