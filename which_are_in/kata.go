package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	result := []string{}
	for _, small := range array1 {
		found := false
		for _, big := range array2 {
			if strings.Contains(big, small) {
				found = true
			}
		}
		if found && !contains(result, small) {
			result = append(result, small)
		}
	}
	sort.Strings(result)
	return result
}

func contains(result []string, small string) bool {
	for _, s := range result {
		if s == small {
			return true
		}
	}
	return false
}
