package main

// 첫 실수, 문자열이 연속으로 중복될 경우를 카운트하는 것인데 전체 문자의 카운트를 해버림. 문제 꼼꼼히 읽자ㅠ

// 이 알고리즘은 문자열을 합치는데 O(n^2)의 시간이 걸리기 때문에 StringBuilder로 수행 시간을 단축시켜야 한다.

import (
	"strconv"
)

func compress(text string) string {
	cnt := 1
	result := ""

	for idx, char := range text {
		if idx+1 < len(text) && string(text[idx+1]) == string(char) {
			cnt += 1
		} else {
			// 이 부분에서 내부적으로 새로운 string을 생성해 기존, 이후의 string을 복사해넣기때문에 속도가 느려짐.
			// ToDo: golang에서의 string add 동작 방법을 찾아보고 싶음.
			result += string(char) + strconv.Itoa(cnt)
			cnt = 1
		}
	}

	if len(result) > len(text) {
		return text
	}
	return result
}

/*
func main() {
	text := "aabccccaaa"
	result := compress(text)
	fmt.Println(result)
}
*/
