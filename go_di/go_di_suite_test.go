package go_di_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoDi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoDi Suite")
}
