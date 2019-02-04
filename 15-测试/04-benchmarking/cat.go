package benchmarking

func CatString(strlist []string, sep string) string {
	if len(strlist) == 0 {
		return ""
	} else if len(strlist) == 1 {
		return strlist[0]
	}
	result := strlist[0]
	for _, str := range strlist[1:] {
		result = result + sep + str
	}
	return result
}
