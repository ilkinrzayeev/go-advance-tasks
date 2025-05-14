package main

import (
	"fmt"
	"sort"
)

func RemainderSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		remI := len(strArr[i]) % 3
		remJ := len(strArr[j]) % 3
		if remI != remJ {
			return remI < remJ
		}
		return strArr[i] < strArr[j]
	})
	return strArr
}

func main() {
	var n int
	fmt.Scan(&n)

	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strArr[i])
	}

	sorted := RemainderSorting(strArr)

	for _, str := range sorted {
		fmt.Println(str)
	}
}
