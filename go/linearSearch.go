package main

func linear_search(haystack []int, needle int) bool {
    for _, v := range haystack {
        if needle == v {
            return true
        }
	}
	return false
}
