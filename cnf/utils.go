package cnf

func absInt(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func maxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func minInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
