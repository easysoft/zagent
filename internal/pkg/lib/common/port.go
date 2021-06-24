package _commonUtils

func GetValidPort(start int, end int, existPortMap *map[int]bool) (ret int) {
	newPort := 0

	for i := 0; i < 99; i++ {
		port := start + i
		if port > end {
			break
		}

		if _, ok := (*existPortMap)[port]; !ok {
			newPort = port
			break
		}
	}

	if newPort > 0 {
		ret = newPort
		(*existPortMap)[newPort] = true
	}
	return
}
