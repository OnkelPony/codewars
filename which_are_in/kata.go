package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	var result []string
	found := false
	for _, small := range array1 {
		for _, big := range array2 {
			if strings.Contains(big, small) {
				found = true
			}
		}
		if found {
			result = append(result, small)
		}
	}
	sort.Strings(result)
	return result
}
