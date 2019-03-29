package main

func isPermutation(text1 string, text2 string) bool {
	// 대소문자 구별, 공백 처리 확인
	// 책에 오류가 있는 것 같다. 282p line -4, 풀이 #2
	if len(text1) != len(text2) {
		return false
	}

	var arr1 [128]int
	var arr2 [128]int

	for _, char := range text1 {
		arr1[int(char)] += 1
	}
	for _, char := range text2 {
		arr2[int(char)] += 1
	}

	if arr1 != arr2 {
		return false
	}
	return true
}

/*
func main() {
	text1 := "abccc"
	text2 := "cecba"

	result := isPermutation(text1, text2)
	fmt.Println(result)
}*/
