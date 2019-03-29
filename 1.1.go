package main

func findOverlap(text string) bool {
	if len(text) > 128 {
		return true
	}
	var arr [128]bool
	for _, char := range text {
		if arr[int(char)] == true {
			return true
		}
		arr[int(char)] = true
	}
	return false
}

/*
func main() {
	fmt.Println("vim-go")
	result := findOverlap("abcde")
	fmt.Println(result)
}
*/
