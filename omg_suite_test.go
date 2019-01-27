package omg_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOmg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Omg Suite")
}
