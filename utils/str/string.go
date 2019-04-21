package str

import (
	"sort"
	"strings"
)

type sortable []byte

func (s sortable) Len() int           { return len(s) }
func (s sortable) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Sort 함수는 주어진 문자열을 알파벳 순서대로 재배열하여 반환한다.
func Sort(s string) string {
	b := sortable(s)
	sort.Sort(b)
	return string(b)
}

// SortAdapter 타입은 문자열 슬라이스를 정렬하기 위한 어댑터다.
type SortAdapter []string

func (a SortAdapter) Len() int           { return len(a) }
func (a SortAdapter) Less(i, j int) bool { return strings.Compare(a[i], a[j]) < 0 }
func (a SortAdapter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Reverse 함수는 주어진 문자열을 뒤집어 반환한다.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
