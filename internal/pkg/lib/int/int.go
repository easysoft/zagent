package _intUtils

func FindInArr(val int, arr []int) bool {
	for _, i := range arr {
		if val == i {
			return true
		}
	}

	return false
}
