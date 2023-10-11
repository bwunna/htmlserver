package SupportingFunctions

import "fmt"

func KeysInString(keys []string) string {
	var answer string
	for _, value := range keys {
		answer += "'" + value + "', "
	}
	if len(answer) > 2 {
		answer = answer[:len(answer)-2]
	}
	fmt.Println(answer)
	return answer
}
