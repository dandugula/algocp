package main

import "math"

/*
func two_crystal_balls(breaks []bool) int {
    n_sqrt := int(math.Sqrt(float64(len(breaks))))
    prev := 0
    i := 1
    j := i * n_sqrt
    for j < len(breaks) {
        if breaks[j] {
            break
        }
        prev = j + 1
        i++
        j = i * n_sqrt
    }
    for prev < j {
        if breaks [prev] {
            return prev
        }
        prev++
    }
    return -1
}
*/

func two_crystal_balls(breaks [1000]bool) int {

	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount

	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jmpAmount
	}

	i -= jmpAmount

	for j := 0; j < jmpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1

}
