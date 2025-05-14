package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// customSorting sıralama funksiyası
func customSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		lenI := len(strArr[i])
		lenJ := len(strArr[j])

		// Tək və cüt uzunluqlu sətirlər üçün
		iIsOdd := lenI%2 == 1
		jIsOdd := lenJ%2 == 1

		// Qayda 1: Tək uzunluqlu sətir cüt uzunluqludan əvvəl gəlir
		if iIsOdd && !jIsOdd {
			return true
		}
		if !iIsOdd && jIsOdd {
			return false
		}

		// Qayda 2: Əgər hər ikisi tək uzunluqludursa, qısa olan əvvəl gəlir
		if iIsOdd && jIsOdd {
			if lenI != lenJ {
				return lenI < lenJ
			}
		}

		// Qayda 3: Əgər hər ikisi cüt uzunluqludursa, uzun olan əvvəl gəlir
		if !iIsOdd && !jIsOdd {
			if lenI != lenJ {
				return lenI > lenJ
			}
		}

		// Qayda 4: Əgər uzunluqlar eynidirsə, əlifba sırası ilə
		return strArr[i] < strArr[j]
	})

	return strArr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Sətirlərin sayını oxuyuruq
	var n int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// Sətirləri oxuyuruq
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		strArr[i] = scanner.Text()
	}

	// Sıralayırıq
	sortedArr := customSorting(strArr)

	// Nəticələri çap edirik
	for _, s := range sortedArr {
		fmt.Println(s)
	}
}
