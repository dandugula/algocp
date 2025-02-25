package main


func maxDiffPriorElem(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	minElemSoFar := arr[0]
	maxDiffSoFar, currDiff := -1, -1

	for _, v := range arr[1:] {

		if v < minElemSoFar {
			minElemSoFar = v
		} else {
			currDiff = v - minElemSoFar
		}

		if maxDiffSoFar < currDiff {
			maxDiffSoFar = currDiff
		}
	}

	return maxDiffSoFar
}
