package integer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInteger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integer Suite")
}
