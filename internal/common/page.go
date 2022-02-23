package common

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

// RemainingPages returns the number of remaining pages to collect.
func RemainingPages(count, size, page, end int) int {
	n := count / size
	if count%size != 0 {
		n++
	}

	return min(n, end) - page + 1
}
