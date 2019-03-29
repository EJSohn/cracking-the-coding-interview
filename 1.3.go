package main

func replace(text *string) {
	result := ""
	for idx, val := range *text {
		if val != ' ' {
			result += string((*text)[idx])
		} else {
			result += "%20"
		}
	}

	*text = result
}

/*
func main() {
	text := "hello world golang"
	replace(&text)
	fmt.Println(text)
}
*/
