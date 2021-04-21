package reverse

import "fmt"

func Reverse(word string) string {
	var drow string
	fmt.Println(len(word))
	for i , _ := range word {
		i = len(word) - 1 - i
		drow += string(word[i])

	}
	fmt.Println(drow)
	return drow
}