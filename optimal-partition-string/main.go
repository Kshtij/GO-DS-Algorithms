package main

func main() {
	partitionString("abacaba")
}

func partitionString(s string) int {
	retList := make([]string, 0, 1)
	var tempStr string
	length := len(s)

	i := 0
	charSet := make(map[byte]interface{})
	for i < length {
		_, ok := charSet[s[i]]
		if ok {
			retList = append(retList, tempStr)
			tempStr = "" + string(s[i])
			charSet = make(map[byte]interface{})
			charSet[s[i]] = nil
		} else {
			charSet[s[i]] = nil
			tempStr = tempStr + string(s[i])
		}

		i++
	}
	retList = append(retList, tempStr)
	return len(retList)

}
