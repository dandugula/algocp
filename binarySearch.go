package main

func bs_list (haystack []int, needle int) bool {
    lo := 0
    hi := len(haystack)
    for lo < hi {
        m := lo + (hi - lo)/2
        v := haystack[m]
        if v == needle {
            return true
        } else if needle > v {
            lo = m + 1
        } else {
            hi = m 
        }
    }
    return false
}
