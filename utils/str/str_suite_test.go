package str_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "String Suite")
}
