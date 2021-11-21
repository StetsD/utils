package utils

func Contain(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}

	}

	return false
}

func ContainInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func Check() string {
	return "check"
}
