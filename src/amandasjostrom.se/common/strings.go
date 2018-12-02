package common
func NumberOfOccurences(find string, list []string) (sum int) {
	for _, value := range list {
		if value == find {
			sum = sum + 1
		}
	}
	return
}