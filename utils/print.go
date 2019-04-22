package utils

import (
	"errors"
	"github.com/go-test/deep"
	"io/ioutil"
	"os"
	"strings"
)

// PrintTest 함수는 인자로 전달하는 함수 내부에서 표준 출력으로 출력한 문자열과 두 번째 인자로 전달한 문자열을 비교한다.
// 만약 그 내용이 다르면 오류를 반환한다.
func PrintTest(f func(), expected []string) error {
	r, w, _ := os.Pipe()
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
		r.Close()
	}()
	os.Stdout = w
	go func() {
		defer w.Close()
		f()
	}()
	stdout, _ := ioutil.ReadAll(r)

	actual := strings.Split(string(stdout), "\n")
	if actual[len(actual)-1] == "" {
		actual = actual[:len(actual)-1]
	}
	diff := deep.Equal(actual, expected)
	if diff != nil {
		return errors.New(strings.Join(diff, "\n"))
	}
	return nil
}
